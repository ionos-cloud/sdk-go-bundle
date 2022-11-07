# \SnapshotsApi

All URIs are relative to *https://api.ionos.com/databases/mongodb*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**ClustersSnapshotsGet**](SnapshotsApi.md#ClustersSnapshotsGet) | **Get** /clusters/{clusterId}/snapshots | Get the snapshots of your cluster|



## ClustersSnapshotsGet

```go
var result SnapshotList = ClustersSnapshotsGet(ctx, clusterId)
                      .Execute()
```

Get the snapshots of your cluster



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    mongo "github.com/ionos-cloud/sdk-go-bundle/services/mongo"
    "github.com/ionos-cloud/sdk-go-bundle/common"
)

func main() {
    clusterId := "clusterId_example" // string | The unique ID of the cluster.

    configuration := common.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := mongo.NewAPIClient(configuration)
    resource, resp, err := apiClient.SnapshotsApi.ClustersSnapshotsGet(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsApi.ClustersSnapshotsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersSnapshotsGet`: SnapshotList
    fmt.Fprintf(os.Stdout, "Response from `SnapshotsApi.ClustersSnapshotsGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The unique ID of the cluster. | |

### Other Parameters

Other parameters are passed through a pointer to an apiClustersSnapshotsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**SnapshotList**](SnapshotList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


