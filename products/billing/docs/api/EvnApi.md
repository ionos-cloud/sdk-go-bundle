# \EvnApi

All URIs are relative to *https://api.ionos.com/billing*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**EvnFindByPeriod**](EvnApi.md#EvnFindByPeriod) | **Get** /{contract}/evn/{period} | |
|[**EvnGet**](EvnApi.md#EvnGet) | **Get** /{contract}/evn | |



## EvnFindByPeriod

```go
var result Evn = EvnFindByPeriod(ctx, contract, period)
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

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := billing.NewAPIClient(configuration)
    resource, resp, err := apiClient.EvnApi.EvnFindByPeriod(context.Background(), contract, period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EvnApi.EvnFindByPeriod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `EvnFindByPeriod`: Evn
    fmt.Fprintf(os.Stdout, "Response from `EvnApi.EvnFindByPeriod`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**contract** | **int32** | Contract number | |
|**period** | **string** | Period of interest in format YYYY-MM | |

### Other Parameters

Other parameters are passed through a pointer to an apiEvnFindByPeriodRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**Evn**](../models/Evn.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## EvnGet

```go
var result Evn = EvnGet(ctx, contract)
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

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := billing.NewAPIClient(configuration)
    resource, resp, err := apiClient.EvnApi.EvnGet(context.Background(), contract).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EvnApi.EvnGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `EvnGet`: Evn
    fmt.Fprintf(os.Stdout, "Response from `EvnApi.EvnGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**contract** | **int32** | Contract number | |

### Other Parameters

Other parameters are passed through a pointer to an apiEvnGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**Evn**](../models/Evn.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


