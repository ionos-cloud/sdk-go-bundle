# \InvoicesApi

All URIs are relative to *https://api.ionos.com/billing*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**InvoicesFindById**](InvoicesApi.md#InvoicesFindById) | **Get** /{contract}/invoices/{id} | |
|[**InvoicesFindByPeriod**](InvoicesApi.md#InvoicesFindByPeriod) | **Get** /invoices/{period} | |
|[**InvoicesGet**](InvoicesApi.md#InvoicesGet) | **Get** /{contract}/invoices | |



## InvoicesFindById

```go
var result Invoice = InvoicesFindById(ctx, contract, id)
                      .Dateformat(dateformat)
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
    id := "id_example" // string | Invoice ID
    dateformat := "dateformat_example" // string | Vendor date format (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := billing.NewAPIClient(configuration)
    resource, resp, err := apiClient.InvoicesApi.InvoicesFindById(context.Background(), contract, id).Dateformat(dateformat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.InvoicesFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `InvoicesFindById`: Invoice
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.InvoicesFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**contract** | **int32** | Contract number | |
|**id** | **string** | Invoice ID | |

### Other Parameters

Other parameters are passed through a pointer to an apiInvoicesFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **dateformat** | **string** | Vendor date format | |

### Return type

[**Invoice**](../models/Invoice.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## InvoicesFindByPeriod

```go
var result []Invoice = InvoicesFindByPeriod(ctx, period)
                      .Contractid(contractid)
                      .Dateformat(dateformat)
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
    period := "2020-01" // string | Period of interest in format YYYY-MM
    contractid := "contractid_example" // string | Filter out the exact contractID (optional)
    dateformat := "dateformat_example" // string | Vendor date format (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := billing.NewAPIClient(configuration)
    resource, resp, err := apiClient.InvoicesApi.InvoicesFindByPeriod(context.Background(), period).Contractid(contractid).Dateformat(dateformat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.InvoicesFindByPeriod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `InvoicesFindByPeriod`: []Invoice
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.InvoicesFindByPeriod`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**period** | **string** | Period of interest in format YYYY-MM | |

### Other Parameters

Other parameters are passed through a pointer to an apiInvoicesFindByPeriodRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **contractid** | **string** | Filter out the exact contractID | |
| **dateformat** | **string** | Vendor date format | |

### Return type

[**[]Invoice**](../models/Invoice.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## InvoicesGet

```go
var result InvoicesGet200Response = InvoicesGet(ctx, contract)
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
    resource, resp, err := apiClient.InvoicesApi.InvoicesGet(context.Background(), contract).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.InvoicesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `InvoicesGet`: InvoicesGet200Response
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.InvoicesGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**contract** | **int32** | Contract number | |

### Other Parameters

Other parameters are passed through a pointer to an apiInvoicesGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**InvoicesGet200Response**](../models/InvoicesGet200Response.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


