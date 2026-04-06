# \TrafficApi

All URIs are relative to *https://api.ionos.com/billing*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**TrafficFindByPeriod**](TrafficApi.md#TrafficFindByPeriod) | **Get** /{contract}/traffic/{period} | |
|[**TrafficGet**](TrafficApi.md#TrafficGet) | **Get** /{contract}/traffic | |



## TrafficFindByPeriod

```go
var result Traffic = TrafficFindByPeriod(ctx, contract, period)
                      .Ip(ip)
                      .Output(output)
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
    ip := true // bool | Provide report grouped by IPs (optional)
    output := "output_example" // string | The output format (object, array, CSV or all of them) (optional) (default to "all")

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := billing.NewAPIClient(configuration)
    resource, resp, err := apiClient.TrafficApi.TrafficFindByPeriod(context.Background(), contract, period).Ip(ip).Output(output).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrafficApi.TrafficFindByPeriod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `TrafficFindByPeriod`: Traffic
    fmt.Fprintf(os.Stdout, "Response from `TrafficApi.TrafficFindByPeriod`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**contract** | **int32** | Contract number | |
|**period** | **string** | Period of interest in format YYYY-MM | |

### Other Parameters

Other parameters are passed through a pointer to an apiTrafficFindByPeriodRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **ip** | **bool** | Provide report grouped by IPs | |
| **output** | **string** | The output format (object, array, CSV or all of them) | [default to &quot;all&quot;]|

### Return type

[**Traffic**](../models/Traffic.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## TrafficGet

```go
var result Traffic = TrafficGet(ctx, contract)
                      .Ip(ip)
                      .Output(output)
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
    ip := true // bool | Provide report grouped by IPs (optional)
    output := "output_example" // string | The output format (object, array, CSV or all of them) (optional) (default to "all")

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := billing.NewAPIClient(configuration)
    resource, resp, err := apiClient.TrafficApi.TrafficGet(context.Background(), contract).Ip(ip).Output(output).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrafficApi.TrafficGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `TrafficGet`: Traffic
    fmt.Fprintf(os.Stdout, "Response from `TrafficApi.TrafficGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**contract** | **int32** | Contract number | |

### Other Parameters

Other parameters are passed through a pointer to an apiTrafficGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **ip** | **bool** | Provide report grouped by IPs | |
| **output** | **string** | The output format (object, array, CSV or all of them) | [default to &quot;all&quot;]|

### Return type

[**Traffic**](../models/Traffic.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


