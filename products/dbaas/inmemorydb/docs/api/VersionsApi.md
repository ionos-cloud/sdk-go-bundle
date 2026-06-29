# \VersionsApi

All URIs are relative to *https://in-memory-db.de-fra.ionos.com/v2*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**VersionsFindById**](VersionsApi.md#VersionsFindById) | **Get** /versions/{versionId} | Retrieve Version|
|[**VersionsGet**](VersionsApi.md#VersionsGet) | **Get** /versions | Retrieve all Versions|



## VersionsFindById

```go
var result VersionRead = VersionsFindById(ctx, versionId)
                      .Execute()
```

Retrieve Version



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
    versionId := "6b901215-a751-5580-a190-a17cc5e0d292" // string | The ID (UUID) of the Version.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := inmemorydb.NewAPIClient(configuration)
    resource, resp, err := apiClient.VersionsApi.VersionsFindById(context.Background(), versionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsApi.VersionsFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `VersionsFindById`: VersionRead
    fmt.Fprintf(os.Stdout, "Response from `VersionsApi.VersionsFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**versionId** | **string** | The ID (UUID) of the Version. | |

### Other Parameters

Other parameters are passed through a pointer to an apiVersionsFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**VersionRead**](../models/VersionRead.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## VersionsGet

```go
var result VersionReadList = VersionsGet(ctx)
                      .Offset(offset)
                      .Limit(limit)
                      .Execute()
```

Retrieve all Versions



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
    resource, resp, err := apiClient.VersionsApi.VersionsGet(context.Background()).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsApi.VersionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `VersionsGet`: VersionReadList
    fmt.Fprintf(os.Stdout, "Response from `VersionsApi.VersionsGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiVersionsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **offset** | **int32** | The first element (of the total list of elements) to include in the response. Use this parameter together with the limit for pagination. | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return. Use this parameter together with the offset for pagination. | [default to 100]|

### Return type

[**VersionReadList**](../models/VersionReadList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


