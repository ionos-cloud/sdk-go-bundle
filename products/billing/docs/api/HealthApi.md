# \HealthApi

All URIs are relative to *https://api.ionos.com/billing*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**HealthCheckGet**](HealthApi.md#HealthCheckGet) | **Get** /intern/ping | |



## HealthCheckGet

```go
var result  = HealthCheckGet(ctx)
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

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := billing.NewAPIClient(configuration)
    resource, resp, err := apiClient.HealthApi.HealthCheckGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthApi.HealthCheckGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to an apiHealthCheckGetRequest struct via the builder pattern


### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


