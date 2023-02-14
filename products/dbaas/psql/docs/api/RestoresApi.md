# \RestoresApi

All URIs are relative to *https://api.ionos.com/databases/postgresql*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**ClusterRestorePost**](RestoresApi.md#ClusterRestorePost) | **Post** /clusters/{clusterId}/restore | In-place restore of a cluster|



## ClusterRestorePost

```go
var result  = ClusterRestorePost(ctx, clusterId)
                      .CreateRestoreRequest(createRestoreRequest)
                      .Execute()
```

In-place restore of a cluster



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
    clusterId := "clusterId_example" // string | The unique ID of the cluster.
    createRestoreRequest := *openapiclient.NewCreateRestoreRequest("dcd31531-3ac8-11eb-9feb-046c59cc737e") // CreateRestoreRequest | The restore request to create.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := psql.NewAPIClient(configuration)
    resource, resp, err := apiClient.RestoresApi.ClusterRestorePost(context.Background(), clusterId).CreateRestoreRequest(createRestoreRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoresApi.ClusterRestorePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The unique ID of the cluster. | |

### Other Parameters

Other parameters are passed through a pointer to an apiClusterRestorePostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **createRestoreRequest** | [**CreateRestoreRequest**](../models/CreateRestoreRequest.md) | The restore request to create. | |

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


