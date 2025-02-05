# \RecordsApi

All URIs are relative to *https://dns.de-fra.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**RecordsGet**](RecordsApi.md#RecordsGet) | **Get** /records | Retrieve all records from primary zones|
|[**SecondaryzonesRecordsGet**](RecordsApi.md#SecondaryzonesRecordsGet) | **Get** /secondaryzones/{secondaryZoneId}/records | Retrieve records for a secondary zone|
|[**ZonesRecordsDelete**](RecordsApi.md#ZonesRecordsDelete) | **Delete** /zones/{zoneId}/records/{recordId} | Delete a record|
|[**ZonesRecordsFindById**](RecordsApi.md#ZonesRecordsFindById) | **Get** /zones/{zoneId}/records/{recordId} | Retrieve a record|
|[**ZonesRecordsGet**](RecordsApi.md#ZonesRecordsGet) | **Get** /zones/{zoneId}/records | Retrieve records|
|[**ZonesRecordsPost**](RecordsApi.md#ZonesRecordsPost) | **Post** /zones/{zoneId}/records | Create a record|
|[**ZonesRecordsPut**](RecordsApi.md#ZonesRecordsPut) | **Put** /zones/{zoneId}/records/{recordId} | Update a record|



## RecordsGet

```go
var result RecordReadList = RecordsGet(ctx)
                      .FilterZoneId(filterZoneId)
                      .FilterName(filterName)
                      .FilterState(filterState)
                      .FilterType(filterType)
                      .Offset(offset)
                      .Limit(limit)
                      .Execute()
```

Retrieve all records from primary zones



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
    filterZoneId := "1d6ca576-7162-4700-8df7-208bbe28fc44" // string | Filter used to fetch only the records that contain specified zoneId. (optional)
    filterName := "app" // string | Filter used to fetch only the records that contain specified record name. (optional)
    filterState := openapiclient.ProvisioningState("PROVISIONING") // ProvisioningState | Filter used to fetch only the records that are in certain state. (optional)
    filterType := openapiclient.RecordType("A") // RecordType | Filter used to fetch only the records with specified type. (optional)
    offset := int32(56) // int32 | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. (optional) (default to 0)
    limit := int32(56) // int32 | The maximum number of elements to return. Use together with offset for pagination. (optional) (default to 100)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := dns.NewAPIClient(configuration)
    resource, resp, err := apiClient.RecordsApi.RecordsGet(context.Background()).FilterZoneId(filterZoneId).FilterName(filterName).FilterState(filterState).FilterType(filterType).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordsApi.RecordsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `RecordsGet`: RecordReadList
    fmt.Fprintf(os.Stdout, "Response from `RecordsApi.RecordsGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiRecordsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **filterZoneId** | **string** | Filter used to fetch only the records that contain specified zoneId. | |
| **filterName** | **string** | Filter used to fetch only the records that contain specified record name. | |
| **filterState** | [**ProvisioningState**](../models/.md) | Filter used to fetch only the records that are in certain state. | |
| **filterType** | [**RecordType**](../models/.md) | Filter used to fetch only the records with specified type. | |
| **offset** | **int32** | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return. Use together with offset for pagination. | [default to 100]|

### Return type

[**RecordReadList**](../models/RecordReadList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## SecondaryzonesRecordsGet

```go
var result SecondaryZoneRecordReadList = SecondaryzonesRecordsGet(ctx, secondaryZoneId)
                      .Offset(offset)
                      .Limit(limit)
                      .Execute()
```

Retrieve records for a secondary zone



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
    secondaryZoneId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID (UUID) of the DNS secondary zone.
    offset := int32(56) // int32 | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. (optional) (default to 0)
    limit := int32(56) // int32 | The maximum number of elements to return. Use together with offset for pagination. (optional) (default to 100)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := dns.NewAPIClient(configuration)
    resource, resp, err := apiClient.RecordsApi.SecondaryzonesRecordsGet(context.Background(), secondaryZoneId).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordsApi.SecondaryzonesRecordsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `SecondaryzonesRecordsGet`: SecondaryZoneRecordReadList
    fmt.Fprintf(os.Stdout, "Response from `RecordsApi.SecondaryzonesRecordsGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**secondaryZoneId** | **string** | The ID (UUID) of the DNS secondary zone. | |

### Other Parameters

Other parameters are passed through a pointer to an apiSecondaryzonesRecordsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **offset** | **int32** | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return. Use together with offset for pagination. | [default to 100]|

### Return type

[**SecondaryZoneRecordReadList**](../models/SecondaryZoneRecordReadList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ZonesRecordsDelete

```go
var result map[string]interface{} = ZonesRecordsDelete(ctx, zoneId, recordId)
                      .Execute()
```

Delete a record



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
    recordId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID (UUID) of the record.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := dns.NewAPIClient(configuration)
    resp, err := apiClient.RecordsApi.ZonesRecordsDelete(context.Background(), zoneId, recordId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordsApi.ZonesRecordsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ZonesRecordsDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RecordsApi.ZonesRecordsDelete`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**zoneId** | **string** | The ID (UUID) of the DNS zone. | |
|**recordId** | **string** | The ID (UUID) of the record. | |

### Other Parameters

Other parameters are passed through a pointer to an apiZonesRecordsDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

**map[string]interface{}**

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ZonesRecordsFindById

```go
var result RecordRead = ZonesRecordsFindById(ctx, zoneId, recordId)
                      .Execute()
```

Retrieve a record



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
    recordId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID (UUID) of the record.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := dns.NewAPIClient(configuration)
    resource, resp, err := apiClient.RecordsApi.ZonesRecordsFindById(context.Background(), zoneId, recordId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordsApi.ZonesRecordsFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ZonesRecordsFindById`: RecordRead
    fmt.Fprintf(os.Stdout, "Response from `RecordsApi.ZonesRecordsFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**zoneId** | **string** | The ID (UUID) of the DNS zone. | |
|**recordId** | **string** | The ID (UUID) of the record. | |

### Other Parameters

Other parameters are passed through a pointer to an apiZonesRecordsFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**RecordRead**](../models/RecordRead.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ZonesRecordsGet

```go
var result RecordReadList = ZonesRecordsGet(ctx, zoneId)
                      .Execute()
```

Retrieve records



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
    resource, resp, err := apiClient.RecordsApi.ZonesRecordsGet(context.Background(), zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordsApi.ZonesRecordsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ZonesRecordsGet`: RecordReadList
    fmt.Fprintf(os.Stdout, "Response from `RecordsApi.ZonesRecordsGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**zoneId** | **string** | The ID (UUID) of the DNS zone. | |

### Other Parameters

Other parameters are passed through a pointer to an apiZonesRecordsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**RecordReadList**](../models/RecordReadList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ZonesRecordsPost

```go
var result RecordRead = ZonesRecordsPost(ctx, zoneId)
                      .RecordCreate(recordCreate)
                      .Execute()
```

Create a record



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
    recordCreate := *openapiclient.NewRecordCreate(*openapiclient.NewRecord("app", openapiclient.RecordType("A"), "1.2.3.4")) // RecordCreate | record

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := dns.NewAPIClient(configuration)
    resource, resp, err := apiClient.RecordsApi.ZonesRecordsPost(context.Background(), zoneId).RecordCreate(recordCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordsApi.ZonesRecordsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ZonesRecordsPost`: RecordRead
    fmt.Fprintf(os.Stdout, "Response from `RecordsApi.ZonesRecordsPost`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**zoneId** | **string** | The ID (UUID) of the DNS zone. | |

### Other Parameters

Other parameters are passed through a pointer to an apiZonesRecordsPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **recordCreate** | [**RecordCreate**](../models/RecordCreate.md) | record | |

### Return type

[**RecordRead**](../models/RecordRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## ZonesRecordsPut

```go
var result RecordRead = ZonesRecordsPut(ctx, zoneId, recordId)
                      .RecordEnsure(recordEnsure)
                      .Execute()
```

Update a record



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
    recordId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID (UUID) of the DNS record.
    recordEnsure := *openapiclient.NewRecordEnsure(*openapiclient.NewRecord("app", openapiclient.RecordType("A"), "1.2.3.4")) // RecordEnsure | 

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := dns.NewAPIClient(configuration)
    resource, resp, err := apiClient.RecordsApi.ZonesRecordsPut(context.Background(), zoneId, recordId).RecordEnsure(recordEnsure).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordsApi.ZonesRecordsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ZonesRecordsPut`: RecordRead
    fmt.Fprintf(os.Stdout, "Response from `RecordsApi.ZonesRecordsPut`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**zoneId** | **string** | The ID (UUID) of the DNS zone. | |
|**recordId** | **string** | The ID (UUID) of the DNS record. | |

### Other Parameters

Other parameters are passed through a pointer to an apiZonesRecordsPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **recordEnsure** | [**RecordEnsure**](../models/RecordEnsure.md) |  | |

### Return type

[**RecordRead**](../models/RecordRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


