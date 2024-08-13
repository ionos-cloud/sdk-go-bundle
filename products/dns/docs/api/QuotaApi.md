# \QuotaApi

All URIs are relative to *https://dns.de-fra.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**QuotaGet**](QuotaApi.md#QuotaGet) | **Get** /quota | Retrieve resources quota|



## QuotaGet

```go
var result Quota = QuotaGet(ctx)
                      .Execute()
```

Retrieve resources quota



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

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := dns.NewAPIClient(configuration)
    resource, resp, err := apiClient.QuotaApi.QuotaGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotaApi.QuotaGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `QuotaGet`: Quota
    fmt.Fprintf(os.Stdout, "Response from `QuotaApi.QuotaGet`: %v\n", resource)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to an apiQuotaGetRequest struct via the builder pattern


### Return type

[**Quota**](../models/Quota.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


