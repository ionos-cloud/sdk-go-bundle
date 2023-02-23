# \DataPlatformClusterApi

All URIs are relative to *https://api.ionos.com/dataplatform*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**CreateCluster**](DataPlatformClusterApi.md#CreateCluster) | **Post** /clusters | Create a DataPlatformCluster|
|[**DeleteCluster**](DataPlatformClusterApi.md#DeleteCluster) | **Delete** /clusters/{clusterId} | Delete DataPlatformCluster|
|[**GetCluster**](DataPlatformClusterApi.md#GetCluster) | **Get** /clusters/{clusterId} | Retrieve a DataPlatformCluster|
|[**GetClusterKubeconfig**](DataPlatformClusterApi.md#GetClusterKubeconfig) | **Get** /clusters/{clusterId}/kubeconfig | Read the kubeconfig|
|[**GetClusters**](DataPlatformClusterApi.md#GetClusters) | **Get** /clusters | List DataPlatformCluster|
|[**PatchCluster**](DataPlatformClusterApi.md#PatchCluster) | **Patch** /clusters/{clusterId} | Partially modify a DataPlatformCluster|



## CreateCluster

```go
var result ClusterResponseData = CreateCluster(ctx)
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
    resource, resp, err := apiClient.DataPlatformClusterApi.CreateCluster(context.Background()).CreateClusterRequest(createClusterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataPlatformClusterApi.CreateCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `CreateCluster`: ClusterResponseData
    fmt.Fprintf(os.Stdout, "Response from `DataPlatformClusterApi.CreateCluster`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiCreateClusterRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **createClusterRequest** | [**CreateClusterRequest**](../models/CreateClusterRequest.md) | Request payload with the properties that defines a new DataPlatformCluster and the credentials to interact with the PaaS API to create it.  | |

### Return type

[**ClusterResponseData**](../models/ClusterResponseData.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DeleteCluster

```go
var result ClusterResponseData = DeleteCluster(ctx, clusterId)
                      .Execute()
```

Delete DataPlatformCluster



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
    clusterId := TODO // string | The unique ID of the cluster. Must conform to the UUID format. 

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := dataplatform.NewAPIClient(configuration)
    resource, resp, err := apiClient.DataPlatformClusterApi.DeleteCluster(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataPlatformClusterApi.DeleteCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `DeleteCluster`: ClusterResponseData
    fmt.Fprintf(os.Stdout, "Response from `DataPlatformClusterApi.DeleteCluster`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | [**string**](../models/.md) | The unique ID of the cluster. Must conform to the UUID format.  | |

### Other Parameters

Other parameters are passed through a pointer to an apiDeleteClusterRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**ClusterResponseData**](../models/ClusterResponseData.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## GetCluster

```go
var result ClusterResponseData = GetCluster(ctx, clusterId)
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
    clusterId := TODO // string | The unique ID of the cluster. Must conform to the UUID format. 

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := dataplatform.NewAPIClient(configuration)
    resource, resp, err := apiClient.DataPlatformClusterApi.GetCluster(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataPlatformClusterApi.GetCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `GetCluster`: ClusterResponseData
    fmt.Fprintf(os.Stdout, "Response from `DataPlatformClusterApi.GetCluster`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | [**string**](../models/.md) | The unique ID of the cluster. Must conform to the UUID format.  | |

### Other Parameters

Other parameters are passed through a pointer to an apiGetClusterRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**ClusterResponseData**](../models/ClusterResponseData.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## GetClusterKubeconfig

```go
var result string = GetClusterKubeconfig(ctx, clusterId)
                      .Execute()
```

Read the kubeconfig



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
    clusterId := TODO // string | The unique ID of the cluster. Must conform to the UUID format. 

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := dataplatform.NewAPIClient(configuration)
    resource, resp, err := apiClient.DataPlatformClusterApi.GetClusterKubeconfig(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataPlatformClusterApi.GetClusterKubeconfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `GetClusterKubeconfig`: string
    fmt.Fprintf(os.Stdout, "Response from `DataPlatformClusterApi.GetClusterKubeconfig`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | [**string**](../models/.md) | The unique ID of the cluster. Must conform to the UUID format.  | |

### Other Parameters

Other parameters are passed through a pointer to an apiGetClusterKubeconfigRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

**string**

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## GetClusters

```go
var result ClusterListResponseData = GetClusters(ctx)
                      .Name(name)
                      .Execute()
```

List DataPlatformCluster



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
    name := "name_example" // string | Response filter to list only the clusters which include the specified name. The value is case insensitive and matched on the `name` property of the cluster. The input is limited to 63 characters with alphanumeric character ([a-z0-9A-Z]) dashes (-), underscores (_), dots (.), and alphanumerics allowed.  (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := dataplatform.NewAPIClient(configuration)
    resource, resp, err := apiClient.DataPlatformClusterApi.GetClusters(context.Background()).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataPlatformClusterApi.GetClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `GetClusters`: ClusterListResponseData
    fmt.Fprintf(os.Stdout, "Response from `DataPlatformClusterApi.GetClusters`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiGetClustersRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **name** | **string** | Response filter to list only the clusters which include the specified name. The value is case insensitive and matched on the &#x60;name&#x60; property of the cluster. The input is limited to 63 characters with alphanumeric character ([a-z0-9A-Z]) dashes (-), underscores (_), dots (.), and alphanumerics allowed.  | |

### Return type

[**ClusterListResponseData**](../models/ClusterListResponseData.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## PatchCluster

```go
var result ClusterResponseData = PatchCluster(ctx, clusterId)
                      .PatchClusterRequest(patchClusterRequest)
                      .Execute()
```

Partially modify a DataPlatformCluster



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
    clusterId := TODO // string | The unique ID of the cluster. Must conform to the UUID format. 
    patchClusterRequest := *openapiclient.NewPatchClusterRequest(*openapiclient.NewPatchClusterProperties()) // PatchClusterRequest | Request payload with the properties that shall be applied to an existing DataPlatformCluster. 

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := dataplatform.NewAPIClient(configuration)
    resource, resp, err := apiClient.DataPlatformClusterApi.PatchCluster(context.Background(), clusterId).PatchClusterRequest(patchClusterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataPlatformClusterApi.PatchCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `PatchCluster`: ClusterResponseData
    fmt.Fprintf(os.Stdout, "Response from `DataPlatformClusterApi.PatchCluster`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | [**string**](../models/.md) | The unique ID of the cluster. Must conform to the UUID format.  | |

### Other Parameters

Other parameters are passed through a pointer to an apiPatchClusterRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **patchClusterRequest** | [**PatchClusterRequest**](../models/PatchClusterRequest.md) | Request payload with the properties that shall be applied to an existing DataPlatformCluster.  | |

### Return type

[**ClusterResponseData**](../models/ClusterResponseData.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


