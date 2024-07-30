# \DataPlatformClusterApi

All URIs are relative to *https://api.ionos.com/dataplatform*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**ClustersDelete**](DataPlatformClusterApi.md#ClustersDelete) | **Delete** /clusters/{clusterId} | Delete a DataPlatformCluster|
|[**ClustersFindById**](DataPlatformClusterApi.md#ClustersFindById) | **Get** /clusters/{clusterId} | Retrieve a DataPlatformCluster|
|[**ClustersGet**](DataPlatformClusterApi.md#ClustersGet) | **Get** /clusters | List the DataPlatformClusters|
|[**ClustersKubeconfigFindByClusterId**](DataPlatformClusterApi.md#ClustersKubeconfigFindByClusterId) | **Get** /clusters/{clusterId}/kubeconfig | Read the Kubeconfig|
|[**ClustersPatch**](DataPlatformClusterApi.md#ClustersPatch) | **Patch** /clusters/{clusterId} | Partially Modify a DataPlatformCluster|
|[**ClustersPost**](DataPlatformClusterApi.md#ClustersPost) | **Post** /clusters | Create a DataPlatformCluster|



## ClustersDelete

```go
var result ClusterResponseData = ClustersDelete(ctx, clusterId)
                      .Execute()
```

Delete a DataPlatformCluster



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    dataplatform "github.com/ionos-cloud/sdk-go-bundle/products/dataplatform"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the cluster. Must conform to the UUID format. 

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := dataplatform.NewAPIClient(configuration)
    resp, err := apiClient.DataPlatformClusterApi.ClustersDelete(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataPlatformClusterApi.ClustersDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersDelete`: ClusterResponseData
    fmt.Fprintf(os.Stdout, "Response from `DataPlatformClusterApi.ClustersDelete`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The unique ID of the cluster. Must conform to the UUID format.  | |

### Other Parameters

Other parameters are passed through a pointer to an apiClustersDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**ClusterResponseData**](../models/ClusterResponseData.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ClustersFindById

```go
var result ClusterResponseData = ClustersFindById(ctx, clusterId)
                      .Execute()
```

Retrieve a DataPlatformCluster



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    dataplatform "github.com/ionos-cloud/sdk-go-bundle/products/dataplatform"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the cluster. Must conform to the UUID format. 

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := dataplatform.NewAPIClient(configuration)
    resource, resp, err := apiClient.DataPlatformClusterApi.ClustersFindById(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataPlatformClusterApi.ClustersFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersFindById`: ClusterResponseData
    fmt.Fprintf(os.Stdout, "Response from `DataPlatformClusterApi.ClustersFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The unique ID of the cluster. Must conform to the UUID format.  | |

### Other Parameters

Other parameters are passed through a pointer to an apiClustersFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**ClusterResponseData**](../models/ClusterResponseData.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ClustersGet

```go
var result ClusterListResponseData = ClustersGet(ctx)
                      .Name(name)
                      .Execute()
```

List the DataPlatformClusters



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    dataplatform "github.com/ionos-cloud/sdk-go-bundle/products/dataplatform"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    name := "name_example" // string | Response filter to list only the clusters which include the specified name. The value is case insensitive and matched on the `name` property of the cluster. The input is limited to 63 characters with alphanumeric characters (`[a-z0-9A-Z]`), dashes (`-`), underscores (`_`), and dots (`.`) allowed.  (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := dataplatform.NewAPIClient(configuration)
    resource, resp, err := apiClient.DataPlatformClusterApi.ClustersGet(context.Background()).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataPlatformClusterApi.ClustersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersGet`: ClusterListResponseData
    fmt.Fprintf(os.Stdout, "Response from `DataPlatformClusterApi.ClustersGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiClustersGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **name** | **string** | Response filter to list only the clusters which include the specified name. The value is case insensitive and matched on the &#x60;name&#x60; property of the cluster. The input is limited to 63 characters with alphanumeric characters (&#x60;[a-z0-9A-Z]&#x60;), dashes (&#x60;-&#x60;), underscores (&#x60;_&#x60;), and dots (&#x60;.&#x60;) allowed.  | |

### Return type

[**ClusterListResponseData**](../models/ClusterListResponseData.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ClustersKubeconfigFindByClusterId

```go
var result map[string]interface{} = ClustersKubeconfigFindByClusterId(ctx, clusterId)
                      .Execute()
```

Read the Kubeconfig



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    dataplatform "github.com/ionos-cloud/sdk-go-bundle/products/dataplatform"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the cluster. Must conform to the UUID format. 

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := dataplatform.NewAPIClient(configuration)
    resource, resp, err := apiClient.DataPlatformClusterApi.ClustersKubeconfigFindByClusterId(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataPlatformClusterApi.ClustersKubeconfigFindByClusterId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersKubeconfigFindByClusterId`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DataPlatformClusterApi.ClustersKubeconfigFindByClusterId`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The unique ID of the cluster. Must conform to the UUID format.  | |

### Other Parameters

Other parameters are passed through a pointer to an apiClustersKubeconfigFindByClusterIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

**map[string]interface{}**

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ClustersPatch

```go
var result ClusterResponseData = ClustersPatch(ctx, clusterId)
                      .PatchClusterRequest(patchClusterRequest)
                      .Execute()
```

Partially Modify a DataPlatformCluster



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    dataplatform "github.com/ionos-cloud/sdk-go-bundle/products/dataplatform"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    clusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the cluster. Must conform to the UUID format. 
    patchClusterRequest := *openapiclient.NewPatchClusterRequest(*openapiclient.NewPatchClusterProperties()) // PatchClusterRequest | Request payload with the properties that shall be applied to an existing DataPlatformCluster. 

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := dataplatform.NewAPIClient(configuration)
    resource, resp, err := apiClient.DataPlatformClusterApi.ClustersPatch(context.Background(), clusterId).PatchClusterRequest(patchClusterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataPlatformClusterApi.ClustersPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersPatch`: ClusterResponseData
    fmt.Fprintf(os.Stdout, "Response from `DataPlatformClusterApi.ClustersPatch`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The unique ID of the cluster. Must conform to the UUID format.  | |

### Other Parameters

Other parameters are passed through a pointer to an apiClustersPatchRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **patchClusterRequest** | [**PatchClusterRequest**](../models/PatchClusterRequest.md) | Request payload with the properties that shall be applied to an existing DataPlatformCluster.  | |

### Return type

[**ClusterResponseData**](../models/ClusterResponseData.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## ClustersPost

```go
var result ClusterResponseData = ClustersPost(ctx)
                      .CreateClusterRequest(createClusterRequest)
                      .Execute()
```

Create a DataPlatformCluster



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    dataplatform "github.com/ionos-cloud/sdk-go-bundle/products/dataplatform"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    createClusterRequest := *openapiclient.NewCreateClusterRequest(*openapiclient.NewCreateClusterProperties("my-cluster", "DatacenterId_example")) // CreateClusterRequest | Request payload with the properties that defines a new DataPlatformCluster and the credentials to interact with the PaaS API to create it. 

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := dataplatform.NewAPIClient(configuration)
    resource, resp, err := apiClient.DataPlatformClusterApi.ClustersPost(context.Background()).CreateClusterRequest(createClusterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataPlatformClusterApi.ClustersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersPost`: ClusterResponseData
    fmt.Fprintf(os.Stdout, "Response from `DataPlatformClusterApi.ClustersPost`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiClustersPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **createClusterRequest** | [**CreateClusterRequest**](../models/CreateClusterRequest.md) | Request payload with the properties that defines a new DataPlatformCluster and the credentials to interact with the PaaS API to create it.  | |

### Return type

[**ClusterResponseData**](../models/ClusterResponseData.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


