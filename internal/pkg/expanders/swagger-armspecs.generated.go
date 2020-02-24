


package expanders

import (
	"github.com/lawrencegripper/azbrowse/pkg/endpoints"	
	"github.com/lawrencegripper/azbrowse/pkg/swagger"	
)

func (e *SwaggerAPISetARMResources) loadResourceTypes() []swagger.ResourceType {
	return  []swagger.ResourceType{ 
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.EnterpriseKnowledgeGraph/operations", "2018-12-03"),
},
{ 
	Display: "services",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.EnterpriseKnowledgeGraph/services", "2018-12-03"),
},
{ 
	Display: "services",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EnterpriseKnowledgeGraph/services", "2018-12-03"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{resourceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EnterpriseKnowledgeGraph/services/{resourceName}", "2018-12-03"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EnterpriseKnowledgeGraph/services/{resourceName}", "2018-12-03"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EnterpriseKnowledgeGraph/services/{resourceName}", "2018-12-03"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EnterpriseKnowledgeGraph/services/{resourceName}", "2018-12-03"),
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Addons/operations", "2018-03-01"),
},
{ 
	Display: "{planTypeName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Addons/supportProviders/{providerName}/supportPlanTypes/{planTypeName}", "2018-03-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Addons/supportProviders/{providerName}/supportPlanTypes/{planTypeName}", "2018-03-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Addons/supportProviders/{providerName}/supportPlanTypes/{planTypeName}", "2018-03-01"),
},
{ 
	Display: "addsservices",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices", "2014-01-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "premiumCheck",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/premiumCheck", "2014-01-01"),
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{serviceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}", "2014-01-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}", "2014-01-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}", "2014-01-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "addomainservicemembers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/addomainservicemembers", "2014-01-01"),
},
{ 
	Display: "addsservicemembers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/addsservicemembers", "2014-01-01"),
},
{ 
	Display: "alerts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/alerts", "2014-01-01"),
},
{ 
	Display: "configuration",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/configuration", "2014-01-01"),
},
{ 
	Display: "forestsummary",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/forestsummary", "2014-01-01"),
},
{ 
	Display: "metricmetadata",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/metricmetadata", "2014-01-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{metricName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/metricmetadata/{metricName}", "2014-01-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{groupName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/metricmetadata/{metricName}/groups/{groupName}", "2014-01-01"),
}, } ,
}, } ,
},
{ 
	Display: "replicationdetails",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/replicationdetails", "2014-01-01"),
},
{ 
	Display: "replicationstatus",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/replicationstatus", "2014-01-01"),
},
{ 
	Display: "replicationsummary",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/replicationsummary", "2014-01-01"),
},
{ 
	Display: "servicemembers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/servicemembers", "2014-01-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{serviceMemberId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/servicemembers/{serviceMemberId}", "2014-01-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/servicemembers/{serviceMemberId}", "2014-01-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "alerts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/servicemembers/{serviceMemberId}/alerts", "2014-01-01"),
},
{ 
	Display: "credentials",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/servicemembers/{serviceMemberId}/credentials", "2014-01-01"),
}, } ,
}, } ,
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{dimension}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/dimensions/{dimension}", "2014-01-01"),
},
{ 
	Display: "userpreference",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/features/{featureName}/userpreference", "2014-01-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/features/{featureName}/userpreference", "2014-01-01"),
},
{ 
	Display: "{groupName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/metrics/{metricName}/groups/{groupName}", "2014-01-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "average",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/metrics/{metricName}/groups/{groupName}/average", "2014-01-01"),
},
{ 
	Display: "sum",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/metrics/{metricName}/groups/{groupName}/sum", "2014-01-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "configuration",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/configuration", "2014-01-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/configuration", "2014-01-01"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/operations", "2014-01-01"),
},
{ 
	Display: "IsDevOps",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/reports/DevOps/IsDevOps", "2014-01-01"),
},
{ 
	Display: "connectors",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/service/{serviceName}/servicemembers/{serviceMemberId}/connectors", "2014-01-01"),
},
{ 
	Display: "services",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services", "2014-01-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "premiumCheck",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/premiumCheck", "2014-01-01"),
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{serviceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}", "2014-01-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}", "2014-01-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}", "2014-01-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "alerts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/alerts", "2014-01-01"),
},
{ 
	Display: "counts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/exporterrors/counts", "2014-01-01"),
},
{ 
	Display: "listV2",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/exporterrors/listV2", "2014-01-01"),
},
{ 
	Display: "exportstatus",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/exportstatus", "2014-01-01"),
},
{ 
	Display: "metricmetadata",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/metricmetadata", "2014-01-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{metricName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/metricmetadata/{metricName}", "2014-01-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{groupName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/metricmetadata/{metricName}/groups/{groupName}", "2014-01-01"),
}, } ,
}, } ,
},
{ 
	Display: "monitoringconfigurations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/monitoringconfigurations", "2014-01-01"),
},
{ 
	Display: "user",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/reports/badpassword/details/user", "2014-01-01"),
},
{ 
	Display: "blobUris",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/reports/riskyIp/blobUris", "2014-01-01"),
},
{ 
	Display: "servicemembers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/servicemembers", "2014-01-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{serviceMemberId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/servicemembers/{serviceMemberId}", "2014-01-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/servicemembers/{serviceMemberId}", "2014-01-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "alerts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/servicemembers/{serviceMemberId}/alerts", "2014-01-01"),
},
{ 
	Display: "credentials",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/servicemembers/{serviceMemberId}/credentials", "2014-01-01"),
},
{ 
	Display: "datafreshness",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/servicemembers/{serviceMemberId}/datafreshness", "2014-01-01"),
},
{ 
	Display: "exportstatus",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/servicemembers/{serviceMemberId}/exportstatus", "2014-01-01"),
},
{ 
	Display: "globalconfiguration",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/servicemembers/{serviceMemberId}/globalconfiguration", "2014-01-01"),
},
{ 
	Display: "serviceconfiguration",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/servicemembers/{serviceMemberId}/serviceconfiguration", "2014-01-01"),
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{metricName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/servicemembers/{serviceMemberId}/metrics/{metricName}", "2014-01-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{groupName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/servicemembers/{serviceMemberId}/metrics/{metricName}/groups/{groupName}", "2014-01-01"),
}, } ,
}, } ,
}, } ,
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{featureName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/TenantWhitelisting/{featureName}", "2014-01-01"),
},
{ 
	Display: "{featureName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/checkServiceFeatureAvailibility/{featureName}", "2014-01-01"),
},
{ 
	Display: "alertfeedback",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/feedbacktype/alerts/{shortName}/alertfeedback", "2014-01-01"),
},
{ 
	Display: "{groupName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/metrics/{metricName}/groups/{groupName}", "2014-01-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "average",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/metrics/{metricName}/groups/{groupName}/average", "2014-01-01"),
},
{ 
	Display: "sum",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/metrics/{metricName}/groups/{groupName}/sum", "2014-01-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "metadata",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Advisor/metadata", "2020-01-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Advisor/metadata/{name}", "2020-01-01"),
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Advisor/operations", "2020-01-01"),
},
{ 
	Display: "configurations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/configurations", "2020-01-01"),
	SubResources: []swagger.ResourceType{  } ,
},
{ 
	Display: "recommendations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/recommendations", "2020-01-01"),
},
{ 
	Display: "suppressions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/suppressions", "2020-01-01"),
},
{ 
	Display: "configurations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Advisor/configurations", "2020-01-01"),
	SubResources: []swagger.ResourceType{  } ,
},
{ 
	Display: "{recommendationId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceUri}/providers/Microsoft.Advisor/recommendations/{recommendationId}", "2020-01-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceUri}/providers/Microsoft.Advisor/recommendations/{recommendationId}/suppressions/{name}", "2020-01-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceUri}/providers/Microsoft.Advisor/recommendations/{recommendationId}/suppressions/{name}", "2020-01-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceUri}/providers/Microsoft.Advisor/recommendations/{recommendationId}/suppressions/{name}", "2020-01-01"),
}, } ,
},
{ 
	Display: "alertsMetaData",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.AlertsManagement/alertsMetaData", "2019-05-05-preview"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.AlertsManagement/operations", "2019-05-05-preview"),
},
{ 
	Display: "actionRules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.AlertsManagement/actionRules", "2019-05-05-preview"),
},
{ 
	Display: "alerts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.AlertsManagement/alerts", "2019-05-05-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{alertId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.AlertsManagement/alerts/{alertId}", "2019-05-05-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "history",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.AlertsManagement/alerts/{alertId}/history", "2019-05-05-preview"),
}, } ,
}, } ,
},
{ 
	Display: "alertsSummary",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.AlertsManagement/alertsSummary", "2019-05-05-preview"),
},
{ 
	Display: "smartGroups",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.AlertsManagement/smartGroups", "2019-05-05-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{smartGroupId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.AlertsManagement/smartGroups/{smartGroupId}", "2019-05-05-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "history",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.AlertsManagement/smartGroups/{smartGroupId}/history", "2019-05-05-preview"),
}, } ,
}, } ,
},
{ 
	Display: "actionRules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AlertsManagement/actionRules", "2019-05-05-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{actionRuleName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AlertsManagement/actionRules/{actionRuleName}", "2019-05-05-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AlertsManagement/actionRules/{actionRuleName}", "2019-05-05-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AlertsManagement/actionRules/{actionRuleName}", "2019-05-05-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AlertsManagement/actionRules/{actionRuleName}", "2019-05-05-preview"),
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.AnalysisServices/operations", "2017-08-01-beta"),
},
{ 
	Display: "{operationId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.AnalysisServices/locations/{location}/operationresults/{operationId}", "2017-08-01-beta"),
},
{ 
	Display: "{operationId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.AnalysisServices/locations/{location}/operationstatuses/{operationId}", "2017-08-01-beta"),
},
{ 
	Display: "servers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.AnalysisServices/servers", "2017-08-01-beta"),
},
{ 
	Display: "skus",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.AnalysisServices/skus", "2017-08-01-beta"),
},
{ 
	Display: "servers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AnalysisServices/servers", "2017-08-01-beta"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{serverName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AnalysisServices/servers/{serverName}", "2017-08-01-beta"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AnalysisServices/servers/{serverName}", "2017-08-01-beta"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AnalysisServices/servers/{serverName}", "2017-08-01-beta"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AnalysisServices/servers/{serverName}", "2017-08-01-beta"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "skus",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AnalysisServices/servers/{serverName}/skus", "2017-08-01-beta"),
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ApiManagement/operations", "2019-12-01-preview"),
},
{ 
	Display: "service",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ApiManagement/service", "2019-12-01-preview"),
},
{ 
	Display: "service",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{serviceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}", "2019-12-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}", "2019-12-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}", "2019-12-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}", "2019-12-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "apiVersionSets",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apiVersionSets", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{versionSetId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apiVersionSets/{versionSetId}", "2019-12-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apiVersionSets/{versionSetId}", "2019-12-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apiVersionSets/{versionSetId}", "2019-12-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apiVersionSets/{versionSetId}", "2019-12-01-preview"),
}, } ,
},
{ 
	Display: "apis",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{apiId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}", "2019-12-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}", "2019-12-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}", "2019-12-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}", "2019-12-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "diagnostics",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/diagnostics", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{diagnosticId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/diagnostics/{diagnosticId}", "2019-12-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/diagnostics/{diagnosticId}", "2019-12-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/diagnostics/{diagnosticId}", "2019-12-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/diagnostics/{diagnosticId}", "2019-12-01-preview"),
}, } ,
},
{ 
	Display: "issues",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{issueId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}", "2019-12-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}", "2019-12-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}", "2019-12-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}", "2019-12-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "attachments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}/attachments", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{attachmentId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}/attachments/{attachmentId}", "2019-12-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}/attachments/{attachmentId}", "2019-12-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}/attachments/{attachmentId}", "2019-12-01-preview"),
}, } ,
},
{ 
	Display: "comments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}/comments", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{commentId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}/comments/{commentId}", "2019-12-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}/comments/{commentId}", "2019-12-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}/comments/{commentId}", "2019-12-01-preview"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{operationId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}", "2019-12-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}", "2019-12-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}", "2019-12-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}", "2019-12-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "policies",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}/policies", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{policyId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}/policies/{policyId}", "2019-12-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}/policies/{policyId}", "2019-12-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}/policies/{policyId}", "2019-12-01-preview"),
}, } ,
},
{ 
	Display: "tags",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}/tags", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{tagId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}/tags/{tagId}", "2019-12-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}/tags/{tagId}", "2019-12-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}/tags/{tagId}", "2019-12-01-preview"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "operationsByTags",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operationsByTags", "2019-12-01-preview"),
},
{ 
	Display: "policies",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/policies", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{policyId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/policies/{policyId}", "2019-12-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/policies/{policyId}", "2019-12-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/policies/{policyId}", "2019-12-01-preview"),
}, } ,
},
{ 
	Display: "products",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/products", "2019-12-01-preview"),
},
{ 
	Display: "releases",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/releases", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{releaseId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/releases/{releaseId}", "2019-12-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/releases/{releaseId}", "2019-12-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/releases/{releaseId}", "2019-12-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/releases/{releaseId}", "2019-12-01-preview"),
}, } ,
},
{ 
	Display: "revisions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/revisions", "2019-12-01-preview"),
},
{ 
	Display: "schemas",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/schemas", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{schemaId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/schemas/{schemaId}", "2019-12-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/schemas/{schemaId}", "2019-12-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/schemas/{schemaId}", "2019-12-01-preview"),
}, } ,
},
{ 
	Display: "tagDescriptions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/tagDescriptions", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{tagDescriptionId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/tagDescriptions/{tagDescriptionId}", "2019-12-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/tagDescriptions/{tagDescriptionId}", "2019-12-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/tagDescriptions/{tagDescriptionId}", "2019-12-01-preview"),
}, } ,
},
{ 
	Display: "tags",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/tags", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{tagId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/tags/{tagId}", "2019-12-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/tags/{tagId}", "2019-12-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/tags/{tagId}", "2019-12-01-preview"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "apisByTags",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apisByTags", "2019-12-01-preview"),
},
{ 
	Display: "authorizationServers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationServers", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{authsid}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationServers/{authsid}", "2019-12-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationServers/{authsid}", "2019-12-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationServers/{authsid}", "2019-12-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationServers/{authsid}", "2019-12-01-preview"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "backends",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/backends", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{backendId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/backends/{backendId}", "2019-12-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/backends/{backendId}", "2019-12-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/backends/{backendId}", "2019-12-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/backends/{backendId}", "2019-12-01-preview"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "caches",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/caches", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{cacheId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/caches/{cacheId}", "2019-12-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/caches/{cacheId}", "2019-12-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/caches/{cacheId}", "2019-12-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/caches/{cacheId}", "2019-12-01-preview"),
}, } ,
},
{ 
	Display: "certificates",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/certificates", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{certificateId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/certificates/{certificateId}", "2019-12-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/certificates/{certificateId}", "2019-12-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/certificates/{certificateId}", "2019-12-01-preview"),
}, } ,
},
{ 
	Display: "diagnostics",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/diagnostics", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{diagnosticId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/diagnostics/{diagnosticId}", "2019-12-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/diagnostics/{diagnosticId}", "2019-12-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/diagnostics/{diagnosticId}", "2019-12-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/diagnostics/{diagnosticId}", "2019-12-01-preview"),
}, } ,
},
{ 
	Display: "groups",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/groups", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{groupId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/groups/{groupId}", "2019-12-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/groups/{groupId}", "2019-12-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/groups/{groupId}", "2019-12-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/groups/{groupId}", "2019-12-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "users",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/groups/{groupId}/users", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{  } ,
}, } ,
}, } ,
},
{ 
	Display: "identityProviders",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/identityProviders", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{identityProviderName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/identityProviders/{identityProviderName}", "2019-12-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/identityProviders/{identityProviderName}", "2019-12-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/identityProviders/{identityProviderName}", "2019-12-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/identityProviders/{identityProviderName}", "2019-12-01-preview"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "issues",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/issues", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{issueId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/issues/{issueId}", "2019-12-01-preview"),
}, } ,
},
{ 
	Display: "loggers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/loggers", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{loggerId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/loggers/{loggerId}", "2019-12-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/loggers/{loggerId}", "2019-12-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/loggers/{loggerId}", "2019-12-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/loggers/{loggerId}", "2019-12-01-preview"),
}, } ,
},
{ 
	Display: "namedValues",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/namedValues", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{namedValueId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/namedValues/{namedValueId}", "2019-12-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/namedValues/{namedValueId}", "2019-12-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/namedValues/{namedValueId}", "2019-12-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/namedValues/{namedValueId}", "2019-12-01-preview"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "networkstatus",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/networkstatus", "2019-12-01-preview"),
},
{ 
	Display: "notifications",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{notificationName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications/{notificationName}", "2019-12-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications/{notificationName}", "2019-12-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "recipientEmails",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications/{notificationName}/recipientEmails", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{  } ,
},
{ 
	Display: "recipientUsers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications/{notificationName}/recipientUsers", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{  } ,
}, } ,
}, } ,
},
{ 
	Display: "openidConnectProviders",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/openidConnectProviders", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{opid}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/openidConnectProviders/{opid}", "2019-12-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/openidConnectProviders/{opid}", "2019-12-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/openidConnectProviders/{opid}", "2019-12-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/openidConnectProviders/{opid}", "2019-12-01-preview"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "policies",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/policies", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{policyId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/policies/{policyId}", "2019-12-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/policies/{policyId}", "2019-12-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/policies/{policyId}", "2019-12-01-preview"),
}, } ,
},
{ 
	Display: "policyDescriptions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/policyDescriptions", "2019-12-01-preview"),
},
{ 
	Display: "delegation",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/delegation", "2019-12-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/delegation", "2019-12-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/delegation", "2019-12-01-preview"),
	Children: []swagger.ResourceType{  } ,
},
{ 
	Display: "signin",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/signin", "2019-12-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/signin", "2019-12-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/signin", "2019-12-01-preview"),
},
{ 
	Display: "signup",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/signup", "2019-12-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/signup", "2019-12-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/signup", "2019-12-01-preview"),
},
{ 
	Display: "products",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{productId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}", "2019-12-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}", "2019-12-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}", "2019-12-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}", "2019-12-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "apis",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/apis", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{  } ,
},
{ 
	Display: "groups",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/groups", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{  } ,
},
{ 
	Display: "policies",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/policies", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{policyId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/policies/{policyId}", "2019-12-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/policies/{policyId}", "2019-12-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/policies/{policyId}", "2019-12-01-preview"),
}, } ,
},
{ 
	Display: "subscriptions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/subscriptions", "2019-12-01-preview"),
},
{ 
	Display: "tags",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/tags", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{tagId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/tags/{tagId}", "2019-12-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/tags/{tagId}", "2019-12-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/tags/{tagId}", "2019-12-01-preview"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "productsByTags",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/productsByTags", "2019-12-01-preview"),
},
{ 
	Display: "regions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/regions", "2019-12-01-preview"),
},
{ 
	Display: "byApi",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/byApi", "2019-12-01-preview"),
},
{ 
	Display: "byGeo",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/byGeo", "2019-12-01-preview"),
},
{ 
	Display: "byOperation",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/byOperation", "2019-12-01-preview"),
},
{ 
	Display: "byProduct",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/byProduct", "2019-12-01-preview"),
},
{ 
	Display: "byRequest",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/byRequest", "2019-12-01-preview"),
},
{ 
	Display: "bySubscription",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/bySubscription", "2019-12-01-preview"),
},
{ 
	Display: "byTime",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/byTime", "2019-12-01-preview"),
},
{ 
	Display: "byUser",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/byUser", "2019-12-01-preview"),
},
{ 
	Display: "skus",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/skus", "2019-12-01-preview"),
},
{ 
	Display: "subscriptions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{sid}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}", "2019-12-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}", "2019-12-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}", "2019-12-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}", "2019-12-01-preview"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "tagResources",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tagResources", "2019-12-01-preview"),
},
{ 
	Display: "tags",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tags", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{tagId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tags/{tagId}", "2019-12-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tags/{tagId}", "2019-12-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tags/{tagId}", "2019-12-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tags/{tagId}", "2019-12-01-preview"),
}, } ,
},
{ 
	Display: "templates",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/templates", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{templateName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/templates/{templateName}", "2019-12-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/templates/{templateName}", "2019-12-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/templates/{templateName}", "2019-12-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/templates/{templateName}", "2019-12-01-preview"),
}, } ,
},
{ 
	Display: "users",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{userId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{userId}", "2019-12-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{userId}", "2019-12-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{userId}", "2019-12-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{userId}", "2019-12-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "groups",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{userId}/groups", "2019-12-01-preview"),
},
{ 
	Display: "identities",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{userId}/identities", "2019-12-01-preview"),
},
{ 
	Display: "subscriptions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{userId}/subscriptions", "2019-12-01-preview"),
}, } ,
}, } ,
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "networkstatus",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/locations/{locationName}/networkstatus", "2019-12-01-preview"),
},
{ 
	Display: "{quotaCounterKey}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/quotas/{quotaCounterKey}", "2019-12-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/quotas/{quotaCounterKey}", "2019-12-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{quotaPeriodKey}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/quotas/{quotaCounterKey}/periods/{quotaPeriodKey}", "2019-12-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/quotas/{quotaCounterKey}/periods/{quotaPeriodKey}", "2019-12-01-preview"),
}, } ,
},
{ 
	Display: "{accessName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{accessName}", "2019-12-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{accessName}", "2019-12-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "git",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{accessName}/git", "2019-12-01-preview"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "syncState",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{configurationName}/syncState", "2019-12-01-preview"),
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.AppConfiguration/operations", "2019-10-01"),
},
{ 
	Display: "configurationStores",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.AppConfiguration/configurationStores", "2019-10-01"),
},
{ 
	Display: "configurationStores",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores", "2019-10-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{configStoreName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}", "2019-10-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}", "2019-10-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}", "2019-10-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}", "2019-10-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "current",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/components/{resourceName}/pricingPlans/current", "2017-10-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/components/{resourceName}/pricingPlans/current", "2017-10-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/components/{resourceName}/pricingPlans/current", "2017-10-01"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.AppPlatform/operations", "2019-05-01-preview"),
},
{ 
	Display: "Spring",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.AppPlatform/Spring", "2019-05-01-preview"),
},
{ 
	Display: "Spring",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring", "2019-05-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{serviceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}", "2019-05-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}", "2019-05-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}", "2019-05-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}", "2019-05-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "apps",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps", "2019-05-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{appName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}", "2019-05-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}", "2019-05-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}", "2019-05-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}", "2019-05-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "bindings",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}/bindings", "2019-05-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{bindingName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}/bindings/{bindingName}", "2019-05-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}/bindings/{bindingName}", "2019-05-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}/bindings/{bindingName}", "2019-05-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}/bindings/{bindingName}", "2019-05-01-preview"),
}, } ,
},
{ 
	Display: "deployments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}/deployments", "2019-05-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{deploymentName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}/deployments/{deploymentName}", "2019-05-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}/deployments/{deploymentName}", "2019-05-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}/deployments/{deploymentName}", "2019-05-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/apps/{appName}/deployments/{deploymentName}", "2019-05-01-preview"),
	Children: []swagger.ResourceType{  } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "deployments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/deployments", "2019-05-01-preview"),
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Attestation/operations", "2018-09-01-preview"),
},
{ 
	Display: "attestationProviders",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Attestation/attestationProviders", "2018-09-01-preview"),
},
{ 
	Display: "attestationProviders",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Attestation/attestationProviders", "2018-09-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{providerName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Attestation/attestationProviders/{providerName}", "2018-09-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Attestation/attestationProviders/{providerName}", "2018-09-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Attestation/attestationProviders/{providerName}", "2018-09-01-preview"),
}, } ,
},
{ 
	Display: "providerOperations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Authorization/providerOperations", "2015-07-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{resourceProviderNamespace}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Authorization/providerOperations/{resourceProviderNamespace}", "2015-07-01"),
}, } ,
},
{ 
	Display: "roleAssignments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/roleAssignments", "2015-07-01"),
},
{ 
	Display: "roleAssignments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Authorization/roleAssignments", "2015-07-01"),
},
{ 
	Display: "permissions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Authorization/permissions", "2015-07-01"),
},
{ 
	Display: "permissions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}/providers/Microsoft.Authorization/permissions", "2015-07-01"),
},
{ 
	Display: "roleAssignments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}/providers/Microsoft.Authorization/roleAssignments", "2015-07-01"),
},
{ 
	Display: "{roleAssignmentId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{roleAssignmentId}", "2015-07-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{roleAssignmentId}", "2015-07-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/{roleAssignmentId}", "2015-07-01"),
},
{ 
	Display: "{roleDefinitionId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{roleDefinitionId}", "2015-07-01"),
},
{ 
	Display: "roleAssignments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Authorization/roleAssignments", "2015-07-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{roleAssignmentName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}", "2015-07-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}", "2015-07-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}", "2015-07-01"),
}, } ,
},
{ 
	Display: "roleDefinitions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Authorization/roleDefinitions", "2015-07-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{roleDefinitionId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Authorization/roleDefinitions/{roleDefinitionId}", "2015-07-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Authorization/roleDefinitions/{roleDefinitionId}", "2015-07-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Authorization/roleDefinitions/{roleDefinitionId}", "2015-07-01"),
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Automation/operations", "2015-10-31"),
},
{ 
	Display: "automationAccounts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Automation/automationAccounts", "2015-10-31"),
},
{ 
	Display: "automationAccounts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts", "2015-10-31"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{automationAccountName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}", "2015-10-31"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}", "2015-10-31"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}", "2015-10-31"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}", "2015-10-31"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "agentRegistrationInformation",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/agentRegistrationInformation", "2015-10-31"),
	Children: []swagger.ResourceType{  } ,
},
{ 
	Display: "certificates",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/certificates", "2015-10-31"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{certificateName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/certificates/{certificateName}", "2015-10-31"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/certificates/{certificateName}", "2015-10-31"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/certificates/{certificateName}", "2015-10-31"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/certificates/{certificateName}", "2015-10-31"),
}, } ,
},
{ 
	Display: "compilationjobs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/compilationjobs", "2015-10-31"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{compilationJobId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/compilationjobs/{compilationJobId}", "2015-10-31"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/compilationjobs/{compilationJobId}", "2015-10-31"),
},
{ 
	Display: "streams",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/compilationjobs/{jobId}/streams", "2015-10-31"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{jobStreamId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/compilationjobs/{jobId}/streams/{jobStreamId}", "2015-10-31"),
}, } ,
}, } ,
},
{ 
	Display: "configurations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/configurations", "2015-10-31"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{configurationName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/configurations/{configurationName}", "2015-10-31"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/configurations/{configurationName}", "2015-10-31"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/configurations/{configurationName}", "2015-10-31"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/configurations/{configurationName}", "2015-10-31"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "content",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/configurations/{configurationName}/content", "2015-10-31"),
}, } ,
}, } ,
},
{ 
	Display: "connectionTypes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/connectionTypes", "2015-10-31"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{connectionTypeName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/connectionTypes/{connectionTypeName}", "2015-10-31"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/connectionTypes/{connectionTypeName}", "2015-10-31"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/connectionTypes/{connectionTypeName}", "2015-10-31"),
}, } ,
},
{ 
	Display: "connections",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/connections", "2015-10-31"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{connectionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/connections/{connectionName}", "2015-10-31"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/connections/{connectionName}", "2015-10-31"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/connections/{connectionName}", "2015-10-31"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/connections/{connectionName}", "2015-10-31"),
}, } ,
},
{ 
	Display: "credentials",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/credentials", "2015-10-31"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{credentialName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/credentials/{credentialName}", "2015-10-31"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/credentials/{credentialName}", "2015-10-31"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/credentials/{credentialName}", "2015-10-31"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/credentials/{credentialName}", "2015-10-31"),
}, } ,
},
{ 
	Display: "hybridRunbookWorkerGroups",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/hybridRunbookWorkerGroups", "2015-10-31"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{hybridRunbookWorkerGroupName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/hybridRunbookWorkerGroups/{hybridRunbookWorkerGroupName}", "2015-10-31"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/hybridRunbookWorkerGroups/{hybridRunbookWorkerGroupName}", "2015-10-31"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/hybridRunbookWorkerGroups/{hybridRunbookWorkerGroupName}", "2015-10-31"),
}, } ,
},
{ 
	Display: "jobSchedules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobSchedules", "2015-10-31"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{jobScheduleId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobSchedules/{jobScheduleId}", "2015-10-31"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobSchedules/{jobScheduleId}", "2015-10-31"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobSchedules/{jobScheduleId}", "2015-10-31"),
}, } ,
},
{ 
	Display: "jobs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobs", "2015-10-31"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{jobId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobs/{jobId}", "2015-10-31"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobs/{jobId}", "2015-10-31"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "output",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobs/{jobId}/output", "2015-10-31"),
},
{ 
	Display: "runbookContent",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobs/{jobId}/runbookContent", "2015-10-31"),
},
{ 
	Display: "streams",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobs/{jobId}/streams", "2015-10-31"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{jobStreamId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobs/{jobId}/streams/{jobStreamId}", "2015-10-31"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "linkedWorkspace",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/linkedWorkspace", "2015-10-31"),
},
{ 
	Display: "modules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/modules", "2015-10-31"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{moduleName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/modules/{moduleName}", "2015-10-31"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/modules/{moduleName}", "2015-10-31"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/modules/{moduleName}", "2015-10-31"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/modules/{moduleName}", "2015-10-31"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "activities",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/modules/{moduleName}/activities", "2015-10-31"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{activityName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/modules/{moduleName}/activities/{activityName}", "2015-10-31"),
}, } ,
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "fields",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/modules/{moduleName}/objectDataTypes/{typeName}/fields", "2015-10-31"),
},
{ 
	Display: "fields",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/modules/{moduleName}/types/{typeName}/fields", "2015-10-31"),
}, } ,
}, } ,
},
{ 
	Display: "nodeConfigurations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodeConfigurations", "2015-10-31"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{nodeConfigurationName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodeConfigurations/{nodeConfigurationName}", "2015-10-31"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodeConfigurations/{nodeConfigurationName}", "2015-10-31"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodeConfigurations/{nodeConfigurationName}", "2015-10-31"),
}, } ,
},
{ 
	Display: "nodes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodes", "2015-10-31"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{nodeId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodes/{nodeId}", "2015-10-31"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodes/{nodeId}", "2015-10-31"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodes/{nodeId}", "2015-10-31"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "reports",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodes/{nodeId}/reports", "2015-10-31"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{reportId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodes/{nodeId}/reports/{reportId}", "2015-10-31"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "content",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodes/{nodeId}/reports/{reportId}/content", "2015-10-31"),
}, } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "runbooks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks", "2015-10-31"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{runbookName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}", "2015-10-31"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}", "2015-10-31"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}", "2015-10-31"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}", "2015-10-31"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "content",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/content", "2015-10-31"),
},
{ 
	Display: "draft",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft", "2015-10-31"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "content",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/content", "2015-10-31"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/content", "2015-10-31"),
},
{ 
	Display: "testJob",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/testJob", "2015-10-31"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/testJob", "2015-10-31"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "streams",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/testJob/streams", "2015-10-31"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{jobStreamId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/testJob/streams/{jobStreamId}", "2015-10-31"),
}, } ,
}, } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "schedules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/schedules", "2015-10-31"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{scheduleName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/schedules/{scheduleName}", "2015-10-31"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/schedules/{scheduleName}", "2015-10-31"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/schedules/{scheduleName}", "2015-10-31"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/schedules/{scheduleName}", "2015-10-31"),
}, } ,
},
{ 
	Display: "statistics",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/statistics", "2015-10-31"),
},
{ 
	Display: "usages",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/usages", "2015-10-31"),
},
{ 
	Display: "variables",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/variables", "2015-10-31"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{variableName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/variables/{variableName}", "2015-10-31"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/variables/{variableName}", "2015-10-31"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/variables/{variableName}", "2015-10-31"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/variables/{variableName}", "2015-10-31"),
}, } ,
},
{ 
	Display: "webhooks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/webhooks", "2015-10-31"),
	Children: []swagger.ResourceType{  } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{webhookName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/webhooks/{webhookName}", "2015-10-31"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/webhooks/{webhookName}", "2015-10-31"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/webhooks/{webhookName}", "2015-10-31"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/webhooks/{webhookName}", "2015-10-31"),
}, } ,
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "fields",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/objectDataTypes/{typeName}/fields", "2015-10-31"),
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Kusto/operations", "2019-09-07"),
},
{ 
	Display: "clusters",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Kusto/clusters", "2019-09-07"),
},
{ 
	Display: "skus",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Kusto/skus", "2019-09-07"),
},
{ 
	Display: "clusters",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters", "2019-09-07"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{clusterName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}", "2019-09-07"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}", "2019-09-07"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}", "2019-09-07"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}", "2019-09-07"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "attachedDatabaseConfigurations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/attachedDatabaseConfigurations", "2019-09-07"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{attachedDatabaseConfigurationName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/attachedDatabaseConfigurations/{attachedDatabaseConfigurationName}", "2019-09-07"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/attachedDatabaseConfigurations/{attachedDatabaseConfigurationName}", "2019-09-07"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/attachedDatabaseConfigurations/{attachedDatabaseConfigurationName}", "2019-09-07"),
}, } ,
},
{ 
	Display: "databases",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases", "2019-09-07"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{databaseName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}", "2019-09-07"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}", "2019-09-07"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}", "2019-09-07"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}", "2019-09-07"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "dataConnections",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/dataConnections", "2019-09-07"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{dataConnectionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/dataConnections/{dataConnectionName}", "2019-09-07"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/dataConnections/{dataConnectionName}", "2019-09-07"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/dataConnections/{dataConnectionName}", "2019-09-07"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/dataConnections/{dataConnectionName}", "2019-09-07"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "skus",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/skus", "2019-09-07"),
}, } ,
}, } ,
},
{ 
	Display: "diagnosticSettings",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/microsoft.aadiam/diagnosticSettings", "2017-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/microsoft.aadiam/diagnosticSettings/{name}", "2017-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/microsoft.aadiam/diagnosticSettings/{name}", "2017-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/microsoft.aadiam/diagnosticSettings/{name}", "2017-04-01"),
}, } ,
},
{ 
	Display: "diagnosticSettingsCategories",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/microsoft.aadiam/diagnosticSettingsCategories", "2017-04-01"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/microsoft.aadiam/operations", "2017-04-01"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.AzureData/operations", "2017-03-01-preview"),
},
{ 
	Display: "sqlServerRegistrations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.AzureData/sqlServerRegistrations", "2017-03-01-preview"),
},
{ 
	Display: "sqlServerRegistrations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureData/sqlServerRegistrations", "2017-03-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{sqlServerRegistrationName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureData/sqlServerRegistrations/{sqlServerRegistrationName}", "2017-03-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureData/sqlServerRegistrations/{sqlServerRegistrationName}", "2017-03-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureData/sqlServerRegistrations/{sqlServerRegistrationName}", "2017-03-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureData/sqlServerRegistrations/{sqlServerRegistrationName}", "2017-03-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "sqlServers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureData/sqlServerRegistrations/{sqlServerRegistrationName}/sqlServers", "2017-03-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{sqlServerName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureData/sqlServerRegistrations/{sqlServerRegistrationName}/sqlServers/{sqlServerName}", "2017-03-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureData/sqlServerRegistrations/{sqlServerRegistrationName}/sqlServers/{sqlServerName}", "2017-03-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureData/sqlServerRegistrations/{sqlServerRegistrationName}/sqlServers/{sqlServerName}", "2017-03-01-preview"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.AzureStack/operations", "2017-06-01"),
},
{ 
	Display: "registrations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations", "2017-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{registrationName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}", "2017-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}", "2017-06-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}", "2017-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}", "2017-06-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "customerSubscriptions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}/customerSubscriptions", "2017-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{customerSubscriptionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}/customerSubscriptions/{customerSubscriptionName}", "2017-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}/customerSubscriptions/{customerSubscriptionName}", "2017-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}/customerSubscriptions/{customerSubscriptionName}", "2017-06-01"),
}, } ,
},
{ 
	Display: "products",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}/products", "2017-06-01"),
	Children: []swagger.ResourceType{  } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{productName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}/products/{productName}", "2017-06-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Batch/operations", "2019-08-01"),
},
{ 
	Display: "batchAccounts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Batch/batchAccounts", "2019-08-01"),
},
{ 
	Display: "quotas",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Batch/locations/{locationName}/quotas", "2019-08-01"),
},
{ 
	Display: "batchAccounts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts", "2019-08-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{accountName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}", "2019-08-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}", "2019-08-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}", "2019-08-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}", "2019-08-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "applications",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications", "2019-08-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{applicationName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationName}", "2019-08-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationName}", "2019-08-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationName}", "2019-08-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationName}", "2019-08-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "versions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationName}/versions", "2019-08-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{versionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationName}/versions/{versionName}", "2019-08-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationName}/versions/{versionName}", "2019-08-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationName}/versions/{versionName}", "2019-08-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "certificates",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/certificates", "2019-08-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{certificateName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/certificates/{certificateName}", "2019-08-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/certificates/{certificateName}", "2019-08-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/certificates/{certificateName}", "2019-08-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/certificates/{certificateName}", "2019-08-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "pools",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/pools", "2019-08-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{poolName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/pools/{poolName}", "2019-08-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/pools/{poolName}", "2019-08-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/pools/{poolName}", "2019-08-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/pools/{poolName}", "2019-08-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.BatchAI/operations", "2018-05-01"),
},
{ 
	Display: "usages",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.BatchAI/locations/{location}/usages", "2018-05-01"),
},
{ 
	Display: "workspaces",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.BatchAI/workspaces", "2018-05-01"),
},
{ 
	Display: "workspaces",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces", "2018-05-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{workspaceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}", "2018-05-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}", "2018-05-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}", "2018-05-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}", "2018-05-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "clusters",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/clusters", "2018-05-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{clusterName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/clusters/{clusterName}", "2018-05-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/clusters/{clusterName}", "2018-05-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/clusters/{clusterName}", "2018-05-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/clusters/{clusterName}", "2018-05-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "experiments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/experiments", "2018-05-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{experimentName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/experiments/{experimentName}", "2018-05-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/experiments/{experimentName}", "2018-05-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/experiments/{experimentName}", "2018-05-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "jobs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/experiments/{experimentName}/jobs", "2018-05-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{jobName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/experiments/{experimentName}/jobs/{jobName}", "2018-05-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/experiments/{experimentName}/jobs/{jobName}", "2018-05-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/experiments/{experimentName}/jobs/{jobName}", "2018-05-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "fileServers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/fileServers", "2018-05-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{fileServerName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/fileServers/{fileServerName}", "2018-05-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/fileServers/{fileServerName}", "2018-05-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/fileServers/{fileServerName}", "2018-05-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "billingAccounts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts", "2019-10-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{billingAccountName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}", "2019-10-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}", "2019-10-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "agreements",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/agreements", "2019-10-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{agreementName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/agreements/{agreementName}", "2019-10-01-preview"),
}, } ,
},
{ 
	Display: "billingPermissions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingPermissions", "2019-10-01-preview"),
},
{ 
	Display: "billingProfiles",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles", "2019-10-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{billingProfileName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}", "2019-10-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}", "2019-10-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}", "2019-10-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "default",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/availableBalance/default", "2019-10-01-preview"),
},
{ 
	Display: "billingPermissions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/billingPermissions", "2019-10-01-preview"),
},
{ 
	Display: "billingRoleAssignments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/billingRoleAssignments", "2019-10-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{billingRoleAssignmentName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/billingRoleAssignments/{billingRoleAssignmentName}", "2019-10-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/billingRoleAssignments/{billingRoleAssignmentName}", "2019-10-01-preview"),
}, } ,
},
{ 
	Display: "billingRoleDefinitions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/billingRoleDefinitions", "2019-10-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{billingRoleDefinitionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/billingRoleDefinitions/{billingRoleDefinitionName}", "2019-10-01-preview"),
}, } ,
},
{ 
	Display: "billingSubscriptions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/billingSubscriptions", "2019-10-01-preview"),
},
{ 
	Display: "customers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/customers", "2019-10-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "transfers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/customers/{customerName}/transfers", "2019-10-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{transferName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/customers/{customerName}/transfers/{transferName}", "2019-10-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/customers/{customerName}/transfers/{transferName}", "2019-10-01-preview"),
}, } ,
}, } ,
},
{ 
	Display: "instructions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/instructions", "2019-10-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{instructionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/instructions/{instructionName}", "2019-10-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/instructions/{instructionName}", "2019-10-01-preview"),
}, } ,
},
{ 
	Display: "invoiceSections",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections", "2019-10-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{invoiceSectionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}", "2019-10-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}", "2019-10-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}", "2019-10-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "billingPermissions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/billingPermissions", "2019-10-01-preview"),
},
{ 
	Display: "billingRoleAssignments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/billingRoleAssignments", "2019-10-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{billingRoleAssignmentName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/billingRoleAssignments/{billingRoleAssignmentName}", "2019-10-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/billingRoleAssignments/{billingRoleAssignmentName}", "2019-10-01-preview"),
}, } ,
},
{ 
	Display: "billingRoleDefinitions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/billingRoleDefinitions", "2019-10-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{billingRoleDefinitionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/billingRoleDefinitions/{billingRoleDefinitionName}", "2019-10-01-preview"),
}, } ,
},
{ 
	Display: "billingSubscriptions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/billingSubscriptions", "2019-10-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{billingSubscriptionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/billingSubscriptions/{billingSubscriptionName}", "2019-10-01-preview"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "products",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/products", "2019-10-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{productName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/products/{productName}", "2019-10-01-preview"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "transactions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/transactions", "2019-10-01-preview"),
},
{ 
	Display: "transfers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/transfers", "2019-10-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{transferName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/transfers/{transferName}", "2019-10-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/transfers/{transferName}", "2019-10-01-preview"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "invoices",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoices", "2019-10-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{invoiceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoices/{invoiceName}", "2019-10-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "transactions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoices/{invoiceName}/transactions", "2019-10-01-preview"),
}, } ,
}, } ,
},
{ 
	Display: "paymentMethods",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/paymentMethods", "2019-10-01-preview"),
},
{ 
	Display: "default",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/policies/default", "2019-10-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/policies/default", "2019-10-01-preview"),
},
{ 
	Display: "transactions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/transactions", "2019-10-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{transactionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/transactions/{transactionName}", "2019-10-01-preview"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "billingRoleAssignments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingRoleAssignments", "2019-10-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{billingRoleAssignmentName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingRoleAssignments/{billingRoleAssignmentName}", "2019-10-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingRoleAssignments/{billingRoleAssignmentName}", "2019-10-01-preview"),
}, } ,
},
{ 
	Display: "billingRoleDefinitions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingRoleDefinitions", "2019-10-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{billingRoleDefinitionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingRoleDefinitions/{billingRoleDefinitionName}", "2019-10-01-preview"),
}, } ,
},
{ 
	Display: "billingSubscriptions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingSubscriptions", "2019-10-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "invoices",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingSubscriptions/{billingSubscriptionName}/invoices", "2019-10-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{invoiceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingSubscriptions/{billingSubscriptionName}/invoices/{invoiceName}", "2019-10-01-preview"),
}, } ,
}, } ,
},
{ 
	Display: "customers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/customers", "2019-10-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{customerName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/customers/{customerName}", "2019-10-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "billingPermissions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/customers/{customerName}/billingPermissions", "2019-10-01-preview"),
},
{ 
	Display: "billingSubscriptions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/customers/{customerName}/billingSubscriptions", "2019-10-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{billingSubscriptionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/customers/{customerName}/billingSubscriptions/{billingSubscriptionName}", "2019-10-01-preview"),
}, } ,
},
{ 
	Display: "default",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/customers/{customerName}/policies/default", "2019-10-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/customers/{customerName}/policies/default", "2019-10-01-preview"),
},
{ 
	Display: "products",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/customers/{customerName}/products", "2019-10-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{productName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/customers/{customerName}/products/{productName}", "2019-10-01-preview"),
}, } ,
},
{ 
	Display: "transactions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/customers/{customerName}/transactions", "2019-10-01-preview"),
}, } ,
}, } ,
},
{ 
	Display: "departments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/departments", "2019-10-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{departmentName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/departments/{departmentName}", "2019-10-01-preview"),
}, } ,
},
{ 
	Display: "enrollmentAccounts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/enrollmentAccounts", "2019-10-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{enrollmentAccountName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/enrollmentAccounts/{enrollmentAccountName}", "2019-10-01-preview"),
}, } ,
},
{ 
	Display: "invoices",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/invoices", "2019-10-01-preview"),
},
{ 
	Display: "paymentMethods",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/paymentMethods", "2019-10-01-preview"),
},
{ 
	Display: "products",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/products", "2019-10-01-preview"),
},
{ 
	Display: "transactions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/transactions", "2019-10-01-preview"),
}, } ,
},
{ 
	Display: "balances",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/providers/Microsoft.Billing/billingPeriods/{billingPeriodName}/providers/Microsoft.Consumption/balances", "2019-05-01-preview"),
},
{ 
	Display: "balances",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/providers/Microsoft.Consumption/balances", "2019-05-01-preview"),
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/operations", "2019-10-01-preview"),
},
{ 
	Display: "transfers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/transfers", "2019-10-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{transferName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/transfers/{transferName}", "2019-10-01-preview"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "default",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Billing/billingProperty/default", "2019-10-01-preview"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Blockchain/operations", "2018-06-01-preview"),
},
{ 
	Display: "blockchainMembers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Blockchain/blockchainMembers", "2018-06-01-preview"),
},
{ 
	Display: "{operationId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Blockchain/locations/{locationName}/blockchainMemberOperationResults/{operationId}", "2018-06-01-preview"),
},
{ 
	Display: "skus",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Blockchain/skus", "2018-06-01-preview"),
},
{ 
	Display: "blockchainMembers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers", "2018-06-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{blockchainMemberName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}", "2018-06-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}", "2018-06-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}", "2018-06-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}", "2018-06-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "consortiumMembers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}/consortiumMembers", "2018-06-01-preview"),
},
{ 
	Display: "transactionNodes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}/transactionNodes", "2018-06-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{transactionNodeName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}/transactionNodes/{transactionNodeName}", "2018-06-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}/transactionNodes/{transactionNodeName}", "2018-06-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}/transactionNodes/{transactionNodeName}", "2018-06-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}/transactionNodes/{transactionNodeName}", "2018-06-01-preview"),
	Children: []swagger.ResourceType{  } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "blueprintAssignments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Blueprint/blueprintAssignments", "2018-11-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{assignmentName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Blueprint/blueprintAssignments/{assignmentName}", "2018-11-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Blueprint/blueprintAssignments/{assignmentName}", "2018-11-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Blueprint/blueprintAssignments/{assignmentName}", "2018-11-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "assignmentOperations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Blueprint/blueprintAssignments/{assignmentName}/assignmentOperations", "2018-11-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{assignmentOperationName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Blueprint/blueprintAssignments/{assignmentName}/assignmentOperations/{assignmentOperationName}", "2018-11-01-preview"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "blueprints",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Blueprint/blueprints", "2018-11-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{blueprintName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Blueprint/blueprints/{blueprintName}", "2018-11-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Blueprint/blueprints/{blueprintName}", "2018-11-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Blueprint/blueprints/{blueprintName}", "2018-11-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "artifacts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Blueprint/blueprints/{blueprintName}/artifacts", "2018-11-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{artifactName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Blueprint/blueprints/{blueprintName}/artifacts/{artifactName}", "2018-11-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Blueprint/blueprints/{blueprintName}/artifacts/{artifactName}", "2018-11-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Blueprint/blueprints/{blueprintName}/artifacts/{artifactName}", "2018-11-01-preview"),
}, } ,
},
{ 
	Display: "versions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Blueprint/blueprints/{blueprintName}/versions", "2018-11-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{versionId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Blueprint/blueprints/{blueprintName}/versions/{versionId}", "2018-11-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Blueprint/blueprints/{blueprintName}/versions/{versionId}", "2018-11-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Blueprint/blueprints/{blueprintName}/versions/{versionId}", "2018-11-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "artifacts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Blueprint/blueprints/{blueprintName}/versions/{versionId}/artifacts", "2018-11-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{artifactName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Blueprint/blueprints/{blueprintName}/versions/{versionId}/artifacts/{artifactName}", "2018-11-01-preview"),
}, } ,
}, } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.BotService/operations", "2018-07-12"),
},
{ 
	Display: "botServices",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.BotService/botServices", "2018-07-12"),
},
{ 
	Display: "botServices",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices", "2018-07-12"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{resourceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}", "2018-07-12"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}", "2018-07-12"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}", "2018-07-12"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}", "2018-07-12"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "channels",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/channels", "2018-07-12"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{channelName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/channels/{channelName}", "2018-07-12"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/channels/{channelName}", "2018-07-12"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/channels/{channelName}", "2018-07-12"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/channels/{channelName}", "2018-07-12"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "connections",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/connections", "2018-07-12"),
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{connectionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/Connections/{connectionName}", "2018-07-12"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/Connections/{connectionName}", "2018-07-12"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/Connections/{connectionName}", "2018-07-12"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/botServices/{resourceName}/Connections/{connectionName}", "2018-07-12"),
	Children: []swagger.ResourceType{  } ,
}, } ,
}, } ,
},
{ 
	Display: "enterpriseChannels",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/enterpriseChannels", "2018-07-12"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{resourceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/enterpriseChannels/{resourceName}", "2018-07-12"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/enterpriseChannels/{resourceName}", "2018-07-12"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/enterpriseChannels/{resourceName}", "2018-07-12"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BotService/enterpriseChannels/{resourceName}", "2018-07-12"),
}, } ,
},
{ 
	Display: "edgenodes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Cdn/edgenodes", "2019-12-31"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Cdn/operations", "2019-12-31"),
},
{ 
	Display: "profiles",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Cdn/profiles", "2019-12-31"),
},
{ 
	Display: "profiles",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles", "2019-12-31"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{profileName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}", "2019-12-31"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}", "2019-12-31"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}", "2019-12-31"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}", "2019-12-31"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "endpoints",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints", "2019-12-31"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{endpointName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}", "2019-12-31"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}", "2019-12-31"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}", "2019-12-31"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}", "2019-12-31"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "customDomains",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/customDomains", "2019-12-31"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{customDomainName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/customDomains/{customDomainName}", "2019-12-31"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/customDomains/{customDomainName}", "2019-12-31"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/customDomains/{customDomainName}", "2019-12-31"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "originGroups",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/originGroups", "2019-12-31"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{originGroupName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/originGroups/{originGroupName}", "2019-12-31"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/originGroups/{originGroupName}", "2019-12-31"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/originGroups/{originGroupName}", "2019-12-31"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/originGroups/{originGroupName}", "2019-12-31"),
}, } ,
},
{ 
	Display: "origins",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/origins", "2019-12-31"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{originName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/origins/{originName}", "2019-12-31"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/origins/{originName}", "2019-12-31"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/origins/{originName}", "2019-12-31"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/origins/{originName}", "2019-12-31"),
}, } ,
}, } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.CognitiveServices/operations", "2017-04-18"),
},
{ 
	Display: "accounts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.CognitiveServices/accounts", "2017-04-18"),
},
{ 
	Display: "skus",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.CognitiveServices/skus", "2017-04-18"),
},
{ 
	Display: "accounts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts", "2017-04-18"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{accountName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}", "2017-04-18"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}", "2017-04-18"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}", "2017-04-18"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}", "2017-04-18"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "skus",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/skus", "2017-04-18"),
},
{ 
	Display: "usages",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/usages", "2017-04-18"),
}, } ,
}, } ,
},
{ 
	Display: "RateCard",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Commerce/RateCard", "2015-06-01-preview"),
},
{ 
	Display: "UsageAggregates",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Commerce/UsageAggregates", "2015-06-01-preview"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Compute/operations", "2017-12-01"),
},
{ 
	Display: "availabilitySets",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/availabilitySets", "2017-12-01"),
},
{ 
	Display: "disks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/disks", "2017-03-30"),
},
{ 
	Display: "images",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/images", "2017-12-01"),
},
{ 
	Display: "publishers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers", "2017-12-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "types",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmextension/types", "2017-12-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "versions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmextension/types/{type}/versions", "2017-12-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{version}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmextension/types/{type}/versions/{version}", "2017-12-01"),
}, } ,
}, } ,
},
{ 
	Display: "offers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmimage/offers", "2017-12-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "skus",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus", "2017-12-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "versions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus/{skus}/versions", "2017-12-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{version}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus/{skus}/versions/{version}", "2017-12-01"),
}, } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "usages",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/usages", "2017-12-01"),
},
{ 
	Display: "virtualMachines",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/virtualMachines", "2017-12-01"),
},
{ 
	Display: "vmSizes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/vmSizes", "2017-12-01"),
},
{ 
	Display: "snapshots",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/snapshots", "2017-03-30"),
},
{ 
	Display: "virtualMachineScaleSets",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/virtualMachineScaleSets", "2017-12-01"),
},
{ 
	Display: "virtualMachines",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/virtualMachines", "2017-12-01"),
},
{ 
	Display: "availabilitySets",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets", "2017-12-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{availabilitySetName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}", "2017-12-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}", "2017-12-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}", "2017-12-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}", "2017-12-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "vmSizes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}/vmSizes", "2017-12-01"),
}, } ,
}, } ,
},
{ 
	Display: "disks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks", "2017-03-30"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{diskName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}", "2017-03-30"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}", "2017-03-30"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}", "2017-03-30"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}", "2017-03-30"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "images",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images", "2017-12-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{imageName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images/{imageName}", "2017-12-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images/{imageName}", "2017-12-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images/{imageName}", "2017-12-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images/{imageName}", "2017-12-01"),
}, } ,
},
{ 
	Display: "snapshots",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots", "2017-03-30"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{snapshotName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}", "2017-03-30"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}", "2017-03-30"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}", "2017-03-30"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}", "2017-03-30"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "virtualMachineScaleSets",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets", "2017-12-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{vmScaleSetName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}", "2017-12-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}", "2017-12-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}", "2017-12-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}", "2017-12-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "extensions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/extensions", "2017-12-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{vmssExtensionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/extensions/{vmssExtensionName}", "2017-12-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/extensions/{vmssExtensionName}", "2017-12-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/extensions/{vmssExtensionName}", "2017-12-01"),
}, } ,
},
{ 
	Display: "instanceView",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/instanceView", "2017-12-01"),
},
{ 
	Display: "osUpgradeHistory",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/osUpgradeHistory", "2017-12-01"),
},
{ 
	Display: "latest",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/rollingUpgrades/latest", "2017-12-01"),
},
{ 
	Display: "skus",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/skus", "2017-12-01"),
},
{ 
	Display: "virtualMachines",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines", "2017-12-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{instanceId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}", "2017-12-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}", "2017-12-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}", "2017-12-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "instanceView",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/instanceView", "2017-12-01"),
}, } ,
},
{ 
	Display: "publicipaddresses",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces/{networkInterfaceName}/ipconfigurations/{ipConfigurationName}/publicipaddresses", "2018-10-01"),
}, } ,
},
{ 
	Display: "publicipaddresses",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/publicipaddresses", "2018-10-01"),
}, } ,
	SubResources: []swagger.ResourceType{  } ,
},
{ 
	Display: "{publicIpAddressName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces/{networkInterfaceName}/ipconfigurations/{ipConfigurationName}/publicipaddresses/{publicIpAddressName}", "2018-10-01"),
}, } ,
},
{ 
	Display: "virtualMachines",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines", "2017-12-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{vmName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}", "2017-12-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}", "2017-12-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}", "2017-12-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}", "2017-12-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "extensions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/extensions", "2017-12-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{vmExtensionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/extensions/{vmExtensionName}", "2017-12-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/extensions/{vmExtensionName}", "2017-12-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/extensions/{vmExtensionName}", "2017-12-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/extensions/{vmExtensionName}", "2017-12-01"),
}, } ,
},
{ 
	Display: "instanceView",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/instanceView", "2017-12-01"),
},
{ 
	Display: "vmSizes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/vmSizes", "2017-12-01"),
},
{ 
	Display: "guestConfigurationAssignments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments", "2018-11-20"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{guestConfigurationAssignmentName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{guestConfigurationAssignmentName}", "2018-11-20"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{guestConfigurationAssignmentName}", "2018-11-20"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{guestConfigurationAssignmentName}", "2018-11-20"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "reports",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{guestConfigurationAssignmentName}/reports", "2018-11-20"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{reportId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{guestConfigurationAssignmentName}/reports/{reportId}", "2018-11-20"),
}, } ,
}, } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "reservationDetails",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Capacity/reservationorders/{reservationOrderId}/providers/Microsoft.Consumption/reservationDetails", "2019-05-01-preview"),
},
{ 
	Display: "reservationSummaries",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Capacity/reservationorders/{reservationOrderId}/providers/Microsoft.Consumption/reservationSummaries", "2019-05-01-preview"),
},
{ 
	Display: "reservationDetails",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Capacity/reservationorders/{reservationOrderId}/reservations/{reservationId}/providers/Microsoft.Consumption/reservationDetails", "2019-05-01-preview"),
},
{ 
	Display: "reservationSummaries",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Capacity/reservationorders/{reservationOrderId}/reservations/{reservationId}/providers/Microsoft.Consumption/reservationSummaries", "2019-05-01-preview"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Consumption/operations", "2019-05-01-preview"),
},
{ 
	Display: "aggregatedcost",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Billing/billingPeriods/{billingPeriodName}/Microsoft.Consumption/aggregatedcost", "2019-05-01-preview"),
},
{ 
	Display: "aggregatedcost",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Consumption/aggregatedcost", "2019-05-01-preview"),
},
{ 
	Display: "default",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Billing/billingPeriods/{billingPeriodName}/providers/Microsoft.Consumption/pricesheets/default", "2019-05-01-preview"),
},
{ 
	Display: "forecasts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Consumption/forecasts", "2019-05-01-preview"),
},
{ 
	Display: "default",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Consumption/pricesheets/default", "2019-05-01-preview"),
},
{ 
	Display: "reservationRecommendations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Consumption/reservationRecommendations", "2019-05-01-preview"),
},
{ 
	Display: "budgets",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Consumption/budgets", "2019-05-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{budgetName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Consumption/budgets/{budgetName}", "2019-05-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Consumption/budgets/{budgetName}", "2019-05-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Consumption/budgets/{budgetName}", "2019-05-01-preview"),
}, } ,
},
{ 
	Display: "charges",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Consumption/charges", "2019-05-01-preview"),
},
{ 
	Display: "marketplaces",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Consumption/marketplaces", "2019-05-01-preview"),
},
{ 
	Display: "tags",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Consumption/tags", "2019-05-01-preview"),
},
{ 
	Display: "usageDetails",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Consumption/usageDetails", "2019-05-01-preview"),
	Children: []swagger.ResourceType{  } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ContainerInstance/operations", "2018-10-01"),
},
{ 
	Display: "containerGroups",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ContainerInstance/containerGroups", "2018-10-01"),
},
{ 
	Display: "cachedImages",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ContainerInstance/locations/{location}/cachedImages", "2018-10-01"),
},
{ 
	Display: "capabilities",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ContainerInstance/locations/{location}/capabilities", "2018-10-01"),
},
{ 
	Display: "usages",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ContainerInstance/locations/{location}/usages", "2018-10-01"),
},
{ 
	Display: "containerGroups",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups", "2018-10-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{containerGroupName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}", "2018-10-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}", "2018-10-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}", "2018-10-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}", "2018-10-01"),
	Children: []swagger.ResourceType{  } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "logs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}/containers/{containerName}/logs", "2018-10-01"),
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ContainerRegistry/operations", "2019-05-01"),
},
{ 
	Display: "registries",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ContainerRegistry/registries", "2019-05-01"),
},
{ 
	Display: "registries",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries", "2019-05-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{registryName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}", "2019-05-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}", "2019-05-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}", "2019-05-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}", "2019-05-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "listUsages",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/listUsages", "2019-05-01"),
},
{ 
	Display: "replications",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/replications", "2019-05-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{replicationName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/replications/{replicationName}", "2019-05-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/replications/{replicationName}", "2019-05-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/replications/{replicationName}", "2019-05-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/replications/{replicationName}", "2019-05-01"),
}, } ,
},
{ 
	Display: "runs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/runs", "2019-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{runId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/runs/{runId}", "2019-04-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/runs/{runId}", "2019-04-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "tasks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/tasks", "2019-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{taskName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/tasks/{taskName}", "2019-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/tasks/{taskName}", "2019-04-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/tasks/{taskName}", "2019-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/tasks/{taskName}", "2019-04-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "webhooks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/webhooks", "2019-05-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{webhookName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/webhooks/{webhookName}", "2019-05-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/webhooks/{webhookName}", "2019-05-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/webhooks/{webhookName}", "2019-05-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/webhooks/{webhookName}", "2019-05-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ContainerService/operations", "2020-02-01"),
},
{ 
	Display: "managedClusters",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ContainerService/managedClusters", "2020-02-01"),
},
{ 
	Display: "managedClusters",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters", "2020-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{resourceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}", "2020-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}", "2020-02-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}", "2020-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}", "2020-02-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "agentPools",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/agentPools", "2020-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{agentPoolName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}", "2020-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}", "2020-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}", "2020-02-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "default",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}/upgradeProfiles/default", "2020-02-01"),
}, } ,
}, } ,
},
{ 
	Display: "availableAgentPoolVersions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/availableAgentPoolVersions", "2020-02-01"),
},
{ 
	Display: "default",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/upgradeProfiles/default", "2020-02-01"),
}, } ,
	SubResources: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.DocumentDB/operations", "2019-12-12"),
},
{ 
	Display: "databaseAccounts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/databaseAccounts", "2019-12-12"),
},
{ 
	Display: "databaseAccounts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts", "2019-12-12"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{accountName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}", "2019-12-12"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}", "2019-12-12"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}", "2019-12-12"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}", "2019-12-12"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "cassandraKeyspaces",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/cassandraKeyspaces", "2019-12-12"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{keyspaceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/cassandraKeyspaces/{keyspaceName}", "2019-12-12"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/cassandraKeyspaces/{keyspaceName}", "2019-12-12"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/cassandraKeyspaces/{keyspaceName}", "2019-12-12"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "tables",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/cassandraKeyspaces/{keyspaceName}/tables", "2019-12-12"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{tableName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/cassandraKeyspaces/{keyspaceName}/tables/{tableName}", "2019-12-12"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/cassandraKeyspaces/{keyspaceName}/tables/{tableName}", "2019-12-12"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/cassandraKeyspaces/{keyspaceName}/tables/{tableName}", "2019-12-12"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "default",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/cassandraKeyspaces/{keyspaceName}/tables/{tableName}/throughputSettings/default", "2019-12-12"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/cassandraKeyspaces/{keyspaceName}/tables/{tableName}/throughputSettings/default", "2019-12-12"),
}, } ,
}, } ,
},
{ 
	Display: "default",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/cassandraKeyspaces/{keyspaceName}/throughputSettings/default", "2019-12-12"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/cassandraKeyspaces/{keyspaceName}/throughputSettings/default", "2019-12-12"),
}, } ,
}, } ,
},
{ 
	Display: "gremlinDatabases",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/gremlinDatabases", "2019-12-12"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{databaseName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/gremlinDatabases/{databaseName}", "2019-12-12"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/gremlinDatabases/{databaseName}", "2019-12-12"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/gremlinDatabases/{databaseName}", "2019-12-12"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "graphs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/gremlinDatabases/{databaseName}/graphs", "2019-12-12"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{graphName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/gremlinDatabases/{databaseName}/graphs/{graphName}", "2019-12-12"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/gremlinDatabases/{databaseName}/graphs/{graphName}", "2019-12-12"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/gremlinDatabases/{databaseName}/graphs/{graphName}", "2019-12-12"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "default",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/gremlinDatabases/{databaseName}/graphs/{graphName}/throughputSettings/default", "2019-12-12"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/gremlinDatabases/{databaseName}/graphs/{graphName}/throughputSettings/default", "2019-12-12"),
}, } ,
}, } ,
},
{ 
	Display: "default",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/gremlinDatabases/{databaseName}/throughputSettings/default", "2019-12-12"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/gremlinDatabases/{databaseName}/throughputSettings/default", "2019-12-12"),
}, } ,
}, } ,
},
{ 
	Display: "metricDefinitions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/metricDefinitions", "2019-12-12"),
},
{ 
	Display: "metrics",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/metrics", "2019-12-12"),
},
{ 
	Display: "mongodbDatabases",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases", "2019-12-12"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{databaseName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}", "2019-12-12"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}", "2019-12-12"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}", "2019-12-12"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "collections",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}/collections", "2019-12-12"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{collectionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}/collections/{collectionName}", "2019-12-12"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}/collections/{collectionName}", "2019-12-12"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}/collections/{collectionName}", "2019-12-12"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "default",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}/collections/{collectionName}/throughputSettings/default", "2019-12-12"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}/collections/{collectionName}/throughputSettings/default", "2019-12-12"),
}, } ,
}, } ,
},
{ 
	Display: "default",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}/throughputSettings/default", "2019-12-12"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}/throughputSettings/default", "2019-12-12"),
}, } ,
}, } ,
},
{ 
	Display: "metrics",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/percentile/metrics", "2019-12-12"),
},
{ 
	Display: "privateEndpointConnections",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/privateEndpointConnections", "2019-08-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{privateEndpointConnectionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/privateEndpointConnections/{privateEndpointConnectionName}", "2019-08-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/privateEndpointConnections/{privateEndpointConnectionName}", "2019-08-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/privateEndpointConnections/{privateEndpointConnectionName}", "2019-08-01-preview"),
}, } ,
},
{ 
	Display: "privateLinkResources",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/privateLinkResources", "2019-08-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{groupName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/privateLinkResources/{groupName}", "2019-08-01-preview"),
}, } ,
},
{ 
	Display: "readonlykeys",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/readonlykeys", "2019-12-12"),
},
{ 
	Display: "sqlDatabases",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases", "2019-12-12"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{databaseName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}", "2019-12-12"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}", "2019-12-12"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}", "2019-12-12"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "containers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers", "2019-12-12"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{containerName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}", "2019-12-12"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}", "2019-12-12"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}", "2019-12-12"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "storedProcedures",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/storedProcedures", "2019-12-12"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{storedProcedureName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/storedProcedures/{storedProcedureName}", "2019-12-12"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/storedProcedures/{storedProcedureName}", "2019-12-12"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/storedProcedures/{storedProcedureName}", "2019-12-12"),
}, } ,
},
{ 
	Display: "default",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/throughputSettings/default", "2019-12-12"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/throughputSettings/default", "2019-12-12"),
},
{ 
	Display: "triggers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/triggers", "2019-12-12"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{triggerName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/triggers/{triggerName}", "2019-12-12"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/triggers/{triggerName}", "2019-12-12"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/triggers/{triggerName}", "2019-12-12"),
}, } ,
},
{ 
	Display: "userDefinedFunctions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/userDefinedFunctions", "2019-12-12"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{userDefinedFunctionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/userDefinedFunctions/{userDefinedFunctionName}", "2019-12-12"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/userDefinedFunctions/{userDefinedFunctionName}", "2019-12-12"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/userDefinedFunctions/{userDefinedFunctionName}", "2019-12-12"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "default",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/throughputSettings/default", "2019-12-12"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/throughputSettings/default", "2019-12-12"),
}, } ,
}, } ,
},
{ 
	Display: "tables",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables", "2019-12-12"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{tableName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}", "2019-12-12"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}", "2019-12-12"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}", "2019-12-12"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "default",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}/throughputSettings/default", "2019-12-12"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}/throughputSettings/default", "2019-12-12"),
}, } ,
}, } ,
},
{ 
	Display: "usages",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/usages", "2019-12-12"),
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "metricDefinitions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/collections/{collectionRid}/metricDefinitions", "2019-12-12"),
},
{ 
	Display: "metrics",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/collections/{collectionRid}/metrics", "2019-12-12"),
},
{ 
	Display: "metrics",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/collections/{collectionRid}/partitionKeyRangeId/{partitionKeyRangeId}/metrics", "2019-12-12"),
},
{ 
	Display: "metrics",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/collections/{collectionRid}/partitions/metrics", "2019-12-12"),
},
{ 
	Display: "usages",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/collections/{collectionRid}/partitions/usages", "2019-12-12"),
},
{ 
	Display: "usages",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/collections/{collectionRid}/usages", "2019-12-12"),
},
{ 
	Display: "metricDefinitions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/metricDefinitions", "2019-12-12"),
},
{ 
	Display: "metrics",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/metrics", "2019-12-12"),
},
{ 
	Display: "usages",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/usages", "2019-12-12"),
},
{ 
	Display: "metrics",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/region/{region}/databases/{databaseRid}/collections/{collectionRid}/metrics", "2019-12-12"),
},
{ 
	Display: "metrics",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/region/{region}/databases/{databaseRid}/collections/{collectionRid}/partitionKeyRangeId/{partitionKeyRangeId}/metrics", "2019-12-12"),
},
{ 
	Display: "metrics",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/region/{region}/databases/{databaseRid}/collections/{collectionRid}/partitions/metrics", "2019-12-12"),
},
{ 
	Display: "metrics",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/region/{region}/metrics", "2019-12-12"),
},
{ 
	Display: "metrics",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sourceRegion/{sourceRegion}/targetRegion/{targetRegion}/percentile/metrics", "2019-12-12"),
},
{ 
	Display: "metrics",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/targetRegion/{targetRegion}/percentile/metrics", "2019-12-12"),
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.CostManagement/operations", "2019-04-01-preview"),
},
{ 
	Display: "views",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.CostManagement/views", "2019-04-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{viewName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.CostManagement/views/{viewName}", "2019-04-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.CostManagement/views/{viewName}", "2019-04-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.CostManagement/views/{viewName}", "2019-04-01-preview"),
}, } ,
},
{ 
	Display: "budgets",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.CostManagement/budgets", "2019-04-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{budgetName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.CostManagement/budgets/{budgetName}", "2019-04-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.CostManagement/budgets/{budgetName}", "2019-04-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.CostManagement/budgets/{budgetName}", "2019-04-01-preview"),
}, } ,
},
{ 
	Display: "views",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.CostManagement/views", "2019-04-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{viewName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.CostManagement/views/{viewName}", "2019-04-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.CostManagement/views/{viewName}", "2019-04-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.CostManagement/views/{viewName}", "2019-04-01-preview"),
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.CustomerInsights/operations", "2017-04-26"),
},
{ 
	Display: "hubs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.CustomerInsights/hubs", "2017-04-26"),
},
{ 
	Display: "hubs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs", "2017-04-26"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{hubName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}", "2017-04-26"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}", "2017-04-26"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}", "2017-04-26"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}", "2017-04-26"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "authorizationPolicies",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/authorizationPolicies", "2017-04-26"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{authorizationPolicyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/authorizationPolicies/{authorizationPolicyName}", "2017-04-26"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/authorizationPolicies/{authorizationPolicyName}", "2017-04-26"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "connectors",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/connectors", "2017-04-26"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{connectorName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/connectors/{connectorName}", "2017-04-26"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/connectors/{connectorName}", "2017-04-26"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/connectors/{connectorName}", "2017-04-26"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "mappings",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/connectors/{connectorName}/mappings", "2017-04-26"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{mappingName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/connectors/{connectorName}/mappings/{mappingName}", "2017-04-26"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/connectors/{connectorName}/mappings/{mappingName}", "2017-04-26"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/connectors/{connectorName}/mappings/{mappingName}", "2017-04-26"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "interactions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/interactions", "2017-04-26"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{interactionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/interactions/{interactionName}", "2017-04-26"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/interactions/{interactionName}", "2017-04-26"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "kpi",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/kpi", "2017-04-26"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{kpiName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/kpi/{kpiName}", "2017-04-26"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/kpi/{kpiName}", "2017-04-26"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/kpi/{kpiName}", "2017-04-26"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "links",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/links", "2017-04-26"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{linkName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/links/{linkName}", "2017-04-26"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/links/{linkName}", "2017-04-26"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/links/{linkName}", "2017-04-26"),
}, } ,
},
{ 
	Display: "predictions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/predictions", "2017-04-26"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{predictionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/predictions/{predictionName}", "2017-04-26"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/predictions/{predictionName}", "2017-04-26"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/predictions/{predictionName}", "2017-04-26"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "profiles",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/profiles", "2017-04-26"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{profileName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/profiles/{profileName}", "2017-04-26"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/profiles/{profileName}", "2017-04-26"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/profiles/{profileName}", "2017-04-26"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "relationshipLinks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationshipLinks", "2017-04-26"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{relationshipLinkName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationshipLinks/{relationshipLinkName}", "2017-04-26"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationshipLinks/{relationshipLinkName}", "2017-04-26"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationshipLinks/{relationshipLinkName}", "2017-04-26"),
}, } ,
},
{ 
	Display: "relationships",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationships", "2017-04-26"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{relationshipName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationships/{relationshipName}", "2017-04-26"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationships/{relationshipName}", "2017-04-26"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationships/{relationshipName}", "2017-04-26"),
}, } ,
},
{ 
	Display: "roleAssignments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/roleAssignments", "2017-04-26"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{assignmentName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/roleAssignments/{assignmentName}", "2017-04-26"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/roleAssignments/{assignmentName}", "2017-04-26"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/roleAssignments/{assignmentName}", "2017-04-26"),
}, } ,
},
{ 
	Display: "roles",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/roles", "2017-04-26"),
},
{ 
	Display: "views",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/views", "2017-04-26"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{viewName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/views/{viewName}", "2017-04-26"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/views/{viewName}", "2017-04-26"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/views/{viewName}", "2017-04-26"),
}, } ,
},
{ 
	Display: "widgetTypes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/widgetTypes", "2017-04-26"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{widgetTypeName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/widgetTypes/{widgetTypeName}", "2017-04-26"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.CustomerLockbox/operations", "2018-02-28-preview"),
},
{ 
	Display: "requests",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.CustomerLockbox/requests", "2018-02-28-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{requestId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.CustomerLockbox/requests/{requestId}", "2018-02-28-preview"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.CustomProviders/operations", "2018-09-01-preview"),
},
{ 
	Display: "resourceProviders",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.CustomProviders/resourceProviders", "2018-09-01-preview"),
},
{ 
	Display: "resourceProviders",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders", "2018-09-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{resourceProviderName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/{resourceProviderName}", "2018-09-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/{resourceProviderName}", "2018-09-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/{resourceProviderName}", "2018-09-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomProviders/resourceProviders/{resourceProviderName}", "2018-09-01-preview"),
}, } ,
},
{ 
	Display: "associations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.CustomProviders/associations", "2018-09-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{associationName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.CustomProviders/associations/{associationName}", "2018-09-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.CustomProviders/associations/{associationName}", "2018-09-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.CustomProviders/associations/{associationName}", "2018-09-01-preview"),
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.DataBox/operations", "2019-09-01"),
},
{ 
	Display: "jobs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DataBox/jobs", "2019-09-01"),
},
{ 
	Display: "jobs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBox/jobs", "2019-09-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{jobName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBox/jobs/{jobName}", "2019-09-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBox/jobs/{jobName}", "2019-09-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBox/jobs/{jobName}", "2019-09-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBox/jobs/{jobName}", "2019-09-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.DataBoxEdge/operations", "2019-08-01"),
},
{ 
	Display: "dataBoxEdgeDevices",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices", "2019-08-01"),
},
{ 
	Display: "skus",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DataBoxEdge/skus", "2019-08-01"),
},
{ 
	Display: "dataBoxEdgeDevices",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices", "2019-08-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{deviceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}", "2019-08-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}", "2019-08-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}", "2019-08-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}", "2019-08-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "alerts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/alerts", "2019-08-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/alerts/{name}", "2019-08-01"),
}, } ,
},
{ 
	Display: "bandwidthSchedules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/bandwidthSchedules", "2019-08-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/bandwidthSchedules/{name}", "2019-08-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/bandwidthSchedules/{name}", "2019-08-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/bandwidthSchedules/{name}", "2019-08-01"),
}, } ,
},
{ 
	Display: "default",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/networkSettings/default", "2019-08-01"),
},
{ 
	Display: "nodes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/nodes", "2019-08-01"),
},
{ 
	Display: "orders",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/orders", "2019-08-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "default",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/orders/default", "2019-08-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/orders/default", "2019-08-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/orders/default", "2019-08-01"),
}, } ,
},
{ 
	Display: "roles",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles", "2019-08-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles/{name}", "2019-08-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles/{name}", "2019-08-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles/{name}", "2019-08-01"),
}, } ,
},
{ 
	Display: "shares",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/shares", "2019-08-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/shares/{name}", "2019-08-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/shares/{name}", "2019-08-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/shares/{name}", "2019-08-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "storageAccountCredentials",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccountCredentials", "2019-08-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccountCredentials/{name}", "2019-08-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccountCredentials/{name}", "2019-08-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccountCredentials/{name}", "2019-08-01"),
}, } ,
},
{ 
	Display: "storageAccounts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccounts", "2019-08-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{storageAccountName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccounts/{storageAccountName}", "2019-08-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccounts/{storageAccountName}", "2019-08-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccounts/{storageAccountName}", "2019-08-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "containers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccounts/{storageAccountName}/containers", "2019-08-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{containerName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccounts/{storageAccountName}/containers/{containerName}", "2019-08-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccounts/{storageAccountName}/containers/{containerName}", "2019-08-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccounts/{storageAccountName}/containers/{containerName}", "2019-08-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "triggers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/triggers", "2019-08-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/triggers/{name}", "2019-08-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/triggers/{name}", "2019-08-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/triggers/{name}", "2019-08-01"),
}, } ,
},
{ 
	Display: "default",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/updateSummary/default", "2019-08-01"),
},
{ 
	Display: "users",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/users", "2019-08-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/users/{name}", "2019-08-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/users/{name}", "2019-08-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/users/{name}", "2019-08-01"),
}, } ,
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/jobs/{name}", "2019-08-01"),
},
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/operationsStatus/{name}", "2019-08-01"),
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Databricks/operations", "2018-04-01"),
},
{ 
	Display: "workspaces",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Databricks/workspaces", "2018-04-01"),
},
{ 
	Display: "workspaces",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Databricks/workspaces", "2018-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{workspaceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Databricks/workspaces/{workspaceName}", "2018-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Databricks/workspaces/{workspaceName}", "2018-04-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Databricks/workspaces/{workspaceName}", "2018-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Databricks/workspaces/{workspaceName}", "2018-04-01"),
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.DataCatalog/operations", "2016-03-30"),
},
{ 
	Display: "catalogs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataCatalog/catalogs", "2016-03-30"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{catalogName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataCatalog/catalogs/{catalogName}", "2016-03-30"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataCatalog/catalogs/{catalogName}", "2016-03-30"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataCatalog/catalogs/{catalogName}", "2016-03-30"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataCatalog/catalogs/{catalogName}", "2016-03-30"),
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.DataFactory/operations", "2018-06-01"),
},
{ 
	Display: "factories",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DataFactory/factories", "2018-06-01"),
},
{ 
	Display: "factories",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories", "2018-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{factoryName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}", "2018-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}", "2018-06-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}", "2018-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}", "2018-06-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "dataflows",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/dataflows", "2018-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{dataFlowName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/dataflows/{dataFlowName}", "2018-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/dataflows/{dataFlowName}", "2018-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/dataflows/{dataFlowName}", "2018-06-01"),
}, } ,
},
{ 
	Display: "datasets",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/datasets", "2018-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{datasetName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/datasets/{datasetName}", "2018-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/datasets/{datasetName}", "2018-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/datasets/{datasetName}", "2018-06-01"),
}, } ,
},
{ 
	Display: "integrationRuntimes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/integrationRuntimes", "2018-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{integrationRuntimeName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/integrationRuntimes/{integrationRuntimeName}", "2018-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/integrationRuntimes/{integrationRuntimeName}", "2018-06-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/integrationRuntimes/{integrationRuntimeName}", "2018-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/integrationRuntimes/{integrationRuntimeName}", "2018-06-01"),
	Children: []swagger.ResourceType{  } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{nodeName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/integrationRuntimes/{integrationRuntimeName}/nodes/{nodeName}", "2018-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/integrationRuntimes/{integrationRuntimeName}/nodes/{nodeName}", "2018-06-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/integrationRuntimes/{integrationRuntimeName}/nodes/{nodeName}", "2018-06-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
}, } ,
},
{ 
	Display: "linkedservices",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/linkedservices", "2018-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{linkedServiceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/linkedservices/{linkedServiceName}", "2018-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/linkedservices/{linkedServiceName}", "2018-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/linkedservices/{linkedServiceName}", "2018-06-01"),
}, } ,
},
{ 
	Display: "pipelines",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelines", "2018-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{pipelineName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelines/{pipelineName}", "2018-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelines/{pipelineName}", "2018-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelines/{pipelineName}", "2018-06-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "triggers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/triggers", "2018-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{triggerName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/triggers/{triggerName}", "2018-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/triggers/{triggerName}", "2018-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/triggers/{triggerName}", "2018-06-01"),
	Children: []swagger.ResourceType{  } ,
	SubResources: []swagger.ResourceType{  } ,
}, } ,
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{runId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelineruns/{runId}", "2018-06-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.DataLakeAnalytics/operations", "2016-11-01"),
},
{ 
	Display: "accounts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DataLakeAnalytics/accounts", "2016-11-01"),
},
{ 
	Display: "capability",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DataLakeAnalytics/locations/{location}/capability", "2016-11-01"),
},
{ 
	Display: "accounts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts", "2016-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{accountName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}", "2016-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}", "2016-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}", "2016-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}", "2016-11-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "computePolicies",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/computePolicies", "2016-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{computePolicyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/computePolicies/{computePolicyName}", "2016-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/computePolicies/{computePolicyName}", "2016-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/computePolicies/{computePolicyName}", "2016-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/computePolicies/{computePolicyName}", "2016-11-01"),
}, } ,
},
{ 
	Display: "dataLakeStoreAccounts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/dataLakeStoreAccounts", "2016-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{dataLakeStoreAccountName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/dataLakeStoreAccounts/{dataLakeStoreAccountName}", "2016-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/dataLakeStoreAccounts/{dataLakeStoreAccountName}", "2016-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/dataLakeStoreAccounts/{dataLakeStoreAccountName}", "2016-11-01"),
}, } ,
},
{ 
	Display: "firewallRules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/firewallRules", "2016-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{firewallRuleName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/firewallRules/{firewallRuleName}", "2016-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/firewallRules/{firewallRuleName}", "2016-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/firewallRules/{firewallRuleName}", "2016-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/firewallRules/{firewallRuleName}", "2016-11-01"),
}, } ,
},
{ 
	Display: "storageAccounts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts", "2016-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{storageAccountName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}", "2016-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}", "2016-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}", "2016-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}", "2016-11-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "containers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}/containers", "2016-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{containerName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}/containers/{containerName}", "2016-11-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
}, } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.DataLakeStore/operations", "2016-11-01"),
},
{ 
	Display: "accounts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DataLakeStore/accounts", "2016-11-01"),
},
{ 
	Display: "capability",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DataLakeStore/locations/{location}/capability", "2016-11-01"),
},
{ 
	Display: "usages",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DataLakeStore/locations/{location}/usages", "2016-11-01"),
},
{ 
	Display: "accounts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts", "2016-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{accountName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}", "2016-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}", "2016-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}", "2016-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}", "2016-11-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "firewallRules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/firewallRules", "2016-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{firewallRuleName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/firewallRules/{firewallRuleName}", "2016-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/firewallRules/{firewallRuleName}", "2016-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/firewallRules/{firewallRuleName}", "2016-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/firewallRules/{firewallRuleName}", "2016-11-01"),
}, } ,
},
{ 
	Display: "trustedIdProviders",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/trustedIdProviders", "2016-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{trustedIdProviderName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/trustedIdProviders/{trustedIdProviderName}", "2016-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/trustedIdProviders/{trustedIdProviderName}", "2016-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/trustedIdProviders/{trustedIdProviderName}", "2016-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/trustedIdProviders/{trustedIdProviderName}", "2016-11-01"),
}, } ,
},
{ 
	Display: "virtualNetworkRules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/virtualNetworkRules", "2016-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{virtualNetworkRuleName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/virtualNetworkRules/{virtualNetworkRuleName}", "2016-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/virtualNetworkRules/{virtualNetworkRuleName}", "2016-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/virtualNetworkRules/{virtualNetworkRuleName}", "2016-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/virtualNetworkRules/{virtualNetworkRuleName}", "2016-11-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.DataMigration/operations", "2018-04-19"),
},
{ 
	Display: "usages",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DataMigration/locations/{location}/usages", "2018-04-19"),
},
{ 
	Display: "services",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DataMigration/services", "2018-04-19"),
},
{ 
	Display: "skus",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DataMigration/skus", "2018-04-19"),
},
{ 
	Display: "services",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services", "2018-04-19"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{serviceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}", "2018-04-19"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}", "2018-04-19"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}", "2018-04-19"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}", "2018-04-19"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "projects",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects", "2018-04-19"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{projectName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}", "2018-04-19"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}", "2018-04-19"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}", "2018-04-19"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}", "2018-04-19"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "tasks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/tasks", "2018-04-19"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{taskName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/tasks/{taskName}", "2018-04-19"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/tasks/{taskName}", "2018-04-19"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/tasks/{taskName}", "2018-04-19"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/tasks/{taskName}", "2018-04-19"),
	Children: []swagger.ResourceType{  } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "skus",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/skus", "2018-04-19"),
}, } ,
}, } ,
},
{ 
	Display: "ListInvitations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.DataShare/ListInvitations", "2019-11-01"),
},
{ 
	Display: "{invitationId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.DataShare/locations/{location}/consumerInvitations/{invitationId}", "2019-11-01"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.DataShare/operations", "2019-11-01"),
},
{ 
	Display: "accounts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DataShare/accounts", "2019-11-01"),
},
{ 
	Display: "accounts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{accountName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}", "2019-11-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "shareSubscriptions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shareSubscriptions", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{shareSubscriptionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shareSubscriptions/{shareSubscriptionName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shareSubscriptions/{shareSubscriptionName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shareSubscriptions/{shareSubscriptionName}", "2019-11-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "ConsumerSourceDataSets",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shareSubscriptions/{shareSubscriptionName}/ConsumerSourceDataSets", "2019-11-01"),
},
{ 
	Display: "dataSetMappings",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shareSubscriptions/{shareSubscriptionName}/dataSetMappings", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{dataSetMappingName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shareSubscriptions/{shareSubscriptionName}/dataSetMappings/{dataSetMappingName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shareSubscriptions/{shareSubscriptionName}/dataSetMappings/{dataSetMappingName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shareSubscriptions/{shareSubscriptionName}/dataSetMappings/{dataSetMappingName}", "2019-11-01"),
}, } ,
},
{ 
	Display: "triggers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shareSubscriptions/{shareSubscriptionName}/triggers", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{triggerName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shareSubscriptions/{shareSubscriptionName}/triggers/{triggerName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shareSubscriptions/{shareSubscriptionName}/triggers/{triggerName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shareSubscriptions/{shareSubscriptionName}/triggers/{triggerName}", "2019-11-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "shares",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{shareName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}", "2019-11-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "dataSets",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/dataSets", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{dataSetName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/dataSets/{dataSetName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/dataSets/{dataSetName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/dataSets/{dataSetName}", "2019-11-01"),
}, } ,
},
{ 
	Display: "invitations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/invitations", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{invitationName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/invitations/{invitationName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/invitations/{invitationName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/invitations/{invitationName}", "2019-11-01"),
}, } ,
},
{ 
	Display: "providerShareSubscriptions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/providerShareSubscriptions", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{providerShareSubscriptionId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/providerShareSubscriptions/{providerShareSubscriptionId}", "2019-11-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "synchronizationSettings",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/synchronizationSettings", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{synchronizationSettingName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/synchronizationSettings/{synchronizationSettingName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/synchronizationSettings/{synchronizationSettingName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/synchronizationSettings/{synchronizationSettingName}", "2019-11-01"),
}, } ,
}, } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.DeploymentManager/operations", "2019-11-01-preview"),
},
{ 
	Display: "artifactSources",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/artifactSources", "2019-11-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{artifactSourceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/artifactSources/{artifactSourceName}", "2019-11-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/artifactSources/{artifactSourceName}", "2019-11-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/artifactSources/{artifactSourceName}", "2019-11-01-preview"),
}, } ,
},
{ 
	Display: "rollouts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/rollouts", "2019-11-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{rolloutName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/rollouts/{rolloutName}", "2019-11-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/rollouts/{rolloutName}", "2019-11-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/rollouts/{rolloutName}", "2019-11-01-preview"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "serviceTopologies",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/serviceTopologies", "2019-11-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{serviceTopologyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/serviceTopologies/{serviceTopologyName}", "2019-11-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/serviceTopologies/{serviceTopologyName}", "2019-11-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/serviceTopologies/{serviceTopologyName}", "2019-11-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "services",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/serviceTopologies/{serviceTopologyName}/services", "2019-11-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{serviceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/serviceTopologies/{serviceTopologyName}/services/{serviceName}", "2019-11-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/serviceTopologies/{serviceTopologyName}/services/{serviceName}", "2019-11-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/serviceTopologies/{serviceTopologyName}/services/{serviceName}", "2019-11-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "serviceUnits",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/serviceTopologies/{serviceTopologyName}/services/{serviceName}/serviceUnits", "2019-11-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{serviceUnitName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/serviceTopologies/{serviceTopologyName}/services/{serviceName}/serviceUnits/{serviceUnitName}", "2019-11-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/serviceTopologies/{serviceTopologyName}/services/{serviceName}/serviceUnits/{serviceUnitName}", "2019-11-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/serviceTopologies/{serviceTopologyName}/services/{serviceName}/serviceUnits/{serviceUnitName}", "2019-11-01-preview"),
}, } ,
}, } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "steps",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/steps", "2019-11-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{stepName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/steps/{stepName}", "2019-11-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/steps/{stepName}", "2019-11-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/steps/{stepName}", "2019-11-01-preview"),
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Devices/operations", "2019-07-01-preview"),
},
{ 
	Display: "provisioningServices",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Devices/provisioningServices", "2018-01-22"),
},
{ 
	Display: "provisioningServices",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices", "2018-01-22"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{provisioningServiceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}", "2018-01-22"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}", "2018-01-22"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}", "2018-01-22"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}", "2018-01-22"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "certificates",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}/certificates", "2018-01-22"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{certificateName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}/certificates/{certificateName}", "2018-01-22"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}/certificates/{certificateName}", "2018-01-22"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}/certificates/{certificateName}", "2018-01-22"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "skus",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}/skus", "2018-01-22"),
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{operationId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}/operationresults/{operationId}", "2018-01-22"),
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.DevOps/operations", "2019-07-01-preview"),
},
{ 
	Display: "pipelineTemplateDefinitions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.DevOps/pipelineTemplateDefinitions", "2019-07-01-preview"),
},
{ 
	Display: "pipelines",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DevOps/pipelines", "2019-07-01-preview"),
},
{ 
	Display: "pipelines",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevOps/pipelines", "2019-07-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{pipelineName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevOps/pipelines/{pipelineName}", "2019-07-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevOps/pipelines/{pipelineName}", "2019-07-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevOps/pipelines/{pipelineName}", "2019-07-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevOps/pipelines/{pipelineName}", "2019-07-01-preview"),
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.DevSpaces/operations", "2019-04-01"),
},
{ 
	Display: "controllers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DevSpaces/controllers", "2019-04-01"),
},
{ 
	Display: "controllers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevSpaces/controllers", "2019-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevSpaces/controllers/{name}", "2019-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevSpaces/controllers/{name}", "2019-04-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevSpaces/controllers/{name}", "2019-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevSpaces/controllers/{name}", "2019-04-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.DevTestLab/operations", "2018-09-15"),
},
{ 
	Display: "labs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DevTestLab/labs", "2018-09-15"),
},
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DevTestLab/locations/{locationName}/operations/{name}", "2018-09-15"),
},
{ 
	Display: "schedules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DevTestLab/schedules", "2018-09-15"),
},
{ 
	Display: "labs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs", "2018-09-15"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "artifactsources",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources", "2018-09-15"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "armtemplates",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{artifactSourceName}/armtemplates", "2018-09-15"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{artifactSourceName}/armtemplates/{name}", "2018-09-15"),
}, } ,
},
{ 
	Display: "artifacts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{artifactSourceName}/artifacts", "2018-09-15"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{artifactSourceName}/artifacts/{name}", "2018-09-15"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{name}", "2018-09-15"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{name}", "2018-09-15"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{name}", "2018-09-15"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{name}", "2018-09-15"),
}, } ,
},
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/costs/{name}", "2018-09-15"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/costs/{name}", "2018-09-15"),
},
{ 
	Display: "customimages",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/customimages", "2018-09-15"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/customimages/{name}", "2018-09-15"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/customimages/{name}", "2018-09-15"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/customimages/{name}", "2018-09-15"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/customimages/{name}", "2018-09-15"),
}, } ,
},
{ 
	Display: "formulas",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/formulas", "2018-09-15"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/formulas/{name}", "2018-09-15"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/formulas/{name}", "2018-09-15"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/formulas/{name}", "2018-09-15"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/formulas/{name}", "2018-09-15"),
}, } ,
},
{ 
	Display: "galleryimages",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/galleryimages", "2018-09-15"),
},
{ 
	Display: "notificationchannels",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/notificationchannels", "2018-09-15"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/notificationchannels/{name}", "2018-09-15"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/notificationchannels/{name}", "2018-09-15"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/notificationchannels/{name}", "2018-09-15"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/notificationchannels/{name}", "2018-09-15"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "policies",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/policysets/{policySetName}/policies", "2018-09-15"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/policysets/{policySetName}/policies/{name}", "2018-09-15"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/policysets/{policySetName}/policies/{name}", "2018-09-15"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/policysets/{policySetName}/policies/{name}", "2018-09-15"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/policysets/{policySetName}/policies/{name}", "2018-09-15"),
}, } ,
},
{ 
	Display: "schedules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/schedules", "2018-09-15"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/schedules/{name}", "2018-09-15"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/schedules/{name}", "2018-09-15"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/schedules/{name}", "2018-09-15"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/schedules/{name}", "2018-09-15"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/servicerunners/{name}", "2018-09-15"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/servicerunners/{name}", "2018-09-15"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/servicerunners/{name}", "2018-09-15"),
},
{ 
	Display: "users",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users", "2018-09-15"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{name}", "2018-09-15"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{name}", "2018-09-15"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{name}", "2018-09-15"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{name}", "2018-09-15"),
},
{ 
	Display: "disks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/disks", "2018-09-15"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/disks/{name}", "2018-09-15"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/disks/{name}", "2018-09-15"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/disks/{name}", "2018-09-15"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/disks/{name}", "2018-09-15"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "environments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/environments", "2018-09-15"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/environments/{name}", "2018-09-15"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/environments/{name}", "2018-09-15"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/environments/{name}", "2018-09-15"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/environments/{name}", "2018-09-15"),
}, } ,
},
{ 
	Display: "secrets",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/secrets", "2018-09-15"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/secrets/{name}", "2018-09-15"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/secrets/{name}", "2018-09-15"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/secrets/{name}", "2018-09-15"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/secrets/{name}", "2018-09-15"),
}, } ,
},
{ 
	Display: "servicefabrics",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics", "2018-09-15"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{name}", "2018-09-15"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{name}", "2018-09-15"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{name}", "2018-09-15"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{name}", "2018-09-15"),
	Children: []swagger.ResourceType{  } ,
},
{ 
	Display: "schedules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{serviceFabricName}/schedules", "2018-09-15"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{serviceFabricName}/schedules/{name}", "2018-09-15"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{serviceFabricName}/schedules/{name}", "2018-09-15"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{serviceFabricName}/schedules/{name}", "2018-09-15"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{serviceFabricName}/schedules/{name}", "2018-09-15"),
	Children: []swagger.ResourceType{  } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "virtualmachines",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines", "2018-09-15"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{name}", "2018-09-15"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{name}", "2018-09-15"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{name}", "2018-09-15"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{name}", "2018-09-15"),
	Children: []swagger.ResourceType{  } ,
},
{ 
	Display: "schedules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{virtualMachineName}/schedules", "2018-09-15"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{virtualMachineName}/schedules/{name}", "2018-09-15"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{virtualMachineName}/schedules/{name}", "2018-09-15"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{virtualMachineName}/schedules/{name}", "2018-09-15"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{virtualMachineName}/schedules/{name}", "2018-09-15"),
	Children: []swagger.ResourceType{  } ,
}, } ,
}, } ,
},
{ 
	Display: "virtualnetworks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualnetworks", "2018-09-15"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualnetworks/{name}", "2018-09-15"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualnetworks/{name}", "2018-09-15"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualnetworks/{name}", "2018-09-15"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualnetworks/{name}", "2018-09-15"),
}, } ,
},
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{name}", "2018-09-15"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{name}", "2018-09-15"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{name}", "2018-09-15"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{name}", "2018-09-15"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "schedules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/schedules", "2018-09-15"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/schedules/{name}", "2018-09-15"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/schedules/{name}", "2018-09-15"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/schedules/{name}", "2018-09-15"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/schedules/{name}", "2018-09-15"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.DigitalTwins/operations", "2020-03-01-preview"),
},
{ 
	Display: "digitalTwinsInstances",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DigitalTwins/digitalTwinsInstances", "2020-03-01-preview"),
},
{ 
	Display: "digitalTwinsInstances",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DigitalTwins/digitalTwinsInstances", "2020-03-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{resourceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{resourceName}", "2020-03-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{resourceName}", "2020-03-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{resourceName}", "2020-03-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{resourceName}", "2020-03-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "endpoints",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{resourceName}/endpoints", "2020-03-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{endpointName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{resourceName}/endpoints/{endpointName}", "2020-03-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{resourceName}/endpoints/{endpointName}", "2020-03-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{resourceName}/endpoints/{endpointName}", "2020-03-01-preview"),
}, } ,
},
{ 
	Display: "integrationResources",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{resourceName}/integrationResources", "2020-03-01-preview"),
}, } ,
}, } ,
},
{ 
	Display: "{integrationResourceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.DigitalTwins/integrationResources/{integrationResourceName}", "2020-03-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.DigitalTwins/integrationResources/{integrationResourceName}", "2020-03-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.DigitalTwins/integrationResources/{integrationResourceName}", "2020-03-01-preview"),
},
{ 
	Display: "dnszones",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/dnszones", "2016-04-01"),
},
{ 
	Display: "dnsZones",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones", "2016-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{zoneName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}", "2016-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}", "2016-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}", "2016-04-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "recordsets",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/recordsets", "2016-04-01"),
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{recordType}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/{recordType}", "2016-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{relativeRecordSetName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/{recordType}/{relativeRecordSetName}", "2016-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/{recordType}/{relativeRecordSetName}", "2016-04-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/{recordType}/{relativeRecordSetName}", "2016-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/{recordType}/{relativeRecordSetName}", "2016-04-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.AAD/operations", "2020-01-01"),
},
{ 
	Display: "domainServices",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.AAD/domainServices", "2020-01-01"),
},
{ 
	Display: "domainServices",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices", "2020-01-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{domainServiceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices/{domainServiceName}", "2020-01-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices/{domainServiceName}", "2020-01-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices/{domainServiceName}", "2020-01-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices/{domainServiceName}", "2020-01-01"),
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.EngagementFabric/operations", "2018-09-01-preview"),
},
{ 
	Display: "Accounts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.EngagementFabric/Accounts", "2018-09-01-preview"),
},
{ 
	Display: "skus",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.EngagementFabric/skus", "2018-09-01-preview"),
},
{ 
	Display: "Accounts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EngagementFabric/Accounts", "2018-09-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{accountName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EngagementFabric/Accounts/{accountName}", "2018-09-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EngagementFabric/Accounts/{accountName}", "2018-09-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EngagementFabric/Accounts/{accountName}", "2018-09-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EngagementFabric/Accounts/{accountName}", "2018-09-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "Channels",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EngagementFabric/Accounts/{accountName}/Channels", "2018-09-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{channelName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EngagementFabric/Accounts/{accountName}/Channels/{channelName}", "2018-09-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EngagementFabric/Accounts/{accountName}/Channels/{channelName}", "2018-09-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EngagementFabric/Accounts/{accountName}/Channels/{channelName}", "2018-09-01-preview"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.EventGrid/operations", "2019-06-01"),
},
{ 
	Display: "topicTypes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.EventGrid/topicTypes", "2019-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{topicTypeName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.EventGrid/topicTypes/{topicTypeName}", "2019-06-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "eventTypes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.EventGrid/topicTypes/{topicTypeName}/eventTypes", "2019-06-01"),
}, } ,
}, } ,
},
{ 
	Display: "domains",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.EventGrid/domains", "2019-06-01"),
},
{ 
	Display: "eventSubscriptions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.EventGrid/eventSubscriptions", "2019-06-01"),
},
{ 
	Display: "eventSubscriptions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.EventGrid/locations/{location}/eventSubscriptions", "2019-06-01"),
},
{ 
	Display: "eventSubscriptions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.EventGrid/locations/{location}/topicTypes/{topicTypeName}/eventSubscriptions", "2019-06-01"),
},
{ 
	Display: "eventSubscriptions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.EventGrid/topicTypes/{topicTypeName}/eventSubscriptions", "2019-06-01"),
},
{ 
	Display: "topics",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.EventGrid/topics", "2019-06-01"),
},
{ 
	Display: "domains",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains", "2019-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{domainName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}", "2019-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}", "2019-06-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}", "2019-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}", "2019-06-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "topics",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/topics", "2019-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{domainTopicName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/topics/{domainTopicName}", "2019-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/topics/{domainTopicName}", "2019-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/topics/{domainTopicName}", "2019-06-01"),
},
{ 
	Display: "eventSubscriptions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/topics/{topicName}/providers/Microsoft.EventGrid/eventSubscriptions", "2019-06-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "eventSubscriptions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/eventSubscriptions", "2019-06-01"),
},
{ 
	Display: "eventSubscriptions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/locations/{location}/eventSubscriptions", "2019-06-01"),
},
{ 
	Display: "eventSubscriptions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/locations/{location}/topicTypes/{topicTypeName}/eventSubscriptions", "2019-06-01"),
},
{ 
	Display: "eventSubscriptions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topicTypes/{topicTypeName}/eventSubscriptions", "2019-06-01"),
},
{ 
	Display: "topics",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics", "2019-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{topicName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics/{topicName}", "2019-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics/{topicName}", "2019-06-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics/{topicName}", "2019-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics/{topicName}", "2019-06-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "eventSubscriptions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{providerNamespace}/{resourceTypeName}/{resourceName}/providers/Microsoft.EventGrid/eventSubscriptions", "2019-06-01"),
},
{ 
	Display: "eventTypes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{providerNamespace}/{resourceTypeName}/{resourceName}/providers/Microsoft.EventGrid/eventTypes", "2019-06-01"),
},
{ 
	Display: "{eventSubscriptionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.EventGrid/eventSubscriptions/{eventSubscriptionName}", "2019-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.EventGrid/eventSubscriptions/{eventSubscriptionName}", "2019-06-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.EventGrid/eventSubscriptions/{eventSubscriptionName}", "2019-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.EventGrid/eventSubscriptions/{eventSubscriptionName}", "2019-06-01"),
	Children: []swagger.ResourceType{  } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.EventHub/operations", "2017-04-01"),
},
{ 
	Display: "namespaces",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.EventHub/namespaces", "2017-04-01"),
},
{ 
	Display: "regions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.EventHub/sku/{sku}/regions", "2017-04-01"),
},
{ 
	Display: "namespaces",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces", "2017-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{namespaceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}", "2017-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}", "2017-04-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}", "2017-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}", "2017-04-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "AuthorizationRules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/AuthorizationRules", "2017-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{authorizationRuleName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "disasterRecoveryConfigs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs", "2017-04-01"),
	Children: []swagger.ResourceType{  } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{alias}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}", "2017-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}", "2017-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}", "2017-04-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "AuthorizationRules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}/AuthorizationRules", "2017-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{authorizationRuleName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "eventhubs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs", "2017-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{eventHubName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}", "2017-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}", "2017-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}", "2017-04-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "authorizationRules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/authorizationRules", "2017-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{authorizationRuleName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "consumergroups",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/consumergroups", "2017-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{consumerGroupName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/consumergroups/{consumerGroupName}", "2017-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/consumergroups/{consumerGroupName}", "2017-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/consumergroups/{consumerGroupName}", "2017-04-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "messagingplan",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/messagingplan", "2017-04-01"),
},
{ 
	Display: "networkRuleSets",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/networkRuleSets", "2017-04-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "default",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/networkRuleSets/default", "2017-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/networkRuleSets/default", "2017-04-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "FrontDoorWebApplicationFirewallManagedRuleSets",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/FrontDoorWebApplicationFirewallManagedRuleSets", "2019-10-01"),
},
{ 
	Display: "NetworkExperimentProfiles",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/NetworkExperimentProfiles", "2019-11-01"),
},
{ 
	Display: "frontDoors",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/frontDoors", "2020-01-01"),
},
{ 
	Display: "FrontDoorWebApplicationFirewallPolicies",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/FrontDoorWebApplicationFirewallPolicies", "2019-10-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{policyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/FrontDoorWebApplicationFirewallPolicies/{policyName}", "2019-10-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/FrontDoorWebApplicationFirewallPolicies/{policyName}", "2019-10-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/FrontDoorWebApplicationFirewallPolicies/{policyName}", "2019-10-01"),
}, } ,
},
{ 
	Display: "NetworkExperimentProfiles",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{profileName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}", "2019-11-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "Experiments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}/Experiments", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{experimentName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}/Experiments/{experimentName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}/Experiments/{experimentName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}/Experiments/{experimentName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}/Experiments/{experimentName}", "2019-11-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "LatencyScorecard",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}/Experiments/{experimentName}/LatencyScorecard", "2019-11-01"),
},
{ 
	Display: "Timeseries",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}/Experiments/{experimentName}/Timeseries", "2019-11-01"),
}, } ,
}, } ,
},
{ 
	Display: "PreconfiguredEndpoints",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}/PreconfiguredEndpoints", "2019-11-01"),
}, } ,
}, } ,
},
{ 
	Display: "frontDoors",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors", "2020-01-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{frontDoorName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}", "2020-01-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}", "2020-01-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}", "2020-01-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "frontendEndpoints",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/frontendEndpoints", "2020-01-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{frontendEndpointName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/frontendEndpoints/{frontendEndpointName}", "2020-01-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "rulesEngines",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/rulesEngines", "2020-01-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{rulesEngineName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/rulesEngines/{rulesEngineName}", "2020-01-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/rulesEngines/{rulesEngineName}", "2020-01-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/rulesEngines/{rulesEngineName}", "2020-01-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.GuestConfiguration/operations", "2018-11-20"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.HanaOnAzure/operations", "2017-11-03-preview"),
},
{ 
	Display: "hanaInstances",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.HanaOnAzure/hanaInstances", "2017-11-03-preview"),
},
{ 
	Display: "sapMonitors",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.HanaOnAzure/sapMonitors", "2017-11-03-preview"),
},
{ 
	Display: "hanaInstances",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HanaOnAzure/hanaInstances", "2017-11-03-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{hanaInstanceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HanaOnAzure/hanaInstances/{hanaInstanceName}", "2017-11-03-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HanaOnAzure/hanaInstances/{hanaInstanceName}", "2017-11-03-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HanaOnAzure/hanaInstances/{hanaInstanceName}", "2017-11-03-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HanaOnAzure/hanaInstances/{hanaInstanceName}", "2017-11-03-preview"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "{sapMonitorName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HanaOnAzure/sapMonitors/{sapMonitorName}", "2017-11-03-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HanaOnAzure/sapMonitors/{sapMonitorName}", "2017-11-03-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HanaOnAzure/sapMonitors/{sapMonitorName}", "2017-11-03-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HanaOnAzure/sapMonitors/{sapMonitorName}", "2017-11-03-preview"),
},
{ 
	Display: "dedicatedHSMs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs", "2018-10-31-preview"),
},
{ 
	Display: "dedicatedHSMs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs", "2018-10-31-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs/{name}", "2018-10-31-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs/{name}", "2018-10-31-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs/{name}", "2018-10-31-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HardwareSecurityModules/dedicatedHSMs/{name}", "2018-10-31-preview"),
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.HDInsight/operations", "2018-06-01-preview"),
},
{ 
	Display: "clusters",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.HDInsight/clusters", "2018-06-01-preview"),
},
{ 
	Display: "billingSpecs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.HDInsight/locations/{location}/billingSpecs", "2018-06-01-preview"),
},
{ 
	Display: "capabilities",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.HDInsight/locations/{location}/capabilities", "2018-06-01-preview"),
},
{ 
	Display: "usages",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.HDInsight/locations/{location}/usages", "2018-06-01-preview"),
},
{ 
	Display: "clusters",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters", "2018-06-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{clusterName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}", "2018-06-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}", "2018-06-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}", "2018-06-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}", "2018-06-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "applications",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/applications", "2018-06-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{applicationName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/applications/{applicationName}", "2018-06-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/applications/{applicationName}", "2018-06-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/applications/{applicationName}", "2018-06-01-preview"),
}, } ,
},
{ 
	Display: "clustermonitoring",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/clustermonitoring", "2018-06-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/clustermonitoring", "2018-06-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/clustermonitoring", "2018-06-01-preview"),
},
{ 
	Display: "scriptActions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/scriptActions", "2018-06-01-preview"),
	SubResources: []swagger.ResourceType{  } ,
},
{ 
	Display: "scriptExecutionHistory",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/scriptExecutionHistory", "2018-06-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{scriptExecutionId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/scriptExecutionHistory/{scriptExecutionId}", "2018-06-01-preview"),
	Children: []swagger.ResourceType{  } ,
}, } ,
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{extensionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/{extensionName}", "2018-06-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/{extensionName}", "2018-06-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/{extensionName}", "2018-06-01-preview"),
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.HealthcareApis/operations", "2019-09-16"),
},
{ 
	Display: "{operationResultId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.HealthcareApis/locations/{locationName}/operationresults/{operationResultId}", "2019-09-16"),
},
{ 
	Display: "services",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.HealthcareApis/services", "2019-09-16"),
},
{ 
	Display: "services",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/services", "2019-09-16"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{resourceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/services/{resourceName}", "2019-09-16"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/services/{resourceName}", "2019-09-16"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/services/{resourceName}", "2019-09-16"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/services/{resourceName}", "2019-09-16"),
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.HybridCompute/operations", "2019-12-12"),
},
{ 
	Display: "machines",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.HybridCompute/machines", "2019-12-12"),
},
{ 
	Display: "machines",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines", "2019-12-12"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{name}", "2019-12-12"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{name}", "2019-12-12"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{name}", "2019-12-12"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{name}", "2019-12-12"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "extensions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{name}/extensions", "2019-12-12"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{extensionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{name}/extensions/{extensionName}", "2019-12-12"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{name}/extensions/{extensionName}", "2019-12-12"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{name}/extensions/{extensionName}", "2019-12-12"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{name}/extensions/{extensionName}", "2019-12-12"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.HybridData/operations", "2016-06-01"),
},
{ 
	Display: "dataManagers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.HybridData/dataManagers", "2016-06-01"),
},
{ 
	Display: "dataManagers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers", "2016-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{dataManagerName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}", "2016-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}", "2016-06-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}", "2016-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}", "2016-06-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "dataServices",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices", "2016-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{dataServiceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}", "2016-06-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "jobDefinitions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions", "2016-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{jobDefinitionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions/{jobDefinitionName}", "2016-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions/{jobDefinitionName}", "2016-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions/{jobDefinitionName}", "2016-06-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "jobs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions/{jobDefinitionName}/jobs", "2016-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{jobId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions/{jobDefinitionName}/jobs/{jobId}", "2016-06-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "jobs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobs", "2016-06-01"),
}, } ,
}, } ,
},
{ 
	Display: "dataStoreTypes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataStoreTypes", "2016-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{dataStoreTypeName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataStoreTypes/{dataStoreTypeName}", "2016-06-01"),
}, } ,
},
{ 
	Display: "dataStores",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataStores", "2016-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{dataStoreName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataStores/{dataStoreName}", "2016-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataStores/{dataStoreName}", "2016-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataStores/{dataStoreName}", "2016-06-01"),
}, } ,
},
{ 
	Display: "jobDefinitions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/jobDefinitions", "2016-06-01"),
},
{ 
	Display: "jobs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/jobs", "2016-06-01"),
},
{ 
	Display: "publicKeys",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/publicKeys", "2016-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{publicKeyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/publicKeys/{publicKeyName}", "2016-06-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.VirtualMachineImages/operations", "2019-05-01-preview"),
},
{ 
	Display: "imageTemplates",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VirtualMachineImages/imageTemplates", "2019-05-01-preview"),
},
{ 
	Display: "imageTemplates",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VirtualMachineImages/imageTemplates", "2019-05-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{imageTemplateName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VirtualMachineImages/imageTemplates/{imageTemplateName}", "2019-05-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VirtualMachineImages/imageTemplates/{imageTemplateName}", "2019-05-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VirtualMachineImages/imageTemplates/{imageTemplateName}", "2019-05-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VirtualMachineImages/imageTemplates/{imageTemplateName}", "2019-05-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "runOutputs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VirtualMachineImages/imageTemplates/{imageTemplateName}/runOutputs", "2019-05-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{runOutputName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VirtualMachineImages/imageTemplates/{imageTemplateName}/runOutputs/{runOutputName}", "2019-05-01-preview"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "locations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Intune/locations", "2015-01-14-privatepreview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "hostName",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Intune/locations/hostName", "2015-01-14-privatepreview"),
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "apps",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Intune/locations/{hostName}/AndroidPolicies/{policyName}/apps", "2015-01-14-privatepreview"),
},
{ 
	Display: "androidPolicies",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Intune/locations/{hostName}/androidPolicies", "2015-01-14-privatepreview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{policyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Intune/locations/{hostName}/androidPolicies/{policyName}", "2015-01-14-privatepreview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Intune/locations/{hostName}/androidPolicies/{policyName}", "2015-01-14-privatepreview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Intune/locations/{hostName}/androidPolicies/{policyName}", "2015-01-14-privatepreview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Intune/locations/{hostName}/androidPolicies/{policyName}", "2015-01-14-privatepreview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "groups",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Intune/locations/{hostName}/androidPolicies/{policyName}/groups", "2015-01-14-privatepreview"),
	SubResources: []swagger.ResourceType{  } ,
}, } ,
	SubResources: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "apps",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Intune/locations/{hostName}/apps", "2015-01-14-privatepreview"),
},
{ 
	Display: "flaggedUsers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Intune/locations/{hostName}/flaggedUsers", "2015-01-14-privatepreview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{userName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Intune/locations/{hostName}/flaggedUsers/{userName}", "2015-01-14-privatepreview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "flaggedEnrolledApps",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Intune/locations/{hostName}/flaggedUsers/{userName}/flaggedEnrolledApps", "2015-01-14-privatepreview"),
}, } ,
}, } ,
},
{ 
	Display: "iosPolicies",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Intune/locations/{hostName}/iosPolicies", "2015-01-14-privatepreview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{policyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Intune/locations/{hostName}/iosPolicies/{policyName}", "2015-01-14-privatepreview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Intune/locations/{hostName}/iosPolicies/{policyName}", "2015-01-14-privatepreview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Intune/locations/{hostName}/iosPolicies/{policyName}", "2015-01-14-privatepreview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Intune/locations/{hostName}/iosPolicies/{policyName}", "2015-01-14-privatepreview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "apps",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Intune/locations/{hostName}/iosPolicies/{policyName}/apps", "2015-01-14-privatepreview"),
	SubResources: []swagger.ResourceType{  } ,
},
{ 
	Display: "groups",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Intune/locations/{hostName}/iosPolicies/{policyName}/groups", "2015-01-14-privatepreview"),
	SubResources: []swagger.ResourceType{  } ,
}, } ,
}, } ,
},
{ 
	Display: "operationResults",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Intune/locations/{hostName}/operationResults", "2015-01-14-privatepreview"),
},
{ 
	Display: "default",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Intune/locations/{hostName}/statuses/default", "2015-01-14-privatepreview"),
},
{ 
	Display: "devices",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Intune/locations/{hostName}/users/{userName}/devices", "2015-01-14-privatepreview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{deviceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Intune/locations/{hostName}/users/{userName}/devices/{deviceName}", "2015-01-14-privatepreview"),
	Children: []swagger.ResourceType{  } ,
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.IoTCentral/operations", "2018-09-01"),
},
{ 
	Display: "IoTApps",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.IoTCentral/IoTApps", "2018-09-01"),
},
{ 
	Display: "IoTApps",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTCentral/IoTApps", "2018-09-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{resourceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTCentral/IoTApps/{resourceName}", "2018-09-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTCentral/IoTApps/{resourceName}", "2018-09-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTCentral/IoTApps/{resourceName}", "2018-09-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTCentral/IoTApps/{resourceName}", "2018-09-01"),
}, } ,
},
{ 
	Display: "IotHubs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Devices/IotHubs", "2019-07-01-preview"),
},
{ 
	Display: "usages",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Devices/usages", "2019-07-01-preview"),
},
{ 
	Display: "IotHubs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs", "2019-07-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "routingEndpointsHealth",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{iotHubName}/routingEndpointsHealth", "2019-07-01-preview"),
},
{ 
	Display: "{resourceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}", "2019-07-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}", "2019-07-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}", "2019-07-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}", "2019-07-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "IotHubStats",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/IotHubStats", "2019-07-01-preview"),
},
{ 
	Display: "certificates",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/certificates", "2019-07-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{certificateName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/certificates/{certificateName}", "2019-07-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/certificates/{certificateName}", "2019-07-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/certificates/{certificateName}", "2019-07-01-preview"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "jobs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/jobs", "2019-07-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{jobId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/jobs/{jobId}", "2019-07-01-preview"),
}, } ,
},
{ 
	Display: "quotaMetrics",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/quotaMetrics", "2019-07-01-preview"),
},
{ 
	Display: "skus",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/skus", "2019-07-01-preview"),
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "ConsumerGroups",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/eventHubEndpoints/{eventHubEndpointName}/ConsumerGroups", "2019-07-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/eventHubEndpoints/{eventHubEndpointName}/ConsumerGroups/{name}", "2019-07-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/eventHubEndpoints/{eventHubEndpointName}/ConsumerGroups/{name}", "2019-07-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/eventHubEndpoints/{eventHubEndpointName}/ConsumerGroups/{name}", "2019-07-01-preview"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.IoTSpaces/operations", "2017-10-01-preview"),
},
{ 
	Display: "Graph",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.IoTSpaces/Graph", "2017-10-01-preview"),
},
{ 
	Display: "Graph",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTSpaces/Graph", "2017-10-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{resourceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTSpaces/Graph/{resourceName}", "2017-10-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTSpaces/Graph/{resourceName}", "2017-10-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTSpaces/Graph/{resourceName}", "2017-10-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTSpaces/Graph/{resourceName}", "2017-10-01-preview"),
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.KeyVault/operations", "2019-09-01"),
},
{ 
	Display: "deletedVaults",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.KeyVault/deletedVaults", "2019-09-01"),
},
{ 
	Display: "{vaultName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.KeyVault/locations/{location}/deletedVaults/{vaultName}", "2019-09-01"),
	Children: []swagger.ResourceType{  } ,
},
{ 
	Display: "vaults",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.KeyVault/vaults", "2019-09-01"),
},
{ 
	Display: "vaults",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults", "2019-09-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{vaultName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}", "2019-09-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}", "2019-09-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}", "2019-09-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}", "2019-09-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "privateLinkResources",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/privateLinkResources", "2019-09-01"),
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{privateEndpointConnectionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/privateEndpointConnections/{privateEndpointConnectionName}", "2019-09-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/privateEndpointConnections/{privateEndpointConnectionName}", "2019-09-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/privateEndpointConnections/{privateEndpointConnectionName}", "2019-09-01"),
}, } ,
}, } ,
},
{ 
	Display: "resources",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resources", "2018-05-01"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.LabServices/operations", "2018-10-15"),
},
{ 
	Display: "labaccounts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.LabServices/labaccounts", "2018-10-15"),
},
{ 
	Display: "{operationName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.LabServices/locations/{locationName}/operations/{operationName}", "2018-10-15"),
},
{ 
	Display: "labaccounts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts", "2018-10-15"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{labAccountName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}", "2018-10-15"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}", "2018-10-15"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}", "2018-10-15"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}", "2018-10-15"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "galleryimages",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/galleryimages", "2018-10-15"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{galleryImageName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/galleryimages/{galleryImageName}", "2018-10-15"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/galleryimages/{galleryImageName}", "2018-10-15"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/galleryimages/{galleryImageName}", "2018-10-15"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/galleryimages/{galleryImageName}", "2018-10-15"),
}, } ,
},
{ 
	Display: "labs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs", "2018-10-15"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{labName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}", "2018-10-15"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}", "2018-10-15"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}", "2018-10-15"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}", "2018-10-15"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "environmentsettings",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings", "2018-10-15"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{environmentSettingName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}", "2018-10-15"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}", "2018-10-15"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}", "2018-10-15"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}", "2018-10-15"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "environments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}/environments", "2018-10-15"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{environmentName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}/environments/{environmentName}", "2018-10-15"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}/environments/{environmentName}", "2018-10-15"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}/environments/{environmentName}", "2018-10-15"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}/environments/{environmentName}", "2018-10-15"),
	Children: []swagger.ResourceType{  } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "users",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/users", "2018-10-15"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{userName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/users/{userName}", "2018-10-15"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/users/{userName}", "2018-10-15"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/users/{userName}", "2018-10-15"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/users/{userName}", "2018-10-15"),
}, } ,
}, } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Logic/operations", "2019-05-01"),
},
{ 
	Display: "integrationAccounts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Logic/integrationAccounts", "2019-05-01"),
},
{ 
	Display: "integrationServiceEnvironments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Logic/integrationServiceEnvironments", "2019-05-01"),
},
{ 
	Display: "workflows",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Logic/workflows", "2019-05-01"),
},
{ 
	Display: "integrationAccounts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts", "2019-05-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{integrationAccountName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}", "2019-05-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}", "2019-05-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}", "2019-05-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}", "2019-05-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "agreements",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/agreements", "2019-05-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{agreementName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/agreements/{agreementName}", "2019-05-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/agreements/{agreementName}", "2019-05-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/agreements/{agreementName}", "2019-05-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "assemblies",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/assemblies", "2019-05-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{assemblyArtifactName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/assemblies/{assemblyArtifactName}", "2019-05-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/assemblies/{assemblyArtifactName}", "2019-05-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/assemblies/{assemblyArtifactName}", "2019-05-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "batchConfigurations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/batchConfigurations", "2019-05-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{batchConfigurationName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/batchConfigurations/{batchConfigurationName}", "2019-05-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/batchConfigurations/{batchConfigurationName}", "2019-05-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/batchConfigurations/{batchConfigurationName}", "2019-05-01"),
}, } ,
},
{ 
	Display: "certificates",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/certificates", "2019-05-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{certificateName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/certificates/{certificateName}", "2019-05-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/certificates/{certificateName}", "2019-05-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/certificates/{certificateName}", "2019-05-01"),
}, } ,
},
{ 
	Display: "maps",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/maps", "2019-05-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{mapName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/maps/{mapName}", "2019-05-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/maps/{mapName}", "2019-05-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/maps/{mapName}", "2019-05-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "partners",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/partners", "2019-05-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{partnerName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/partners/{partnerName}", "2019-05-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/partners/{partnerName}", "2019-05-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/partners/{partnerName}", "2019-05-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "schemas",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/schemas", "2019-05-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{schemaName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/schemas/{schemaName}", "2019-05-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/schemas/{schemaName}", "2019-05-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/schemas/{schemaName}", "2019-05-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "sessions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/sessions", "2019-05-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{sessionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/sessions/{sessionName}", "2019-05-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/sessions/{sessionName}", "2019-05-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/sessions/{sessionName}", "2019-05-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "workflows",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows", "2019-05-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{workflowName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}", "2019-05-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}", "2019-05-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}", "2019-05-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}", "2019-05-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "runs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs", "2019-05-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{runName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}", "2019-05-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "actions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions", "2019-05-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{actionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}", "2019-05-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "repetitions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}/repetitions", "2019-05-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{repetitionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}/repetitions/{repetitionName}", "2019-05-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "requestHistories",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}/repetitions/{repetitionName}/requestHistories", "2019-05-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{requestHistoryName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}/repetitions/{repetitionName}/requestHistories/{requestHistoryName}", "2019-05-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "requestHistories",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}/requestHistories", "2019-05-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{requestHistoryName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}/requestHistories/{requestHistoryName}", "2019-05-01"),
}, } ,
},
{ 
	Display: "scopeRepetitions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}/scopeRepetitions", "2019-05-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{repetitionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}/scopeRepetitions/{repetitionName}", "2019-05-01"),
}, } ,
}, } ,
}, } ,
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{operationId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/operations/{operationId}", "2019-05-01"),
}, } ,
}, } ,
},
{ 
	Display: "triggers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/triggers", "2019-05-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{triggerName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/triggers/{triggerName}", "2019-05-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "histories",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/triggers/{triggerName}/histories", "2019-05-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{historyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/triggers/{triggerName}/histories/{historyName}", "2019-05-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "json",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/triggers/{triggerName}/schemas/json", "2019-05-01"),
}, } ,
}, } ,
},
{ 
	Display: "versions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/versions", "2019-05-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{versionId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/versions/{versionId}", "2019-05-01"),
	SubResources: []swagger.ResourceType{  } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "integrationServiceEnvironments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments", "2019-05-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{integrationServiceEnvironmentName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}", "2019-05-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}", "2019-05-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}", "2019-05-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}", "2019-05-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "network",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}/health/network", "2019-05-01"),
},
{ 
	Display: "managedApis",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}/managedApis", "2019-05-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{apiName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}/managedApis/{apiName}", "2019-05-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}/managedApis/{apiName}", "2019-05-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}/managedApis/{apiName}", "2019-05-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "apiOperations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}/managedApis/{apiName}/apiOperations", "2019-05-01"),
}, } ,
}, } ,
},
{ 
	Display: "skus",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}/skus", "2019-05-01"),
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.MachineLearning/operations", "2019-10-01"),
},
{ 
	Display: "workspaces",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.MachineLearning/workspaces", "2019-10-01"),
},
{ 
	Display: "workspaces",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearning/workspaces", "2019-10-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{workspaceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearning/workspaces/{workspaceName}", "2019-10-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearning/workspaces/{workspaceName}", "2019-10-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearning/workspaces/{workspaceName}", "2019-10-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearning/workspaces/{workspaceName}", "2019-10-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.MachineLearningCompute/operations", "2017-08-01-preview"),
},
{ 
	Display: "operationalizationClusters",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.MachineLearningCompute/operationalizationClusters", "2017-08-01-preview"),
},
{ 
	Display: "operationalizationClusters",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningCompute/operationalizationClusters", "2017-08-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{clusterName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningCompute/operationalizationClusters/{clusterName}", "2017-08-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningCompute/operationalizationClusters/{clusterName}", "2017-08-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningCompute/operationalizationClusters/{clusterName}", "2017-08-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningCompute/operationalizationClusters/{clusterName}", "2017-08-01-preview"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.MachineLearningExperimentation/operations", "2017-05-01-preview"),
},
{ 
	Display: "accounts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.MachineLearningExperimentation/accounts", "2017-05-01-preview"),
},
{ 
	Display: "accounts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningExperimentation/accounts", "2017-05-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{accountName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningExperimentation/accounts/{accountName}", "2017-05-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningExperimentation/accounts/{accountName}", "2017-05-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningExperimentation/accounts/{accountName}", "2017-05-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningExperimentation/accounts/{accountName}", "2017-05-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "workspaces",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningExperimentation/accounts/{accountName}/workspaces", "2017-05-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{workspaceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningExperimentation/accounts/{accountName}/workspaces/{workspaceName}", "2017-05-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningExperimentation/accounts/{accountName}/workspaces/{workspaceName}", "2017-05-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningExperimentation/accounts/{accountName}/workspaces/{workspaceName}", "2017-05-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningExperimentation/accounts/{accountName}/workspaces/{workspaceName}", "2017-05-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{projectName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningExperimentation/accounts/{accountName}/workspaces/{workspaceName}/projects/{projectName}", "2017-05-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningExperimentation/accounts/{accountName}/workspaces/{workspaceName}/projects/{projectName}", "2017-05-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningExperimentation/accounts/{accountName}/workspaces/{workspaceName}/projects/{projectName}", "2017-05-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningExperimentation/accounts/{accountName}/workspaces/{workspaceName}/projects/{projectName}", "2017-05-01-preview"),
}, } ,
}, } ,
},
{ 
	Display: "projects",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningExperimentation/accounts/{accountName}/workspaces{workspaceName}/projects", "2017-05-01-preview"),
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.MachineLearningServices/operations", "2020-01-01"),
},
{ 
	Display: "Quotas",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.MachineLearningServices/locations/{location}/Quotas", "2020-01-01"),
},
{ 
	Display: "usages",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.MachineLearningServices/locations/{location}/usages", "2020-01-01"),
},
{ 
	Display: "vmSizes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.MachineLearningServices/locations/{location}/vmSizes", "2020-01-01"),
},
{ 
	Display: "workspaces",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.MachineLearningServices/workspaces", "2020-01-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "skus",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.MachineLearningServices/workspaces/skus", "2020-01-01"),
}, } ,
},
{ 
	Display: "workspaces",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces", "2020-01-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{workspaceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}", "2020-01-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}", "2020-01-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}", "2020-01-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}", "2020-01-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "computes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes", "2020-01-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{computeName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}", "2020-01-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}", "2020-01-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}", "2020-01-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}", "2020-01-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "features",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/features", "2020-01-01"),
},
{ 
	Display: "privateLinkResources",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/privateLinkResources", "2020-01-01"),
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{privateEndpointConnectionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/privateEndpointConnections/{privateEndpointConnectionName}", "2020-01-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/privateEndpointConnections/{privateEndpointConnectionName}", "2020-01-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/privateEndpointConnections/{privateEndpointConnectionName}", "2020-01-01"),
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Maintenance/operations", "2018-06-01-preview"),
},
{ 
	Display: "maintenanceConfigurations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Maintenance/maintenanceConfigurations", "2018-06-01-preview"),
},
{ 
	Display: "{resourceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Maintenance/maintenanceConfigurations/{resourceName}", "2018-06-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Maintenance/maintenanceConfigurations/{resourceName}", "2018-06-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Maintenance/maintenanceConfigurations/{resourceName}", "2018-06-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Maintenance/maintenanceConfigurations/{resourceName}", "2018-06-01-preview"),
},
{ 
	Display: "{applyUpdateName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceParentType}/{resourceParentName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/applyUpdates/{applyUpdateName}", "2018-06-01-preview"),
},
{ 
	Display: "configurationAssignments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceParentType}/{resourceParentName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/configurationAssignments", "2018-06-01-preview"),
	SubResources: []swagger.ResourceType{  } ,
},
{ 
	Display: "updates",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceParentType}/{resourceParentName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/updates", "2018-06-01-preview"),
},
{ 
	Display: "{applyUpdateName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/applyUpdates/{applyUpdateName}", "2018-06-01-preview"),
},
{ 
	Display: "configurationAssignments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/configurationAssignments", "2018-06-01-preview"),
	SubResources: []swagger.ResourceType{  } ,
},
{ 
	Display: "updates",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/updates", "2018-06-01-preview"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ManagedNetwork/operations", "2019-06-01-preview"),
},
{ 
	Display: "managedNetworks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ManagedNetwork/managedNetworks", "2019-06-01-preview"),
},
{ 
	Display: "managedNetworks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks", "2019-06-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{managedNetworkName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}", "2019-06-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}", "2019-06-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}", "2019-06-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}", "2019-06-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "managedNetworkGroups",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}/managedNetworkGroups", "2019-06-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{managedNetworkGroupName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}/managedNetworkGroups/{managedNetworkGroupName}", "2019-06-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}/managedNetworkGroups/{managedNetworkGroupName}", "2019-06-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}/managedNetworkGroups/{managedNetworkGroupName}", "2019-06-01-preview"),
}, } ,
},
{ 
	Display: "managedNetworkPeeringPolicies",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}/managedNetworkPeeringPolicies", "2019-06-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{managedNetworkPeeringPolicyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}/managedNetworkPeeringPolicies/{managedNetworkPeeringPolicyName}", "2019-06-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}/managedNetworkPeeringPolicies/{managedNetworkPeeringPolicyName}", "2019-06-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}/managedNetworkPeeringPolicies/{managedNetworkPeeringPolicyName}", "2019-06-01-preview"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "scopeAssignments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.ManagedNetwork/scopeAssignments", "2019-06-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{scopeAssignmentName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.ManagedNetwork/scopeAssignments/{scopeAssignmentName}", "2019-06-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.ManagedNetwork/scopeAssignments/{scopeAssignmentName}", "2019-06-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.ManagedNetwork/scopeAssignments/{scopeAssignmentName}", "2019-06-01-preview"),
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ManagedServices/operations", "2019-09-01"),
},
{ 
	Display: "marketplaceRegistrationDefinitions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.ManagedServices/marketplaceRegistrationDefinitions", "2019-09-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{marketplaceIdentifier}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.ManagedServices/marketplaceRegistrationDefinitions/{marketplaceIdentifier}", "2019-09-01"),
}, } ,
},
{ 
	Display: "registrationAssignments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.ManagedServices/registrationAssignments", "2019-09-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{registrationAssignmentId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.ManagedServices/registrationAssignments/{registrationAssignmentId}", "2019-09-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.ManagedServices/registrationAssignments/{registrationAssignmentId}", "2019-09-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.ManagedServices/registrationAssignments/{registrationAssignmentId}", "2019-09-01"),
}, } ,
},
{ 
	Display: "registrationDefinitions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.ManagedServices/registrationDefinitions", "2019-09-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{registrationDefinitionId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.ManagedServices/registrationDefinitions/{registrationDefinitionId}", "2019-09-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.ManagedServices/registrationDefinitions/{registrationDefinitionId}", "2019-09-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.ManagedServices/registrationDefinitions/{registrationDefinitionId}", "2019-09-01"),
}, } ,
},
{ 
	Display: "managementGroups",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementGroups", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{groupId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementGroups/{groupId}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementGroups/{groupId}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementGroups/{groupId}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementGroups/{groupId}", "2019-11-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "descendants",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementGroups/{groupId}/descendants", "2019-11-01"),
}, } ,
	SubResources: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Management/operations", "2019-11-01"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ManagementPartner/operations", "2018-02-01"),
},
{ 
	Display: "partners",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ManagementPartner/partners", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{partnerId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ManagementPartner/partners/{partnerId}", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ManagementPartner/partners/{partnerId}", "2018-02-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ManagementPartner/partners/{partnerId}", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ManagementPartner/partners/{partnerId}", "2018-02-01"),
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Maps/operations", "2020-02-01-preview"),
},
{ 
	Display: "accounts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Maps/accounts", "2020-02-01-preview"),
},
{ 
	Display: "accounts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Maps/accounts", "2020-02-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{accountName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Maps/accounts/{accountName}", "2020-02-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Maps/accounts/{accountName}", "2020-02-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Maps/accounts/{accountName}", "2020-02-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Maps/accounts/{accountName}", "2020-02-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "privateAtlases",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Maps/accounts/{accountName}/privateAtlases", "2020-02-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{privateAtlasName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Maps/accounts/{accountName}/privateAtlases/{privateAtlasName}", "2020-02-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Maps/accounts/{accountName}/privateAtlases/{privateAtlasName}", "2020-02-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Maps/accounts/{accountName}/privateAtlases/{privateAtlasName}", "2020-02-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Maps/accounts/{accountName}/privateAtlases/{privateAtlasName}", "2020-02-01-preview"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "keys",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/keys", "2020-01-01-privatepreview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{keyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/keys/{keyName}", "2020-01-01-privatepreview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/keys/{keyName}", "2020-01-01-privatepreview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/keys/{keyName}", "2020-01-01-privatepreview"),
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Marketplace/operations", "2020-01-01"),
},
{ 
	Display: "privateStores",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Marketplace/privateStores", "2020-01-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{PrivateStoreId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Marketplace/privateStores/{PrivateStoreId}", "2020-01-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Marketplace/privateStores/{PrivateStoreId}", "2020-01-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Marketplace/privateStores/{PrivateStoreId}", "2020-01-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "offers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Marketplace/privateStores/{PrivateStoreId}/offers", "2020-01-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{OfferId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Marketplace/privateStores/{PrivateStoreId}/offers/{OfferId}", "2020-01-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Marketplace/privateStores/{PrivateStoreId}/offers/{OfferId}", "2020-01-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Marketplace/privateStores/{PrivateStoreId}/offers/{OfferId}", "2020-01-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.MarketplaceOrdering/operations", "2015-06-01"),
},
{ 
	Display: "agreements",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.MarketplaceOrdering/agreements", "2015-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{planId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.MarketplaceOrdering/agreements/{publisherId}/offers/{offerId}/plans/{planId}", "2015-06-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "current",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.MarketplaceOrdering/offerTypes/{offerType}/publishers/{publisherId}/offers/{offerId}/plans/{planId}/agreements/current", "2015-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.MarketplaceOrdering/offerTypes/{offerType}/publishers/{publisherId}/offers/{offerId}/plans/{planId}/agreements/current", "2015-06-01"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Media/operations", "2018-07-01"),
},
{ 
	Display: "mediaservices",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Media/mediaservices", "2018-07-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{accountName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Media/mediaservices/{accountName}", "2018-07-01"),
}, } ,
},
{ 
	Display: "accountFilters",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/accountFilters", "2018-07-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{filterName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/accountFilters/{filterName}", "2018-07-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/accountFilters/{filterName}", "2018-07-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/accountFilters/{filterName}", "2018-07-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/accountFilters/{filterName}", "2018-07-01"),
}, } ,
},
{ 
	Display: "assets",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets", "2018-07-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{assetName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}", "2018-07-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}", "2018-07-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}", "2018-07-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}", "2018-07-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "assetFilters",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}/assetFilters", "2018-07-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{filterName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}/assetFilters/{filterName}", "2018-07-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}/assetFilters/{filterName}", "2018-07-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}/assetFilters/{filterName}", "2018-07-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}/assetFilters/{filterName}", "2018-07-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "contentKeyPolicies",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/contentKeyPolicies", "2018-07-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{contentKeyPolicyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/contentKeyPolicies/{contentKeyPolicyName}", "2018-07-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/contentKeyPolicies/{contentKeyPolicyName}", "2018-07-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/contentKeyPolicies/{contentKeyPolicyName}", "2018-07-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/contentKeyPolicies/{contentKeyPolicyName}", "2018-07-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "streamingLocators",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingLocators", "2018-07-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{streamingLocatorName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingLocators/{streamingLocatorName}", "2018-07-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingLocators/{streamingLocatorName}", "2018-07-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingLocators/{streamingLocatorName}", "2018-07-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "streamingPolicies",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingPolicies", "2018-07-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{streamingPolicyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingPolicies/{streamingPolicyName}", "2018-07-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingPolicies/{streamingPolicyName}", "2018-07-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingPolicies/{streamingPolicyName}", "2018-07-01"),
}, } ,
},
{ 
	Display: "transforms",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/transforms", "2018-07-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{transformName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/transforms/{transformName}", "2018-07-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/transforms/{transformName}", "2018-07-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/transforms/{transformName}", "2018-07-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/transforms/{transformName}", "2018-07-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "jobs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/transforms/{transformName}/jobs", "2018-07-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{jobName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/transforms/{transformName}/jobs/{jobName}", "2018-07-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/transforms/{transformName}/jobs/{jobName}", "2018-07-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/transforms/{transformName}/jobs/{jobName}", "2018-07-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/transforms/{transformName}/jobs/{jobName}", "2018-07-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "mediaservices",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices", "2018-07-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{accountName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}", "2018-07-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}", "2018-07-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}", "2018-07-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}", "2018-07-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "liveEvents",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents", "2018-07-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{liveEventName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}", "2018-07-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}", "2018-07-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}", "2018-07-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}", "2018-07-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "liveOutputs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}/liveOutputs", "2018-07-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{liveOutputName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}/liveOutputs/{liveOutputName}", "2018-07-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}/liveOutputs/{liveOutputName}", "2018-07-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}/liveOutputs/{liveOutputName}", "2018-07-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "streamingEndpoints",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/streamingEndpoints", "2018-07-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{streamingEndpointName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/streamingEndpoints/{streamingEndpointName}", "2018-07-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/streamingEndpoints/{streamingEndpointName}", "2018-07-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/streamingEndpoints/{streamingEndpointName}", "2018-07-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/streamingEndpoints/{streamingEndpointName}", "2018-07-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Migrate/operations", "2018-09-01-preview"),
},
{ 
	Display: "assessmentProjects",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Migrate/assessmentProjects", "2019-10-01"),
},
{ 
	Display: "assessments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/assessments", "2019-10-01"),
},
{ 
	Display: "groups",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/groups", "2019-10-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{groupName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/groups/{groupName}", "2019-10-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/groups/{groupName}", "2019-10-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/groups/{groupName}", "2019-10-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "assessments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/groups/{groupName}/assessments", "2019-10-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{assessmentName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/groups/{groupName}/assessments/{assessmentName}", "2019-10-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/groups/{groupName}/assessments/{assessmentName}", "2019-10-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/groups/{groupName}/assessments/{assessmentName}", "2019-10-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "assessedMachines",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/groups/{groupName}/assessments/{assessmentName}/assessedMachines", "2019-10-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{assessedMachineName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/groups/{groupName}/assessments/{assessmentName}/assessedMachines/{assessedMachineName}", "2019-10-01"),
}, } ,
}, } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "hypervcollectors",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/hypervcollectors", "2019-10-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{hyperVCollectorName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/hypervcollectors/{hyperVCollectorName}", "2019-10-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/hypervcollectors/{hyperVCollectorName}", "2019-10-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/hypervcollectors/{hyperVCollectorName}", "2019-10-01"),
}, } ,
},
{ 
	Display: "machines",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/machines", "2019-10-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{machineName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/machines/{machineName}", "2019-10-01"),
}, } ,
},
{ 
	Display: "vmwarecollectors",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/vmwarecollectors", "2019-10-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{vmWareCollectorName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/vmwarecollectors/{vmWareCollectorName}", "2019-10-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/vmwarecollectors/{vmWareCollectorName}", "2019-10-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/vmwarecollectors/{vmWareCollectorName}", "2019-10-01"),
}, } ,
},
{ 
	Display: "assessmentProjects",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects", "2019-10-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{projectName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}", "2019-10-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}", "2019-10-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}", "2019-10-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}", "2019-10-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "assessmentOptions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/assessmentOptions", "2019-10-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{assessmentOptionsName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/assessmentOptions/{assessmentOptionsName}", "2019-10-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "{migrateProjectName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}", "2018-09-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}", "2018-09-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}", "2018-09-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}", "2018-09-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "databaseInstances",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}/databaseInstances", "2018-09-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{databaseInstanceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}/databaseInstances/{databaseInstanceName}", "2018-09-01-preview"),
}, } ,
},
{ 
	Display: "databases",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}/databases", "2018-09-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{databaseName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}/databases/{databaseName}", "2018-09-01-preview"),
}, } ,
},
{ 
	Display: "machines",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}/machines", "2018-09-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{machineName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}/machines/{machineName}", "2018-09-01-preview"),
}, } ,
},
{ 
	Display: "migrateEvents",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}/migrateEvents", "2018-09-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{eventName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}/migrateEvents/{eventName}", "2018-09-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}/migrateEvents/{eventName}", "2018-09-01-preview"),
}, } ,
},
{ 
	Display: "solutions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}/solutions", "2018-09-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{solutionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}/solutions/{solutionName}", "2018-09-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}/solutions/{solutionName}", "2018-09-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}/solutions/{solutionName}", "2018-09-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}/solutions/{solutionName}", "2018-09-01-preview"),
	Children: []swagger.ResourceType{  } ,
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.MixedReality/operations", "2019-12-02-preview"),
},
{ 
	Display: "remoteRenderingAccounts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.MixedReality/remoteRenderingAccounts", "2019-12-02-preview"),
},
{ 
	Display: "spatialAnchorsAccounts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.MixedReality/spatialAnchorsAccounts", "2019-12-02-preview"),
},
{ 
	Display: "remoteRenderingAccounts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MixedReality/remoteRenderingAccounts", "2019-12-02-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{accountName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MixedReality/remoteRenderingAccounts/{accountName}", "2019-12-02-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MixedReality/remoteRenderingAccounts/{accountName}", "2019-12-02-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MixedReality/remoteRenderingAccounts/{accountName}", "2019-12-02-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MixedReality/remoteRenderingAccounts/{accountName}", "2019-12-02-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "keys",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MixedReality/remoteRenderingAccounts/{accountName}/keys", "2019-12-02-preview"),
}, } ,
}, } ,
},
{ 
	Display: "spatialAnchorsAccounts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MixedReality/spatialAnchorsAccounts", "2019-12-02-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{accountName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MixedReality/spatialAnchorsAccounts/{accountName}", "2019-12-02-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MixedReality/spatialAnchorsAccounts/{accountName}", "2019-12-02-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MixedReality/spatialAnchorsAccounts/{accountName}", "2019-12-02-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MixedReality/spatialAnchorsAccounts/{accountName}", "2019-12-02-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "keys",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MixedReality/spatialAnchorsAccounts/{accountName}/keys", "2019-12-02-preview"),
}, } ,
}, } ,
},
{ 
	Display: "eventcategories",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/microsoft.insights/eventcategories", "2015-04-01"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/microsoft.insights/operations", "2015-04-01"),
},
{ 
	Display: "diagnosticSettings",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceUri}/providers/microsoft.insights/diagnosticSettings", "2017-05-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceUri}/providers/microsoft.insights/diagnosticSettings/{name}", "2017-05-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceUri}/providers/microsoft.insights/diagnosticSettings/{name}", "2017-05-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceUri}/providers/microsoft.insights/diagnosticSettings/{name}", "2017-05-01-preview"),
}, } ,
},
{ 
	Display: "diagnosticSettingsCategories",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceUri}/providers/microsoft.insights/diagnosticSettingsCategories", "2017-05-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceUri}/providers/microsoft.insights/diagnosticSettingsCategories/{name}", "2017-05-01-preview"),
}, } ,
},
{ 
	Display: "metricDefinitions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceUri}/providers/microsoft.insights/metricDefinitions", "2018-01-01"),
},
{ 
	Display: "metrics",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceUri}/providers/microsoft.insights/metrics", "2018-01-01"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ManagedIdentity/operations", "2018-11-30"),
},
{ 
	Display: "userAssignedIdentities",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ManagedIdentity/userAssignedIdentities", "2018-11-30"),
},
{ 
	Display: "userAssignedIdentities",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities", "2018-11-30"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{resourceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{resourceName}", "2018-11-30"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{resourceName}", "2018-11-30"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{resourceName}", "2018-11-30"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{resourceName}", "2018-11-30"),
}, } ,
},
{ 
	Display: "default",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.ManagedIdentity/identities/default", "2018-11-30"),
},
{ 
	Display: "keys",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/keys", "2020-01-01-privatepreview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{keyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/keys/{keyName}", "2020-01-01-privatepreview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/keys/{keyName}", "2020-01-01-privatepreview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/keys/{keyName}", "2020-01-01-privatepreview"),
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.NetApp/operations", "2019-11-01"),
},
{ 
	Display: "netAppAccounts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{accountName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}", "2019-11-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "capacityPools",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{poolName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}", "2019-11-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "volumes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{volumeName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}", "2019-11-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "replicationStatus",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/replicationStatus", "2019-11-01"),
},
{ 
	Display: "snapshots",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/snapshots", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{snapshotName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/snapshots/{snapshotName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/snapshots/{snapshotName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/snapshots/{snapshotName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/snapshots/{snapshotName}", "2019-11-01"),
}, } ,
}, } ,
}, } ,
}, } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Network/operations", "2019-11-01"),
},
{ 
	Display: "ApplicationGatewayWebApplicationFirewallPolicies",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies", "2019-11-01"),
},
{ 
	Display: "ExpressRoutePorts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/ExpressRoutePorts", "2019-11-01"),
},
{ 
	Display: "ExpressRoutePortsLocations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/ExpressRoutePortsLocations", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{locationName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/ExpressRoutePortsLocations/{locationName}", "2019-11-01"),
}, } ,
},
{ 
	Display: "ServiceEndpointPolicies",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/ServiceEndpointPolicies", "2019-11-01"),
},
{ 
	Display: "applicationGatewayAvailableRequestHeaders",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationGatewayAvailableRequestHeaders", "2019-11-01"),
},
{ 
	Display: "applicationGatewayAvailableResponseHeaders",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationGatewayAvailableResponseHeaders", "2019-11-01"),
},
{ 
	Display: "applicationGatewayAvailableServerVariables",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationGatewayAvailableServerVariables", "2019-11-01"),
},
{ 
	Display: "default",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationGatewayAvailableSslOptions/default", "2019-11-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "predefinedPolicies",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationGatewayAvailableSslOptions/default/predefinedPolicies", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{predefinedPolicyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationGatewayAvailableSslOptions/default/predefinedPolicies/{predefinedPolicyName}", "2019-11-01"),
}, } ,
}, } ,
},
{ 
	Display: "applicationGatewayAvailableWafRuleSets",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationGatewayAvailableWafRuleSets", "2019-11-01"),
},
{ 
	Display: "applicationGateways",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationGateways", "2019-11-01"),
},
{ 
	Display: "applicationSecurityGroups",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationSecurityGroups", "2019-11-01"),
},
{ 
	Display: "azureFirewallFqdnTags",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/azureFirewallFqdnTags", "2019-11-01"),
},
{ 
	Display: "azureFirewalls",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/azureFirewalls", "2019-11-01"),
},
{ 
	Display: "bastionHosts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/bastionHosts", "2019-11-01"),
},
{ 
	Display: "bgpServiceCommunities",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/bgpServiceCommunities", "2019-11-01"),
},
{ 
	Display: "ddosProtectionPlans",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/ddosProtectionPlans", "2019-11-01"),
},
{ 
	Display: "expressRouteCircuits",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteCircuits", "2019-11-01"),
},
{ 
	Display: "expressRouteCrossConnections",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteCrossConnections", "2019-11-01"),
},
{ 
	Display: "expressRouteGateways",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteGateways", "2019-11-01"),
},
{ 
	Display: "expressRouteServiceProviders",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteServiceProviders", "2019-11-01"),
},
{ 
	Display: "firewallPolicies",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/firewallPolicies", "2019-11-01"),
},
{ 
	Display: "ipGroups",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/ipGroups", "2019-11-01"),
},
{ 
	Display: "loadBalancers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/loadBalancers", "2019-11-01"),
},
{ 
	Display: "CheckDnsNameAvailability",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/CheckDnsNameAvailability", "2019-11-01"),
},
{ 
	Display: "autoApprovedPrivateLinkServices",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/autoApprovedPrivateLinkServices", "2019-11-01"),
},
{ 
	Display: "availableDelegations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/availableDelegations", "2019-11-01"),
},
{ 
	Display: "availablePrivateEndpointTypes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/availablePrivateEndpointTypes", "2019-11-01"),
},
{ 
	Display: "availableServiceAliases",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/availableServiceAliases", "2019-11-01"),
},
{ 
	Display: "serviceTags",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/serviceTags", "2019-11-01"),
},
{ 
	Display: "usages",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/usages", "2019-11-01"),
},
{ 
	Display: "virtualNetworkAvailableEndpointServices",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/virtualNetworkAvailableEndpointServices", "2019-11-01"),
},
{ 
	Display: "natGateways",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/natGateways", "2019-11-01"),
},
{ 
	Display: "networkInterfaces",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkInterfaces", "2019-11-01"),
},
{ 
	Display: "networkProfiles",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkProfiles", "2019-11-01"),
},
{ 
	Display: "networkSecurityGroups",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkSecurityGroups", "2019-11-01"),
},
{ 
	Display: "networkWatchers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkWatchers", "2019-11-01"),
},
{ 
	Display: "p2svpnGateways",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/p2svpnGateways", "2019-11-01"),
},
{ 
	Display: "privateEndpoints",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/privateEndpoints", "2019-11-01"),
},
{ 
	Display: "privateLinkServices",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/privateLinkServices", "2019-11-01"),
},
{ 
	Display: "publicIPAddresses",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/publicIPAddresses", "2019-11-01"),
},
{ 
	Display: "publicIPPrefixes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/publicIPPrefixes", "2019-11-01"),
},
{ 
	Display: "routeFilters",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/routeFilters", "2019-11-01"),
},
{ 
	Display: "routeTables",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/routeTables", "2019-11-01"),
},
{ 
	Display: "virtualHubs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualHubs", "2019-11-01"),
},
{ 
	Display: "virtualNetworkTaps",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualNetworkTaps", "2019-11-01"),
},
{ 
	Display: "virtualNetworks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualNetworks", "2019-11-01"),
},
{ 
	Display: "virtualRouters",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualRouters", "2019-11-01"),
},
{ 
	Display: "virtualWans",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualWans", "2019-11-01"),
},
{ 
	Display: "vpnGateways",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/vpnGateways", "2019-11-01"),
},
{ 
	Display: "vpnServerConfigurations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/vpnServerConfigurations", "2019-11-01"),
},
{ 
	Display: "vpnSites",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/vpnSites", "2019-11-01"),
},
{ 
	Display: "ApplicationGatewayWebApplicationFirewallPolicies",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{policyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies/{policyName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies/{policyName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies/{policyName}", "2019-11-01"),
}, } ,
},
{ 
	Display: "ExpressRoutePorts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{expressRoutePortName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}", "2019-11-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "links",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}/links", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{linkName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}/links/{linkName}", "2019-11-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "applicationGateways",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{applicationGatewayName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}", "2019-11-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "applicationSecurityGroups",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{applicationSecurityGroupName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups/{applicationSecurityGroupName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups/{applicationSecurityGroupName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups/{applicationSecurityGroupName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups/{applicationSecurityGroupName}", "2019-11-01"),
}, } ,
},
{ 
	Display: "azureFirewalls",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/azureFirewalls", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{azureFirewallName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/azureFirewalls/{azureFirewallName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/azureFirewalls/{azureFirewallName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/azureFirewalls/{azureFirewallName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/azureFirewalls/{azureFirewallName}", "2019-11-01"),
}, } ,
},
{ 
	Display: "bastionHosts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{bastionHostName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts/{bastionHostName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts/{bastionHostName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts/{bastionHostName}", "2019-11-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "connections",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{virtualNetworkGatewayConnectionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}", "2019-11-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "sharedkey",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}/sharedkey", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}/sharedkey", "2019-11-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
}, } ,
},
{ 
	Display: "{ddosCustomPolicyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosCustomPolicies/{ddosCustomPolicyName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosCustomPolicies/{ddosCustomPolicyName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosCustomPolicies/{ddosCustomPolicyName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosCustomPolicies/{ddosCustomPolicyName}", "2019-11-01"),
},
{ 
	Display: "ddosProtectionPlans",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosProtectionPlans", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{ddosProtectionPlanName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosProtectionPlans/{ddosProtectionPlanName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosProtectionPlans/{ddosProtectionPlanName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosProtectionPlans/{ddosProtectionPlanName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosProtectionPlans/{ddosProtectionPlanName}", "2019-11-01"),
}, } ,
},
{ 
	Display: "expressRouteCircuits",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{circuitName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}", "2019-11-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "authorizations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/authorizations", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{authorizationName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/authorizations/{authorizationName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/authorizations/{authorizationName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/authorizations/{authorizationName}", "2019-11-01"),
}, } ,
},
{ 
	Display: "peerings",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{peeringName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}", "2019-11-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "connections",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/connections", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{connectionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/connections/{connectionName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/connections/{connectionName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/connections/{connectionName}", "2019-11-01"),
}, } ,
},
{ 
	Display: "peerConnections",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/peerConnections", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{connectionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/peerConnections/{connectionName}", "2019-11-01"),
}, } ,
},
{ 
	Display: "stats",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/stats", "2019-11-01"),
}, } ,
	SubResources: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "stats",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/stats", "2019-11-01"),
}, } ,
}, } ,
},
{ 
	Display: "expressRouteCrossConnections",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{crossConnectionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}", "2019-11-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "peerings",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{peeringName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}", "2019-11-01"),
	SubResources: []swagger.ResourceType{  } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "expressRouteGateways",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{expressRouteGatewayName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}", "2019-11-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "expressRouteConnections",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}/expressRouteConnections", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{connectionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}/expressRouteConnections/{connectionName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}/expressRouteConnections/{connectionName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}/expressRouteConnections/{connectionName}", "2019-11-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "firewallPolicies",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{firewallPolicyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}", "2019-11-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "ruleGroups",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleGroups", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{ruleGroupName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleGroups/{ruleGroupName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleGroups/{ruleGroupName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleGroups/{ruleGroupName}", "2019-11-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "ipGroups",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ipGroups", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{ipGroupsName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ipGroups/{ipGroupsName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ipGroups/{ipGroupsName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ipGroups/{ipGroupsName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ipGroups/{ipGroupsName}", "2019-11-01"),
}, } ,
},
{ 
	Display: "loadBalancers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{loadBalancerName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}", "2019-11-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "backendAddressPools",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/backendAddressPools", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{backendAddressPoolName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/backendAddressPools/{backendAddressPoolName}", "2019-11-01"),
}, } ,
},
{ 
	Display: "frontendIPConfigurations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/frontendIPConfigurations", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{frontendIPConfigurationName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/frontendIPConfigurations/{frontendIPConfigurationName}", "2019-11-01"),
}, } ,
},
{ 
	Display: "inboundNatRules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/inboundNatRules", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{inboundNatRuleName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/inboundNatRules/{inboundNatRuleName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/inboundNatRules/{inboundNatRuleName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/inboundNatRules/{inboundNatRuleName}", "2019-11-01"),
}, } ,
},
{ 
	Display: "loadBalancingRules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/loadBalancingRules", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{loadBalancingRuleName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/loadBalancingRules/{loadBalancingRuleName}", "2019-11-01"),
}, } ,
},
{ 
	Display: "networkInterfaces",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/networkInterfaces", "2019-11-01"),
},
{ 
	Display: "outboundRules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/outboundRules", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{outboundRuleName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/outboundRules/{outboundRuleName}", "2019-11-01"),
}, } ,
},
{ 
	Display: "probes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/probes", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{probeName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/probes/{probeName}", "2019-11-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "localNetworkGateways",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{localNetworkGatewayName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways/{localNetworkGatewayName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways/{localNetworkGatewayName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways/{localNetworkGatewayName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways/{localNetworkGatewayName}", "2019-11-01"),
}, } ,
},
{ 
	Display: "autoApprovedPrivateLinkServices",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/locations/{location}/autoApprovedPrivateLinkServices", "2019-11-01"),
},
{ 
	Display: "availableDelegations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/locations/{location}/availableDelegations", "2019-11-01"),
},
{ 
	Display: "availablePrivateEndpointTypes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/locations/{location}/availablePrivateEndpointTypes", "2019-11-01"),
},
{ 
	Display: "availableServiceAliases",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/locations/{location}/availableServiceAliases", "2019-11-01"),
},
{ 
	Display: "natGateways",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{natGatewayName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways/{natGatewayName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways/{natGatewayName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways/{natGatewayName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways/{natGatewayName}", "2019-11-01"),
}, } ,
},
{ 
	Display: "networkInterfaces",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{networkInterfaceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}", "2019-11-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "ipConfigurations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/ipConfigurations", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{ipConfigurationName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/ipConfigurations/{ipConfigurationName}", "2019-11-01"),
}, } ,
},
{ 
	Display: "loadBalancers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/loadBalancers", "2019-11-01"),
},
{ 
	Display: "tapConfigurations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/tapConfigurations", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{tapConfigurationName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/tapConfigurations/{tapConfigurationName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/tapConfigurations/{tapConfigurationName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/tapConfigurations/{tapConfigurationName}", "2019-11-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "networkProfiles",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkProfiles", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{networkProfileName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkProfiles/{networkProfileName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkProfiles/{networkProfileName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkProfiles/{networkProfileName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkProfiles/{networkProfileName}", "2019-11-01"),
}, } ,
},
{ 
	Display: "networkSecurityGroups",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{networkSecurityGroupName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}", "2019-11-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "defaultSecurityRules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/defaultSecurityRules", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{defaultSecurityRuleName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/defaultSecurityRules/{defaultSecurityRuleName}", "2019-11-01"),
}, } ,
},
{ 
	Display: "securityRules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{securityRuleName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules/{securityRuleName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules/{securityRuleName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules/{securityRuleName}", "2019-11-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "networkWatchers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{networkWatcherName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}", "2019-11-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "connectionMonitors",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{connectionMonitorName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors/{connectionMonitorName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors/{connectionMonitorName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors/{connectionMonitorName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/connectionMonitors/{connectionMonitorName}", "2019-11-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "flowLogs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/flowLogs", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{flowLogName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/flowLogs/{flowLogName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/flowLogs/{flowLogName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/flowLogs/{flowLogName}", "2019-11-01"),
}, } ,
},
{ 
	Display: "packetCaptures",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{packetCaptureName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}", "2019-11-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "p2svpnGateways",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{gatewayName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}", "2019-11-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "privateEndpoints",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{privateEndpointName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}", "2019-11-01"),
}, } ,
},
{ 
	Display: "privateLinkServices",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateLinkServices", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{serviceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateLinkServices/{serviceName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateLinkServices/{serviceName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateLinkServices/{serviceName}", "2019-11-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "privateEndpointConnections",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateLinkServices/{serviceName}/privateEndpointConnections", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{peConnectionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateLinkServices/{serviceName}/privateEndpointConnections/{peConnectionName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateLinkServices/{serviceName}/privateEndpointConnections/{peConnectionName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateLinkServices/{serviceName}/privateEndpointConnections/{peConnectionName}", "2019-11-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "publicIPAddresses",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{publicIpAddressName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}", "2019-11-01"),
}, } ,
},
{ 
	Display: "publicIPPrefixes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{publicIpPrefixName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIpPrefixName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIpPrefixName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIpPrefixName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIpPrefixName}", "2019-11-01"),
}, } ,
},
{ 
	Display: "routeFilters",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{routeFilterName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}", "2019-11-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "routeFilterRules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}/routeFilterRules", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{ruleName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}/routeFilterRules/{ruleName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}/routeFilterRules/{ruleName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}/routeFilterRules/{ruleName}", "2019-11-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "routeTables",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{routeTableName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}", "2019-11-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "routes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}/routes", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{routeName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}/routes/{routeName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}/routes/{routeName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}/routes/{routeName}", "2019-11-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "serviceEndpointPolicies",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{serviceEndpointPolicyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}", "2019-11-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "serviceEndpointPolicyDefinitions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}/serviceEndpointPolicyDefinitions", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{serviceEndpointPolicyDefinitionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}/serviceEndpointPolicyDefinitions/{serviceEndpointPolicyDefinitionName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}/serviceEndpointPolicyDefinitions/{serviceEndpointPolicyDefinitionName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}/serviceEndpointPolicyDefinitions/{serviceEndpointPolicyDefinitionName}", "2019-11-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "virtualHubs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{virtualHubName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}", "2019-11-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "hubVirtualNetworkConnections",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubVirtualNetworkConnections", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{connectionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubVirtualNetworkConnections/{connectionName}", "2019-11-01"),
}, } ,
},
{ 
	Display: "routeTables",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/routeTables", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{routeTableName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/routeTables/{routeTableName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/routeTables/{routeTableName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/routeTables/{routeTableName}", "2019-11-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "virtualNetworkGateways",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{virtualNetworkGatewayName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}", "2019-11-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "connections",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}/connections", "2019-11-01"),
}, } ,
}, } ,
},
{ 
	Display: "virtualNetworkTaps",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{tapName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps/{tapName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps/{tapName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps/{tapName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps/{tapName}", "2019-11-01"),
}, } ,
},
{ 
	Display: "virtualNetworks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{virtualNetworkName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}", "2019-11-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "CheckIPAddressAvailability",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/CheckIPAddressAvailability", "2019-11-01"),
},
{ 
	Display: "subnets",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{subnetName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}", "2019-11-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "ResourceNavigationLinks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}/ResourceNavigationLinks", "2019-11-01"),
},
{ 
	Display: "ServiceAssociationLinks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}/ServiceAssociationLinks", "2019-11-01"),
}, } ,
}, } ,
},
{ 
	Display: "usages",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/usages", "2019-11-01"),
},
{ 
	Display: "virtualNetworkPeerings",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/virtualNetworkPeerings", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{virtualNetworkPeeringName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/virtualNetworkPeerings/{virtualNetworkPeeringName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/virtualNetworkPeerings/{virtualNetworkPeeringName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/virtualNetworkPeerings/{virtualNetworkPeeringName}", "2019-11-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "virtualRouters",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualRouters", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{virtualRouterName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualRouters/{virtualRouterName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualRouters/{virtualRouterName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualRouters/{virtualRouterName}", "2019-11-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "peerings",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualRouters/{virtualRouterName}/peerings", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{peeringName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualRouters/{virtualRouterName}/peerings/{peeringName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualRouters/{virtualRouterName}/peerings/{peeringName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualRouters/{virtualRouterName}/peerings/{peeringName}", "2019-11-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "virtualWans",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{VirtualWANName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{VirtualWANName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{VirtualWANName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{VirtualWANName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{VirtualWANName}", "2019-11-01"),
},
{ 
	Display: "supportedSecurityProviders",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{virtualWANName}/supportedSecurityProviders", "2019-11-01"),
}, } ,
},
{ 
	Display: "vpnGateways",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{gatewayName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}", "2019-11-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "vpnConnections",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{connectionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}", "2019-11-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "vpnLinkConnections",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}/vpnLinkConnections", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{linkConnectionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}/vpnLinkConnections/{linkConnectionName}", "2019-11-01"),
}, } ,
}, } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "vpnServerConfigurations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{vpnServerConfigurationName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations/{vpnServerConfigurationName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations/{vpnServerConfigurationName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations/{vpnServerConfigurationName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations/{vpnServerConfigurationName}", "2019-11-01"),
}, } ,
},
{ 
	Display: "vpnSites",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{vpnSiteName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}", "2019-11-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "vpnSiteLinks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}/vpnSiteLinks", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{vpnSiteLinkName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}/vpnSiteLinks/{vpnSiteLinkName}", "2019-11-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "networkInterfaces",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/networkInterfaces", "2018-10-01"),
},
{ 
	Display: "networkInterfaces",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces", "2018-10-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{networkInterfaceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces/{networkInterfaceName}", "2018-10-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "ipConfigurations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces/{networkInterfaceName}/ipConfigurations", "2018-10-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{ipConfigurationName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces/{networkInterfaceName}/ipConfigurations/{ipConfigurationName}", "2018-10-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.NotificationHubs/operations", "2017-04-01"),
},
{ 
	Display: "namespaces",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.NotificationHubs/namespaces", "2017-04-01"),
},
{ 
	Display: "namespaces",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces", "2017-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{namespaceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}", "2017-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}", "2017-04-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}", "2017-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}", "2017-04-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "AuthorizationRules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/AuthorizationRules", "2017-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{authorizationRuleName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "notificationHubs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/notificationHubs", "2017-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{notificationHubName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/notificationHubs/{notificationHubName}", "2017-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/notificationHubs/{notificationHubName}", "2017-04-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/notificationHubs/{notificationHubName}", "2017-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/notificationHubs/{notificationHubName}", "2017-04-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "AuthorizationRules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/notificationHubs/{notificationHubName}/AuthorizationRules", "2017-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{authorizationRuleName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/notificationHubs/{notificationHubName}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/notificationHubs/{notificationHubName}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/notificationHubs/{notificationHubName}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
}, } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.OperationalInsights/operations", "2015-03-20"),
},
{ 
	Display: "linkTargets",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.OperationalInsights/linkTargets", "2015-03-20"),
},
{ 
	Display: "{purgeId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/operations/{purgeId}", "2015-03-20"),
},
{ 
	Display: "savedSearches",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/savedSearches", "2015-03-20"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{savedSearchId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/savedSearches/{savedSearchId}", "2015-03-20"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/savedSearches/{savedSearchId}", "2015-03-20"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/savedSearches/{savedSearchId}", "2015-03-20"),
}, } ,
},
{ 
	Display: "storageInsightConfigs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/storageInsightConfigs", "2015-03-20"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{storageInsightName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/storageInsightConfigs/{storageInsightName}", "2015-03-20"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/storageInsightConfigs/{storageInsightName}", "2015-03-20"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/storageInsightConfigs/{storageInsightName}", "2015-03-20"),
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.OperationsManagement/operations", "2015-11-01-preview"),
},
{ 
	Display: "ManagementAssociations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.OperationsManagement/ManagementAssociations", "2015-11-01-preview"),
},
{ 
	Display: "ManagementConfigurations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.OperationsManagement/ManagementConfigurations", "2015-11-01-preview"),
},
{ 
	Display: "solutions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.OperationsManagement/solutions", "2015-11-01-preview"),
},
{ 
	Display: "{managementConfigurationName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationsManagement/ManagementConfigurations/{managementConfigurationName}", "2015-11-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationsManagement/ManagementConfigurations/{managementConfigurationName}", "2015-11-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationsManagement/ManagementConfigurations/{managementConfigurationName}", "2015-11-01-preview"),
},
{ 
	Display: "solutions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationsManagement/solutions", "2015-11-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{solutionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationsManagement/solutions/{solutionName}", "2015-11-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationsManagement/solutions/{solutionName}", "2015-11-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationsManagement/solutions/{solutionName}", "2015-11-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationsManagement/solutions/{solutionName}", "2015-11-01-preview"),
}, } ,
},
{ 
	Display: "{managementAssociationName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.OperationsManagement/ManagementAssociations/{managementAssociationName}", "2015-11-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.OperationsManagement/ManagementAssociations/{managementAssociationName}", "2015-11-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.OperationsManagement/ManagementAssociations/{managementAssociationName}", "2015-11-01-preview"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Peering/operations", "2020-01-01-preview"),
},
{ 
	Display: "legacyPeerings",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Peering/legacyPeerings", "2020-01-01-preview"),
},
{ 
	Display: "peerAsns",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Peering/peerAsns", "2020-01-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{peerAsnName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Peering/peerAsns/{peerAsnName}", "2020-01-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Peering/peerAsns/{peerAsnName}", "2020-01-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Peering/peerAsns/{peerAsnName}", "2020-01-01-preview"),
}, } ,
},
{ 
	Display: "peeringLocations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Peering/peeringLocations", "2020-01-01-preview"),
},
{ 
	Display: "peeringServiceCountries",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Peering/peeringServiceCountries", "2020-01-01-preview"),
},
{ 
	Display: "peeringServiceLocations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Peering/peeringServiceLocations", "2020-01-01-preview"),
},
{ 
	Display: "peeringServiceProviders",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Peering/peeringServiceProviders", "2020-01-01-preview"),
},
{ 
	Display: "peeringServices",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Peering/peeringServices", "2020-01-01-preview"),
},
{ 
	Display: "peerings",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Peering/peerings", "2020-01-01-preview"),
},
{ 
	Display: "peeringServices",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peeringServices", "2020-01-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{peeringServiceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peeringServices/{peeringServiceName}", "2020-01-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peeringServices/{peeringServiceName}", "2020-01-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peeringServices/{peeringServiceName}", "2020-01-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peeringServices/{peeringServiceName}", "2020-01-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "prefixes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peeringServices/{peeringServiceName}/prefixes", "2020-01-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{prefixName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peeringServices/{peeringServiceName}/prefixes/{prefixName}", "2020-01-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peeringServices/{peeringServiceName}/prefixes/{prefixName}", "2020-01-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peeringServices/{peeringServiceName}/prefixes/{prefixName}", "2020-01-01-preview"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "peerings",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings", "2020-01-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{peeringName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings/{peeringName}", "2020-01-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings/{peeringName}", "2020-01-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings/{peeringName}", "2020-01-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings/{peeringName}", "2020-01-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "registeredAsns",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings/{peeringName}/registeredAsns", "2020-01-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{registeredAsnName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings/{peeringName}/registeredAsns/{registeredAsnName}", "2020-01-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings/{peeringName}/registeredAsns/{registeredAsnName}", "2020-01-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings/{peeringName}/registeredAsns/{registeredAsnName}", "2020-01-01-preview"),
}, } ,
},
{ 
	Display: "registeredPrefixes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings/{peeringName}/registeredPrefixes", "2020-01-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{registeredPrefixName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings/{peeringName}/registeredPrefixes/{registeredPrefixName}", "2020-01-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings/{peeringName}/registeredPrefixes/{registeredPrefixName}", "2020-01-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings/{peeringName}/registeredPrefixes/{registeredPrefixName}", "2020-01-01-preview"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.PolicyInsights/operations", "2019-10-01"),
},
{ 
	Display: "policyMetadata",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.PolicyInsights/policyMetadata", "2019-10-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{resourceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.PolicyInsights/policyMetadata/{resourceName}", "2019-10-01"),
}, } ,
},
{ 
	Display: "remediations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/{managementGroupsNamespace}/managementGroups/{managementGroupId}/providers/Microsoft.PolicyInsights/remediations", "2019-07-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{remediationName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/{managementGroupsNamespace}/managementGroups/{managementGroupId}/providers/Microsoft.PolicyInsights/remediations/{remediationName}", "2019-07-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/{managementGroupsNamespace}/managementGroups/{managementGroupId}/providers/Microsoft.PolicyInsights/remediations/{remediationName}", "2019-07-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/{managementGroupsNamespace}/managementGroups/{managementGroupId}/providers/Microsoft.PolicyInsights/remediations/{remediationName}", "2019-07-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "remediations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.PolicyInsights/remediations", "2019-07-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{remediationName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.PolicyInsights/remediations/{remediationName}", "2019-07-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.PolicyInsights/remediations/{remediationName}", "2019-07-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.PolicyInsights/remediations/{remediationName}", "2019-07-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "remediations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PolicyInsights/remediations", "2019-07-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{remediationName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PolicyInsights/remediations/{remediationName}", "2019-07-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PolicyInsights/remediations/{remediationName}", "2019-07-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PolicyInsights/remediations/{remediationName}", "2019-07-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "remediations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceId}/providers/Microsoft.PolicyInsights/remediations", "2019-07-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{remediationName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceId}/providers/Microsoft.PolicyInsights/remediations/{remediationName}", "2019-07-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceId}/providers/Microsoft.PolicyInsights/remediations/{remediationName}", "2019-07-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceId}/providers/Microsoft.PolicyInsights/remediations/{remediationName}", "2019-07-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "$metadata",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.PolicyInsights/policyEvents/$metadata", "2018-04-04"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Portal/operations", "2019-01-01-preview"),
},
{ 
	Display: "dashboards",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Portal/dashboards", "2019-01-01-preview"),
},
{ 
	Display: "dashboards",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Portal/dashboards", "2019-01-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{dashboardName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Portal/dashboards/{dashboardName}", "2019-01-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Portal/dashboards/{dashboardName}", "2019-01-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Portal/dashboards/{dashboardName}", "2019-01-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Portal/dashboards/{dashboardName}", "2019-01-01-preview"),
}, } ,
},
{ 
	Display: "keys",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/keys", "2020-01-01-privatepreview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{keyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/keys/{keyName}", "2020-01-01-privatepreview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/keys/{keyName}", "2020-01-01-privatepreview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/keys/{keyName}", "2020-01-01-privatepreview"),
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.PowerBIDedicated/operations", "2017-10-01"),
},
{ 
	Display: "capacities",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.PowerBIDedicated/capacities", "2017-10-01"),
},
{ 
	Display: "skus",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.PowerBIDedicated/skus", "2017-10-01"),
},
{ 
	Display: "capacities",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBIDedicated/capacities", "2017-10-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{dedicatedCapacityName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBIDedicated/capacities/{dedicatedCapacityName}", "2017-10-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBIDedicated/capacities/{dedicatedCapacityName}", "2017-10-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBIDedicated/capacities/{dedicatedCapacityName}", "2017-10-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBIDedicated/capacities/{dedicatedCapacityName}", "2017-10-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "skus",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBIDedicated/capacities/{dedicatedCapacityName}/skus", "2017-10-01"),
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.PowerBI/operations", "2016-01-29"),
},
{ 
	Display: "workspaceCollections",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.PowerBI/workspaceCollections", "2016-01-29"),
},
{ 
	Display: "workspaceCollections",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/workspaceCollections", "2016-01-29"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{workspaceCollectionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/workspaceCollections/{workspaceCollectionName}", "2016-01-29"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/workspaceCollections/{workspaceCollectionName}", "2016-01-29"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/workspaceCollections/{workspaceCollectionName}", "2016-01-29"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/workspaceCollections/{workspaceCollectionName}", "2016-01-29"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "workspaces",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/workspaceCollections/{workspaceCollectionName}/workspaces", "2016-01-29"),
}, } ,
}, } ,
},
{ 
	Display: "privateDnsZones",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/privateDnsZones", "2018-09-01"),
},
{ 
	Display: "privateDnsZones",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones", "2018-09-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{privateZoneName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}", "2018-09-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}", "2018-09-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}", "2018-09-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}", "2018-09-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "ALL",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/ALL", "2018-09-01"),
},
{ 
	Display: "virtualNetworkLinks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/virtualNetworkLinks", "2018-09-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{virtualNetworkLinkName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/virtualNetworkLinks/{virtualNetworkLinkName}", "2018-09-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/virtualNetworkLinks/{virtualNetworkLinkName}", "2018-09-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/virtualNetworkLinks/{virtualNetworkLinkName}", "2018-09-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/virtualNetworkLinks/{virtualNetworkLinkName}", "2018-09-01"),
}, } ,
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{recordType}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/{recordType}", "2018-09-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{relativeRecordSetName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/{recordType}/{relativeRecordSetName}", "2018-09-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/{recordType}/{relativeRecordSetName}", "2018-09-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/{recordType}/{relativeRecordSetName}", "2018-09-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/{recordType}/{relativeRecordSetName}", "2018-09-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "replicationUsages",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/replicationUsages", "2016-06-01"),
},
{ 
	Display: "usages",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/usages", "2016-06-01"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.RecoveryServices/operations", "2016-08-10"),
},
{ 
	Display: "vaults",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.RecoveryServices/vaults", "2016-06-01"),
},
{ 
	Display: "vaults",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults", "2016-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{vaultName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}", "2016-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}", "2016-06-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}", "2016-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}", "2016-06-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "vaultExtendedInfo",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/extendedInformation/vaultExtendedInfo", "2016-06-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/extendedInformation/vaultExtendedInfo", "2016-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/extendedInformation/vaultExtendedInfo", "2016-06-01"),
}, } ,
}, } ,
},
{ 
	Display: "backupEngines",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupEngines", "2016-12-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{backupEngineName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupEngines/{backupEngineName}", "2016-12-01"),
}, } ,
},
{ 
	Display: "{intentObjectName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/backupProtectionIntent/{intentObjectName}", "2017-07-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/backupProtectionIntent/{intentObjectName}", "2017-07-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/backupProtectionIntent/{intentObjectName}", "2017-07-01"),
},
{ 
	Display: "{operationId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/operationResults/{operationId}", "2016-12-01"),
},
{ 
	Display: "protectableContainers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectableContainers", "2016-12-01"),
},
{ 
	Display: "{containerName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}", "2016-12-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}", "2016-12-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}", "2016-12-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "items",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/items", "2016-12-01"),
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{operationId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/operationResults/{operationId}", "2016-12-01"),
},
{ 
	Display: "{protectedItemName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/protectedItems/{protectedItemName}", "2019-06-15"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/protectedItems/{protectedItemName}", "2019-06-15"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/protectedItems/{protectedItemName}", "2019-06-15"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "recoveryPoints",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/protectedItems/{protectedItemName}/recoveryPoints", "2019-06-15"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{recoveryPointId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/protectedItems/{protectedItemName}/recoveryPoints/{recoveryPointId}", "2019-06-15"),
	Children: []swagger.ResourceType{  } ,
}, } ,
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{operationId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/protectedItems/{protectedItemName}/operationResults/{operationId}", "2019-06-15"),
},
{ 
	Display: "{operationId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/protectedItems/{protectedItemName}/operationsStatus/{operationId}", "2016-12-01"),
}, } ,
}, } ,
},
{ 
	Display: "backupJobs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupJobs", "2019-06-15"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{operationId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupJobs/operationResults/{operationId}", "2019-06-15"),
},
{ 
	Display: "{jobName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupJobs/{jobName}", "2019-06-15"),
	Children: []swagger.ResourceType{  } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{operationId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupJobs/{jobName}/operationResults/{operationId}", "2019-06-15"),
}, } ,
}, } ,
},
{ 
	Display: "{operationId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupOperationResults/{operationId}", "2016-12-01"),
},
{ 
	Display: "{operationId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupOperations/{operationId}", "2016-12-01"),
},
{ 
	Display: "backupPolicies",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupPolicies", "2019-06-15"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{policyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupPolicies/{policyName}", "2019-06-15"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupPolicies/{policyName}", "2016-12-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupPolicies/{policyName}", "2019-06-15"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{operationId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupPolicies/{policyName}/operationResults/{operationId}", "2019-06-15"),
},
{ 
	Display: "{operationId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupPolicies/{policyName}/operations/{operationId}", "2016-12-01"),
}, } ,
}, } ,
},
{ 
	Display: "backupProtectableItems",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupProtectableItems", "2016-12-01"),
},
{ 
	Display: "backupProtectedItems",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupProtectedItems", "2019-06-15"),
},
{ 
	Display: "backupProtectionContainers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupProtectionContainers", "2016-12-01"),
},
{ 
	Display: "backupProtectionIntents",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupProtectionIntents", "2017-07-01"),
},
{ 
	Display: "backupUsageSummaries",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupUsageSummaries", "2017-07-01"),
},
{ 
	Display: "vaultconfig",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupconfig/vaultconfig", "2019-06-15"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupconfig/vaultconfig", "2019-06-15"),
},
{ 
	Display: "vaultstorageconfig",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupstorageconfig/vaultstorageconfig", "2016-12-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupstorageconfig/vaultstorageconfig", "2016-12-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupstorageconfig/vaultstorageconfig", "2016-12-01"),
},
{ 
	Display: "replicationEligibilityResults",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{virtualMachineName}/providers/Microsoft.RecoveryServices/replicationEligibilityResults", "2018-07-10"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "default",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{virtualMachineName}/providers/Microsoft.RecoveryServices/replicationEligibilityResults/default", "2018-07-10"),
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/operations", "2018-07-10"),
},
{ 
	Display: "replicationAlertSettings",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationAlertSettings", "2018-07-10"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{alertSettingName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationAlertSettings/{alertSettingName}", "2018-07-10"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationAlertSettings/{alertSettingName}", "2018-07-10"),
}, } ,
},
{ 
	Display: "replicationEvents",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationEvents", "2018-07-10"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{eventName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationEvents/{eventName}", "2018-07-10"),
}, } ,
},
{ 
	Display: "replicationFabrics",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics", "2018-07-10"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{fabricName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}", "2018-07-10"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}", "2018-07-10"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}", "2018-07-10"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "replicationLogicalNetworks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationLogicalNetworks", "2018-07-10"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{logicalNetworkName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationLogicalNetworks/{logicalNetworkName}", "2018-07-10"),
}, } ,
},
{ 
	Display: "replicationNetworks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationNetworks", "2018-07-10"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{networkName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationNetworks/{networkName}", "2018-07-10"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "replicationNetworkMappings",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationNetworks/{networkName}/replicationNetworkMappings", "2018-07-10"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{networkMappingName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationNetworks/{networkName}/replicationNetworkMappings/{networkMappingName}", "2018-07-10"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationNetworks/{networkName}/replicationNetworkMappings/{networkMappingName}", "2018-07-10"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationNetworks/{networkName}/replicationNetworkMappings/{networkMappingName}", "2018-07-10"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationNetworks/{networkName}/replicationNetworkMappings/{networkMappingName}", "2018-07-10"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "replicationProtectionContainers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers", "2018-07-10"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{protectionContainerName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}", "2018-07-10"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}", "2018-07-10"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "replicationMigrationItems",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationMigrationItems", "2018-07-10"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{migrationItemName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationMigrationItems/{migrationItemName}", "2018-07-10"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationMigrationItems/{migrationItemName}", "2018-07-10"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationMigrationItems/{migrationItemName}", "2018-07-10"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationMigrationItems/{migrationItemName}", "2018-07-10"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "migrationRecoveryPoints",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationMigrationItems/{migrationItemName}/migrationRecoveryPoints", "2018-07-10"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{migrationRecoveryPointName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationMigrationItems/{migrationItemName}/migrationRecoveryPoints/{migrationRecoveryPointName}", "2018-07-10"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "replicationProtectableItems",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectableItems", "2018-07-10"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{protectableItemName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectableItems/{protectableItemName}", "2018-07-10"),
}, } ,
},
{ 
	Display: "replicationProtectedItems",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectedItems", "2018-07-10"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{replicatedProtectedItemName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectedItems/{replicatedProtectedItemName}", "2018-07-10"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectedItems/{replicatedProtectedItemName}", "2018-07-10"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectedItems/{replicatedProtectedItemName}", "2018-07-10"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectedItems/{replicatedProtectedItemName}", "2018-07-10"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "recoveryPoints",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectedItems/{replicatedProtectedItemName}/recoveryPoints", "2018-07-10"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{recoveryPointName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectedItems/{replicatedProtectedItemName}/recoveryPoints/{recoveryPointName}", "2018-07-10"),
}, } ,
},
{ 
	Display: "targetComputeSizes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectedItems/{replicatedProtectedItemName}/targetComputeSizes", "2018-07-10"),
}, } ,
}, } ,
},
{ 
	Display: "replicationProtectionContainerMappings",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectionContainerMappings", "2018-07-10"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{mappingName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectionContainerMappings/{mappingName}", "2018-07-10"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectionContainerMappings/{mappingName}", "2018-07-10"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectionContainerMappings/{mappingName}", "2018-07-10"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectionContainerMappings/{mappingName}", "2018-07-10"),
	Children: []swagger.ResourceType{  } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "replicationRecoveryServicesProviders",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationRecoveryServicesProviders", "2018-07-10"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{providerName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationRecoveryServicesProviders/{providerName}", "2018-07-10"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationRecoveryServicesProviders/{providerName}", "2018-07-10"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationRecoveryServicesProviders/{providerName}", "2018-07-10"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "replicationStorageClassifications",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationStorageClassifications", "2018-07-10"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{storageClassificationName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationStorageClassifications/{storageClassificationName}", "2018-07-10"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "replicationStorageClassificationMappings",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationStorageClassifications/{storageClassificationName}/replicationStorageClassificationMappings", "2018-07-10"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{storageClassificationMappingName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationStorageClassifications/{storageClassificationName}/replicationStorageClassificationMappings/{storageClassificationMappingName}", "2018-07-10"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationStorageClassifications/{storageClassificationName}/replicationStorageClassificationMappings/{storageClassificationMappingName}", "2018-07-10"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationStorageClassifications/{storageClassificationName}/replicationStorageClassificationMappings/{storageClassificationMappingName}", "2018-07-10"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "replicationvCenters",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationvCenters", "2018-07-10"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{vCenterName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationvCenters/{vCenterName}", "2018-07-10"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationvCenters/{vCenterName}", "2018-07-10"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationvCenters/{vCenterName}", "2018-07-10"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationvCenters/{vCenterName}", "2018-07-10"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "replicationJobs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationJobs", "2018-07-10"),
	Children: []swagger.ResourceType{  } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{jobName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationJobs/{jobName}", "2018-07-10"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "replicationMigrationItems",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationMigrationItems", "2018-07-10"),
},
{ 
	Display: "replicationNetworkMappings",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationNetworkMappings", "2018-07-10"),
},
{ 
	Display: "replicationNetworks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationNetworks", "2018-07-10"),
},
{ 
	Display: "replicationPolicies",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationPolicies", "2018-07-10"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{policyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationPolicies/{policyName}", "2018-07-10"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationPolicies/{policyName}", "2018-07-10"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationPolicies/{policyName}", "2018-07-10"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationPolicies/{policyName}", "2018-07-10"),
}, } ,
},
{ 
	Display: "replicationProtectedItems",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationProtectedItems", "2018-07-10"),
},
{ 
	Display: "replicationProtectionContainerMappings",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationProtectionContainerMappings", "2018-07-10"),
},
{ 
	Display: "replicationProtectionContainers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationProtectionContainers", "2018-07-10"),
},
{ 
	Display: "replicationRecoveryPlans",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationRecoveryPlans", "2018-07-10"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{recoveryPlanName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationRecoveryPlans/{recoveryPlanName}", "2018-07-10"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationRecoveryPlans/{recoveryPlanName}", "2018-07-10"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationRecoveryPlans/{recoveryPlanName}", "2018-07-10"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationRecoveryPlans/{recoveryPlanName}", "2018-07-10"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "replicationRecoveryServicesProviders",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationRecoveryServicesProviders", "2018-07-10"),
},
{ 
	Display: "replicationStorageClassificationMappings",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationStorageClassificationMappings", "2018-07-10"),
},
{ 
	Display: "replicationStorageClassifications",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationStorageClassifications", "2018-07-10"),
},
{ 
	Display: "replicationSupportedOperatingSystems",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationSupportedOperatingSystems", "2018-07-10"),
},
{ 
	Display: "replicationVaultHealth",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationVaultHealth", "2018-07-10"),
	Children: []swagger.ResourceType{  } ,
},
{ 
	Display: "replicationVaultSettings",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationVaultSettings", "2018-07-10"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{vaultSettingName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationVaultSettings/{vaultSettingName}", "2018-07-10"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationVaultSettings/{vaultSettingName}", "2018-07-10"),
}, } ,
},
{ 
	Display: "replicationvCenters",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationvCenters", "2018-07-10"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Cache/operations", "2018-03-01"),
},
{ 
	Display: "Redis",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Cache/Redis", "2018-03-01"),
},
{ 
	Display: "Redis",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis", "2018-03-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "firewallRules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{cacheName}/firewallRules", "2018-03-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{ruleName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{cacheName}/firewallRules/{ruleName}", "2018-03-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{cacheName}/firewallRules/{ruleName}", "2018-03-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{cacheName}/firewallRules/{ruleName}", "2018-03-01"),
}, } ,
},
{ 
	Display: "patchSchedules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{cacheName}/patchSchedules", "2018-03-01"),
},
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}", "2018-03-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}", "2018-03-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}", "2018-03-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}", "2018-03-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "linkedServers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/linkedServers", "2018-03-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{linkedServerName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/linkedServers/{linkedServerName}", "2018-03-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/linkedServers/{linkedServerName}", "2018-03-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/linkedServers/{linkedServerName}", "2018-03-01"),
}, } ,
},
{ 
	Display: "listUpgradeNotifications",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/listUpgradeNotifications", "2018-03-01"),
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{default}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/patchSchedules/{default}", "2018-03-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/patchSchedules/{default}", "2018-03-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/patchSchedules/{default}", "2018-03-01"),
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Relay/operations", "2017-04-01"),
},
{ 
	Display: "namespaces",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Relay/namespaces", "2017-04-01"),
},
{ 
	Display: "namespaces",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces", "2017-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{namespaceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}", "2017-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}", "2017-04-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}", "2017-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}", "2017-04-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "authorizationRules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/authorizationRules", "2017-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{authorizationRuleName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "hybridConnections",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/hybridConnections", "2017-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{hybridConnectionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/hybridConnections/{hybridConnectionName}", "2017-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/hybridConnections/{hybridConnectionName}", "2017-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/hybridConnections/{hybridConnectionName}", "2017-04-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "authorizationRules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/hybridConnections/{hybridConnectionName}/authorizationRules", "2017-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{authorizationRuleName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/hybridConnections/{hybridConnectionName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/hybridConnections/{hybridConnectionName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/hybridConnections/{hybridConnectionName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "wcfRelays",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/wcfRelays", "2017-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{relayName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/wcfRelays/{relayName}", "2017-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/wcfRelays/{relayName}", "2017-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/wcfRelays/{relayName}", "2017-04-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "authorizationRules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/wcfRelays/{relayName}/authorizationRules", "2017-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{authorizationRuleName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/wcfRelays/{relayName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/wcfRelays/{relayName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/wcfRelays/{relayName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
}, } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Capacity/operations", "2019-04-01"),
},
{ 
	Display: "reservationOrders",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Capacity/reservationOrders", "2019-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{reservationOrderId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}", "2019-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}", "2019-04-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "reservations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/reservations", "2019-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{reservationId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/reservations/{reservationId}", "2019-04-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/reservations/{reservationId}", "2019-04-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "revisions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/reservations/{reservationId}/revisions", "2019-04-01"),
}, } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "appliedReservations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Capacity/appliedReservations", "2019-04-01"),
},
{ 
	Display: "autoQuotaIncrease",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Capacity/autoQuotaIncrease", "2019-07-19-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Capacity/autoQuotaIncrease", "2019-07-19-preview"),
},
{ 
	Display: "catalogs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Capacity/catalogs", "2019-04-01"),
},
{ 
	Display: "serviceLimits",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Capacity/resourceProviders/{providerId}/locations/{location}/serviceLimits", "2019-07-19-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{resourceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Capacity/resourceProviders/{providerId}/locations/{location}/serviceLimits/{resourceName}", "2019-07-19-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Capacity/resourceProviders/{providerId}/locations/{location}/serviceLimits/{resourceName}", "2019-07-19-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Capacity/resourceProviders/{providerId}/locations/{location}/serviceLimits/{resourceName}", "2019-07-19-preview"),
}, } ,
},
{ 
	Display: "serviceLimitsRequests",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Capacity/resourceProviders/{providerId}/locations/{location}/serviceLimitsRequests", "2019-07-19-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{id}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Capacity/resourceProviders/{providerId}/locations/{location}/serviceLimitsRequests/{id}", "2019-07-19-preview"),
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ResourceGraph/operations", "2019-04-01"),
},
{ 
	Display: "emergingIssues",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ResourceHealth/emergingIssues", "2018-07-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{issueName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ResourceHealth/emergingIssues/{issueName}", "2018-07-01"),
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ResourceHealth/operations", "2018-07-01"),
},
{ 
	Display: "availabilityStatuses",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ResourceHealth/availabilityStatuses", "2018-07-01"),
},
{ 
	Display: "events",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ResourceHealth/events", "2018-07-01"),
},
{ 
	Display: "availabilityStatuses",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ResourceHealth/availabilityStatuses", "2018-07-01"),
},
{ 
	Display: "availabilityStatuses",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceUri}/providers/Microsoft.ResourceHealth/availabilityStatuses", "2018-07-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "current",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceUri}/providers/Microsoft.ResourceHealth/availabilityStatuses/current", "2018-07-01"),
}, } ,
},
{ 
	Display: "events",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceUri}/providers/Microsoft.ResourceHealth/events", "2018-07-01"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Authorization/operations", "2016-09-01"),
},
{ 
	Display: "policyDefinitions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Authorization/policyDefinitions", "2016-12-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{policyDefinitionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", "2016-12-01"),
}, } ,
},
{ 
	Display: "policyDefinitions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policyDefinitions", "2016-12-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{policyDefinitionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", "2016-12-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", "2016-12-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", "2016-12-01"),
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Resources/operations", "2018-05-01"),
},
{ 
	Display: "subscriptions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions", "2016-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{subscriptionId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}", "2016-06-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "locations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/locations", "2016-06-01"),
},
{ 
	Display: "providers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers", "2018-05-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "locks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/locks", "2016-09-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{lockName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/locks/{lockName}", "2016-09-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/locks/{lockName}", "2016-09-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/locks/{lockName}", "2016-09-01"),
}, } ,
},
{ 
	Display: "policyAssignments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyAssignments", "2016-12-01"),
},
{ 
	Display: "policyDefinitions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyDefinitions", "2016-12-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{policyDefinitionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", "2016-12-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", "2016-12-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", "2016-12-01"),
}, } ,
},
{ 
	Display: "deployments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Resources/deployments", "2018-05-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{deploymentName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Resources/deployments/{deploymentName}", "2018-05-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Resources/deployments/{deploymentName}", "2018-05-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Resources/deployments/{deploymentName}", "2018-05-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Resources/deployments/{deploymentName}/operations", "2018-05-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{operationId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Resources/deployments/{deploymentName}/operations/{operationId}", "2018-05-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "jobCollections",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Scheduler/jobCollections", "2016-03-01"),
},
{ 
	Display: "searchServices",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Search/searchServices", "2015-08-19"),
},
{ 
	Display: "alerts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/alerts", "2019-01-01"),
},
{ 
	Display: "allowedConnections",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/allowedConnections", "2015-06-01-preview"),
},
{ 
	Display: "applicationWhitelistings",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/applicationWhitelistings", "2015-06-01-preview"),
},
{ 
	Display: "assessmentMetadata",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/assessmentMetadata", "2020-01-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{assessmentMetadataName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/assessmentMetadata/{assessmentMetadataName}", "2020-01-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/assessmentMetadata/{assessmentMetadataName}", "2020-01-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/assessmentMetadata/{assessmentMetadataName}", "2020-01-01"),
}, } ,
},
{ 
	Display: "autoProvisioningSettings",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/autoProvisioningSettings", "2017-08-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{settingName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/autoProvisioningSettings/{settingName}", "2017-08-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/autoProvisioningSettings/{settingName}", "2017-08-01-preview"),
}, } ,
},
{ 
	Display: "automations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/automations", "2019-01-01-preview"),
},
{ 
	Display: "discoveredSecuritySolutions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/discoveredSecuritySolutions", "2015-06-01-preview"),
},
{ 
	Display: "externalSecuritySolutions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/externalSecuritySolutions", "2015-06-01-preview"),
},
{ 
	Display: "iotSecuritySolutions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/iotSecuritySolutions", "2019-08-01"),
},
{ 
	Display: "jitNetworkAccessPolicies",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/jitNetworkAccessPolicies", "2015-06-01-preview"),
},
{ 
	Display: "locations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations", "2015-06-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{ascLocation}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations/{ascLocation}", "2015-06-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "ExternalSecuritySolutions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations/{ascLocation}/ExternalSecuritySolutions", "2015-06-01-preview"),
},
{ 
	Display: "alerts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations/{ascLocation}/alerts", "2019-01-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{alertName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations/{ascLocation}/alerts/{alertName}", "2019-01-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "allowedConnections",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations/{ascLocation}/allowedConnections", "2015-06-01-preview"),
},
{ 
	Display: "discoveredSecuritySolutions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations/{ascLocation}/discoveredSecuritySolutions", "2015-06-01-preview"),
},
{ 
	Display: "jitNetworkAccessPolicies",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations/{ascLocation}/jitNetworkAccessPolicies", "2015-06-01-preview"),
},
{ 
	Display: "tasks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations/{ascLocation}/tasks", "2015-06-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{taskName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations/{ascLocation}/tasks/{taskName}", "2015-06-01-preview"),
	SubResources: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "topologies",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations/{ascLocation}/topologies", "2015-06-01-preview"),
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{groupName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations/{ascLocation}/applicationWhitelistings/{groupName}", "2015-06-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations/{ascLocation}/applicationWhitelistings/{groupName}", "2015-06-01-preview"),
}, } ,
}, } ,
},
{ 
	Display: "pricings",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/pricings", "2018-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{pricingName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/pricings/{pricingName}", "2018-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/pricings/{pricingName}", "2018-06-01"),
}, } ,
},
{ 
	Display: "regulatoryComplianceStandards",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/regulatoryComplianceStandards", "2019-01-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{regulatoryComplianceStandardName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/regulatoryComplianceStandards/{regulatoryComplianceStandardName}", "2019-01-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "regulatoryComplianceControls",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/regulatoryComplianceStandards/{regulatoryComplianceStandardName}/regulatoryComplianceControls", "2019-01-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{regulatoryComplianceControlName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/regulatoryComplianceStandards/{regulatoryComplianceStandardName}/regulatoryComplianceControls/{regulatoryComplianceControlName}", "2019-01-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "regulatoryComplianceAssessments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/regulatoryComplianceStandards/{regulatoryComplianceStandardName}/regulatoryComplianceControls/{regulatoryComplianceControlName}/regulatoryComplianceAssessments", "2019-01-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{regulatoryComplianceAssessmentName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/regulatoryComplianceStandards/{regulatoryComplianceStandardName}/regulatoryComplianceControls/{regulatoryComplianceControlName}/regulatoryComplianceAssessments/{regulatoryComplianceAssessmentName}", "2019-01-01-preview"),
}, } ,
}, } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "securityContacts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/securityContacts", "2017-08-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{securityContactName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/securityContacts/{securityContactName}", "2017-08-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/securityContacts/{securityContactName}", "2017-08-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/securityContacts/{securityContactName}", "2017-08-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/securityContacts/{securityContactName}", "2017-08-01-preview"),
}, } ,
},
{ 
	Display: "settings",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/settings", "2019-01-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{settingName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/settings/{settingName}", "2019-01-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/settings/{settingName}", "2019-01-01"),
}, } ,
},
{ 
	Display: "tasks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/tasks", "2015-06-01-preview"),
},
{ 
	Display: "topologies",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/topologies", "2015-06-01-preview"),
},
{ 
	Display: "workspaceSettings",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/workspaceSettings", "2017-08-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{workspaceSettingName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/workspaceSettings/{workspaceSettingName}", "2017-08-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/workspaceSettings/{workspaceSettingName}", "2017-08-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/workspaceSettings/{workspaceSettingName}", "2017-08-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/workspaceSettings/{workspaceSettingName}", "2017-08-01-preview"),
}, } ,
},
{ 
	Display: "namespaces",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceBus/namespaces", "2017-04-01"),
},
{ 
	Display: "premiumMessagingRegions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceBus/premiumMessagingRegions", "2017-04-01"),
},
{ 
	Display: "clusters",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/clusters", "2019-03-01"),
},
{ 
	Display: "applications",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabricMesh/applications", "2018-09-01-preview"),
},
{ 
	Display: "gateways",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabricMesh/gateways", "2018-09-01-preview"),
},
{ 
	Display: "networks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabricMesh/networks", "2018-09-01-preview"),
},
{ 
	Display: "secrets",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabricMesh/secrets", "2018-09-01-preview"),
},
{ 
	Display: "volumes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabricMesh/volumes", "2018-09-01-preview"),
},
{ 
	Display: "SignalR",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.SignalRService/SignalR", "2018-10-01"),
},
{ 
	Display: "servers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Sql/servers", "2014-04-01"),
},
{ 
	Display: "sqlVirtualMachineGroups",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups", "2017-03-01-preview"),
},
{ 
	Display: "sqlVirtualMachines",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines", "2017-03-01-preview"),
},
{ 
	Display: "managers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.StorSimple/managers", "2017-06-01"),
},
{ 
	Display: "skus",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Storage/skus", "2019-06-01"),
},
{ 
	Display: "storageAccounts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Storage/storageAccounts", "2019-06-01"),
},
{ 
	Display: "caches",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.StorageCache/caches", "2019-11-01"),
},
{ 
	Display: "skus",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.StorageCache/skus", "2019-11-01"),
},
{ 
	Display: "usageModels",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.StorageCache/usageModels", "2019-11-01"),
},
{ 
	Display: "jobs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ImportExport/jobs", "2016-11-01"),
},
{ 
	Display: "storageSyncServices",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.StorageSync/storageSyncServices", "2019-06-01"),
},
{ 
	Display: "streamingjobs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.StreamAnalytics/streamingjobs", "2016-03-01"),
},
{ 
	Display: "supportTickets",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Support/supportTickets", "2019-05-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{supportTicketName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Support/supportTickets/{supportTicketName}", "2019-05-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Support/supportTickets/{supportTicketName}", "2019-05-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Support/supportTickets/{supportTicketName}", "2019-05-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "communications",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Support/supportTickets/{supportTicketName}/communications", "2019-05-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{communicationName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Support/supportTickets/{supportTicketName}/communications/{communicationName}", "2019-05-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Support/supportTickets/{supportTicketName}/communications/{communicationName}", "2019-05-01-preview"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "workspaces",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Synapse/workspaces", "2019-06-01-preview"),
},
{ 
	Display: "environments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.TimeSeriesInsights/environments", "2017-11-15"),
},
{ 
	Display: "default",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/trafficManagerUserMetricsKeys/default", "2018-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/trafficManagerUserMetricsKeys/default", "2018-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/trafficManagerUserMetricsKeys/default", "2018-04-01"),
},
{ 
	Display: "trafficmanagerprofiles",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/trafficmanagerprofiles", "2018-04-01"),
},
{ 
	Display: "dedicatedCloudNodes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes", "2019-04-01"),
},
{ 
	Display: "dedicatedCloudServices",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudServices", "2019-04-01"),
},
{ 
	Display: "virtualMachines",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/virtualMachines", "2019-04-01"),
},
{ 
	Display: "privateClouds",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareVirtustream/privateClouds", "2019-08-09-preview"),
},
{ 
	Display: "availableStacks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/availableStacks", "2018-02-01"),
},
{ 
	Display: "billingMeters",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/billingMeters", "2018-02-01"),
},
{ 
	Display: "certificates",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/certificates", "2018-02-01"),
},
{ 
	Display: "deploymentLocations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/deploymentLocations", "2018-02-01"),
},
{ 
	Display: "geoRegions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/geoRegions", "2018-02-01"),
},
{ 
	Display: "premieraddonoffers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/premieraddonoffers", "2018-02-01"),
},
{ 
	Display: "recommendations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/recommendations", "2018-02-01"),
	Children: []swagger.ResourceType{  } ,
	SubResources: []swagger.ResourceType{  } ,
},
{ 
	Display: "serverfarms",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/serverfarms", "2018-02-01"),
},
{ 
	Display: "sites",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/sites", "2018-02-01"),
},
{ 
	Display: "skus",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/skus", "2018-02-01"),
},
{ 
	Display: "multipleActivationKeys",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.WindowsESU/multipleActivationKeys", "2019-09-16-preview"),
},
{ 
	Display: "deviceServices",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.WindowsIoT/deviceServices", "2019-06-01"),
},
{ 
	Display: "componentsSummary",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.WorkloadMonitor/componentsSummary", "2018-08-31-preview"),
},
{ 
	Display: "monitorInstancesSummary",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.WorkloadMonitor/monitorInstancesSummary", "2018-08-31-preview"),
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{resourceProviderNamespace}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/{resourceProviderNamespace}", "2018-05-01"),
	Children: []swagger.ResourceType{  } ,
},
{ 
	Display: "{default}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.SerialConsole/consoleServices/{default}", "2018-05-01"),
	Children: []swagger.ResourceType{  } ,
},
{ 
	Display: "regions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceBus/sku/{sku}/regions", "2017-04-01"),
},
{ 
	Display: "clusterVersions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/clusterVersions", "2019-03-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{clusterVersion}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/clusterVersions/{clusterVersion}", "2019-03-01"),
}, } ,
},
{ 
	Display: "clusterVersions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/environments/{environment}/clusterVersions", "2019-03-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{clusterVersion}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/environments/{environment}/clusterVersions/{clusterVersion}", "2019-03-01"),
}, } ,
},
{ 
	Display: "usages",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.SignalRService/locations/{location}/usages", "2018-10-01"),
},
{ 
	Display: "capabilities",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Sql/locations/{locationId}/capabilities", "2014-04-01"),
},
{ 
	Display: "usages",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Storage/locations/{location}/usages", "2019-06-01"),
},
{ 
	Display: "quotas",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.StreamAnalytics/locations/{location}/quotas", "2016-03-01"),
},
{ 
	Display: "availabilities",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/availabilities", "2019-04-01"),
},
{ 
	Display: "{operationId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/operationResults/{operationId}", "2019-04-01"),
},
{ 
	Display: "privateClouds",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds", "2019-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{pcName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}", "2019-04-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "customizationPolicies",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/customizationPolicies", "2019-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{customizationPolicyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/customizationPolicies/{customizationPolicyName}", "2019-04-01"),
}, } ,
},
{ 
	Display: "resourcePools",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/resourcePools", "2019-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{resourcePoolName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/resourcePools/{resourcePoolName}", "2019-04-01"),
}, } ,
},
{ 
	Display: "virtualMachineTemplates",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/virtualMachineTemplates", "2019-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{virtualMachineTemplateName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/virtualMachineTemplates/{virtualMachineTemplateName}", "2019-04-01"),
}, } ,
},
{ 
	Display: "virtualNetworks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/virtualNetworks", "2019-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{virtualNetworkName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/virtualNetworks/{virtualNetworkName}", "2019-04-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "usages",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/usages", "2019-04-01"),
}, } ,
},
{ 
	Display: "resourcegroups",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups", "2018-05-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{resourceGroupName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}", "2018-05-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}", "2018-05-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}", "2018-05-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}", "2018-05-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "deployments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Resources/deployments", "2018-05-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{deploymentName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Resources/deployments/{deploymentName}", "2018-05-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Resources/deployments/{deploymentName}", "2018-05-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Resources/deployments/{deploymentName}", "2018-05-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "clusters",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters", "2019-03-01"),
},
{ 
	Display: "caches",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{cacheName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}", "2019-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}", "2019-11-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "storageTargets",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}/storageTargets", "2019-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{storageTargetName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}/storageTargets/{storageTargetName}", "2019-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}/storageTargets/{storageTargetName}", "2019-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}/storageTargets/{storageTargetName}", "2019-11-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "streamingjobs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs", "2016-03-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{jobName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}", "2016-03-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}", "2016-03-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}", "2016-03-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}", "2016-03-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "functions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/functions", "2016-03-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{functionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/functions/{functionName}", "2016-03-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/functions/{functionName}", "2016-03-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/functions/{functionName}", "2016-03-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/functions/{functionName}", "2016-03-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "inputs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/inputs", "2016-03-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{inputName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/inputs/{inputName}", "2016-03-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/inputs/{inputName}", "2016-03-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/inputs/{inputName}", "2016-03-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/inputs/{inputName}", "2016-03-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "outputs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/outputs", "2016-03-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{outputName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/outputs/{outputName}", "2016-03-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/outputs/{outputName}", "2016-03-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/outputs/{outputName}", "2016-03-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/outputs/{outputName}", "2016-03-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{transformationName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/transformations/{transformationName}", "2016-03-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/transformations/{transformationName}", "2016-03-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/transformations/{transformationName}", "2016-03-01"),
}, } ,
}, } ,
},
{ 
	Display: "account",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.visualstudio/account", "2014-04-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "extension",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.visualstudio/account/{accountResourceName}/extension", "2014-04-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{extensionResourceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.visualstudio/account/{accountResourceName}/extension/{extensionResourceName}", "2014-04-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.visualstudio/account/{accountResourceName}/extension/{extensionResourceName}", "2014-04-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.visualstudio/account/{accountResourceName}/extension/{extensionResourceName}", "2014-04-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.visualstudio/account/{accountResourceName}/extension/{extensionResourceName}", "2014-04-01-preview"),
}, } ,
},
{ 
	Display: "{resourceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.visualstudio/account/{resourceName}", "2014-04-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.visualstudio/account/{resourceName}", "2014-04-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.visualstudio/account/{resourceName}", "2014-04-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.visualstudio/account/{resourceName}", "2014-04-01-preview"),
}, } ,
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/deployments/{deploymentName}/operations", "2018-05-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{operationId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/deployments/{deploymentName}/operations/{operationId}", "2018-05-01"),
}, } ,
},
{ 
	Display: "{resourceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}", "2018-05-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}", "2018-05-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}", "2018-05-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}", "2018-05-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "locks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}/providers/Microsoft.Authorization/locks", "2016-09-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{lockName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}/providers/Microsoft.Authorization/locks/{lockName}", "2016-09-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}/providers/Microsoft.Authorization/locks/{lockName}", "2016-09-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}/providers/Microsoft.Authorization/locks/{lockName}", "2016-09-01"),
}, } ,
},
{ 
	Display: "policyAssignments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}/providers/Microsoft.Authorization/policyAssignments", "2016-12-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "tagNames",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/tagNames", "2018-05-01"),
	SubResources: []swagger.ResourceType{  } ,
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "locks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Authorization/locks", "2016-09-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{lockName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Authorization/locks/{lockName}", "2016-09-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Authorization/locks/{lockName}", "2016-09-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Authorization/locks/{lockName}", "2016-09-01"),
}, } ,
},
{ 
	Display: "policyAssignments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Authorization/policyAssignments", "2016-12-01"),
},
{ 
	Display: "resources",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/resources", "2018-05-01"),
},
{ 
	Display: "jobCollections",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections", "2016-03-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{jobCollectionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}", "2016-03-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}", "2016-03-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}", "2016-03-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}", "2016-03-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "jobs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs", "2016-03-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{jobName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs/{jobName}", "2016-03-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs/{jobName}", "2016-03-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs/{jobName}", "2016-03-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs/{jobName}", "2016-03-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "history",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs/{jobName}/history", "2016-03-01"),
}, } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "searchServices",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices", "2015-08-19"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{searchServiceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}", "2015-08-19"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}", "2015-08-19"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}", "2015-08-19"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}", "2015-08-19"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "listAdminKeys",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/listAdminKeys", "2015-08-19"),
	Verb: "POST",
},
{ 
	Display: "listQueryKeys",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/listQueryKeys", "2015-08-19"),
}, } ,
	SubResources: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "alerts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/alerts", "2019-01-01"),
},
{ 
	Display: "automations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/automations", "2019-01-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{automationName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/automations/{automationName}", "2019-01-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/automations/{automationName}", "2019-01-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/automations/{automationName}", "2019-01-01-preview"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "iotSecuritySolutions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions", "2019-08-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{solutionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}", "2019-08-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}", "2019-08-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}", "2019-08-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}", "2019-08-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "analyticsModels",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}/analyticsModels", "2019-08-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "default",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}/analyticsModels/default", "2019-08-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "aggregatedAlerts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}/analyticsModels/default/aggregatedAlerts", "2019-08-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{aggregatedAlertName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}/analyticsModels/default/aggregatedAlerts/{aggregatedAlertName}", "2019-08-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "aggregatedRecommendations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}/analyticsModels/default/aggregatedRecommendations", "2019-08-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{aggregatedRecommendationName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}/analyticsModels/default/aggregatedRecommendations/{aggregatedRecommendationName}", "2019-08-01"),
}, } ,
}, } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "jitNetworkAccessPolicies",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/jitNetworkAccessPolicies", "2015-06-01-preview"),
},
{ 
	Display: "{externalSecuritySolutionsName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/locations/{ascLocation}/ExternalSecuritySolutions/{externalSecuritySolutionsName}", "2015-06-01-preview"),
},
{ 
	Display: "alerts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/locations/{ascLocation}/alerts", "2019-01-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{alertName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/locations/{ascLocation}/alerts/{alertName}", "2019-01-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "{connectionType}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/locations/{ascLocation}/allowedConnections/{connectionType}", "2015-06-01-preview"),
},
{ 
	Display: "{discoveredSecuritySolutionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/locations/{ascLocation}/discoveredSecuritySolutions/{discoveredSecuritySolutionName}", "2015-06-01-preview"),
},
{ 
	Display: "jitNetworkAccessPolicies",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/locations/{ascLocation}/jitNetworkAccessPolicies", "2015-06-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{jitNetworkAccessPolicyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/locations/{ascLocation}/jitNetworkAccessPolicies/{jitNetworkAccessPolicyName}", "2015-06-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/locations/{ascLocation}/jitNetworkAccessPolicies/{jitNetworkAccessPolicyName}", "2015-06-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/locations/{ascLocation}/jitNetworkAccessPolicies/{jitNetworkAccessPolicyName}", "2015-06-01-preview"),
	SubResources: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "tasks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/locations/{ascLocation}/tasks", "2015-06-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{taskName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/locations/{ascLocation}/tasks/{taskName}", "2015-06-01-preview"),
	SubResources: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "{topologyResourceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/locations/{ascLocation}/topologies/{topologyResourceName}", "2015-06-01-preview"),
},
{ 
	Display: "serverVulnerabilityAssessments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceNamespace}/{resourceType}/{resourceName}/providers/Microsoft.Security/serverVulnerabilityAssessments", "2019-01-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{serverVulnerabilityAssessment}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceNamespace}/{resourceType}/{resourceName}/providers/Microsoft.Security/serverVulnerabilityAssessments/{serverVulnerabilityAssessment}", "2019-01-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceNamespace}/{resourceType}/{resourceName}/providers/Microsoft.Security/serverVulnerabilityAssessments/{serverVulnerabilityAssessment}", "2019-01-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceNamespace}/{resourceType}/{resourceName}/providers/Microsoft.Security/serverVulnerabilityAssessments/{serverVulnerabilityAssessment}", "2019-01-01-preview"),
}, } ,
},
{ 
	Display: "alertRules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/alertRules", "2020-01-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{ruleId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/alertRules/{ruleId}", "2020-01-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/alertRules/{ruleId}", "2020-01-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/alertRules/{ruleId}", "2020-01-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "actions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/alertRules/{ruleId}/actions", "2020-01-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{actionId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/alertRules/{ruleId}/actions/{actionId}", "2020-01-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/alertRules/{ruleId}/actions/{actionId}", "2020-01-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/alertRules/{ruleId}/actions/{actionId}", "2020-01-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "dataConnectors",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/dataConnectors", "2020-01-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{dataConnectorId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/dataConnectors/{dataConnectorId}", "2020-01-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/dataConnectors/{dataConnectorId}", "2020-01-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/dataConnectors/{dataConnectorId}", "2020-01-01"),
}, } ,
},
{ 
	Display: "{clientGroupName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/clientGroups/{clientGroupName}", "2015-11-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "members",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/clientGroups/{clientGroupName}/members", "2015-11-01-preview"),
},
{ 
	Display: "membersCount",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/clientGroups/{clientGroupName}/membersCount", "2015-11-01-preview"),
}, } ,
},
{ 
	Display: "machineGroups",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/machineGroups", "2015-11-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{machineGroupName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/machineGroups/{machineGroupName}", "2015-11-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/machineGroups/{machineGroupName}", "2015-11-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/machineGroups/{machineGroupName}", "2015-11-01-preview"),
}, } ,
},
{ 
	Display: "machines",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/machines", "2015-11-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{machineName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/machines/{machineName}", "2015-11-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "connections",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/machines/{machineName}/connections", "2015-11-01-preview"),
},
{ 
	Display: "liveness",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/machines/{machineName}/liveness", "2015-11-01-preview"),
},
{ 
	Display: "machineGroups",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/machines/{machineName}/machineGroups", "2015-11-01-preview"),
},
{ 
	Display: "ports",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/machines/{machineName}/ports", "2015-11-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{portName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/machines/{machineName}/ports/{portName}", "2015-11-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "acceptingProcesses",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/machines/{machineName}/ports/{portName}/acceptingProcesses", "2015-11-01-preview"),
},
{ 
	Display: "connections",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/machines/{machineName}/ports/{portName}/connections", "2015-11-01-preview"),
},
{ 
	Display: "liveness",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/machines/{machineName}/ports/{portName}/liveness", "2015-11-01-preview"),
}, } ,
}, } ,
},
{ 
	Display: "processes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/machines/{machineName}/processes", "2015-11-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{processName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/machines/{machineName}/processes/{processName}", "2015-11-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "acceptingPorts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/machines/{machineName}/processes/{processName}/acceptingPorts", "2015-11-01-preview"),
},
{ 
	Display: "connections",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/machines/{machineName}/processes/{processName}/connections", "2015-11-01-preview"),
},
{ 
	Display: "liveness",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/machines/{machineName}/processes/{processName}/liveness", "2015-11-01-preview"),
}, } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "machines",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/summaries/machines", "2015-11-01-preview"),
},
{ 
	Display: "namespaces",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces", "2017-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{namespaceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}", "2017-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}", "2017-04-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}", "2017-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}", "2017-04-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "AuthorizationRules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/AuthorizationRules", "2017-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{authorizationRuleName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "disasterRecoveryConfigs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/disasterRecoveryConfigs", "2017-04-01"),
	Children: []swagger.ResourceType{  } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{alias}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}", "2017-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}", "2017-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}", "2017-04-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "AuthorizationRules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}/AuthorizationRules", "2017-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{authorizationRuleName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "eventhubs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/eventhubs", "2017-04-01"),
},
{ 
	Display: "migrationConfigurations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/migrationConfigurations", "2017-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{configName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/migrationConfigurations/{configName}", "2017-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/migrationConfigurations/{configName}", "2017-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/migrationConfigurations/{configName}", "2017-04-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "networkRuleSets",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/networkRuleSets", "2017-04-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "default",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/networkRuleSets/default", "2017-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/networkRuleSets/default", "2017-04-01"),
}, } ,
},
{ 
	Display: "queues",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues", "2017-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{queueName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}", "2017-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}", "2017-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}", "2017-04-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "authorizationRules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}/authorizationRules", "2017-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{authorizationRuleName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "topics",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics", "2017-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{topicName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}", "2017-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}", "2017-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}", "2017-04-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "authorizationRules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/authorizationRules", "2017-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{authorizationRuleName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "subscriptions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions", "2017-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{subscriptionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions/{subscriptionName}", "2017-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions/{subscriptionName}", "2017-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions/{subscriptionName}", "2017-04-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "rules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions/{subscriptionName}/rules", "2017-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{ruleName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions/{subscriptionName}/rules/{ruleName}", "2017-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions/{subscriptionName}/rules/{ruleName}", "2017-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions/{subscriptionName}/rules/{ruleName}", "2017-04-01"),
}, } ,
}, } ,
}, } ,
}, } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "{clusterName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}", "2019-03-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}", "2019-03-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}", "2019-03-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}", "2019-03-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "applicationTypes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes", "2019-03-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{applicationTypeName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes/{applicationTypeName}", "2019-03-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes/{applicationTypeName}", "2019-03-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes/{applicationTypeName}", "2019-03-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "versions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes/{applicationTypeName}/versions", "2019-03-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{version}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes/{applicationTypeName}/versions/{version}", "2019-03-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes/{applicationTypeName}/versions/{version}", "2019-03-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes/{applicationTypeName}/versions/{version}", "2019-03-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "applications",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applications", "2019-03-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{applicationName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applications/{applicationName}", "2019-03-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applications/{applicationName}", "2019-03-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applications/{applicationName}", "2019-03-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applications/{applicationName}", "2019-03-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "services",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applications/{applicationName}/services", "2019-03-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{serviceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applications/{applicationName}/services/{serviceName}", "2019-03-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applications/{applicationName}/services/{serviceName}", "2019-03-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applications/{applicationName}/services/{serviceName}", "2019-03-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applications/{applicationName}/services/{serviceName}", "2019-03-01"),
}, } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "applications",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/applications", "2018-09-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{applicationResourceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/applications/{applicationResourceName}", "2018-09-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/applications/{applicationResourceName}", "2018-09-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/applications/{applicationResourceName}", "2018-09-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "services",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/applications/{applicationResourceName}/services", "2018-09-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{serviceResourceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/applications/{applicationResourceName}/services/{serviceResourceName}", "2018-09-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "replicas",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/applications/{applicationResourceName}/services/{serviceResourceName}/replicas", "2018-09-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{replicaName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/applications/{applicationResourceName}/services/{serviceResourceName}/replicas/{replicaName}", "2018-09-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "logs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/applications/{applicationResourceName}/services/{serviceResourceName}/replicas/{replicaName}/codePackages/{codePackageName}/logs", "2018-09-01-preview"),
}, } ,
}, } ,
}, } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "gateways",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/gateways", "2018-09-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{gatewayResourceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/gateways/{gatewayResourceName}", "2018-09-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/gateways/{gatewayResourceName}", "2018-09-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/gateways/{gatewayResourceName}", "2018-09-01-preview"),
}, } ,
},
{ 
	Display: "networks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/networks", "2018-09-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{networkResourceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/networks/{networkResourceName}", "2018-09-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/networks/{networkResourceName}", "2018-09-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/networks/{networkResourceName}", "2018-09-01-preview"),
}, } ,
},
{ 
	Display: "secrets",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/secrets", "2018-09-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{secretResourceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/secrets/{secretResourceName}", "2018-09-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/secrets/{secretResourceName}", "2018-09-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/secrets/{secretResourceName}", "2018-09-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "values",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/secrets/{secretResourceName}/values", "2018-09-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{secretValueResourceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/secrets/{secretResourceName}/values/{secretValueResourceName}", "2018-09-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/secrets/{secretResourceName}/values/{secretValueResourceName}", "2018-09-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/secrets/{secretResourceName}/values/{secretValueResourceName}", "2018-09-01-preview"),
	Children: []swagger.ResourceType{  } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "volumes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/volumes", "2018-09-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{volumeResourceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/volumes/{volumeResourceName}", "2018-09-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/volumes/{volumeResourceName}", "2018-09-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/volumes/{volumeResourceName}", "2018-09-01-preview"),
}, } ,
},
{ 
	Display: "SignalR",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/SignalR", "2018-10-01"),
	SubResources: []swagger.ResourceType{  } ,
},
{ 
	Display: "{resourceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/signalR/{resourceName}", "2018-10-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/signalR/{resourceName}", "2018-10-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/signalR/{resourceName}", "2018-10-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/signalR/{resourceName}", "2018-10-01"),
	Children: []swagger.ResourceType{  } ,
},
{ 
	Display: "servers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers", "2014-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{serverName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}", "2014-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}", "2014-04-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}", "2014-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}", "2014-04-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "administrators",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/administrators", "2014-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{administratorName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/administrators/{administratorName}", "2014-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/administrators/{administratorName}", "2014-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/administrators/{administratorName}", "2014-04-01"),
}, } ,
},
{ 
	Display: "advisors",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/advisors", "2014-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{advisorName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/advisors/{advisorName}", "2014-04-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/advisors/{advisorName}", "2014-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/advisors/{advisorName}", "2014-04-01"),
}, } ,
},
{ 
	Display: "auditingPolicies",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/auditingPolicies", "2014-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{tableAuditingPolicyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/auditingPolicies/{tableAuditingPolicyName}", "2014-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/auditingPolicies/{tableAuditingPolicyName}", "2014-04-01"),
}, } ,
},
{ 
	Display: "backupLongTermRetentionVaults",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/backupLongTermRetentionVaults", "2014-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{backupLongTermRetentionVaultName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/backupLongTermRetentionVaults/{backupLongTermRetentionVaultName}", "2014-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/backupLongTermRetentionVaults/{backupLongTermRetentionVaultName}", "2014-04-01"),
}, } ,
},
{ 
	Display: "communicationLinks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/communicationLinks", "2014-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{communicationLinkName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/communicationLinks/{communicationLinkName}", "2014-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/communicationLinks/{communicationLinkName}", "2014-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/communicationLinks/{communicationLinkName}", "2014-04-01"),
}, } ,
},
{ 
	Display: "databases",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases", "2014-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{databaseName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}", "2014-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}", "2014-04-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}", "2014-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}", "2014-04-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "advisors",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/advisors", "2014-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{advisorName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/advisors/{advisorName}", "2014-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/advisors/{advisorName}", "2014-04-01"),
}, } ,
},
{ 
	Display: "auditingPolicies",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/auditingPolicies", "2014-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{tableAuditingPolicyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/auditingPolicies/{tableAuditingPolicyName}", "2014-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/auditingPolicies/{tableAuditingPolicyName}", "2014-04-01"),
}, } ,
},
{ 
	Display: "backupLongTermRetentionPolicies",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/backupLongTermRetentionPolicies", "2014-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{backupLongTermRetentionPolicyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/backupLongTermRetentionPolicies/{backupLongTermRetentionPolicyName}", "2014-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/backupLongTermRetentionPolicies/{backupLongTermRetentionPolicyName}", "2014-04-01"),
}, } ,
},
{ 
	Display: "extensions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/extensions", "2014-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{extensionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/extensions/{extensionName}", "2014-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/extensions/{extensionName}", "2014-04-01"),
}, } ,
},
{ 
	Display: "geoBackupPolicies",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/geoBackupPolicies", "2014-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{geoBackupPolicyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/geoBackupPolicies/{geoBackupPolicyName}", "2014-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/geoBackupPolicies/{geoBackupPolicyName}", "2014-04-01"),
}, } ,
},
{ 
	Display: "metricDefinitions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/metricDefinitions", "2014-04-01"),
},
{ 
	Display: "metrics",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/metrics", "2014-04-01"),
},
{ 
	Display: "replicationLinks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/replicationLinks", "2014-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{linkId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/replicationLinks/{linkId}", "2014-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/replicationLinks/{linkId}", "2014-04-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "restorePoints",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/restorePoints", "2014-04-01"),
},
{ 
	Display: "serviceTierAdvisors",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/serviceTierAdvisors", "2014-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{serviceTierAdvisorName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/serviceTierAdvisors/{serviceTierAdvisorName}", "2014-04-01"),
}, } ,
},
{ 
	Display: "topQueries",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/topQueries", "2014-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "queryText",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/topQueries/{queryId}/queryText", "2014-04-01"),
},
{ 
	Display: "statistics",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/topQueries/{queryId}/statistics", "2014-04-01"),
}, } ,
},
{ 
	Display: "transparentDataEncryption",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/transparentDataEncryption", "2014-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{transparentDataEncryptionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/transparentDataEncryption/{transparentDataEncryptionName}", "2014-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/transparentDataEncryption/{transparentDataEncryptionName}", "2014-04-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "operationResults",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/transparentDataEncryption/{transparentDataEncryptionName}/operationResults", "2014-04-01"),
}, } ,
}, } ,
},
{ 
	Display: "usages",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/usages", "2014-04-01"),
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{connectionPolicyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/connectionPolicies/{connectionPolicyName}", "2014-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/connectionPolicies/{connectionPolicyName}", "2014-04-01"),
},
{ 
	Display: "{dataMaskingPolicyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/dataMaskingPolicies/{dataMaskingPolicyName}", "2014-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/dataMaskingPolicies/{dataMaskingPolicyName}", "2014-04-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "rules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/dataMaskingPolicies/{dataMaskingPolicyName}/rules", "2014-04-01"),
	SubResources: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "{securityAlertPolicyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/securityAlertPolicies/{securityAlertPolicyName}", "2014-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/securityAlertPolicies/{securityAlertPolicyName}", "2014-04-01"),
}, } ,
}, } ,
},
{ 
	Display: "disasterRecoveryConfiguration",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/disasterRecoveryConfiguration", "2014-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{disasterRecoveryConfigurationName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/disasterRecoveryConfiguration/{disasterRecoveryConfigurationName}", "2014-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/disasterRecoveryConfiguration/{disasterRecoveryConfigurationName}", "2014-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/disasterRecoveryConfiguration/{disasterRecoveryConfigurationName}", "2014-04-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "elasticPools",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools", "2014-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{elasticPoolName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}", "2014-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}", "2014-04-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}", "2014-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}", "2014-04-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "databases",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}/databases", "2014-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{databaseName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}/databases/{databaseName}", "2014-04-01"),
}, } ,
},
{ 
	Display: "elasticPoolActivity",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}/elasticPoolActivity", "2014-04-01"),
},
{ 
	Display: "elasticPoolDatabaseActivity",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}/elasticPoolDatabaseActivity", "2014-04-01"),
},
{ 
	Display: "metricDefinitions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}/metricDefinitions", "2014-04-01"),
},
{ 
	Display: "metrics",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}/metrics", "2014-04-01"),
}, } ,
}, } ,
},
{ 
	Display: "firewallRules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/firewallRules", "2014-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{firewallRuleName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/firewallRules/{firewallRuleName}", "2014-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/firewallRules/{firewallRuleName}", "2014-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/firewallRules/{firewallRuleName}", "2014-04-01"),
}, } ,
},
{ 
	Display: "recommendedElasticPools",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/recommendedElasticPools", "2014-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{recommendedElasticPoolName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/recommendedElasticPools/{recommendedElasticPoolName}", "2014-04-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "databases",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/recommendedElasticPools/{recommendedElasticPoolName}/databases", "2014-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{databaseName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/recommendedElasticPools/{recommendedElasticPoolName}/databases/{databaseName}", "2014-04-01"),
}, } ,
},
{ 
	Display: "metrics",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/recommendedElasticPools/{recommendedElasticPoolName}/metrics", "2014-04-01"),
}, } ,
}, } ,
},
{ 
	Display: "recoverableDatabases",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/recoverableDatabases", "2014-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{databaseName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/recoverableDatabases/{databaseName}", "2014-04-01"),
}, } ,
},
{ 
	Display: "restorableDroppedDatabases",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/restorableDroppedDatabases", "2014-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{restorableDroppededDatabaseId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/restorableDroppedDatabases/{restorableDroppededDatabaseId}", "2014-04-01"),
}, } ,
},
{ 
	Display: "serviceObjectives",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/serviceObjectives", "2014-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{serviceObjectiveName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/serviceObjectives/{serviceObjectiveName}", "2014-04-01"),
}, } ,
},
{ 
	Display: "usages",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/usages", "2014-04-01"),
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{connectionPolicyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/connectionPolicies/{connectionPolicyName}", "2014-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/connectionPolicies/{connectionPolicyName}", "2014-04-01"),
}, } ,
}, } ,
},
{ 
	Display: "sqlVirtualMachineGroups",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups", "2017-03-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{sqlVirtualMachineGroupName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/{sqlVirtualMachineGroupName}", "2017-03-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/{sqlVirtualMachineGroupName}", "2017-03-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/{sqlVirtualMachineGroupName}", "2017-03-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/{sqlVirtualMachineGroupName}", "2017-03-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "availabilityGroupListeners",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/{sqlVirtualMachineGroupName}/availabilityGroupListeners", "2017-03-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{availabilityGroupListenerName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/{sqlVirtualMachineGroupName}/availabilityGroupListeners/{availabilityGroupListenerName}", "2017-03-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/{sqlVirtualMachineGroupName}/availabilityGroupListeners/{availabilityGroupListenerName}", "2017-03-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/{sqlVirtualMachineGroupName}/availabilityGroupListeners/{availabilityGroupListenerName}", "2017-03-01-preview"),
}, } ,
},
{ 
	Display: "sqlVirtualMachines",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/{sqlVirtualMachineGroupName}/sqlVirtualMachines", "2017-03-01-preview"),
}, } ,
}, } ,
},
{ 
	Display: "sqlVirtualMachines",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines", "2017-03-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{sqlVirtualMachineName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/{sqlVirtualMachineName}", "2017-03-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/{sqlVirtualMachineName}", "2017-03-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/{sqlVirtualMachineName}", "2017-03-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/{sqlVirtualMachineName}", "2017-03-01-preview"),
}, } ,
},
{ 
	Display: "managers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers", "2017-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{managerName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}", "2017-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}", "2017-06-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}", "2017-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}", "2017-06-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "accessControlRecords",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/accessControlRecords", "2017-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{accessControlRecordName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/accessControlRecords/{accessControlRecordName}", "2017-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/accessControlRecords/{accessControlRecordName}", "2017-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/accessControlRecords/{accessControlRecordName}", "2017-06-01"),
}, } ,
},
{ 
	Display: "alerts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/alerts", "2017-06-01"),
},
{ 
	Display: "backups",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/backups", "2016-10-01"),
},
{ 
	Display: "devices",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices", "2017-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{deviceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}", "2017-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}", "2017-06-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}", "2017-06-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "default",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/alertSettings/default", "2017-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/alertSettings/default", "2017-06-01"),
},
{ 
	Display: "backupScheduleGroups",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupScheduleGroups", "2016-10-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{scheduleGroupName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupScheduleGroups/{scheduleGroupName}", "2016-10-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupScheduleGroups/{scheduleGroupName}", "2016-10-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupScheduleGroups/{scheduleGroupName}", "2016-10-01"),
}, } ,
},
{ 
	Display: "backups",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backups", "2017-06-01"),
	SubResources: []swagger.ResourceType{  } ,
},
{ 
	Display: "chapSettings",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/chapSettings", "2016-10-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{chapUserName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/chapSettings/{chapUserName}", "2016-10-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/chapSettings/{chapUserName}", "2016-10-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/chapSettings/{chapUserName}", "2016-10-01"),
}, } ,
},
{ 
	Display: "disks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/disks", "2016-10-01"),
},
{ 
	Display: "failoverTargets",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/failoverTargets", "2016-10-01"),
},
{ 
	Display: "fileservers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/fileservers", "2016-10-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{fileServerName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/fileservers/{fileServerName}", "2016-10-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/fileservers/{fileServerName}", "2016-10-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/fileservers/{fileServerName}", "2016-10-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "metrics",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/fileservers/{fileServerName}/metrics", "2016-10-01"),
},
{ 
	Display: "metricsDefinitions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/fileservers/{fileServerName}/metricsDefinitions", "2016-10-01"),
},
{ 
	Display: "shares",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/fileservers/{fileServerName}/shares", "2016-10-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{shareName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/fileservers/{fileServerName}/shares/{shareName}", "2016-10-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/fileservers/{fileServerName}/shares/{shareName}", "2016-10-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/fileservers/{fileServerName}/shares/{shareName}", "2016-10-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "metrics",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/fileservers/{fileServerName}/shares/{shareName}/metrics", "2016-10-01"),
},
{ 
	Display: "metricsDefinitions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/fileservers/{fileServerName}/shares/{shareName}/metricsDefinitions", "2016-10-01"),
}, } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "iscsiservers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers", "2016-10-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{iscsiServerName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}", "2016-10-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}", "2016-10-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}", "2016-10-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "disks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}/disks", "2016-10-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{diskName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}/disks/{diskName}", "2016-10-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}/disks/{diskName}", "2016-10-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}/disks/{diskName}", "2016-10-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "metrics",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}/disks/{diskName}/metrics", "2016-10-01"),
},
{ 
	Display: "metricsDefinitions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}/disks/{diskName}/metricsDefinitions", "2016-10-01"),
}, } ,
}, } ,
},
{ 
	Display: "metrics",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}/metrics", "2016-10-01"),
},
{ 
	Display: "metricsDefinitions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}/metricsDefinitions", "2016-10-01"),
}, } ,
}, } ,
},
{ 
	Display: "jobs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/jobs", "2017-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{jobName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/jobs/{jobName}", "2017-06-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "metrics",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/metrics", "2017-06-01"),
},
{ 
	Display: "metricsDefinitions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/metricsDefinitions", "2017-06-01"),
},
{ 
	Display: "default",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/networkSettings/default", "2017-06-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/networkSettings/default", "2017-06-01"),
},
{ 
	Display: "shares",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/shares", "2016-10-01"),
},
{ 
	Display: "default",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/timeSettings/default", "2017-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/timeSettings/default", "2017-06-01"),
},
{ 
	Display: "default",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/updateSummary/default", "2017-06-01"),
},
{ 
	Display: "backupPolicies",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies", "2017-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{backupPolicyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}", "2017-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}", "2017-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}", "2017-06-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "schedules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}/schedules", "2017-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{backupScheduleName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}/schedules/{backupScheduleName}", "2017-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}/schedules/{backupScheduleName}", "2017-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}/schedules/{backupScheduleName}", "2017-06-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "hardwareComponentGroups",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/hardwareComponentGroups", "2017-06-01"),
	SubResources: []swagger.ResourceType{  } ,
},
{ 
	Display: "default",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/securitySettings/default", "2017-06-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/securitySettings/default", "2017-06-01"),
	Children: []swagger.ResourceType{  } ,
},
{ 
	Display: "volumeContainers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers", "2017-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{volumeContainerName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers/{volumeContainerName}", "2017-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers/{volumeContainerName}", "2017-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers/{volumeContainerName}", "2017-06-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "metrics",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers/{volumeContainerName}/metrics", "2017-06-01"),
},
{ 
	Display: "metricsDefinitions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers/{volumeContainerName}/metricsDefinitions", "2017-06-01"),
},
{ 
	Display: "volumes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers/{volumeContainerName}/volumes", "2017-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{volumeName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers/{volumeContainerName}/volumes/{volumeName}", "2017-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers/{volumeContainerName}/volumes/{volumeName}", "2017-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers/{volumeContainerName}/volumes/{volumeName}", "2017-06-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "metrics",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers/{volumeContainerName}/volumes/{volumeName}/metrics", "2017-06-01"),
},
{ 
	Display: "metricsDefinitions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers/{volumeContainerName}/volumes/{volumeName}/metricsDefinitions", "2017-06-01"),
}, } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "volumes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumes", "2017-06-01"),
}, } ,
}, } ,
},
{ 
	Display: "default",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/encryptionSettings/default", "2017-06-01"),
},
{ 
	Display: "vaultExtendedInfo",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/extendedInformation/vaultExtendedInfo", "2017-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/extendedInformation/vaultExtendedInfo", "2017-06-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/extendedInformation/vaultExtendedInfo", "2017-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/extendedInformation/vaultExtendedInfo", "2017-06-01"),
},
{ 
	Display: "fileservers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/fileservers", "2016-10-01"),
},
{ 
	Display: "iscsiservers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/iscsiservers", "2016-10-01"),
},
{ 
	Display: "jobs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/jobs", "2017-06-01"),
},
{ 
	Display: "metrics",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/metrics", "2017-06-01"),
},
{ 
	Display: "metricsDefinitions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/metricsDefinitions", "2017-06-01"),
},
{ 
	Display: "storageAccountCredentials",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageAccountCredentials", "2017-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{credentialName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageAccountCredentials/{credentialName}", "2016-10-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageAccountCredentials/{credentialName}", "2016-10-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageAccountCredentials/{credentialName}", "2016-10-01"),
},
{ 
	Display: "{storageAccountCredentialName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageAccountCredentials/{storageAccountCredentialName}", "2017-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageAccountCredentials/{storageAccountCredentialName}", "2017-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageAccountCredentials/{storageAccountCredentialName}", "2017-06-01"),
}, } ,
},
{ 
	Display: "storageDomains",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageDomains", "2016-10-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{storageDomainName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageDomains/{storageDomainName}", "2016-10-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageDomains/{storageDomainName}", "2016-10-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageDomains/{storageDomainName}", "2016-10-01"),
}, } ,
},
{ 
	Display: "bandwidthSettings",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/bandwidthSettings", "2017-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{bandwidthSettingName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/bandwidthSettings/{bandwidthSettingName}", "2017-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/bandwidthSettings/{bandwidthSettingName}", "2017-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/bandwidthSettings/{bandwidthSettingName}", "2017-06-01"),
}, } ,
},
{ 
	Display: "cloudApplianceConfigurations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/cloudApplianceConfigurations", "2017-06-01"),
},
{ 
	Display: "features",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/features", "2017-06-01"),
}, } ,
	SubResources: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "storageAccounts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts", "2019-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{accountName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}", "2019-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}", "2019-06-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}", "2019-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}", "2019-06-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "blobServices",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices", "2019-06-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "containers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers", "2019-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{containerName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}", "2019-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}", "2019-06-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}", "2019-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}", "2019-06-01"),
	Children: []swagger.ResourceType{  } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{immutabilityPolicyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}/immutabilityPolicies/{immutabilityPolicyName}", "2019-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}/immutabilityPolicies/{immutabilityPolicyName}", "2019-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}/immutabilityPolicies/{immutabilityPolicyName}", "2019-06-01"),
}, } ,
}, } ,
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{BlobServicesName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/{BlobServicesName}", "2019-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/{BlobServicesName}", "2019-06-01"),
}, } ,
},
{ 
	Display: "encryptionScopes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/encryptionScopes", "2019-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{encryptionScopeName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/encryptionScopes/{encryptionScopeName}", "2019-06-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/encryptionScopes/{encryptionScopeName}", "2019-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/encryptionScopes/{encryptionScopeName}", "2019-06-01"),
}, } ,
},
{ 
	Display: "fileServices",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices", "2019-06-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "shares",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices/default/shares", "2019-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{shareName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices/default/shares/{shareName}", "2019-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices/default/shares/{shareName}", "2019-06-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices/default/shares/{shareName}", "2019-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices/default/shares/{shareName}", "2019-06-01"),
}, } ,
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{FileServicesName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices/{FileServicesName}", "2019-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices/{FileServicesName}", "2019-06-01"),
}, } ,
},
{ 
	Display: "privateLinkResources",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/privateLinkResources", "2019-06-01"),
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{managementPolicyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/managementPolicies/{managementPolicyName}", "2019-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/managementPolicies/{managementPolicyName}", "2019-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/managementPolicies/{managementPolicyName}", "2019-06-01"),
},
{ 
	Display: "{privateEndpointConnectionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/privateEndpointConnections/{privateEndpointConnectionName}", "2019-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/privateEndpointConnections/{privateEndpointConnectionName}", "2019-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/privateEndpointConnections/{privateEndpointConnectionName}", "2019-06-01"),
}, } ,
}, } ,
},
{ 
	Display: "jobs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ImportExport/jobs", "2016-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{jobName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ImportExport/jobs/{jobName}", "2016-11-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ImportExport/jobs/{jobName}", "2016-11-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ImportExport/jobs/{jobName}", "2016-11-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ImportExport/jobs/{jobName}", "2016-11-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "{operationId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/locations/{locationName}/workflows/{workflowId}/operations/{operationId}", "2019-06-01"),
},
{ 
	Display: "storageSyncServices",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices", "2019-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{storageSyncServiceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}", "2019-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}", "2019-06-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}", "2019-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}", "2019-06-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "registeredServers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/registeredServers", "2019-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{serverId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/registeredServers/{serverId}", "2019-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/registeredServers/{serverId}", "2019-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/registeredServers/{serverId}", "2019-06-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "syncGroups",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups", "2019-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{syncGroupName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}", "2019-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}", "2019-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}", "2019-06-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "cloudEndpoints",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/cloudEndpoints", "2019-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{cloudEndpointName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/cloudEndpoints/{cloudEndpointName}", "2019-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/cloudEndpoints/{cloudEndpointName}", "2019-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/cloudEndpoints/{cloudEndpointName}", "2019-06-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "serverEndpoints",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/serverEndpoints", "2019-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{serverEndpointName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/serverEndpoints/{serverEndpointName}", "2019-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/serverEndpoints/{serverEndpointName}", "2019-06-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/serverEndpoints/{serverEndpointName}", "2019-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/serverEndpoints/{serverEndpointName}", "2019-06-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "workflows",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/workflows", "2019-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{workflowId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/workflows/{workflowId}", "2019-06-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "workspaces",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces", "2019-06-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{workspaceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}", "2019-06-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}", "2019-06-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}", "2019-06-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}", "2019-06-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "activeDirectory",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/administrators/activeDirectory", "2019-06-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/administrators/activeDirectory", "2019-06-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/administrators/activeDirectory", "2019-06-01-preview"),
},
{ 
	Display: "bigDataPools",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/bigDataPools", "2019-06-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{bigDataPoolName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/bigDataPools/{bigDataPoolName}", "2019-06-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/bigDataPools/{bigDataPoolName}", "2019-06-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/bigDataPools/{bigDataPoolName}", "2019-06-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/bigDataPools/{bigDataPoolName}", "2019-06-01-preview"),
}, } ,
},
{ 
	Display: "firewallRules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/firewallRules", "2019-06-01-preview"),
	SubResources: []swagger.ResourceType{  } ,
},
{ 
	Display: "default",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/managedIdentitySqlControlSettings/default", "2019-06-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/managedIdentitySqlControlSettings/default", "2019-06-01-preview"),
},
{ 
	Display: "sqlPools",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools", "2019-06-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{sqlPoolName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}", "2019-06-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}", "2019-06-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}", "2019-06-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}", "2019-06-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "currentSensitivityLabels",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/currentSensitivityLabels", "2019-06-01-preview"),
},
{ 
	Display: "config",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/metadataSync/config", "2019-06-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/metadataSync/config", "2019-06-01-preview"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/operations", "2019-06-01-preview"),
},
{ 
	Display: "recommendedSensitivityLabels",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/recommendedSensitivityLabels", "2019-06-01-preview"),
},
{ 
	Display: "replicationLinks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/replicationLinks", "2019-06-01-preview"),
},
{ 
	Display: "restorePoints",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/restorePoints", "2019-06-01-preview"),
},
{ 
	Display: "schemas",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/schemas", "2019-06-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "tables",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/schemas/{schemaName}/tables", "2019-06-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "columns",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/schemas/{schemaName}/tables/{tableName}/columns", "2019-06-01-preview"),
	SubResources: []swagger.ResourceType{  } ,
}, } ,
}, } ,
},
{ 
	Display: "usages",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/usages", "2019-06-01-preview"),
},
{ 
	Display: "vulnerabilityAssessments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/vulnerabilityAssessments", "2019-06-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{vulnerabilityAssessmentName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/vulnerabilityAssessments/{vulnerabilityAssessmentName}", "2019-06-01-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/vulnerabilityAssessments/{vulnerabilityAssessmentName}", "2019-06-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/vulnerabilityAssessments/{vulnerabilityAssessmentName}", "2019-06-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "scans",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/vulnerabilityAssessments/{vulnerabilityAssessmentName}/scans", "2019-06-01-preview"),
	SubResources: []swagger.ResourceType{  } ,
}, } ,
	SubResources: []swagger.ResourceType{  } ,
}, } ,
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{blobAuditingPolicyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/auditingSettings/{blobAuditingPolicyName}", "2019-06-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/auditingSettings/{blobAuditingPolicyName}", "2019-06-01-preview"),
},
{ 
	Display: "{connectionPolicyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/connectionPolicies/{connectionPolicyName}", "2019-06-01-preview"),
},
{ 
	Display: "{dataWarehouseUserActivityName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/dataWarehouseUserActivities/{dataWarehouseUserActivityName}", "2019-06-01-preview"),
},
{ 
	Display: "{geoBackupPolicyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/geoBackupPolicies/{geoBackupPolicyName}", "2019-06-01-preview"),
},
{ 
	Display: "{operationId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/operationResults/{operationId}", "2019-06-01-preview"),
},
{ 
	Display: "{securityAlertPolicyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/securityAlertPolicies/{securityAlertPolicyName}", "2019-06-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/securityAlertPolicies/{securityAlertPolicyName}", "2019-06-01-preview"),
},
{ 
	Display: "{transparentDataEncryptionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/transparentDataEncryption/{transparentDataEncryptionName}", "2019-06-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/transparentDataEncryption/{transparentDataEncryptionName}", "2019-06-01-preview"),
}, } ,
}, } ,
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{operationId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/operationResults/{operationId}", "2019-06-01-preview"),
},
{ 
	Display: "{operationId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/operationStatuses/{operationId}", "2019-06-01-preview"),
}, } ,
}, } ,
},
{ 
	Display: "environments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments", "2017-11-15"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{environmentName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}", "2017-11-15"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}", "2017-11-15"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}", "2017-11-15"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}", "2017-11-15"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "accessPolicies",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/accessPolicies", "2017-11-15"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{accessPolicyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/accessPolicies/{accessPolicyName}", "2017-11-15"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/accessPolicies/{accessPolicyName}", "2017-11-15"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/accessPolicies/{accessPolicyName}", "2017-11-15"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/accessPolicies/{accessPolicyName}", "2017-11-15"),
}, } ,
},
{ 
	Display: "eventSources",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/eventSources", "2017-11-15"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{eventSourceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/eventSources/{eventSourceName}", "2017-11-15"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/eventSources/{eventSourceName}", "2017-11-15"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/eventSources/{eventSourceName}", "2017-11-15"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/eventSources/{eventSourceName}", "2017-11-15"),
}, } ,
},
{ 
	Display: "referenceDataSets",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/referenceDataSets", "2017-11-15"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{referenceDataSetName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/referenceDataSets/{referenceDataSetName}", "2017-11-15"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/referenceDataSets/{referenceDataSetName}", "2017-11-15"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/referenceDataSets/{referenceDataSetName}", "2017-11-15"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/referenceDataSets/{referenceDataSetName}", "2017-11-15"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "trafficmanagerprofiles",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles", "2018-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{profileName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}", "2018-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}", "2018-04-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}", "2018-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}", "2018-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{heatMapType}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}/heatMaps/{heatMapType}", "2018-04-01"),
},
{ 
	Display: "{endpointName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}/{endpointType}/{endpointName}", "2018-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}/{endpointType}/{endpointName}", "2018-04-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}/{endpointType}/{endpointName}", "2018-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}/{endpointType}/{endpointName}", "2018-04-01"),
}, } ,
}, } ,
},
{ 
	Display: "project",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.visualstudio/account/{rootResourceName}/project", "2014-04-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{resourceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.visualstudio/account/{rootResourceName}/project/{resourceName}", "2014-04-01-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.visualstudio/account/{rootResourceName}/project/{resourceName}", "2014-04-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.visualstudio/account/{rootResourceName}/project/{resourceName}", "2014-04-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "status",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.visualstudio/account/{rootResourceName}/project/{resourceName}/subContainers/{subContainerName}/status", "2014-04-01-preview"),
}, } ,
}, } ,
},
{ 
	Display: "dedicatedCloudNodes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes", "2019-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{dedicatedCloudNodeName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes/{dedicatedCloudNodeName}", "2019-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes/{dedicatedCloudNodeName}", "2019-04-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes/{dedicatedCloudNodeName}", "2019-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes/{dedicatedCloudNodeName}", "2019-04-01"),
}, } ,
},
{ 
	Display: "dedicatedCloudServices",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudServices", "2019-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{dedicatedCloudServiceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudServices/{dedicatedCloudServiceName}", "2019-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudServices/{dedicatedCloudServiceName}", "2019-04-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudServices/{dedicatedCloudServiceName}", "2019-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudServices/{dedicatedCloudServiceName}", "2019-04-01"),
}, } ,
},
{ 
	Display: "virtualMachines",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/virtualMachines", "2019-04-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{virtualMachineName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/virtualMachines/{virtualMachineName}", "2019-04-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/virtualMachines/{virtualMachineName}", "2019-04-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/virtualMachines/{virtualMachineName}", "2019-04-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/virtualMachines/{virtualMachineName}", "2019-04-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "privateClouds",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareVirtustream/privateClouds", "2019-08-09-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{privateCloudName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareVirtustream/privateClouds/{privateCloudName}", "2019-08-09-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareVirtustream/privateClouds/{privateCloudName}", "2019-08-09-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareVirtustream/privateClouds/{privateCloudName}", "2019-08-09-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareVirtustream/privateClouds/{privateCloudName}", "2019-08-09-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "clusters",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareVirtustream/privateClouds/{privateCloudName}/clusters", "2019-08-09-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{clusterName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareVirtustream/privateClouds/{privateCloudName}/clusters/{clusterName}", "2019-08-09-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareVirtustream/privateClouds/{privateCloudName}/clusters/{clusterName}", "2019-08-09-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareVirtustream/privateClouds/{privateCloudName}/clusters/{clusterName}", "2019-08-09-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareVirtustream/privateClouds/{privateCloudName}/clusters/{clusterName}", "2019-08-09-preview"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "certificates",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/certificates", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/certificates/{name}", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/certificates/{name}", "2018-02-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/certificates/{name}", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/certificates/{name}", "2018-02-01"),
}, } ,
},
{ 
	Display: "recommendationHistory",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{hostingEnvironmentName}/recommendationHistory", "2018-02-01"),
},
{ 
	Display: "recommendations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{hostingEnvironmentName}/recommendations", "2018-02-01"),
	Children: []swagger.ResourceType{  } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{hostingEnvironmentName}/recommendations/{name}", "2018-02-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "serverfarms",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}", "2018-02-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}", "2018-02-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "capabilities",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/capabilities", "2018-02-01"),
},
{ 
	Display: "limit",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/hybridConnectionPlanLimits/limit", "2018-02-01"),
},
{ 
	Display: "hybridConnectionRelays",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/hybridConnectionRelays", "2018-02-01"),
},
{ 
	Display: "metricdefinitions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/metricdefinitions", "2018-02-01"),
},
{ 
	Display: "metrics",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/metrics", "2018-02-01"),
},
{ 
	Display: "sites",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/sites", "2018-02-01"),
},
{ 
	Display: "skus",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/skus", "2018-02-01"),
},
{ 
	Display: "usages",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/usages", "2018-02-01"),
},
{ 
	Display: "virtualNetworkConnections",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/virtualNetworkConnections", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{vnetName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/virtualNetworkConnections/{vnetName}", "2018-02-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "routes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/virtualNetworkConnections/{vnetName}/routes", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{routeName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/virtualNetworkConnections/{vnetName}/routes/{routeName}", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/virtualNetworkConnections/{vnetName}/routes/{routeName}", "2018-02-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/virtualNetworkConnections/{vnetName}/routes/{routeName}", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/virtualNetworkConnections/{vnetName}/routes/{routeName}", "2018-02-01"),
}, } ,
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{gatewayName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/virtualNetworkConnections/{vnetName}/gateways/{gatewayName}", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/virtualNetworkConnections/{vnetName}/gateways/{gatewayName}", "2018-02-01"),
}, } ,
}, } ,
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{relayName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}", "2018-02-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "sites",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}/sites", "2018-02-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "sites",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}", "2018-02-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}", "2018-02-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "analyzeCustomHostname",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/analyzeCustomHostname", "2018-02-01"),
},
{ 
	Display: "backups",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/backups", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{backupId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/backups/{backupId}", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/backups/{backupId}", "2018-02-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "config",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config", "2018-02-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "appsettings",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/appsettings/list", "2018-02-01"),
	Verb: "POST",
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/appsettings", "2018-02-01"),
},
{ 
	Display: "authsettings",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/authsettings/list", "2018-02-01"),
	Verb: "POST",
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/authsettings", "2018-02-01"),
},
{ 
	Display: "azurestorageaccounts",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/azurestorageaccounts/list", "2018-02-01"),
	Verb: "POST",
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/azurestorageaccounts", "2018-02-01"),
},
{ 
	Display: "backup",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/backup/list", "2018-02-01"),
	Verb: "POST",
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/backup", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/backup", "2018-02-01"),
},
{ 
	Display: "connectionstrings",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/connectionstrings/list", "2018-02-01"),
	Verb: "POST",
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/connectionstrings", "2018-02-01"),
},
{ 
	Display: "logs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/logs", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/logs", "2018-02-01"),
},
{ 
	Display: "metadata",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/metadata/list", "2018-02-01"),
	Verb: "POST",
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/metadata", "2018-02-01"),
},
{ 
	Display: "publishingcredentials",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/publishingcredentials/list", "2018-02-01"),
	Verb: "POST",
},
{ 
	Display: "pushsettings",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/pushsettings/list", "2018-02-01"),
	Verb: "POST",
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/pushsettings", "2018-02-01"),
},
{ 
	Display: "slotConfigNames",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/slotConfigNames", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/slotConfigNames", "2018-02-01"),
},
{ 
	Display: "web",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/web", "2018-02-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/web", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/web", "2018-02-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "snapshots",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/web/snapshots", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{snapshotId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/web/snapshots/{snapshotId}", "2018-02-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "continuouswebjobs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/continuouswebjobs", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{webJobName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/continuouswebjobs/{webJobName}", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/continuouswebjobs/{webJobName}", "2018-02-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "deployments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/deployments", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{id}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/deployments/{id}", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/deployments/{id}", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/deployments/{id}", "2018-02-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "log",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/deployments/{id}/log", "2018-02-01"),
}, } ,
}, } ,
},
{ 
	Display: "domainOwnershipIdentifiers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/domainOwnershipIdentifiers", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{domainOwnershipIdentifierName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/domainOwnershipIdentifiers/{domainOwnershipIdentifierName}", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/domainOwnershipIdentifiers/{domainOwnershipIdentifierName}", "2018-02-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/domainOwnershipIdentifiers/{domainOwnershipIdentifierName}", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/domainOwnershipIdentifiers/{domainOwnershipIdentifierName}", "2018-02-01"),
}, } ,
},
{ 
	Display: "MSDeploy",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/extensions/MSDeploy", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/extensions/MSDeploy", "2018-02-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "log",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/extensions/MSDeploy/log", "2018-02-01"),
}, } ,
},
{ 
	Display: "functions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/functions", "2018-02-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "token",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/functions/admin/token", "2018-02-01"),
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{functionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/functions/{functionName}", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/functions/{functionName}", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/functions/{functionName}", "2018-02-01"),
	Children: []swagger.ResourceType{  } ,
	SubResources: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "hostNameBindings",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hostNameBindings", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{hostName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hostNameBindings/{hostName}", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hostNameBindings/{hostName}", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hostNameBindings/{hostName}", "2018-02-01"),
}, } ,
},
{ 
	Display: "hybridConnectionRelays",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hybridConnectionRelays", "2018-02-01"),
},
{ 
	Display: "hybridconnection",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hybridconnection", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{entityName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hybridconnection/{entityName}", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hybridconnection/{entityName}", "2018-02-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hybridconnection/{entityName}", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hybridconnection/{entityName}", "2018-02-01"),
}, } ,
},
{ 
	Display: "instances",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "MSDeploy",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/extensions/MSDeploy", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/extensions/MSDeploy", "2018-02-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "log",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/extensions/MSDeploy/log", "2018-02-01"),
}, } ,
},
{ 
	Display: "processes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{processId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes/{processId}", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes/{processId}", "2018-02-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "dump",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes/{processId}/dump", "2018-02-01"),
},
{ 
	Display: "modules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes/{processId}/modules", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{baseAddress}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes/{processId}/modules/{baseAddress}", "2018-02-01"),
}, } ,
},
{ 
	Display: "threads",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes/{processId}/threads", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{threadId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes/{processId}/threads/{threadId}", "2018-02-01"),
}, } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "metricdefinitions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/metricdefinitions", "2018-02-01"),
},
{ 
	Display: "metrics",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/metrics", "2018-02-01"),
},
{ 
	Display: "virtualNetwork",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/networkConfig/virtualNetwork", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/networkConfig/virtualNetwork", "2018-02-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/networkConfig/virtualNetwork", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/networkConfig/virtualNetwork", "2018-02-01"),
},
{ 
	Display: "perfcounters",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/perfcounters", "2018-02-01"),
},
{ 
	Display: "phplogging",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/phplogging", "2018-02-01"),
},
{ 
	Display: "premieraddons",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/premieraddons", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{premierAddOnName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/premieraddons/{premierAddOnName}", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/premieraddons/{premierAddOnName}", "2018-02-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/premieraddons/{premierAddOnName}", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/premieraddons/{premierAddOnName}", "2018-02-01"),
}, } ,
},
{ 
	Display: "virtualNetworks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/privateAccess/virtualNetworks", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/privateAccess/virtualNetworks", "2018-02-01"),
},
{ 
	Display: "processes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/processes", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{processId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/processes/{processId}", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/processes/{processId}", "2018-02-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "dump",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/processes/{processId}/dump", "2018-02-01"),
},
{ 
	Display: "modules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/processes/{processId}/modules", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{baseAddress}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/processes/{processId}/modules/{baseAddress}", "2018-02-01"),
}, } ,
},
{ 
	Display: "threads",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/processes/{processId}/threads", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{threadId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/processes/{processId}/threads/{threadId}", "2018-02-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "publicCertificates",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/publicCertificates", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{publicCertificateName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/publicCertificates/{publicCertificateName}", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/publicCertificates/{publicCertificateName}", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/publicCertificates/{publicCertificateName}", "2018-02-01"),
}, } ,
},
{ 
	Display: "siteextensions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/siteextensions", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{siteExtensionId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/siteextensions/{siteExtensionId}", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/siteextensions/{siteExtensionId}", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/siteextensions/{siteExtensionId}", "2018-02-01"),
}, } ,
},
{ 
	Display: "slots",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{slot}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}", "2018-02-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}", "2018-02-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "analyzeCustomHostname",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/analyzeCustomHostname", "2018-02-01"),
},
{ 
	Display: "backups",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/backups", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{backupId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/backups/{backupId}", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/backups/{backupId}", "2018-02-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "config",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config", "2018-02-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "logs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/logs", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/logs", "2018-02-01"),
},
{ 
	Display: "web",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/web", "2018-02-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/web", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/web", "2018-02-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "snapshots",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/web/snapshots", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{snapshotId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/web/snapshots/{snapshotId}", "2018-02-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "continuouswebjobs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/continuouswebjobs", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{webJobName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/continuouswebjobs/{webJobName}", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/continuouswebjobs/{webJobName}", "2018-02-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "deployments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/deployments", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{id}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/deployments/{id}", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/deployments/{id}", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/deployments/{id}", "2018-02-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "log",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/deployments/{id}/log", "2018-02-01"),
}, } ,
}, } ,
},
{ 
	Display: "domainOwnershipIdentifiers",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/domainOwnershipIdentifiers", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{domainOwnershipIdentifierName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/domainOwnershipIdentifiers/{domainOwnershipIdentifierName}", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/domainOwnershipIdentifiers/{domainOwnershipIdentifierName}", "2018-02-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/domainOwnershipIdentifiers/{domainOwnershipIdentifierName}", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/domainOwnershipIdentifiers/{domainOwnershipIdentifierName}", "2018-02-01"),
}, } ,
},
{ 
	Display: "MSDeploy",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/extensions/MSDeploy", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/extensions/MSDeploy", "2018-02-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "log",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/extensions/MSDeploy/log", "2018-02-01"),
}, } ,
},
{ 
	Display: "functions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/functions", "2018-02-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "token",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/functions/admin/token", "2018-02-01"),
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{functionName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/functions/{functionName}", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/functions/{functionName}", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/functions/{functionName}", "2018-02-01"),
	Children: []swagger.ResourceType{  } ,
	SubResources: []swagger.ResourceType{  } ,
}, } ,
},
{ 
	Display: "hostNameBindings",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hostNameBindings", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{hostName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hostNameBindings/{hostName}", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hostNameBindings/{hostName}", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hostNameBindings/{hostName}", "2018-02-01"),
}, } ,
},
{ 
	Display: "hybridConnectionRelays",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hybridConnectionRelays", "2018-02-01"),
},
{ 
	Display: "hybridconnection",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hybridconnection", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{entityName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hybridconnection/{entityName}", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hybridconnection/{entityName}", "2018-02-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hybridconnection/{entityName}", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hybridconnection/{entityName}", "2018-02-01"),
}, } ,
},
{ 
	Display: "instances",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "MSDeploy",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/extensions/MSDeploy", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/extensions/MSDeploy", "2018-02-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "log",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/extensions/MSDeploy/log", "2018-02-01"),
}, } ,
},
{ 
	Display: "processes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{processId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes/{processId}", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes/{processId}", "2018-02-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "dump",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes/{processId}/dump", "2018-02-01"),
},
{ 
	Display: "modules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes/{processId}/modules", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{baseAddress}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes/{processId}/modules/{baseAddress}", "2018-02-01"),
}, } ,
},
{ 
	Display: "threads",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes/{processId}/threads", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{threadId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes/{processId}/threads/{threadId}", "2018-02-01"),
}, } ,
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "metricdefinitions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/metricdefinitions", "2018-02-01"),
},
{ 
	Display: "metrics",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/metrics", "2018-02-01"),
},
{ 
	Display: "status",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/migratemysql/status", "2018-02-01"),
},
{ 
	Display: "virtualNetwork",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/networkConfig/virtualNetwork", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/networkConfig/virtualNetwork", "2018-02-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/networkConfig/virtualNetwork", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/networkConfig/virtualNetwork", "2018-02-01"),
},
{ 
	Display: "perfcounters",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/perfcounters", "2018-02-01"),
},
{ 
	Display: "phplogging",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/phplogging", "2018-02-01"),
},
{ 
	Display: "premieraddons",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/premieraddons", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{premierAddOnName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/premieraddons/{premierAddOnName}", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/premieraddons/{premierAddOnName}", "2018-02-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/premieraddons/{premierAddOnName}", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/premieraddons/{premierAddOnName}", "2018-02-01"),
}, } ,
},
{ 
	Display: "virtualNetworks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/privateAccess/virtualNetworks", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/privateAccess/virtualNetworks", "2018-02-01"),
},
{ 
	Display: "processes",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/processes", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{processId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/processes/{processId}", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/processes/{processId}", "2018-02-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "dump",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/processes/{processId}/dump", "2018-02-01"),
},
{ 
	Display: "modules",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/processes/{processId}/modules", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{baseAddress}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/processes/{processId}/modules/{baseAddress}", "2018-02-01"),
}, } ,
},
{ 
	Display: "threads",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/processes/{processId}/threads", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{threadId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/processes/{processId}/threads/{threadId}", "2018-02-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "publicCertificates",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/publicCertificates", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{publicCertificateName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/publicCertificates/{publicCertificateName}", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/publicCertificates/{publicCertificateName}", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/publicCertificates/{publicCertificateName}", "2018-02-01"),
}, } ,
},
{ 
	Display: "siteextensions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/siteextensions", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{siteExtensionId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/siteextensions/{siteExtensionId}", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/siteextensions/{siteExtensionId}", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/siteextensions/{siteExtensionId}", "2018-02-01"),
}, } ,
},
{ 
	Display: "snapshots",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/snapshots", "2018-02-01"),
},
{ 
	Display: "snapshotsdr",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/snapshotsdr", "2018-02-01"),
},
{ 
	Display: "web",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/sourcecontrols/web", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/sourcecontrols/web", "2018-02-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/sourcecontrols/web", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/sourcecontrols/web", "2018-02-01"),
},
{ 
	Display: "triggeredwebjobs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/triggeredwebjobs", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{webJobName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/triggeredwebjobs/{webJobName}", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/triggeredwebjobs/{webJobName}", "2018-02-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "history",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/triggeredwebjobs/{webJobName}/history", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{id}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/triggeredwebjobs/{webJobName}/history/{id}", "2018-02-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "usages",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/usages", "2018-02-01"),
},
{ 
	Display: "virtualNetworkConnections",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/virtualNetworkConnections", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{vnetName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/virtualNetworkConnections/{vnetName}", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/virtualNetworkConnections/{vnetName}", "2018-02-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/virtualNetworkConnections/{vnetName}", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/virtualNetworkConnections/{vnetName}", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{gatewayName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/virtualNetworkConnections/{vnetName}/gateways/{gatewayName}", "2018-02-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/virtualNetworkConnections/{vnetName}/gateways/{gatewayName}", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/virtualNetworkConnections/{vnetName}/gateways/{gatewayName}", "2018-02-01"),
}, } ,
}, } ,
},
{ 
	Display: "webjobs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/webjobs", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{webJobName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/webjobs/{webJobName}", "2018-02-01"),
}, } ,
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{relayName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}", "2018-02-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}", "2018-02-01"),
	Children: []swagger.ResourceType{  } ,
},
{ 
	Display: "{view}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/networkFeatures/{view}", "2018-02-01"),
},
{ 
	Display: "{operationId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/networkTrace/operationresults/{operationId}", "2018-02-01"),
},
{ 
	Display: "{operationId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/networkTrace/{operationId}", "2018-02-01"),
},
{ 
	Display: "{operationId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/networkTraces/current/operationresults/{operationId}", "2018-02-01"),
},
{ 
	Display: "{operationId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/networkTraces/{operationId}", "2018-02-01"),
}, } ,
}, } ,
},
{ 
	Display: "snapshots",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/snapshots", "2018-02-01"),
},
{ 
	Display: "snapshotsdr",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/snapshotsdr", "2018-02-01"),
},
{ 
	Display: "web",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/sourcecontrols/web", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/sourcecontrols/web", "2018-02-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/sourcecontrols/web", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/sourcecontrols/web", "2018-02-01"),
},
{ 
	Display: "triggeredwebjobs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/triggeredwebjobs", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{webJobName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/triggeredwebjobs/{webJobName}", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/triggeredwebjobs/{webJobName}", "2018-02-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "history",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/triggeredwebjobs/{webJobName}/history", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{id}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/triggeredwebjobs/{webJobName}/history/{id}", "2018-02-01"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "usages",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/usages", "2018-02-01"),
},
{ 
	Display: "virtualNetworkConnections",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/virtualNetworkConnections", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{vnetName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/virtualNetworkConnections/{vnetName}", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/virtualNetworkConnections/{vnetName}", "2018-02-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/virtualNetworkConnections/{vnetName}", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/virtualNetworkConnections/{vnetName}", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{gatewayName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/virtualNetworkConnections/{vnetName}/gateways/{gatewayName}", "2018-02-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/virtualNetworkConnections/{vnetName}/gateways/{gatewayName}", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/virtualNetworkConnections/{vnetName}/gateways/{gatewayName}", "2018-02-01"),
}, } ,
}, } ,
},
{ 
	Display: "webjobs",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/webjobs", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{webJobName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/webjobs/{webJobName}", "2018-02-01"),
}, } ,
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{relayName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}", "2018-02-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}", "2018-02-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}", "2018-02-01"),
	Children: []swagger.ResourceType{  } ,
},
{ 
	Display: "{view}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/networkFeatures/{view}", "2018-02-01"),
},
{ 
	Display: "{operationId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/networkTrace/operationresults/{operationId}", "2018-02-01"),
},
{ 
	Display: "{operationId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/networkTrace/{operationId}", "2018-02-01"),
},
{ 
	Display: "{operationId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/networkTraces/current/operationresults/{operationId}", "2018-02-01"),
},
{ 
	Display: "{operationId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/networkTraces/{operationId}", "2018-02-01"),
}, } ,
},
{ 
	Display: "recommendationHistory",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/recommendationHistory", "2018-02-01"),
},
{ 
	Display: "recommendations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/recommendations", "2018-02-01"),
	Children: []swagger.ResourceType{  } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{name}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/recommendations/{name}", "2018-02-01"),
	Children: []swagger.ResourceType{  } ,
}, } ,
}, } ,
},
{ 
	Display: "multipleActivationKeys",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.WindowsESU/multipleActivationKeys", "2019-09-16-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{multipleActivationKeyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.WindowsESU/multipleActivationKeys/{multipleActivationKeyName}", "2019-09-16-preview"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.WindowsESU/multipleActivationKeys/{multipleActivationKeyName}", "2019-09-16-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.WindowsESU/multipleActivationKeys/{multipleActivationKeyName}", "2019-09-16-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.WindowsESU/multipleActivationKeys/{multipleActivationKeyName}", "2019-09-16-preview"),
}, } ,
},
{ 
	Display: "deviceServices",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.WindowsIoT/deviceServices", "2019-06-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{deviceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.WindowsIoT/deviceServices/{deviceName}", "2019-06-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.WindowsIoT/deviceServices/{deviceName}", "2019-06-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.WindowsIoT/deviceServices/{deviceName}", "2019-06-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.WindowsIoT/deviceServices/{deviceName}", "2019-06-01"),
}, } ,
},
{ 
	Display: "components",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceNamespace}/{resourceType}/{resourceName}/providers/Microsoft.WorkloadMonitor/components", "2018-08-31-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{componentId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceNamespace}/{resourceType}/{resourceName}/providers/Microsoft.WorkloadMonitor/components/{componentId}", "2018-08-31-preview"),
}, } ,
},
{ 
	Display: "monitorInstances",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceNamespace}/{resourceType}/{resourceName}/providers/Microsoft.WorkloadMonitor/monitorInstances", "2018-08-31-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{monitorInstanceId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceNamespace}/{resourceType}/{resourceName}/providers/Microsoft.WorkloadMonitor/monitorInstances/{monitorInstanceId}", "2018-08-31-preview"),
}, } ,
},
{ 
	Display: "monitors",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceNamespace}/{resourceType}/{resourceName}/providers/Microsoft.WorkloadMonitor/monitors", "2018-08-31-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{monitorId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceNamespace}/{resourceType}/{resourceName}/providers/Microsoft.WorkloadMonitor/monitors/{monitorId}", "2018-08-31-preview"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceNamespace}/{resourceType}/{resourceName}/providers/Microsoft.WorkloadMonitor/monitors/{monitorId}", "2018-08-31-preview"),
}, } ,
},
{ 
	Display: "notificationSettings",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceNamespace}/{resourceType}/{resourceName}/providers/Microsoft.WorkloadMonitor/notificationSettings", "2018-08-31-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{notificationSettingName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceNamespace}/{resourceType}/{resourceName}/providers/Microsoft.WorkloadMonitor/notificationSettings/{notificationSettingName}", "2018-08-31-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceNamespace}/{resourceType}/{resourceName}/providers/Microsoft.WorkloadMonitor/notificationSettings/{notificationSettingName}", "2018-08-31-preview"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "tenants",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/tenants", "2016-06-01"),
},
{ 
	Display: "{policyAssignmentId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{policyAssignmentId}", "2016-12-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{policyAssignmentId}", "2016-12-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/{policyAssignmentId}", "2016-12-01"),
},
{ 
	Display: "{resourceId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceId}", "2018-05-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceId}", "2018-05-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceId}", "2018-05-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceId}", "2018-05-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "deviceSecurityGroups",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceId}/providers/Microsoft.Security/deviceSecurityGroups", "2019-08-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{deviceSecurityGroupName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceId}/providers/Microsoft.Security/deviceSecurityGroups/{deviceSecurityGroupName}", "2019-08-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceId}/providers/Microsoft.Security/deviceSecurityGroups/{deviceSecurityGroupName}", "2019-08-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceId}/providers/Microsoft.Security/deviceSecurityGroups/{deviceSecurityGroupName}", "2019-08-01"),
}, } ,
}, } ,
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{settingName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceId}/providers/Microsoft.Security/advancedThreatProtectionSettings/{settingName}", "2019-01-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceId}/providers/Microsoft.Security/advancedThreatProtectionSettings/{settingName}", "2019-01-01"),
},
{ 
	Display: "{assessmentName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceId}/providers/Microsoft.Security/assessments/{assessmentName}", "2020-01-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceId}/providers/Microsoft.Security/assessments/{assessmentName}", "2020-01-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceId}/providers/Microsoft.Security/assessments/{assessmentName}", "2020-01-01"),
},
{ 
	Display: "{complianceResultName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceId}/providers/Microsoft.Security/complianceResults/{complianceResultName}", "2017-08-01"),
}, } ,
},
{ 
	Display: "locks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Authorization/locks", "2016-09-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{lockName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Authorization/locks/{lockName}", "2016-09-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Authorization/locks/{lockName}", "2016-09-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Authorization/locks/{lockName}", "2016-09-01"),
}, } ,
},
{ 
	Display: "{policyAssignmentName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Authorization/policyAssignments/{policyAssignmentName}", "2016-12-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Authorization/policyAssignments/{policyAssignmentName}", "2016-12-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Authorization/policyAssignments/{policyAssignmentName}", "2016-12-01"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Search/operations", "2015-08-19"),
},
{ 
	Display: "assessmentMetadata",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Security/assessmentMetadata", "2020-01-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{assessmentMetadataName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Security/assessmentMetadata/{assessmentMetadataName}", "2020-01-01"),
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Security/operations", "2015-06-01-preview"),
},
{ 
	Display: "assessments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Security/assessments", "2020-01-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "subAssessments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Security/assessments/{assessmentName}/subAssessments", "2019-01-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{subAssessmentName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Security/assessments/{assessmentName}/subAssessments/{subAssessmentName}", "2019-01-01-preview"),
}, } ,
}, } ,
},
{ 
	Display: "complianceResults",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Security/complianceResults", "2017-08-01"),
},
{ 
	Display: "compliances",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Security/compliances", "2017-08-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{complianceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Security/compliances/{complianceName}", "2017-08-01-preview"),
}, } ,
},
{ 
	Display: "informationProtectionPolicies",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Security/informationProtectionPolicies", "2017-08-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{informationProtectionPolicyName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Security/informationProtectionPolicies/{informationProtectionPolicyName}", "2017-08-01-preview"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Security/informationProtectionPolicies/{informationProtectionPolicyName}", "2017-08-01-preview"),
}, } ,
},
{ 
	Display: "subAssessments",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Security/subAssessments", "2019-01-01-preview"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.SecurityInsights/operations", "2020-01-01"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.SerialConsole/operations", "2018-05-01"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ServiceBus/operations", "2017-04-01"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ServiceFabric/operations", "2019-03-01"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ServiceFabricMesh/operations", "2018-09-01-preview"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.SignalRService/operations", "2018-10-01"),
},
{ 
	Display: "hybridUseBenefits",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.SoftwarePlan/hybridUseBenefits", "2019-12-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{planId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.SoftwarePlan/hybridUseBenefits/{planId}", "2019-12-01"),
	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.SoftwarePlan/hybridUseBenefits/{planId}", "2019-12-01"),
	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.SoftwarePlan/hybridUseBenefits/{planId}", "2019-12-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.SoftwarePlan/hybridUseBenefits/{planId}", "2019-12-01"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "revisions",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.SoftwarePlan/hybridUseBenefits/{planId}/revisions", "2019-12-01"),
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.SoftwarePlan/operations", "2019-12-01"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Sql/operations", "2014-04-01"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.SqlVirtualMachine/operations", "2017-03-01-preview"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.StorSimple/operations", "2017-06-01"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Storage/operations", "2019-06-01"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.StorageCache/operations", "2019-11-01"),
},
{ 
	Display: "locations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ImportExport/locations", "2016-11-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{locationName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ImportExport/locations/{locationName}", "2016-11-01"),
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ImportExport/operations", "2016-11-01"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.StorageSync/operations", "2019-06-01"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.StreamAnalytics/operations", "2016-03-01"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Subscription/operations", "2020-01-01"),
},
{ 
	Display: "{operationId}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Subscription/subscriptionOperations/{operationId}", "2020-01-01"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Support/operations", "2019-05-01-preview"),
},
{ 
	Display: "services",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Support/services", "2019-05-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{serviceName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Support/services/{serviceName}", "2019-05-01-preview"),
	Children: []swagger.ResourceType{ 
{ 
	Display: "problemClassifications",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Support/services/{serviceName}/problemClassifications", "2019-05-01-preview"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{problemClassificationName}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Support/services/{serviceName}/problemClassifications/{problemClassificationName}", "2019-05-01-preview"),
}, } ,
}, } ,
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Synapse/operations", "2019-06-01-preview"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.TimeSeriesInsights/operations", "2017-11-15"),
},
{ 
	Display: "default",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Network/trafficManagerGeographicHierarchies/default", "2018-04-01"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/microsoft.visualstudio/operations", "2014-04-01-preview"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.VMwareCloudSimple/operations", "2019-04-01"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.VMwareVirtustream/operations", "2019-08-09-preview"),
},
{ 
	Display: "availableStacks",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Web/availableStacks", "2018-02-01"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Web/operations", "2018-02-01"),
},
{ 
	Display: "web",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Web/publishingUsers/web", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Web/publishingUsers/web", "2018-02-01"),
},
{ 
	Display: "sourcecontrols",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Web/sourcecontrols", "2018-02-01"),
	SubResources: []swagger.ResourceType{ 
{ 
	Display: "{sourceControlType}",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Web/sourcecontrols/{sourceControlType}", "2018-02-01"),
	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Web/sourcecontrols/{sourceControlType}", "2018-02-01"),
}, } ,
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.WindowsESU/operations", "2019-09-16-preview"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.WindowsIoT/operations", "2019-06-01"),
},
{ 
	Display: "operations",
	Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.WorkloadMonitor/operations", "2018-08-31-preview"),
}, } 

}
