# \NamesApi

All URIs are relative to *https://api.ionos.com/containerregistries*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**NamesCheckUsage**](NamesApi.md#NamesCheckUsage) | **Head** /names/{name} | Get container registry name availability|



## NamesCheckUsage

```go
var result  = NamesCheckUsage(ctx, name)
                      .Execute()
```

Get container registry name availability



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    containerregistry "github.com/ionos-cloud/sdk-go-bundle/products/containerregistry"
    "github.com/ionos-cloud/sdk-go-bundle/common"
)

func main() {
    name := "name_example" // string | The desired registry name

    configuration := common.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := containerregistry.NewAPIClient(configuration)
    resource, resp, err := apiClient.NamesApi.NamesCheckUsage(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamesApi.NamesCheckUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**name** | **string** | The desired registry name | |

### Other Parameters

Other parameters are passed through a pointer to an apiNamesCheckUsageRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


