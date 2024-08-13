# \APIGatewaysApi

All URIs are relative to *https://apigateway.de-txl.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**ApigatewaysDelete**](APIGatewaysApi.md#ApigatewaysDelete) | **Delete** /gateways/{apigatewayId} | Delete Gateway|
|[**ApigatewaysFindById**](APIGatewaysApi.md#ApigatewaysFindById) | **Get** /gateways/{apigatewayId} | Retrieve Gateway|
|[**ApigatewaysGet**](APIGatewaysApi.md#ApigatewaysGet) | **Get** /gateways | Retrieve all Apigateways|
|[**ApigatewaysPost**](APIGatewaysApi.md#ApigatewaysPost) | **Post** /gateways | Create Gateway|
|[**ApigatewaysPut**](APIGatewaysApi.md#ApigatewaysPut) | **Put** /gateways/{apigatewayId} | Ensure Gateway|



## ApigatewaysDelete

```go
var result  = ApigatewaysDelete(ctx, apigatewayId)
                      .Execute()
```

Delete Gateway



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    apigateway "github.com/ionos-cloud/sdk-go-bundle/products/apigateway"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    apigatewayId := "0620c174-dd3c-5eb4-87c8-e2b516553a00" // string | The ID (UUID) of the Gateway.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := apigateway.NewAPIClient(configuration)
    resp, err := apiClient.APIGatewaysApi.ApigatewaysDelete(context.Background(), apigatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIGatewaysApi.ApigatewaysDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**apigatewayId** | **string** | The ID (UUID) of the Gateway. | |

### Other Parameters

Other parameters are passed through a pointer to an apiApigatewaysDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ApigatewaysFindById

```go
var result GatewayRead = ApigatewaysFindById(ctx, apigatewayId)
                      .Execute()
```

Retrieve Gateway



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    apigateway "github.com/ionos-cloud/sdk-go-bundle/products/apigateway"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    apigatewayId := "0620c174-dd3c-5eb4-87c8-e2b516553a00" // string | The ID (UUID) of the Gateway.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := apigateway.NewAPIClient(configuration)
    resource, resp, err := apiClient.APIGatewaysApi.ApigatewaysFindById(context.Background(), apigatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIGatewaysApi.ApigatewaysFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ApigatewaysFindById`: GatewayRead
    fmt.Fprintf(os.Stdout, "Response from `APIGatewaysApi.ApigatewaysFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**apigatewayId** | **string** | The ID (UUID) of the Gateway. | |

### Other Parameters

Other parameters are passed through a pointer to an apiApigatewaysFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**GatewayRead**](../models/GatewayRead.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ApigatewaysGet

```go
var result GatewayReadList = ApigatewaysGet(ctx)
                      .Offset(offset)
                      .Limit(limit)
                      .OrderBy(orderBy)
                      .Execute()
```

Retrieve all Apigateways



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    apigateway "github.com/ionos-cloud/sdk-go-bundle/products/apigateway"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    offset := int32(0) // int32 | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. (optional) (default to 0)
    limit := int32(100) // int32 | The maximum number of elements to return. Use together with offset for pagination. (optional) (default to 100)
    orderBy := "orderBy_example" // string | The field to order the results by. If not provided, the results will be ordered by the default field. (optional) (default to "-createdDate")

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := apigateway.NewAPIClient(configuration)
    resource, resp, err := apiClient.APIGatewaysApi.ApigatewaysGet(context.Background()).Offset(offset).Limit(limit).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIGatewaysApi.ApigatewaysGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ApigatewaysGet`: GatewayReadList
    fmt.Fprintf(os.Stdout, "Response from `APIGatewaysApi.ApigatewaysGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiApigatewaysGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **offset** | **int32** | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return. Use together with offset for pagination. | [default to 100]|
| **orderBy** | **string** | The field to order the results by. If not provided, the results will be ordered by the default field. | [default to &quot;-createdDate&quot;]|

### Return type

[**GatewayReadList**](../models/GatewayReadList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ApigatewaysPost

```go
var result GatewayRead = ApigatewaysPost(ctx)
                      .GatewayCreate(gatewayCreate)
                      .Execute()
```

Create Gateway



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    apigateway "github.com/ionos-cloud/sdk-go-bundle/products/apigateway"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    gatewayCreate := *openapiclient.NewGatewayCreate(*openapiclient.NewGateway("APIGateway 1")) // GatewayCreate | Gateway to create.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := apigateway.NewAPIClient(configuration)
    resource, resp, err := apiClient.APIGatewaysApi.ApigatewaysPost(context.Background()).GatewayCreate(gatewayCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIGatewaysApi.ApigatewaysPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ApigatewaysPost`: GatewayRead
    fmt.Fprintf(os.Stdout, "Response from `APIGatewaysApi.ApigatewaysPost`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiApigatewaysPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **gatewayCreate** | [**GatewayCreate**](../models/GatewayCreate.md) | Gateway to create. | |

### Return type

[**GatewayRead**](../models/GatewayRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## ApigatewaysPut

```go
var result GatewayRead = ApigatewaysPut(ctx, apigatewayId)
                      .GatewayEnsure(gatewayEnsure)
                      .Execute()
```

Ensure Gateway



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    apigateway "github.com/ionos-cloud/sdk-go-bundle/products/apigateway"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    apigatewayId := "0620c174-dd3c-5eb4-87c8-e2b516553a00" // string | The ID (UUID) of the Gateway.
    gatewayEnsure := *openapiclient.NewGatewayEnsure("0620c174-dd3c-5eb4-87c8-e2b516553a00", *openapiclient.NewGateway("APIGateway 1")) // GatewayEnsure | update Gateway

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := apigateway.NewAPIClient(configuration)
    resource, resp, err := apiClient.APIGatewaysApi.ApigatewaysPut(context.Background(), apigatewayId).GatewayEnsure(gatewayEnsure).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIGatewaysApi.ApigatewaysPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ApigatewaysPut`: GatewayRead
    fmt.Fprintf(os.Stdout, "Response from `APIGatewaysApi.ApigatewaysPut`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**apigatewayId** | **string** | The ID (UUID) of the Gateway. | |

### Other Parameters

Other parameters are passed through a pointer to an apiApigatewaysPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **gatewayEnsure** | [**GatewayEnsure**](../models/GatewayEnsure.md) | update Gateway | |

### Return type

[**GatewayRead**](../models/GatewayRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


