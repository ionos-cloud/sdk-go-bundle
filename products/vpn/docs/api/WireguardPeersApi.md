# \WireguardPeersApi

All URIs are relative to *https://vpn.de-fra.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**WireguardgatewaysPeersDelete**](WireguardPeersApi.md#WireguardgatewaysPeersDelete) | **Delete** /wireguardgateways/{gatewayId}/peers/{peerId} | Delete WireguardPeer|
|[**WireguardgatewaysPeersFindById**](WireguardPeersApi.md#WireguardgatewaysPeersFindById) | **Get** /wireguardgateways/{gatewayId}/peers/{peerId} | Retrieve WireguardPeer|
|[**WireguardgatewaysPeersGet**](WireguardPeersApi.md#WireguardgatewaysPeersGet) | **Get** /wireguardgateways/{gatewayId}/peers | Retrieve all WireguardPeers|
|[**WireguardgatewaysPeersPost**](WireguardPeersApi.md#WireguardgatewaysPeersPost) | **Post** /wireguardgateways/{gatewayId}/peers | Create WireguardPeer|
|[**WireguardgatewaysPeersPut**](WireguardPeersApi.md#WireguardgatewaysPeersPut) | **Put** /wireguardgateways/{gatewayId}/peers/{peerId} | Ensure WireguardPeer|



## WireguardgatewaysPeersDelete

```go
var result  = WireguardgatewaysPeersDelete(ctx, gatewayId, peerId)
                      .Execute()
```

Delete WireguardPeer



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
    peerId := "b62b3a40-adee-5b6c-b98d-be20bfcbdd91" // string | The ID (UUID) of the WireguardPeer.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := vpn.NewAPIClient(configuration)
    resp, err := apiClient.WireguardPeersApi.WireguardgatewaysPeersDelete(context.Background(), gatewayId, peerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WireguardPeersApi.WireguardgatewaysPeersDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**gatewayId** | **string** | The ID (UUID) of the WireguardGateway. | |
|**peerId** | **string** | The ID (UUID) of the WireguardPeer. | |

### Other Parameters

Other parameters are passed through a pointer to an apiWireguardgatewaysPeersDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## WireguardgatewaysPeersFindById

```go
var result WireguardPeerRead = WireguardgatewaysPeersFindById(ctx, gatewayId, peerId)
                      .Execute()
```

Retrieve WireguardPeer



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
    peerId := "b62b3a40-adee-5b6c-b98d-be20bfcbdd91" // string | The ID (UUID) of the WireguardPeer.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := vpn.NewAPIClient(configuration)
    resource, resp, err := apiClient.WireguardPeersApi.WireguardgatewaysPeersFindById(context.Background(), gatewayId, peerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WireguardPeersApi.WireguardgatewaysPeersFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `WireguardgatewaysPeersFindById`: WireguardPeerRead
    fmt.Fprintf(os.Stdout, "Response from `WireguardPeersApi.WireguardgatewaysPeersFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**gatewayId** | **string** | The ID (UUID) of the WireguardGateway. | |
|**peerId** | **string** | The ID (UUID) of the WireguardPeer. | |

### Other Parameters

Other parameters are passed through a pointer to an apiWireguardgatewaysPeersFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**WireguardPeerRead**](../models/WireguardPeerRead.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## WireguardgatewaysPeersGet

```go
var result WireguardPeerReadList = WireguardgatewaysPeersGet(ctx, gatewayId)
                      .Offset(offset)
                      .Limit(limit)
                      .Execute()
```

Retrieve all WireguardPeers



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
    offset := int32(0) // int32 | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. (optional) (default to 0)
    limit := int32(100) // int32 | The maximum number of elements to return. Use together with offset for pagination. (optional) (default to 100)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := vpn.NewAPIClient(configuration)
    resource, resp, err := apiClient.WireguardPeersApi.WireguardgatewaysPeersGet(context.Background(), gatewayId).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WireguardPeersApi.WireguardgatewaysPeersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `WireguardgatewaysPeersGet`: WireguardPeerReadList
    fmt.Fprintf(os.Stdout, "Response from `WireguardPeersApi.WireguardgatewaysPeersGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**gatewayId** | **string** | The ID (UUID) of the WireguardGateway. | |

### Other Parameters

Other parameters are passed through a pointer to an apiWireguardgatewaysPeersGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **offset** | **int32** | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return. Use together with offset for pagination. | [default to 100]|

### Return type

[**WireguardPeerReadList**](../models/WireguardPeerReadList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## WireguardgatewaysPeersPost

```go
var result WireguardPeerRead = WireguardgatewaysPeersPost(ctx, gatewayId)
                      .WireguardPeerCreate(wireguardPeerCreate)
                      .Execute()
```

Create WireguardPeer



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
    wireguardPeerCreate := *openapiclient.NewWireguardPeerCreate(*openapiclient.NewWireguardPeer("My Company Gateway Peer", []string{"AllowedIPs_example"}, "no8iaSEoqfbI6PVYsdEiUU5efYdtKX8VAhKity19MWI=")) // WireguardPeerCreate | WireguardPeer to create.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := vpn.NewAPIClient(configuration)
    resource, resp, err := apiClient.WireguardPeersApi.WireguardgatewaysPeersPost(context.Background(), gatewayId).WireguardPeerCreate(wireguardPeerCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WireguardPeersApi.WireguardgatewaysPeersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `WireguardgatewaysPeersPost`: WireguardPeerRead
    fmt.Fprintf(os.Stdout, "Response from `WireguardPeersApi.WireguardgatewaysPeersPost`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**gatewayId** | **string** | The ID (UUID) of the WireguardGateway. | |

### Other Parameters

Other parameters are passed through a pointer to an apiWireguardgatewaysPeersPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **wireguardPeerCreate** | [**WireguardPeerCreate**](../models/WireguardPeerCreate.md) | WireguardPeer to create. | |

### Return type

[**WireguardPeerRead**](../models/WireguardPeerRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## WireguardgatewaysPeersPut

```go
var result WireguardPeerRead = WireguardgatewaysPeersPut(ctx, gatewayId, peerId)
                      .WireguardPeerEnsure(wireguardPeerEnsure)
                      .Execute()
```

Ensure WireguardPeer



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
    peerId := "b62b3a40-adee-5b6c-b98d-be20bfcbdd91" // string | The ID (UUID) of the WireguardPeer.
    wireguardPeerEnsure := *openapiclient.NewWireguardPeerEnsure("b62b3a40-adee-5b6c-b98d-be20bfcbdd91", *openapiclient.NewWireguardPeer("My Company Gateway Peer", []string{"AllowedIPs_example"}, "no8iaSEoqfbI6PVYsdEiUU5efYdtKX8VAhKity19MWI=")) // WireguardPeerEnsure | update WireguardPeer

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := vpn.NewAPIClient(configuration)
    resource, resp, err := apiClient.WireguardPeersApi.WireguardgatewaysPeersPut(context.Background(), gatewayId, peerId).WireguardPeerEnsure(wireguardPeerEnsure).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WireguardPeersApi.WireguardgatewaysPeersPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `WireguardgatewaysPeersPut`: WireguardPeerRead
    fmt.Fprintf(os.Stdout, "Response from `WireguardPeersApi.WireguardgatewaysPeersPut`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**gatewayId** | **string** | The ID (UUID) of the WireguardGateway. | |
|**peerId** | **string** | The ID (UUID) of the WireguardPeer. | |

### Other Parameters

Other parameters are passed through a pointer to an apiWireguardgatewaysPeersPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **wireguardPeerEnsure** | [**WireguardPeerEnsure**](../models/WireguardPeerEnsure.md) | update WireguardPeer | |

### Return type

[**WireguardPeerRead**](../models/WireguardPeerRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


