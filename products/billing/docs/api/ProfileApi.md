# \ProfileApi

All URIs are relative to *https://api.ionos.com/billing*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**ProfilesGet**](ProfileApi.md#ProfilesGet) | **Get** /profile | |



## ProfilesGet

```go
var result ProfilesGet200Response = ProfilesGet(ctx)
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
    resource, resp, err := apiClient.ProfileApi.ProfilesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.ProfilesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ProfilesGet`: ProfilesGet200Response
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.ProfilesGet`: %v\n", resource)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to an apiProfilesGetRequest struct via the builder pattern


### Return type

[**ProfilesGet200Response**](../models/ProfilesGet200Response.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


