# \IPSecTunnelsApi

All URIs are relative to *https://vpn.de-fra.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**IpsecgatewaysTunnelsDelete**](IPSecTunnelsApi.md#IpsecgatewaysTunnelsDelete) | **Delete** /ipsecgateways/{gatewayId}/tunnels/{tunnelId} | Delete IPSecTunnel|
|[**IpsecgatewaysTunnelsFindById**](IPSecTunnelsApi.md#IpsecgatewaysTunnelsFindById) | **Get** /ipsecgateways/{gatewayId}/tunnels/{tunnelId} | Retrieve IPSecTunnel|
|[**IpsecgatewaysTunnelsGet**](IPSecTunnelsApi.md#IpsecgatewaysTunnelsGet) | **Get** /ipsecgateways/{gatewayId}/tunnels | Retrieve all IPSecTunnels|
|[**IpsecgatewaysTunnelsPost**](IPSecTunnelsApi.md#IpsecgatewaysTunnelsPost) | **Post** /ipsecgateways/{gatewayId}/tunnels | Create IPSecTunnel|
|[**IpsecgatewaysTunnelsPut**](IPSecTunnelsApi.md#IpsecgatewaysTunnelsPut) | **Put** /ipsecgateways/{gatewayId}/tunnels/{tunnelId} | Ensure IPSecTunnel|



## IpsecgatewaysTunnelsDelete

```go
var result  = IpsecgatewaysTunnelsDelete(ctx, gatewayId, tunnelId)
                      .Execute()
```

Delete IPSecTunnel



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
    tunnelId := "c28b2d3e-7b15-53ca-ae88-6ae9378d6efe" // string | The ID (UUID) of the IPSecTunnel.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := vpn.NewAPIClient(configuration)
    resp, err := apiClient.IPSecTunnelsApi.IpsecgatewaysTunnelsDelete(context.Background(), gatewayId, tunnelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPSecTunnelsApi.IpsecgatewaysTunnelsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**gatewayId** | **string** | The ID (UUID) of the IPSecGateway. | |
|**tunnelId** | **string** | The ID (UUID) of the IPSecTunnel. | |

### Other Parameters

Other parameters are passed through a pointer to an apiIpsecgatewaysTunnelsDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## IpsecgatewaysTunnelsFindById

```go
var result IPSecTunnelRead = IpsecgatewaysTunnelsFindById(ctx, gatewayId, tunnelId)
                      .Execute()
```

Retrieve IPSecTunnel



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
    tunnelId := "c28b2d3e-7b15-53ca-ae88-6ae9378d6efe" // string | The ID (UUID) of the IPSecTunnel.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := vpn.NewAPIClient(configuration)
    resource, resp, err := apiClient.IPSecTunnelsApi.IpsecgatewaysTunnelsFindById(context.Background(), gatewayId, tunnelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPSecTunnelsApi.IpsecgatewaysTunnelsFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `IpsecgatewaysTunnelsFindById`: IPSecTunnelRead
    fmt.Fprintf(os.Stdout, "Response from `IPSecTunnelsApi.IpsecgatewaysTunnelsFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**gatewayId** | **string** | The ID (UUID) of the IPSecGateway. | |
|**tunnelId** | **string** | The ID (UUID) of the IPSecTunnel. | |

### Other Parameters

Other parameters are passed through a pointer to an apiIpsecgatewaysTunnelsFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**IPSecTunnelRead**](../models/IPSecTunnelRead.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## IpsecgatewaysTunnelsGet

```go
var result IPSecTunnelReadList = IpsecgatewaysTunnelsGet(ctx, gatewayId)
                      .Offset(offset)
                      .Limit(limit)
                      .Execute()
```

Retrieve all IPSecTunnels



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
    offset := int32(0) // int32 | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. (optional) (default to 0)
    limit := int32(100) // int32 | The maximum number of elements to return. Use together with offset for pagination. (optional) (default to 100)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := vpn.NewAPIClient(configuration)
    resource, resp, err := apiClient.IPSecTunnelsApi.IpsecgatewaysTunnelsGet(context.Background(), gatewayId).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPSecTunnelsApi.IpsecgatewaysTunnelsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `IpsecgatewaysTunnelsGet`: IPSecTunnelReadList
    fmt.Fprintf(os.Stdout, "Response from `IPSecTunnelsApi.IpsecgatewaysTunnelsGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**gatewayId** | **string** | The ID (UUID) of the IPSecGateway. | |

### Other Parameters

Other parameters are passed through a pointer to an apiIpsecgatewaysTunnelsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **offset** | **int32** | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return. Use together with offset for pagination. | [default to 100]|

### Return type

[**IPSecTunnelReadList**](../models/IPSecTunnelReadList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## IpsecgatewaysTunnelsPost

```go
var result IPSecTunnelRead = IpsecgatewaysTunnelsPost(ctx, gatewayId)
                      .IPSecTunnelCreate(iPSecTunnelCreate)
                      .Execute()
```

Create IPSecTunnel



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
    iPSecTunnelCreate := *openapiclient.NewIPSecTunnelCreate(*openapiclient.NewIPSecTunnel("My Company Gateway Tunnel", "vpn.mycompany.com", *openapiclient.NewIPSecTunnelAuth("Method_example"), *openapiclient.NewIKEEncryption(), *openapiclient.NewESPEncryption(), []string{"CloudNetworkCIDRs_example"}, []string{"PeerNetworkCIDRs_example"})) // IPSecTunnelCreate | IPSecTunnel to create.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := vpn.NewAPIClient(configuration)
    resource, resp, err := apiClient.IPSecTunnelsApi.IpsecgatewaysTunnelsPost(context.Background(), gatewayId).IPSecTunnelCreate(iPSecTunnelCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPSecTunnelsApi.IpsecgatewaysTunnelsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `IpsecgatewaysTunnelsPost`: IPSecTunnelRead
    fmt.Fprintf(os.Stdout, "Response from `IPSecTunnelsApi.IpsecgatewaysTunnelsPost`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**gatewayId** | **string** | The ID (UUID) of the IPSecGateway. | |

### Other Parameters

Other parameters are passed through a pointer to an apiIpsecgatewaysTunnelsPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **iPSecTunnelCreate** | [**IPSecTunnelCreate**](../models/IPSecTunnelCreate.md) | IPSecTunnel to create. | |

### Return type

[**IPSecTunnelRead**](../models/IPSecTunnelRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## IpsecgatewaysTunnelsPut

```go
var result IPSecTunnelRead = IpsecgatewaysTunnelsPut(ctx, gatewayId, tunnelId)
                      .IPSecTunnelEnsure(iPSecTunnelEnsure)
                      .Execute()
```

Ensure IPSecTunnel



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
    tunnelId := "c28b2d3e-7b15-53ca-ae88-6ae9378d6efe" // string | The ID (UUID) of the IPSecTunnel.
    iPSecTunnelEnsure := *openapiclient.NewIPSecTunnelEnsure("c28b2d3e-7b15-53ca-ae88-6ae9378d6efe", *openapiclient.NewIPSecTunnel("My Company Gateway Tunnel", "vpn.mycompany.com", *openapiclient.NewIPSecTunnelAuth("Method_example"), *openapiclient.NewIKEEncryption(), *openapiclient.NewESPEncryption(), []string{"CloudNetworkCIDRs_example"}, []string{"PeerNetworkCIDRs_example"})) // IPSecTunnelEnsure | update IPSecTunnel

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := vpn.NewAPIClient(configuration)
    resource, resp, err := apiClient.IPSecTunnelsApi.IpsecgatewaysTunnelsPut(context.Background(), gatewayId, tunnelId).IPSecTunnelEnsure(iPSecTunnelEnsure).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPSecTunnelsApi.IpsecgatewaysTunnelsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `IpsecgatewaysTunnelsPut`: IPSecTunnelRead
    fmt.Fprintf(os.Stdout, "Response from `IPSecTunnelsApi.IpsecgatewaysTunnelsPut`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**gatewayId** | **string** | The ID (UUID) of the IPSecGateway. | |
|**tunnelId** | **string** | The ID (UUID) of the IPSecTunnel. | |

### Other Parameters

Other parameters are passed through a pointer to an apiIpsecgatewaysTunnelsPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **iPSecTunnelEnsure** | [**IPSecTunnelEnsure**](../models/IPSecTunnelEnsure.md) | update IPSecTunnel | |

### Return type

[**IPSecTunnelRead**](../models/IPSecTunnelRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


