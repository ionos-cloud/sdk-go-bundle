# \PipelinesApi

All URIs are relative to *https://logging.de-txl.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**PipelineKey**](PipelinesApi.md#PipelineKey) | **Post** /pipelines/{pipelineId}/key | Renews the key of a Pipeline|
|[**PipelinesDelete**](PipelinesApi.md#PipelinesDelete) | **Delete** /pipelines/{pipelineId} | Delete a pipeline|
|[**PipelinesFindById**](PipelinesApi.md#PipelinesFindById) | **Get** /pipelines/{pipelineId} | Fetch a pipeline|
|[**PipelinesGet**](PipelinesApi.md#PipelinesGet) | **Get** /pipelines | List pipelines|
|[**PipelinesPatch**](PipelinesApi.md#PipelinesPatch) | **Patch** /pipelines/{pipelineId} | Patch a pipeline|
|[**PipelinesPost**](PipelinesApi.md#PipelinesPost) | **Post** /pipelines | Create a pipeline|



## PipelineKey

```go
var result InlineResponse200 = PipelineKey(ctx, pipelineId)
                      .Execute()
```

Renews the key of a Pipeline



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    logging "github.com/ionos-cloud/sdk-go-bundle/products/logging"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    pipelineId := "pipelineId_example" // string | The unique ID of the pipeline

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := logging.NewAPIClient(configuration)
    resource, resp, err := apiClient.PipelinesApi.PipelineKey(context.Background(), pipelineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.PipelineKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `PipelineKey`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.PipelineKey`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**pipelineId** | **string** | The unique ID of the pipeline | |

### Other Parameters

Other parameters are passed through a pointer to an apiPipelineKeyRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**InlineResponse200**](../models/InlineResponse200.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## PipelinesDelete

```go
var result Pipeline = PipelinesDelete(ctx, pipelineId)
                      .Execute()
```

Delete a pipeline



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    logging "github.com/ionos-cloud/sdk-go-bundle/products/logging"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    pipelineId := "pipelineId_example" // string | The unique ID of the pipeline

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := logging.NewAPIClient(configuration)
    resp, err := apiClient.PipelinesApi.PipelinesDelete(context.Background(), pipelineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.PipelinesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `PipelinesDelete`: Pipeline
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.PipelinesDelete`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**pipelineId** | **string** | The unique ID of the pipeline | |

### Other Parameters

Other parameters are passed through a pointer to an apiPipelinesDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**Pipeline**](../models/Pipeline.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## PipelinesFindById

```go
var result Pipeline = PipelinesFindById(ctx, pipelineId)
                      .Execute()
```

Fetch a pipeline



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    logging "github.com/ionos-cloud/sdk-go-bundle/products/logging"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    pipelineId := "pipelineId_example" // string | The unique ID of the pipeline

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := logging.NewAPIClient(configuration)
    resource, resp, err := apiClient.PipelinesApi.PipelinesFindById(context.Background(), pipelineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.PipelinesFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `PipelinesFindById`: Pipeline
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.PipelinesFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**pipelineId** | **string** | The unique ID of the pipeline | |

### Other Parameters

Other parameters are passed through a pointer to an apiPipelinesFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**Pipeline**](../models/Pipeline.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## PipelinesGet

```go
var result PipelineListResponse = PipelinesGet(ctx)
                      .Limit(limit)
                      .Offset(offset)
                      .OrderBy(orderBy)
                      .Execute()
```

List pipelines



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    logging "github.com/ionos-cloud/sdk-go-bundle/products/logging"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    limit := int32(56) // int32 | the maximum number of elements to return (use together with offset for pagination). Default to 100 (optional) (default to 0)
    offset := int32(56) // int32 | the first element (of the total list of elements) to include in the response (use together with limit for pagination). Default to 0 (optional) (default to 0)
    orderBy := "orderBy_example" // string | Sorts the results alphanumerically in ascending order based on the specified property (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := logging.NewAPIClient(configuration)
    resource, resp, err := apiClient.PipelinesApi.PipelinesGet(context.Background()).Limit(limit).Offset(offset).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.PipelinesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `PipelinesGet`: PipelineListResponse
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.PipelinesGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiPipelinesGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **limit** | **int32** | the maximum number of elements to return (use together with offset for pagination). Default to 100 | [default to 0]|
| **offset** | **int32** | the first element (of the total list of elements) to include in the response (use together with limit for pagination). Default to 0 | [default to 0]|
| **orderBy** | **string** | Sorts the results alphanumerically in ascending order based on the specified property | |

### Return type

[**PipelineListResponse**](../models/PipelineListResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## PipelinesPatch

```go
var result Pipeline = PipelinesPatch(ctx, pipelineId)
                      .Pipeline(pipeline)
                      .Execute()
```

Patch a pipeline



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    logging "github.com/ionos-cloud/sdk-go-bundle/products/logging"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    pipelineId := "pipelineId_example" // string | The unique ID of the pipeline
    pipeline := *openapiclient.NewPatchRequest(*openapiclient.NewPatchRequestProperties()) // PatchRequest | The modified pipeline.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := logging.NewAPIClient(configuration)
    resource, resp, err := apiClient.PipelinesApi.PipelinesPatch(context.Background(), pipelineId).Pipeline(pipeline).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.PipelinesPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `PipelinesPatch`: Pipeline
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.PipelinesPatch`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**pipelineId** | **string** | The unique ID of the pipeline | |

### Other Parameters

Other parameters are passed through a pointer to an apiPipelinesPatchRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pipeline** | [**PatchRequest**](../models/PatchRequest.md) | The modified pipeline. | |

### Return type

[**Pipeline**](../models/Pipeline.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## PipelinesPost

```go
var result Pipeline = PipelinesPost(ctx)
                      .Pipeline(pipeline)
                      .Execute()
```

Create a pipeline



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    logging "github.com/ionos-cloud/sdk-go-bundle/products/logging"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    pipeline := *openapiclient.NewCreateRequest(*openapiclient.NewCreateRequestProperties("Name_example", []openapiclient.CreateRequestPipeline{*openapiclient.NewCreateRequestPipeline()})) // CreateRequest | The pipeline to be created.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := logging.NewAPIClient(configuration)
    resource, resp, err := apiClient.PipelinesApi.PipelinesPost(context.Background()).Pipeline(pipeline).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.PipelinesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `PipelinesPost`: Pipeline
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.PipelinesPost`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiPipelinesPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pipeline** | [**CreateRequest**](../models/CreateRequest.md) | The pipeline to be created. | |

### Return type

[**Pipeline**](../models/Pipeline.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


