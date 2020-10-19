package expanders

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/lawrencegripper/azbrowse/internal/pkg/tracing"
	"github.com/lawrencegripper/azbrowse/pkg/armclient"
)

// NewStorageBlobExpander creates a new instance of StorageBlobExpander
func NewStorageBlobExpander(armclient *armclient.Client) *StorageBlobExpander {
	return &StorageBlobExpander{
		client:    &http.Client{},
		armClient: armclient,
	}
}

// Check interface
var _ Expander = &StorageBlobExpander{}

func (e *StorageBlobExpander) setClient(c *armclient.Client) {
	e.armClient = c
}

// StorageBlobExpander expands the blob  data-plane aspects of a Storage Account
type StorageBlobExpander struct {
	ExpanderBase
	client    *http.Client
	armClient *armclient.Client
}

// Name returns the name of the expander
func (e *StorageBlobExpander) Name() string {
	return "StorageBlobExpander"
}

// DoesExpand checks if this is a storage account
func (e *StorageBlobExpander) DoesExpand(ctx context.Context, currentItem *TreeNode) (bool, error) {
	swaggerResourceType := currentItem.SwaggerResourceType
	if currentItem.ItemType == SubResourceType && swaggerResourceType != nil {
		if swaggerResourceType.Endpoint.TemplateURL == "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}" {
			return true, nil
		}
	}
	if currentItem.Namespace == "storageBlob" {
		return true, nil
	}
	return false, nil
}

// Expand returns blobs in the StorageAccount congtainer
func (e *StorageBlobExpander) Expand(ctx context.Context, currentItem *TreeNode) ExpanderResult {

	swaggerResourceType := currentItem.SwaggerResourceType
	if currentItem.Namespace != "storageBlob" &&
		swaggerResourceType != nil &&
		swaggerResourceType.Endpoint.TemplateURL == "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}" {
		newItems := []*TreeNode{}
		newItems = append(newItems, &TreeNode{
			Parentid:  currentItem.ID,
			ID:        currentItem.ID + "/<repositories>",
			Namespace: "storageBlob",
			Name:      "Blobs",
			Display:   "Blobs",
			ItemType:  SubResourceType,
			ExpandURL: ExpandURLNotSupported,
			Metadata: map[string]string{
				"ContainerID":           currentItem.ExpandURL, // save resourceID of blob
				"SuppressSwaggerExpand": "true",
				"SuppressGenericExpand": "true",
			},
		})

		return ExpanderResult{
			Err:               nil,
			Response:          ExpanderResponse{Response: ""}, // Swagger expander will supply the response
			SourceDescription: "StorageBlobExpander request",
			Nodes:             newItems,
			IsPrimaryResponse: false,
		}
	}

	if currentItem.Namespace == "storageBlob" && currentItem.ItemType == SubResourceType {
		return e.expandBlobs(ctx, currentItem)
	}
	// if currentItem.Namespace == "containerRegistry" && currentItem.ItemType == SubResourceType {
	// 	return e.expandRepositories(ctx, currentItem)
	// } else if currentItem.ItemType == "containerRegistry.repository" {
	// 	return e.expandRepository(ctx, currentItem)
	// } else if currentItem.ItemType == "containerRegistry.repository.tags" {
	// 	return e.expandRepositoryTags(ctx, currentItem)
	// } else if currentItem.ItemType == "containerRegistry.repository.tag" {
	// 	return e.expandRepositoryTag(ctx, currentItem)
	// } else if currentItem.ItemType == "containerRegistry.repository.manifests" {
	// 	return e.expandRepositoryManifests(ctx, currentItem)
	// } else if currentItem.ItemType == "containerRegistry.repository.manifest" {
	// 	return e.expandRepositoryManifest(ctx, currentItem)
	// }

	return ExpanderResult{
		Err:               fmt.Errorf("Error - unhandled Expand"),
		Response:          ExpanderResponse{Response: "Error!"},
		SourceDescription: "ContainerRegistryExpander request",
	}
}

// Delete attempts to delete the item. Returns true if deleted, false if not handled, an error if an error occurred attempting to delete
func (e *StorageBlobExpander) Delete(ctx context.Context, currentItem *TreeNode) (bool, error) {
	// TODO - implement
	// switch currentItem.ItemType {
	// case "containerRegistry.repository":
	// 	return e.deleteRepository(ctx, currentItem)
	// case "containerRegistry.repository.tag":
	// 	return e.deleteRepositoryTag(ctx, currentItem)
	// case "containerRegistry.repository.manifest":
	// 	return e.deleteRepositoryManifest(ctx, currentItem)
	// }
	return false, nil
}

