# \RoutesApi

All URIs are relative to *https://apigateway.de-txl.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**ApigatewaysRoutesDelete**](RoutesApi.md#ApigatewaysRoutesDelete) | **Delete** /gateways/{apigatewayId}/routes/{routeId} | Delete Route|
|[**ApigatewaysRoutesFindById**](RoutesApi.md#ApigatewaysRoutesFindById) | **Get** /gateways/{apigatewayId}/routes/{routeId} | Retrieve Route|
|[**ApigatewaysRoutesGet**](RoutesApi.md#ApigatewaysRoutesGet) | **Get** /gateways/{apigatewayId}/routes | Retrieve all Routes|
|[**ApigatewaysRoutesPost**](RoutesApi.md#ApigatewaysRoutesPost) | **Post** /gateways/{apigatewayId}/routes | Create Route|
|[**ApigatewaysRoutesPut**](RoutesApi.md#ApigatewaysRoutesPut) | **Put** /gateways/{apigatewayId}/routes/{routeId} | Ensure Route|



## ApigatewaysRoutesDelete

```go
var result  = ApigatewaysRoutesDelete(ctx, apigatewayId, routeId)
                      .Execute()
```

Delete Route



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
    routeId := "50982018-bb17-5cb9-bcd4-97f8bbc7dc23" // string | The ID (UUID) of the Route.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := apigateway.NewAPIClient(configuration)
    resp, err := apiClient.RoutesApi.ApigatewaysRoutesDelete(context.Background(), apigatewayId, routeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutesApi.ApigatewaysRoutesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**apigatewayId** | **string** | The ID (UUID) of the Gateway. | |
|**routeId** | **string** | The ID (UUID) of the Route. | |

### Other Parameters

Other parameters are passed through a pointer to an apiApigatewaysRoutesDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ApigatewaysRoutesFindById

```go
var result RouteRead = ApigatewaysRoutesFindById(ctx, apigatewayId, routeId)
                      .Execute()
```

Retrieve Route



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
    routeId := "50982018-bb17-5cb9-bcd4-97f8bbc7dc23" // string | The ID (UUID) of the Route.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := apigateway.NewAPIClient(configuration)
    resource, resp, err := apiClient.RoutesApi.ApigatewaysRoutesFindById(context.Background(), apigatewayId, routeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutesApi.ApigatewaysRoutesFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ApigatewaysRoutesFindById`: RouteRead
    fmt.Fprintf(os.Stdout, "Response from `RoutesApi.ApigatewaysRoutesFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**apigatewayId** | **string** | The ID (UUID) of the Gateway. | |
|**routeId** | **string** | The ID (UUID) of the Route. | |

### Other Parameters

Other parameters are passed through a pointer to an apiApigatewaysRoutesFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**RouteRead**](../models/RouteRead.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ApigatewaysRoutesGet

```go
var result RouteReadList = ApigatewaysRoutesGet(ctx, apigatewayId)
                      .Offset(offset)
                      .Limit(limit)
                      .OrderBy(orderBy)
                      .Execute()
```

Retrieve all Routes



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
    offset := int32(0) // int32 | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. (optional) (default to 0)
    limit := int32(100) // int32 | The maximum number of elements to return. Use together with offset for pagination. (optional) (default to 100)
    orderBy := "orderBy_example" // string | The field to order the results by. If not provided, the results will be ordered by the default field. (optional) (default to "-createdDate")

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := apigateway.NewAPIClient(configuration)
    resource, resp, err := apiClient.RoutesApi.ApigatewaysRoutesGet(context.Background(), apigatewayId).Offset(offset).Limit(limit).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutesApi.ApigatewaysRoutesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ApigatewaysRoutesGet`: RouteReadList
    fmt.Fprintf(os.Stdout, "Response from `RoutesApi.ApigatewaysRoutesGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**apigatewayId** | **string** | The ID (UUID) of the Gateway. | |

### Other Parameters

Other parameters are passed through a pointer to an apiApigatewaysRoutesGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **offset** | **int32** | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return. Use together with offset for pagination. | [default to 100]|
| **orderBy** | **string** | The field to order the results by. If not provided, the results will be ordered by the default field. | [default to &quot;-createdDate&quot;]|

### Return type

[**RouteReadList**](../models/RouteReadList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ApigatewaysRoutesPost

```go
var result RouteRead = ApigatewaysRoutesPost(ctx, apigatewayId)
                      .RouteCreate(routeCreate)
                      .Execute()
```

Create Route



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
    routeCreate := *openapiclient.NewRouteCreate(*openapiclient.NewRoute("Name_example", "Type_example", []string{"Paths_example"}, []string{"Methods_example"}, []openapiclient.RouteUpstreams{*openapiclient.NewRouteUpstreams("Scheme_example", "Loadbalancer_example", "Host_example", int32(123))})) // RouteCreate | Route to create.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := apigateway.NewAPIClient(configuration)
    resource, resp, err := apiClient.RoutesApi.ApigatewaysRoutesPost(context.Background(), apigatewayId).RouteCreate(routeCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutesApi.ApigatewaysRoutesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ApigatewaysRoutesPost`: RouteRead
    fmt.Fprintf(os.Stdout, "Response from `RoutesApi.ApigatewaysRoutesPost`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**apigatewayId** | **string** | The ID (UUID) of the Gateway. | |

### Other Parameters

Other parameters are passed through a pointer to an apiApigatewaysRoutesPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **routeCreate** | [**RouteCreate**](../models/RouteCreate.md) | Route to create. | |

### Return type

[**RouteRead**](../models/RouteRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## ApigatewaysRoutesPut

```go
var result RouteRead = ApigatewaysRoutesPut(ctx, apigatewayId, routeId)
                      .RouteEnsure(routeEnsure)
                      .Execute()
```

Ensure Route



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
    routeId := "50982018-bb17-5cb9-bcd4-97f8bbc7dc23" // string | The ID (UUID) of the Route.
    routeEnsure := *openapiclient.NewRouteEnsure("50982018-bb17-5cb9-bcd4-97f8bbc7dc23", *openapiclient.NewRoute("Name_example", "Type_example", []string{"Paths_example"}, []string{"Methods_example"}, []openapiclient.RouteUpstreams{*openapiclient.NewRouteUpstreams("Scheme_example", "Loadbalancer_example", "Host_example", int32(123))})) // RouteEnsure | update Route

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := apigateway.NewAPIClient(configuration)
    resource, resp, err := apiClient.RoutesApi.ApigatewaysRoutesPut(context.Background(), apigatewayId, routeId).RouteEnsure(routeEnsure).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutesApi.ApigatewaysRoutesPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ApigatewaysRoutesPut`: RouteRead
    fmt.Fprintf(os.Stdout, "Response from `RoutesApi.ApigatewaysRoutesPut`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**apigatewayId** | **string** | The ID (UUID) of the Gateway. | |
|**routeId** | **string** | The ID (UUID) of the Route. | |

### Other Parameters

Other parameters are passed through a pointer to an apiApigatewaysRoutesPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **routeEnsure** | [**RouteEnsure**](../models/RouteEnsure.md) | update Route | |

### Return type

[**RouteRead**](../models/RouteRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


