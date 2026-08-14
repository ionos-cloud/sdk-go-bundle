# \SnapshotLocationsApi

All URIs are relative to *https://in-memory-db.de-fra.ionos.com/v2*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**SnapshotlocationsFindById**](SnapshotLocationsApi.md#SnapshotlocationsFindById) | **Get** /snapshot-locations/{snapshotLocationId} | Retrieve SnapshotLocation|
|[**SnapshotlocationsGet**](SnapshotLocationsApi.md#SnapshotlocationsGet) | **Get** /snapshot-locations | Retrieve all SnapshotLocations|



## SnapshotlocationsFindById

```go
var result SnapshotLocationRead = SnapshotlocationsFindById(ctx, snapshotLocationId)
                      .Execute()
```

Retrieve SnapshotLocation



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
    snapshotLocationId := "caac592a-83a4-55d9-aee1-b43e78a098d2" // string | The ID (UUID) of the SnapshotLocation.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := inmemorydb.NewAPIClient(configuration)
    resource, resp, err := apiClient.SnapshotLocationsApi.SnapshotlocationsFindById(context.Background(), snapshotLocationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotLocationsApi.SnapshotlocationsFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `SnapshotlocationsFindById`: SnapshotLocationRead
    fmt.Fprintf(os.Stdout, "Response from `SnapshotLocationsApi.SnapshotlocationsFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**snapshotLocationId** | **string** | The ID (UUID) of the SnapshotLocation. | |

### Other Parameters

Other parameters are passed through a pointer to an apiSnapshotlocationsFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**SnapshotLocationRead**](../models/SnapshotLocationRead.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## SnapshotlocationsGet

```go
var result SnapshotLocationReadList = SnapshotlocationsGet(ctx)
                      .Offset(offset)
                      .Limit(limit)
                      .Execute()
```

Retrieve all SnapshotLocations



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

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := inmemorydb.NewAPIClient(configuration)
    resource, resp, err := apiClient.SnapshotLocationsApi.SnapshotlocationsGet(context.Background()).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotLocationsApi.SnapshotlocationsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `SnapshotlocationsGet`: SnapshotLocationReadList
    fmt.Fprintf(os.Stdout, "Response from `SnapshotLocationsApi.SnapshotlocationsGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiSnapshotlocationsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **offset** | **int32** | The first element (of the total list of elements) to include in the response. Use this parameter together with the limit for pagination. | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return. Use this parameter together with the offset for pagination. | [default to 100]|

### Return type

[**SnapshotLocationReadList**](../models/SnapshotLocationReadList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