func (e *StorageBlobExpander) expandBlobs(ctx context.Context, currentItem *TreeNode) ExpanderResult {

	// TODO cache this in the metadata
	containerID := currentItem.Metadata["ContainerID"]
	containerName := e.getContainerName(containerID)
	accountName := e.getAccountName(containerID)
	accountKey, err := e.getAccountKey(ctx, containerID)
	if err != nil {
		err = fmt.Errorf("Error getting account key: %s", err)
		return ExpanderResult{
			Err:               err,
			SourceDescription: "StorageBlobExpander request",
		}
	}
	blobEndpoint, err := e.getStorageBlobEndpoint(ctx, containerID)
	if err != nil {
		err = fmt.Errorf("Error getting blob endpoint: %s", err)
		return ExpanderResult{
			Err:               err,
			SourceDescription: "StorageBlobExpander request",
		}
	}

	// ListBlob docs: https://docs.microsoft.com/en-us/rest/api/storageservices/list-blobs
	url := blobEndpoint + containerName + "?restype=container&comp=list"
	buf, err := e.doRequest(ctx, "GET", url, accountName, accountKey, "/"+accountName+"/"+containerName)

	if err != nil {
		return ExpanderResult{
			Err:               fmt.Errorf("Error listing blobs: %s", err),
			SourceDescription: "StorageBlobExpander request",
		}
	}
	result := string(buf)
	return ExpanderResult{
		Response:          ExpanderResponse{Response: result, ResponseType: ResponseXML},
		SourceDescription: "StorageBlobExpander request",
		Nodes:             []*TreeNode{},
		IsPrimaryResponse: true,
	}
}

// StorageListKeyResponse is used to unmarshal a call to listKeys on a storage account
type StorageListKeyResponse struct {
	Keys []struct {
		KeyName     string `json:"keyName"`
		Value       string `json:"value"`
		Permissions string `json:"permissions"`
	} `json:"keys"`
}

// StorageAccountResponse is a partial representation of the storage account response
type StorageAccountResponse struct {
	Properties struct {
		PrimaryEndpoints struct {
			Blob string `json:"blob"`
			Dfs  string `json:"dfs"`
		} `json:"primaryEndpoints`
	} `json: "properties`
}

func (e *StorageBlobExpander) getAccountName(containerID string) string {
	i := strings.Index(containerID, "/blobServices")
	accountID := containerID[0:i]
	i = strings.LastIndex(accountID, "/")
	return accountID[i+1:]
}
func (e *StorageBlobExpander) getContainerName(containerID string) string {
	i := strings.LastIndex(containerID, "/")
	containerName := containerID[i+1:]
	i = strings.Index(containerName, "?")
	if i >= 0 {
		containerName = containerName[:i] // strip query string
	}
	return containerName
}

func (e *StorageBlobExpander) getAccountKey(ctx context.Context, containerID string) (string, error) {

	i := strings.Index(containerID, "/blobServices")
	rootURL := containerID[0:i]
	listKeysURL := rootURL + "/listKeys?api-version=2019-06-01"

	data, err := e.armClient.DoRequest(ctx, "POST", listKeysURL)
	if err != nil {
		return "", fmt.Errorf("Error calling listKeys: %s", err)
	}
	response := StorageListKeyResponse{}
	err = json.Unmarshal([]byte(data), &response)
	if err != nil {
		err = fmt.Errorf("Error unmarshalling response: %s\nURL:%s", err, listKeysURL)
		return "", err
	}
	if len(response.Keys) == 0 {
		err = fmt.Errorf("No keys in response: %s", err)
		return "", err
	}

	return response.Keys[0].Value, nil
}
func (e *StorageBlobExpander) getStorageBlobEndpoint(ctx context.Context, containerID string) (string, error) {

	i := strings.Index(containerID, "/blobServices")
	rootURL := containerID[0:i]
	storageAccountURL := rootURL + "?api-version=2019-06-01"

	data, err := e.armClient.DoRequest(ctx, "GET", storageAccountURL)
	if err != nil {
		return "", fmt.Errorf("Error getting storage account: %s", err)
	}
	response := StorageAccountResponse{}
	err = json.Unmarshal([]byte(data), &response)
	if err != nil {
		err = fmt.Errorf("Error unmarshalling response: %s\nURL:%s", err, storageAccountURL)
		return "", err
	}

	return response.Properties.PrimaryEndpoints.Blob, nil
}

