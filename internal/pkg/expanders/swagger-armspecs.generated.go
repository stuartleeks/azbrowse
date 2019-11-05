package expanders

import (
	"github.com/lawrencegripper/azbrowse/pkg/endpoints"
	"github.com/lawrencegripper/azbrowse/pkg/swagger"
)

func (e *SwaggerAPISetARMResources) loadResourceTypes() []swagger.ResourceType {
	return []swagger.ResourceType{
		{
			Display:  "addsservices",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices", "2014-01-01"),
			Children: []swagger.ResourceType{
				{
					Display:  "premiumCheck",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/premiumCheck", "2014-01-01"),
				}},
			SubResources: []swagger.ResourceType{
				{
					Display:        "{serviceName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}", "2014-01-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}", "2014-01-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}", "2014-01-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "addomainservicemembers",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/addomainservicemembers", "2014-01-01"),
						},
						{
							Display:  "addsservicemembers",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/addsservicemembers", "2014-01-01"),
						},
						{
							Display:  "alerts",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/alerts", "2014-01-01"),
						},
						{
							Display:  "configuration",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/configuration", "2014-01-01"),
						},
						{
							Display:  "forestsummary",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/forestsummary", "2014-01-01"),
						},
						{
							Display:  "metricmetadata",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/metricmetadata", "2014-01-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:  "{metricName}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/metricmetadata/{metricName}", "2014-01-01"),
									SubResources: []swagger.ResourceType{
										{
											Display:  "{groupName}",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/metricmetadata/{metricName}/groups/{groupName}", "2014-01-01"),
										}},
								}},
						},
						{
							Display:  "replicationdetails",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/replicationdetails", "2014-01-01"),
						},
						{
							Display:  "replicationstatus",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/replicationstatus", "2014-01-01"),
						},
						{
							Display:  "replicationsummary",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/replicationsummary", "2014-01-01"),
						},
						{
							Display:  "servicemembers",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/servicemembers", "2014-01-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{serviceMemberId}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/servicemembers/{serviceMemberId}", "2014-01-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/servicemembers/{serviceMemberId}", "2014-01-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "alerts",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/servicemembers/{serviceMemberId}/alerts", "2014-01-01"),
										},
										{
											Display:  "credentials",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/servicemembers/{serviceMemberId}/credentials", "2014-01-01"),
										}},
								}},
						}},
					SubResources: []swagger.ResourceType{
						{
							Display:  "{dimension}",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/dimensions/{dimension}", "2014-01-01"),
						},
						{
							Display:        "userpreference",
							Endpoint:       endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/features/{featureName}/userpreference", "2014-01-01"),
							DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/features/{featureName}/userpreference", "2014-01-01"),
						},
						{
							Display:  "{groupName}",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/metrics/{metricName}/groups/{groupName}", "2014-01-01"),
							Children: []swagger.ResourceType{
								{
									Display:  "average",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/metrics/{metricName}/groups/{groupName}/average", "2014-01-01"),
								},
								{
									Display:  "sum",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/addsservices/{serviceName}/metrics/{metricName}/groups/{groupName}/sum", "2014-01-01"),
								}},
						}},
				}},
		},
		{
			Display:       "configuration",
			Endpoint:      endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/configuration", "2014-01-01"),
			PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/configuration", "2014-01-01"),
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/operations", "2014-01-01"),
		},
		{
			Display:  "IsDevOps",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/reports/DevOps/IsDevOps", "2014-01-01"),
		},
		{
			Display:  "connectors",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/service/{serviceName}/servicemembers/{serviceMemberId}/connectors", "2014-01-01"),
		},
		{
			Display:  "services",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services", "2014-01-01"),
			Children: []swagger.ResourceType{
				{
					Display:  "premiumCheck",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/premiumCheck", "2014-01-01"),
				}},
			SubResources: []swagger.ResourceType{
				{
					Display:        "{serviceName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}", "2014-01-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}", "2014-01-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}", "2014-01-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "alerts",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/alerts", "2014-01-01"),
						},
						{
							Display:  "counts",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/exporterrors/counts", "2014-01-01"),
						},
						{
							Display:  "listV2",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/exporterrors/listV2", "2014-01-01"),
						},
						{
							Display:  "exportstatus",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/exportstatus", "2014-01-01"),
						},
						{
							Display:  "metricmetadata",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/metricmetadata", "2014-01-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:  "{metricName}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/metricmetadata/{metricName}", "2014-01-01"),
									SubResources: []swagger.ResourceType{
										{
											Display:  "{groupName}",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/metricmetadata/{metricName}/groups/{groupName}", "2014-01-01"),
										}},
								}},
						},
						{
							Display:  "monitoringconfigurations",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/monitoringconfigurations", "2014-01-01"),
						},
						{
							Display:  "user",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/reports/badpassword/details/user", "2014-01-01"),
						},
						{
							Display:  "blobUris",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/reports/riskyIp/blobUris", "2014-01-01"),
						},
						{
							Display:  "servicemembers",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/servicemembers", "2014-01-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{serviceMemberId}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/servicemembers/{serviceMemberId}", "2014-01-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/servicemembers/{serviceMemberId}", "2014-01-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "alerts",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/servicemembers/{serviceMemberId}/alerts", "2014-01-01"),
										},
										{
											Display:  "credentials",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/servicemembers/{serviceMemberId}/credentials", "2014-01-01"),
										},
										{
											Display:  "datafreshness",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/servicemembers/{serviceMemberId}/datafreshness", "2014-01-01"),
										},
										{
											Display:  "exportstatus",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/servicemembers/{serviceMemberId}/exportstatus", "2014-01-01"),
										},
										{
											Display:  "globalconfiguration",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/servicemembers/{serviceMemberId}/globalconfiguration", "2014-01-01"),
										},
										{
											Display:  "serviceconfiguration",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/servicemembers/{serviceMemberId}/serviceconfiguration", "2014-01-01"),
										}},
									SubResources: []swagger.ResourceType{
										{
											Display:  "{metricName}",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/servicemembers/{serviceMemberId}/metrics/{metricName}", "2014-01-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:  "{groupName}",
													Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/servicemembers/{serviceMemberId}/metrics/{metricName}/groups/{groupName}", "2014-01-01"),
												}},
										}},
								}},
						}},
					SubResources: []swagger.ResourceType{
						{
							Display:  "{featureName}",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/TenantWhitelisting/{featureName}", "2014-01-01"),
						},
						{
							Display:  "{featureName}",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/checkServiceFeatureAvailibility/{featureName}", "2014-01-01"),
						},
						{
							Display:  "alertfeedback",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/feedbacktype/alerts/{shortName}/alertfeedback", "2014-01-01"),
						},
						{
							Display:  "{groupName}",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/metrics/{metricName}/groups/{groupName}", "2014-01-01"),
							Children: []swagger.ResourceType{
								{
									Display:  "average",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/metrics/{metricName}/groups/{groupName}/average", "2014-01-01"),
								},
								{
									Display:  "sum",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/metrics/{metricName}/groups/{groupName}/sum", "2014-01-01"),
								}},
						}},
				}},
		},
		{
			Display:  "metadata",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Advisor/metadata", "2017-04-19"),
			SubResources: []swagger.ResourceType{
				{
					Display:  "{name}",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Advisor/metadata/{name}", "2017-04-19"),
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Advisor/operations", "2017-04-19"),
		},
		{
			Display:     "configurations",
			Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/configurations", "2017-04-19"),
			PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/configurations", "2017-04-19"),
		},
		{
			Display:  "recommendations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/recommendations", "2017-04-19"),
		},
		{
			Display:  "suppressions",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/suppressions", "2017-04-19"),
		},
		{
			Display:     "configurations",
			Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Advisor/configurations", "2017-04-19"),
			PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Advisor/configurations", "2017-04-19"),
		},
		{
			Display:  "{recommendationId}",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceUri}/providers/Microsoft.Advisor/recommendations/{recommendationId}", "2017-04-19"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{name}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/{resourceUri}/providers/Microsoft.Advisor/recommendations/{recommendationId}/suppressions/{name}", "2017-04-19"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceUri}/providers/Microsoft.Advisor/recommendations/{recommendationId}/suppressions/{name}", "2017-04-19"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/{resourceUri}/providers/Microsoft.Advisor/recommendations/{recommendationId}/suppressions/{name}", "2017-04-19"),
				}},
		},
		{
			Display:  "smartDetectorAlertRules",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/microsoft.alertsManagement/smartDetectorAlertRules", "2019-06-01"),
		},
		{
			Display:  "smartDetectorAlertRules",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.alertsManagement/smartDetectorAlertRules", "2019-06-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{alertRuleName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.alertsManagement/smartDetectorAlertRules/{alertRuleName}", "2019-06-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.alertsManagement/smartDetectorAlertRules/{alertRuleName}", "2019-06-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.alertsManagement/smartDetectorAlertRules/{alertRuleName}", "2019-06-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.alertsManagement/smartDetectorAlertRules/{alertRuleName}", "2019-06-01"),
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.AnalysisServices/operations", "2017-08-01"),
		},
		{
			Display:  "{operationId}",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.AnalysisServices/locations/{location}/operationresults/{operationId}", "2017-08-01"),
		},
		{
			Display:  "{operationId}",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.AnalysisServices/locations/{location}/operationstatuses/{operationId}", "2017-08-01"),
		},
		{
			Display:  "servers",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.AnalysisServices/servers", "2017-08-01"),
		},
		{
			Display:  "skus",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.AnalysisServices/skus", "2017-08-01"),
		},
		{
			Display:  "servers",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AnalysisServices/servers", "2017-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{serverName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AnalysisServices/servers/{serverName}", "2017-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AnalysisServices/servers/{serverName}", "2017-08-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AnalysisServices/servers/{serverName}", "2017-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AnalysisServices/servers/{serverName}", "2017-08-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "skus",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AnalysisServices/servers/{serverName}/skus", "2017-08-01"),
						}},
				}},
		},
		{
			Display:  "apis",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis", "2019-01-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{apiId}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}", "2019-01-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}", "2019-01-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}", "2019-01-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}", "2019-01-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "diagnostics",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/diagnostics", "2019-01-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{diagnosticId}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/diagnostics/{diagnosticId}", "2019-01-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/diagnostics/{diagnosticId}", "2019-01-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/diagnostics/{diagnosticId}", "2019-01-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/diagnostics/{diagnosticId}", "2019-01-01"),
								}},
						},
						{
							Display:  "issues",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues", "2019-01-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{issueId}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}", "2019-01-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}", "2019-01-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}", "2019-01-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}", "2019-01-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "attachments",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}/attachments", "2019-01-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{attachmentId}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}/attachments/{attachmentId}", "2019-01-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}/attachments/{attachmentId}", "2019-01-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}/attachments/{attachmentId}", "2019-01-01"),
												}},
										},
										{
											Display:  "comments",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}/comments", "2019-01-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{commentId}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}/comments/{commentId}", "2019-01-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}/comments/{commentId}", "2019-01-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}/comments/{commentId}", "2019-01-01"),
												}},
										}},
								}},
						},
						{
							Display:  "operations",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations", "2019-01-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{operationId}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}", "2019-01-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}", "2019-01-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}", "2019-01-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}", "2019-01-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "policies",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}/policies", "2019-01-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{policyId}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}/policies/{policyId}", "2019-01-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}/policies/{policyId}", "2019-01-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}/policies/{policyId}", "2019-01-01"),
												}},
										},
										{
											Display:  "tags",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}/tags", "2019-01-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{tagId}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}/tags/{tagId}", "2019-01-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}/tags/{tagId}", "2019-01-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operations/{operationId}/tags/{tagId}", "2019-01-01"),
												}},
										}},
								}},
						},
						{
							Display:  "operationsByTags",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operationsByTags", "2019-01-01"),
						},
						{
							Display:  "policies",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/policies", "2019-01-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{policyId}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/policies/{policyId}", "2019-01-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/policies/{policyId}", "2019-01-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/policies/{policyId}", "2019-01-01"),
								}},
						},
						{
							Display:  "products",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/products", "2019-01-01"),
						},
						{
							Display:  "releases",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/releases", "2019-01-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{releaseId}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/releases/{releaseId}", "2019-01-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/releases/{releaseId}", "2019-01-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/releases/{releaseId}", "2019-01-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/releases/{releaseId}", "2019-01-01"),
								}},
						},
						{
							Display:  "revisions",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/revisions", "2019-01-01"),
						},
						{
							Display:  "schemas",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/schemas", "2019-01-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{schemaId}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/schemas/{schemaId}", "2019-01-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/schemas/{schemaId}", "2019-01-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/schemas/{schemaId}", "2019-01-01"),
								}},
						},
						{
							Display:  "tagDescriptions",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/tagDescriptions", "2019-01-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{tagId}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/tagDescriptions/{tagId}", "2019-01-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/tagDescriptions/{tagId}", "2019-01-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/tagDescriptions/{tagId}", "2019-01-01"),
								}},
						},
						{
							Display:  "tags",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/tags", "2019-01-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{tagId}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/tags/{tagId}", "2019-01-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/tags/{tagId}", "2019-01-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/tags/{tagId}", "2019-01-01"),
								}},
						}},
				}},
		},
		{
			Display:  "apisByTags",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apisByTags", "2019-01-01"),
		},
		{
			Display:  "apiVersionSets",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apiVersionSets", "2019-01-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{versionSetId}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apiVersionSets/{versionSetId}", "2019-01-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apiVersionSets/{versionSetId}", "2019-01-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apiVersionSets/{versionSetId}", "2019-01-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apiVersionSets/{versionSetId}", "2019-01-01"),
				}},
		},
		{
			Display:  "authorizationServers",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationServers", "2019-01-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{authsid}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationServers/{authsid}", "2019-01-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationServers/{authsid}", "2019-01-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationServers/{authsid}", "2019-01-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationServers/{authsid}", "2019-01-01"),
				}},
		},
		{
			Display:  "backends",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/backends", "2019-01-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{backendId}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/backends/{backendId}", "2019-01-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/backends/{backendId}", "2019-01-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/backends/{backendId}", "2019-01-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/backends/{backendId}", "2019-01-01"),
					Children:       []swagger.ResourceType{},
				}},
		},
		{
			Display:  "caches",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/caches", "2019-01-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{cacheId}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/caches/{cacheId}", "2019-01-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/caches/{cacheId}", "2019-01-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/caches/{cacheId}", "2019-01-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/caches/{cacheId}", "2019-01-01"),
				}},
		},
		{
			Display:  "certificates",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/certificates", "2019-01-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{certificateId}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/certificates/{certificateId}", "2019-01-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/certificates/{certificateId}", "2019-01-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/certificates/{certificateId}", "2019-01-01"),
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ApiManagement/operations", "2019-01-01"),
		},
		{
			Display:  "service",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ApiManagement/service", "2019-01-01"),
		},
		{
			Display:  "service",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service", "2019-01-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{serviceName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}", "2019-01-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}", "2019-01-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}", "2019-01-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}", "2019-01-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "skus",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/skus", "2019-01-01"),
						},
						{
							Display:  "diagnostics",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/diagnostics", "2019-01-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{diagnosticId}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/diagnostics/{diagnosticId}", "2019-01-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/diagnostics/{diagnosticId}", "2019-01-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/diagnostics/{diagnosticId}", "2019-01-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/diagnostics/{diagnosticId}", "2019-01-01"),
								}},
						},
						{
							Display:  "templates",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/templates", "2019-01-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{templateName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/templates/{templateName}", "2019-01-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/templates/{templateName}", "2019-01-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/templates/{templateName}", "2019-01-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/templates/{templateName}", "2019-01-01"),
								}},
						},
						{
							Display:  "groups",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/groups", "2019-01-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{groupId}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/groups/{groupId}", "2019-01-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/groups/{groupId}", "2019-01-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/groups/{groupId}", "2019-01-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/groups/{groupId}", "2019-01-01"),
									Children: []swagger.ResourceType{
										{
											Display:      "users",
											Endpoint:     endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/groups/{groupId}/users", "2019-01-01"),
											SubResources: []swagger.ResourceType{},
										}},
								}},
						},
						{
							Display:  "identityProviders",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/identityProviders", "2019-01-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{identityProviderName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/identityProviders/{identityProviderName}", "2019-01-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/identityProviders/{identityProviderName}", "2019-01-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/identityProviders/{identityProviderName}", "2019-01-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/identityProviders/{identityProviderName}", "2019-01-01"),
								}},
						},
						{
							Display:  "issues",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/issues", "2019-01-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:  "{issueId}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/issues/{issueId}", "2019-01-01"),
								}},
						},
						{
							Display:  "loggers",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/loggers", "2019-01-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{loggerId}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/loggers/{loggerId}", "2019-01-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/loggers/{loggerId}", "2019-01-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/loggers/{loggerId}", "2019-01-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/loggers/{loggerId}", "2019-01-01"),
								}},
						},
						{
							Display:  "networkstatus",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/networkstatus", "2019-01-01"),
						},
						{
							Display:  "notifications",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications", "2019-01-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:     "{notificationName}",
									Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications/{notificationName}", "2019-01-01"),
									PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications/{notificationName}", "2019-01-01"),
									Children: []swagger.ResourceType{
										{
											Display:      "recipientEmails",
											Endpoint:     endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications/{notificationName}/recipientEmails", "2019-01-01"),
											SubResources: []swagger.ResourceType{},
										},
										{
											Display:      "recipientUsers",
											Endpoint:     endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/notifications/{notificationName}/recipientUsers", "2019-01-01"),
											SubResources: []swagger.ResourceType{},
										}},
								}},
						},
						{
							Display:  "openidConnectProviders",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/openidConnectProviders", "2019-01-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{opid}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/openidConnectProviders/{opid}", "2019-01-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/openidConnectProviders/{opid}", "2019-01-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/openidConnectProviders/{opid}", "2019-01-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/openidConnectProviders/{opid}", "2019-01-01"),
								}},
						},
						{
							Display:  "policies",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/policies", "2019-01-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{policyId}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/policies/{policyId}", "2019-01-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/policies/{policyId}", "2019-01-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/policies/{policyId}", "2019-01-01"),
								}},
						},
						{
							Display:  "policySnippets",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/policySnippets", "2019-01-01"),
						},
						{
							Display:       "delegation",
							Endpoint:      endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/delegation", "2019-01-01"),
							PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/delegation", "2019-01-01"),
							PutEndpoint:   endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/delegation", "2019-01-01"),
						},
						{
							Display:       "signin",
							Endpoint:      endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/signin", "2019-01-01"),
							PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/signin", "2019-01-01"),
							PutEndpoint:   endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/signin", "2019-01-01"),
						},
						{
							Display:       "signup",
							Endpoint:      endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/signup", "2019-01-01"),
							PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/signup", "2019-01-01"),
							PutEndpoint:   endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/signup", "2019-01-01"),
						},
						{
							Display:  "products",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products", "2019-01-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{productId}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}", "2019-01-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}", "2019-01-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}", "2019-01-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}", "2019-01-01"),
									Children: []swagger.ResourceType{
										{
											Display:      "apis",
											Endpoint:     endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/apis", "2019-01-01"),
											SubResources: []swagger.ResourceType{},
										},
										{
											Display:      "groups",
											Endpoint:     endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/groups", "2019-01-01"),
											SubResources: []swagger.ResourceType{},
										},
										{
											Display:  "policies",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/policies", "2019-01-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{policyId}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/policies/{policyId}", "2019-01-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/policies/{policyId}", "2019-01-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/policies/{policyId}", "2019-01-01"),
												}},
										},
										{
											Display:  "subscriptions",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/subscriptions", "2019-01-01"),
										},
										{
											Display:  "tags",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/tags", "2019-01-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{tagId}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/tags/{tagId}", "2019-01-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/tags/{tagId}", "2019-01-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/tags/{tagId}", "2019-01-01"),
												}},
										}},
								}},
						},
						{
							Display:  "productsByTags",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/productsByTags", "2019-01-01"),
						},
						{
							Display:  "properties",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/properties", "2019-01-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{propId}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/properties/{propId}", "2019-01-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/properties/{propId}", "2019-01-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/properties/{propId}", "2019-01-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/properties/{propId}", "2019-01-01"),
								}},
						},
						{
							Display:  "regions",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/regions", "2019-01-01"),
						},
						{
							Display:  "byApi",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/byApi", "2019-01-01"),
						},
						{
							Display:  "byGeo",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/byGeo", "2019-01-01"),
						},
						{
							Display:  "byOperation",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/byOperation", "2019-01-01"),
						},
						{
							Display:  "byProduct",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/byProduct", "2019-01-01"),
						},
						{
							Display:  "byRequest",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/byRequest", "2019-01-01"),
						},
						{
							Display:  "bySubscription",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/bySubscription", "2019-01-01"),
						},
						{
							Display:  "byTime",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/byTime", "2019-01-01"),
						},
						{
							Display:  "byUser",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/byUser", "2019-01-01"),
						},
						{
							Display:  "subscriptions",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions", "2019-01-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{sid}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}", "2019-01-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}", "2019-01-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}", "2019-01-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}", "2019-01-01"),
									Children:       []swagger.ResourceType{},
								}},
						},
						{
							Display:  "tagResources",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tagResources", "2019-01-01"),
						},
						{
							Display:  "tags",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tags", "2019-01-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{tagId}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tags/{tagId}", "2019-01-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tags/{tagId}", "2019-01-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tags/{tagId}", "2019-01-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tags/{tagId}", "2019-01-01"),
								}},
						},
						{
							Display:  "users",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users", "2019-01-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{userId}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{userId}", "2019-01-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{userId}", "2019-01-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{userId}", "2019-01-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{userId}", "2019-01-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "groups",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{userId}/groups", "2019-01-01"),
										},
										{
											Display:  "identities",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{userId}/identities", "2019-01-01"),
										},
										{
											Display:  "subscriptions",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{userId}/subscriptions", "2019-01-01"),
										}},
								}},
						}},
					SubResources: []swagger.ResourceType{
						{
							Display:  "networkstatus",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/locations/{locationName}/networkstatus", "2019-01-01"),
						},
						{
							Display:       "{quotaCounterKey}",
							Endpoint:      endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/quotas/{quotaCounterKey}", "2019-01-01"),
							PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/quotas/{quotaCounterKey}", "2019-01-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:       "{quotaPeriodKey}",
									Endpoint:      endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/quotas/{quotaCounterKey}/periods/{quotaPeriodKey}", "2019-01-01"),
									PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/quotas/{quotaCounterKey}/periods/{quotaPeriodKey}", "2019-01-01"),
								}},
						},
						{
							Display:       "{accessName}",
							Endpoint:      endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{accessName}", "2019-01-01"),
							PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{accessName}", "2019-01-01"),
							Children: []swagger.ResourceType{
								{
									Display:  "git",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{accessName}/git", "2019-01-01"),
									Children: []swagger.ResourceType{},
								}},
						},
						{
							Display:  "syncState",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{configurationName}/syncState", "2019-01-01"),
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.AppConfiguration/operations", "2019-10-01"),
		},
		{
			Display:  "configurationStores",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.AppConfiguration/configurationStores", "2019-10-01"),
		},
		{
			Display:  "configurationStores",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores", "2019-10-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{configStoreName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}", "2019-10-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}", "2019-10-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}", "2019-10-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}", "2019-10-01"),
					Children:       []swagger.ResourceType{},
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Insights/operations", "2015-05-01"),
		},
		{
			Display:  "{scopePath}",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/components/{resourceName}/{scopePath}", "2015-05-01"),
			Children: []swagger.ResourceType{
				{
					Display:        "item",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/components/{resourceName}/{scopePath}/item", "2015-05-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/components/{resourceName}/{scopePath}/item", "2015-05-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/components/{resourceName}/{scopePath}/item", "2015-05-01"),
				}},
		},
		{
			Display:     "Annotations",
			Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/Annotations", "2015-05-01"),
			PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/Annotations", "2015-05-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{annotationId}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/Annotations/{annotationId}", "2015-05-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/Annotations/{annotationId}", "2015-05-01"),
				}},
		},
		{
			Display:        "{keyId}",
			Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/APIKeys/{keyId}", "2015-05-01"),
			DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/APIKeys/{keyId}", "2015-05-01"),
		},
		{
			Display:  "ApiKeys",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/ApiKeys", "2015-05-01"),
		},
		{
			Display:  "exportconfiguration",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/exportconfiguration", "2015-05-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{exportId}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/exportconfiguration/{exportId}", "2015-05-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/exportconfiguration/{exportId}", "2015-05-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/exportconfiguration/{exportId}", "2015-05-01"),
				}},
		},
		{
			Display:     "currentbillingfeatures",
			Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/currentbillingfeatures", "2015-05-01"),
			PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/currentbillingfeatures", "2015-05-01"),
		},
		{
			Display:  "featurecapabilities",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/featurecapabilities", "2015-05-01"),
		},
		{
			Display:  "getavailablebillingfeatures",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/getavailablebillingfeatures", "2015-05-01"),
		},
		{
			Display:  "quotastatus",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/quotastatus", "2015-05-01"),
		},
		{
			Display:  "ProactiveDetectionConfigs",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/ProactiveDetectionConfigs", "2015-05-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:     "{ConfigurationId}",
					Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/ProactiveDetectionConfigs/{ConfigurationId}", "2015-05-01"),
					PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/ProactiveDetectionConfigs/{ConfigurationId}", "2015-05-01"),
				}},
		},
		{
			Display:  "DefaultWorkItemConfig",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/DefaultWorkItemConfig", "2015-05-01"),
		},
		{
			Display:  "WorkItemConfigs",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/WorkItemConfigs", "2015-05-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{workItemConfigId}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/WorkItemConfigs/{workItemConfigId}", "2015-05-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/WorkItemConfigs/{workItemConfigId}", "2015-05-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/WorkItemConfigs/{workItemConfigId}", "2015-05-01"),
				}},
		},
		{
			Display:  "components",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Insights/components", "2015-05-01"),
		},
		{
			Display:  "components",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components", "2015-05-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{resourceName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}", "2015-05-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}", "2015-05-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}", "2015-05-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}", "2015-05-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "favorites",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/favorites", "2015-05-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{favoriteId}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/favorites/{favoriteId}", "2015-05-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/favorites/{favoriteId}", "2015-05-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/favorites/{favoriteId}", "2015-05-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/favorites/{favoriteId}", "2015-05-01"),
								}},
						},
						{
							Display:  "syntheticmonitorlocations",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/syntheticmonitorlocations", "2015-05-01"),
						}},
					SubResources: []swagger.ResourceType{
						{
							Display:  "{purgeId}",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/operations/{purgeId}", "2015-05-01"),
						}},
				},
				{
					Display:  "webtests",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{componentName}/webtests", "2015-05-01"),
				}},
		},
		{
			Display:  "webtests",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Insights/webtests", "2015-05-01"),
		},
		{
			Display:  "webtests",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/webtests", "2015-05-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{webTestName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/webtests/{webTestName}", "2015-05-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/webtests/{webTestName}", "2015-05-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/webtests/{webTestName}", "2015-05-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/webtests/{webTestName}", "2015-05-01"),
				}},
		},
		{
			Display:  "workbooks",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroup/{resourceGroupName}/providers/microsoft.insights/workbooks", "2015-05-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{resourceName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroup/{resourceGroupName}/providers/microsoft.insights/workbooks/{resourceName}", "2015-05-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroup/{resourceGroupName}/providers/microsoft.insights/workbooks/{resourceName}", "2015-05-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroup/{resourceGroupName}/providers/microsoft.insights/workbooks/{resourceName}", "2015-05-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroup/{resourceGroupName}/providers/microsoft.insights/workbooks/{resourceName}", "2015-05-01"),
				}},
		},
		{
			Display:  "classicAdministrators",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/classicAdministrators", "2015-07-01"),
		},
		{
			Display:  "providerOperations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Authorization/providerOperations", "2015-07-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:  "{resourceProviderNamespace}",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Authorization/providerOperations/{resourceProviderNamespace}", "2015-07-01"),
				}},
		},
		{
			Display:  "roleAssignments",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/roleAssignments", "2015-07-01"),
		},
		{
			Display:  "roleAssignments",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Authorization/roleAssignments", "2015-07-01"),
		},
		{
			Display:  "roleAssignments",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}/providers/Microsoft.Authorization/roleAssignments", "2015-07-01"),
		},
		{
			Display:        "{roleAssignmentId}",
			Endpoint:       endpoints.MustGetEndpointInfoFromURL("/{roleAssignmentId}", "2015-07-01"),
			DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{roleAssignmentId}", "2015-07-01"),
			PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/{roleAssignmentId}", "2015-07-01"),
		},
		{
			Display:  "roleAssignments",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Authorization/roleAssignments", "2015-07-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{roleAssignmentName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}", "2015-07-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}", "2015-07-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}", "2015-07-01"),
				}},
		},
		{
			Display:  "permissions",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Authorization/permissions", "2015-07-01"),
		},
		{
			Display:  "permissions",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}/providers/Microsoft.Authorization/permissions", "2015-07-01"),
		},
		{
			Display:  "{roleDefinitionId}",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/{roleDefinitionId}", "2015-07-01"),
		},
		{
			Display:  "roleDefinitions",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Authorization/roleDefinitions", "2015-07-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{roleDefinitionId}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Authorization/roleDefinitions/{roleDefinitionId}", "2015-07-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Authorization/roleDefinitions/{roleDefinitionId}", "2015-07-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Authorization/roleDefinitions/{roleDefinitionId}", "2015-07-01"),
				}},
		},
		{
			Display:  "python2Packages",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/python2Packages", "2018-06-30"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{packageName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/python2Packages/{packageName}", "2018-06-30"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/python2Packages/{packageName}", "2018-06-30"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/python2Packages/{packageName}", "2018-06-30"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/python2Packages/{packageName}", "2018-06-30"),
				}},
		},
		{
			Display:  "runbooks",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks", "2018-06-30"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{runbookName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}", "2018-06-30"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}", "2018-06-30"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}", "2018-06-30"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}", "2018-06-30"),
					Children: []swagger.ResourceType{
						{
							Display:  "content",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/content", "2018-06-30"),
						},
						{
							Display:  "draft",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft", "2018-06-30"),
							Children: []swagger.ResourceType{
								{
									Display:     "content",
									Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/content", "2018-06-30"),
									PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/content", "2018-06-30"),
								},
								{
									Display:     "testJob",
									Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/testJob", "2018-06-30"),
									PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/testJob", "2018-06-30"),
									Children: []swagger.ResourceType{
										{
											Display:  "streams",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/testJob/streams", "2018-06-30"),
											SubResources: []swagger.ResourceType{
												{
													Display:  "{jobStreamId}",
													Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/testJob/streams/{jobStreamId}", "2018-06-30"),
												}},
										}},
								}},
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Kusto/operations", "2019-09-07"),
		},
		{
			Display:  "clusters",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Kusto/clusters", "2019-09-07"),
		},
		{
			Display:  "skus",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Kusto/skus", "2019-09-07"),
		},
		{
			Display:  "clusters",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters", "2019-09-07"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{clusterName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}", "2019-09-07"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}", "2019-09-07"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}", "2019-09-07"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}", "2019-09-07"),
					Children: []swagger.ResourceType{
						{
							Display:  "attachedDatabaseConfigurations",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/attachedDatabaseConfigurations", "2019-09-07"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{attachedDatabaseConfigurationName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/attachedDatabaseConfigurations/{attachedDatabaseConfigurationName}", "2019-09-07"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/attachedDatabaseConfigurations/{attachedDatabaseConfigurationName}", "2019-09-07"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/attachedDatabaseConfigurations/{attachedDatabaseConfigurationName}", "2019-09-07"),
								}},
						},
						{
							Display:  "databases",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases", "2019-09-07"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{databaseName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}", "2019-09-07"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}", "2019-09-07"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}", "2019-09-07"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}", "2019-09-07"),
									Children: []swagger.ResourceType{
										{
											Display:  "dataConnections",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/dataConnections", "2019-09-07"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{dataConnectionName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/dataConnections/{dataConnectionName}", "2019-09-07"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/dataConnections/{dataConnectionName}", "2019-09-07"),
													PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/dataConnections/{dataConnectionName}", "2019-09-07"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/dataConnections/{dataConnectionName}", "2019-09-07"),
												}},
										}},
								}},
						},
						{
							Display:  "skus",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/skus", "2019-09-07"),
						}},
				}},
		},
		{
			Display:  "diagnosticSettings",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/microsoft.aadiam/diagnosticSettings", "2017-04-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{name}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/providers/microsoft.aadiam/diagnosticSettings/{name}", "2017-04-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/microsoft.aadiam/diagnosticSettings/{name}", "2017-04-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/providers/microsoft.aadiam/diagnosticSettings/{name}", "2017-04-01"),
				}},
		},
		{
			Display:  "diagnosticSettingsCategories",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/microsoft.aadiam/diagnosticSettingsCategories", "2017-04-01"),
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/microsoft.aadiam/operations", "2017-04-01"),
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.AzureStack/operations", "2017-06-01"),
		},
		{
			Display:  "customerSubscriptions",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}/customerSubscriptions", "2017-06-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{customerSubscriptionName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}/customerSubscriptions/{customerSubscriptionName}", "2017-06-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}/customerSubscriptions/{customerSubscriptionName}", "2017-06-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}/customerSubscriptions/{customerSubscriptionName}", "2017-06-01"),
				}},
		},
		{
			Display:  "products",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}/products", "2017-06-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:  "{productName}",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}/products/{productName}", "2017-06-01"),
					Children: []swagger.ResourceType{},
				}},
		},
		{
			Display:  "registrations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations", "2017-06-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{registrationName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}", "2017-06-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}", "2017-06-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}", "2017-06-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.AzureStack/registrations/{registrationName}", "2017-06-01"),
					Children:       []swagger.ResourceType{},
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Batch/operations", "2019-08-01"),
		},
		{
			Display:  "batchAccounts",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Batch/batchAccounts", "2019-08-01"),
		},
		{
			Display:  "quotas",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Batch/locations/{locationName}/quotas", "2019-08-01"),
		},
		{
			Display:  "batchAccounts",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{accountName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}", "2019-08-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}", "2019-08-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "applications",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{applicationName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationName}", "2019-08-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationName}", "2019-08-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationName}", "2019-08-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationName}", "2019-08-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "versions",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationName}/versions", "2019-08-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{versionName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationName}/versions/{versionName}", "2019-08-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationName}/versions/{versionName}", "2019-08-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationName}/versions/{versionName}", "2019-08-01"),
													Children:       []swagger.ResourceType{},
												}},
										}},
								}},
						},
						{
							Display:  "certificates",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/certificates", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{certificateName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/certificates/{certificateName}", "2019-08-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/certificates/{certificateName}", "2019-08-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/certificates/{certificateName}", "2019-08-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/certificates/{certificateName}", "2019-08-01"),
									Children:       []swagger.ResourceType{},
								}},
						},
						{
							Display:  "pools",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/pools", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{poolName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/pools/{poolName}", "2019-08-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/pools/{poolName}", "2019-08-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/pools/{poolName}", "2019-08-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/pools/{poolName}", "2019-08-01"),
									Children:       []swagger.ResourceType{},
								}},
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.BatchAI/operations", "2018-05-01"),
		},
		{
			Display:  "usages",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.BatchAI/locations/{location}/usages", "2018-05-01"),
		},
		{
			Display:  "workspaces",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.BatchAI/workspaces", "2018-05-01"),
		},
		{
			Display:  "workspaces",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces", "2018-05-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{workspaceName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}", "2018-05-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}", "2018-05-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}", "2018-05-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}", "2018-05-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "clusters",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/clusters", "2018-05-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{clusterName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/clusters/{clusterName}", "2018-05-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/clusters/{clusterName}", "2018-05-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/clusters/{clusterName}", "2018-05-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/clusters/{clusterName}", "2018-05-01"),
									Children:       []swagger.ResourceType{},
								}},
						},
						{
							Display:  "experiments",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/experiments", "2018-05-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{experimentName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/experiments/{experimentName}", "2018-05-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/experiments/{experimentName}", "2018-05-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/experiments/{experimentName}", "2018-05-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "jobs",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/experiments/{experimentName}/jobs", "2018-05-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{jobName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/experiments/{experimentName}/jobs/{jobName}", "2018-05-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/experiments/{experimentName}/jobs/{jobName}", "2018-05-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/experiments/{experimentName}/jobs/{jobName}", "2018-05-01"),
													Children:       []swagger.ResourceType{},
												}},
										}},
								}},
						},
						{
							Display:  "fileServers",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/fileServers", "2018-05-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{fileServerName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/fileServers/{fileServerName}", "2018-05-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/fileServers/{fileServerName}", "2018-05-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BatchAI/workspaces/{workspaceName}/fileServers/{fileServerName}", "2018-05-01"),
								}},
						}},
				}},
		},
		{
			Display:  "edgenodes",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Cdn/edgenodes", "2019-04-15"),
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Cdn/operations", "2019-04-15"),
		},
		{
			Display:  "profiles",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Cdn/profiles", "2019-04-15"),
		},
		{
			Display:  "profiles",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles", "2019-04-15"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{profileName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}", "2019-04-15"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}", "2019-04-15"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}", "2019-04-15"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}", "2019-04-15"),
					Children: []swagger.ResourceType{
						{
							Display:  "endpoints",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints", "2019-04-15"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{endpointName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}", "2019-04-15"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}", "2019-04-15"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}", "2019-04-15"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}", "2019-04-15"),
									Children: []swagger.ResourceType{
										{
											Display:  "customDomains",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/customDomains", "2019-04-15"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{customDomainName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/customDomains/{customDomainName}", "2019-04-15"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/customDomains/{customDomainName}", "2019-04-15"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/customDomains/{customDomainName}", "2019-04-15"),
													Children:       []swagger.ResourceType{},
												}},
										},
										{
											Display:  "origins",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/origins", "2019-04-15"),
											SubResources: []swagger.ResourceType{
												{
													Display:       "{originName}",
													Endpoint:      endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/origins/{originName}", "2019-04-15"),
													PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/origins/{originName}", "2019-04-15"),
												}},
										}},
								}},
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.CognitiveServices/operations", "2017-04-18"),
		},
		{
			Display:  "accounts",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.CognitiveServices/accounts", "2017-04-18"),
		},
		{
			Display:  "skus",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.CognitiveServices/skus", "2017-04-18"),
		},
		{
			Display:  "accounts",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts", "2017-04-18"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{accountName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}", "2017-04-18"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}", "2017-04-18"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}", "2017-04-18"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}", "2017-04-18"),
					Children: []swagger.ResourceType{
						{
							Display:  "skus",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/skus", "2017-04-18"),
						},
						{
							Display:  "usages",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/usages", "2017-04-18"),
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Compute/operations", "2019-07-01"),
		},
		{
			Display:  "availabilitySets",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/availabilitySets", "2019-07-01"),
		},
		{
			Display:  "hostGroups",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/hostGroups", "2019-07-01"),
		},
		{
			Display:  "images",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/images", "2019-07-01"),
		},
		{
			Display:  "publishers",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers", "2019-07-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:  "types",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmextension/types", "2019-07-01"),
					SubResources: []swagger.ResourceType{
						{
							Display:  "versions",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmextension/types/{type}/versions", "2019-07-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:  "{version}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmextension/types/{type}/versions/{version}", "2019-07-01"),
								}},
						}},
				},
				{
					Display:  "offers",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmimage/offers", "2019-07-01"),
					SubResources: []swagger.ResourceType{
						{
							Display:  "skus",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus", "2019-07-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:  "versions",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus/{skus}/versions", "2019-07-01"),
									SubResources: []swagger.ResourceType{
										{
											Display:  "{version}",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus/{skus}/versions/{version}", "2019-07-01"),
										}},
								}},
						}},
				}},
		},
		{
			Display:  "usages",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/usages", "2019-07-01"),
		},
		{
			Display:  "virtualMachines",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/virtualMachines", "2019-07-01"),
		},
		{
			Display:  "vmSizes",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/vmSizes", "2019-07-01"),
		},
		{
			Display:  "proximityPlacementGroups",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/proximityPlacementGroups", "2019-07-01"),
		},
		{
			Display:  "virtualMachineScaleSets",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/virtualMachineScaleSets", "2019-07-01"),
		},
		{
			Display:  "virtualMachines",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/virtualMachines", "2019-07-01"),
		},
		{
			Display:  "availabilitySets",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets", "2019-07-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{availabilitySetName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}", "2019-07-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}", "2019-07-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}", "2019-07-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}", "2019-07-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "vmSizes",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}/vmSizes", "2019-07-01"),
						}},
				}},
		},
		{
			Display:  "hostGroups",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups", "2019-07-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{hostGroupName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}", "2019-07-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}", "2019-07-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}", "2019-07-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}", "2019-07-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "hosts",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}/hosts", "2019-07-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{hostName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}/hosts/{hostName}", "2019-07-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}/hosts/{hostName}", "2019-07-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}/hosts/{hostName}", "2019-07-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}/hosts/{hostName}", "2019-07-01"),
								}},
						}},
				}},
		},
		{
			Display:  "images",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images", "2019-07-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{imageName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images/{imageName}", "2019-07-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images/{imageName}", "2019-07-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images/{imageName}", "2019-07-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images/{imageName}", "2019-07-01"),
				}},
		},
		{
			Display:  "proximityPlacementGroups",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/proximityPlacementGroups", "2019-07-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{proximityPlacementGroupName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/proximityPlacementGroups/{proximityPlacementGroupName}", "2019-07-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/proximityPlacementGroups/{proximityPlacementGroupName}", "2019-07-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/proximityPlacementGroups/{proximityPlacementGroupName}", "2019-07-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/proximityPlacementGroups/{proximityPlacementGroupName}", "2019-07-01"),
				}},
		},
		{
			Display:  "virtualMachineScaleSets",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets", "2019-07-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{vmScaleSetName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}", "2019-07-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}", "2019-07-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}", "2019-07-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}", "2019-07-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "extensions",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/extensions", "2019-07-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{vmssExtensionName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/extensions/{vmssExtensionName}", "2019-07-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/extensions/{vmssExtensionName}", "2019-07-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/extensions/{vmssExtensionName}", "2019-07-01"),
								}},
						},
						{
							Display:  "instanceView",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/instanceView", "2019-07-01"),
						},
						{
							Display:  "osUpgradeHistory",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/osUpgradeHistory", "2019-07-01"),
						},
						{
							Display:  "latest",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/rollingUpgrades/latest", "2019-07-01"),
						},
						{
							Display:  "skus",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/skus", "2019-07-01"),
						},
						{
							Display:  "virtualMachines",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines", "2019-07-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{instanceId}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}", "2019-07-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}", "2019-07-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}", "2019-07-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "extensions",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/extensions", "2019-07-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{vmExtensionName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/extensions/{vmExtensionName}", "2019-07-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/extensions/{vmExtensionName}", "2019-07-01"),
													PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/extensions/{vmExtensionName}", "2019-07-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/extensions/{vmExtensionName}", "2019-07-01"),
												}},
										},
										{
											Display:  "instanceView",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/instanceView", "2019-07-01"),
										}},
								},
								{
									Display:  "publicipaddresses",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces/{networkInterfaceName}/ipconfigurations/{ipConfigurationName}/publicipaddresses", "2018-10-01"),
								}},
						},
						{
							Display:  "publicipaddresses",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/publicipaddresses", "2018-10-01"),
						}},
					SubResources: []swagger.ResourceType{},
				},
				{
					Display:  "{publicIpAddressName}",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces/{networkInterfaceName}/ipconfigurations/{ipConfigurationName}/publicipaddresses/{publicIpAddressName}", "2018-10-01"),
				}},
		},
		{
			Display:  "virtualMachines",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines", "2019-07-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{vmName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}", "2019-07-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}", "2019-07-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}", "2019-07-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}", "2019-07-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "extensions",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/extensions", "2019-07-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{vmExtensionName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/extensions/{vmExtensionName}", "2019-07-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/extensions/{vmExtensionName}", "2019-07-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/extensions/{vmExtensionName}", "2019-07-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/extensions/{vmExtensionName}", "2019-07-01"),
								}},
						},
						{
							Display:  "instanceView",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/instanceView", "2019-07-01"),
						},
						{
							Display:  "vmSizes",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/vmSizes", "2019-07-01"),
						},
						{
							Display:  "guestConfigurationAssignments",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments", "2018-11-20"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{guestConfigurationAssignmentName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{guestConfigurationAssignmentName}", "2018-11-20"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{guestConfigurationAssignmentName}", "2018-11-20"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{guestConfigurationAssignmentName}", "2018-11-20"),
									Children: []swagger.ResourceType{
										{
											Display:  "reports",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{guestConfigurationAssignmentName}/reports", "2018-11-20"),
											SubResources: []swagger.ResourceType{
												{
													Display:  "{reportId}",
													Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{guestConfigurationAssignmentName}/reports/{reportId}", "2018-11-20"),
												}},
										}},
								}},
						}},
				}},
		},
		{
			Display:  "diskEncryptionSets",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/diskEncryptionSets", "2019-07-01"),
		},
		{
			Display:  "disks",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/disks", "2019-07-01"),
		},
		{
			Display:  "snapshots",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/snapshots", "2019-07-01"),
		},
		{
			Display:  "diskEncryptionSets",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets", "2019-07-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{diskEncryptionSetName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}", "2019-07-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}", "2019-07-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}", "2019-07-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}", "2019-07-01"),
				}},
		},
		{
			Display:  "disks",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks", "2019-07-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{diskName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}", "2019-07-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}", "2019-07-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}", "2019-07-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}", "2019-07-01"),
					Children:       []swagger.ResourceType{},
				}},
		},
		{
			Display:  "snapshots",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots", "2019-07-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{snapshotName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}", "2019-07-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}", "2019-07-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}", "2019-07-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}", "2019-07-01"),
					Children:       []swagger.ResourceType{},
				}},
		},
		{
			Display:  "galleries",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/galleries", "2019-07-01"),
		},
		{
			Display:  "galleries",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries", "2019-07-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{galleryName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}", "2019-07-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}", "2019-07-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}", "2019-07-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "applications",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications", "2019-07-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{galleryApplicationName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}", "2019-07-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}", "2019-07-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}", "2019-07-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "versions",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}/versions", "2019-07-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{galleryApplicationVersionName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}/versions/{galleryApplicationVersionName}", "2019-07-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}/versions/{galleryApplicationVersionName}", "2019-07-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}/versions/{galleryApplicationVersionName}", "2019-07-01"),
												}},
										}},
								}},
						},
						{
							Display:  "images",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/images", "2019-07-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{galleryImageName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/images/{galleryImageName}", "2019-07-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/images/{galleryImageName}", "2019-07-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/images/{galleryImageName}", "2019-07-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "versions",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/images/{galleryImageName}/versions", "2019-07-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{galleryImageVersionName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/images/{galleryImageName}/versions/{galleryImageVersionName}", "2019-07-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/images/{galleryImageName}/versions/{galleryImageVersionName}", "2019-07-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/images/{galleryImageName}/versions/{galleryImageVersionName}", "2019-07-01"),
												}},
										}},
								}},
						}},
				}},
		},
		{
			Display:  "runCommands",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/runCommands", "2019-07-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:  "{commandId}",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/runCommands/{commandId}", "2019-07-01"),
				}},
		},
		{
			Display:  "containerServices",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ContainerService/containerServices", "2017-01-31"),
		},
		{
			Display:  "containerServices",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices", "2017-01-31"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{containerServiceName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices/{containerServiceName}", "2017-01-31"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices/{containerServiceName}", "2017-01-31"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices/{containerServiceName}", "2017-01-31"),
				}},
		},
		{
			Display:  "balanceSummary",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/providers/Microsoft.Consumption/credits/balanceSummary", "2019-10-01"),
		},
		{
			Display:  "events",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/providers/Microsoft.Consumption/events", "2019-10-01"),
		},
		{
			Display:  "lots",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/providers/Microsoft.Consumption/lots", "2019-10-01"),
		},
		{
			Display:  "balances",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/providers/Microsoft.Billing/billingPeriods/{billingPeriodName}/providers/Microsoft.Consumption/balances", "2019-10-01"),
		},
		{
			Display:  "balances",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/providers/Microsoft.Consumption/balances", "2019-10-01"),
		},
		{
			Display:  "reservationTransactions",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/providers/Microsoft.Consumption/reservationTransactions", "2019-10-01"),
		},
		{
			Display:  "reservationDetails",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Capacity/reservationorders/{reservationOrderId}/providers/Microsoft.Consumption/reservationDetails", "2019-10-01"),
		},
		{
			Display:  "reservationSummaries",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Capacity/reservationorders/{reservationOrderId}/providers/Microsoft.Consumption/reservationSummaries", "2019-10-01"),
		},
		{
			Display:  "reservationDetails",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Capacity/reservationorders/{reservationOrderId}/reservations/{reservationId}/providers/Microsoft.Consumption/reservationDetails", "2019-10-01"),
		},
		{
			Display:  "reservationSummaries",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Capacity/reservationorders/{reservationOrderId}/reservations/{reservationId}/providers/Microsoft.Consumption/reservationSummaries", "2019-10-01"),
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Consumption/operations", "2019-10-01"),
		},
		{
			Display:  "aggregatedcost",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Billing/billingPeriods/{billingPeriodName}/Microsoft.Consumption/aggregatedcost", "2019-10-01"),
		},
		{
			Display:  "aggregatedcost",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Consumption/aggregatedcost", "2019-10-01"),
		},
		{
			Display:  "default",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Billing/billingPeriods/{billingPeriodName}/providers/Microsoft.Consumption/pricesheets/default", "2019-10-01"),
		},
		{
			Display:  "forecasts",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Consumption/forecasts", "2019-10-01"),
		},
		{
			Display:  "default",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Consumption/pricesheets/default", "2019-10-01"),
		},
		{
			Display:  "budgets",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Consumption/budgets", "2019-10-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{budgetName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Consumption/budgets/{budgetName}", "2019-10-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Consumption/budgets/{budgetName}", "2019-10-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Consumption/budgets/{budgetName}", "2019-10-01"),
				}},
		},
		{
			Display:  "charges",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Consumption/charges", "2019-10-01"),
		},
		{
			Display:  "marketplaces",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Consumption/marketplaces", "2019-10-01"),
		},
		{
			Display:  "reservationDetails",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Consumption/reservationDetails", "2019-10-01"),
		},
		{
			Display:  "reservationRecommendations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Consumption/reservationRecommendations", "2019-10-01"),
		},
		{
			Display:  "reservationSummaries",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Consumption/reservationSummaries", "2019-10-01"),
		},
		{
			Display:  "tags",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Consumption/tags", "2019-10-01"),
		},
		{
			Display:  "usageDetails",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Consumption/usageDetails", "2019-10-01"),
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ContainerInstance/operations", "2018-10-01"),
		},
		{
			Display:  "containerGroups",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ContainerInstance/containerGroups", "2018-10-01"),
		},
		{
			Display:  "cachedImages",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ContainerInstance/locations/{location}/cachedImages", "2018-10-01"),
		},
		{
			Display:  "capabilities",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ContainerInstance/locations/{location}/capabilities", "2018-10-01"),
		},
		{
			Display:  "usages",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ContainerInstance/locations/{location}/usages", "2018-10-01"),
		},
		{
			Display:  "containerGroups",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups", "2018-10-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{containerGroupName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}", "2018-10-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}", "2018-10-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}", "2018-10-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}", "2018-10-01"),
					Children:       []swagger.ResourceType{},
					SubResources: []swagger.ResourceType{
						{
							Display:  "logs",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}/containers/{containerName}/logs", "2018-10-01"),
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ContainerRegistry/operations", "2019-05-01"),
		},
		{
			Display:  "registries",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ContainerRegistry/registries", "2019-05-01"),
		},
		{
			Display:  "registries",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries", "2019-05-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{registryName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}", "2019-05-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}", "2019-05-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}", "2019-05-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}", "2019-05-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "listUsages",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/listUsages", "2019-05-01"),
						},
						{
							Display:  "replications",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/replications", "2019-05-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{replicationName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/replications/{replicationName}", "2019-05-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/replications/{replicationName}", "2019-05-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/replications/{replicationName}", "2019-05-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/replications/{replicationName}", "2019-05-01"),
								}},
						},
						{
							Display:  "webhooks",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/webhooks", "2019-05-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{webhookName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/webhooks/{webhookName}", "2019-05-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/webhooks/{webhookName}", "2019-05-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/webhooks/{webhookName}", "2019-05-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/webhooks/{webhookName}", "2019-05-01"),
									Children:       []swagger.ResourceType{},
								}},
						}},
				}},
		},
		{
			Display:  "orchestrators",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ContainerService/locations/{location}/orchestrators", "2019-08-01"),
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ContainerService/operations", "2019-08-01"),
		},
		{
			Display:  "managedClusters",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ContainerService/managedClusters", "2019-08-01"),
		},
		{
			Display:  "managedClusters",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{resourceName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}", "2019-08-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}", "2019-08-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "agentPools",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/agentPools", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{agentPoolName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}", "2019-08-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}", "2019-08-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}", "2019-08-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "default",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}/upgradeProfiles/default", "2019-08-01"),
										}},
								}},
						},
						{
							Display:  "availableAgentPoolVersions",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/availableAgentPoolVersions", "2019-08-01"),
						},
						{
							Display:  "default",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/upgradeProfiles/default", "2019-08-01"),
						}},
					SubResources: []swagger.ResourceType{},
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.DocumentDB/operations", "2019-08-01"),
		},
		{
			Display:  "databaseAccounts",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/databaseAccounts", "2019-08-01"),
		},
		{
			Display:  "databaseAccounts",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{accountName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}", "2019-08-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}", "2019-08-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "cassandraKeyspaces",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/cassandraKeyspaces", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{keyspaceName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/cassandraKeyspaces/{keyspaceName}", "2019-08-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/cassandraKeyspaces/{keyspaceName}", "2019-08-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/cassandraKeyspaces/{keyspaceName}", "2019-08-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "tables",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/cassandraKeyspaces/{keyspaceName}/tables", "2019-08-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{tableName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/cassandraKeyspaces/{keyspaceName}/tables/{tableName}", "2019-08-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/cassandraKeyspaces/{keyspaceName}/tables/{tableName}", "2019-08-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/cassandraKeyspaces/{keyspaceName}/tables/{tableName}", "2019-08-01"),
													Children: []swagger.ResourceType{
														{
															Display:     "default",
															Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/cassandraKeyspaces/{keyspaceName}/tables/{tableName}/throughputSettings/default", "2019-08-01"),
															PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/cassandraKeyspaces/{keyspaceName}/tables/{tableName}/throughputSettings/default", "2019-08-01"),
														}},
												}},
										},
										{
											Display:     "default",
											Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/cassandraKeyspaces/{keyspaceName}/throughputSettings/default", "2019-08-01"),
											PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/cassandraKeyspaces/{keyspaceName}/throughputSettings/default", "2019-08-01"),
										}},
								}},
						},
						{
							Display:  "gremlinDatabases",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/gremlinDatabases", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{databaseName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/gremlinDatabases/{databaseName}", "2019-08-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/gremlinDatabases/{databaseName}", "2019-08-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/gremlinDatabases/{databaseName}", "2019-08-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "graphs",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/gremlinDatabases/{databaseName}/graphs", "2019-08-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{graphName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/gremlinDatabases/{databaseName}/graphs/{graphName}", "2019-08-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/gremlinDatabases/{databaseName}/graphs/{graphName}", "2019-08-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/gremlinDatabases/{databaseName}/graphs/{graphName}", "2019-08-01"),
													Children: []swagger.ResourceType{
														{
															Display:     "default",
															Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/gremlinDatabases/{databaseName}/graphs/{graphName}/throughputSettings/default", "2019-08-01"),
															PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/gremlinDatabases/{databaseName}/graphs/{graphName}/throughputSettings/default", "2019-08-01"),
														}},
												}},
										},
										{
											Display:     "default",
											Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/gremlinDatabases/{databaseName}/throughputSettings/default", "2019-08-01"),
											PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/gremlinDatabases/{databaseName}/throughputSettings/default", "2019-08-01"),
										}},
								}},
						},
						{
							Display:  "metricDefinitions",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/metricDefinitions", "2019-08-01"),
						},
						{
							Display:  "metrics",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/metrics", "2019-08-01"),
						},
						{
							Display:  "mongodbDatabases",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{databaseName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}", "2019-08-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}", "2019-08-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}", "2019-08-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "collections",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}/collections", "2019-08-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{collectionName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}/collections/{collectionName}", "2019-08-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}/collections/{collectionName}", "2019-08-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}/collections/{collectionName}", "2019-08-01"),
													Children: []swagger.ResourceType{
														{
															Display:     "default",
															Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}/collections/{collectionName}/throughputSettings/default", "2019-08-01"),
															PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}/collections/{collectionName}/throughputSettings/default", "2019-08-01"),
														}},
												}},
										},
										{
											Display:     "default",
											Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}/throughputSettings/default", "2019-08-01"),
											PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}/throughputSettings/default", "2019-08-01"),
										}},
								}},
						},
						{
							Display:  "metrics",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/percentile/metrics", "2019-08-01"),
						},
						{
							Display:  "readonlykeys",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/readonlykeys", "2019-08-01"),
						},
						{
							Display:  "sqlDatabases",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{databaseName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}", "2019-08-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}", "2019-08-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}", "2019-08-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "containers",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers", "2019-08-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{containerName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}", "2019-08-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}", "2019-08-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}", "2019-08-01"),
													Children: []swagger.ResourceType{
														{
															Display:  "storedProcedures",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/storedProcedures", "2019-08-01"),
															SubResources: []swagger.ResourceType{
																{
																	Display:        "{storedProcedureName}",
																	Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/storedProcedures/{storedProcedureName}", "2019-08-01"),
																	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/storedProcedures/{storedProcedureName}", "2019-08-01"),
																	PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/storedProcedures/{storedProcedureName}", "2019-08-01"),
																}},
														},
														{
															Display:     "default",
															Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/throughputSettings/default", "2019-08-01"),
															PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/throughputSettings/default", "2019-08-01"),
														},
														{
															Display:  "triggers",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/triggers", "2019-08-01"),
															SubResources: []swagger.ResourceType{
																{
																	Display:        "{triggerName}",
																	Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/triggers/{triggerName}", "2019-08-01"),
																	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/triggers/{triggerName}", "2019-08-01"),
																	PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/triggers/{triggerName}", "2019-08-01"),
																}},
														},
														{
															Display:  "userDefinedFunctions",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/userDefinedFunctions", "2019-08-01"),
															SubResources: []swagger.ResourceType{
																{
																	Display:        "{userDefinedFunctionName}",
																	Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/userDefinedFunctions/{userDefinedFunctionName}", "2019-08-01"),
																	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/userDefinedFunctions/{userDefinedFunctionName}", "2019-08-01"),
																	PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/userDefinedFunctions/{userDefinedFunctionName}", "2019-08-01"),
																}},
														}},
												}},
										},
										{
											Display:     "default",
											Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/throughputSettings/default", "2019-08-01"),
											PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/throughputSettings/default", "2019-08-01"),
										}},
								}},
						},
						{
							Display:  "tables",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{tableName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}", "2019-08-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}", "2019-08-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}", "2019-08-01"),
									Children: []swagger.ResourceType{
										{
											Display:     "default",
											Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}/throughputSettings/default", "2019-08-01"),
											PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}/throughputSettings/default", "2019-08-01"),
										}},
								}},
						},
						{
							Display:  "usages",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/usages", "2019-08-01"),
						}},
					SubResources: []swagger.ResourceType{
						{
							Display:  "metricDefinitions",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/collections/{collectionRid}/metricDefinitions", "2019-08-01"),
						},
						{
							Display:  "metrics",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/collections/{collectionRid}/metrics", "2019-08-01"),
						},
						{
							Display:  "metrics",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/collections/{collectionRid}/partitionKeyRangeId/{partitionKeyRangeId}/metrics", "2019-08-01"),
						},
						{
							Display:  "metrics",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/collections/{collectionRid}/partitions/metrics", "2019-08-01"),
						},
						{
							Display:  "usages",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/collections/{collectionRid}/partitions/usages", "2019-08-01"),
						},
						{
							Display:  "usages",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/collections/{collectionRid}/usages", "2019-08-01"),
						},
						{
							Display:  "metricDefinitions",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/metricDefinitions", "2019-08-01"),
						},
						{
							Display:  "metrics",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/metrics", "2019-08-01"),
						},
						{
							Display:  "usages",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/usages", "2019-08-01"),
						},
						{
							Display:  "metrics",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/region/{region}/databases/{databaseRid}/collections/{collectionRid}/metrics", "2019-08-01"),
						},
						{
							Display:  "metrics",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/region/{region}/databases/{databaseRid}/collections/{collectionRid}/partitionKeyRangeId/{partitionKeyRangeId}/metrics", "2019-08-01"),
						},
						{
							Display:  "metrics",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/region/{region}/databases/{databaseRid}/collections/{collectionRid}/partitions/metrics", "2019-08-01"),
						},
						{
							Display:  "metrics",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/region/{region}/metrics", "2019-08-01"),
						},
						{
							Display:  "metrics",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sourceRegion/{sourceRegion}/targetRegion/{targetRegion}/percentile/metrics", "2019-08-01"),
						},
						{
							Display:  "metrics",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/targetRegion/{targetRegion}/percentile/metrics", "2019-08-01"),
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.CostManagement/operations", "2019-10-01"),
		},
		{
			Display:  "dimensions",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.CostManagement/dimensions", "2019-10-01"),
		},
		{
			Display:  "exports",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.CostManagement/exports", "2019-10-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{exportName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.CostManagement/exports/{exportName}", "2019-10-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.CostManagement/exports/{exportName}", "2019-10-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.CostManagement/exports/{exportName}", "2019-10-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "runHistory",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.CostManagement/exports/{exportName}/runHistory", "2019-10-01"),
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.CustomerInsights/operations", "2017-04-26"),
		},
		{
			Display:  "hubs",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.CustomerInsights/hubs", "2017-04-26"),
		},
		{
			Display:  "hubs",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs", "2017-04-26"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{hubName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}", "2017-04-26"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}", "2017-04-26"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}", "2017-04-26"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}", "2017-04-26"),
					Children: []swagger.ResourceType{
						{
							Display:  "authorizationPolicies",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/authorizationPolicies", "2017-04-26"),
							SubResources: []swagger.ResourceType{
								{
									Display:     "{authorizationPolicyName}",
									Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/authorizationPolicies/{authorizationPolicyName}", "2017-04-26"),
									PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/authorizationPolicies/{authorizationPolicyName}", "2017-04-26"),
									Children:    []swagger.ResourceType{},
								}},
						},
						{
							Display:  "connectors",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/connectors", "2017-04-26"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{connectorName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/connectors/{connectorName}", "2017-04-26"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/connectors/{connectorName}", "2017-04-26"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/connectors/{connectorName}", "2017-04-26"),
									Children: []swagger.ResourceType{
										{
											Display:  "mappings",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/connectors/{connectorName}/mappings", "2017-04-26"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{mappingName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/connectors/{connectorName}/mappings/{mappingName}", "2017-04-26"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/connectors/{connectorName}/mappings/{mappingName}", "2017-04-26"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/connectors/{connectorName}/mappings/{mappingName}", "2017-04-26"),
												}},
										}},
								}},
						},
						{
							Display:  "interactions",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/interactions", "2017-04-26"),
							SubResources: []swagger.ResourceType{
								{
									Display:     "{interactionName}",
									Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/interactions/{interactionName}", "2017-04-26"),
									PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/interactions/{interactionName}", "2017-04-26"),
									Children:    []swagger.ResourceType{},
								}},
						},
						{
							Display:  "kpi",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/kpi", "2017-04-26"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{kpiName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/kpi/{kpiName}", "2017-04-26"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/kpi/{kpiName}", "2017-04-26"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/kpi/{kpiName}", "2017-04-26"),
									Children:       []swagger.ResourceType{},
								}},
						},
						{
							Display:  "links",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/links", "2017-04-26"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{linkName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/links/{linkName}", "2017-04-26"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/links/{linkName}", "2017-04-26"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/links/{linkName}", "2017-04-26"),
								}},
						},
						{
							Display:  "predictions",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/predictions", "2017-04-26"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{predictionName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/predictions/{predictionName}", "2017-04-26"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/predictions/{predictionName}", "2017-04-26"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/predictions/{predictionName}", "2017-04-26"),
									Children:       []swagger.ResourceType{},
								}},
						},
						{
							Display:  "profiles",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/profiles", "2017-04-26"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{profileName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/profiles/{profileName}", "2017-04-26"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/profiles/{profileName}", "2017-04-26"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/profiles/{profileName}", "2017-04-26"),
									Children:       []swagger.ResourceType{},
								}},
						},
						{
							Display:  "relationshipLinks",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationshipLinks", "2017-04-26"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{relationshipLinkName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationshipLinks/{relationshipLinkName}", "2017-04-26"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationshipLinks/{relationshipLinkName}", "2017-04-26"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationshipLinks/{relationshipLinkName}", "2017-04-26"),
								}},
						},
						{
							Display:  "relationships",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationships", "2017-04-26"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{relationshipName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationships/{relationshipName}", "2017-04-26"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationships/{relationshipName}", "2017-04-26"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationships/{relationshipName}", "2017-04-26"),
								}},
						},
						{
							Display:  "roleAssignments",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/roleAssignments", "2017-04-26"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{assignmentName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/roleAssignments/{assignmentName}", "2017-04-26"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/roleAssignments/{assignmentName}", "2017-04-26"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/roleAssignments/{assignmentName}", "2017-04-26"),
								}},
						},
						{
							Display:  "roles",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/roles", "2017-04-26"),
						},
						{
							Display:  "views",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/views", "2017-04-26"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{viewName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/views/{viewName}", "2017-04-26"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/views/{viewName}", "2017-04-26"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/views/{viewName}", "2017-04-26"),
								}},
						},
						{
							Display:  "widgetTypes",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/widgetTypes", "2017-04-26"),
							SubResources: []swagger.ResourceType{
								{
									Display:  "{widgetTypeName}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/widgetTypes/{widgetTypeName}", "2017-04-26"),
								}},
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.DataBox/operations", "2019-09-01"),
		},
		{
			Display:  "jobs",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DataBox/jobs", "2019-09-01"),
		},
		{
			Display:  "jobs",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBox/jobs", "2019-09-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{jobName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBox/jobs/{jobName}", "2019-09-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBox/jobs/{jobName}", "2019-09-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBox/jobs/{jobName}", "2019-09-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBox/jobs/{jobName}", "2019-09-01"),
					Children:       []swagger.ResourceType{},
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.DataBoxEdge/operations", "2019-07-01"),
		},
		{
			Display:  "dataBoxEdgeDevices",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices", "2019-07-01"),
		},
		{
			Display:  "dataBoxEdgeDevices",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices", "2019-07-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{deviceName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}", "2019-07-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}", "2019-07-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}", "2019-07-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}", "2019-07-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "alerts",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/alerts", "2019-07-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:  "{name}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/alerts/{name}", "2019-07-01"),
								}},
						},
						{
							Display:  "bandwidthSchedules",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/bandwidthSchedules", "2019-07-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{name}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/bandwidthSchedules/{name}", "2019-07-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/bandwidthSchedules/{name}", "2019-07-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/bandwidthSchedules/{name}", "2019-07-01"),
								}},
						},
						{
							Display:  "default",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/networkSettings/default", "2019-07-01"),
						},
						{
							Display:  "nodes",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/nodes", "2019-07-01"),
						},
						{
							Display:  "orders",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/orders", "2019-07-01"),
							Children: []swagger.ResourceType{
								{
									Display:        "default",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/orders/default", "2019-07-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/orders/default", "2019-07-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/orders/default", "2019-07-01"),
								}},
						},
						{
							Display:  "roles",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles", "2019-07-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{name}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles/{name}", "2019-07-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles/{name}", "2019-07-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles/{name}", "2019-07-01"),
								}},
						},
						{
							Display:  "shares",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/shares", "2019-07-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{name}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/shares/{name}", "2019-07-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/shares/{name}", "2019-07-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/shares/{name}", "2019-07-01"),
									Children:       []swagger.ResourceType{},
								}},
						},
						{
							Display:  "storageAccountCredentials",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccountCredentials", "2019-07-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{name}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccountCredentials/{name}", "2019-07-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccountCredentials/{name}", "2019-07-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccountCredentials/{name}", "2019-07-01"),
								}},
						},
						{
							Display:  "triggers",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/triggers", "2019-07-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{name}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/triggers/{name}", "2019-07-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/triggers/{name}", "2019-07-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/triggers/{name}", "2019-07-01"),
								}},
						},
						{
							Display:  "default",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/updateSummary/default", "2019-07-01"),
						},
						{
							Display:  "users",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/users", "2019-07-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{name}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/users/{name}", "2019-07-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/users/{name}", "2019-07-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/users/{name}", "2019-07-01"),
								}},
						}},
					SubResources: []swagger.ResourceType{
						{
							Display:  "{name}",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/jobs/{name}", "2019-07-01"),
						},
						{
							Display:  "{name}",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/operationsStatus/{name}", "2019-07-01"),
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Databricks/operations", "2018-04-01"),
		},
		{
			Display:  "workspaces",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Databricks/workspaces", "2018-04-01"),
		},
		{
			Display:  "workspaces",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Databricks/workspaces", "2018-04-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{workspaceName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Databricks/workspaces/{workspaceName}", "2018-04-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Databricks/workspaces/{workspaceName}", "2018-04-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Databricks/workspaces/{workspaceName}", "2018-04-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Databricks/workspaces/{workspaceName}", "2018-04-01"),
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.DataCatalog/operations", "2016-03-30"),
		},
		{
			Display:  "catalogs",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataCatalog/catalogs", "2016-03-30"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{catalogName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataCatalog/catalogs/{catalogName}", "2016-03-30"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataCatalog/catalogs/{catalogName}", "2016-03-30"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataCatalog/catalogs/{catalogName}", "2016-03-30"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataCatalog/catalogs/{catalogName}", "2016-03-30"),
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.DataFactory/operations", "2018-06-01"),
		},
		{
			Display:  "factories",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DataFactory/factories", "2018-06-01"),
		},
		{
			Display:  "factories",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories", "2018-06-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{factoryName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}", "2018-06-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}", "2018-06-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}", "2018-06-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}", "2018-06-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "dataflows",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/dataflows", "2018-06-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{dataFlowName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/dataflows/{dataFlowName}", "2018-06-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/dataflows/{dataFlowName}", "2018-06-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/dataflows/{dataFlowName}", "2018-06-01"),
								}},
						},
						{
							Display:  "datasets",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/datasets", "2018-06-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{datasetName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/datasets/{datasetName}", "2018-06-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/datasets/{datasetName}", "2018-06-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/datasets/{datasetName}", "2018-06-01"),
								}},
						},
						{
							Display:  "integrationRuntimes",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/integrationRuntimes", "2018-06-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{integrationRuntimeName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/integrationRuntimes/{integrationRuntimeName}", "2018-06-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/integrationRuntimes/{integrationRuntimeName}", "2018-06-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/integrationRuntimes/{integrationRuntimeName}", "2018-06-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/integrationRuntimes/{integrationRuntimeName}", "2018-06-01"),
									Children:       []swagger.ResourceType{},
									SubResources: []swagger.ResourceType{
										{
											Display:        "{nodeName}",
											Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/integrationRuntimes/{integrationRuntimeName}/nodes/{nodeName}", "2018-06-01"),
											DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/integrationRuntimes/{integrationRuntimeName}/nodes/{nodeName}", "2018-06-01"),
											PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/integrationRuntimes/{integrationRuntimeName}/nodes/{nodeName}", "2018-06-01"),
											Children:       []swagger.ResourceType{},
										}},
								}},
						},
						{
							Display:  "linkedservices",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/linkedservices", "2018-06-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{linkedServiceName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/linkedservices/{linkedServiceName}", "2018-06-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/linkedservices/{linkedServiceName}", "2018-06-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/linkedservices/{linkedServiceName}", "2018-06-01"),
								}},
						},
						{
							Display:  "pipelines",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelines", "2018-06-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{pipelineName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelines/{pipelineName}", "2018-06-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelines/{pipelineName}", "2018-06-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelines/{pipelineName}", "2018-06-01"),
									Children:       []swagger.ResourceType{},
								}},
						},
						{
							Display:  "triggers",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/triggers", "2018-06-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{triggerName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/triggers/{triggerName}", "2018-06-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/triggers/{triggerName}", "2018-06-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/triggers/{triggerName}", "2018-06-01"),
									Children: []swagger.ResourceType{
										{
											Display:      "rerunTriggers",
											Endpoint:     endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/triggers/{triggerName}/rerunTriggers", "2018-06-01"),
											SubResources: []swagger.ResourceType{},
										}},
									SubResources: []swagger.ResourceType{},
								}},
						}},
					SubResources: []swagger.ResourceType{
						{
							Display:  "{runId}",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelineruns/{runId}", "2018-06-01"),
							Children: []swagger.ResourceType{},
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.DataLakeAnalytics/operations", "2016-11-01"),
		},
		{
			Display:  "accounts",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DataLakeAnalytics/accounts", "2016-11-01"),
		},
		{
			Display:  "capability",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DataLakeAnalytics/locations/{location}/capability", "2016-11-01"),
		},
		{
			Display:  "accounts",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts", "2016-11-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{accountName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}", "2016-11-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}", "2016-11-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}", "2016-11-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}", "2016-11-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "computePolicies",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/computePolicies", "2016-11-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{computePolicyName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/computePolicies/{computePolicyName}", "2016-11-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/computePolicies/{computePolicyName}", "2016-11-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/computePolicies/{computePolicyName}", "2016-11-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/computePolicies/{computePolicyName}", "2016-11-01"),
								}},
						},
						{
							Display:  "dataLakeStoreAccounts",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/dataLakeStoreAccounts", "2016-11-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{dataLakeStoreAccountName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/dataLakeStoreAccounts/{dataLakeStoreAccountName}", "2016-11-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/dataLakeStoreAccounts/{dataLakeStoreAccountName}", "2016-11-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/dataLakeStoreAccounts/{dataLakeStoreAccountName}", "2016-11-01"),
								}},
						},
						{
							Display:  "firewallRules",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/firewallRules", "2016-11-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{firewallRuleName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/firewallRules/{firewallRuleName}", "2016-11-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/firewallRules/{firewallRuleName}", "2016-11-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/firewallRules/{firewallRuleName}", "2016-11-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/firewallRules/{firewallRuleName}", "2016-11-01"),
								}},
						},
						{
							Display:  "storageAccounts",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts", "2016-11-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{storageAccountName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}", "2016-11-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}", "2016-11-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}", "2016-11-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}", "2016-11-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "containers",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}/containers", "2016-11-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:  "{containerName}",
													Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}/containers/{containerName}", "2016-11-01"),
													Children: []swagger.ResourceType{},
												}},
										}},
								}},
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.DataLakeStore/operations", "2016-11-01"),
		},
		{
			Display:  "accounts",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DataLakeStore/accounts", "2016-11-01"),
		},
		{
			Display:  "capability",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DataLakeStore/locations/{location}/capability", "2016-11-01"),
		},
		{
			Display:  "usages",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DataLakeStore/locations/{location}/usages", "2016-11-01"),
		},
		{
			Display:  "accounts",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts", "2016-11-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{accountName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}", "2016-11-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}", "2016-11-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}", "2016-11-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}", "2016-11-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "firewallRules",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/firewallRules", "2016-11-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{firewallRuleName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/firewallRules/{firewallRuleName}", "2016-11-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/firewallRules/{firewallRuleName}", "2016-11-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/firewallRules/{firewallRuleName}", "2016-11-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/firewallRules/{firewallRuleName}", "2016-11-01"),
								}},
						},
						{
							Display:  "trustedIdProviders",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/trustedIdProviders", "2016-11-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{trustedIdProviderName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/trustedIdProviders/{trustedIdProviderName}", "2016-11-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/trustedIdProviders/{trustedIdProviderName}", "2016-11-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/trustedIdProviders/{trustedIdProviderName}", "2016-11-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/trustedIdProviders/{trustedIdProviderName}", "2016-11-01"),
								}},
						},
						{
							Display:  "virtualNetworkRules",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/virtualNetworkRules", "2016-11-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{virtualNetworkRuleName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/virtualNetworkRules/{virtualNetworkRuleName}", "2016-11-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/virtualNetworkRules/{virtualNetworkRuleName}", "2016-11-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/virtualNetworkRules/{virtualNetworkRuleName}", "2016-11-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/virtualNetworkRules/{virtualNetworkRuleName}", "2016-11-01"),
								}},
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.DataMigration/operations", "2018-04-19"),
		},
		{
			Display:  "usages",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DataMigration/locations/{location}/usages", "2018-04-19"),
		},
		{
			Display:  "services",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DataMigration/services", "2018-04-19"),
		},
		{
			Display:  "skus",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DataMigration/skus", "2018-04-19"),
		},
		{
			Display:  "services",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services", "2018-04-19"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{serviceName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}", "2018-04-19"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}", "2018-04-19"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}", "2018-04-19"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}", "2018-04-19"),
					Children: []swagger.ResourceType{
						{
							Display:  "projects",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects", "2018-04-19"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{projectName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}", "2018-04-19"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}", "2018-04-19"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}", "2018-04-19"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}", "2018-04-19"),
									Children: []swagger.ResourceType{
										{
											Display:  "tasks",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/tasks", "2018-04-19"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{taskName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/tasks/{taskName}", "2018-04-19"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/tasks/{taskName}", "2018-04-19"),
													PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/tasks/{taskName}", "2018-04-19"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/tasks/{taskName}", "2018-04-19"),
													Children:       []swagger.ResourceType{},
												}},
										}},
								}},
						},
						{
							Display:  "skus",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/skus", "2018-04-19"),
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Devices/operations", "2019-03-22"),
		},
		{
			Display:  "provisioningServices",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Devices/provisioningServices", "2018-01-22"),
		},
		{
			Display:  "provisioningServices",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices", "2018-01-22"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{provisioningServiceName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}", "2018-01-22"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}", "2018-01-22"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}", "2018-01-22"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}", "2018-01-22"),
					Children: []swagger.ResourceType{
						{
							Display:  "certificates",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}/certificates", "2018-01-22"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{certificateName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}/certificates/{certificateName}", "2018-01-22"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}/certificates/{certificateName}", "2018-01-22"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}/certificates/{certificateName}", "2018-01-22"),
									Children:       []swagger.ResourceType{},
								}},
						},
						{
							Display:  "skus",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}/skus", "2018-01-22"),
						}},
					SubResources: []swagger.ResourceType{
						{
							Display:  "{operationId}",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}/operationresults/{operationId}", "2018-01-22"),
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.DevSpaces/operations", "2019-04-01"),
		},
		{
			Display:  "controllers",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DevSpaces/controllers", "2019-04-01"),
		},
		{
			Display:  "controllers",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevSpaces/controllers", "2019-04-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{name}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevSpaces/controllers/{name}", "2019-04-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevSpaces/controllers/{name}", "2019-04-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevSpaces/controllers/{name}", "2019-04-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevSpaces/controllers/{name}", "2019-04-01"),
					Children:       []swagger.ResourceType{},
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.DevTestLab/operations", "2018-09-15"),
		},
		{
			Display:  "labs",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DevTestLab/labs", "2018-09-15"),
		},
		{
			Display:  "{name}",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DevTestLab/locations/{locationName}/operations/{name}", "2018-09-15"),
		},
		{
			Display:  "schedules",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DevTestLab/schedules", "2018-09-15"),
		},
		{
			Display:  "labs",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs", "2018-09-15"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{name}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{name}", "2018-09-15"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{name}", "2018-09-15"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{name}", "2018-09-15"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{name}", "2018-09-15"),
					Children:       []swagger.ResourceType{},
				},
				{
					Display:  "artifactsources",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources", "2018-09-15"),
					SubResources: []swagger.ResourceType{
						{
							Display:        "{name}",
							Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{name}", "2018-09-15"),
							DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{name}", "2018-09-15"),
							PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{name}", "2018-09-15"),
							PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{name}", "2018-09-15"),
						},
						{
							Display:  "armtemplates",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{artifactSourceName}/armtemplates", "2018-09-15"),
							SubResources: []swagger.ResourceType{
								{
									Display:  "{name}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{artifactSourceName}/armtemplates/{name}", "2018-09-15"),
								}},
						},
						{
							Display:  "artifacts",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{artifactSourceName}/artifacts", "2018-09-15"),
							SubResources: []swagger.ResourceType{
								{
									Display:  "{name}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{artifactSourceName}/artifacts/{name}", "2018-09-15"),
									Children: []swagger.ResourceType{},
								}},
						}},
				},
				{
					Display:     "{name}",
					Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/costs/{name}", "2018-09-15"),
					PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/costs/{name}", "2018-09-15"),
				},
				{
					Display:  "customimages",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/customimages", "2018-09-15"),
					SubResources: []swagger.ResourceType{
						{
							Display:        "{name}",
							Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/customimages/{name}", "2018-09-15"),
							DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/customimages/{name}", "2018-09-15"),
							PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/customimages/{name}", "2018-09-15"),
							PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/customimages/{name}", "2018-09-15"),
						}},
				},
				{
					Display:  "formulas",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/formulas", "2018-09-15"),
					SubResources: []swagger.ResourceType{
						{
							Display:        "{name}",
							Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/formulas/{name}", "2018-09-15"),
							DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/formulas/{name}", "2018-09-15"),
							PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/formulas/{name}", "2018-09-15"),
							PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/formulas/{name}", "2018-09-15"),
						}},
				},
				{
					Display:  "galleryimages",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/galleryimages", "2018-09-15"),
				},
				{
					Display:  "notificationchannels",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/notificationchannels", "2018-09-15"),
					SubResources: []swagger.ResourceType{
						{
							Display:        "{name}",
							Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/notificationchannels/{name}", "2018-09-15"),
							DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/notificationchannels/{name}", "2018-09-15"),
							PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/notificationchannels/{name}", "2018-09-15"),
							PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/notificationchannels/{name}", "2018-09-15"),
							Children:       []swagger.ResourceType{},
						}},
				},
				{
					Display:  "policies",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/policysets/{policySetName}/policies", "2018-09-15"),
					SubResources: []swagger.ResourceType{
						{
							Display:        "{name}",
							Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/policysets/{policySetName}/policies/{name}", "2018-09-15"),
							DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/policysets/{policySetName}/policies/{name}", "2018-09-15"),
							PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/policysets/{policySetName}/policies/{name}", "2018-09-15"),
							PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/policysets/{policySetName}/policies/{name}", "2018-09-15"),
						}},
				},
				{
					Display:  "schedules",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/schedules", "2018-09-15"),
					SubResources: []swagger.ResourceType{
						{
							Display:        "{name}",
							Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/schedules/{name}", "2018-09-15"),
							DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/schedules/{name}", "2018-09-15"),
							PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/schedules/{name}", "2018-09-15"),
							PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/schedules/{name}", "2018-09-15"),
							Children:       []swagger.ResourceType{},
						}},
				},
				{
					Display:        "{name}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/servicerunners/{name}", "2018-09-15"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/servicerunners/{name}", "2018-09-15"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/servicerunners/{name}", "2018-09-15"),
				},
				{
					Display:  "users",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users", "2018-09-15"),
					SubResources: []swagger.ResourceType{
						{
							Display:        "{name}",
							Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{name}", "2018-09-15"),
							DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{name}", "2018-09-15"),
							PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{name}", "2018-09-15"),
							PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{name}", "2018-09-15"),
						},
						{
							Display:  "disks",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/disks", "2018-09-15"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{name}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/disks/{name}", "2018-09-15"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/disks/{name}", "2018-09-15"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/disks/{name}", "2018-09-15"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/disks/{name}", "2018-09-15"),
									Children:       []swagger.ResourceType{},
								}},
						},
						{
							Display:  "environments",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/environments", "2018-09-15"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{name}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/environments/{name}", "2018-09-15"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/environments/{name}", "2018-09-15"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/environments/{name}", "2018-09-15"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/environments/{name}", "2018-09-15"),
								}},
						},
						{
							Display:  "secrets",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/secrets", "2018-09-15"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{name}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/secrets/{name}", "2018-09-15"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/secrets/{name}", "2018-09-15"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/secrets/{name}", "2018-09-15"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/secrets/{name}", "2018-09-15"),
								}},
						},
						{
							Display:  "servicefabrics",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics", "2018-09-15"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{name}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{name}", "2018-09-15"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{name}", "2018-09-15"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{name}", "2018-09-15"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{name}", "2018-09-15"),
									Children:       []swagger.ResourceType{},
								},
								{
									Display:  "schedules",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{serviceFabricName}/schedules", "2018-09-15"),
									SubResources: []swagger.ResourceType{
										{
											Display:        "{name}",
											Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{serviceFabricName}/schedules/{name}", "2018-09-15"),
											DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{serviceFabricName}/schedules/{name}", "2018-09-15"),
											PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{serviceFabricName}/schedules/{name}", "2018-09-15"),
											PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/servicefabrics/{serviceFabricName}/schedules/{name}", "2018-09-15"),
											Children:       []swagger.ResourceType{},
										}},
								}},
						}},
				},
				{
					Display:  "virtualmachines",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines", "2018-09-15"),
					SubResources: []swagger.ResourceType{
						{
							Display:        "{name}",
							Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{name}", "2018-09-15"),
							DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{name}", "2018-09-15"),
							PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{name}", "2018-09-15"),
							PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{name}", "2018-09-15"),
							Children:       []swagger.ResourceType{},
						},
						{
							Display:  "schedules",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{virtualMachineName}/schedules", "2018-09-15"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{name}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{virtualMachineName}/schedules/{name}", "2018-09-15"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{virtualMachineName}/schedules/{name}", "2018-09-15"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{virtualMachineName}/schedules/{name}", "2018-09-15"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualmachines/{virtualMachineName}/schedules/{name}", "2018-09-15"),
									Children:       []swagger.ResourceType{},
								}},
						}},
				},
				{
					Display:  "virtualnetworks",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualnetworks", "2018-09-15"),
					SubResources: []swagger.ResourceType{
						{
							Display:        "{name}",
							Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualnetworks/{name}", "2018-09-15"),
							DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualnetworks/{name}", "2018-09-15"),
							PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualnetworks/{name}", "2018-09-15"),
							PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualnetworks/{name}", "2018-09-15"),
						}},
				}},
		},
		{
			Display:  "schedules",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/schedules", "2018-09-15"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{name}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/schedules/{name}", "2018-09-15"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/schedules/{name}", "2018-09-15"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/schedules/{name}", "2018-09-15"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/schedules/{name}", "2018-09-15"),
					Children:       []swagger.ResourceType{},
				}},
		},
		{
			Display:  "dnszones",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/dnszones", "2018-05-01"),
		},
		{
			Display:  "dnsZones",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones", "2018-05-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{zoneName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}", "2018-05-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}", "2018-05-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}", "2018-05-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}", "2018-05-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "all",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/all", "2018-05-01"),
						},
						{
							Display:  "recordsets",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/recordsets", "2018-05-01"),
						}},
					SubResources: []swagger.ResourceType{
						{
							Display:  "{recordType}",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/{recordType}", "2018-05-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{relativeRecordSetName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/{recordType}/{relativeRecordSetName}", "2018-05-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/{recordType}/{relativeRecordSetName}", "2018-05-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/{recordType}/{relativeRecordSetName}", "2018-05-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/{recordType}/{relativeRecordSetName}", "2018-05-01"),
								}},
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.AAD/operations", "2017-06-01"),
		},
		{
			Display:  "domainServices",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.AAD/domainServices", "2017-06-01"),
		},
		{
			Display:  "domainServices",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices", "2017-06-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{domainServiceName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices/{domainServiceName}", "2017-06-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices/{domainServiceName}", "2017-06-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices/{domainServiceName}", "2017-06-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices/{domainServiceName}", "2017-06-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "replicaSets",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices/{domainServiceName}/replicaSets", "2017-06-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{replicaSetName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices/{domainServiceName}/replicaSets/{replicaSetName}", "2017-06-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices/{domainServiceName}/replicaSets/{replicaSetName}", "2017-06-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices/{domainServiceName}/replicaSets/{replicaSetName}", "2017-06-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AAD/domainServices/{domainServiceName}/replicaSets/{replicaSetName}", "2017-06-01"),
								}},
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.EventGrid/operations", "2019-06-01"),
		},
		{
			Display:  "topicTypes",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.EventGrid/topicTypes", "2019-06-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:  "{topicTypeName}",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.EventGrid/topicTypes/{topicTypeName}", "2019-06-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "eventTypes",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.EventGrid/topicTypes/{topicTypeName}/eventTypes", "2019-06-01"),
						}},
				}},
		},
		{
			Display:  "domains",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.EventGrid/domains", "2019-06-01"),
		},
		{
			Display:  "eventSubscriptions",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.EventGrid/eventSubscriptions", "2019-06-01"),
		},
		{
			Display:  "eventSubscriptions",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.EventGrid/locations/{location}/eventSubscriptions", "2019-06-01"),
		},
		{
			Display:  "eventSubscriptions",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.EventGrid/locations/{location}/topicTypes/{topicTypeName}/eventSubscriptions", "2019-06-01"),
		},
		{
			Display:  "eventSubscriptions",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.EventGrid/topicTypes/{topicTypeName}/eventSubscriptions", "2019-06-01"),
		},
		{
			Display:  "topics",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.EventGrid/topics", "2019-06-01"),
		},
		{
			Display:  "domains",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains", "2019-06-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{domainName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}", "2019-06-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}", "2019-06-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}", "2019-06-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}", "2019-06-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "topics",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/topics", "2019-06-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{domainTopicName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/topics/{domainTopicName}", "2019-06-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/topics/{domainTopicName}", "2019-06-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/topics/{domainTopicName}", "2019-06-01"),
								},
								{
									Display:  "eventSubscriptions",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/topics/{topicName}/providers/Microsoft.EventGrid/eventSubscriptions", "2019-06-01"),
								}},
						}},
				}},
		},
		{
			Display:  "eventSubscriptions",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/eventSubscriptions", "2019-06-01"),
		},
		{
			Display:  "eventSubscriptions",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/locations/{location}/eventSubscriptions", "2019-06-01"),
		},
		{
			Display:  "eventSubscriptions",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/locations/{location}/topicTypes/{topicTypeName}/eventSubscriptions", "2019-06-01"),
		},
		{
			Display:  "eventSubscriptions",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topicTypes/{topicTypeName}/eventSubscriptions", "2019-06-01"),
		},
		{
			Display:  "topics",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics", "2019-06-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{topicName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics/{topicName}", "2019-06-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics/{topicName}", "2019-06-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics/{topicName}", "2019-06-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics/{topicName}", "2019-06-01"),
					Children:       []swagger.ResourceType{},
				}},
		},
		{
			Display:  "eventSubscriptions",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{providerNamespace}/{resourceTypeName}/{resourceName}/providers/Microsoft.EventGrid/eventSubscriptions", "2019-06-01"),
		},
		{
			Display:  "eventTypes",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{providerNamespace}/{resourceTypeName}/{resourceName}/providers/Microsoft.EventGrid/eventTypes", "2019-06-01"),
		},
		{
			Display:        "{eventSubscriptionName}",
			Endpoint:       endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.EventGrid/eventSubscriptions/{eventSubscriptionName}", "2019-06-01"),
			DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.EventGrid/eventSubscriptions/{eventSubscriptionName}", "2019-06-01"),
			PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.EventGrid/eventSubscriptions/{eventSubscriptionName}", "2019-06-01"),
			PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.EventGrid/eventSubscriptions/{eventSubscriptionName}", "2019-06-01"),
			Children:       []swagger.ResourceType{},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.EventHub/operations", "2017-04-01"),
		},
		{
			Display:  "namespaces",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.EventHub/namespaces", "2017-04-01"),
		},
		{
			Display:  "regions",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.EventHub/sku/{sku}/regions", "2017-04-01"),
		},
		{
			Display:  "namespaces",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces", "2017-04-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{namespaceName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}", "2017-04-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}", "2017-04-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}", "2017-04-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}", "2017-04-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "AuthorizationRules",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/AuthorizationRules", "2017-04-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{authorizationRuleName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
									Children:       []swagger.ResourceType{},
								}},
						},
						{
							Display:  "disasterRecoveryConfigs",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs", "2017-04-01"),
							Children: []swagger.ResourceType{},
							SubResources: []swagger.ResourceType{
								{
									Display:        "{alias}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}", "2017-04-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}", "2017-04-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}", "2017-04-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "AuthorizationRules",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}/AuthorizationRules", "2017-04-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:  "{authorizationRuleName}",
													Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
													Children: []swagger.ResourceType{},
												}},
										}},
								}},
						},
						{
							Display:  "eventhubs",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs", "2017-04-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{eventHubName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}", "2017-04-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}", "2017-04-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}", "2017-04-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "authorizationRules",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/authorizationRules", "2017-04-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{authorizationRuleName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
													Children:       []swagger.ResourceType{},
												}},
										},
										{
											Display:  "consumergroups",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/consumergroups", "2017-04-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{consumerGroupName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/consumergroups/{consumerGroupName}", "2017-04-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/consumergroups/{consumerGroupName}", "2017-04-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/consumergroups/{consumerGroupName}", "2017-04-01"),
												}},
										}},
								}},
						},
						{
							Display:  "messagingplan",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/messagingplan", "2017-04-01"),
						},
						{
							Display:  "networkRuleSets",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/networkRuleSets", "2017-04-01"),
							Children: []swagger.ResourceType{
								{
									Display:     "default",
									Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/networkRuleSets/default", "2017-04-01"),
									PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/networkRuleSets/default", "2017-04-01"),
								}},
						}},
				}},
		},
		{
			Display:  "NetworkExperimentProfiles",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/NetworkExperimentProfiles", "2019-11-01"),
		},
		{
			Display:  "NetworkExperimentProfiles",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles", "2019-11-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{profileName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}", "2019-11-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}", "2019-11-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}", "2019-11-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}", "2019-11-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "Experiments",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}/Experiments", "2019-11-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{experimentName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}/Experiments/{experimentName}", "2019-11-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}/Experiments/{experimentName}", "2019-11-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}/Experiments/{experimentName}", "2019-11-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}/Experiments/{experimentName}", "2019-11-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "LatencyScorecard",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}/Experiments/{experimentName}/LatencyScorecard", "2019-11-01"),
										},
										{
											Display:  "Timeseries",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}/Experiments/{experimentName}/Timeseries", "2019-11-01"),
										}},
								}},
						},
						{
							Display:  "PreconfiguredEndpoints",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}/PreconfiguredEndpoints", "2019-11-01"),
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.GuestConfiguration/operations", "2018-11-20"),
		},
		{
			Display:  "applications",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/applications", "2018-06-01-preview"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{applicationName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/applications/{applicationName}", "2018-06-01-preview"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/applications/{applicationName}", "2018-06-01-preview"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/applications/{applicationName}", "2018-06-01-preview"),
				}},
		},
		{
			Display:  "clusters",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.HDInsight/clusters", "2018-06-01-preview"),
		},
		{
			Display:  "clusters",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters", "2018-06-01-preview"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{clusterName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}", "2018-06-01-preview"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}", "2018-06-01-preview"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}", "2018-06-01-preview"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}", "2018-06-01-preview"),
					Children: []swagger.ResourceType{
						{
							Display:        "clustermonitoring",
							Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/clustermonitoring", "2018-06-01-preview"),
							DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/clustermonitoring", "2018-06-01-preview"),
							PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/clustermonitoring", "2018-06-01-preview"),
						},
						{
							Display:      "scriptActions",
							Endpoint:     endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/scriptActions", "2018-06-01-preview"),
							SubResources: []swagger.ResourceType{},
						},
						{
							Display:  "scriptExecutionHistory",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/scriptExecutionHistory", "2018-06-01-preview"),
							SubResources: []swagger.ResourceType{
								{
									Display:  "{scriptExecutionId}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/scriptExecutionHistory/{scriptExecutionId}", "2018-06-01-preview"),
									Children: []swagger.ResourceType{},
								}},
						}},
					SubResources: []swagger.ResourceType{
						{
							Display:        "{extensionName}",
							Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/{extensionName}", "2018-06-01-preview"),
							DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/{extensionName}", "2018-06-01-preview"),
							PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/extensions/{extensionName}", "2018-06-01-preview"),
						}},
				}},
		},
		{
			Display:  "billingSpecs",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.HDInsight/locations/{location}/billingSpecs", "2018-06-01-preview"),
		},
		{
			Display:  "capabilities",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.HDInsight/locations/{location}/capabilities", "2018-06-01-preview"),
		},
		{
			Display:  "usages",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.HDInsight/locations/{location}/usages", "2018-06-01-preview"),
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.HDInsight/operations", "2018-06-01-preview"),
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.HealthcareApis/operations", "2019-09-16"),
		},
		{
			Display:  "{operationResultId}",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.HealthcareApis/locations/{locationName}/operationresults/{operationResultId}", "2019-09-16"),
		},
		{
			Display:  "services",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.HealthcareApis/services", "2019-09-16"),
		},
		{
			Display:  "services",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/services", "2019-09-16"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{resourceName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/services/{resourceName}", "2019-09-16"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/services/{resourceName}", "2019-09-16"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/services/{resourceName}", "2019-09-16"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/services/{resourceName}", "2019-09-16"),
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.HybridData/operations", "2016-06-01"),
		},
		{
			Display:  "dataManagers",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.HybridData/dataManagers", "2016-06-01"),
		},
		{
			Display:  "dataManagers",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers", "2016-06-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{dataManagerName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}", "2016-06-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}", "2016-06-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}", "2016-06-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}", "2016-06-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "dataServices",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices", "2016-06-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:  "{dataServiceName}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}", "2016-06-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "jobDefinitions",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions", "2016-06-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{jobDefinitionName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions/{jobDefinitionName}", "2016-06-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions/{jobDefinitionName}", "2016-06-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions/{jobDefinitionName}", "2016-06-01"),
													Children: []swagger.ResourceType{
														{
															Display:  "jobs",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions/{jobDefinitionName}/jobs", "2016-06-01"),
															SubResources: []swagger.ResourceType{
																{
																	Display:  "{jobId}",
																	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobDefinitions/{jobDefinitionName}/jobs/{jobId}", "2016-06-01"),
																	Children: []swagger.ResourceType{},
																}},
														}},
												}},
										},
										{
											Display:  "jobs",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}/jobs", "2016-06-01"),
										}},
								}},
						},
						{
							Display:  "dataStoreTypes",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataStoreTypes", "2016-06-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:  "{dataStoreTypeName}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataStoreTypes/{dataStoreTypeName}", "2016-06-01"),
								}},
						},
						{
							Display:  "dataStores",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataStores", "2016-06-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{dataStoreName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataStores/{dataStoreName}", "2016-06-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataStores/{dataStoreName}", "2016-06-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataStores/{dataStoreName}", "2016-06-01"),
								}},
						},
						{
							Display:  "jobDefinitions",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/jobDefinitions", "2016-06-01"),
						},
						{
							Display:  "jobs",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/jobs", "2016-06-01"),
						},
						{
							Display:  "publicKeys",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/publicKeys", "2016-06-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:  "{publicKeyName}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/publicKeys/{publicKeyName}", "2016-06-01"),
								}},
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.IoTCentral/operations", "2018-09-01"),
		},
		{
			Display:  "IoTApps",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.IoTCentral/IoTApps", "2018-09-01"),
		},
		{
			Display:  "IoTApps",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTCentral/IoTApps", "2018-09-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{resourceName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTCentral/IoTApps/{resourceName}", "2018-09-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTCentral/IoTApps/{resourceName}", "2018-09-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTCentral/IoTApps/{resourceName}", "2018-09-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTCentral/IoTApps/{resourceName}", "2018-09-01"),
				}},
		},
		{
			Display:  "IotHubs",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Devices/IotHubs", "2019-03-22"),
		},
		{
			Display:  "usages",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Devices/usages", "2019-03-22"),
		},
		{
			Display:  "IotHubs",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs", "2019-03-22"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{resourceName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}", "2019-03-22"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}", "2019-03-22"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}", "2019-03-22"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}", "2019-03-22"),
					Children: []swagger.ResourceType{
						{
							Display:  "IotHubStats",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/IotHubStats", "2019-03-22"),
						},
						{
							Display:  "certificates",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/certificates", "2019-03-22"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{certificateName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/certificates/{certificateName}", "2019-03-22"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/certificates/{certificateName}", "2019-03-22"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/certificates/{certificateName}", "2019-03-22"),
									Children:       []swagger.ResourceType{},
								}},
						},
						{
							Display:  "jobs",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/jobs", "2019-03-22"),
							SubResources: []swagger.ResourceType{
								{
									Display:  "{jobId}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/jobs/{jobId}", "2019-03-22"),
								}},
						},
						{
							Display:  "quotaMetrics",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/quotaMetrics", "2019-03-22"),
						},
						{
							Display:  "skus",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/skus", "2019-03-22"),
						}},
					SubResources: []swagger.ResourceType{
						{
							Display:  "ConsumerGroups",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/eventHubEndpoints/{eventHubEndpointName}/ConsumerGroups", "2019-03-22"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{name}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/eventHubEndpoints/{eventHubEndpointName}/ConsumerGroups/{name}", "2019-03-22"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/eventHubEndpoints/{eventHubEndpointName}/ConsumerGroups/{name}", "2019-03-22"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{resourceName}/eventHubEndpoints/{eventHubEndpointName}/ConsumerGroups/{name}", "2019-03-22"),
								}},
						}},
				},
				{
					Display:  "routingEndpointsHealth",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/IotHubs/{iotHubName}/routingEndpointsHealth", "2019-03-22"),
				}},
		},
		{
			Display:  "deletedVaults",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.KeyVault/deletedVaults", "2018-02-14"),
		},
		{
			Display:  "{vaultName}",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.KeyVault/locations/{location}/deletedVaults/{vaultName}", "2018-02-14"),
			Children: []swagger.ResourceType{},
		},
		{
			Display:  "vaults",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.KeyVault/vaults", "2018-02-14"),
		},
		{
			Display:  "vaults",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults", "2018-02-14"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{vaultName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}", "2018-02-14"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}", "2018-02-14"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}", "2018-02-14"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}", "2018-02-14"),
					Children: []swagger.ResourceType{
						{
							Display:  "secrets",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/secrets", "2018-02-14"),
							SubResources: []swagger.ResourceType{
								{
									Display:       "{secretName}",
									Endpoint:      endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/secrets/{secretName}", "2018-02-14"),
									PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/secrets/{secretName}", "2018-02-14"),
									PutEndpoint:   endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/secrets/{secretName}", "2018-02-14"),
								}},
						}},
					SubResources: []swagger.ResourceType{},
				}},
		},
		{
			Display:  "resources",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resources", "2019-08-01"),
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.KeyVault/operations", "2018-02-14"),
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.LabServices/operations", "2018-10-15"),
		},
		{
			Display:  "labaccounts",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.LabServices/labaccounts", "2018-10-15"),
		},
		{
			Display:  "{operationName}",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.LabServices/locations/{locationName}/operations/{operationName}", "2018-10-15"),
		},
		{
			Display:  "labaccounts",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts", "2018-10-15"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{labAccountName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}", "2018-10-15"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}", "2018-10-15"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}", "2018-10-15"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}", "2018-10-15"),
					Children: []swagger.ResourceType{
						{
							Display:  "galleryimages",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/galleryimages", "2018-10-15"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{galleryImageName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/galleryimages/{galleryImageName}", "2018-10-15"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/galleryimages/{galleryImageName}", "2018-10-15"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/galleryimages/{galleryImageName}", "2018-10-15"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/galleryimages/{galleryImageName}", "2018-10-15"),
								}},
						},
						{
							Display:  "labs",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs", "2018-10-15"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{labName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}", "2018-10-15"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}", "2018-10-15"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}", "2018-10-15"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}", "2018-10-15"),
									Children: []swagger.ResourceType{
										{
											Display:  "environmentsettings",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings", "2018-10-15"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{environmentSettingName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}", "2018-10-15"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}", "2018-10-15"),
													PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}", "2018-10-15"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}", "2018-10-15"),
													Children: []swagger.ResourceType{
														{
															Display:  "environments",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}/environments", "2018-10-15"),
															SubResources: []swagger.ResourceType{
																{
																	Display:        "{environmentName}",
																	Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}/environments/{environmentName}", "2018-10-15"),
																	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}/environments/{environmentName}", "2018-10-15"),
																	PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}/environments/{environmentName}", "2018-10-15"),
																	PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/environmentsettings/{environmentSettingName}/environments/{environmentName}", "2018-10-15"),
																	Children:       []swagger.ResourceType{},
																}},
														}},
												}},
										},
										{
											Display:  "users",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/users", "2018-10-15"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{userName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/users/{userName}", "2018-10-15"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/users/{userName}", "2018-10-15"),
													PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/users/{userName}", "2018-10-15"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/labs/{labName}/users/{userName}", "2018-10-15"),
												}},
										}},
								}},
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Logic/operations", "2019-05-01"),
		},
		{
			Display:  "integrationAccounts",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Logic/integrationAccounts", "2019-05-01"),
		},
		{
			Display:  "integrationServiceEnvironments",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Logic/integrationServiceEnvironments", "2019-05-01"),
		},
		{
			Display:  "workflows",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Logic/workflows", "2019-05-01"),
		},
		{
			Display:  "integrationAccounts",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts", "2019-05-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{integrationAccountName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}", "2019-05-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}", "2019-05-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}", "2019-05-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}", "2019-05-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "agreements",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/agreements", "2019-05-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{agreementName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/agreements/{agreementName}", "2019-05-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/agreements/{agreementName}", "2019-05-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/agreements/{agreementName}", "2019-05-01"),
									Children:       []swagger.ResourceType{},
								}},
						},
						{
							Display:  "assemblies",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/assemblies", "2019-05-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{assemblyArtifactName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/assemblies/{assemblyArtifactName}", "2019-05-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/assemblies/{assemblyArtifactName}", "2019-05-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/assemblies/{assemblyArtifactName}", "2019-05-01"),
									Children:       []swagger.ResourceType{},
								}},
						},
						{
							Display:  "batchConfigurations",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/batchConfigurations", "2019-05-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{batchConfigurationName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/batchConfigurations/{batchConfigurationName}", "2019-05-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/batchConfigurations/{batchConfigurationName}", "2019-05-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/batchConfigurations/{batchConfigurationName}", "2019-05-01"),
								}},
						},
						{
							Display:  "certificates",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/certificates", "2019-05-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{certificateName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/certificates/{certificateName}", "2019-05-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/certificates/{certificateName}", "2019-05-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/certificates/{certificateName}", "2019-05-01"),
								}},
						},
						{
							Display:  "maps",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/maps", "2019-05-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{mapName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/maps/{mapName}", "2019-05-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/maps/{mapName}", "2019-05-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/maps/{mapName}", "2019-05-01"),
									Children:       []swagger.ResourceType{},
								}},
						},
						{
							Display:  "partners",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/partners", "2019-05-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{partnerName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/partners/{partnerName}", "2019-05-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/partners/{partnerName}", "2019-05-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/partners/{partnerName}", "2019-05-01"),
									Children:       []swagger.ResourceType{},
								}},
						},
						{
							Display:  "schemas",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/schemas", "2019-05-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{schemaName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/schemas/{schemaName}", "2019-05-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/schemas/{schemaName}", "2019-05-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/schemas/{schemaName}", "2019-05-01"),
									Children:       []swagger.ResourceType{},
								}},
						},
						{
							Display:  "sessions",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/sessions", "2019-05-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{sessionName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/sessions/{sessionName}", "2019-05-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/sessions/{sessionName}", "2019-05-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/sessions/{sessionName}", "2019-05-01"),
								}},
						}},
				}},
		},
		{
			Display:  "integrationServiceEnvironments",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments", "2019-05-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{integrationServiceEnvironmentName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}", "2019-05-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}", "2019-05-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}", "2019-05-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}", "2019-05-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "network",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}/health/network", "2019-05-01"),
						},
						{
							Display:  "managedApis",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}/managedApis", "2019-05-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{apiName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}/managedApis/{apiName}", "2019-05-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}/managedApis/{apiName}", "2019-05-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}/managedApis/{apiName}", "2019-05-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "apiOperations",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}/managedApis/{apiName}/apiOperations", "2019-05-01"),
										}},
								}},
						},
						{
							Display:  "skus",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}/skus", "2019-05-01"),
						}},
				}},
		},
		{
			Display:  "workflows",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows", "2019-05-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{workflowName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}", "2019-05-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}", "2019-05-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}", "2019-05-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}", "2019-05-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "runs",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs", "2019-05-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:  "{runName}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}", "2019-05-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "actions",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions", "2019-05-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:  "{actionName}",
													Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}", "2019-05-01"),
													Children: []swagger.ResourceType{
														{
															Display:  "repetitions",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}/repetitions", "2019-05-01"),
															SubResources: []swagger.ResourceType{
																{
																	Display:  "{repetitionName}",
																	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}/repetitions/{repetitionName}", "2019-05-01"),
																	Children: []swagger.ResourceType{
																		{
																			Display:  "requestHistories",
																			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}/repetitions/{repetitionName}/requestHistories", "2019-05-01"),
																			SubResources: []swagger.ResourceType{
																				{
																					Display:  "{requestHistoryName}",
																					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}/repetitions/{repetitionName}/requestHistories/{requestHistoryName}", "2019-05-01"),
																				}},
																		}},
																}},
														},
														{
															Display:  "requestHistories",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}/requestHistories", "2019-05-01"),
															SubResources: []swagger.ResourceType{
																{
																	Display:  "{requestHistoryName}",
																	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}/requestHistories/{requestHistoryName}", "2019-05-01"),
																}},
														},
														{
															Display:  "scopeRepetitions",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}/scopeRepetitions", "2019-05-01"),
															SubResources: []swagger.ResourceType{
																{
																	Display:  "{repetitionName}",
																	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}/scopeRepetitions/{repetitionName}", "2019-05-01"),
																}},
														}},
												}},
										}},
									SubResources: []swagger.ResourceType{
										{
											Display:  "{operationId}",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/operations/{operationId}", "2019-05-01"),
										}},
								}},
						},
						{
							Display:  "triggers",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/triggers", "2019-05-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:  "{triggerName}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/triggers/{triggerName}", "2019-05-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "histories",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/triggers/{triggerName}/histories", "2019-05-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:  "{historyName}",
													Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/triggers/{triggerName}/histories/{historyName}", "2019-05-01"),
													Children: []swagger.ResourceType{},
												}},
										},
										{
											Display:  "json",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/triggers/{triggerName}/schemas/json", "2019-05-01"),
										}},
								}},
						},
						{
							Display:  "versions",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/versions", "2019-05-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:      "{versionId}",
									Endpoint:     endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/versions/{versionId}", "2019-05-01"),
									SubResources: []swagger.ResourceType{},
								}},
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.MachineLearning/operations", "2017-01-01"),
		},
		{
			Display:  "webServices",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.MachineLearning/webServices", "2017-01-01"),
		},
		{
			Display:  "webServices",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearning/webServices", "2017-01-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{webServiceName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearning/webServices/{webServiceName}", "2017-01-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearning/webServices/{webServiceName}", "2017-01-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearning/webServices/{webServiceName}", "2017-01-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearning/webServices/{webServiceName}", "2017-01-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "listKeys",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearning/webServices/{webServiceName}/listKeys", "2017-01-01"),
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.MachineLearningServices/operations", "2019-06-01"),
		},
		{
			Display:  "Quotas",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.MachineLearningServices/locations/{location}/Quotas", "2019-06-01"),
		},
		{
			Display:  "usages",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.MachineLearningServices/locations/{location}/usages", "2019-06-01"),
		},
		{
			Display:  "vmSizes",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.MachineLearningServices/locations/{location}/vmSizes", "2019-06-01"),
		},
		{
			Display:  "workspaces",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.MachineLearningServices/workspaces", "2019-06-01"),
		},
		{
			Display:  "workspaces",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces", "2019-06-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{workspaceName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}", "2019-06-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}", "2019-06-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}", "2019-06-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}", "2019-06-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "computes",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes", "2019-06-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{computeName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}", "2019-06-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}", "2019-06-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}", "2019-06-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/computes/{computeName}", "2019-06-01"),
									Children:       []swagger.ResourceType{},
								}},
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ManagedServices/operations", "2019-06-01"),
		},
		{
			Display:  "registrationAssignments",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.ManagedServices/registrationAssignments", "2019-06-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{registrationAssignmentId}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.ManagedServices/registrationAssignments/{registrationAssignmentId}", "2019-06-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.ManagedServices/registrationAssignments/{registrationAssignmentId}", "2019-06-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.ManagedServices/registrationAssignments/{registrationAssignmentId}", "2019-06-01"),
				}},
		},
		{
			Display:  "registrationDefinitions",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.ManagedServices/registrationDefinitions", "2019-06-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{registrationDefinitionId}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.ManagedServices/registrationDefinitions/{registrationDefinitionId}", "2019-06-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.ManagedServices/registrationDefinitions/{registrationDefinitionId}", "2019-06-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.ManagedServices/registrationDefinitions/{registrationDefinitionId}", "2019-06-01"),
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Maps/operations", "2018-05-01"),
		},
		{
			Display:  "accounts",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Maps/accounts", "2018-05-01"),
		},
		{
			Display:  "accounts",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Maps/accounts", "2018-05-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{accountName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Maps/accounts/{accountName}", "2018-05-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Maps/accounts/{accountName}", "2018-05-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Maps/accounts/{accountName}", "2018-05-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Maps/accounts/{accountName}", "2018-05-01"),
					Children:       []swagger.ResourceType{},
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.DBforMariaDB/operations", "2018-06-01"),
		},
		{
			Display:  "performanceTiers",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DBforMariaDB/locations/{locationName}/performanceTiers", "2018-06-01"),
		},
		{
			Display:  "servers",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DBforMariaDB/servers", "2018-06-01"),
		},
		{
			Display:  "servers",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers", "2018-06-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{serverName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}", "2018-06-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}", "2018-06-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}", "2018-06-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}", "2018-06-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "configurations",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/configurations", "2018-06-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:     "{configurationName}",
									Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/configurations/{configurationName}", "2018-06-01"),
									PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/configurations/{configurationName}", "2018-06-01"),
								}},
						},
						{
							Display:  "databases",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/databases", "2018-06-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{databaseName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/databases/{databaseName}", "2018-06-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/databases/{databaseName}", "2018-06-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/databases/{databaseName}", "2018-06-01"),
								}},
						},
						{
							Display:  "firewallRules",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/firewallRules", "2018-06-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{firewallRuleName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/firewallRules/{firewallRuleName}", "2018-06-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/firewallRules/{firewallRuleName}", "2018-06-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/firewallRules/{firewallRuleName}", "2018-06-01"),
								}},
						},
						{
							Display:  "logFiles",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/logFiles", "2018-06-01"),
						},
						{
							Display:  "replicas",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/replicas", "2018-06-01"),
						},
						{
							Display:  "virtualNetworkRules",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/virtualNetworkRules", "2018-06-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{virtualNetworkRuleName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/virtualNetworkRules/{virtualNetworkRuleName}", "2018-06-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/virtualNetworkRules/{virtualNetworkRuleName}", "2018-06-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/virtualNetworkRules/{virtualNetworkRuleName}", "2018-06-01"),
								}},
						}},
					SubResources: []swagger.ResourceType{
						{
							Display:     "{securityAlertPolicyName}",
							Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/securityAlertPolicies/{securityAlertPolicyName}", "2018-06-01"),
							PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/securityAlertPolicies/{securityAlertPolicyName}", "2018-06-01"),
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.MarketplaceOrdering/operations", "2015-06-01"),
		},
		{
			Display:  "agreements",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.MarketplaceOrdering/agreements", "2015-06-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:  "{planId}",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.MarketplaceOrdering/agreements/{publisherId}/offers/{offerId}/plans/{planId}", "2015-06-01"),
					Children: []swagger.ResourceType{},
				}},
		},
		{
			Display:     "current",
			Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.MarketplaceOrdering/offerTypes/{offerType}/publishers/{publisherId}/offers/{offerId}/plans/{planId}/agreements/current", "2015-06-01"),
			PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.MarketplaceOrdering/offerTypes/{offerType}/publishers/{publisherId}/offers/{offerId}/plans/{planId}/agreements/current", "2015-06-01"),
		},
		{
			Display:  "accountFilters",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/accountFilters", "2018-07-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{filterName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/accountFilters/{filterName}", "2018-07-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/accountFilters/{filterName}", "2018-07-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/accountFilters/{filterName}", "2018-07-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/accountFilters/{filterName}", "2018-07-01"),
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Media/operations", "2018-07-01"),
		},
		{
			Display:  "mediaservices",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Media/mediaservices", "2018-07-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:  "{accountName}",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Media/mediaservices/{accountName}", "2018-07-01"),
				}},
		},
		{
			Display:  "mediaservices",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices", "2018-07-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{accountName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}", "2018-07-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}", "2018-07-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}", "2018-07-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}", "2018-07-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "liveEvents",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents", "2018-07-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{liveEventName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}", "2018-07-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}", "2018-07-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}", "2018-07-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}", "2018-07-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "liveOutputs",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}/liveOutputs", "2018-07-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{liveOutputName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}/liveOutputs/{liveOutputName}", "2018-07-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}/liveOutputs/{liveOutputName}", "2018-07-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}/liveOutputs/{liveOutputName}", "2018-07-01"),
												}},
										}},
								}},
						},
						{
							Display:  "streamingEndpoints",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/streamingEndpoints", "2018-07-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{streamingEndpointName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/streamingEndpoints/{streamingEndpointName}", "2018-07-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/streamingEndpoints/{streamingEndpointName}", "2018-07-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/streamingEndpoints/{streamingEndpointName}", "2018-07-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/streamingEndpoints/{streamingEndpointName}", "2018-07-01"),
									Children:       []swagger.ResourceType{},
								}},
						}},
				}},
		},
		{
			Display:  "assets",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets", "2018-07-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{assetName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}", "2018-07-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}", "2018-07-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}", "2018-07-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}", "2018-07-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "assetFilters",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}/assetFilters", "2018-07-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{filterName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}/assetFilters/{filterName}", "2018-07-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}/assetFilters/{filterName}", "2018-07-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}/assetFilters/{filterName}", "2018-07-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}/assetFilters/{filterName}", "2018-07-01"),
								}},
						}},
				}},
		},
		{
			Display:  "contentKeyPolicies",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/contentKeyPolicies", "2018-07-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{contentKeyPolicyName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/contentKeyPolicies/{contentKeyPolicyName}", "2018-07-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/contentKeyPolicies/{contentKeyPolicyName}", "2018-07-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/contentKeyPolicies/{contentKeyPolicyName}", "2018-07-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/contentKeyPolicies/{contentKeyPolicyName}", "2018-07-01"),
					Children:       []swagger.ResourceType{},
				}},
		},
		{
			Display:  "transforms",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/transforms", "2018-07-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{transformName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/transforms/{transformName}", "2018-07-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/transforms/{transformName}", "2018-07-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/transforms/{transformName}", "2018-07-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/transforms/{transformName}", "2018-07-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "jobs",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/transforms/{transformName}/jobs", "2018-07-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{jobName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/transforms/{transformName}/jobs/{jobName}", "2018-07-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/transforms/{transformName}/jobs/{jobName}", "2018-07-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/transforms/{transformName}/jobs/{jobName}", "2018-07-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/transforms/{transformName}/jobs/{jobName}", "2018-07-01"),
									Children:       []swagger.ResourceType{},
								}},
						}},
				}},
		},
		{
			Display:  "streamingLocators",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingLocators", "2018-07-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{streamingLocatorName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingLocators/{streamingLocatorName}", "2018-07-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingLocators/{streamingLocatorName}", "2018-07-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingLocators/{streamingLocatorName}", "2018-07-01"),
					Children:       []swagger.ResourceType{},
				}},
		},
		{
			Display:  "streamingPolicies",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingPolicies", "2018-07-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{streamingPolicyName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingPolicies/{streamingPolicyName}", "2018-07-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingPolicies/{streamingPolicyName}", "2018-07-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/streamingPolicies/{streamingPolicyName}", "2018-07-01"),
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Migrate/operations", "2018-02-02"),
		},
		{
			Display:  "assessmentOptions",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Migrate/locations/{locationName}/assessmentOptions", "2018-02-02"),
		},
		{
			Display:  "projects",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Migrate/projects", "2018-02-02"),
		},
		{
			Display:  "assessments",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}/assessments", "2018-02-02"),
		},
		{
			Display:  "groups",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}/groups", "2018-02-02"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{groupName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}/groups/{groupName}", "2018-02-02"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}/groups/{groupName}", "2018-02-02"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}/groups/{groupName}", "2018-02-02"),
					Children: []swagger.ResourceType{
						{
							Display:  "assessments",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}/groups/{groupName}/assessments", "2018-02-02"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{assessmentName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}/groups/{groupName}/assessments/{assessmentName}", "2018-02-02"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}/groups/{groupName}/assessments/{assessmentName}", "2018-02-02"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}/groups/{groupName}/assessments/{assessmentName}", "2018-02-02"),
									Children: []swagger.ResourceType{
										{
											Display:  "assessedMachines",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}/groups/{groupName}/assessments/{assessmentName}/assessedMachines", "2018-02-02"),
											SubResources: []swagger.ResourceType{
												{
													Display:  "{assessedMachineName}",
													Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}/groups/{groupName}/assessments/{assessmentName}/assessedMachines/{assessedMachineName}", "2018-02-02"),
												}},
										}},
								}},
						}},
				}},
		},
		{
			Display:  "machines",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}/machines", "2018-02-02"),
			SubResources: []swagger.ResourceType{
				{
					Display:  "{machineName}",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}/machines/{machineName}", "2018-02-02"),
				}},
		},
		{
			Display:  "projects",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Migrate/projects", "2018-02-02"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{projectName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}", "2018-02-02"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}", "2018-02-02"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}", "2018-02-02"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}", "2018-02-02"),
					Children:       []swagger.ResourceType{},
				}},
		},
		{
			Display:  "actionGroups",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/microsoft.insights/actionGroups", "2019-06-01"),
		},
		{
			Display:  "actionGroups",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/actionGroups", "2019-06-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{actionGroupName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/actionGroups/{actionGroupName}", "2019-06-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/actionGroups/{actionGroupName}", "2019-06-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/actionGroups/{actionGroupName}", "2019-06-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/actionGroups/{actionGroupName}", "2019-06-01"),
					Children:       []swagger.ResourceType{},
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ManagedIdentity/operations", "2018-11-30"),
		},
		{
			Display:  "userAssignedIdentities",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ManagedIdentity/userAssignedIdentities", "2018-11-30"),
		},
		{
			Display:  "userAssignedIdentities",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities", "2018-11-30"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{resourceName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{resourceName}", "2018-11-30"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{resourceName}", "2018-11-30"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{resourceName}", "2018-11-30"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{resourceName}", "2018-11-30"),
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.DBforMySQL/operations", "2017-12-01"),
		},
		{
			Display:  "performanceTiers",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DBforMySQL/locations/{locationName}/performanceTiers", "2017-12-01"),
		},
		{
			Display:  "servers",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DBforMySQL/servers", "2017-12-01"),
		},
		{
			Display:  "servers",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers", "2017-12-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{serverName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}", "2017-12-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}", "2017-12-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}", "2017-12-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}", "2017-12-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "configurations",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/configurations", "2017-12-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:     "{configurationName}",
									Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/configurations/{configurationName}", "2017-12-01"),
									PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/configurations/{configurationName}", "2017-12-01"),
								}},
						},
						{
							Display:  "databases",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/databases", "2017-12-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{databaseName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/databases/{databaseName}", "2017-12-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/databases/{databaseName}", "2017-12-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/databases/{databaseName}", "2017-12-01"),
								}},
						},
						{
							Display:  "firewallRules",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/firewallRules", "2017-12-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{firewallRuleName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/firewallRules/{firewallRuleName}", "2017-12-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/firewallRules/{firewallRuleName}", "2017-12-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/firewallRules/{firewallRuleName}", "2017-12-01"),
								}},
						},
						{
							Display:  "logFiles",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/logFiles", "2017-12-01"),
						},
						{
							Display:  "replicas",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/replicas", "2017-12-01"),
						},
						{
							Display:  "virtualNetworkRules",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/virtualNetworkRules", "2017-12-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{virtualNetworkRuleName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/virtualNetworkRules/{virtualNetworkRuleName}", "2017-12-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/virtualNetworkRules/{virtualNetworkRuleName}", "2017-12-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/virtualNetworkRules/{virtualNetworkRuleName}", "2017-12-01"),
								}},
						}},
					SubResources: []swagger.ResourceType{
						{
							Display:     "{securityAlertPolicyName}",
							Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/securityAlertPolicies/{securityAlertPolicyName}", "2017-12-01"),
							PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/securityAlertPolicies/{securityAlertPolicyName}", "2017-12-01"),
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.NetApp/operations", "2019-07-01"),
		},
		{
			Display:  "netAppAccounts",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts", "2019-07-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{accountName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}", "2019-07-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}", "2019-07-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}", "2019-07-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}", "2019-07-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "capacityPools",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools", "2019-07-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{poolName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}", "2019-07-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}", "2019-07-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}", "2019-07-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}", "2019-07-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "volumes",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes", "2019-07-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{volumeName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}", "2019-07-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}", "2019-07-01"),
													PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}", "2019-07-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}", "2019-07-01"),
													Children: []swagger.ResourceType{
														{
															Display:  "mountTargets",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/mountTargets", "2019-07-01"),
														},
														{
															Display:  "snapshots",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/snapshots", "2019-07-01"),
															SubResources: []swagger.ResourceType{
																{
																	Display:        "{snapshotName}",
																	Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/snapshots/{snapshotName}", "2019-07-01"),
																	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/snapshots/{snapshotName}", "2019-07-01"),
																	PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/snapshots/{snapshotName}", "2019-07-01"),
																	PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/snapshots/{snapshotName}", "2019-07-01"),
																}},
														}},
												}},
										}},
								}},
						}},
				}},
		},
		{
			Display:  "applicationGatewayAvailableRequestHeaders",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationGatewayAvailableRequestHeaders", "2019-08-01"),
		},
		{
			Display:  "applicationGatewayAvailableResponseHeaders",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationGatewayAvailableResponseHeaders", "2019-08-01"),
		},
		{
			Display:  "applicationGatewayAvailableServerVariables",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationGatewayAvailableServerVariables", "2019-08-01"),
		},
		{
			Display:  "default",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationGatewayAvailableSslOptions/default", "2019-08-01"),
			Children: []swagger.ResourceType{
				{
					Display:  "predefinedPolicies",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationGatewayAvailableSslOptions/default/predefinedPolicies", "2019-08-01"),
					SubResources: []swagger.ResourceType{
						{
							Display:  "{predefinedPolicyName}",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationGatewayAvailableSslOptions/default/predefinedPolicies/{predefinedPolicyName}", "2019-08-01"),
						}},
				}},
		},
		{
			Display:  "applicationGatewayAvailableWafRuleSets",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationGatewayAvailableWafRuleSets", "2019-08-01"),
		},
		{
			Display:  "applicationGateways",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationGateways", "2019-08-01"),
		},
		{
			Display:  "applicationGateways",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{applicationGatewayName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}", "2019-08-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}", "2019-08-01"),
					Children:       []swagger.ResourceType{},
				}},
		},
		{
			Display:  "applicationSecurityGroups",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationSecurityGroups", "2019-08-01"),
		},
		{
			Display:  "applicationSecurityGroups",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{applicationSecurityGroupName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups/{applicationSecurityGroupName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups/{applicationSecurityGroupName}", "2019-08-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups/{applicationSecurityGroupName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups/{applicationSecurityGroupName}", "2019-08-01"),
				}},
		},
		{
			Display:  "availableDelegations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/availableDelegations", "2019-08-01"),
		},
		{
			Display:  "availableDelegations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/locations/{location}/availableDelegations", "2019-08-01"),
		},
		{
			Display:  "availableServiceAliases",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/availableServiceAliases", "2019-08-01"),
		},
		{
			Display:  "availableServiceAliases",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/locations/{location}/availableServiceAliases", "2019-08-01"),
		},
		{
			Display:  "azureFirewalls",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/azureFirewalls", "2019-08-01"),
		},
		{
			Display:  "azureFirewalls",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/azureFirewalls", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{azureFirewallName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/azureFirewalls/{azureFirewallName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/azureFirewalls/{azureFirewallName}", "2019-08-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/azureFirewalls/{azureFirewallName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/azureFirewalls/{azureFirewallName}", "2019-08-01"),
				}},
		},
		{
			Display:  "azureFirewallFqdnTags",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/azureFirewallFqdnTags", "2019-08-01"),
		},
		{
			Display:  "bastionHosts",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/bastionHosts", "2019-08-01"),
		},
		{
			Display:  "bastionHosts",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{bastionHostName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts/{bastionHostName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts/{bastionHostName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts/{bastionHostName}", "2019-08-01"),
				}},
		},
		{
			Display:  "CheckDnsNameAvailability",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/CheckDnsNameAvailability", "2019-08-01"),
		},
		{
			Display:        "{ddosCustomPolicyName}",
			Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosCustomPolicies/{ddosCustomPolicyName}", "2019-08-01"),
			DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosCustomPolicies/{ddosCustomPolicyName}", "2019-08-01"),
			PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosCustomPolicies/{ddosCustomPolicyName}", "2019-08-01"),
			PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosCustomPolicies/{ddosCustomPolicyName}", "2019-08-01"),
		},
		{
			Display:  "ddosProtectionPlans",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/ddosProtectionPlans", "2019-08-01"),
		},
		{
			Display:  "ddosProtectionPlans",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosProtectionPlans", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{ddosProtectionPlanName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosProtectionPlans/{ddosProtectionPlanName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosProtectionPlans/{ddosProtectionPlanName}", "2019-08-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosProtectionPlans/{ddosProtectionPlanName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosProtectionPlans/{ddosProtectionPlanName}", "2019-08-01"),
				}},
		},
		{
			Display:  "virtualNetworkAvailableEndpointServices",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/virtualNetworkAvailableEndpointServices", "2019-08-01"),
		},
		{
			Display:  "expressRouteCircuits",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteCircuits", "2019-08-01"),
		},
		{
			Display:  "expressRouteServiceProviders",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteServiceProviders", "2019-08-01"),
		},
		{
			Display:  "expressRouteCircuits",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{circuitName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}", "2019-08-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}", "2019-08-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "authorizations",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/authorizations", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{authorizationName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/authorizations/{authorizationName}", "2019-08-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/authorizations/{authorizationName}", "2019-08-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/authorizations/{authorizationName}", "2019-08-01"),
								}},
						},
						{
							Display:  "peerings",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{peeringName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}", "2019-08-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}", "2019-08-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}", "2019-08-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "connections",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/connections", "2019-08-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{connectionName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/connections/{connectionName}", "2019-08-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/connections/{connectionName}", "2019-08-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/connections/{connectionName}", "2019-08-01"),
												}},
										},
										{
											Display:  "peerConnections",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/peerConnections", "2019-08-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:  "{connectionName}",
													Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/peerConnections/{connectionName}", "2019-08-01"),
												}},
										},
										{
											Display:  "stats",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/stats", "2019-08-01"),
										}},
									SubResources: []swagger.ResourceType{},
								}},
						},
						{
							Display:  "stats",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/stats", "2019-08-01"),
						}},
				}},
		},
		{
			Display:  "expressRouteCrossConnections",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteCrossConnections", "2019-08-01"),
		},
		{
			Display:  "expressRouteCrossConnections",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:       "{crossConnectionName}",
					Endpoint:      endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}", "2019-08-01"),
					PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}", "2019-08-01"),
					PutEndpoint:   endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}", "2019-08-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "peerings",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{peeringName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}", "2019-08-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}", "2019-08-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}", "2019-08-01"),
									SubResources:   []swagger.ResourceType{},
								}},
						}},
				}},
		},
		{
			Display:  "expressRouteGateways",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteGateways", "2019-08-01"),
		},
		{
			Display:  "expressRouteGateways",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{expressRouteGatewayName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}", "2019-08-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "expressRouteConnections",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}/expressRouteConnections", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{connectionName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}/expressRouteConnections/{connectionName}", "2019-08-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}/expressRouteConnections/{connectionName}", "2019-08-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}/expressRouteConnections/{connectionName}", "2019-08-01"),
								}},
						}},
				}},
		},
		{
			Display:  "ExpressRoutePorts",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/ExpressRoutePorts", "2019-08-01"),
		},
		{
			Display:  "ExpressRoutePortsLocations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/ExpressRoutePortsLocations", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:  "{locationName}",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/ExpressRoutePortsLocations/{locationName}", "2019-08-01"),
				}},
		},
		{
			Display:  "ExpressRoutePorts",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{expressRoutePortName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}", "2019-08-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}", "2019-08-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "links",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}/links", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:  "{linkName}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}/links/{linkName}", "2019-08-01"),
								}},
						}},
				}},
		},
		{
			Display:  "firewallPolicies",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/firewallPolicies", "2019-08-01"),
		},
		{
			Display:  "firewallPolicies",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{firewallPolicyName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}", "2019-08-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}", "2019-08-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "ruleGroups",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleGroups", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{ruleGroupName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleGroups/{ruleGroupName}", "2019-08-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleGroups/{ruleGroupName}", "2019-08-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleGroups/{ruleGroupName}", "2019-08-01"),
								}},
						}},
				}},
		},
		{
			Display:  "loadBalancers",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/loadBalancers", "2019-08-01"),
		},
		{
			Display:  "loadBalancers",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{loadBalancerName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}", "2019-08-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}", "2019-08-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "backendAddressPools",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/backendAddressPools", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:  "{backendAddressPoolName}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/backendAddressPools/{backendAddressPoolName}", "2019-08-01"),
								}},
						},
						{
							Display:  "frontendIPConfigurations",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/frontendIPConfigurations", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:  "{frontendIPConfigurationName}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/frontendIPConfigurations/{frontendIPConfigurationName}", "2019-08-01"),
								}},
						},
						{
							Display:  "inboundNatRules",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/inboundNatRules", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{inboundNatRuleName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/inboundNatRules/{inboundNatRuleName}", "2019-08-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/inboundNatRules/{inboundNatRuleName}", "2019-08-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/inboundNatRules/{inboundNatRuleName}", "2019-08-01"),
								}},
						},
						{
							Display:  "loadBalancingRules",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/loadBalancingRules", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:  "{loadBalancingRuleName}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/loadBalancingRules/{loadBalancingRuleName}", "2019-08-01"),
								}},
						},
						{
							Display:  "networkInterfaces",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/networkInterfaces", "2019-08-01"),
						},
						{
							Display:  "outboundRules",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/outboundRules", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:  "{outboundRuleName}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/outboundRules/{outboundRuleName}", "2019-08-01"),
								}},
						},
						{
							Display:  "probes",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/probes", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:  "{probeName}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/probes/{probeName}", "2019-08-01"),
								}},
						}},
				}},
		},
		{
			Display:  "natGateways",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/natGateways", "2019-08-01"),
		},
		{
			Display:  "natGateways",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{natGatewayName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways/{natGatewayName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways/{natGatewayName}", "2019-08-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways/{natGatewayName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways/{natGatewayName}", "2019-08-01"),
				}},
		},
		{
			Display:  "networkInterfaces",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkInterfaces", "2019-08-01"),
		},
		{
			Display:  "networkInterfaces",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{networkInterfaceName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}", "2019-08-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}", "2019-08-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "ipConfigurations",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/ipConfigurations", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:  "{ipConfigurationName}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/ipConfigurations/{ipConfigurationName}", "2019-08-01"),
								}},
						},
						{
							Display:  "loadBalancers",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/loadBalancers", "2019-08-01"),
						},
						{
							Display:  "tapConfigurations",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/tapConfigurations", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{tapConfigurationName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/tapConfigurations/{tapConfigurationName}", "2019-08-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/tapConfigurations/{tapConfigurationName}", "2019-08-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/tapConfigurations/{tapConfigurationName}", "2019-08-01"),
								}},
						}},
				}},
		},
		{
			Display:  "networkProfiles",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkProfiles", "2019-08-01"),
		},
		{
			Display:  "networkProfiles",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkProfiles", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{networkProfileName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkProfiles/{networkProfileName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkProfiles/{networkProfileName}", "2019-08-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkProfiles/{networkProfileName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkProfiles/{networkProfileName}", "2019-08-01"),
				}},
		},
		{
			Display:  "networkSecurityGroups",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkSecurityGroups", "2019-08-01"),
		},
		{
			Display:  "networkSecurityGroups",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{networkSecurityGroupName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}", "2019-08-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}", "2019-08-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "defaultSecurityRules",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/defaultSecurityRules", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:  "{defaultSecurityRuleName}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/defaultSecurityRules/{defaultSecurityRuleName}", "2019-08-01"),
								}},
						},
						{
							Display:  "securityRules",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{securityRuleName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules/{securityRuleName}", "2019-08-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules/{securityRuleName}", "2019-08-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules/{securityRuleName}", "2019-08-01"),
								}},
						}},
				}},
		},
		{
			Display:  "networkWatchers",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkWatchers", "2019-08-01"),
		},
		{
			Display:  "networkWatchers",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{networkWatcherName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}", "2019-08-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}", "2019-08-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "packetCaptures",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{packetCaptureName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}", "2019-08-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}", "2019-08-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}", "2019-08-01"),
									Children:       []swagger.ResourceType{},
								}},
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Network/operations", "2019-08-01"),
		},
		{
			Display:  "availablePrivateEndpointTypes",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/availablePrivateEndpointTypes", "2019-08-01"),
		},
		{
			Display:  "privateEndpoints",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/privateEndpoints", "2019-08-01"),
		},
		{
			Display:  "availablePrivateEndpointTypes",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/locations/{location}/availablePrivateEndpointTypes", "2019-08-01"),
		},
		{
			Display:  "privateEndpoints",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{privateEndpointName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}", "2019-08-01"),
				}},
		},
		{
			Display:  "autoApprovedPrivateLinkServices",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/autoApprovedPrivateLinkServices", "2019-08-01"),
		},
		{
			Display:  "privateLinkServices",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/privateLinkServices", "2019-08-01"),
		},
		{
			Display:  "autoApprovedPrivateLinkServices",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/locations/{location}/autoApprovedPrivateLinkServices", "2019-08-01"),
		},
		{
			Display:  "privateLinkServices",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateLinkServices", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{serviceName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateLinkServices/{serviceName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateLinkServices/{serviceName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateLinkServices/{serviceName}", "2019-08-01"),
					SubResources:   []swagger.ResourceType{},
				}},
		},
		{
			Display:  "publicIPAddresses",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/publicIPAddresses", "2019-08-01"),
		},
		{
			Display:  "publicIPAddresses",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{publicIpAddressName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}", "2019-08-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}", "2019-08-01"),
				}},
		},
		{
			Display:  "publicIPPrefixes",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/publicIPPrefixes", "2019-08-01"),
		},
		{
			Display:  "publicIPPrefixes",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{publicIpPrefixName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIpPrefixName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIpPrefixName}", "2019-08-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIpPrefixName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIpPrefixName}", "2019-08-01"),
				}},
		},
		{
			Display:  "routeFilters",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/routeFilters", "2019-08-01"),
		},
		{
			Display:  "routeFilters",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{routeFilterName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}", "2019-08-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}", "2019-08-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "routeFilterRules",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}/routeFilterRules", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{ruleName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}/routeFilterRules/{ruleName}", "2019-08-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}/routeFilterRules/{ruleName}", "2019-08-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}/routeFilterRules/{ruleName}", "2019-08-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}/routeFilterRules/{ruleName}", "2019-08-01"),
								}},
						}},
				}},
		},
		{
			Display:  "routeTables",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/routeTables", "2019-08-01"),
		},
		{
			Display:  "routeTables",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{routeTableName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}", "2019-08-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}", "2019-08-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "routes",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}/routes", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{routeName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}/routes/{routeName}", "2019-08-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}/routes/{routeName}", "2019-08-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}/routes/{routeName}", "2019-08-01"),
								}},
						}},
				}},
		},
		{
			Display:  "bgpServiceCommunities",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/bgpServiceCommunities", "2019-08-01"),
		},
		{
			Display:  "ServiceEndpointPolicies",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/ServiceEndpointPolicies", "2019-08-01"),
		},
		{
			Display:  "serviceEndpointPolicies",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{serviceEndpointPolicyName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}", "2019-08-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}", "2019-08-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "serviceEndpointPolicyDefinitions",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}/serviceEndpointPolicyDefinitions", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{serviceEndpointPolicyDefinitionName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}/serviceEndpointPolicyDefinitions/{serviceEndpointPolicyDefinitionName}", "2019-08-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}/serviceEndpointPolicyDefinitions/{serviceEndpointPolicyDefinitionName}", "2019-08-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}/serviceEndpointPolicyDefinitions/{serviceEndpointPolicyDefinitionName}", "2019-08-01"),
								}},
						}},
				}},
		},
		{
			Display:  "serviceTags",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/serviceTags", "2019-08-01"),
		},
		{
			Display:  "usages",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/usages", "2019-08-01"),
		},
		{
			Display:  "virtualNetworks",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualNetworks", "2019-08-01"),
		},
		{
			Display:  "virtualNetworks",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{virtualNetworkName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}", "2019-08-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}", "2019-08-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "CheckIPAddressAvailability",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/CheckIPAddressAvailability", "2019-08-01"),
						},
						{
							Display:  "subnets",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{subnetName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}", "2019-08-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}", "2019-08-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}", "2019-08-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "ResourceNavigationLinks",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}/ResourceNavigationLinks", "2019-08-01"),
										},
										{
											Display:  "ServiceAssociationLinks",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}/ServiceAssociationLinks", "2019-08-01"),
										}},
								}},
						},
						{
							Display:  "usages",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/usages", "2019-08-01"),
						},
						{
							Display:  "virtualNetworkPeerings",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/virtualNetworkPeerings", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{virtualNetworkPeeringName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/virtualNetworkPeerings/{virtualNetworkPeeringName}", "2019-08-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/virtualNetworkPeerings/{virtualNetworkPeeringName}", "2019-08-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/virtualNetworkPeerings/{virtualNetworkPeeringName}", "2019-08-01"),
								}},
						}},
				}},
		},
		{
			Display:  "connections",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{virtualNetworkGatewayConnectionName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}", "2019-08-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}", "2019-08-01"),
					Children: []swagger.ResourceType{
						{
							Display:     "sharedkey",
							Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}/sharedkey", "2019-08-01"),
							PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}/sharedkey", "2019-08-01"),
							Children:    []swagger.ResourceType{},
						}},
				}},
		},
		{
			Display:  "localNetworkGateways",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{localNetworkGatewayName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways/{localNetworkGatewayName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways/{localNetworkGatewayName}", "2019-08-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways/{localNetworkGatewayName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways/{localNetworkGatewayName}", "2019-08-01"),
				}},
		},
		{
			Display:  "virtualNetworkGateways",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{virtualNetworkGatewayName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}", "2019-08-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}", "2019-08-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "connections",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}/connections", "2019-08-01"),
						}},
				}},
		},
		{
			Display:  "virtualNetworkTaps",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualNetworkTaps", "2019-08-01"),
		},
		{
			Display:  "virtualNetworkTaps",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{tapName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps/{tapName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps/{tapName}", "2019-08-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps/{tapName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkTaps/{tapName}", "2019-08-01"),
				}},
		},
		{
			Display:  "virtualRouters",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualRouters", "2019-08-01"),
		},
		{
			Display:  "virtualRouters",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualRouters", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{virtualRouterName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualRouters/{virtualRouterName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualRouters/{virtualRouterName}", "2019-08-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualRouters/{virtualRouterName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualRouters/{virtualRouterName}", "2019-08-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "peerings",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualRouters/{virtualRouterName}/peerings", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{peeringName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualRouters/{virtualRouterName}/peerings/{peeringName}", "2019-08-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualRouters/{virtualRouterName}/peerings/{peeringName}", "2019-08-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualRouters/{virtualRouterName}/peerings/{peeringName}", "2019-08-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualRouters/{virtualRouterName}/peerings/{peeringName}", "2019-08-01"),
								}},
						}},
				}},
		},
		{
			Display:  "p2svpnGateways",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/p2svpnGateways", "2019-08-01"),
		},
		{
			Display:  "virtualHubs",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualHubs", "2019-08-01"),
		},
		{
			Display:  "virtualWans",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualWans", "2019-08-01"),
		},
		{
			Display:  "vpnGateways",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/vpnGateways", "2019-08-01"),
		},
		{
			Display:  "vpnServerConfigurations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/vpnServerConfigurations", "2019-08-01"),
		},
		{
			Display:  "vpnSites",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/vpnSites", "2019-08-01"),
		},
		{
			Display:  "p2svpnGateways",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{gatewayName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}", "2019-08-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}", "2019-08-01"),
					Children:       []swagger.ResourceType{},
				}},
		},
		{
			Display:  "virtualHubs",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{virtualHubName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}", "2019-08-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}", "2019-08-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "hubVirtualNetworkConnections",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubVirtualNetworkConnections", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:  "{connectionName}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubVirtualNetworkConnections/{connectionName}", "2019-08-01"),
								}},
						}},
				}},
		},
		{
			Display:  "virtualWans",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{VirtualWANName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{VirtualWANName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{VirtualWANName}", "2019-08-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{VirtualWANName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{VirtualWANName}", "2019-08-01"),
				},
				{
					Display:  "supportedSecurityProviders",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{virtualWANName}/supportedSecurityProviders", "2019-08-01"),
				}},
		},
		{
			Display:  "vpnGateways",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{gatewayName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}", "2019-08-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}", "2019-08-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "vpnConnections",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{connectionName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}", "2019-08-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}", "2019-08-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}", "2019-08-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "vpnLinkConnections",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}/vpnLinkConnections", "2019-08-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:  "{linkConnectionName}",
													Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}/vpnLinkConnections/{linkConnectionName}", "2019-08-01"),
												}},
										}},
								}},
						}},
				}},
		},
		{
			Display:  "vpnServerConfigurations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{vpnServerConfigurationName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations/{vpnServerConfigurationName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations/{vpnServerConfigurationName}", "2019-08-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations/{vpnServerConfigurationName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations/{vpnServerConfigurationName}", "2019-08-01"),
				}},
		},
		{
			Display:  "vpnSites",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{vpnSiteName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}", "2019-08-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}", "2019-08-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "vpnSiteLinks",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}/vpnSiteLinks", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:  "{vpnSiteLinkName}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}/vpnSiteLinks/{vpnSiteLinkName}", "2019-08-01"),
								}},
						}},
				}},
		},
		{
			Display:  "networkInterfaces",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/networkInterfaces", "2018-10-01"),
		},
		{
			Display:  "networkInterfaces",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces", "2018-10-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:  "{networkInterfaceName}",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces/{networkInterfaceName}", "2018-10-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "ipConfigurations",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces/{networkInterfaceName}/ipConfigurations", "2018-10-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:  "{ipConfigurationName}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces/{networkInterfaceName}/ipConfigurations/{ipConfigurationName}", "2018-10-01"),
								}},
						}},
				}},
		},
		{
			Display:  "ApplicationGatewayWebApplicationFirewallPolicies",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies", "2019-08-01"),
		},
		{
			Display:  "ApplicationGatewayWebApplicationFirewallPolicies",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{policyName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies/{policyName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies/{policyName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies/{policyName}", "2019-08-01"),
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.NotificationHubs/operations", "2017-04-01"),
		},
		{
			Display:  "namespaces",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.NotificationHubs/namespaces", "2017-04-01"),
		},
		{
			Display:  "namespaces",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces", "2017-04-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{namespaceName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}", "2017-04-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}", "2017-04-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}", "2017-04-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}", "2017-04-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "AuthorizationRules",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/AuthorizationRules", "2017-04-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{authorizationRuleName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
									Children:       []swagger.ResourceType{},
								}},
						},
						{
							Display:  "notificationHubs",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/notificationHubs", "2017-04-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{notificationHubName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/notificationHubs/{notificationHubName}", "2017-04-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/notificationHubs/{notificationHubName}", "2017-04-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/notificationHubs/{notificationHubName}", "2017-04-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/notificationHubs/{notificationHubName}", "2017-04-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "AuthorizationRules",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/notificationHubs/{notificationHubName}/AuthorizationRules", "2017-04-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{authorizationRuleName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/notificationHubs/{notificationHubName}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/notificationHubs/{notificationHubName}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/notificationHubs/{notificationHubName}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
													Children:       []swagger.ResourceType{},
												}},
										}},
								}},
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.OperationalInsights/operations", "2015-03-20"),
		},
		{
			Display:  "linkTargets",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.OperationalInsights/linkTargets", "2015-03-20"),
		},
		{
			Display:  "{purgeId}",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/operations/{purgeId}", "2015-03-20"),
		},
		{
			Display:  "savedSearches",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/savedSearches", "2015-03-20"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{savedSearchId}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/savedSearches/{savedSearchId}", "2015-03-20"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/savedSearches/{savedSearchId}", "2015-03-20"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/savedSearches/{savedSearchId}", "2015-03-20"),
				}},
		},
		{
			Display:  "storageInsightConfigs",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/storageInsightConfigs", "2015-03-20"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{storageInsightName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/storageInsightConfigs/{storageInsightName}", "2015-03-20"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/storageInsightConfigs/{storageInsightName}", "2015-03-20"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/storageInsightConfigs/{storageInsightName}", "2015-03-20"),
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.PolicyInsights/operations", "2019-10-01"),
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.DBforPostgreSQL/operations", "2017-12-01"),
		},
		{
			Display:  "performanceTiers",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DBforPostgreSQL/locations/{locationName}/performanceTiers", "2017-12-01"),
		},
		{
			Display:  "servers",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DBforPostgreSQL/servers", "2017-12-01"),
		},
		{
			Display:  "servers",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers", "2017-12-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{serverName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}", "2017-12-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}", "2017-12-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}", "2017-12-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}", "2017-12-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "Replicas",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/Replicas", "2017-12-01"),
						},
						{
							Display:  "configurations",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/configurations", "2017-12-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:     "{configurationName}",
									Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/configurations/{configurationName}", "2017-12-01"),
									PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/configurations/{configurationName}", "2017-12-01"),
								}},
						},
						{
							Display:  "databases",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/databases", "2017-12-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{databaseName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/databases/{databaseName}", "2017-12-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/databases/{databaseName}", "2017-12-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/databases/{databaseName}", "2017-12-01"),
								}},
						},
						{
							Display:  "firewallRules",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/firewallRules", "2017-12-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{firewallRuleName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/firewallRules/{firewallRuleName}", "2017-12-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/firewallRules/{firewallRuleName}", "2017-12-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/firewallRules/{firewallRuleName}", "2017-12-01"),
								}},
						},
						{
							Display:  "logFiles",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/logFiles", "2017-12-01"),
						},
						{
							Display:  "virtualNetworkRules",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/virtualNetworkRules", "2017-12-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{virtualNetworkRuleName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/virtualNetworkRules/{virtualNetworkRuleName}", "2017-12-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/virtualNetworkRules/{virtualNetworkRuleName}", "2017-12-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/virtualNetworkRules/{virtualNetworkRuleName}", "2017-12-01"),
								}},
						}},
					SubResources: []swagger.ResourceType{
						{
							Display:     "{securityAlertPolicyName}",
							Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/securityAlertPolicies/{securityAlertPolicyName}", "2017-12-01"),
							PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/securityAlertPolicies/{securityAlertPolicyName}", "2017-12-01"),
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.PowerBIDedicated/operations", "2017-10-01"),
		},
		{
			Display:  "capacities",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.PowerBIDedicated/capacities", "2017-10-01"),
		},
		{
			Display:  "skus",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.PowerBIDedicated/skus", "2017-10-01"),
		},
		{
			Display:  "capacities",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBIDedicated/capacities", "2017-10-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{dedicatedCapacityName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBIDedicated/capacities/{dedicatedCapacityName}", "2017-10-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBIDedicated/capacities/{dedicatedCapacityName}", "2017-10-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBIDedicated/capacities/{dedicatedCapacityName}", "2017-10-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBIDedicated/capacities/{dedicatedCapacityName}", "2017-10-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "skus",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBIDedicated/capacities/{dedicatedCapacityName}/skus", "2017-10-01"),
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.PowerBI/operations", "2016-01-29"),
		},
		{
			Display:  "workspaceCollections",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.PowerBI/workspaceCollections", "2016-01-29"),
		},
		{
			Display:  "workspaceCollections",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/workspaceCollections", "2016-01-29"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{workspaceCollectionName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/workspaceCollections/{workspaceCollectionName}", "2016-01-29"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/workspaceCollections/{workspaceCollectionName}", "2016-01-29"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/workspaceCollections/{workspaceCollectionName}", "2016-01-29"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/workspaceCollections/{workspaceCollectionName}", "2016-01-29"),
					Children: []swagger.ResourceType{
						{
							Display:  "workspaces",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/workspaceCollections/{workspaceCollectionName}/workspaces", "2016-01-29"),
						}},
				}},
		},
		{
			Display:  "privateDnsZones",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/privateDnsZones", "2018-09-01"),
		},
		{
			Display:  "privateDnsZones",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones", "2018-09-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{privateZoneName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}", "2018-09-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}", "2018-09-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}", "2018-09-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}", "2018-09-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "ALL",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/ALL", "2018-09-01"),
						},
						{
							Display:  "virtualNetworkLinks",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/virtualNetworkLinks", "2018-09-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{virtualNetworkLinkName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/virtualNetworkLinks/{virtualNetworkLinkName}", "2018-09-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/virtualNetworkLinks/{virtualNetworkLinkName}", "2018-09-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/virtualNetworkLinks/{virtualNetworkLinkName}", "2018-09-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/virtualNetworkLinks/{virtualNetworkLinkName}", "2018-09-01"),
								}},
						}},
					SubResources: []swagger.ResourceType{
						{
							Display:  "{recordType}",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/{recordType}", "2018-09-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{relativeRecordSetName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/{recordType}/{relativeRecordSetName}", "2018-09-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/{recordType}/{relativeRecordSetName}", "2018-09-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/{recordType}/{relativeRecordSetName}", "2018-09-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/{recordType}/{relativeRecordSetName}", "2018-09-01"),
								}},
						}},
				}},
		},
		{
			Display:  "replicationUsages",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/replicationUsages", "2016-06-01"),
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.RecoveryServices/operations", "2016-06-01"),
		},
		{
			Display:  "vaults",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.RecoveryServices/vaults", "2016-06-01"),
		},
		{
			Display:  "vaults",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults", "2016-06-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{vaultName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}", "2016-06-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}", "2016-06-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}", "2016-06-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}", "2016-06-01"),
					Children: []swagger.ResourceType{
						{
							Display:       "vaultExtendedInfo",
							Endpoint:      endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/extendedInformation/vaultExtendedInfo", "2016-06-01"),
							PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/extendedInformation/vaultExtendedInfo", "2016-06-01"),
							PutEndpoint:   endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/extendedInformation/vaultExtendedInfo", "2016-06-01"),
						}},
				}},
		},
		{
			Display:  "usages",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/usages", "2016-06-01"),
		},
		{
			Display:        "{protectedItemName}",
			Endpoint:       endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/protectedItems/{protectedItemName}", "2019-05-13"),
			DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/protectedItems/{protectedItemName}", "2019-05-13"),
			PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/protectedItems/{protectedItemName}", "2019-05-13"),
			Children: []swagger.ResourceType{
				{
					Display:  "recoveryPoints",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/protectedItems/{protectedItemName}/recoveryPoints", "2019-05-13"),
					SubResources: []swagger.ResourceType{
						{
							Display:  "{recoveryPointId}",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/protectedItems/{protectedItemName}/recoveryPoints/{recoveryPointId}", "2019-05-13"),
							Children: []swagger.ResourceType{},
						}},
				}},
			SubResources: []swagger.ResourceType{
				{
					Display:  "{operationId}",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/protectedItems/{protectedItemName}/operationResults/{operationId}", "2019-05-13"),
				}},
		},
		{
			Display:  "backupJobs",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupJobs", "2019-05-13"),
			SubResources: []swagger.ResourceType{
				{
					Display:  "{operationId}",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupJobs/operationResults/{operationId}", "2019-05-13"),
				},
				{
					Display:  "{jobName}",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupJobs/{jobName}", "2019-05-13"),
					Children: []swagger.ResourceType{},
					SubResources: []swagger.ResourceType{
						{
							Display:  "{operationId}",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupJobs/{jobName}/operationResults/{operationId}", "2019-05-13"),
						}},
				}},
		},
		{
			Display:  "backupPolicies",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupPolicies", "2019-05-13"),
			SubResources: []swagger.ResourceType{
				{
					Display:     "{policyName}",
					Endpoint:    endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupPolicies/{policyName}", "2019-05-13"),
					PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupPolicies/{policyName}", "2019-05-13"),
					SubResources: []swagger.ResourceType{
						{
							Display:  "{operationId}",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupPolicies/{policyName}/operationResults/{operationId}", "2019-05-13"),
						}},
				}},
		},
		{
			Display:  "backupProtectedItems",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupProtectedItems", "2019-05-13"),
		},
		{
			Display:       "vaultconfig",
			Endpoint:      endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupconfig/vaultconfig", "2019-05-13"),
			PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupconfig/vaultconfig", "2019-05-13"),
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/operations", "2018-07-10"),
		},
		{
			Display:  "replicationAlertSettings",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationAlertSettings", "2018-07-10"),
			SubResources: []swagger.ResourceType{
				{
					Display:     "{alertSettingName}",
					Endpoint:    endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationAlertSettings/{alertSettingName}", "2018-07-10"),
					PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationAlertSettings/{alertSettingName}", "2018-07-10"),
				}},
		},
		{
			Display:  "replicationEvents",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationEvents", "2018-07-10"),
			SubResources: []swagger.ResourceType{
				{
					Display:  "{eventName}",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationEvents/{eventName}", "2018-07-10"),
				}},
		},
		{
			Display:  "replicationFabrics",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics", "2018-07-10"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{fabricName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}", "2018-07-10"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}", "2018-07-10"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}", "2018-07-10"),
					Children: []swagger.ResourceType{
						{
							Display:  "replicationLogicalNetworks",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationLogicalNetworks", "2018-07-10"),
							SubResources: []swagger.ResourceType{
								{
									Display:  "{logicalNetworkName}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationLogicalNetworks/{logicalNetworkName}", "2018-07-10"),
								}},
						},
						{
							Display:  "replicationNetworks",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationNetworks", "2018-07-10"),
							SubResources: []swagger.ResourceType{
								{
									Display:  "{networkName}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationNetworks/{networkName}", "2018-07-10"),
									Children: []swagger.ResourceType{
										{
											Display:  "replicationNetworkMappings",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationNetworks/{networkName}/replicationNetworkMappings", "2018-07-10"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{networkMappingName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationNetworks/{networkName}/replicationNetworkMappings/{networkMappingName}", "2018-07-10"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationNetworks/{networkName}/replicationNetworkMappings/{networkMappingName}", "2018-07-10"),
													PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationNetworks/{networkName}/replicationNetworkMappings/{networkMappingName}", "2018-07-10"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationNetworks/{networkName}/replicationNetworkMappings/{networkMappingName}", "2018-07-10"),
												}},
										}},
								}},
						},
						{
							Display:  "replicationProtectionContainers",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers", "2018-07-10"),
							SubResources: []swagger.ResourceType{
								{
									Display:     "{protectionContainerName}",
									Endpoint:    endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}", "2018-07-10"),
									PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}", "2018-07-10"),
									Children: []swagger.ResourceType{
										{
											Display:  "replicationMigrationItems",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationMigrationItems", "2018-07-10"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{migrationItemName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationMigrationItems/{migrationItemName}", "2018-07-10"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationMigrationItems/{migrationItemName}", "2018-07-10"),
													PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationMigrationItems/{migrationItemName}", "2018-07-10"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationMigrationItems/{migrationItemName}", "2018-07-10"),
													Children: []swagger.ResourceType{
														{
															Display:  "migrationRecoveryPoints",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationMigrationItems/{migrationItemName}/migrationRecoveryPoints", "2018-07-10"),
															SubResources: []swagger.ResourceType{
																{
																	Display:  "{migrationRecoveryPointName}",
																	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationMigrationItems/{migrationItemName}/migrationRecoveryPoints/{migrationRecoveryPointName}", "2018-07-10"),
																}},
														}},
												}},
										},
										{
											Display:  "replicationProtectableItems",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectableItems", "2018-07-10"),
											SubResources: []swagger.ResourceType{
												{
													Display:  "{protectableItemName}",
													Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectableItems/{protectableItemName}", "2018-07-10"),
												}},
										},
										{
											Display:  "replicationProtectedItems",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectedItems", "2018-07-10"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{replicatedProtectedItemName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectedItems/{replicatedProtectedItemName}", "2018-07-10"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectedItems/{replicatedProtectedItemName}", "2018-07-10"),
													PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectedItems/{replicatedProtectedItemName}", "2018-07-10"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectedItems/{replicatedProtectedItemName}", "2018-07-10"),
													Children: []swagger.ResourceType{
														{
															Display:  "recoveryPoints",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectedItems/{replicatedProtectedItemName}/recoveryPoints", "2018-07-10"),
															SubResources: []swagger.ResourceType{
																{
																	Display:  "{recoveryPointName}",
																	Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectedItems/{replicatedProtectedItemName}/recoveryPoints/{recoveryPointName}", "2018-07-10"),
																}},
														},
														{
															Display:  "targetComputeSizes",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectedItems/{replicatedProtectedItemName}/targetComputeSizes", "2018-07-10"),
														}},
												}},
										},
										{
											Display:  "replicationProtectionContainerMappings",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectionContainerMappings", "2018-07-10"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{mappingName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectionContainerMappings/{mappingName}", "2018-07-10"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectionContainerMappings/{mappingName}", "2018-07-10"),
													PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectionContainerMappings/{mappingName}", "2018-07-10"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectionContainerMappings/{mappingName}", "2018-07-10"),
													Children:       []swagger.ResourceType{},
												}},
										}},
								}},
						},
						{
							Display:  "replicationRecoveryServicesProviders",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationRecoveryServicesProviders", "2018-07-10"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{providerName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationRecoveryServicesProviders/{providerName}", "2018-07-10"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationRecoveryServicesProviders/{providerName}", "2018-07-10"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationRecoveryServicesProviders/{providerName}", "2018-07-10"),
									Children:       []swagger.ResourceType{},
								}},
						},
						{
							Display:  "replicationStorageClassifications",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationStorageClassifications", "2018-07-10"),
							SubResources: []swagger.ResourceType{
								{
									Display:  "{storageClassificationName}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationStorageClassifications/{storageClassificationName}", "2018-07-10"),
									Children: []swagger.ResourceType{
										{
											Display:  "replicationStorageClassificationMappings",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationStorageClassifications/{storageClassificationName}/replicationStorageClassificationMappings", "2018-07-10"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{storageClassificationMappingName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationStorageClassifications/{storageClassificationName}/replicationStorageClassificationMappings/{storageClassificationMappingName}", "2018-07-10"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationStorageClassifications/{storageClassificationName}/replicationStorageClassificationMappings/{storageClassificationMappingName}", "2018-07-10"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationStorageClassifications/{storageClassificationName}/replicationStorageClassificationMappings/{storageClassificationMappingName}", "2018-07-10"),
												}},
										}},
								}},
						},
						{
							Display:  "replicationvCenters",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationvCenters", "2018-07-10"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{vCenterName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationvCenters/{vCenterName}", "2018-07-10"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationvCenters/{vCenterName}", "2018-07-10"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationvCenters/{vCenterName}", "2018-07-10"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationvCenters/{vCenterName}", "2018-07-10"),
								}},
						}},
				}},
		},
		{
			Display:  "replicationJobs",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationJobs", "2018-07-10"),
			Children: []swagger.ResourceType{},
			SubResources: []swagger.ResourceType{
				{
					Display:  "{jobName}",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationJobs/{jobName}", "2018-07-10"),
					Children: []swagger.ResourceType{},
				}},
		},
		{
			Display:  "replicationMigrationItems",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationMigrationItems", "2018-07-10"),
		},
		{
			Display:  "replicationNetworkMappings",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationNetworkMappings", "2018-07-10"),
		},
		{
			Display:  "replicationNetworks",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationNetworks", "2018-07-10"),
		},
		{
			Display:  "replicationPolicies",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationPolicies", "2018-07-10"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{policyName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationPolicies/{policyName}", "2018-07-10"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationPolicies/{policyName}", "2018-07-10"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationPolicies/{policyName}", "2018-07-10"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationPolicies/{policyName}", "2018-07-10"),
				}},
		},
		{
			Display:  "replicationProtectedItems",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationProtectedItems", "2018-07-10"),
		},
		{
			Display:  "replicationProtectionContainerMappings",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationProtectionContainerMappings", "2018-07-10"),
		},
		{
			Display:  "replicationProtectionContainers",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationProtectionContainers", "2018-07-10"),
		},
		{
			Display:  "replicationRecoveryPlans",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationRecoveryPlans", "2018-07-10"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{recoveryPlanName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationRecoveryPlans/{recoveryPlanName}", "2018-07-10"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationRecoveryPlans/{recoveryPlanName}", "2018-07-10"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationRecoveryPlans/{recoveryPlanName}", "2018-07-10"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationRecoveryPlans/{recoveryPlanName}", "2018-07-10"),
					Children:       []swagger.ResourceType{},
				}},
		},
		{
			Display:  "replicationRecoveryServicesProviders",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationRecoveryServicesProviders", "2018-07-10"),
		},
		{
			Display:  "replicationStorageClassificationMappings",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationStorageClassificationMappings", "2018-07-10"),
		},
		{
			Display:  "replicationStorageClassifications",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationStorageClassifications", "2018-07-10"),
		},
		{
			Display:  "replicationSupportedOperatingSystems",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationSupportedOperatingSystems", "2018-07-10"),
		},
		{
			Display:  "replicationVaultHealth",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationVaultHealth", "2018-07-10"),
			Children: []swagger.ResourceType{},
		},
		{
			Display:  "replicationVaultSettings",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationVaultSettings", "2018-07-10"),
			SubResources: []swagger.ResourceType{
				{
					Display:     "{vaultSettingName}",
					Endpoint:    endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationVaultSettings/{vaultSettingName}", "2018-07-10"),
					PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationVaultSettings/{vaultSettingName}", "2018-07-10"),
				}},
		},
		{
			Display:  "replicationvCenters",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationvCenters", "2018-07-10"),
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Cache/operations", "2018-03-01"),
		},
		{
			Display:  "Redis",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Cache/Redis", "2018-03-01"),
		},
		{
			Display:  "Redis",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis", "2018-03-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{name}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}", "2018-03-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}", "2018-03-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}", "2018-03-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}", "2018-03-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "linkedServers",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/linkedServers", "2018-03-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{linkedServerName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/linkedServers/{linkedServerName}", "2018-03-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/linkedServers/{linkedServerName}", "2018-03-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/linkedServers/{linkedServerName}", "2018-03-01"),
								}},
						},
						{
							Display:  "listUpgradeNotifications",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/listUpgradeNotifications", "2018-03-01"),
						}},
					SubResources: []swagger.ResourceType{
						{
							Display:        "{default}",
							Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/patchSchedules/{default}", "2018-03-01"),
							DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/patchSchedules/{default}", "2018-03-01"),
							PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/patchSchedules/{default}", "2018-03-01"),
						}},
				},
				{
					Display:  "firewallRules",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{cacheName}/firewallRules", "2018-03-01"),
					SubResources: []swagger.ResourceType{
						{
							Display:        "{ruleName}",
							Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{cacheName}/firewallRules/{ruleName}", "2018-03-01"),
							DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{cacheName}/firewallRules/{ruleName}", "2018-03-01"),
							PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{cacheName}/firewallRules/{ruleName}", "2018-03-01"),
						}},
				},
				{
					Display:  "patchSchedules",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{cacheName}/patchSchedules", "2018-03-01"),
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Relay/operations", "2017-04-01"),
		},
		{
			Display:  "namespaces",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Relay/namespaces", "2017-04-01"),
		},
		{
			Display:  "namespaces",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces", "2017-04-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{namespaceName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}", "2017-04-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}", "2017-04-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}", "2017-04-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}", "2017-04-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "authorizationRules",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/authorizationRules", "2017-04-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{authorizationRuleName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
									Children:       []swagger.ResourceType{},
								}},
						},
						{
							Display:  "hybridConnections",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/hybridConnections", "2017-04-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{hybridConnectionName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/hybridConnections/{hybridConnectionName}", "2017-04-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/hybridConnections/{hybridConnectionName}", "2017-04-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/hybridConnections/{hybridConnectionName}", "2017-04-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "authorizationRules",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/hybridConnections/{hybridConnectionName}/authorizationRules", "2017-04-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{authorizationRuleName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/hybridConnections/{hybridConnectionName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/hybridConnections/{hybridConnectionName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/hybridConnections/{hybridConnectionName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
													Children:       []swagger.ResourceType{},
												}},
										}},
								}},
						},
						{
							Display:  "wcfRelays",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/wcfRelays", "2017-04-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{relayName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/wcfRelays/{relayName}", "2017-04-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/wcfRelays/{relayName}", "2017-04-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/wcfRelays/{relayName}", "2017-04-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "authorizationRules",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/wcfRelays/{relayName}/authorizationRules", "2017-04-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{authorizationRuleName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/wcfRelays/{relayName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/wcfRelays/{relayName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/wcfRelays/{relayName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
													Children:       []swagger.ResourceType{},
												}},
										}},
								}},
						}},
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Capacity/operations", "2017-11-01"),
		},
		{
			Display:  "reservationOrders",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Capacity/reservationOrders", "2017-11-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:  "{reservationOrderId}",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}", "2017-11-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "reservations",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/reservations", "2017-11-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:       "{reservationId}",
									Endpoint:      endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/reservations/{reservationId}", "2017-11-01"),
									PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/reservations/{reservationId}", "2017-11-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "revisions",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/reservations/{reservationId}/revisions", "2017-11-01"),
										}},
								}},
						}},
				}},
		},
		{
			Display:  "appliedReservations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Capacity/appliedReservations", "2017-11-01"),
		},
		{
			Display:  "catalogs",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Capacity/catalogs", "2017-11-01"),
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ResourceGraph/operations", "2019-04-01"),
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ResourceHealth/operations", "2018-07-01"),
		},
		{
			Display:  "availabilityStatuses",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ResourceHealth/availabilityStatuses", "2018-07-01"),
		},
		{
			Display:  "events",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ResourceHealth/events", "2018-07-01"),
		},
		{
			Display:  "availabilityStatuses",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ResourceHealth/availabilityStatuses", "2018-07-01"),
		},
		{
			Display:  "availabilityStatuses",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceUri}/providers/Microsoft.ResourceHealth/availabilityStatuses", "2018-07-01"),
			Children: []swagger.ResourceType{
				{
					Display:  "current",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceUri}/providers/Microsoft.ResourceHealth/availabilityStatuses/current", "2018-07-01"),
				}},
		},
		{
			Display:  "events",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceUri}/providers/Microsoft.ResourceHealth/events", "2018-07-01"),
		},
		{
			Display:  "policyAssignments",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyAssignments", "2019-06-01"),
		},
		{
			Display:  "policyAssignments",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Authorization/policyAssignments", "2019-06-01"),
		},
		{
			Display:  "policyAssignments",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}/providers/Microsoft.Authorization/policyAssignments", "2019-06-01"),
		},
		{
			Display:        "{policyAssignmentId}",
			Endpoint:       endpoints.MustGetEndpointInfoFromURL("/{policyAssignmentId}", "2019-06-01"),
			DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{policyAssignmentId}", "2019-06-01"),
			PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/{policyAssignmentId}", "2019-06-01"),
		},
		{
			Display:        "{policyAssignmentName}",
			Endpoint:       endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Authorization/policyAssignments/{policyAssignmentName}", "2019-06-01"),
			DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Authorization/policyAssignments/{policyAssignmentName}", "2019-06-01"),
			PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Authorization/policyAssignments/{policyAssignmentName}", "2019-06-01"),
		},
		{
			Display:  "policyDefinitions",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Authorization/policyDefinitions", "2019-06-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:  "{policyDefinitionName}",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", "2019-06-01"),
				}},
		},
		{
			Display:  "policyDefinitions",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policyDefinitions", "2019-06-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{policyDefinitionName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", "2019-06-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", "2019-06-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", "2019-06-01"),
				}},
		},
		{
			Display:  "policyDefinitions",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyDefinitions", "2019-06-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{policyDefinitionName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", "2019-06-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", "2019-06-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", "2019-06-01"),
				}},
		},
		{
			Display:  "policySetDefinitions",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Authorization/policySetDefinitions", "2019-06-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:  "{policySetDefinitionName}",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Authorization/policySetDefinitions/{policySetDefinitionName}", "2019-06-01"),
				}},
		},
		{
			Display:  "policySetDefinitions",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policySetDefinitions", "2019-06-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{policySetDefinitionName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policySetDefinitions/{policySetDefinitionName}", "2019-06-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policySetDefinitions/{policySetDefinitionName}", "2019-06-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policySetDefinitions/{policySetDefinitionName}", "2019-06-01"),
				}},
		},
		{
			Display:  "policySetDefinitions",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policySetDefinitions", "2019-06-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{policySetDefinitionName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policySetDefinitions/{policySetDefinitionName}", "2019-06-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policySetDefinitions/{policySetDefinitionName}", "2019-06-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policySetDefinitions/{policySetDefinitionName}", "2019-06-01"),
				}},
		},
		{
			Display:  "operations",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Features/operations", "2015-12-01"),
		},
		{
			Display:  "features",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Features/features", "2015-12-01"),
		},
		{
			Display:  "features",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Features/providers/{resourceProviderNamespace}/features", "2015-12-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:  "{featureName}",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Features/providers/{resourceProviderNamespace}/features/{featureName}", "2015-12-01"),
					Children: []swagger.ResourceType{},
				}},
		},
		{
			Display:  "providers",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers", "2019-08-01"),
			Children: []swagger.ResourceType{
				{
					Display:  "deployments",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Resources/deployments", "2019-08-01"),
					SubResources: []swagger.ResourceType{
						{
							Display:        "{deploymentName}",
							Endpoint:       endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Resources/deployments/{deploymentName}", "2019-08-01"),
							DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Resources/deployments/{deploymentName}", "2019-08-01"),
							PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Resources/deployments/{deploymentName}", "2019-08-01"),
							Children: []swagger.ResourceType{
								{
									Display:  "operations",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Resources/deployments/{deploymentName}/operations", "2019-08-01"),
									SubResources: []swagger.ResourceType{
										{
											Display:  "{operationId}",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Resources/deployments/{deploymentName}/operations/{operationId}", "2019-08-01"),
										}},
								}},
						}},
				},
				{
					Display:  "operations",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Resources/operations", "2019-08-01"),
				},
				{
					Display:  "operations",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Search/operations", "2015-08-19"),
				},
				{
					Display:  "operations",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ServiceBus/operations", "2017-04-01"),
				},
				{
					Display:  "operations",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ServiceFabric/operations", "2019-03-01"),
				},
				{
					Display:  "operations",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.SignalRService/operations", "2018-10-01"),
				},
				{
					Display:  "operations",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.StorSimple/operations", "2017-06-01"),
				},
				{
					Display:  "operations",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Storage/operations", "2019-06-01"),
				},
				{
					Display:  "operations",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.StorageCache/operations", "2019-11-01"),
				},
				{
					Display:  "locations",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ImportExport/locations", "2016-11-01"),
					SubResources: []swagger.ResourceType{
						{
							Display:  "{locationName}",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ImportExport/locations/{locationName}", "2016-11-01"),
						}},
				},
				{
					Display:  "operations",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.ImportExport/operations", "2016-11-01"),
				},
				{
					Display:  "operations",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.StorageSync/operations", "2019-06-01"),
				},
				{
					Display:  "operations",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.StreamAnalytics/operations", "2016-03-01"),
				},
				{
					Display:  "operations",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.TimeSeriesInsights/operations", "2017-11-15"),
				},
				{
					Display:  "default",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Network/trafficManagerGeographicHierarchies/default", "2018-04-01"),
				},
				{
					Display:  "operations",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.VMwareCloudSimple/operations", "2019-04-01"),
				},
				{
					Display:  "operations",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.CertificateRegistration/operations", "2018-02-01"),
				},
				{
					Display:  "operations",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.DomainRegistration/operations", "2018-02-01"),
				},
				{
					Display:  "availableStacks",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Web/availableStacks", "2018-02-01"),
				},
				{
					Display:  "operations",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Web/operations", "2018-02-01"),
				},
				{
					Display:     "web",
					Endpoint:    endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Web/publishingUsers/web", "2018-02-01"),
					PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Web/publishingUsers/web", "2018-02-01"),
				},
				{
					Display:  "sourcecontrols",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Web/sourcecontrols", "2018-02-01"),
					SubResources: []swagger.ResourceType{
						{
							Display:     "{sourceControlType}",
							Endpoint:    endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Web/sourcecontrols/{sourceControlType}", "2018-02-01"),
							PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Web/sourcecontrols/{sourceControlType}", "2018-02-01"),
						}},
				},
				{
					Display:  "operations",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.WindowsIoT/operations", "2019-06-01"),
				}},
			SubResources: []swagger.ResourceType{
				{
					Display:  "deployments",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementGroups/{groupId}/providers/Microsoft.Resources/deployments", "2019-08-01"),
					SubResources: []swagger.ResourceType{
						{
							Display:        "{deploymentName}",
							Endpoint:       endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementGroups/{groupId}/providers/Microsoft.Resources/deployments/{deploymentName}", "2019-08-01"),
							DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementGroups/{groupId}/providers/Microsoft.Resources/deployments/{deploymentName}", "2019-08-01"),
							PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementGroups/{groupId}/providers/Microsoft.Resources/deployments/{deploymentName}", "2019-08-01"),
							Children: []swagger.ResourceType{
								{
									Display:  "operations",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementGroups/{groupId}/providers/Microsoft.Resources/deployments/{deploymentName}/operations", "2019-08-01"),
									SubResources: []swagger.ResourceType{
										{
											Display:  "{operationId}",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/Microsoft.Management/managementGroups/{groupId}/providers/Microsoft.Resources/deployments/{deploymentName}/operations/{operationId}", "2019-08-01"),
										}},
								}},
						}},
				},
				{
					Display:  "{resourceProviderNamespace}",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/providers/{resourceProviderNamespace}", "2019-08-01"),
				}},
		},
		{
			Display:  "providers",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers", "2019-08-01"),
			Children: []swagger.ResourceType{
				{
					Display:  "deployments",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Resources/deployments", "2019-08-01"),
					SubResources: []swagger.ResourceType{
						{
							Display:        "{deploymentName}",
							Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Resources/deployments/{deploymentName}", "2019-08-01"),
							DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Resources/deployments/{deploymentName}", "2019-08-01"),
							PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Resources/deployments/{deploymentName}", "2019-08-01"),
							Children: []swagger.ResourceType{
								{
									Display:  "operations",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Resources/deployments/{deploymentName}/operations", "2019-08-01"),
									SubResources: []swagger.ResourceType{
										{
											Display:  "{operationId}",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Resources/deployments/{deploymentName}/operations/{operationId}", "2019-08-01"),
										}},
								}},
						}},
				},
				{
					Display:  "applications",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Solutions/applications", "2018-06-01"),
				},
				{
					Display:  "jobCollections",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Scheduler/jobCollections", "2016-03-01"),
				},
				{
					Display:  "searchServices",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Search/searchServices", "2015-08-19"),
				},
				{
					Display:  "iotSecuritySolutions",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Security/iotSecuritySolutions", "2019-08-01"),
				},
				{
					Display:  "operations",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.SerialConsole/operations", "2018-05-01"),
				},
				{
					Display:  "namespaces",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceBus/namespaces", "2017-04-01"),
				},
				{
					Display:  "premiumMessagingRegions",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceBus/premiumMessagingRegions", "2017-04-01"),
				},
				{
					Display:  "clusters",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/clusters", "2019-03-01"),
				},
				{
					Display:  "SignalR",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.SignalRService/SignalR", "2018-10-01"),
				},
				{
					Display:  "managers",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.StorSimple/managers", "2017-06-01"),
				},
				{
					Display:  "skus",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Storage/skus", "2019-06-01"),
				},
				{
					Display:  "storageAccounts",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Storage/storageAccounts", "2019-06-01"),
				},
				{
					Display:  "caches",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.StorageCache/caches", "2019-11-01"),
				},
				{
					Display:  "skus",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.StorageCache/skus", "2019-11-01"),
				},
				{
					Display:  "usageModels",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.StorageCache/usageModels", "2019-11-01"),
				},
				{
					Display:  "jobs",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ImportExport/jobs", "2016-11-01"),
				},
				{
					Display:  "storageSyncServices",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.StorageSync/storageSyncServices", "2019-06-01"),
				},
				{
					Display:  "streamingjobs",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.StreamAnalytics/streamingjobs", "2016-03-01"),
				},
				{
					Display:  "environments",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.TimeSeriesInsights/environments", "2017-11-15"),
				},
				{
					Display:        "default",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/trafficManagerUserMetricsKeys/default", "2018-04-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/trafficManagerUserMetricsKeys/default", "2018-04-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/trafficManagerUserMetricsKeys/default", "2018-04-01"),
				},
				{
					Display:  "trafficmanagerprofiles",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Network/trafficmanagerprofiles", "2018-04-01"),
				},
				{
					Display:  "dedicatedCloudNodes",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes", "2019-04-01"),
				},
				{
					Display:  "dedicatedCloudServices",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudServices", "2019-04-01"),
				},
				{
					Display:  "virtualMachines",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/virtualMachines", "2019-04-01"),
				},
				{
					Display:  "certificateOrders",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.CertificateRegistration/certificateOrders", "2018-02-01"),
				},
				{
					Display:  "domains",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DomainRegistration/domains", "2018-02-01"),
				},
				{
					Display:  "topLevelDomains",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DomainRegistration/topLevelDomains", "2018-02-01"),
					SubResources: []swagger.ResourceType{
						{
							Display:  "{name}",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.DomainRegistration/topLevelDomains/{name}", "2018-02-01"),
							Children: []swagger.ResourceType{},
						}},
				},
				{
					Display:  "hostingEnvironments",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/hostingEnvironments", "2018-02-01"),
				},
				{
					Display:  "serverfarms",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/serverfarms", "2018-02-01"),
				},
				{
					Display:  "certificates",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/certificates", "2018-02-01"),
				},
				{
					Display:  "deletedSites",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/deletedSites", "2018-02-01"),
				},
				{
					Display:  "availableStacks",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/availableStacks", "2018-02-01"),
				},
				{
					Display:      "recommendations",
					Endpoint:     endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/recommendations", "2018-02-01"),
					Children:     []swagger.ResourceType{},
					SubResources: []swagger.ResourceType{},
				},
				{
					Display:  "resourceHealthMetadata",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/resourceHealthMetadata", "2018-02-01"),
				},
				{
					Display:  "billingMeters",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/billingMeters", "2018-02-01"),
				},
				{
					Display:  "deploymentLocations",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/deploymentLocations", "2018-02-01"),
				},
				{
					Display:  "geoRegions",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/geoRegions", "2018-02-01"),
				},
				{
					Display:  "premieraddonoffers",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/premieraddonoffers", "2018-02-01"),
				},
				{
					Display:  "skus",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/skus", "2018-02-01"),
				},
				{
					Display:  "sites",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/sites", "2018-02-01"),
				},
				{
					Display:  "deviceServices",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.WindowsIoT/deviceServices", "2019-06-01"),
				}},
			SubResources: []swagger.ResourceType{
				{
					Display:  "{resourceProviderNamespace}",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/{resourceProviderNamespace}", "2019-08-01"),
					Children: []swagger.ResourceType{},
				},
				{
					Display:  "{default}",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.SerialConsole/consoleServices/{default}", "2018-05-01"),
					Children: []swagger.ResourceType{},
				},
				{
					Display:  "regions",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceBus/sku/{sku}/regions", "2017-04-01"),
				},
				{
					Display:  "clusterVersions",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/clusterVersions", "2019-03-01"),
					SubResources: []swagger.ResourceType{
						{
							Display:  "{clusterVersion}",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/clusterVersions/{clusterVersion}", "2019-03-01"),
						}},
				},
				{
					Display:  "clusterVersions",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/environments/{environment}/clusterVersions", "2019-03-01"),
					SubResources: []swagger.ResourceType{
						{
							Display:  "{clusterVersion}",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/environments/{environment}/clusterVersions/{clusterVersion}", "2019-03-01"),
						}},
				},
				{
					Display:  "usages",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.SignalRService/locations/{location}/usages", "2018-10-01"),
				},
				{
					Display:  "capabilities",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Sql/locations/{locationName}/capabilities", "2015-05-01"),
				},
				{
					Display:  "usages",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Sql/locations/{locationName}/usages", "2015-05-01"),
					SubResources: []swagger.ResourceType{
						{
							Display:  "{usageName}",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Sql/locations/{locationName}/usages/{usageName}", "2015-05-01"),
						}},
				},
				{
					Display:  "usages",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Storage/locations/{location}/usages", "2019-06-01"),
				},
				{
					Display:  "quotas",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.StreamAnalytics/locations/{location}/quotas", "2016-03-01"),
				},
				{
					Display:  "availabilities",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/availabilities", "2019-04-01"),
				},
				{
					Display:  "{operationId}",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/operationResults/{operationId}", "2019-04-01"),
				},
				{
					Display:  "privateClouds",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds", "2019-04-01"),
					SubResources: []swagger.ResourceType{
						{
							Display:  "{pcName}",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}", "2019-04-01"),
							Children: []swagger.ResourceType{
								{
									Display:  "resourcePools",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/resourcePools", "2019-04-01"),
									SubResources: []swagger.ResourceType{
										{
											Display:  "{resourcePoolName}",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/resourcePools/{resourcePoolName}", "2019-04-01"),
										}},
								},
								{
									Display:  "virtualMachineTemplates",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/virtualMachineTemplates", "2019-04-01"),
									SubResources: []swagger.ResourceType{
										{
											Display:  "{virtualMachineTemplateName}",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/virtualMachineTemplates/{virtualMachineTemplateName}", "2019-04-01"),
										}},
								},
								{
									Display:  "virtualNetworks",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/virtualNetworks", "2019-04-01"),
									SubResources: []swagger.ResourceType{
										{
											Display:  "{virtualNetworkName}",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/virtualNetworks/{virtualNetworkName}", "2019-04-01"),
										}},
								}},
						}},
				},
				{
					Display:  "usages",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/usages", "2019-04-01"),
				},
				{
					Display:  "deletedSites",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/locations/{location}/deletedSites", "2018-02-01"),
					SubResources: []swagger.ResourceType{
						{
							Display:  "{deletedSiteId}",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/providers/Microsoft.Web/locations/{location}/deletedSites/{deletedSiteId}", "2018-02-01"),
						}},
				}},
		},
		{
			Display:  "resources",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/resources", "2019-08-01"),
		},
		{
			Display:  "resourcegroups",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{resourceGroupName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}", "2019-08-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}", "2019-08-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "deployments",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Resources/deployments", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{deploymentName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Resources/deployments/{deploymentName}", "2019-08-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Resources/deployments/{deploymentName}", "2019-08-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Resources/deployments/{deploymentName}", "2019-08-01"),
									Children:       []swagger.ResourceType{},
								}},
						},
						{
							Display:  "clusters",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters", "2019-03-01"),
						},
						{
							Display:  "caches",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches", "2019-11-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{cacheName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}", "2019-11-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}", "2019-11-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}", "2019-11-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}", "2019-11-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "storageTargets",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}/storageTargets", "2019-11-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{storageTargetName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}/storageTargets/{storageTargetName}", "2019-11-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}/storageTargets/{storageTargetName}", "2019-11-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}/storageTargets/{storageTargetName}", "2019-11-01"),
												}},
										}},
								}},
						},
						{
							Display:  "streamingjobs",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs", "2016-03-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{jobName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}", "2016-03-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}", "2016-03-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}", "2016-03-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}", "2016-03-01"),
									Children:       []swagger.ResourceType{},
									SubResources: []swagger.ResourceType{
										{
											Display:       "{transformationName}",
											Endpoint:      endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/transformations/{transformationName}", "2016-03-01"),
											PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/transformations/{transformationName}", "2016-03-01"),
											PutEndpoint:   endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/transformations/{transformationName}", "2016-03-01"),
										}},
								}},
						}},
					SubResources: []swagger.ResourceType{
						{
							Display:  "operations",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/deployments/{deploymentName}/operations", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:  "{operationId}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/deployments/{deploymentName}/operations/{operationId}", "2019-08-01"),
								}},
						},
						{
							Display:        "{resourceName}",
							Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}", "2019-08-01"),
							DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}", "2019-08-01"),
							PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}", "2019-08-01"),
							PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}", "2019-08-01"),
						},
						{
							Display:  "functions",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/functions", "2016-03-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{functionName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/functions/{functionName}", "2016-03-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/functions/{functionName}", "2016-03-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/functions/{functionName}", "2016-03-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/functions/{functionName}", "2016-03-01"),
									Children:       []swagger.ResourceType{},
								}},
						},
						{
							Display:  "inputs",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/inputs", "2016-03-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{inputName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/inputs/{inputName}", "2016-03-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/inputs/{inputName}", "2016-03-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/inputs/{inputName}", "2016-03-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/inputs/{inputName}", "2016-03-01"),
									Children:       []swagger.ResourceType{},
								}},
						},
						{
							Display:  "outputs",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/outputs", "2016-03-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{outputName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/outputs/{outputName}", "2016-03-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/outputs/{outputName}", "2016-03-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/outputs/{outputName}", "2016-03-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/outputs/{outputName}", "2016-03-01"),
									Children:       []swagger.ResourceType{},
								}},
						}},
				}},
		},
		{
			Display:      "tagNames",
			Endpoint:     endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/tagNames", "2019-08-01"),
			SubResources: []swagger.ResourceType{},
		},
		{
			Display:        "{resourceId}",
			Endpoint:       endpoints.MustGetEndpointInfoFromURL("/{resourceId}", "2019-08-01"),
			DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceId}", "2019-08-01"),
			PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/{resourceId}", "2019-08-01"),
			PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/{resourceId}", "2019-08-01"),
			Children: []swagger.ResourceType{
				{
					Display:  "deviceSecurityGroups",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceId}/providers/Microsoft.Security/deviceSecurityGroups", "2019-08-01"),
					SubResources: []swagger.ResourceType{
						{
							Display:        "{deviceSecurityGroupName}",
							Endpoint:       endpoints.MustGetEndpointInfoFromURL("/{resourceId}/providers/Microsoft.Security/deviceSecurityGroups/{deviceSecurityGroupName}", "2019-08-01"),
							DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{resourceId}/providers/Microsoft.Security/deviceSecurityGroups/{deviceSecurityGroupName}", "2019-08-01"),
							PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/{resourceId}/providers/Microsoft.Security/deviceSecurityGroups/{deviceSecurityGroupName}", "2019-08-01"),
						}},
				}},
		},
		{
			Display:  "deployments",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Resources/deployments", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{deploymentName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Resources/deployments/{deploymentName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Resources/deployments/{deploymentName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Resources/deployments/{deploymentName}", "2019-08-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "operations",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Resources/deployments/{deploymentName}/operations", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:  "{operationId}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/{scope}/providers/Microsoft.Resources/deployments/{deploymentName}/operations/{operationId}", "2019-08-01"),
								}},
						}},
				}},
		},
		{
			Display:  "applicationDefinitions",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions", "2018-06-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{applicationDefinitionName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions/{applicationDefinitionName}", "2018-06-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions/{applicationDefinitionName}", "2018-06-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions/{applicationDefinitionName}", "2018-06-01"),
				}},
		},
		{
			Display:  "applications",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applications", "2018-06-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{applicationName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applications/{applicationName}", "2018-06-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applications/{applicationName}", "2018-06-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applications/{applicationName}", "2018-06-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applications/{applicationName}", "2018-06-01"),
				}},
		},
		{
			Display:        "{applicationId}",
			Endpoint:       endpoints.MustGetEndpointInfoFromURL("/{applicationId}", "2018-06-01"),
			DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/{applicationId}", "2018-06-01"),
			PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/{applicationId}", "2018-06-01"),
			PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/{applicationId}", "2018-06-01"),
		},
		{
			Display:  "jobCollections",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections", "2016-03-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{jobCollectionName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}", "2016-03-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}", "2016-03-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}", "2016-03-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}", "2016-03-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "jobs",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs", "2016-03-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{jobName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs/{jobName}", "2016-03-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs/{jobName}", "2016-03-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs/{jobName}", "2016-03-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs/{jobName}", "2016-03-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "history",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/jobs/{jobName}/history", "2016-03-01"),
										}},
								}},
						}},
				}},
		},
		{
			Display:  "searchServices",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices", "2015-08-19"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{searchServiceName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}", "2015-08-19"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}", "2015-08-19"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}", "2015-08-19"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}", "2015-08-19"),
					Children: []swagger.ResourceType{
						{
							Display:  "listQueryKeys",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/listQueryKeys", "2015-08-19"),
						}},
					SubResources: []swagger.ResourceType{},
				}},
		},
		{
			Display:  "analyticsModels",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}/analyticsModels", "2019-08-01"),
			Children: []swagger.ResourceType{
				{
					Display:  "default",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}/analyticsModels/default", "2019-08-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "aggregatedAlerts",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}/analyticsModels/default/aggregatedAlerts", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:  "{aggregatedAlertName}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}/analyticsModels/default/aggregatedAlerts/{aggregatedAlertName}", "2019-08-01"),
									Children: []swagger.ResourceType{},
								}},
						},
						{
							Display:  "aggregatedRecommendations",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}/analyticsModels/default/aggregatedRecommendations", "2019-08-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:  "{aggregatedRecommendationName}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}/analyticsModels/default/aggregatedRecommendations/{aggregatedRecommendationName}", "2019-08-01"),
								}},
						}},
				}},
		},
		{
			Display:  "iotSecuritySolutions",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions", "2019-08-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{solutionName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}", "2019-08-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}", "2019-08-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}", "2019-08-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}", "2019-08-01"),
				}},
		},
		{
			Display:  "namespaces",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces", "2017-04-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{namespaceName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}", "2017-04-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}", "2017-04-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}", "2017-04-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}", "2017-04-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "AuthorizationRules",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/AuthorizationRules", "2017-04-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{authorizationRuleName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
									Children:       []swagger.ResourceType{},
								}},
						},
						{
							Display:  "disasterRecoveryConfigs",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/disasterRecoveryConfigs", "2017-04-01"),
							Children: []swagger.ResourceType{},
							SubResources: []swagger.ResourceType{
								{
									Display:        "{alias}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}", "2017-04-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}", "2017-04-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}", "2017-04-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "AuthorizationRules",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}/AuthorizationRules", "2017-04-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:  "{authorizationRuleName}",
													Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/disasterRecoveryConfigs/{alias}/AuthorizationRules/{authorizationRuleName}", "2017-04-01"),
													Children: []swagger.ResourceType{},
												}},
										}},
								}},
						},
						{
							Display:  "eventhubs",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/eventhubs", "2017-04-01"),
						},
						{
							Display:  "migrationConfigurations",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/migrationConfigurations", "2017-04-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{configName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/migrationConfigurations/{configName}", "2017-04-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/migrationConfigurations/{configName}", "2017-04-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/migrationConfigurations/{configName}", "2017-04-01"),
									Children:       []swagger.ResourceType{},
								}},
						},
						{
							Display:  "networkRuleSets",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/networkRuleSets", "2017-04-01"),
							Children: []swagger.ResourceType{
								{
									Display:     "default",
									Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/networkRuleSets/default", "2017-04-01"),
									PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/networkRuleSets/default", "2017-04-01"),
								}},
						},
						{
							Display:  "queues",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues", "2017-04-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{queueName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}", "2017-04-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}", "2017-04-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}", "2017-04-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "authorizationRules",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}/authorizationRules", "2017-04-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{authorizationRuleName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/queues/{queueName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
													Children:       []swagger.ResourceType{},
												}},
										}},
								}},
						},
						{
							Display:  "topics",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics", "2017-04-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{topicName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}", "2017-04-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}", "2017-04-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}", "2017-04-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "authorizationRules",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/authorizationRules", "2017-04-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{authorizationRuleName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/authorizationRules/{authorizationRuleName}", "2017-04-01"),
													Children:       []swagger.ResourceType{},
												}},
										},
										{
											Display:  "subscriptions",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions", "2017-04-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{subscriptionName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions/{subscriptionName}", "2017-04-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions/{subscriptionName}", "2017-04-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions/{subscriptionName}", "2017-04-01"),
													Children: []swagger.ResourceType{
														{
															Display:  "rules",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions/{subscriptionName}/rules", "2017-04-01"),
															SubResources: []swagger.ResourceType{
																{
																	Display:        "{ruleName}",
																	Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions/{subscriptionName}/rules/{ruleName}", "2017-04-01"),
																	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions/{subscriptionName}/rules/{ruleName}", "2017-04-01"),
																	PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/topics/{topicName}/subscriptions/{subscriptionName}/rules/{ruleName}", "2017-04-01"),
																}},
														}},
												}},
										}},
								}},
						}},
				}},
		},
		{
			Display:  "applicationTypes",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes", "2019-03-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{applicationTypeName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes/{applicationTypeName}", "2019-03-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes/{applicationTypeName}", "2019-03-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes/{applicationTypeName}", "2019-03-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "versions",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes/{applicationTypeName}/versions", "2019-03-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{version}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes/{applicationTypeName}/versions/{version}", "2019-03-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes/{applicationTypeName}/versions/{version}", "2019-03-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes/{applicationTypeName}/versions/{version}", "2019-03-01"),
								}},
						}},
				}},
		},
		{
			Display:  "applications",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applications", "2019-03-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{applicationName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applications/{applicationName}", "2019-03-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applications/{applicationName}", "2019-03-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applications/{applicationName}", "2019-03-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applications/{applicationName}", "2019-03-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "services",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applications/{applicationName}/services", "2019-03-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{serviceName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applications/{applicationName}/services/{serviceName}", "2019-03-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applications/{applicationName}/services/{serviceName}", "2019-03-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applications/{applicationName}/services/{serviceName}", "2019-03-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applications/{applicationName}/services/{serviceName}", "2019-03-01"),
								}},
						}},
				}},
		},
		{
			Display:        "{clusterName}",
			Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}", "2019-03-01"),
			DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}", "2019-03-01"),
			PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}", "2019-03-01"),
			PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}", "2019-03-01"),
		},
		{
			Display:      "SignalR",
			Endpoint:     endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/SignalR", "2018-10-01"),
			SubResources: []swagger.ResourceType{},
		},
		{
			Display:        "{resourceName}",
			Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/signalR/{resourceName}", "2018-10-01"),
			DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/signalR/{resourceName}", "2018-10-01"),
			PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/signalR/{resourceName}", "2018-10-01"),
			PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/signalR/{resourceName}", "2018-10-01"),
			Children:       []swagger.ResourceType{},
		},
		{
			Display:  "managers",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers", "2017-06-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{managerName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}", "2017-06-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}", "2017-06-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}", "2017-06-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}", "2017-06-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "accessControlRecords",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/accessControlRecords", "2017-06-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{accessControlRecordName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/accessControlRecords/{accessControlRecordName}", "2017-06-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/accessControlRecords/{accessControlRecordName}", "2017-06-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/accessControlRecords/{accessControlRecordName}", "2017-06-01"),
								}},
						},
						{
							Display:  "alerts",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/alerts", "2017-06-01"),
						},
						{
							Display:  "backups",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/backups", "2016-10-01"),
						},
						{
							Display:  "devices",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices", "2017-06-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{deviceName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}", "2017-06-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}", "2017-06-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}", "2017-06-01"),
									Children: []swagger.ResourceType{
										{
											Display:     "default",
											Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/alertSettings/default", "2017-06-01"),
											PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/alertSettings/default", "2017-06-01"),
										},
										{
											Display:  "backupScheduleGroups",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupScheduleGroups", "2016-10-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{scheduleGroupName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupScheduleGroups/{scheduleGroupName}", "2016-10-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupScheduleGroups/{scheduleGroupName}", "2016-10-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupScheduleGroups/{scheduleGroupName}", "2016-10-01"),
												}},
										},
										{
											Display:      "backups",
											Endpoint:     endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backups", "2017-06-01"),
											SubResources: []swagger.ResourceType{},
										},
										{
											Display:  "chapSettings",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/chapSettings", "2016-10-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{chapUserName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/chapSettings/{chapUserName}", "2016-10-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/chapSettings/{chapUserName}", "2016-10-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/chapSettings/{chapUserName}", "2016-10-01"),
												}},
										},
										{
											Display:  "disks",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/disks", "2016-10-01"),
										},
										{
											Display:  "failoverTargets",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/failoverTargets", "2016-10-01"),
										},
										{
											Display:  "fileservers",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/fileservers", "2016-10-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{fileServerName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/fileservers/{fileServerName}", "2016-10-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/fileservers/{fileServerName}", "2016-10-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/fileservers/{fileServerName}", "2016-10-01"),
													Children: []swagger.ResourceType{
														{
															Display:  "metrics",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/fileservers/{fileServerName}/metrics", "2016-10-01"),
														},
														{
															Display:  "metricsDefinitions",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/fileservers/{fileServerName}/metricsDefinitions", "2016-10-01"),
														},
														{
															Display:  "shares",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/fileservers/{fileServerName}/shares", "2016-10-01"),
															SubResources: []swagger.ResourceType{
																{
																	Display:        "{shareName}",
																	Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/fileservers/{fileServerName}/shares/{shareName}", "2016-10-01"),
																	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/fileservers/{fileServerName}/shares/{shareName}", "2016-10-01"),
																	PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/fileservers/{fileServerName}/shares/{shareName}", "2016-10-01"),
																	Children: []swagger.ResourceType{
																		{
																			Display:  "metrics",
																			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/fileservers/{fileServerName}/shares/{shareName}/metrics", "2016-10-01"),
																		},
																		{
																			Display:  "metricsDefinitions",
																			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/fileservers/{fileServerName}/shares/{shareName}/metricsDefinitions", "2016-10-01"),
																		}},
																}},
														}},
												}},
										},
										{
											Display:  "iscsiservers",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers", "2016-10-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{iscsiServerName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}", "2016-10-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}", "2016-10-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}", "2016-10-01"),
													Children: []swagger.ResourceType{
														{
															Display:  "disks",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}/disks", "2016-10-01"),
															SubResources: []swagger.ResourceType{
																{
																	Display:        "{diskName}",
																	Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}/disks/{diskName}", "2016-10-01"),
																	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}/disks/{diskName}", "2016-10-01"),
																	PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}/disks/{diskName}", "2016-10-01"),
																	Children: []swagger.ResourceType{
																		{
																			Display:  "metrics",
																			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}/disks/{diskName}/metrics", "2016-10-01"),
																		},
																		{
																			Display:  "metricsDefinitions",
																			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}/disks/{diskName}/metricsDefinitions", "2016-10-01"),
																		}},
																}},
														},
														{
															Display:  "metrics",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}/metrics", "2016-10-01"),
														},
														{
															Display:  "metricsDefinitions",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/iscsiservers/{iscsiServerName}/metricsDefinitions", "2016-10-01"),
														}},
												}},
										},
										{
											Display:  "jobs",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/jobs", "2017-06-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:  "{jobName}",
													Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/jobs/{jobName}", "2017-06-01"),
													Children: []swagger.ResourceType{},
												}},
										},
										{
											Display:  "metrics",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/metrics", "2017-06-01"),
										},
										{
											Display:  "metricsDefinitions",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/metricsDefinitions", "2017-06-01"),
										},
										{
											Display:       "default",
											Endpoint:      endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/networkSettings/default", "2017-06-01"),
											PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/networkSettings/default", "2017-06-01"),
										},
										{
											Display:  "shares",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/shares", "2016-10-01"),
										},
										{
											Display:     "default",
											Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/timeSettings/default", "2017-06-01"),
											PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/timeSettings/default", "2017-06-01"),
										},
										{
											Display:  "default",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/updateSummary/default", "2017-06-01"),
										},
										{
											Display:  "backupPolicies",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies", "2017-06-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{backupPolicyName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}", "2017-06-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}", "2017-06-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}", "2017-06-01"),
													Children: []swagger.ResourceType{
														{
															Display:  "schedules",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}/schedules", "2017-06-01"),
															SubResources: []swagger.ResourceType{
																{
																	Display:        "{backupScheduleName}",
																	Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}/schedules/{backupScheduleName}", "2017-06-01"),
																	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}/schedules/{backupScheduleName}", "2017-06-01"),
																	PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backupPolicies/{backupPolicyName}/schedules/{backupScheduleName}", "2017-06-01"),
																}},
														}},
												}},
										},
										{
											Display:      "hardwareComponentGroups",
											Endpoint:     endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/hardwareComponentGroups", "2017-06-01"),
											SubResources: []swagger.ResourceType{},
										},
										{
											Display:       "default",
											Endpoint:      endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/securitySettings/default", "2017-06-01"),
											PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/securitySettings/default", "2017-06-01"),
											Children:      []swagger.ResourceType{},
										},
										{
											Display:  "volumeContainers",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers", "2017-06-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{volumeContainerName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers/{volumeContainerName}", "2017-06-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers/{volumeContainerName}", "2017-06-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers/{volumeContainerName}", "2017-06-01"),
													Children: []swagger.ResourceType{
														{
															Display:  "metrics",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers/{volumeContainerName}/metrics", "2017-06-01"),
														},
														{
															Display:  "metricsDefinitions",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers/{volumeContainerName}/metricsDefinitions", "2017-06-01"),
														},
														{
															Display:  "volumes",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers/{volumeContainerName}/volumes", "2017-06-01"),
															SubResources: []swagger.ResourceType{
																{
																	Display:        "{volumeName}",
																	Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers/{volumeContainerName}/volumes/{volumeName}", "2017-06-01"),
																	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers/{volumeContainerName}/volumes/{volumeName}", "2017-06-01"),
																	PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers/{volumeContainerName}/volumes/{volumeName}", "2017-06-01"),
																	Children: []swagger.ResourceType{
																		{
																			Display:  "metrics",
																			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers/{volumeContainerName}/volumes/{volumeName}/metrics", "2017-06-01"),
																		},
																		{
																			Display:  "metricsDefinitions",
																			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumeContainers/{volumeContainerName}/volumes/{volumeName}/metricsDefinitions", "2017-06-01"),
																		}},
																}},
														}},
												}},
										},
										{
											Display:  "volumes",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/volumes", "2017-06-01"),
										}},
								}},
						},
						{
							Display:  "default",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/encryptionSettings/default", "2017-06-01"),
						},
						{
							Display:        "vaultExtendedInfo",
							Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/extendedInformation/vaultExtendedInfo", "2017-06-01"),
							DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/extendedInformation/vaultExtendedInfo", "2017-06-01"),
							PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/extendedInformation/vaultExtendedInfo", "2017-06-01"),
							PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/extendedInformation/vaultExtendedInfo", "2017-06-01"),
						},
						{
							Display:  "fileservers",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/fileservers", "2016-10-01"),
						},
						{
							Display:  "iscsiservers",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/iscsiservers", "2016-10-01"),
						},
						{
							Display:  "jobs",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/jobs", "2017-06-01"),
						},
						{
							Display:  "metrics",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/metrics", "2017-06-01"),
						},
						{
							Display:  "metricsDefinitions",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/metricsDefinitions", "2017-06-01"),
						},
						{
							Display:  "storageAccountCredentials",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageAccountCredentials", "2017-06-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{credentialName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageAccountCredentials/{credentialName}", "2016-10-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageAccountCredentials/{credentialName}", "2016-10-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageAccountCredentials/{credentialName}", "2016-10-01"),
								},
								{
									Display:        "{storageAccountCredentialName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageAccountCredentials/{storageAccountCredentialName}", "2017-06-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageAccountCredentials/{storageAccountCredentialName}", "2017-06-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageAccountCredentials/{storageAccountCredentialName}", "2017-06-01"),
								}},
						},
						{
							Display:  "storageDomains",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageDomains", "2016-10-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{storageDomainName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageDomains/{storageDomainName}", "2016-10-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageDomains/{storageDomainName}", "2016-10-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageDomains/{storageDomainName}", "2016-10-01"),
								}},
						},
						{
							Display:  "bandwidthSettings",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/bandwidthSettings", "2017-06-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{bandwidthSettingName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/bandwidthSettings/{bandwidthSettingName}", "2017-06-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/bandwidthSettings/{bandwidthSettingName}", "2017-06-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/bandwidthSettings/{bandwidthSettingName}", "2017-06-01"),
								}},
						},
						{
							Display:  "cloudApplianceConfigurations",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/cloudApplianceConfigurations", "2017-06-01"),
						},
						{
							Display:  "features",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/features", "2017-06-01"),
						}},
					SubResources: []swagger.ResourceType{},
				}},
		},
		{
			Display:  "blobServices",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices", "2019-06-01"),
			Children: []swagger.ResourceType{
				{
					Display:  "containers",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers", "2019-06-01"),
					SubResources: []swagger.ResourceType{
						{
							Display:        "{containerName}",
							Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}", "2019-06-01"),
							DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}", "2019-06-01"),
							PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}", "2019-06-01"),
							PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}", "2019-06-01"),
							Children:       []swagger.ResourceType{},
							SubResources: []swagger.ResourceType{
								{
									Display:        "{immutabilityPolicyName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}/immutabilityPolicies/{immutabilityPolicyName}", "2019-06-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}/immutabilityPolicies/{immutabilityPolicyName}", "2019-06-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}/immutabilityPolicies/{immutabilityPolicyName}", "2019-06-01"),
								}},
						}},
				}},
			SubResources: []swagger.ResourceType{
				{
					Display:     "{BlobServicesName}",
					Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/{BlobServicesName}", "2019-06-01"),
					PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/{BlobServicesName}", "2019-06-01"),
				}},
		},
		{
			Display:  "fileServices",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices", "2019-06-01"),
			Children: []swagger.ResourceType{
				{
					Display:  "shares",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices/default/shares", "2019-06-01"),
					SubResources: []swagger.ResourceType{
						{
							Display:        "{shareName}",
							Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices/default/shares/{shareName}", "2019-06-01"),
							DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices/default/shares/{shareName}", "2019-06-01"),
							PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices/default/shares/{shareName}", "2019-06-01"),
							PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices/default/shares/{shareName}", "2019-06-01"),
						}},
				}},
			SubResources: []swagger.ResourceType{
				{
					Display:     "{FileServicesName}",
					Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices/{FileServicesName}", "2019-06-01"),
					PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices/{FileServicesName}", "2019-06-01"),
				}},
		},
		{
			Display:  "storageAccounts",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts", "2019-06-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{accountName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}", "2019-06-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}", "2019-06-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}", "2019-06-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}", "2019-06-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "privateLinkResources",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/privateLinkResources", "2019-06-01"),
						}},
					SubResources: []swagger.ResourceType{
						{
							Display:        "{managementPolicyName}",
							Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/managementPolicies/{managementPolicyName}", "2019-06-01"),
							DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/managementPolicies/{managementPolicyName}", "2019-06-01"),
							PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/managementPolicies/{managementPolicyName}", "2019-06-01"),
						},
						{
							Display:        "{privateEndpointConnectionName}",
							Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/privateEndpointConnections/{privateEndpointConnectionName}", "2019-06-01"),
							DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/privateEndpointConnections/{privateEndpointConnectionName}", "2019-06-01"),
							PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/privateEndpointConnections/{privateEndpointConnectionName}", "2019-06-01"),
						}},
				}},
		},
		{
			Display:  "jobs",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ImportExport/jobs", "2016-11-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{jobName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ImportExport/jobs/{jobName}", "2016-11-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ImportExport/jobs/{jobName}", "2016-11-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ImportExport/jobs/{jobName}", "2016-11-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ImportExport/jobs/{jobName}", "2016-11-01"),
					Children:       []swagger.ResourceType{},
				}},
		},
		{
			Display:  "{operationId}",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/locations/{locationName}/workflows/{workflowId}/operations/{operationId}", "2019-06-01"),
		},
		{
			Display:  "storageSyncServices",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices", "2019-06-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:        "{storageSyncServiceName}",
					Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}", "2019-06-01"),
					DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}", "2019-06-01"),
					PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}", "2019-06-01"),
					PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}", "2019-06-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "registeredServers",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/registeredServers", "2019-06-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{serverId}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/registeredServers/{serverId}", "2019-06-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/registeredServers/{serverId}", "2019-06-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/registeredServers/{serverId}", "2019-06-01"),
									Children:       []swagger.ResourceType{},
								}},
						},
						{
							Display:  "syncGroups",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups", "2019-06-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{syncGroupName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}", "2019-06-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}", "2019-06-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}", "2019-06-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "cloudEndpoints",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/cloudEndpoints", "2019-06-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{cloudEndpointName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/cloudEndpoints/{cloudEndpointName}", "2019-06-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/cloudEndpoints/{cloudEndpointName}", "2019-06-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/cloudEndpoints/{cloudEndpointName}", "2019-06-01"),
													Children:       []swagger.ResourceType{},
												}},
										},
										{
											Display:  "serverEndpoints",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/serverEndpoints", "2019-06-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{serverEndpointName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/serverEndpoints/{serverEndpointName}", "2019-06-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/serverEndpoints/{serverEndpointName}", "2019-06-01"),
													PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/serverEndpoints/{serverEndpointName}", "2019-06-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}/serverEndpoints/{serverEndpointName}", "2019-06-01"),
													Children:       []swagger.ResourceType{},
												}},
										}},
								}},
						},
						{
							Display:  "workflows",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/workflows", "2019-06-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:  "{workflowId}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/workflows/{workflowId}", "2019-06-01"),
									Children: []swagger.ResourceType{},
								}},
						}},
				}},
		},
		{
			Display:  "subscriptions",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions", "2016-06-01"),
			SubResources: []swagger.ResourceType{
				{
					Display:  "{subscriptionId}",
					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}", "2016-06-01"),
					Children: []swagger.ResourceType{
						{
							Display:  "locations",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/locations", "2016-06-01"),
						}},
					SubResources: []swagger.ResourceType{
						{
							Display:  "environments",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments", "2017-11-15"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{environmentName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}", "2017-11-15"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}", "2017-11-15"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}", "2017-11-15"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}", "2017-11-15"),
									Children: []swagger.ResourceType{
										{
											Display:  "accessPolicies",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/accessPolicies", "2017-11-15"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{accessPolicyName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/accessPolicies/{accessPolicyName}", "2017-11-15"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/accessPolicies/{accessPolicyName}", "2017-11-15"),
													PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/accessPolicies/{accessPolicyName}", "2017-11-15"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/accessPolicies/{accessPolicyName}", "2017-11-15"),
												}},
										},
										{
											Display:  "eventSources",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/eventSources", "2017-11-15"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{eventSourceName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/eventSources/{eventSourceName}", "2017-11-15"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/eventSources/{eventSourceName}", "2017-11-15"),
													PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/eventSources/{eventSourceName}", "2017-11-15"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/eventSources/{eventSourceName}", "2017-11-15"),
												}},
										},
										{
											Display:  "referenceDataSets",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/referenceDataSets", "2017-11-15"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{referenceDataSetName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/referenceDataSets/{referenceDataSetName}", "2017-11-15"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/referenceDataSets/{referenceDataSetName}", "2017-11-15"),
													PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/referenceDataSets/{referenceDataSetName}", "2017-11-15"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/referenceDataSets/{referenceDataSetName}", "2017-11-15"),
												}},
										}},
								}},
						},
						{
							Display:  "trafficmanagerprofiles",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles", "2018-04-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{profileName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}", "2018-04-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}", "2018-04-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}", "2018-04-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}", "2018-04-01"),
									SubResources: []swagger.ResourceType{
										{
											Display:  "{heatMapType}",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}/heatMaps/{heatMapType}", "2018-04-01"),
										},
										{
											Display:        "{endpointName}",
											Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}/{endpointType}/{endpointName}", "2018-04-01"),
											DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}/{endpointType}/{endpointName}", "2018-04-01"),
											PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}/{endpointType}/{endpointName}", "2018-04-01"),
											PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}/{endpointType}/{endpointName}", "2018-04-01"),
										}},
								}},
						},
						{
							Display:  "dedicatedCloudNodes",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes", "2019-04-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{dedicatedCloudNodeName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes/{dedicatedCloudNodeName}", "2019-04-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes/{dedicatedCloudNodeName}", "2019-04-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes/{dedicatedCloudNodeName}", "2019-04-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes/{dedicatedCloudNodeName}", "2019-04-01"),
								}},
						},
						{
							Display:  "dedicatedCloudServices",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudServices", "2019-04-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{dedicatedCloudServiceName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudServices/{dedicatedCloudServiceName}", "2019-04-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudServices/{dedicatedCloudServiceName}", "2019-04-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudServices/{dedicatedCloudServiceName}", "2019-04-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudServices/{dedicatedCloudServiceName}", "2019-04-01"),
								}},
						},
						{
							Display:  "virtualMachines",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/virtualMachines", "2019-04-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{virtualMachineName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/virtualMachines/{virtualMachineName}", "2019-04-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/virtualMachines/{virtualMachineName}", "2019-04-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/virtualMachines/{virtualMachineName}", "2019-04-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/virtualMachines/{virtualMachineName}", "2019-04-01"),
									Children:       []swagger.ResourceType{},
								}},
						},
						{
							Display:  "certificateOrders",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CertificateRegistration/certificateOrders", "2018-02-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{certificateOrderName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CertificateRegistration/certificateOrders/{certificateOrderName}", "2018-02-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CertificateRegistration/certificateOrders/{certificateOrderName}", "2018-02-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CertificateRegistration/certificateOrders/{certificateOrderName}", "2018-02-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CertificateRegistration/certificateOrders/{certificateOrderName}", "2018-02-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "certificates",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CertificateRegistration/certificateOrders/{certificateOrderName}/certificates", "2018-02-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{name}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CertificateRegistration/certificateOrders/{certificateOrderName}/certificates/{name}", "2018-02-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CertificateRegistration/certificateOrders/{certificateOrderName}/certificates/{name}", "2018-02-01"),
													PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CertificateRegistration/certificateOrders/{certificateOrderName}/certificates/{name}", "2018-02-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CertificateRegistration/certificateOrders/{certificateOrderName}/certificates/{name}", "2018-02-01"),
												}},
										}},
								}},
						},
						{
							Display:  "domains",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DomainRegistration/domains", "2018-02-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{domainName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DomainRegistration/domains/{domainName}", "2018-02-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DomainRegistration/domains/{domainName}", "2018-02-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DomainRegistration/domains/{domainName}", "2018-02-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DomainRegistration/domains/{domainName}", "2018-02-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "domainOwnershipIdentifiers",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DomainRegistration/domains/{domainName}/domainOwnershipIdentifiers", "2018-02-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{name}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DomainRegistration/domains/{domainName}/domainOwnershipIdentifiers/{name}", "2018-02-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DomainRegistration/domains/{domainName}/domainOwnershipIdentifiers/{name}", "2018-02-01"),
													PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DomainRegistration/domains/{domainName}/domainOwnershipIdentifiers/{name}", "2018-02-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DomainRegistration/domains/{domainName}/domainOwnershipIdentifiers/{name}", "2018-02-01"),
												}},
										}},
								}},
						},
						{
							Display:  "hostingEnvironments",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments", "2018-02-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{name}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}", "2018-02-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}", "2018-02-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}", "2018-02-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}", "2018-02-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "compute",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/capacities/compute", "2018-02-01"),
										},
										{
											Display:  "virtualip",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/capacities/virtualip", "2018-02-01"),
										},
										{
											Display:  "diagnostics",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/diagnostics", "2018-02-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:  "{diagnosticsName}",
													Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/diagnostics/{diagnosticsName}", "2018-02-01"),
												}},
										},
										{
											Display:  "inboundNetworkDependenciesEndpoints",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/inboundNetworkDependenciesEndpoints", "2018-02-01"),
										},
										{
											Display:  "metricdefinitions",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/metricdefinitions", "2018-02-01"),
										},
										{
											Display:  "metrics",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/metrics", "2018-02-01"),
										},
										{
											Display:  "multiRolePools",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/multiRolePools", "2018-02-01"),
											Children: []swagger.ResourceType{
												{
													Display:       "default",
													Endpoint:      endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/multiRolePools/default", "2018-02-01"),
													PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/multiRolePools/default", "2018-02-01"),
													PutEndpoint:   endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/multiRolePools/default", "2018-02-01"),
													Children: []swagger.ResourceType{
														{
															Display:  "metricdefinitions",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/multiRolePools/default/metricdefinitions", "2018-02-01"),
														},
														{
															Display:  "metrics",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/multiRolePools/default/metrics", "2018-02-01"),
														},
														{
															Display:  "skus",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/multiRolePools/default/skus", "2018-02-01"),
														},
														{
															Display:  "usages",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/multiRolePools/default/usages", "2018-02-01"),
														}},
													SubResources: []swagger.ResourceType{
														{
															Display:  "metricdefinitions",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/multiRolePools/default/instances/{instance}/metricdefinitions", "2018-02-01"),
														},
														{
															Display:  "metrics",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/multiRolePools/default/instances/{instance}/metrics", "2018-02-01"),
														}},
												}},
										},
										{
											Display:  "operations",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/operations", "2018-02-01"),
										},
										{
											Display:  "outboundNetworkDependenciesEndpoints",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/outboundNetworkDependenciesEndpoints", "2018-02-01"),
										},
										{
											Display:  "serverfarms",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/serverfarms", "2018-02-01"),
										},
										{
											Display:  "sites",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/sites", "2018-02-01"),
										},
										{
											Display:  "usages",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/usages", "2018-02-01"),
										},
										{
											Display:  "workerPools",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/workerPools", "2018-02-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:       "{workerPoolName}",
													Endpoint:      endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/workerPools/{workerPoolName}", "2018-02-01"),
													PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/workerPools/{workerPoolName}", "2018-02-01"),
													PutEndpoint:   endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/workerPools/{workerPoolName}", "2018-02-01"),
													Children: []swagger.ResourceType{
														{
															Display:  "metricdefinitions",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/workerPools/{workerPoolName}/metricdefinitions", "2018-02-01"),
														},
														{
															Display:  "metrics",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/workerPools/{workerPoolName}/metrics", "2018-02-01"),
														},
														{
															Display:  "skus",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/workerPools/{workerPoolName}/skus", "2018-02-01"),
														},
														{
															Display:  "usages",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/workerPools/{workerPoolName}/usages", "2018-02-01"),
														}},
													SubResources: []swagger.ResourceType{
														{
															Display:  "metricdefinitions",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/workerPools/{workerPoolName}/instances/{instance}/metricdefinitions", "2018-02-01"),
														},
														{
															Display:  "metrics",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/workerPools/{workerPoolName}/instances/{instance}/metrics", "2018-02-01"),
														}},
												}},
										},
										{
											Display:  "detectors",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/detectors", "2018-02-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:  "{detectorName}",
													Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}/detectors/{detectorName}", "2018-02-01"),
												}},
										}},
								},
								{
									Display:  "recommendationHistory",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{hostingEnvironmentName}/recommendationHistory", "2018-02-01"),
								},
								{
									Display:  "recommendations",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{hostingEnvironmentName}/recommendations", "2018-02-01"),
									Children: []swagger.ResourceType{},
									SubResources: []swagger.ResourceType{
										{
											Display:  "{name}",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{hostingEnvironmentName}/recommendations/{name}", "2018-02-01"),
											Children: []swagger.ResourceType{},
										}},
								}},
						},
						{
							Display:  "serverfarms",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms", "2018-02-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{name}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}", "2018-02-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}", "2018-02-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}", "2018-02-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}", "2018-02-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "capabilities",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/capabilities", "2018-02-01"),
										},
										{
											Display:  "limit",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/hybridConnectionPlanLimits/limit", "2018-02-01"),
										},
										{
											Display:  "hybridConnectionRelays",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/hybridConnectionRelays", "2018-02-01"),
										},
										{
											Display:  "metricdefinitions",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/metricdefinitions", "2018-02-01"),
										},
										{
											Display:  "metrics",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/metrics", "2018-02-01"),
										},
										{
											Display:  "sites",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/sites", "2018-02-01"),
										},
										{
											Display:  "skus",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/skus", "2018-02-01"),
										},
										{
											Display:  "usages",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/usages", "2018-02-01"),
										},
										{
											Display:  "virtualNetworkConnections",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/virtualNetworkConnections", "2018-02-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:  "{vnetName}",
													Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/virtualNetworkConnections/{vnetName}", "2018-02-01"),
													Children: []swagger.ResourceType{
														{
															Display:  "routes",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/virtualNetworkConnections/{vnetName}/routes", "2018-02-01"),
															SubResources: []swagger.ResourceType{
																{
																	Display:        "{routeName}",
																	Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/virtualNetworkConnections/{vnetName}/routes/{routeName}", "2018-02-01"),
																	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/virtualNetworkConnections/{vnetName}/routes/{routeName}", "2018-02-01"),
																	PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/virtualNetworkConnections/{vnetName}/routes/{routeName}", "2018-02-01"),
																	PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/virtualNetworkConnections/{vnetName}/routes/{routeName}", "2018-02-01"),
																}},
														}},
													SubResources: []swagger.ResourceType{
														{
															Display:     "{gatewayName}",
															Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/virtualNetworkConnections/{vnetName}/gateways/{gatewayName}", "2018-02-01"),
															PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/virtualNetworkConnections/{vnetName}/gateways/{gatewayName}", "2018-02-01"),
														}},
												}},
										}},
									SubResources: []swagger.ResourceType{
										{
											Display:        "{relayName}",
											Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}", "2018-02-01"),
											DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}", "2018-02-01"),
											Children: []swagger.ResourceType{
												{
													Display:  "sites",
													Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}/sites", "2018-02-01"),
												}},
										}},
								}},
						},
						{
							Display:  "certificates",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/certificates", "2018-02-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{name}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/certificates/{name}", "2018-02-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/certificates/{name}", "2018-02-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/certificates/{name}", "2018-02-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/certificates/{name}", "2018-02-01"),
								}},
						},
						{
							Display:  "detectors",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/detectors", "2018-02-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:  "{detectorName}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/detectors/{detectorName}", "2018-02-01"),
								}},
						},
						{
							Display:  "diagnostics",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/diagnostics", "2018-02-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:  "{diagnosticCategory}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/diagnostics/{diagnosticCategory}", "2018-02-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "analyses",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/diagnostics/{diagnosticCategory}/analyses", "2018-02-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:  "{analysisName}",
													Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/diagnostics/{diagnosticCategory}/analyses/{analysisName}", "2018-02-01"),
													Children: []swagger.ResourceType{},
												}},
										},
										{
											Display:  "detectors",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/diagnostics/{diagnosticCategory}/detectors", "2018-02-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:  "{detectorName}",
													Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/diagnostics/{diagnosticCategory}/detectors/{detectorName}", "2018-02-01"),
													Children: []swagger.ResourceType{},
												}},
										}},
								}},
						},
						{
							Display:  "detectors",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/slots/{slot}/detectors", "2018-02-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:  "{detectorName}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/slots/{slot}/detectors/{detectorName}", "2018-02-01"),
								}},
						},
						{
							Display:  "diagnostics",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/slots/{slot}/diagnostics", "2018-02-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:  "{diagnosticCategory}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/slots/{slot}/diagnostics/{diagnosticCategory}", "2018-02-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "analyses",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/slots/{slot}/diagnostics/{diagnosticCategory}/analyses", "2018-02-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:  "{analysisName}",
													Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/slots/{slot}/diagnostics/{diagnosticCategory}/analyses/{analysisName}", "2018-02-01"),
													Children: []swagger.ResourceType{},
												}},
										},
										{
											Display:  "detectors",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/slots/{slot}/diagnostics/{diagnosticCategory}/detectors", "2018-02-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:  "{detectorName}",
													Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/slots/{slot}/diagnostics/{diagnosticCategory}/detectors/{detectorName}", "2018-02-01"),
													Children: []swagger.ResourceType{},
												}},
										}},
								}},
						},
						{
							Display:  "recommendationHistory",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/recommendationHistory", "2018-02-01"),
						},
						{
							Display:  "recommendations",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/recommendations", "2018-02-01"),
							Children: []swagger.ResourceType{},
							SubResources: []swagger.ResourceType{
								{
									Display:  "{name}",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{siteName}/recommendations/{name}", "2018-02-01"),
									Children: []swagger.ResourceType{},
								}},
						},
						{
							Display:  "resourceHealthMetadata",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/resourceHealthMetadata", "2018-02-01"),
						},
						{
							Display:  "resourceHealthMetadata",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/resourceHealthMetadata", "2018-02-01"),
							Children: []swagger.ResourceType{
								{
									Display:  "default",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/resourceHealthMetadata/default", "2018-02-01"),
								}},
						},
						{
							Display:  "resourceHealthMetadata",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/resourceHealthMetadata", "2018-02-01"),
							Children: []swagger.ResourceType{
								{
									Display:  "default",
									Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/resourceHealthMetadata/default", "2018-02-01"),
								}},
						},
						{
							Display:  "sites",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites", "2018-02-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{name}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}", "2018-02-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}", "2018-02-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}", "2018-02-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}", "2018-02-01"),
									Children: []swagger.ResourceType{
										{
											Display:  "analyzeCustomHostname",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/analyzeCustomHostname", "2018-02-01"),
										},
										{
											Display:  "backups",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/backups", "2018-02-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{backupId}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/backups/{backupId}", "2018-02-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/backups/{backupId}", "2018-02-01"),
													Children:       []swagger.ResourceType{},
												}},
										},
										{
											Display:  "config",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config", "2018-02-01"),
											Children: []swagger.ResourceType{
												{
													Display:     "appsettings",
													Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/appsettings/list", "2018-02-01"),
													Verb:        "POST",
													PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/appsettings", "2018-02-01"),
												},
												{
													Display:     "authsettings",
													Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/authsettings/list", "2018-02-01"),
													Verb:        "POST",
													PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/authsettings", "2018-02-01"),
												},
												{
													Display:     "azurestorageaccounts",
													Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/azurestorageaccounts/list", "2018-02-01"),
													Verb:        "POST",
													PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/azurestorageaccounts", "2018-02-01"),
												},
												{
													Display:        "backup",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/backup/list", "2018-02-01"),
													Verb:           "POST",
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/backup", "2018-02-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/backup", "2018-02-01"),
												},
												{
													Display:     "connectionstrings",
													Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/connectionstrings/list", "2018-02-01"),
													Verb:        "POST",
													PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/connectionstrings", "2018-02-01"),
												},
												{
													Display:     "logs",
													Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/logs", "2018-02-01"),
													PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/logs", "2018-02-01"),
												},
												{
													Display:     "metadata",
													Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/metadata/list", "2018-02-01"),
													Verb:        "POST",
													PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/metadata", "2018-02-01"),
												},
												{
													Display:  "publishingcredentials",
													Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/publishingcredentials/list", "2018-02-01"),
													Verb:     "POST",
												},
												{
													Display:     "pushsettings",
													Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/pushsettings/list", "2018-02-01"),
													Verb:        "POST",
													PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/pushsettings", "2018-02-01"),
												},
												{
													Display:     "slotConfigNames",
													Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/slotConfigNames", "2018-02-01"),
													PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/slotConfigNames", "2018-02-01"),
												},
												{
													Display:       "web",
													Endpoint:      endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/web", "2018-02-01"),
													PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/web", "2018-02-01"),
													PutEndpoint:   endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/web", "2018-02-01"),
													Children: []swagger.ResourceType{
														{
															Display:  "snapshots",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/web/snapshots", "2018-02-01"),
															SubResources: []swagger.ResourceType{
																{
																	Display:  "{snapshotId}",
																	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/web/snapshots/{snapshotId}", "2018-02-01"),
																	Children: []swagger.ResourceType{},
																}},
														}},
												}},
										},
										{
											Display:  "continuouswebjobs",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/continuouswebjobs", "2018-02-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{webJobName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/continuouswebjobs/{webJobName}", "2018-02-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/continuouswebjobs/{webJobName}", "2018-02-01"),
													Children:       []swagger.ResourceType{},
												}},
										},
										{
											Display:  "deployments",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/deployments", "2018-02-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{id}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/deployments/{id}", "2018-02-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/deployments/{id}", "2018-02-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/deployments/{id}", "2018-02-01"),
													Children: []swagger.ResourceType{
														{
															Display:  "log",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/deployments/{id}/log", "2018-02-01"),
														}},
												}},
										},
										{
											Display:  "domainOwnershipIdentifiers",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/domainOwnershipIdentifiers", "2018-02-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{domainOwnershipIdentifierName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/domainOwnershipIdentifiers/{domainOwnershipIdentifierName}", "2018-02-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/domainOwnershipIdentifiers/{domainOwnershipIdentifierName}", "2018-02-01"),
													PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/domainOwnershipIdentifiers/{domainOwnershipIdentifierName}", "2018-02-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/domainOwnershipIdentifiers/{domainOwnershipIdentifierName}", "2018-02-01"),
												}},
										},
										{
											Display:     "MSDeploy",
											Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/extensions/MSDeploy", "2018-02-01"),
											PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/extensions/MSDeploy", "2018-02-01"),
											Children: []swagger.ResourceType{
												{
													Display:  "log",
													Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/extensions/MSDeploy/log", "2018-02-01"),
												}},
										},
										{
											Display:  "functions",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/functions", "2018-02-01"),
											Children: []swagger.ResourceType{
												{
													Display:  "token",
													Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/functions/admin/token", "2018-02-01"),
												}},
											SubResources: []swagger.ResourceType{
												{
													Display:        "{functionName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/functions/{functionName}", "2018-02-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/functions/{functionName}", "2018-02-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/functions/{functionName}", "2018-02-01"),
													Children:       []swagger.ResourceType{},
													SubResources:   []swagger.ResourceType{},
												}},
										},
										{
											Display:  "hostNameBindings",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hostNameBindings", "2018-02-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{hostName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hostNameBindings/{hostName}", "2018-02-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hostNameBindings/{hostName}", "2018-02-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hostNameBindings/{hostName}", "2018-02-01"),
												}},
										},
										{
											Display:  "hybridConnectionRelays",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hybridConnectionRelays", "2018-02-01"),
										},
										{
											Display:  "hybridconnection",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hybridconnection", "2018-02-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{entityName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hybridconnection/{entityName}", "2018-02-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hybridconnection/{entityName}", "2018-02-01"),
													PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hybridconnection/{entityName}", "2018-02-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hybridconnection/{entityName}", "2018-02-01"),
												}},
										},
										{
											Display:  "instances",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances", "2018-02-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:     "MSDeploy",
													Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/extensions/MSDeploy", "2018-02-01"),
													PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/extensions/MSDeploy", "2018-02-01"),
													Children: []swagger.ResourceType{
														{
															Display:  "log",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/extensions/MSDeploy/log", "2018-02-01"),
														}},
												},
												{
													Display:  "processes",
													Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes", "2018-02-01"),
													SubResources: []swagger.ResourceType{
														{
															Display:        "{processId}",
															Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes/{processId}", "2018-02-01"),
															DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes/{processId}", "2018-02-01"),
															Children: []swagger.ResourceType{
																{
																	Display:  "dump",
																	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes/{processId}/dump", "2018-02-01"),
																},
																{
																	Display:  "modules",
																	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes/{processId}/modules", "2018-02-01"),
																	SubResources: []swagger.ResourceType{
																		{
																			Display:  "{baseAddress}",
																			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes/{processId}/modules/{baseAddress}", "2018-02-01"),
																		}},
																},
																{
																	Display:  "threads",
																	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes/{processId}/threads", "2018-02-01"),
																	SubResources: []swagger.ResourceType{
																		{
																			Display:  "{threadId}",
																			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/instances/{instanceId}/processes/{processId}/threads/{threadId}", "2018-02-01"),
																		}},
																}},
														}},
												}},
										},
										{
											Display:  "metricdefinitions",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/metricdefinitions", "2018-02-01"),
										},
										{
											Display:  "metrics",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/metrics", "2018-02-01"),
										},
										{
											Display:        "virtualNetwork",
											Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/networkConfig/virtualNetwork", "2018-02-01"),
											DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/networkConfig/virtualNetwork", "2018-02-01"),
											PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/networkConfig/virtualNetwork", "2018-02-01"),
											PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/networkConfig/virtualNetwork", "2018-02-01"),
										},
										{
											Display:  "perfcounters",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/perfcounters", "2018-02-01"),
										},
										{
											Display:  "phplogging",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/phplogging", "2018-02-01"),
										},
										{
											Display:  "premieraddons",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/premieraddons", "2018-02-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{premierAddOnName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/premieraddons/{premierAddOnName}", "2018-02-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/premieraddons/{premierAddOnName}", "2018-02-01"),
													PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/premieraddons/{premierAddOnName}", "2018-02-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/premieraddons/{premierAddOnName}", "2018-02-01"),
												}},
										},
										{
											Display:     "virtualNetworks",
											Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/privateAccess/virtualNetworks", "2018-02-01"),
											PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/privateAccess/virtualNetworks", "2018-02-01"),
										},
										{
											Display:  "processes",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/processes", "2018-02-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{processId}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/processes/{processId}", "2018-02-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/processes/{processId}", "2018-02-01"),
													Children: []swagger.ResourceType{
														{
															Display:  "dump",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/processes/{processId}/dump", "2018-02-01"),
														},
														{
															Display:  "modules",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/processes/{processId}/modules", "2018-02-01"),
															SubResources: []swagger.ResourceType{
																{
																	Display:  "{baseAddress}",
																	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/processes/{processId}/modules/{baseAddress}", "2018-02-01"),
																}},
														},
														{
															Display:  "threads",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/processes/{processId}/threads", "2018-02-01"),
															SubResources: []swagger.ResourceType{
																{
																	Display:  "{threadId}",
																	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/processes/{processId}/threads/{threadId}", "2018-02-01"),
																}},
														}},
												}},
										},
										{
											Display:  "publicCertificates",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/publicCertificates", "2018-02-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{publicCertificateName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/publicCertificates/{publicCertificateName}", "2018-02-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/publicCertificates/{publicCertificateName}", "2018-02-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/publicCertificates/{publicCertificateName}", "2018-02-01"),
												}},
										},
										{
											Display:  "siteextensions",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/siteextensions", "2018-02-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{siteExtensionId}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/siteextensions/{siteExtensionId}", "2018-02-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/siteextensions/{siteExtensionId}", "2018-02-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/siteextensions/{siteExtensionId}", "2018-02-01"),
												}},
										},
										{
											Display:  "slots",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots", "2018-02-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{slot}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}", "2018-02-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}", "2018-02-01"),
													PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}", "2018-02-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}", "2018-02-01"),
													Children: []swagger.ResourceType{
														{
															Display:  "analyzeCustomHostname",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/analyzeCustomHostname", "2018-02-01"),
														},
														{
															Display:  "backups",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/backups", "2018-02-01"),
															SubResources: []swagger.ResourceType{
																{
																	Display:        "{backupId}",
																	Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/backups/{backupId}", "2018-02-01"),
																	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/backups/{backupId}", "2018-02-01"),
																	Children:       []swagger.ResourceType{},
																}},
														},
														{
															Display:  "config",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config", "2018-02-01"),
															Children: []swagger.ResourceType{
																{
																	Display:     "logs",
																	Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/logs", "2018-02-01"),
																	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/logs", "2018-02-01"),
																},
																{
																	Display:       "web",
																	Endpoint:      endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/web", "2018-02-01"),
																	PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/web", "2018-02-01"),
																	PutEndpoint:   endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/web", "2018-02-01"),
																	Children: []swagger.ResourceType{
																		{
																			Display:  "snapshots",
																			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/web/snapshots", "2018-02-01"),
																			SubResources: []swagger.ResourceType{
																				{
																					Display:  "{snapshotId}",
																					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/config/web/snapshots/{snapshotId}", "2018-02-01"),
																					Children: []swagger.ResourceType{},
																				}},
																		}},
																}},
														},
														{
															Display:  "continuouswebjobs",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/continuouswebjobs", "2018-02-01"),
															SubResources: []swagger.ResourceType{
																{
																	Display:        "{webJobName}",
																	Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/continuouswebjobs/{webJobName}", "2018-02-01"),
																	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/continuouswebjobs/{webJobName}", "2018-02-01"),
																	Children:       []swagger.ResourceType{},
																}},
														},
														{
															Display:  "deployments",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/deployments", "2018-02-01"),
															SubResources: []swagger.ResourceType{
																{
																	Display:        "{id}",
																	Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/deployments/{id}", "2018-02-01"),
																	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/deployments/{id}", "2018-02-01"),
																	PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/deployments/{id}", "2018-02-01"),
																	Children: []swagger.ResourceType{
																		{
																			Display:  "log",
																			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/deployments/{id}/log", "2018-02-01"),
																		}},
																}},
														},
														{
															Display:  "domainOwnershipIdentifiers",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/domainOwnershipIdentifiers", "2018-02-01"),
															SubResources: []swagger.ResourceType{
																{
																	Display:        "{domainOwnershipIdentifierName}",
																	Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/domainOwnershipIdentifiers/{domainOwnershipIdentifierName}", "2018-02-01"),
																	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/domainOwnershipIdentifiers/{domainOwnershipIdentifierName}", "2018-02-01"),
																	PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/domainOwnershipIdentifiers/{domainOwnershipIdentifierName}", "2018-02-01"),
																	PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/domainOwnershipIdentifiers/{domainOwnershipIdentifierName}", "2018-02-01"),
																}},
														},
														{
															Display:     "MSDeploy",
															Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/extensions/MSDeploy", "2018-02-01"),
															PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/extensions/MSDeploy", "2018-02-01"),
															Children: []swagger.ResourceType{
																{
																	Display:  "log",
																	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/extensions/MSDeploy/log", "2018-02-01"),
																}},
														},
														{
															Display:  "functions",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/functions", "2018-02-01"),
															Children: []swagger.ResourceType{
																{
																	Display:  "token",
																	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/functions/admin/token", "2018-02-01"),
																}},
															SubResources: []swagger.ResourceType{
																{
																	Display:        "{functionName}",
																	Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/functions/{functionName}", "2018-02-01"),
																	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/functions/{functionName}", "2018-02-01"),
																	PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/functions/{functionName}", "2018-02-01"),
																	Children:       []swagger.ResourceType{},
																	SubResources:   []swagger.ResourceType{},
																}},
														},
														{
															Display:  "hostNameBindings",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hostNameBindings", "2018-02-01"),
															SubResources: []swagger.ResourceType{
																{
																	Display:        "{hostName}",
																	Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hostNameBindings/{hostName}", "2018-02-01"),
																	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hostNameBindings/{hostName}", "2018-02-01"),
																	PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hostNameBindings/{hostName}", "2018-02-01"),
																}},
														},
														{
															Display:  "hybridConnectionRelays",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hybridConnectionRelays", "2018-02-01"),
														},
														{
															Display:  "hybridconnection",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hybridconnection", "2018-02-01"),
															SubResources: []swagger.ResourceType{
																{
																	Display:        "{entityName}",
																	Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hybridconnection/{entityName}", "2018-02-01"),
																	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hybridconnection/{entityName}", "2018-02-01"),
																	PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hybridconnection/{entityName}", "2018-02-01"),
																	PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hybridconnection/{entityName}", "2018-02-01"),
																}},
														},
														{
															Display:  "instances",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances", "2018-02-01"),
															SubResources: []swagger.ResourceType{
																{
																	Display:     "MSDeploy",
																	Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/extensions/MSDeploy", "2018-02-01"),
																	PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/extensions/MSDeploy", "2018-02-01"),
																	Children: []swagger.ResourceType{
																		{
																			Display:  "log",
																			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/extensions/MSDeploy/log", "2018-02-01"),
																		}},
																},
																{
																	Display:  "processes",
																	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes", "2018-02-01"),
																	SubResources: []swagger.ResourceType{
																		{
																			Display:        "{processId}",
																			Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes/{processId}", "2018-02-01"),
																			DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes/{processId}", "2018-02-01"),
																			Children: []swagger.ResourceType{
																				{
																					Display:  "dump",
																					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes/{processId}/dump", "2018-02-01"),
																				},
																				{
																					Display:  "modules",
																					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes/{processId}/modules", "2018-02-01"),
																					SubResources: []swagger.ResourceType{
																						{
																							Display:  "{baseAddress}",
																							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes/{processId}/modules/{baseAddress}", "2018-02-01"),
																						}},
																				},
																				{
																					Display:  "threads",
																					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes/{processId}/threads", "2018-02-01"),
																					SubResources: []swagger.ResourceType{
																						{
																							Display:  "{threadId}",
																							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/instances/{instanceId}/processes/{processId}/threads/{threadId}", "2018-02-01"),
																						}},
																				}},
																		}},
																}},
														},
														{
															Display:  "metricdefinitions",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/metricdefinitions", "2018-02-01"),
														},
														{
															Display:  "metrics",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/metrics", "2018-02-01"),
														},
														{
															Display:  "status",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/migratemysql/status", "2018-02-01"),
														},
														{
															Display:        "virtualNetwork",
															Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/networkConfig/virtualNetwork", "2018-02-01"),
															DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/networkConfig/virtualNetwork", "2018-02-01"),
															PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/networkConfig/virtualNetwork", "2018-02-01"),
															PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/networkConfig/virtualNetwork", "2018-02-01"),
														},
														{
															Display:  "perfcounters",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/perfcounters", "2018-02-01"),
														},
														{
															Display:  "phplogging",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/phplogging", "2018-02-01"),
														},
														{
															Display:  "premieraddons",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/premieraddons", "2018-02-01"),
															SubResources: []swagger.ResourceType{
																{
																	Display:        "{premierAddOnName}",
																	Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/premieraddons/{premierAddOnName}", "2018-02-01"),
																	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/premieraddons/{premierAddOnName}", "2018-02-01"),
																	PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/premieraddons/{premierAddOnName}", "2018-02-01"),
																	PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/premieraddons/{premierAddOnName}", "2018-02-01"),
																}},
														},
														{
															Display:     "virtualNetworks",
															Endpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/privateAccess/virtualNetworks", "2018-02-01"),
															PutEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/privateAccess/virtualNetworks", "2018-02-01"),
														},
														{
															Display:  "processes",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/processes", "2018-02-01"),
															SubResources: []swagger.ResourceType{
																{
																	Display:        "{processId}",
																	Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/processes/{processId}", "2018-02-01"),
																	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/processes/{processId}", "2018-02-01"),
																	Children: []swagger.ResourceType{
																		{
																			Display:  "dump",
																			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/processes/{processId}/dump", "2018-02-01"),
																		},
																		{
																			Display:  "modules",
																			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/processes/{processId}/modules", "2018-02-01"),
																			SubResources: []swagger.ResourceType{
																				{
																					Display:  "{baseAddress}",
																					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/processes/{processId}/modules/{baseAddress}", "2018-02-01"),
																				}},
																		},
																		{
																			Display:  "threads",
																			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/processes/{processId}/threads", "2018-02-01"),
																			SubResources: []swagger.ResourceType{
																				{
																					Display:  "{threadId}",
																					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/processes/{processId}/threads/{threadId}", "2018-02-01"),
																				}},
																		}},
																}},
														},
														{
															Display:  "publicCertificates",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/publicCertificates", "2018-02-01"),
															SubResources: []swagger.ResourceType{
																{
																	Display:        "{publicCertificateName}",
																	Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/publicCertificates/{publicCertificateName}", "2018-02-01"),
																	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/publicCertificates/{publicCertificateName}", "2018-02-01"),
																	PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/publicCertificates/{publicCertificateName}", "2018-02-01"),
																}},
														},
														{
															Display:  "siteextensions",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/siteextensions", "2018-02-01"),
															SubResources: []swagger.ResourceType{
																{
																	Display:        "{siteExtensionId}",
																	Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/siteextensions/{siteExtensionId}", "2018-02-01"),
																	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/siteextensions/{siteExtensionId}", "2018-02-01"),
																	PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/siteextensions/{siteExtensionId}", "2018-02-01"),
																}},
														},
														{
															Display:  "snapshots",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/snapshots", "2018-02-01"),
														},
														{
															Display:  "snapshotsdr",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/snapshotsdr", "2018-02-01"),
														},
														{
															Display:        "web",
															Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/sourcecontrols/web", "2018-02-01"),
															DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/sourcecontrols/web", "2018-02-01"),
															PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/sourcecontrols/web", "2018-02-01"),
															PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/sourcecontrols/web", "2018-02-01"),
														},
														{
															Display:  "triggeredwebjobs",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/triggeredwebjobs", "2018-02-01"),
															SubResources: []swagger.ResourceType{
																{
																	Display:        "{webJobName}",
																	Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/triggeredwebjobs/{webJobName}", "2018-02-01"),
																	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/triggeredwebjobs/{webJobName}", "2018-02-01"),
																	Children: []swagger.ResourceType{
																		{
																			Display:  "history",
																			Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/triggeredwebjobs/{webJobName}/history", "2018-02-01"),
																			SubResources: []swagger.ResourceType{
																				{
																					Display:  "{id}",
																					Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/triggeredwebjobs/{webJobName}/history/{id}", "2018-02-01"),
																				}},
																		}},
																}},
														},
														{
															Display:  "usages",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/usages", "2018-02-01"),
														},
														{
															Display:  "virtualNetworkConnections",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/virtualNetworkConnections", "2018-02-01"),
															SubResources: []swagger.ResourceType{
																{
																	Display:        "{vnetName}",
																	Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/virtualNetworkConnections/{vnetName}", "2018-02-01"),
																	DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/virtualNetworkConnections/{vnetName}", "2018-02-01"),
																	PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/virtualNetworkConnections/{vnetName}", "2018-02-01"),
																	PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/virtualNetworkConnections/{vnetName}", "2018-02-01"),
																	SubResources: []swagger.ResourceType{
																		{
																			Display:       "{gatewayName}",
																			Endpoint:      endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/virtualNetworkConnections/{vnetName}/gateways/{gatewayName}", "2018-02-01"),
																			PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/virtualNetworkConnections/{vnetName}/gateways/{gatewayName}", "2018-02-01"),
																			PutEndpoint:   endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/virtualNetworkConnections/{vnetName}/gateways/{gatewayName}", "2018-02-01"),
																		}},
																}},
														},
														{
															Display:  "webjobs",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/webjobs", "2018-02-01"),
															SubResources: []swagger.ResourceType{
																{
																	Display:  "{webJobName}",
																	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/webjobs/{webJobName}", "2018-02-01"),
																}},
														}},
													SubResources: []swagger.ResourceType{
														{
															Display:        "{relayName}",
															Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}", "2018-02-01"),
															DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}", "2018-02-01"),
															PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}", "2018-02-01"),
															PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}", "2018-02-01"),
															Children:       []swagger.ResourceType{},
														},
														{
															Display:  "{view}",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/networkFeatures/{view}", "2018-02-01"),
														},
														{
															Display:  "{operationId}",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/networkTrace/operationresults/{operationId}", "2018-02-01"),
														},
														{
															Display:  "{operationId}",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/networkTrace/{operationId}", "2018-02-01"),
														},
														{
															Display:  "{operationId}",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/networkTraces/current/operationresults/{operationId}", "2018-02-01"),
														},
														{
															Display:  "{operationId}",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/networkTraces/{operationId}", "2018-02-01"),
														}},
												}},
										},
										{
											Display:  "snapshots",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/snapshots", "2018-02-01"),
										},
										{
											Display:  "snapshotsdr",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/snapshotsdr", "2018-02-01"),
										},
										{
											Display:        "web",
											Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/sourcecontrols/web", "2018-02-01"),
											DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/sourcecontrols/web", "2018-02-01"),
											PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/sourcecontrols/web", "2018-02-01"),
											PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/sourcecontrols/web", "2018-02-01"),
										},
										{
											Display:  "triggeredwebjobs",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/triggeredwebjobs", "2018-02-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{webJobName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/triggeredwebjobs/{webJobName}", "2018-02-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/triggeredwebjobs/{webJobName}", "2018-02-01"),
													Children: []swagger.ResourceType{
														{
															Display:  "history",
															Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/triggeredwebjobs/{webJobName}/history", "2018-02-01"),
															SubResources: []swagger.ResourceType{
																{
																	Display:  "{id}",
																	Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/triggeredwebjobs/{webJobName}/history/{id}", "2018-02-01"),
																}},
														}},
												}},
										},
										{
											Display:  "usages",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/usages", "2018-02-01"),
										},
										{
											Display:  "virtualNetworkConnections",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/virtualNetworkConnections", "2018-02-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:        "{vnetName}",
													Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/virtualNetworkConnections/{vnetName}", "2018-02-01"),
													DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/virtualNetworkConnections/{vnetName}", "2018-02-01"),
													PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/virtualNetworkConnections/{vnetName}", "2018-02-01"),
													PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/virtualNetworkConnections/{vnetName}", "2018-02-01"),
													SubResources: []swagger.ResourceType{
														{
															Display:       "{gatewayName}",
															Endpoint:      endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/virtualNetworkConnections/{vnetName}/gateways/{gatewayName}", "2018-02-01"),
															PatchEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/virtualNetworkConnections/{vnetName}/gateways/{gatewayName}", "2018-02-01"),
															PutEndpoint:   endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/virtualNetworkConnections/{vnetName}/gateways/{gatewayName}", "2018-02-01"),
														}},
												}},
										},
										{
											Display:  "webjobs",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/webjobs", "2018-02-01"),
											SubResources: []swagger.ResourceType{
												{
													Display:  "{webJobName}",
													Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/webjobs/{webJobName}", "2018-02-01"),
												}},
										}},
									SubResources: []swagger.ResourceType{
										{
											Display:        "{relayName}",
											Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}", "2018-02-01"),
											DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}", "2018-02-01"),
											PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}", "2018-02-01"),
											PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hybridConnectionNamespaces/{namespaceName}/relays/{relayName}", "2018-02-01"),
											Children:       []swagger.ResourceType{},
										},
										{
											Display:  "{view}",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/networkFeatures/{view}", "2018-02-01"),
										},
										{
											Display:  "{operationId}",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/networkTrace/operationresults/{operationId}", "2018-02-01"),
										},
										{
											Display:  "{operationId}",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/networkTrace/{operationId}", "2018-02-01"),
										},
										{
											Display:  "{operationId}",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/networkTraces/current/operationresults/{operationId}", "2018-02-01"),
										},
										{
											Display:  "{operationId}",
											Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/networkTraces/{operationId}", "2018-02-01"),
										}},
								}},
						},
						{
							Display:  "deviceServices",
							Endpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.WindowsIoT/deviceServices", "2019-06-01"),
							SubResources: []swagger.ResourceType{
								{
									Display:        "{deviceName}",
									Endpoint:       endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.WindowsIoT/deviceServices/{deviceName}", "2019-06-01"),
									DeleteEndpoint: endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.WindowsIoT/deviceServices/{deviceName}", "2019-06-01"),
									PatchEndpoint:  endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.WindowsIoT/deviceServices/{deviceName}", "2019-06-01"),
									PutEndpoint:    endpoints.MustGetEndpointInfoFromURL("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.WindowsIoT/deviceServices/{deviceName}", "2019-06-01"),
								}},
						}},
				}},
		},
		{
			Display:  "tenants",
			Endpoint: endpoints.MustGetEndpointInfoFromURL("/tenants", "2016-06-01"),
		}}

}