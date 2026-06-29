# \ClustersApi

All URIs are relative to *https://in-memory-db.de-fra.ionos.com/v2*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**ClustersDelete**](ClustersApi.md#ClustersDelete) | **Delete** /clusters/{clusterId} | Delete Cluster|
|[**ClustersFindById**](ClustersApi.md#ClustersFindById) | **Get** /clusters/{clusterId} | Retrieve Cluster|
|[**ClustersGet**](ClustersApi.md#ClustersGet) | **Get** /clusters | Retrieve all Clusters|
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

    inmemorydb "github.com/ionos-cloud/sdk-go-bundle/products/inmemorydb"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    clusterId := "e69b22a5-8fee-56b1-b6fb-4a07e4205ead" // string | The ID (UUID) of the Cluster.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := inmemorydb.NewAPIClient(configuration)
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
|**clusterId** | **string** | The ID (UUID) of the Cluster. | |

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

    inmemorydb "github.com/ionos-cloud/sdk-go-bundle/products/inmemorydb"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    clusterId := "e69b22a5-8fee-56b1-b6fb-4a07e4205ead" // string | The ID (UUID) of the Cluster.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := inmemorydb.NewAPIClient(configuration)
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
|**clusterId** | **string** | The ID (UUID) of the Cluster. | |

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
                      .FilterName(filterName)
                      .FilterState(filterState)
                      .Execute()
```

Retrieve all Clusters



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    inmemorydb "github.com/ionos-cloud/sdk-go-bundle/products/inmemorydb"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    offset := int32(0) // int32 | The first element (of the total list of elements) to include in the response. Use this parameter together with the limit for pagination. (optional) (default to 0)
    limit := int32(100) // int32 | The maximum number of elements to return. Use this parameter together with the offset for pagination. (optional) (default to 100)
    filterName := "filterName_example" // string | Filters clusters by name. Matches cluster names that contain the provided string (case-insensitive). (optional)
    filterState := openapiclient.ClusterState("PROVISIONING") // ClusterState | Filters clusters by their current lifecycle state. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := inmemorydb.NewAPIClient(configuration)
    resource, resp, err := apiClient.ClustersApi.ClustersGet(context.Background()).Offset(offset).Limit(limit).FilterName(filterName).FilterState(filterState).Execute()
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
| **offset** | **int32** | The first element (of the total list of elements) to include in the response. Use this parameter together with the limit for pagination. | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return. Use this parameter together with the offset for pagination. | [default to 100]|
| **filterName** | **string** | Filters clusters by name. Matches cluster names that contain the provided string (case-insensitive). | |
| **filterState** | [**ClusterState**](../models/.md) | Filters clusters by their current lifecycle state. | |

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

    inmemorydb "github.com/ionos-cloud/sdk-go-bundle/products/inmemorydb"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    clusterCreate := *openapiclient.NewClusterCreate(*openapiclient.NewClusterCreateProperties("my-cluster", "9.0", *openapiclient.NewInstanceConfiguration(int32(1), int32(4), int32(8)), *openapiclient.NewClusterConnection("5a029f4a-72e5-11ec-90d6-0242ac120003", "2", "192.168.2.101/24"), openapiclient.EvictionPolicy("noeviction"), *openapiclient.NewSnapshotConfiguration("eu-central-3", int32(14), []int32{int32(123)}), *openapiclient.NewMaintenanceWindow("16:30:59", openapiclient.DayOfTheWeek("Sunday")), *openapiclient.NewClusterCredentials("cache_admin", *openapiclient.NewHashedPassword("SHA-256", "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8")))) // ClusterCreate | Cluster to create.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := inmemorydb.NewAPIClient(configuration)
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

    inmemorydb "github.com/ionos-cloud/sdk-go-bundle/products/inmemorydb"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    clusterId := "e69b22a5-8fee-56b1-b6fb-4a07e4205ead" // string | The ID (UUID) of the Cluster.
    clusterEnsure := *openapiclient.NewClusterEnsure("e69b22a5-8fee-56b1-b6fb-4a07e4205ead", *openapiclient.NewCluster("my-cluster", "9.0", *openapiclient.NewInstanceConfiguration(int32(1), int32(4), int32(8)), *openapiclient.NewClusterConnection("5a029f4a-72e5-11ec-90d6-0242ac120003", "2", "192.168.2.101/24"), openapiclient.EvictionPolicy("noeviction"), *openapiclient.NewSnapshotConfiguration("eu-central-3", int32(14), []int32{int32(123)}), *openapiclient.NewMaintenanceWindow("16:30:59", openapiclient.DayOfTheWeek("Sunday")))) // ClusterEnsure | update Cluster

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := inmemorydb.NewAPIClient(configuration)
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
| **clusterEnsure** | [**ClusterEnsure**](../models/ClusterEnsure.md) | update Cluster | |

### Return type

[**ClusterRead**](../models/ClusterRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


