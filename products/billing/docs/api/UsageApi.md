# \UsageApi

All URIs are relative to *https://api.ionos.com/billing*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**UsageFindByDatacenter**](UsageApi.md#UsageFindByDatacenter) | **Get** /{contract}/usage/{dc} | |
|[**UsageGet**](UsageApi.md#UsageGet) | **Get** /{contract}/usage | |



## UsageFindByDatacenter

```go
var result UsageGet200Response = UsageFindByDatacenter(ctx, contract, dc)
                      .Period(period)
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
    dc := "ad34b997-43c7-4666-889b-57acbeaeeb8b" // string | Get data for the particular data-center ID (VDC UUID)
    period := "period_example" // string | Period of interest in format YYYY-MM (current month by default if not specified) (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := billing.NewAPIClient(configuration)
    resource, resp, err := apiClient.UsageApi.UsageFindByDatacenter(context.Background(), contract, dc).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageApi.UsageFindByDatacenter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `UsageFindByDatacenter`: UsageGet200Response
    fmt.Fprintf(os.Stdout, "Response from `UsageApi.UsageFindByDatacenter`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**contract** | **int32** | Contract number | |
|**dc** | **string** | Get data for the particular data-center ID (VDC UUID) | |

### Other Parameters

Other parameters are passed through a pointer to an apiUsageFindByDatacenterRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **period** | **string** | Period of interest in format YYYY-MM (current month by default if not specified) | |

### Return type

[**UsageGet200Response**](../models/UsageGet200Response.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## UsageGet

```go
var result UsageGet200Response = UsageGet(ctx, contract)
                      .Period(period)
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
    period := "period_example" // string | Period of interest in format YYYY-MM (current month by default if not specified) (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := billing.NewAPIClient(configuration)
    resource, resp, err := apiClient.UsageApi.UsageGet(context.Background(), contract).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsageApi.UsageGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `UsageGet`: UsageGet200Response
    fmt.Fprintf(os.Stdout, "Response from `UsageApi.UsageGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**contract** | **int32** | Contract number | |

### Other Parameters

Other parameters are passed through a pointer to an apiUsageGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **period** | **string** | Period of interest in format YYYY-MM (current month by default if not specified) | |

### Return type

[**UsageGet200Response**](../models/UsageGet200Response.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


