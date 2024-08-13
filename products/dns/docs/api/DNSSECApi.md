# \DNSSECApi

All URIs are relative to *https://dns.de-fra.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**ZonesKeysDelete**](DNSSECApi.md#ZonesKeysDelete) | **Delete** /zones/{zoneId}/keys | Delete a DNSSEC key|
|[**ZonesKeysGet**](DNSSECApi.md#ZonesKeysGet) | **Get** /zones/{zoneId}/keys | Retrieve a DNSSEC key|
|[**ZonesKeysPost**](DNSSECApi.md#ZonesKeysPost) | **Post** /zones/{zoneId}/keys | Create a DNSSEC key|



## ZonesKeysDelete

```go
var result map[string]interface{} = ZonesKeysDelete(ctx, zoneId)
                      .Execute()
```

Delete a DNSSEC key



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    dns "github.com/ionos-cloud/sdk-go-bundle/products/dns"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    zoneId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID (UUID) of the DNS zone.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := dns.NewAPIClient(configuration)
    resource, resp, err := apiClient.DNSSECApi.ZonesKeysDelete(context.Background(), zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNSSECApi.ZonesKeysDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ZonesKeysDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DNSSECApi.ZonesKeysDelete`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**zoneId** | **string** | The ID (UUID) of the DNS zone. | |

### Other Parameters

Other parameters are passed through a pointer to an apiZonesKeysDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

**map[string]interface{}**

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ZonesKeysGet

```go
var result DnssecKeyReadList = ZonesKeysGet(ctx, zoneId)
                      .Execute()
```

Retrieve a DNSSEC key



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    dns "github.com/ionos-cloud/sdk-go-bundle/products/dns"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    zoneId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID (UUID) of the DNS zone.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := dns.NewAPIClient(configuration)
    resource, resp, err := apiClient.DNSSECApi.ZonesKeysGet(context.Background(), zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNSSECApi.ZonesKeysGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ZonesKeysGet`: DnssecKeyReadList
    fmt.Fprintf(os.Stdout, "Response from `DNSSECApi.ZonesKeysGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**zoneId** | **string** | The ID (UUID) of the DNS zone. | |

### Other Parameters

Other parameters are passed through a pointer to an apiZonesKeysGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**DnssecKeyReadList**](../models/DnssecKeyReadList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ZonesKeysPost

```go
var result DnssecKeyReadCreation = ZonesKeysPost(ctx, zoneId)
                      .DnssecKeyCreate(dnssecKeyCreate)
                      .Execute()
```

Create a DNSSEC key



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    dns "github.com/ionos-cloud/sdk-go-bundle/products/dns"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    zoneId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID (UUID) of the DNS zone.
    dnssecKeyCreate := *openapiclient.NewDnssecKeyCreate(*openapiclient.NewDnssecKeyParameters(*openapiclient.NewKeyParameters(openapiclient.algorithm("RSASHA256"), openapiclient.kskBits(1024), openapiclient.zskBits(1024)), *openapiclient.NewNsecParameters(openapiclient.nsecMode("NSEC"), int32(21), int32(128)), int32(120))) // DnssecKeyCreate | Enable DNSSEC request.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := dns.NewAPIClient(configuration)
    resource, resp, err := apiClient.DNSSECApi.ZonesKeysPost(context.Background(), zoneId).DnssecKeyCreate(dnssecKeyCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNSSECApi.ZonesKeysPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ZonesKeysPost`: DnssecKeyReadCreation
    fmt.Fprintf(os.Stdout, "Response from `DNSSECApi.ZonesKeysPost`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**zoneId** | **string** | The ID (UUID) of the DNS zone. | |

### Other Parameters

Other parameters are passed through a pointer to an apiZonesKeysPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **dnssecKeyCreate** | [**DnssecKeyCreate**](../models/DnssecKeyCreate.md) | Enable DNSSEC request. | |

### Return type

[**DnssecKeyReadCreation**](../models/DnssecKeyReadCreation.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


