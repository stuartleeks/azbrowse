package handlers

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/lawrencegripper/azbrowse/pkg/swagger"
)

type kubernetesListResponse struct {
	Items []kubernetesItem `json:"items"`
}
type kubernetesItem struct {
	Metadata struct {
		Name     string `yaml:"name"`
		SelfLink string `yaml:"selfLink"`
	} `yaml:"metadata"`
}
type podResponse struct {
	Spec struct {
		Containers []struct {
			Name string `yaml:"name"`
		} `yaml:"containers"`
	} `yaml:"spec"`
}

// SwaggerAPISetContainerService holds the config for working with an AKS cluster API
type SwaggerAPISetContainerService struct {
	resourceTypes []swagger.ResourceType
	httpClient    http.Client
	clusterID     string
	serverURL     string
}

// NewSwaggerAPISetContainerService creates a new SwaggerAPISetContainerService
func NewSwaggerAPISetContainerService(resourceTypes []swagger.ResourceType, httpClient http.Client, clusterID string, serverURL string) SwaggerAPISetContainerService {
	c := SwaggerAPISetContainerService{}
	c.resourceTypes = resourceTypes
	c.httpClient = httpClient
	c.clusterID = clusterID
	c.serverURL = serverURL
	return c
}

// ID returns the ID for the APISet
func (c SwaggerAPISetContainerService) ID() string {
	return c.clusterID
}

// MatchChildNodesByName indicates whether child nodes should be matched by name (or position)
func (c SwaggerAPISetContainerService) MatchChildNodesByName() bool {
	return false
}

// AppliesToNode is called by the Swagger exapnder to test whether the node applies to this APISet
func (c SwaggerAPISetContainerService) AppliesToNode(node *TreeNode) bool {
	// this function is only called for nodes that don't have the SwaggerAPISetID set
	// this should never happen for containerService nodes
	return false
}

// GetResourceTypes returns the ResourceTypes for the API Set
func (c SwaggerAPISetContainerService) GetResourceTypes() []swagger.ResourceType {
	return c.resourceTypes
}

func (c SwaggerAPISetContainerService) doRequest(verb string, url string) (string, error) {
	request, err := http.NewRequest("GET", url, bytes.NewReader([]byte("")))
	if err != nil {
		err = fmt.Errorf("Failed to create request" + err.Error() + url)
		return "", err
	}

	request.Header.Set("Accept", "application/yaml")
	response, err := c.httpClient.Do(request)
	if err != nil {
		err = fmt.Errorf("Failed" + err.Error() + url)
		return "", err
	}
	if 200 <= response.StatusCode && response.StatusCode < 300 {
		defer response.Body.Close() //nolint: errcheck
		buf, err := ioutil.ReadAll(response.Body)
		if err != nil {
			err = fmt.Errorf("Failed to read body: %s", err)
			return "", err
		}
		data := string(buf)
		return data, nil
	}
	return "", fmt.Errorf("Response failed with %s (%s)", response.Status, url)
}

// ExpandResource returns metadata about child resources of the specified resource node
func (c SwaggerAPISetContainerService) ExpandResource(ctx context.Context, currentItem *TreeNode, resourceType swagger.ResourceType) (APISetExpandResponse, error) {

	if resourceType.Endpoint.TemplateURL == "/api/v1/namespaces/{namespace}/pods/{name}/log" {
		if strings.Contains(currentItem.ExpandURL, "?") { // we haven't already set the container name!

			logURL := c.serverURL + currentItem.ExpandURL
			containerURL := logURL[:len(logURL)-3]
			data, err := c.doRequest("GET", containerURL)
			if err != nil {
				err = fmt.Errorf("Failed to make request: %s", err)
				return APISetExpandResponse{}, err
			}

			var podInfo podResponse
			err = yaml.Unmarshal([]byte(data), &podInfo)
			if err != nil {
				err = fmt.Errorf("Error parsing YAML response: %s", err)
				return APISetExpandResponse{Response: data}, err
			}
			if podInfo.Spec.Containers == nil || len(podInfo.Spec.Containers) == 0 {
				err = fmt.Errorf("No containers in response: %s", err)
				return APISetExpandResponse{}, err
			}

			if len(podInfo.Spec.Containers) > 1 { // if only a single resopnse then fall through to just return logs for the single container
				subResources := []SubResource{}
				for _, container := range podInfo.Spec.Containers {
					subResource := SubResource{
						ID:           currentItem.ID + "/" + container.Name,
						Name:         container.Name,
						ResourceType: resourceType,
						ExpandURL:    currentItem.ExpandURL + "?container=" + container.Name,
					}
					subResources = append(subResources, subResource)
				}
				return APISetExpandResponse{
					Response:     "Pick a container to view logs",
					SubResources: subResources,
				}, nil
			}
		}
	}

	url := c.serverURL + currentItem.ExpandURL
	data, err := c.doRequest("GET", url)
	if err != nil {
		err = fmt.Errorf("Failed to make request: %s", err)
		return APISetExpandResponse{}, err
	}

	subResources := []SubResource{}

	if len(resourceType.SubResources) > 0 {
		// We have defined subResources - Unmarshal the response and add these to newItems

		var listResponse kubernetesListResponse
		err = yaml.Unmarshal([]byte(data), &listResponse)
		if err != nil {
			err = fmt.Errorf("Error parsing YAML response: %s", err)
			return APISetExpandResponse{Response: data}, err
		}

		for _, item := range listResponse.Items {
			subResourceType := getResourceTypeForURL(ctx, item.Metadata.SelfLink, resourceType.SubResources)
			if subResourceType == nil {
				err = fmt.Errorf("SubResource type not found! %s", item.Metadata.SelfLink)
				return APISetExpandResponse{Response: data}, err
			}
			name := item.Metadata.Name
			deleteURL := ""
			if subResourceType.DeleteEndpoint != nil {
				subResourceTemplateValues := subResourceType.Endpoint.Match(item.Metadata.SelfLink).Values
				deleteURL, err = subResourceType.DeleteEndpoint.BuildURL(subResourceTemplateValues)
				if err != nil {
					err = fmt.Errorf("Error building subresource delete url '%s': %s", subResourceType.DeleteEndpoint.TemplateURL, err)
					return APISetExpandResponse{Response: data}, err
				}
			}
			subResource := SubResource{
				ID:           c.clusterID + item.Metadata.SelfLink,
				Name:         name,
				ResourceType: *subResourceType,
				ExpandURL:    item.Metadata.SelfLink,
				DeleteURL:    deleteURL,
			}
			subResources = append(subResources, subResource)
		}
	}

	return APISetExpandResponse{
		Response:     data,
		ResponseType: ResponseYAML,
		SubResources: subResources,
	}, nil
}

// Delete attempts to delete the item. Returns true if deleted, false if not handled, an error if an error occurred attempting to delete
func (c SwaggerAPISetContainerService) Delete(ctx context.Context, item *TreeNode) (bool, error) {
	if item.DeleteURL == "" {
		return false, fmt.Errorf("Item cannot be deleted (No DeleteURL)")
	}

	url := c.serverURL + item.DeleteURL
	_, err := c.doRequest("DELETE", url)
	if err != nil {
		err = fmt.Errorf("Failed to delete: %s (%s)", err.Error(), item.DeleteURL)
		return false, err
	}
	return true, nil
}