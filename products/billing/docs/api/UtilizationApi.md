# \UtilizationApi

All URIs are relative to *https://api.ionos.com/billing*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**UtilizationDailyFindByDate**](UtilizationApi.md#UtilizationDailyFindByDate) | **Get** /{contract}/utilization/daily/{date} | |
|[**UtilizationFindByPeriod**](UtilizationApi.md#UtilizationFindByPeriod) | **Get** /{contract}/utilization/{period} | |
|[**UtilizationGet**](UtilizationApi.md#UtilizationGet) | **Get** /{contract}/utilization | |



## UtilizationDailyFindByDate

```go
var result UtilizationDailyFindByDate200Response = UtilizationDailyFindByDate(ctx, contract, date)
                      .Dc(dc)
                      .Resource(resource)
                      .Type_(type_)
                      .Execute()
```





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    billing "github.com/ionos-cloud/sdk-go-bundle/products/billing"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    contract := int32(56) // int32 | Contract number
    date := "2025-09-15" // string | Date of interest in format YYYY-MM-DD
    dc := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter out by particular VDC UUID (or \"Bucket\" UUID for non-vdc related resources) (optional)
    resource := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter out by particular resource UUID (e.g.VM UUID, NIC UUID) (optional)
    type_ := openapiclient.ResourceType("SERVER") // ResourceType | Filter out by particular type of items (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := billing.NewAPIClient(configuration)
    resource, resp, err := apiClient.UtilizationApi.UtilizationDailyFindByDate(context.Background(), contract, date).Dc(dc).Resource(resource).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UtilizationApi.UtilizationDailyFindByDate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `UtilizationDailyFindByDate`: UtilizationDailyFindByDate200Response
    fmt.Fprintf(os.Stdout, "Response from `UtilizationApi.UtilizationDailyFindByDate`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**contract** | **int32** | Contract number | |
|**date** | **string** | Date of interest in format YYYY-MM-DD | |

### Other Parameters

Other parameters are passed through a pointer to an apiUtilizationDailyFindByDateRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **dc** | **string** | Filter out by particular VDC UUID (or \&quot;Bucket\&quot; UUID for non-vdc related resources) | |
| **resource** | **string** | Filter out by particular resource UUID (e.g.VM UUID, NIC UUID) | |
| **type_** | [**ResourceType**](../models/.md) | Filter out by particular type of items | |

### Return type

[**UtilizationDailyFindByDate200Response**](../models/UtilizationDailyFindByDate200Response.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## UtilizationFindByPeriod

```go
var result UtilizationGet200Response = UtilizationFindByPeriod(ctx, contract, period)
                      .Dc(dc)
                      .Resource(resource)
                      .Type_(type_)
                      .Execute()
```





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    billing "github.com/ionos-cloud/sdk-go-bundle/products/billing"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    contract := int32(56) // int32 | Contract number
    period := "2020-01" // string | Period of interest in format YYYY-MM
    dc := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter out by particular VDC UUID (or \"Bucket\" UUID for non-vdc related resources) (optional)
    resource := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter out by particular resource UUID (e.g.VM UUID, NIC UUID) (optional)
    type_ := openapiclient.ResourceType("SERVER") // ResourceType | Filter out by particular type of items (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := billing.NewAPIClient(configuration)
    resource, resp, err := apiClient.UtilizationApi.UtilizationFindByPeriod(context.Background(), contract, period).Dc(dc).Resource(resource).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UtilizationApi.UtilizationFindByPeriod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `UtilizationFindByPeriod`: UtilizationGet200Response
    fmt.Fprintf(os.Stdout, "Response from `UtilizationApi.UtilizationFindByPeriod`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**contract** | **int32** | Contract number | |
|**period** | **string** | Period of interest in format YYYY-MM | |

### Other Parameters

Other parameters are passed through a pointer to an apiUtilizationFindByPeriodRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **dc** | **string** | Filter out by particular VDC UUID (or \&quot;Bucket\&quot; UUID for non-vdc related resources) | |
| **resource** | **string** | Filter out by particular resource UUID (e.g.VM UUID, NIC UUID) | |
| **type_** | [**ResourceType**](../models/.md) | Filter out by particular type of items | |

### Return type

[**UtilizationGet200Response**](../models/UtilizationGet200Response.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## UtilizationGet

```go
var result UtilizationGet200Response = UtilizationGet(ctx, contract)
                      .Dc(dc)
                      .Resource(resource)
                      .Type_(type_)
                      .Execute()
```





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    billing "github.com/ionos-cloud/sdk-go-bundle/products/billing"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    contract := int32(56) // int32 | Contract number
    dc := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter out by particular VDC UUID (or \"Bucket\" UUID for non-vdc related resources) (optional)
    resource := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter out by particular resource UUID (e.g.VM UUID, NIC UUID) (optional)
    type_ := openapiclient.ResourceType("SERVER") // ResourceType | Filter out by particular type of items (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := billing.NewAPIClient(configuration)
    resource, resp, err := apiClient.UtilizationApi.UtilizationGet(context.Background(), contract).Dc(dc).Resource(resource).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UtilizationApi.UtilizationGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `UtilizationGet`: UtilizationGet200Response
    fmt.Fprintf(os.Stdout, "Response from `UtilizationApi.UtilizationGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**contract** | **int32** | Contract number | |

### Other Parameters

Other parameters are passed through a pointer to an apiUtilizationGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **dc** | **string** | Filter out by particular VDC UUID (or \&quot;Bucket\&quot; UUID for non-vdc related resources) | |
| **resource** | **string** | Filter out by particular resource UUID (e.g.VM UUID, NIC UUID) | |
| **type_** | [**ResourceType**](../models/.md) | Filter out by particular type of items | |

### Return type

[**UtilizationGet200Response**](../models/UtilizationGet200Response.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


