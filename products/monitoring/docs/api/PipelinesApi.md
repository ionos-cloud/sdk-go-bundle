# \PipelinesApi

All URIs are relative to *https://monitoring.de-fra.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**PipelinesDelete**](PipelinesApi.md#PipelinesDelete) | **Delete** /pipelines/{pipelineId} | Delete Pipeline|
|[**PipelinesFindById**](PipelinesApi.md#PipelinesFindById) | **Get** /pipelines/{pipelineId} | Retrieve Pipeline|
|[**PipelinesGet**](PipelinesApi.md#PipelinesGet) | **Get** /pipelines | Retrieve all Pipelines|
|[**PipelinesPost**](PipelinesApi.md#PipelinesPost) | **Post** /pipelines | Create Pipeline|
|[**PipelinesPut**](PipelinesApi.md#PipelinesPut) | **Put** /pipelines/{pipelineId} | Ensure Pipeline|



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

    monitoring "github.com/ionos-cloud/sdk-go-bundle/products/monitoring"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    pipelineId := "f72521ba-1590-5998-bf96-6eb997a5887d" // string | The ID (UUID) of the Pipeline.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := monitoring.NewAPIClient(configuration)
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

    monitoring "github.com/ionos-cloud/sdk-go-bundle/products/monitoring"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    pipelineId := "f72521ba-1590-5998-bf96-6eb997a5887d" // string | The ID (UUID) of the Pipeline.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := monitoring.NewAPIClient(configuration)
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

    monitoring "github.com/ionos-cloud/sdk-go-bundle/products/monitoring"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    orderBy := "orderBy_example" // string | The field to order the results by. If not provided, the results will be ordered by the default field. (optional) (default to "-createdDate")

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := monitoring.NewAPIClient(configuration)
    resource, resp, err := apiClient.PipelinesApi.PipelinesGet(context.Background()).OrderBy(orderBy).Execute()
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
| **orderBy** | **string** | The field to order the results by. If not provided, the results will be ordered by the default field. | [default to &quot;-createdDate&quot;]|

### Return type

[**PipelineReadList**](../models/PipelineReadList.md)

### HTTP request headers

- **Content-Type**: Not defined
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

    monitoring "github.com/ionos-cloud/sdk-go-bundle/products/monitoring"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    pipelineCreate := *openapiclient.NewPipelineCreate(*openapiclient.NewPipeline("Pipeline1")) // PipelineCreate | Pipeline to create.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := monitoring.NewAPIClient(configuration)
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



## PipelinesPut

```go
var result PipelineRead = PipelinesPut(ctx, pipelineId)
                      .PipelineEnsure(pipelineEnsure)
                      .Execute()
```

Ensure Pipeline



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
    pipelineEnsure := *openapiclient.NewPipelineEnsure(*openapiclient.NewPipeline("Pipeline1")) // PipelineEnsure | update Pipeline

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := monitoring.NewAPIClient(configuration)
    resource, resp, err := apiClient.PipelinesApi.PipelinesPut(context.Background(), pipelineId).PipelineEnsure(pipelineEnsure).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PipelinesApi.PipelinesPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `PipelinesPut`: PipelineRead
    fmt.Fprintf(os.Stdout, "Response from `PipelinesApi.PipelinesPut`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**pipelineId** | **string** | The ID (UUID) of the Pipeline. | |

### Other Parameters

Other parameters are passed through a pointer to an apiPipelinesPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pipelineEnsure** | [**PipelineEnsure**](../models/PipelineEnsure.md) | update Pipeline | |

### Return type

[**PipelineRead**](../models/PipelineRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


