# \SnapshotApi

All URIs are relative to *https://in-memory-db.de-fra.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**SnapshotsFindById**](SnapshotApi.md#SnapshotsFindById) | **Get** /snapshots/{snapshotId} | Retrieve Snapshot|
|[**SnapshotsGet**](SnapshotApi.md#SnapshotsGet) | **Get** /snapshots | Retrieve all Snapshot|



## SnapshotsFindById

```go
var result SnapshotRead = SnapshotsFindById(ctx, snapshotId)
                      .Execute()
```

Retrieve Snapshot



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
    snapshotId := "a8784665-3d99-5464-af32-30a2967f58e7" // string | The ID (UUID) of the Snapshot.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := inmemorydb.NewAPIClient(configuration)
    resource, resp, err := apiClient.SnapshotApi.SnapshotsFindById(context.Background(), snapshotId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotApi.SnapshotsFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `SnapshotsFindById`: SnapshotRead
    fmt.Fprintf(os.Stdout, "Response from `SnapshotApi.SnapshotsFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**snapshotId** | **string** | The ID (UUID) of the Snapshot. | |

### Other Parameters

Other parameters are passed through a pointer to an apiSnapshotsFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**SnapshotRead**](../models/SnapshotRead.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## SnapshotsGet

```go
var result SnapshotReadList = SnapshotsGet(ctx)
                      .Offset(offset)
                      .Limit(limit)
                      .Execute()
```

Retrieve all Snapshot



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

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := inmemorydb.NewAPIClient(configuration)
    resource, resp, err := apiClient.SnapshotApi.SnapshotsGet(context.Background()).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotApi.SnapshotsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `SnapshotsGet`: SnapshotReadList
    fmt.Fprintf(os.Stdout, "Response from `SnapshotApi.SnapshotsGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiSnapshotsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **offset** | **int32** | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return. Use together with offset for pagination. | [default to 100]|

### Return type

[**SnapshotReadList**](../models/SnapshotReadList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


