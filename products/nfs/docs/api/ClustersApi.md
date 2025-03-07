# \ClustersApi

All URIs are relative to *https://nfs.de-fra.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**ClustersDelete**](ClustersApi.md#ClustersDelete) | **Delete** /clusters/{clusterId} | Delete Cluster|
|[**ClustersFindById**](ClustersApi.md#ClustersFindById) | **Get** /clusters/{clusterId} | Retrieve Cluster|
|[**ClustersGet**](ClustersApi.md#ClustersGet) | **Get** /clusters | Retrieve Clusters|
|[**ClustersPost**](ClustersApi.md#ClustersPost) | **Post** /clusters | Create Cluster|
|[**ClustersPut**](ClustersApi.md#ClustersPut) | **Put** /clusters/{clusterId} | Ensure Cluster|



## ClustersDelete

```go
var result  = ClustersDelete(ctx, clusterId)
                      .Execute()
```

Delete Cluster



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    nfs "github.com/ionos-cloud/sdk-go-bundle/products/nfs"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    clusterId := "e69b22a5-8fee-56b1-b6fb-4a07e4205ead" // string | The identifier (UUID) of the cluster.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := nfs.NewAPIClient(configuration)
    resp, err := apiClient.ClustersApi.ClustersDelete(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The identifier (UUID) of the cluster. | |

### Other Parameters

Other parameters are passed through a pointer to an apiClustersDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ClustersFindById

```go
var result ClusterRead = ClustersFindById(ctx, clusterId)
                      .Execute()
```

Retrieve Cluster



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    nfs "github.com/ionos-cloud/sdk-go-bundle/products/nfs"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    clusterId := "e69b22a5-8fee-56b1-b6fb-4a07e4205ead" // string | The identifier (UUID) of the cluster.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := nfs.NewAPIClient(configuration)
    resource, resp, err := apiClient.ClustersApi.ClustersFindById(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersFindById`: ClusterRead
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ClustersFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The identifier (UUID) of the cluster. | |

### Other Parameters

Other parameters are passed through a pointer to an apiClustersFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**ClusterRead**](../models/ClusterRead.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ClustersGet

```go
var result ClusterReadList = ClustersGet(ctx)
                      .Offset(offset)
                      .Limit(limit)
                      .FilterDatacenterId(filterDatacenterId)
                      .Execute()
```

Retrieve Clusters



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    nfs "github.com/ionos-cloud/sdk-go-bundle/products/nfs"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    offset := int32(0) // int32 | The first element from the total list of elements to include in the response. Use this parameter together with the limit for pagination. (optional) (default to 0)
    limit := int32(100) // int32 | The maximum number of elements to return. Use this parameter together with the offset for pagination. (optional) (default to 100)
    filterDatacenterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The datacenter identifier to filter by. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := nfs.NewAPIClient(configuration)
    resource, resp, err := apiClient.ClustersApi.ClustersGet(context.Background()).Offset(offset).Limit(limit).FilterDatacenterId(filterDatacenterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersGet`: ClusterReadList
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ClustersGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiClustersGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **offset** | **int32** | The first element from the total list of elements to include in the response. Use this parameter together with the limit for pagination. | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return. Use this parameter together with the offset for pagination. | [default to 100]|
| **filterDatacenterId** | **string** | The datacenter identifier to filter by. | |

### Return type

[**ClusterReadList**](../models/ClusterReadList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ClustersPost

```go
var result ClusterRead = ClustersPost(ctx)
                      .ClusterCreate(clusterCreate)
                      .Execute()
```

Create Cluster



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    nfs "github.com/ionos-cloud/sdk-go-bundle/products/nfs"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    clusterCreate := *openapiclient.NewClusterCreate(*openapiclient.NewCluster("Cluster 1", []openapiclient.ClusterConnections{*openapiclient.NewClusterConnections("123e4567-e89b-12d3-a456-426614174001", "1", "10.254.64.1/24")})) // ClusterCreate | Cluster to create.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := nfs.NewAPIClient(configuration)
    resource, resp, err := apiClient.ClustersApi.ClustersPost(context.Background()).ClusterCreate(clusterCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersPost`: ClusterRead
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ClustersPost`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiClustersPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **clusterCreate** | [**ClusterCreate**](../models/ClusterCreate.md) | Cluster to create. | |

### Return type

[**ClusterRead**](../models/ClusterRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## ClustersPut

```go
var result ClusterRead = ClustersPut(ctx, clusterId)
                      .ClusterEnsure(clusterEnsure)
                      .Execute()
```

Ensure Cluster



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    nfs "github.com/ionos-cloud/sdk-go-bundle/products/nfs"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    clusterId := "e69b22a5-8fee-56b1-b6fb-4a07e4205ead" // string | The ID (UUID) of the Cluster.
    clusterEnsure := *openapiclient.NewClusterEnsure("e69b22a5-8fee-56b1-b6fb-4a07e4205ead", *openapiclient.NewCluster("Cluster 1", []openapiclient.ClusterConnections{*openapiclient.NewClusterConnections("123e4567-e89b-12d3-a456-426614174001", "1", "10.254.64.1/24")})) // ClusterEnsure | Update Cluster

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := nfs.NewAPIClient(configuration)
    resource, resp, err := apiClient.ClustersApi.ClustersPut(context.Background(), clusterId).ClusterEnsure(clusterEnsure).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersPut`: ClusterRead
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ClustersPut`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The ID (UUID) of the Cluster. | |

### Other Parameters

Other parameters are passed through a pointer to an apiClustersPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **clusterEnsure** | [**ClusterEnsure**](../models/ClusterEnsure.md) | Update Cluster | |

### Return type

[**ClusterRead**](../models/ClusterRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


