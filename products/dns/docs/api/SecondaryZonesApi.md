# \SecondaryZonesApi

All URIs are relative to *https://dns.de-fra.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**SecondaryzonesAxfrGet**](SecondaryZonesApi.md#SecondaryzonesAxfrGet) | **Get** /secondaryzones/{secondaryZoneId}/axfr | Get status of zone transfer|
|[**SecondaryzonesAxfrPut**](SecondaryZonesApi.md#SecondaryzonesAxfrPut) | **Put** /secondaryzones/{secondaryZoneId}/axfr | Start zone transfer|
|[**SecondaryzonesDelete**](SecondaryZonesApi.md#SecondaryzonesDelete) | **Delete** /secondaryzones/{secondaryZoneId} | Delete a secondary zone|
|[**SecondaryzonesFindById**](SecondaryZonesApi.md#SecondaryzonesFindById) | **Get** /secondaryzones/{secondaryZoneId} | Retrieve a secondary zone|
|[**SecondaryzonesGet**](SecondaryZonesApi.md#SecondaryzonesGet) | **Get** /secondaryzones | Retrieve secondary zones|
|[**SecondaryzonesPost**](SecondaryZonesApi.md#SecondaryzonesPost) | **Post** /secondaryzones | Create a secondary zone|
|[**SecondaryzonesPut**](SecondaryZonesApi.md#SecondaryzonesPut) | **Put** /secondaryzones/{secondaryZoneId} | Update a secondary zone|



## SecondaryzonesAxfrGet

```go
var result ZoneTransferPrimaryIpsStatus = SecondaryzonesAxfrGet(ctx, secondaryZoneId)
                      .Execute()
```

Get status of zone transfer



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
    secondaryZoneId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID (UUID) of the DNS zone.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := dns.NewAPIClient(configuration)
    resource, resp, err := apiClient.SecondaryZonesApi.SecondaryzonesAxfrGet(context.Background(), secondaryZoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecondaryZonesApi.SecondaryzonesAxfrGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `SecondaryzonesAxfrGet`: ZoneTransferPrimaryIpsStatus
    fmt.Fprintf(os.Stdout, "Response from `SecondaryZonesApi.SecondaryzonesAxfrGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**secondaryZoneId** | **string** | The ID (UUID) of the DNS zone. | |

### Other Parameters

Other parameters are passed through a pointer to an apiSecondaryzonesAxfrGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**ZoneTransferPrimaryIpsStatus**](../models/ZoneTransferPrimaryIpsStatus.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## SecondaryzonesAxfrPut

```go
var result map[string]interface{} = SecondaryzonesAxfrPut(ctx, secondaryZoneId)
                      .Execute()
```

Start zone transfer



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
    secondaryZoneId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID (UUID) of the DNS zone.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := dns.NewAPIClient(configuration)
    resource, resp, err := apiClient.SecondaryZonesApi.SecondaryzonesAxfrPut(context.Background(), secondaryZoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecondaryZonesApi.SecondaryzonesAxfrPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `SecondaryzonesAxfrPut`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SecondaryZonesApi.SecondaryzonesAxfrPut`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**secondaryZoneId** | **string** | The ID (UUID) of the DNS zone. | |

### Other Parameters

Other parameters are passed through a pointer to an apiSecondaryzonesAxfrPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

**map[string]interface{}**

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## SecondaryzonesDelete

```go
var result map[string]interface{} = SecondaryzonesDelete(ctx, secondaryZoneId)
                      .Execute()
```

Delete a secondary zone



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
    secondaryZoneId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID (UUID) of the DNS zone.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := dns.NewAPIClient(configuration)
    resp, err := apiClient.SecondaryZonesApi.SecondaryzonesDelete(context.Background(), secondaryZoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecondaryZonesApi.SecondaryzonesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `SecondaryzonesDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SecondaryZonesApi.SecondaryzonesDelete`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**secondaryZoneId** | **string** | The ID (UUID) of the DNS zone. | |

### Other Parameters

Other parameters are passed through a pointer to an apiSecondaryzonesDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

**map[string]interface{}**

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## SecondaryzonesFindById

```go
var result SecondaryZoneRead = SecondaryzonesFindById(ctx, secondaryZoneId)
                      .Execute()
```

Retrieve a secondary zone



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
    secondaryZoneId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID (UUID) of the DNS zone.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := dns.NewAPIClient(configuration)
    resource, resp, err := apiClient.SecondaryZonesApi.SecondaryzonesFindById(context.Background(), secondaryZoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecondaryZonesApi.SecondaryzonesFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `SecondaryzonesFindById`: SecondaryZoneRead
    fmt.Fprintf(os.Stdout, "Response from `SecondaryZonesApi.SecondaryzonesFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**secondaryZoneId** | **string** | The ID (UUID) of the DNS zone. | |

### Other Parameters

Other parameters are passed through a pointer to an apiSecondaryzonesFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**SecondaryZoneRead**](../models/SecondaryZoneRead.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## SecondaryzonesGet

```go
var result SecondaryZoneReadList = SecondaryzonesGet(ctx)
                      .FilterState(filterState)
                      .FilterZoneName(filterZoneName)
                      .Offset(offset)
                      .Limit(limit)
                      .Execute()
```

Retrieve secondary zones



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
    filterState := openapiclient.ProvisioningState("PROVISIONING") // ProvisioningState | Filter used to fetch all zones in a particular state. (optional)
    filterZoneName := "example.com" // string | Filter used to fetch only the zones that contain the specified zone name. (optional)
    offset := int32(56) // int32 | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. (optional) (default to 0)
    limit := int32(56) // int32 | The maximum number of elements to return. Use together with offset for pagination. (optional) (default to 100)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := dns.NewAPIClient(configuration)
    resource, resp, err := apiClient.SecondaryZonesApi.SecondaryzonesGet(context.Background()).FilterState(filterState).FilterZoneName(filterZoneName).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecondaryZonesApi.SecondaryzonesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `SecondaryzonesGet`: SecondaryZoneReadList
    fmt.Fprintf(os.Stdout, "Response from `SecondaryZonesApi.SecondaryzonesGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiSecondaryzonesGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **filterState** | [**ProvisioningState**](../models/.md) | Filter used to fetch all zones in a particular state. | |
| **filterZoneName** | **string** | Filter used to fetch only the zones that contain the specified zone name. | |
| **offset** | **int32** | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return. Use together with offset for pagination. | [default to 100]|

### Return type

[**SecondaryZoneReadList**](../models/SecondaryZoneReadList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## SecondaryzonesPost

```go
var result SecondaryZoneRead = SecondaryzonesPost(ctx)
                      .SecondaryZoneCreate(secondaryZoneCreate)
                      .Execute()
```

Create a secondary zone



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
    secondaryZoneCreate := *openapiclient.NewSecondaryZoneCreate(*openapiclient.NewSecondaryZone("example.com", []string{"PrimaryIps_example"})) // SecondaryZoneCreate | zone

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := dns.NewAPIClient(configuration)
    resource, resp, err := apiClient.SecondaryZonesApi.SecondaryzonesPost(context.Background()).SecondaryZoneCreate(secondaryZoneCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecondaryZonesApi.SecondaryzonesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `SecondaryzonesPost`: SecondaryZoneRead
    fmt.Fprintf(os.Stdout, "Response from `SecondaryZonesApi.SecondaryzonesPost`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiSecondaryzonesPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **secondaryZoneCreate** | [**SecondaryZoneCreate**](../models/SecondaryZoneCreate.md) | zone | |

### Return type

[**SecondaryZoneRead**](../models/SecondaryZoneRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## SecondaryzonesPut

```go
var result SecondaryZoneRead = SecondaryzonesPut(ctx, secondaryZoneId)
                      .SecondaryZoneEnsure(secondaryZoneEnsure)
                      .Execute()
```

Update a secondary zone



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
    secondaryZoneId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID (UUID) of the DNS zone.
    secondaryZoneEnsure := *openapiclient.NewSecondaryZoneEnsure(*openapiclient.NewSecondaryZone("example.com", []string{"PrimaryIps_example"})) // SecondaryZoneEnsure | update zone

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := dns.NewAPIClient(configuration)
    resource, resp, err := apiClient.SecondaryZonesApi.SecondaryzonesPut(context.Background(), secondaryZoneId).SecondaryZoneEnsure(secondaryZoneEnsure).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecondaryZonesApi.SecondaryzonesPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `SecondaryzonesPut`: SecondaryZoneRead
    fmt.Fprintf(os.Stdout, "Response from `SecondaryZonesApi.SecondaryzonesPut`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**secondaryZoneId** | **string** | The ID (UUID) of the DNS zone. | |

### Other Parameters

Other parameters are passed through a pointer to an apiSecondaryzonesPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **secondaryZoneEnsure** | [**SecondaryZoneEnsure**](../models/SecondaryZoneEnsure.md) | update zone | |

### Return type

[**SecondaryZoneRead**](../models/SecondaryZoneRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


