# \BackupsApi

All URIs are relative to *https://api.ionos.com/databases/postgresql*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**ClusterBackupsGet**](BackupsApi.md#ClusterBackupsGet) | **Get** /clusters/{clusterId}/backups | List backups of cluster|
|[**ClustersBackupsFindById**](BackupsApi.md#ClustersBackupsFindById) | **Get** /clusters/backups/{backupId} | Fetch a cluster backup|
|[**ClustersBackupsGet**](BackupsApi.md#ClustersBackupsGet) | **Get** /clusters/backups | List cluster backups|



## ClusterBackupsGet

```go
var result ClusterBackupList = ClusterBackupsGet(ctx, clusterId)
                      .Limit(limit)
                      .Offset(offset)
                      .Execute()
```

List backups of cluster



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    psql "github.com/ionos-cloud/sdk-go-bundle/products/dbaas/psql"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    clusterId := "498ae72f-411f-11eb-9d07-046c59cc737e" // string | The unique ID of the cluster.
    limit := int32(100) // int32 | The maximum number of elements to return. Use together with 'offset' for pagination. (optional) (default to 100)
    offset := int32(200) // int32 | The first element to return. Use together with 'limit' for pagination. (optional) (default to 0)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := psql.NewAPIClient(configuration)
    resource, resp, err := apiClient.BackupsApi.ClusterBackupsGet(context.Background(), clusterId).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.ClusterBackupsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClusterBackupsGet`: ClusterBackupList
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.ClusterBackupsGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The unique ID of the cluster. | |

### Other Parameters

Other parameters are passed through a pointer to an apiClusterBackupsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **limit** | **int32** | The maximum number of elements to return. Use together with &#39;offset&#39; for pagination. | [default to 100]|
| **offset** | **int32** | The first element to return. Use together with &#39;limit&#39; for pagination. | [default to 0]|

### Return type

[**ClusterBackupList**](../models/ClusterBackupList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ClustersBackupsFindById

```go
var result BackupResponse = ClustersBackupsFindById(ctx, backupId)
                      .Execute()
```

Fetch a cluster backup



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    psql "github.com/ionos-cloud/sdk-go-bundle/products/dbaas/psql"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    backupId := "498ae72f-411f-11eb-9d07-046c59cc737e-4oymiqu-12" // string | The unique ID of the backup.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := psql.NewAPIClient(configuration)
    resource, resp, err := apiClient.BackupsApi.ClustersBackupsFindById(context.Background(), backupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.ClustersBackupsFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersBackupsFindById`: BackupResponse
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.ClustersBackupsFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**backupId** | **string** | The unique ID of the backup. | |

### Other Parameters

Other parameters are passed through a pointer to an apiClustersBackupsFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**BackupResponse**](../models/BackupResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ClustersBackupsGet

```go
var result ClusterBackupList = ClustersBackupsGet(ctx)
                      .Limit(limit)
                      .Offset(offset)
                      .Execute()
```

List cluster backups



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    psql "github.com/ionos-cloud/sdk-go-bundle/products/dbaas/psql"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    limit := int32(100) // int32 | The maximum number of elements to return. Use together with 'offset' for pagination. (optional) (default to 100)
    offset := int32(200) // int32 | The first element to return. Use together with 'limit' for pagination. (optional) (default to 0)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := psql.NewAPIClient(configuration)
    resource, resp, err := apiClient.BackupsApi.ClustersBackupsGet(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.ClustersBackupsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersBackupsGet`: ClusterBackupList
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.ClustersBackupsGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiClustersBackupsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **limit** | **int32** | The maximum number of elements to return. Use together with &#39;offset&#39; for pagination. | [default to 100]|
| **offset** | **int32** | The first element to return. Use together with &#39;limit&#39; for pagination. | [default to 0]|

### Return type

[**ClusterBackupList**](../models/ClusterBackupList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


