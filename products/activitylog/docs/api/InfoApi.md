# \InfoApi

All URIs are relative to *https://api.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**GetInfo**](InfoApi.md#GetInfo) | **Get** /activitylog/v1 | Display API information|



## GetInfo

```go
var result Info = GetInfo(ctx)
                      .Execute()
```

Display API information



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    activitylog "github.com/ionos-cloud/sdk-go-bundle/products/activitylog"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := activitylog.NewAPIClient(configuration)
    resource, resp, err := apiClient.InfoApi.GetInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfoApi.GetInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `GetInfo`: Info
    fmt.Fprintf(os.Stdout, "Response from `InfoApi.GetInfo`: %v\n", resource)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to an apiGetInfoRequest struct via the builder pattern


### Return type

[**Info**](../models/Info.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


