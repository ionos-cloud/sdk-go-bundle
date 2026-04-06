# \ProductsApi

All URIs are relative to *https://api.ionos.com/billing*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**ProductsGet**](ProductsApi.md#ProductsGet) | **Get** /{contract}/products | |



## ProductsGet

```go
var result ProductsGet200Response = ProductsGet(ctx, contract)
                      .Date(date)
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
    date := time.Now() // string | Date of interest in format of YYYY-MM-DD (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := billing.NewAPIClient(configuration)
    resource, resp, err := apiClient.ProductsApi.ProductsGet(context.Background(), contract).Date(date).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.ProductsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ProductsGet`: ProductsGet200Response
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.ProductsGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**contract** | **int32** | Contract number | |

### Other Parameters

Other parameters are passed through a pointer to an apiProductsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **date** | **string** | Date of interest in format of YYYY-MM-DD | |

### Return type

[**ProductsGet200Response**](../models/ProductsGet200Response.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


