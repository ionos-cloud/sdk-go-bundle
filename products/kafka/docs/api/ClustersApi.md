# \ClustersApi

All URIs are relative to *https://kafka.de-fra.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**ClustersDelete**](ClustersApi.md#ClustersDelete) | **Delete** /clusters/{clusterId} | Delete Cluster|
|[**ClustersFindById**](ClustersApi.md#ClustersFindById) | **Get** /clusters/{clusterId} | Retrieve Cluster|
|[**ClustersGet**](ClustersApi.md#ClustersGet) | **Get** /clusters | Retrieve all Clusters|
|[**ClustersPost**](ClustersApi.md#ClustersPost) | **Post** /clusters | Create Cluster|



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

    kafka "github.com/ionos-cloud/sdk-go-bundle/products/kafka"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    clusterId := "e69b22a5-8fee-56b1-b6fb-4a07e4205ead" // string | The ID (UUID) of the Cluster.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := kafka.NewAPIClient(configuration)
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

    kafka "github.com/ionos-cloud/sdk-go-bundle/products/kafka"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    clusterId := "e69b22a5-8fee-56b1-b6fb-4a07e4205ead" // string | The ID (UUID) of the Cluster.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := kafka.NewAPIClient(configuration)
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

    kafka "github.com/ionos-cloud/sdk-go-bundle/products/kafka"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    filterName := "filterName_example" // string | Only return Kafka clusters that contain the given name. Case insensitive. (optional)
    filterState := "filterState_example" // string | Only return Kafka clusters with a given state. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := kafka.NewAPIClient(configuration)
    resource, resp, err := apiClient.ClustersApi.ClustersGet(context.Background()).FilterName(filterName).FilterState(filterState).Execute()
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
| **filterName** | **string** | Only return Kafka clusters that contain the given name. Case insensitive. | |
| **filterState** | **string** | Only return Kafka clusters with a given state. | |

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

    kafka "github.com/ionos-cloud/sdk-go-bundle/products/kafka"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    clusterCreate := *openapiclient.NewClusterCreate(*openapiclient.NewCluster("my-kafka-cluster", "3.7.0", "XS", []openapiclient.KafkaClusterConnection{*openapiclient.NewKafkaClusterConnection("5a029f4a-72e5-11ec-90d6-0242ac120003", "2", []string{"192.168.1.101/24"})})) // ClusterCreate | Cluster to create.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := kafka.NewAPIClient(configuration)
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


