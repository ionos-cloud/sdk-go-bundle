# \IPSecGatewaysApi

All URIs are relative to *https://vpn.de-fra.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**IpsecgatewaysDelete**](IPSecGatewaysApi.md#IpsecgatewaysDelete) | **Delete** /ipsecgateways/{gatewayId} | Delete IPSecGateway|
|[**IpsecgatewaysFindById**](IPSecGatewaysApi.md#IpsecgatewaysFindById) | **Get** /ipsecgateways/{gatewayId} | Retrieve IPSecGateway|
|[**IpsecgatewaysGet**](IPSecGatewaysApi.md#IpsecgatewaysGet) | **Get** /ipsecgateways | Retrieve all IPSecGateways|
|[**IpsecgatewaysPost**](IPSecGatewaysApi.md#IpsecgatewaysPost) | **Post** /ipsecgateways | Create IPSecGateway|
|[**IpsecgatewaysPut**](IPSecGatewaysApi.md#IpsecgatewaysPut) | **Put** /ipsecgateways/{gatewayId} | Ensure IPSecGateway|



## IpsecgatewaysDelete

```go
var result  = IpsecgatewaysDelete(ctx, gatewayId)
                      .Execute()
```

Delete IPSecGateway



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
    gatewayId := "66a114c7-2ddd-5119-9ddf-5a789f5a5a44" // string | The ID (UUID) of the IPSecGateway.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := vpn.NewAPIClient(configuration)
    resp, err := apiClient.IPSecGatewaysApi.IpsecgatewaysDelete(context.Background(), gatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPSecGatewaysApi.IpsecgatewaysDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**gatewayId** | **string** | The ID (UUID) of the IPSecGateway. | |

### Other Parameters

Other parameters are passed through a pointer to an apiIpsecgatewaysDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## IpsecgatewaysFindById

```go
var result IPSecGatewayRead = IpsecgatewaysFindById(ctx, gatewayId)
                      .Execute()
```

Retrieve IPSecGateway



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
    gatewayId := "66a114c7-2ddd-5119-9ddf-5a789f5a5a44" // string | The ID (UUID) of the IPSecGateway.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := vpn.NewAPIClient(configuration)
    resource, resp, err := apiClient.IPSecGatewaysApi.IpsecgatewaysFindById(context.Background(), gatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPSecGatewaysApi.IpsecgatewaysFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `IpsecgatewaysFindById`: IPSecGatewayRead
    fmt.Fprintf(os.Stdout, "Response from `IPSecGatewaysApi.IpsecgatewaysFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**gatewayId** | **string** | The ID (UUID) of the IPSecGateway. | |

### Other Parameters

Other parameters are passed through a pointer to an apiIpsecgatewaysFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**IPSecGatewayRead**](../models/IPSecGatewayRead.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## IpsecgatewaysGet

```go
var result IPSecGatewayReadList = IpsecgatewaysGet(ctx)
                      .Offset(offset)
                      .Limit(limit)
                      .Execute()
```

Retrieve all IPSecGateways



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
    resource, resp, err := apiClient.IPSecGatewaysApi.IpsecgatewaysGet(context.Background()).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPSecGatewaysApi.IpsecgatewaysGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `IpsecgatewaysGet`: IPSecGatewayReadList
    fmt.Fprintf(os.Stdout, "Response from `IPSecGatewaysApi.IpsecgatewaysGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiIpsecgatewaysGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **offset** | **int32** | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return. Use together with offset for pagination. | [default to 100]|

### Return type

[**IPSecGatewayReadList**](../models/IPSecGatewayReadList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## IpsecgatewaysPost

```go
var result IPSecGatewayRead = IpsecgatewaysPost(ctx)
                      .IPSecGatewayCreate(iPSecGatewayCreate)
                      .Execute()
```

Create IPSecGateway



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
    iPSecGatewayCreate := *openapiclient.NewIPSecGatewayCreate(*openapiclient.NewIPSecGateway("My Company Gateway", "81.173.1.2", []openapiclient.Connection{*openapiclient.NewConnection("5a029f4a-72e5-11ec-90d6-0242ac120003", "2", "192.168.1.100/24")})) // IPSecGatewayCreate | IPSecGateway to create.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := vpn.NewAPIClient(configuration)
    resource, resp, err := apiClient.IPSecGatewaysApi.IpsecgatewaysPost(context.Background()).IPSecGatewayCreate(iPSecGatewayCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPSecGatewaysApi.IpsecgatewaysPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `IpsecgatewaysPost`: IPSecGatewayRead
    fmt.Fprintf(os.Stdout, "Response from `IPSecGatewaysApi.IpsecgatewaysPost`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiIpsecgatewaysPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **iPSecGatewayCreate** | [**IPSecGatewayCreate**](../models/IPSecGatewayCreate.md) | IPSecGateway to create. | |

### Return type

[**IPSecGatewayRead**](../models/IPSecGatewayRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## IpsecgatewaysPut

```go
var result IPSecGatewayRead = IpsecgatewaysPut(ctx, gatewayId)
                      .IPSecGatewayEnsure(iPSecGatewayEnsure)
                      .Execute()
```

Ensure IPSecGateway



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
    gatewayId := "66a114c7-2ddd-5119-9ddf-5a789f5a5a44" // string | The ID (UUID) of the IPSecGateway.
    iPSecGatewayEnsure := *openapiclient.NewIPSecGatewayEnsure("66a114c7-2ddd-5119-9ddf-5a789f5a5a44", *openapiclient.NewIPSecGateway("My Company Gateway", "81.173.1.2", []openapiclient.Connection{*openapiclient.NewConnection("5a029f4a-72e5-11ec-90d6-0242ac120003", "2", "192.168.1.100/24")})) // IPSecGatewayEnsure | update IPSecGateway

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := vpn.NewAPIClient(configuration)
    resource, resp, err := apiClient.IPSecGatewaysApi.IpsecgatewaysPut(context.Background(), gatewayId).IPSecGatewayEnsure(iPSecGatewayEnsure).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPSecGatewaysApi.IpsecgatewaysPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `IpsecgatewaysPut`: IPSecGatewayRead
    fmt.Fprintf(os.Stdout, "Response from `IPSecGatewaysApi.IpsecgatewaysPut`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**gatewayId** | **string** | The ID (UUID) of the IPSecGateway. | |

### Other Parameters

Other parameters are passed through a pointer to an apiIpsecgatewaysPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **iPSecGatewayEnsure** | [**IPSecGatewayEnsure**](../models/IPSecGatewayEnsure.md) | update IPSecGateway | |

### Return type

[**IPSecGatewayRead**](../models/IPSecGatewayRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