func (e *StorageBlobExpander) doRequest(ctx context.Context, verb string, url string, accountName string, accountKey string, accountAndPath string) ([]byte, error) {
	span, _ := tracing.StartSpanFromContext(ctx, "doRequest(blobexpnder):"+url, tracing.SetTag("url", url))
	defer span.Finish()

	req, err := http.NewRequest(verb, url, nil)
	if err != nil {
		return []byte{}, fmt.Errorf("Failed to create request: %s", err)
	}
	if req.Header.Get("x-ms-version") == "" {
		req.Header.Set("x-ms-version", "2018-03-28")
	}
	dateString := time.Now().UTC().Format(http.TimeFormat)
	req.Header.Set("x-ms-date", dateString)
	// signature := e.getAuthSignature(req, accountKey, accountAndPath)
	// req.Header.Set("Authorization", fmt.Sprintf("SharedKeyLite %s:%s", accountName, signature))
	err = e.addAuthHeader(req, accountName, accountKey)
	if err != nil {
		return []byte{}, fmt.Errorf("Failed to add auth header: %s", err)
	}

	response, err := e.client.Do(req.WithContext(ctx))
	if err != nil {
		return []byte{}, fmt.Errorf("Request failed: %s", err)
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		defer response.Body.Close() //nolint: errcheck
		buftemp, _ := ioutil.ReadAll(response.Body)
		foo := string(buftemp)
		_ = foo
		return []byte{}, fmt.Errorf("DoRequest failed %v for '%s'", response.Status, url)
	}

	defer response.Body.Close() //nolint: errcheck
	buf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("Failed to read body: %s", err)
	}

	return buf, nil
}

// func (e *StorageBlobExpander) getAuthSignature(req *http.Request, accountKey string, accountAndPath string) string {
// 	verb := req.Method
// 	dateString := req.Header.Get("x-ms-date")

// 	// stringToSign := verb + "\n\n\n\n\n\n\n\n\n\n\n\nx-ms-date:" + dateString + "\nx-ms-version:2019-06-01\n" + accountAndPath
// 	// Based on https://docs.microsoft.com/en-us/rest/api/storageservices/authorize-with-shared-key#blob-queue-and-file-services-shared-key-lite-authorization
// 	stringToSign := verb +
// 		"\n" + // Content-MD5
// 		"\n" + // Content-Type
// 		"\nx-ms-date:" + dateString +
// 		e.getCanonicalisedHeaderString(req) + // Canonicalized Headers
// 		accountAndPath

// 	h := hmac.New(sha256.New, []byte(accountKey))
// 	h.Write([]byte(stringToSign))
// 	return base64.StdEncoding.EncodeToString(h.Sum(nil))
// }
// func (e *StorageBlobExpander) getCanonicalisedHeaderString(req *http.Request) string {

// 	headers := map[string]string{}

// 	for k, v := range req.Header {
// 		headerName := strings.TrimSpace(strings.ToLower(k))
// 		if strings.HasPrefix(headerName, "x-ms-") {
// 			headers[headerName] = v[0] // Think this is safe as the x-ms-* headers only seem to have a single value
// 		}
// 	}

// 	if len(headers) == 0 {
// 		return ""
// 	}
// 	keys := []string{}
// 	for key := range headers {
// 		keys = append(keys, key)
// 	}

// 	sort.Strings(keys)

// 	result := ""
// 	separator := ""
// 	for _, key := range keys {
// 		result += separator + key + ":" + headers[key]
// 		separator = "\n"
// 	}

// 	return result
// }

func (e *StorageBlobExpander) testCases() (bool, *[]expanderTestCase) {
	return false, nil
}

// Auth helper code based on https://github.com/Azure/azure-storage-blob-go
// (https://github.com/Azure/azure-storage-blob-go/blob/3efca72bd11c050222deab57e25ea90df03b9692/azblob/zc_credential_shared_key.go#L55)
func (e *StorageBlobExpander) addAuthHeader(request *http.Request, accountName string, accountKey string) error {

	// Add a x-ms-date header if it doesn't already exist
	if d := request.Header.Get(headerXmsDate); d == "" {
		request.Header[headerXmsDate] = []string{time.Now().UTC().Format(http.TimeFormat)}
	}
	stringToSign, err := e.buildStringToSign(request, accountName)
	if err != nil {
		return fmt.Errorf("Failed to build string to sign: %s", err)
		return err
	}
	signature, err := e.ComputeHMACSHA256(stringToSign, accountKey)
	if err != nil {
		return fmt.Errorf("Failed to compute signature: %s", err)
	}
	authHeader := strings.Join([]string{"SharedKey ", accountName, ":", signature}, "")
	request.Header[headerAuthorization] = []string{authHeader}
	return nil
}

