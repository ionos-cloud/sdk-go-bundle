# \BackupLocationsApi

All URIs are relative to *https://mariadb.de-txl.ionos.com/v2*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**BackuplocationsFindById**](BackupLocationsApi.md#BackuplocationsFindById) | **Get** /backup-locations/{backupLocationId} | Retrieve BackupLocation|
|[**BackuplocationsGet**](BackupLocationsApi.md#BackuplocationsGet) | **Get** /backup-locations | Retrieve all BackupLocations|



## BackuplocationsFindById

```go
var result BackupLocationRead = BackuplocationsFindById(ctx, backupLocationId)
                      .Execute()
```

Retrieve BackupLocation



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    mariadb "github.com/ionos-cloud/sdk-go-bundle/products/mariadb"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    backupLocationId := "7fa1dd11-59dd-53a5-ab67-50f649c8e3eb" // string | The ID (UUID) of the BackupLocation.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := mariadb.NewAPIClient(configuration)
    resource, resp, err := apiClient.BackupLocationsApi.BackuplocationsFindById(context.Background(), backupLocationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupLocationsApi.BackuplocationsFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `BackuplocationsFindById`: BackupLocationRead
    fmt.Fprintf(os.Stdout, "Response from `BackupLocationsApi.BackuplocationsFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**backupLocationId** | **string** | The ID (UUID) of the BackupLocation. | |

### Other Parameters

Other parameters are passed through a pointer to an apiBackuplocationsFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**BackupLocationRead**](../models/BackupLocationRead.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## BackuplocationsGet

```go
var result BackupLocationReadList = BackuplocationsGet(ctx)
                      .Offset(offset)
                      .Limit(limit)
                      .Execute()
```

Retrieve all BackupLocations



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    mariadb "github.com/ionos-cloud/sdk-go-bundle/products/mariadb"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    offset := int32(0) // int32 | The first element (of the total list of elements) to include in the response. Use this parameter together with the limit for pagination. (optional) (default to 0)
    limit := int32(100) // int32 | The maximum number of elements to return. Use this parameter together with the offset for pagination. (optional) (default to 100)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := mariadb.NewAPIClient(configuration)
    resource, resp, err := apiClient.BackupLocationsApi.BackuplocationsGet(context.Background()).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupLocationsApi.BackuplocationsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `BackuplocationsGet`: BackupLocationReadList
    fmt.Fprintf(os.Stdout, "Response from `BackupLocationsApi.BackuplocationsGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiBackuplocationsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **offset** | **int32** | The first element (of the total list of elements) to include in the response. Use this parameter together with the limit for pagination. | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return. Use this parameter together with the offset for pagination. | [default to 100]|

### Return type

[**BackupLocationReadList**](../models/BackupLocationReadList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


