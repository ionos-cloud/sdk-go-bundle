# \ReplicaSetApi

All URIs are relative to *https://in-memory-db.de-fra.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**ReplicasetsDelete**](ReplicaSetApi.md#ReplicasetsDelete) | **Delete** /replicasets/{replicasetId} | Delete ReplicaSet|
|[**ReplicasetsFindById**](ReplicaSetApi.md#ReplicasetsFindById) | **Get** /replicasets/{replicasetId} | Retrieve ReplicaSet|
|[**ReplicasetsGet**](ReplicaSetApi.md#ReplicasetsGet) | **Get** /replicasets | Retrieve all ReplicaSet|
|[**ReplicasetsPost**](ReplicaSetApi.md#ReplicasetsPost) | **Post** /replicasets | Create ReplicaSet|
|[**ReplicasetsPut**](ReplicaSetApi.md#ReplicasetsPut) | **Put** /replicasets/{replicasetId} | Ensure ReplicaSet|



## ReplicasetsDelete

```go
var result  = ReplicasetsDelete(ctx, replicasetId)
                      .Execute()
```

Delete ReplicaSet



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
    replicasetId := "1046e9bf-dbc0-5bd3-9291-713d36ab77e9" // string | The ID (UUID) of the ReplicaSet.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := inmemorydb.NewAPIClient(configuration)
    resp, err := apiClient.ReplicaSetApi.ReplicasetsDelete(context.Background(), replicasetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicaSetApi.ReplicasetsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**replicasetId** | **string** | The ID (UUID) of the ReplicaSet. | |

### Other Parameters

Other parameters are passed through a pointer to an apiReplicasetsDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ReplicasetsFindById

```go
var result ReplicaSetRead = ReplicasetsFindById(ctx, replicasetId)
                      .Execute()
```

Retrieve ReplicaSet



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
    replicasetId := "1046e9bf-dbc0-5bd3-9291-713d36ab77e9" // string | The ID (UUID) of the ReplicaSet.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := inmemorydb.NewAPIClient(configuration)
    resource, resp, err := apiClient.ReplicaSetApi.ReplicasetsFindById(context.Background(), replicasetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicaSetApi.ReplicasetsFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ReplicasetsFindById`: ReplicaSetRead
    fmt.Fprintf(os.Stdout, "Response from `ReplicaSetApi.ReplicasetsFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**replicasetId** | **string** | The ID (UUID) of the ReplicaSet. | |

### Other Parameters

Other parameters are passed through a pointer to an apiReplicasetsFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**ReplicaSetRead**](../models/ReplicaSetRead.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ReplicasetsGet

```go
var result ReplicaSetReadList = ReplicasetsGet(ctx)
                      .Offset(offset)
                      .Limit(limit)
                      .FilterName(filterName)
                      .Execute()
```

Retrieve all ReplicaSet



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
    offset := int32(0) // int32 | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. (optional) (default to 0)
    limit := int32(100) // int32 | The maximum number of elements to return. Use together with offset for pagination. (optional) (default to 100)
    filterName := "filterName_example" // string | Response filter to list only items contain the specified name. The value is case insensitive and matched on the 'displayName' field.  (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := inmemorydb.NewAPIClient(configuration)
    resource, resp, err := apiClient.ReplicaSetApi.ReplicasetsGet(context.Background()).Offset(offset).Limit(limit).FilterName(filterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicaSetApi.ReplicasetsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ReplicasetsGet`: ReplicaSetReadList
    fmt.Fprintf(os.Stdout, "Response from `ReplicaSetApi.ReplicasetsGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiReplicasetsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **offset** | **int32** | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return. Use together with offset for pagination. | [default to 100]|
| **filterName** | **string** | Response filter to list only items contain the specified name. The value is case insensitive and matched on the &#39;displayName&#39; field.  | |

### Return type

[**ReplicaSetReadList**](../models/ReplicaSetReadList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ReplicasetsPost

```go
var result ReplicaSetRead = ReplicasetsPost(ctx)
                      .ReplicaSetCreate(replicaSetCreate)
                      .Execute()
```

Create ReplicaSet



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
    replicaSetCreate := *openapiclient.NewReplicaSetCreate(*openapiclient.NewReplicaSet("In-Memory DB replica set", "7.2", int32(2), *openapiclient.NewResources(int32(4), int32(4)), openapiclient.PersistenceMode("None"), openapiclient.EvictionPolicy("noeviction"), []openapiclient.Connection{*openapiclient.NewConnection("5a029f4a-72e5-11ec-90d6-0242ac120003", "2", "192.168.1.100/24")}, *openapiclient.NewUser("DatabaseAdmin"))) // ReplicaSetCreate | ReplicaSet to create.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := inmemorydb.NewAPIClient(configuration)
    resource, resp, err := apiClient.ReplicaSetApi.ReplicasetsPost(context.Background()).ReplicaSetCreate(replicaSetCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicaSetApi.ReplicasetsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ReplicasetsPost`: ReplicaSetRead
    fmt.Fprintf(os.Stdout, "Response from `ReplicaSetApi.ReplicasetsPost`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiReplicasetsPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **replicaSetCreate** | [**ReplicaSetCreate**](../models/ReplicaSetCreate.md) | ReplicaSet to create. | |

### Return type

[**ReplicaSetRead**](../models/ReplicaSetRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## ReplicasetsPut

```go
var result ReplicaSetRead = ReplicasetsPut(ctx, replicasetId)
                      .ReplicaSetEnsure(replicaSetEnsure)
                      .Execute()
```

Ensure ReplicaSet



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
    replicasetId := "1046e9bf-dbc0-5bd3-9291-713d36ab77e9" // string | The ID (UUID) of the ReplicaSet.
    replicaSetEnsure := *openapiclient.NewReplicaSetEnsure("1046e9bf-dbc0-5bd3-9291-713d36ab77e9", *openapiclient.NewReplicaSet("In-Memory DB replica set", "7.2", int32(2), *openapiclient.NewResources(int32(4), int32(4)), openapiclient.PersistenceMode("None"), openapiclient.EvictionPolicy("noeviction"), []openapiclient.Connection{*openapiclient.NewConnection("5a029f4a-72e5-11ec-90d6-0242ac120003", "2", "192.168.1.100/24")}, *openapiclient.NewUser("DatabaseAdmin"))) // ReplicaSetEnsure | update ReplicaSet

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := inmemorydb.NewAPIClient(configuration)
    resource, resp, err := apiClient.ReplicaSetApi.ReplicasetsPut(context.Background(), replicasetId).ReplicaSetEnsure(replicaSetEnsure).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicaSetApi.ReplicasetsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ReplicasetsPut`: ReplicaSetRead
    fmt.Fprintf(os.Stdout, "Response from `ReplicaSetApi.ReplicasetsPut`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**replicasetId** | **string** | The ID (UUID) of the ReplicaSet. | |

### Other Parameters

Other parameters are passed through a pointer to an apiReplicasetsPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **replicaSetEnsure** | [**ReplicaSetEnsure**](../models/ReplicaSetEnsure.md) | update ReplicaSet | |

### Return type

[**ReplicaSetRead**](../models/ReplicaSetRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


