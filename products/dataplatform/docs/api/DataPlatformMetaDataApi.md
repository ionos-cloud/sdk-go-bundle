# \DataPlatformMetaDataApi

All URIs are relative to *https://api.ionos.com/dataplatform*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**VersionsGet**](DataPlatformMetaDataApi.md#VersionsGet) | **Get** /versions | Managed Stackable Data Platform API Versions|



## VersionsGet

```go
var result []string = VersionsGet(ctx)
                      .Execute()
```

Managed Stackable Data Platform API Versions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    dataplatform "github.com/ionos-cloud/sdk-go-bundle/products/dataplatform"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := dataplatform.NewAPIClient(configuration)
    resource, resp, err := apiClient.DataPlatformMetaDataApi.VersionsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataPlatformMetaDataApi.VersionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `VersionsGet`: []string
    fmt.Fprintf(os.Stdout, "Response from `DataPlatformMetaDataApi.VersionsGet`: %v\n", resource)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to an apiVersionsGetRequest struct via the builder pattern


### Return type

**[]string**

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


