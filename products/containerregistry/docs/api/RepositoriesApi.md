# \RepositoriesApi

All URIs are relative to *https://api.ionos.com/containerregistries*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**RegistriesRepositoriesDelete**](RepositoriesApi.md#RegistriesRepositoriesDelete) | **Delete** /registries/{registryId}/repositories/{name} | Delete repository|



## RegistriesRepositoriesDelete

```go
var result  = RegistriesRepositoriesDelete(ctx, registryId, name)
                      .Execute()
```

Delete repository



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
    registryId := TODO // string | The unique ID of the registry
    name := "ubuntu" // string | The name of the repository

    configuration := common.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := containerregistry.NewAPIClient(configuration)
    resource, resp, err := apiClient.RepositoriesApi.RegistriesRepositoriesDelete(context.Background(), registryId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.RegistriesRepositoriesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**registryId** | [**string**](../models/.md) | The unique ID of the registry | |
|**name** | **string** | The name of the repository | |

### Other Parameters

Other parameters are passed through a pointer to an apiRegistriesRepositoriesDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


