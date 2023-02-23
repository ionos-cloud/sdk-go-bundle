# \DataPlatformNodePoolApi

All URIs are relative to *https://api.ionos.com/dataplatform*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**CreateClusterNodepool**](DataPlatformNodePoolApi.md#CreateClusterNodepool) | **Post** /clusters/{clusterId}/nodepools | Create a DataPlatformNodePool for a distinct DataPlatformCluster|
|[**DeleteClusterNodepool**](DataPlatformNodePoolApi.md#DeleteClusterNodepool) | **Delete** /clusters/{clusterId}/nodepools/{nodepoolId} | Remove node pool from DataPlatformCluster.|
|[**GetClusterNodepool**](DataPlatformNodePoolApi.md#GetClusterNodepool) | **Get** /clusters/{clusterId}/nodepools/{nodepoolId} | Retrieve a DataPlatformNodePool|
|[**GetClusterNodepools**](DataPlatformNodePoolApi.md#GetClusterNodepools) | **Get** /clusters/{clusterId}/nodepools | List the DataPlatformNodePools of a  DataPlatformCluster|
|[**PatchClusterNodepool**](DataPlatformNodePoolApi.md#PatchClusterNodepool) | **Patch** /clusters/{clusterId}/nodepools/{nodepoolId} | Partially modify a DataPlatformNodePool|



## CreateClusterNodepool

```go
var result NodePoolResponseData = CreateClusterNodepool(ctx, clusterId)
                      .CreateNodePoolRequest(createNodePoolRequest)
                      .Execute()
```

Create a DataPlatformNodePool for a distinct DataPlatformCluster



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
    createNodePoolRequest := *openapiclient.NewCreateNodePoolRequest(*openapiclient.NewCreateNodePoolProperties("my-node-pool", int32(2))) // CreateNodePoolRequest | Request payload with the properties that defines a DataPlatformNodePool. 

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := dataplatform.NewAPIClient(configuration)
    resource, resp, err := apiClient.DataPlatformNodePoolApi.CreateClusterNodepool(context.Background(), clusterId).CreateNodePoolRequest(createNodePoolRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataPlatformNodePoolApi.CreateClusterNodepool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `CreateClusterNodepool`: NodePoolResponseData
    fmt.Fprintf(os.Stdout, "Response from `DataPlatformNodePoolApi.CreateClusterNodepool`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | [**string**](../models/.md) | The unique ID of the cluster. Must conform to the UUID format.  | |

### Other Parameters

Other parameters are passed through a pointer to an apiCreateClusterNodepoolRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **createNodePoolRequest** | [**CreateNodePoolRequest**](../models/CreateNodePoolRequest.md) | Request payload with the properties that defines a DataPlatformNodePool.  | |

### Return type

[**NodePoolResponseData**](../models/NodePoolResponseData.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DeleteClusterNodepool

```go
var result NodePoolResponseData = DeleteClusterNodepool(ctx, clusterId, nodepoolId)
                      .Execute()
```

Remove node pool from DataPlatformCluster.



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
    nodepoolId := TODO // string | The unique ID of the node pool. Must conform to the UUID format. 

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := dataplatform.NewAPIClient(configuration)
    resource, resp, err := apiClient.DataPlatformNodePoolApi.DeleteClusterNodepool(context.Background(), clusterId, nodepoolId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataPlatformNodePoolApi.DeleteClusterNodepool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `DeleteClusterNodepool`: NodePoolResponseData
    fmt.Fprintf(os.Stdout, "Response from `DataPlatformNodePoolApi.DeleteClusterNodepool`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | [**string**](../models/.md) | The unique ID of the cluster. Must conform to the UUID format.  | |
|**nodepoolId** | [**string**](../models/.md) | The unique ID of the node pool. Must conform to the UUID format.  | |

### Other Parameters

Other parameters are passed through a pointer to an apiDeleteClusterNodepoolRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**NodePoolResponseData**](../models/NodePoolResponseData.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## GetClusterNodepool

```go
var result NodePoolResponseData = GetClusterNodepool(ctx, clusterId, nodepoolId)
                      .Execute()
```

Retrieve a DataPlatformNodePool



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
    nodepoolId := TODO // string | The unique ID of the node pool. Must conform to the UUID format. 

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := dataplatform.NewAPIClient(configuration)
    resource, resp, err := apiClient.DataPlatformNodePoolApi.GetClusterNodepool(context.Background(), clusterId, nodepoolId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataPlatformNodePoolApi.GetClusterNodepool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `GetClusterNodepool`: NodePoolResponseData
    fmt.Fprintf(os.Stdout, "Response from `DataPlatformNodePoolApi.GetClusterNodepool`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | [**string**](../models/.md) | The unique ID of the cluster. Must conform to the UUID format.  | |
|**nodepoolId** | [**string**](../models/.md) | The unique ID of the node pool. Must conform to the UUID format.  | |

### Other Parameters

Other parameters are passed through a pointer to an apiGetClusterNodepoolRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**NodePoolResponseData**](../models/NodePoolResponseData.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## GetClusterNodepools

```go
var result NodePoolListResponseData = GetClusterNodepools(ctx, clusterId)
                      .Execute()
```

List the DataPlatformNodePools of a  DataPlatformCluster



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
    resource, resp, err := apiClient.DataPlatformNodePoolApi.GetClusterNodepools(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataPlatformNodePoolApi.GetClusterNodepools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `GetClusterNodepools`: NodePoolListResponseData
    fmt.Fprintf(os.Stdout, "Response from `DataPlatformNodePoolApi.GetClusterNodepools`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | [**string**](../models/.md) | The unique ID of the cluster. Must conform to the UUID format.  | |

### Other Parameters

Other parameters are passed through a pointer to an apiGetClusterNodepoolsRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**NodePoolListResponseData**](../models/NodePoolListResponseData.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## PatchClusterNodepool

```go
var result NodePoolResponseData = PatchClusterNodepool(ctx, clusterId, nodepoolId)
                      .PatchNodePoolRequest(patchNodePoolRequest)
                      .Execute()
```

Partially modify a DataPlatformNodePool



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
    nodepoolId := TODO // string | The unique ID of the node pool. Must conform to the UUID format. 
    patchNodePoolRequest := *openapiclient.NewPatchNodePoolRequest(*openapiclient.NewPatchNodePoolProperties()) // PatchNodePoolRequest | Request payload with the properties that shall be applied to an existing DataPlatformNodePool. 

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := dataplatform.NewAPIClient(configuration)
    resource, resp, err := apiClient.DataPlatformNodePoolApi.PatchClusterNodepool(context.Background(), clusterId, nodepoolId).PatchNodePoolRequest(patchNodePoolRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataPlatformNodePoolApi.PatchClusterNodepool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `PatchClusterNodepool`: NodePoolResponseData
    fmt.Fprintf(os.Stdout, "Response from `DataPlatformNodePoolApi.PatchClusterNodepool`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | [**string**](../models/.md) | The unique ID of the cluster. Must conform to the UUID format.  | |
|**nodepoolId** | [**string**](../models/.md) | The unique ID of the node pool. Must conform to the UUID format.  | |

### Other Parameters

Other parameters are passed through a pointer to an apiPatchClusterNodepoolRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **patchNodePoolRequest** | [**PatchNodePoolRequest**](../models/PatchNodePoolRequest.md) | Request payload with the properties that shall be applied to an existing DataPlatformNodePool.  | |

### Return type

[**NodePoolResponseData**](../models/NodePoolResponseData.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


