package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
	"github.com/lawrencegripper/azbrowse/pkg/swagger"
)

// The input folder structure is as below
// The bash script that generates this ensures that there is only a single version
// spec folder for each resource type. It is most likely to be `stable`, but it could be
// `preview` if no `stable` version exists for that type
//
//  swagger-specs
//   |- top-level
//          |-service1 (e.g. `cdn` or `compute`)
//          |   |-common   (want these)
//          |   |-quickstart-templates
//          |   |-data-plane
//          |   |-resource-manager (we're only interested in the contents of this folder)
//          |       |- resource-type1 (e.g. `Microsoft.Compute`)
//          |       |    |- common
//          |       |    |   |- *.json (want these)
//          |       |    |- stable (NB - may preview if no stable)
//          |       |    |    |- 2018-10-01
//          |       |    |        |- *.json   (want these)
//          |       |- readme.md <-- this contains metadata describing the versions and what files are needed for that version
//          |       |- misc files (e.g. readme.??.md)
//           ...

func main() {

	fmt.Println("*******************************************")
	fmt.Println("  Processing ARM Specs ")
	fmt.Println("*******************************************")
	config := getARMConfig()
	paths := loadARMSwagger(config)
	writeOutput(paths, config, "./internal/pkg/expanders/swagger-armspecs.generated.go", "SwaggerAPISetARMResources")
	fmt.Println()

	// fmt.Println("*******************************************")
	// fmt.Println("  Processing Azure Search Data-plane Specs ")
	// fmt.Println("*******************************************")
	// config = getAzureSearchDataPlaneConfig()
	// paths = loadAzureSearchDataPlaneSpecs(config)
	// writeOutput(paths, config, "./internal/pkg/expanders/search.generated.go", "AzureSearchServiceExpander")
	// fmt.Println()

}

func loadARMSwagger(config *swagger.Config) []*swagger.Path {
	var paths []*swagger.Path
	serviceFileInfos, err := ioutil.ReadDir("swagger-specs")
	if err != nil {
		panic(err)
	}
	for _, serviceFileInfo := range serviceFileInfos {
		if serviceFileInfo.IsDir() && serviceFileInfo.Name() != "common-types" {
			fmt.Printf("Processing service folder: %s\n", serviceFileInfo.Name())
			readmePath := fmt.Sprintf("swagger-specs/%s/resource-manager/readme.md", serviceFileInfo.Name())
			versions, err := swagger.GetVersionsFromAutoRestReadme(readmePath)
			if err != nil {
				// Could be just
				fmt.Printf("Error reading '%s': %s\n", readmePath, err)
				continue
			}
			if len(versions) == 0 {
				fmt.Printf("No versions found in '%s'\n", readmePath)
				continue
			}
			// TODO: add logic to determin which version to take!
			version := versions[0]

			apiPaths := []swagger.Path{}

			for _, inputFile := range version.Files {
				doc := loadDoc(inputFile)
				newPaths, err := swagger.GetPathsFromSwagger(doc, config, "")
				if err != nil {
					panic(err)
				}
				apiPaths = append(apiPaths, newPaths...)
			}

			if len(apiPaths) > 0 {
				paths, err = swagger.MergeSwaggerPaths(paths, config, apiPaths, true, "")
				if err != nil {
					panic(err)
				}
			}
		}
	}
	return paths
}

// getARMConfig returns the config for ARM Swagger processing
func getARMConfig() *swagger.Config {
	config := &swagger.Config{
		Overrides: map[string]swagger.PathOverride{
			// App Service patches
			"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/appsettings/list": {
				Path:    "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/appsettings",
				GetVerb: "post",
			},
			"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/authsettings/list": {
				Path:    "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/authsettings",
				GetVerb: "post",
			},
			"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/azurestorageaccounts/list": {
				Path:    "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/azurestorageaccounts",
				GetVerb: "post",
			},
			"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/backup/list": {
				Path:    "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/backup",
				GetVerb: "post",
			},
			"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/connectionstrings/list": {
				Path:    "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/connectionstrings",
				GetVerb: "post",
			},
			"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/metadata/list": {
				Path:    "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/metadata",
				GetVerb: "post",
			},
			"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/publishingcredentials/list": {
				Path:    "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/publishingcredentials",
				GetVerb: "post",
			},
			"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/pushsettings/list": {
				Path:    "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/pushsettings",
				GetVerb: "post",
			},
			// Search patches
			"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/listAdminKeys": {
				Path:    "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/listAdminKeys", // no change to path
				GetVerb: "post",
			},
			// VM Scale Sets patches
			"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines": {
				Path:        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines",
				RewritePath: true,
			},
			"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualmachines/{instanceId}": {
				Path:        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}",
				RewritePath: true,
			},
			"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualmachines/{instanceId}/instanceView": {
				Path:        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/instanceView",
				RewritePath: true,
			},
			"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/publicipaddresses": {
				Path:        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/publicipaddresses",
				RewritePath: true,
			},
			"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces/{networkInterfaceName}/ipconfigurations/{ipConfigurationName}/publicipaddresses": {
				Path:        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces/{networkInterfaceName}/ipconfigurations/{ipConfigurationName}/publicipaddresses",
				RewritePath: true,
			},
		},
	}
	return config
}

