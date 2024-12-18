# \KeyApi

All URIs are relative to *https://monitoring.de-fra.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**PipelinesKeyPost**](KeyApi.md#PipelinesKeyPost) | **Post** /pipelines/{pipelineId}/key | Create Key|



## PipelinesKeyPost

```go
var result KeyRead = PipelinesKeyPost(ctx, pipelineId)
                      .Body(body)
                      .Execute()
```

Create Key



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    monitoring "github.com/ionos-cloud/sdk-go-bundle/products/monitoring"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    pipelineId := "f72521ba-1590-5998-bf96-6eb997a5887d" // string | The ID (UUID) of the Pipeline.
    body := map[string]interface{}{ ... } // map[string]interface{} | Key to create.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := monitoring.NewAPIClient(configuration)
    resource, resp, err := apiClient.KeyApi.PipelinesKeyPost(context.Background(), pipelineId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyApi.PipelinesKeyPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `PipelinesKeyPost`: KeyRead
    fmt.Fprintf(os.Stdout, "Response from `KeyApi.PipelinesKeyPost`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**pipelineId** | **string** | The ID (UUID) of the Pipeline. | |

### Other Parameters

Other parameters are passed through a pointer to an apiPipelinesKeyPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **body** | **map[string]interface{}** | Key to create. | |

### Return type

[**KeyRead**](../models/KeyRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


