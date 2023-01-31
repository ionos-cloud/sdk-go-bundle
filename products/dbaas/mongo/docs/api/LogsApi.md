# \LogsApi

All URIs are relative to *https://api.ionos.com/databases/mongodb*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**ClustersLogsGet**](LogsApi.md#ClustersLogsGet) | **Get** /clusters/{clusterId}/logs | Get logs of your cluster|



## ClustersLogsGet

```go
var result ClusterLogs = ClustersLogsGet(ctx, clusterId)
                      .Start(start)
                      .End(end)
                      .Direction(direction)
                      .Limit(limit)
                      .Execute()
```

Get logs of your cluster



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    mongo "github.com/ionos-cloud/sdk-go-bundle/products/dbaas/mongo"
    "github.com/ionos-cloud/sdk-go-bundle/common"
)

func main() {
    clusterId := "clusterId_example" // string | The unique ID of the cluster.
    start := time.Now() // time.Time | The start time for the query in RFC3339 format. Must not be more than 30 days ago but before the end parameter. The default is 30 days ago. (optional)
    end := time.Now() // time.Time | The end time for the query in RFC3339 format. Must not be greater than now. The default is the current timestamp. (optional)
    direction := "direction_example" // string | The direction in which to scan through the logs. The logs are returned in order of the direction. (optional) (default to "BACKWARD")
    limit := int32(56) // int32 | The maximal number of log lines to return.  If the limit is reached then log lines will be cut at the end (respecting the scan direction). (optional) (default to 100)

    configuration := common.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := mongo.NewAPIClient(configuration)
    resource, resp, err := apiClient.LogsApi.ClustersLogsGet(context.Background(), clusterId).Start(start).End(end).Direction(direction).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.ClustersLogsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersLogsGet`: ClusterLogs
    fmt.Fprintf(os.Stdout, "Response from `LogsApi.ClustersLogsGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The unique ID of the cluster. | |

### Other Parameters

Other parameters are passed through a pointer to an apiClustersLogsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **start** | **time.Time** | The start time for the query in RFC3339 format. Must not be more than 30 days ago but before the end parameter. The default is 30 days ago. | |
| **end** | **time.Time** | The end time for the query in RFC3339 format. Must not be greater than now. The default is the current timestamp. | |
| **direction** | **string** | The direction in which to scan through the logs. The logs are returned in order of the direction. | [default to &quot;BACKWARD&quot;]|
| **limit** | **int32** | The maximal number of log lines to return.  If the limit is reached then log lines will be cut at the end (respecting the scan direction). | [default to 100]|

### Return type

[**ClusterLogs**](ClusterLogs.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


