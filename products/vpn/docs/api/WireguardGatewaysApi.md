# \WireguardGatewaysApi

All URIs are relative to *https://vpn.de-fra.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**WireguardgatewaysDelete**](WireguardGatewaysApi.md#WireguardgatewaysDelete) | **Delete** /wireguardgateways/{gatewayId} | Delete WireguardGateway|
|[**WireguardgatewaysFindById**](WireguardGatewaysApi.md#WireguardgatewaysFindById) | **Get** /wireguardgateways/{gatewayId} | Retrieve WireguardGateway|
|[**WireguardgatewaysGet**](WireguardGatewaysApi.md#WireguardgatewaysGet) | **Get** /wireguardgateways | Retrieve all WireguardGateways|
|[**WireguardgatewaysPost**](WireguardGatewaysApi.md#WireguardgatewaysPost) | **Post** /wireguardgateways | Create WireguardGateway|
|[**WireguardgatewaysPut**](WireguardGatewaysApi.md#WireguardgatewaysPut) | **Put** /wireguardgateways/{gatewayId} | Ensure WireguardGateway|



## WireguardgatewaysDelete

```go
var result  = WireguardgatewaysDelete(ctx, gatewayId)
                      .Execute()
```

Delete WireguardGateway



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    vpn "github.com/ionos-cloud/sdk-go-bundle/products/vpn"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    gatewayId := "85c79b4b-5b40-570a-b788-58dd46ea71e2" // string | The ID (UUID) of the WireguardGateway.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := vpn.NewAPIClient(configuration)
    resp, err := apiClient.WireguardGatewaysApi.WireguardgatewaysDelete(context.Background(), gatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WireguardGatewaysApi.WireguardgatewaysDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**gatewayId** | **string** | The ID (UUID) of the WireguardGateway. | |

### Other Parameters

Other parameters are passed through a pointer to an apiWireguardgatewaysDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## WireguardgatewaysFindById

```go
var result WireguardGatewayRead = WireguardgatewaysFindById(ctx, gatewayId)
                      .Execute()
```

Retrieve WireguardGateway



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    vpn "github.com/ionos-cloud/sdk-go-bundle/products/vpn"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    gatewayId := "85c79b4b-5b40-570a-b788-58dd46ea71e2" // string | The ID (UUID) of the WireguardGateway.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := vpn.NewAPIClient(configuration)
    resource, resp, err := apiClient.WireguardGatewaysApi.WireguardgatewaysFindById(context.Background(), gatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WireguardGatewaysApi.WireguardgatewaysFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `WireguardgatewaysFindById`: WireguardGatewayRead
    fmt.Fprintf(os.Stdout, "Response from `WireguardGatewaysApi.WireguardgatewaysFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**gatewayId** | **string** | The ID (UUID) of the WireguardGateway. | |

### Other Parameters

Other parameters are passed through a pointer to an apiWireguardgatewaysFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**WireguardGatewayRead**](../models/WireguardGatewayRead.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## WireguardgatewaysGet

```go
var result WireguardGatewayReadList = WireguardgatewaysGet(ctx)
                      .Offset(offset)
                      .Limit(limit)
                      .Execute()
```

Retrieve all WireguardGateways



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    vpn "github.com/ionos-cloud/sdk-go-bundle/products/vpn"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    offset := int32(0) // int32 | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. (optional) (default to 0)
    limit := int32(100) // int32 | The maximum number of elements to return. Use together with offset for pagination. (optional) (default to 100)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := vpn.NewAPIClient(configuration)
    resource, resp, err := apiClient.WireguardGatewaysApi.WireguardgatewaysGet(context.Background()).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WireguardGatewaysApi.WireguardgatewaysGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `WireguardgatewaysGet`: WireguardGatewayReadList
    fmt.Fprintf(os.Stdout, "Response from `WireguardGatewaysApi.WireguardgatewaysGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiWireguardgatewaysGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **offset** | **int32** | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return. Use together with offset for pagination. | [default to 100]|

### Return type

[**WireguardGatewayReadList**](../models/WireguardGatewayReadList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## WireguardgatewaysPost

```go
var result WireguardGatewayRead = WireguardgatewaysPost(ctx)
                      .WireguardGatewayCreate(wireguardGatewayCreate)
                      .Execute()
```

Create WireguardGateway



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    vpn "github.com/ionos-cloud/sdk-go-bundle/products/vpn"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    wireguardGatewayCreate := *openapiclient.NewWireguardGatewayCreate(*openapiclient.NewWireguardGateway("My Company Gateway", "81.173.1.2", []openapiclient.Connection{*openapiclient.NewConnection("5a029f4a-72e5-11ec-90d6-0242ac120003", "2", "192.168.1.100/24")}, "0HpE4BNwGHabeaC4aY/GFxB6fBSc0d49Db0qAzRVSVc=")) // WireguardGatewayCreate | WireguardGateway to create.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := vpn.NewAPIClient(configuration)
    resource, resp, err := apiClient.WireguardGatewaysApi.WireguardgatewaysPost(context.Background()).WireguardGatewayCreate(wireguardGatewayCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WireguardGatewaysApi.WireguardgatewaysPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `WireguardgatewaysPost`: WireguardGatewayRead
    fmt.Fprintf(os.Stdout, "Response from `WireguardGatewaysApi.WireguardgatewaysPost`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiWireguardgatewaysPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **wireguardGatewayCreate** | [**WireguardGatewayCreate**](../models/WireguardGatewayCreate.md) | WireguardGateway to create. | |

### Return type

[**WireguardGatewayRead**](../models/WireguardGatewayRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## WireguardgatewaysPut

```go
var result WireguardGatewayRead = WireguardgatewaysPut(ctx, gatewayId)
                      .WireguardGatewayEnsure(wireguardGatewayEnsure)
                      .Execute()
```

Ensure WireguardGateway



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    vpn "github.com/ionos-cloud/sdk-go-bundle/products/vpn"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    gatewayId := "85c79b4b-5b40-570a-b788-58dd46ea71e2" // string | The ID (UUID) of the WireguardGateway.
    wireguardGatewayEnsure := *openapiclient.NewWireguardGatewayEnsure("85c79b4b-5b40-570a-b788-58dd46ea71e2", *openapiclient.NewWireguardGateway("My Company Gateway", "81.173.1.2", []openapiclient.Connection{*openapiclient.NewConnection("5a029f4a-72e5-11ec-90d6-0242ac120003", "2", "192.168.1.100/24")}, "0HpE4BNwGHabeaC4aY/GFxB6fBSc0d49Db0qAzRVSVc=")) // WireguardGatewayEnsure | update WireguardGateway

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := vpn.NewAPIClient(configuration)
    resource, resp, err := apiClient.WireguardGatewaysApi.WireguardgatewaysPut(context.Background(), gatewayId).WireguardGatewayEnsure(wireguardGatewayEnsure).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WireguardGatewaysApi.WireguardgatewaysPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `WireguardgatewaysPut`: WireguardGatewayRead
    fmt.Fprintf(os.Stdout, "Response from `WireguardGatewaysApi.WireguardgatewaysPut`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**gatewayId** | **string** | The ID (UUID) of the WireguardGateway. | |

### Other Parameters

Other parameters are passed through a pointer to an apiWireguardgatewaysPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **wireguardGatewayEnsure** | [**WireguardGatewayEnsure**](../models/WireguardGatewayEnsure.md) | update WireguardGateway | |

### Return type

[**WireguardGatewayRead**](../models/WireguardGatewayRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