// Constants ensuring that header names are correctly spelled and consistently cased.
const (
	headerAuthorization      = "Authorization"
	headerCacheControl       = "Cache-Control"
	headerContentEncoding    = "Content-Encoding"
	headerContentDisposition = "Content-Disposition"
	headerContentLanguage    = "Content-Language"
	headerContentLength      = "Content-Length"
	headerContentMD5         = "Content-MD5"
	headerContentType        = "Content-Type"
	headerDate               = "Date"
	headerIfMatch            = "If-Match"
	headerIfModifiedSince    = "If-Modified-Since"
	headerIfNoneMatch        = "If-None-Match"
	headerIfUnmodifiedSince  = "If-Unmodified-Since"
	headerRange              = "Range"
	headerUserAgent          = "User-Agent"
	headerXmsDate            = "x-ms-date"
	headerXmsVersion         = "x-ms-version"
)

// ComputeHMACSHA256 generates a hash signature for an HTTP request or for a SAS.
func (e *StorageBlobExpander) ComputeHMACSHA256(message string, accountKey string) (string, error) {
	bytes, err := base64.StdEncoding.DecodeString(accountKey)
	if err != nil {
		return "", fmt.Errorf("Failed to decode storage account key: %s", err)
	}
	h := hmac.New(sha256.New, bytes)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}

func (e *StorageBlobExpander) buildStringToSign(request *http.Request, accountName string) (string, error) {
	// https://docs.microsoft.com/en-us/rest/api/storageservices/authentication-for-the-azure-storage-services
	headers := request.Header
	contentLength := headers.Get(headerContentLength)
	if contentLength == "0" {
		contentLength = ""
	}

	canonicalizedResource, err := e.buildCanonicalizedResource(request.URL, accountName)
	if err != nil {
		return "", err
	}

	stringToSign := strings.Join([]string{
		request.Method,
		headers.Get(headerContentEncoding),
		headers.Get(headerContentLanguage),
		contentLength,
		headers.Get(headerContentMD5),
		headers.Get(headerContentType),
		"", // Empty date because x-ms-date is expected (as per web page above)
		headers.Get(headerIfModifiedSince),
		headers.Get(headerIfMatch),
		headers.Get(headerIfNoneMatch),
		headers.Get(headerIfUnmodifiedSince),
		headers.Get(headerRange),
		buildCanonicalizedHeader(headers),
		canonicalizedResource,
	}, "\n")
	return stringToSign, nil
}

func buildCanonicalizedHeader(headers http.Header) string {
	cm := map[string][]string{}
	for k, v := range headers {
		headerName := strings.TrimSpace(strings.ToLower(k))
		if strings.HasPrefix(headerName, "x-ms-") {
			cm[headerName] = v // NOTE: the value must not have any whitespace around it.
		}
	}
	if len(cm) == 0 {
		return ""
	}

	keys := make([]string, 0, len(cm))
	for key := range cm {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	ch := bytes.NewBufferString("")
	for i, key := range keys {
		if i > 0 {
			ch.WriteRune('\n')
		}
		ch.WriteString(key)
		ch.WriteRune(':')
		ch.WriteString(strings.Join(cm[key], ","))
	}
	return string(ch.Bytes())
}

func (e *StorageBlobExpander) buildCanonicalizedResource(u *url.URL, accountName string) (string, error) {
	// https://docs.microsoft.com/en-us/rest/api/storageservices/authentication-for-the-azure-storage-services
	cr := bytes.NewBufferString("/")
	cr.WriteString(accountName)

	if len(u.Path) > 0 {
		// Any portion of the CanonicalizedResource string that is derived from
		// the resource's URI should be encoded exactly as it is in the URI.
		// -- https://msdn.microsoft.com/en-gb/library/azure/dd179428.aspx
		cr.WriteString(u.EscapedPath())
	} else {
		// a slash is required to indicate the root path
		cr.WriteString("/")
	}

	// params is a map[string][]string; param name is key; params values is []string
	params, err := url.ParseQuery(u.RawQuery) // Returns URL decoded values
	if err != nil {
		return "", errors.New("parsing query parameters must succeed, otherwise there might be serious problems in the SDK/generated code")
	}

	if len(params) > 0 { // There is at least 1 query parameter
		paramNames := []string{} // We use this to sort the parameter key names
		for paramName := range params {
			paramNames = append(paramNames, paramName) // paramNames must be lowercase
		}
		sort.Strings(paramNames)

		for _, paramName := range paramNames {
			paramValues := params[paramName]
			sort.Strings(paramValues)

			// Join the sorted key values separated by ','
			// Then prepend "keyName:"; then add this string to the buffer
			cr.WriteString("\n" + paramName + ":" + strings.Join(paramValues, ","))
		}
	}
	return string(cr.Bytes()), nil
}
