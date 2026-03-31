# \BackupsApi

All URIs are relative to *https://postgresql.de-txl.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**BackupsFindById**](BackupsApi.md#BackupsFindById) | **Get** /backups/{backupId} | Retrieve Backup|
|[**BackupsGet**](BackupsApi.md#BackupsGet) | **Get** /backups | Retrieve all Backups|



## BackupsFindById

```go
var result BackupRead = BackupsFindById(ctx, backupId)
                      .Execute()
```

Retrieve Backup



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    psql "github.com/ionos-cloud/sdk-go-bundle/products/psql"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    backupId := "45ca67fb-8b07-5783-9c97-2d35acceb084" // string | The ID (UUID) of the Backup.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := psql.NewAPIClient(configuration)
    resource, resp, err := apiClient.BackupsApi.BackupsFindById(context.Background(), backupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.BackupsFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `BackupsFindById`: BackupRead
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.BackupsFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**backupId** | **string** | The ID (UUID) of the Backup. | |

### Other Parameters

Other parameters are passed through a pointer to an apiBackupsFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**BackupRead**](../models/BackupRead.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## BackupsGet

```go
var result BackupReadList = BackupsGet(ctx)
                      .Offset(offset)
                      .Limit(limit)
                      .FilterClusterId(filterClusterId)
                      .Execute()
```

Retrieve all Backups



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    psql "github.com/ionos-cloud/sdk-go-bundle/products/psql"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    offset := int32(0) // int32 | The first element (of the total list of elements) to include in the response. Use this parameter together with the limit for pagination. (optional) (default to 0)
    limit := int32(100) // int32 | The maximum number of elements to return. Use this parameter together with the offset for pagination. (optional) (default to 100)
    filterClusterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter backups by cluster Id. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := psql.NewAPIClient(configuration)
    resource, resp, err := apiClient.BackupsApi.BackupsGet(context.Background()).Offset(offset).Limit(limit).FilterClusterId(filterClusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.BackupsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `BackupsGet`: BackupReadList
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.BackupsGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiBackupsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **offset** | **int32** | The first element (of the total list of elements) to include in the response. Use this parameter together with the limit for pagination. | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return. Use this parameter together with the offset for pagination. | [default to 100]|
| **filterClusterId** | **string** | Filter backups by cluster Id. | |

### Return type

[**BackupReadList**](../models/BackupReadList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