// getARMConfig returns the config for ARM Swagger processing
func getAzureSearchDataPlaneConfig() *swagger.Config {
	return &swagger.Config{
		Overrides: map[string]swagger.PathOverride{
			"/indexes('{indexName}')/docs('{key}')": {
				PutPath:    "/indexes('{indexName}')/docs/index", // update requires POST, but the APISet will handle that
				DeletePath: "/indexes('{indexName}')/docs/index", // update requires POST, but the APISet will handle that
			},
		},
	}
}
func loadAzureSearchDataPlaneSpecs(config *swagger.Config) []*swagger.Path {
	var paths []*swagger.Path

	directoryNames := []string{"Microsoft.Azure.Search.Service", "Microsoft.Azure.Search.Data"} // need to control the document load order
	for _, directoryName := range directoryNames {
		swaggerPath := getFirstNonCommonPath(getFirstNonCommonPath(fmt.Sprintf("swagger-specs/search/data-plane/%s", directoryName)))
		swaggerFileInfos, err := ioutil.ReadDir(swaggerPath)
		if err != nil {
			panic(err)
		}
		for _, swaggerFileInfo := range swaggerFileInfos {
			if !swaggerFileInfo.IsDir() && strings.HasSuffix(swaggerFileInfo.Name(), ".json") {
				fmt.Printf("\tprocessing %s/%s\n", swaggerPath, swaggerFileInfo.Name())
				doc := loadDoc(swaggerPath + "/" + swaggerFileInfo.Name())
				pathPrefix := ""
				if swaggerFileInfo.Name() == "searchindex.json" {
					// searchindex.json uses a custom property to set a base URL that the paths in that file are relative to
					// I couldn't find a way to retrieve it with the swagger library so adding some config here
					pathPrefix = "/indexes('{indexName}')"
				}
				paths, err = swagger.MergeSwaggerDoc(paths, config, doc, true, pathPrefix)
				if err != nil {
					panic(err)
				}
			}
		}
	}
	return paths
}

func loadDoc(path string) *loads.Document {

	document, err := loads.Spec(path)
	if err != nil {
		log.Panicf("Error opening Swagger: %s", err)
	}

	document, err = document.Expanded(&spec.ExpandOptions{RelativeBase: path})
	if err != nil {
		log.Panicf("Error expanding Swagger: %s", err)
	}

	return document
}
func writeOutput(paths []*swagger.Path, config *swagger.Config, filename string, structName string) {
	writer, err := os.Create(filename)
	if err != nil {
		panic(fmt.Errorf("Error opening file: %s", err))
	}
	defer func() {
		err := writer.Close()
		if err != nil {
			panic(fmt.Errorf("Failed to close output file: %s", err))
		}
	}()

	writeTemplate(writer, paths, config, structName)
}
func writeTemplate(w io.Writer, paths []*swagger.Path, config *swagger.Config, structName string) {

	funcMap := template.FuncMap{
		"upper": strings.ToUpper,
	}
	t := template.Must(template.New("code-gen").Funcs(funcMap).Parse(tmpl))

	type Context struct {
		Paths      []*swagger.Path
		StructName string
	}

	context := Context{
		Paths:      paths,
		StructName: structName,
	}

	err := t.Execute(w, context)
	if err != nil {
		panic(err)
	}
}

// getFirstNonCommonPath returns the first subfolder under path that is not named 'common'
func getFirstNonCommonPath(path string) string {
	// get the first non `common` path

	subfolders, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, subpath := range subfolders {
		if subpath.IsDir() && subpath.Name() != "common" {
			return path + "/" + subpath.Name()
		}
	}
	panic(fmt.Errorf("No suitable path found"))
}
