# \PipelinesApi

All URIs are relative to *https://logging.de-fra.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**PipelinesDelete**](PipelinesApi.md#PipelinesDelete) | **Delete** /pipelines/{pipelineId} | Delete Pipeline|
|[**PipelinesFindById**](PipelinesApi.md#PipelinesFindById) | **Get** /pipelines/{pipelineId} | Retrieve Pipeline|
|[**PipelinesGet**](PipelinesApi.md#PipelinesGet) | **Get** /pipelines | Retrieve all Pipelines|
|[**PipelinesPatch**](PipelinesApi.md#PipelinesPatch) | **Patch** /pipelines/{pipelineId} | Updates Pipeline|
|[**PipelinesPost**](PipelinesApi.md#PipelinesPost) | **Post** /pipelines | Create Pipeline|



## PipelinesDelete

```go
var result  = PipelinesDelete(ctx, pipelineId)
                      .Execute()
```

Delete Pipeline



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
    pipelineId := "f72521ba-1590-5998-bf96-6eb997a5887d" // string | The ID (UUID) of the Pipeline.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := logging.NewAPIClient(configuration)
    resp, err := apiClient.PipelinesApi.PipelinesDelete(context.Background(), pipelineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.PipelinesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**pipelineId** | **string** | The ID (UUID) of the Pipeline. | |

### Other Parameters

Other parameters are passed through a pointer to an apiPipelinesDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## PipelinesFindById

```go
var result PipelineRead = PipelinesFindById(ctx, pipelineId)
                      .Execute()
```

Retrieve Pipeline



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
    pipelineId := "f72521ba-1590-5998-bf96-6eb997a5887d" // string | The ID (UUID) of the Pipeline.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := logging.NewAPIClient(configuration)
    resource, resp, err := apiClient.PipelinesApi.PipelinesFindById(context.Background(), pipelineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.PipelinesFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `PipelinesFindById`: PipelineRead
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.PipelinesFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**pipelineId** | **string** | The ID (UUID) of the Pipeline. | |

### Other Parameters

Other parameters are passed through a pointer to an apiPipelinesFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**PipelineRead**](../models/PipelineRead.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## PipelinesGet

```go
var result PipelineReadList = PipelinesGet(ctx)
                      .Offset(offset)
                      .Limit(limit)
                      .OrderBy(orderBy)
                      .Execute()
```

Retrieve all Pipelines



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
    offset := int32(0) // int32 | The first element (of the total list of elements) to include in the response. Use this parameter together with the limit for pagination. (optional) (default to 0)
    limit := int32(100) // int32 | The maximum number of elements to return. Use this parameter together with the offset for pagination. (optional) (default to 100)
    orderBy := "orderBy_example" // string | The field to order the results by. If not provided, the results will be ordered by the default field. (optional) (default to "-createdDate")

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := logging.NewAPIClient(configuration)
    resource, resp, err := apiClient.PipelinesApi.PipelinesGet(context.Background()).Offset(offset).Limit(limit).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.PipelinesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `PipelinesGet`: PipelineReadList
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.PipelinesGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiPipelinesGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **offset** | **int32** | The first element (of the total list of elements) to include in the response. Use this parameter together with the limit for pagination. | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return. Use this parameter together with the offset for pagination. | [default to 100]|
| **orderBy** | **string** | The field to order the results by. If not provided, the results will be ordered by the default field. | [default to &quot;-createdDate&quot;]|

### Return type

[**PipelineReadList**](../models/PipelineReadList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## PipelinesPatch

```go
var result PipelineRead = PipelinesPatch(ctx, pipelineId)
                      .PipelinePatch(pipelinePatch)
                      .Execute()
```

Updates Pipeline



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
    pipelineId := "f72521ba-1590-5998-bf96-6eb997a5887d" // string | The ID (UUID) of the Pipeline.
    pipelinePatch := *openapiclient.NewPipelinePatch(*openapiclient.NewPipelineNoAddr("Pipeline1")) // PipelinePatch | patch Pipeline

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := logging.NewAPIClient(configuration)
    resource, resp, err := apiClient.PipelinesApi.PipelinesPatch(context.Background(), pipelineId).PipelinePatch(pipelinePatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.PipelinesPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `PipelinesPatch`: PipelineRead
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.PipelinesPatch`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**pipelineId** | **string** | The ID (UUID) of the Pipeline. | |

### Other Parameters

Other parameters are passed through a pointer to an apiPipelinesPatchRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pipelinePatch** | [**PipelinePatch**](../models/PipelinePatch.md) | patch Pipeline | |

### Return type

[**PipelineRead**](../models/PipelineRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## PipelinesPost

```go
var result PipelineRead = PipelinesPost(ctx)
                      .PipelineCreate(pipelineCreate)
                      .Execute()
```

Create Pipeline



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
    pipelineCreate := *openapiclient.NewPipelineCreate(*openapiclient.NewPipelineNoAddr("Pipeline1")) // PipelineCreate | Pipeline to create.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := logging.NewAPIClient(configuration)
    resource, resp, err := apiClient.PipelinesApi.PipelinesPost(context.Background()).PipelineCreate(pipelineCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.PipelinesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `PipelinesPost`: PipelineRead
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.PipelinesPost`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiPipelinesPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pipelineCreate** | [**PipelineCreate**](../models/PipelineCreate.md) | Pipeline to create. | |

### Return type

[**PipelineRead**](../models/PipelineRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


