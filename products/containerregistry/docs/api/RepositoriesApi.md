# \RepositoriesApi

All URIs are relative to *https://api.ionos.com/containerregistries*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**RegistriesRepositoriesDelete**](RepositoriesApi.md#RegistriesRepositoriesDelete) | **Delete** /registries/{registryId}/repositories/{repositoryName} | Delete repository|
|[**RegistriesRepositoriesFindByName**](RepositoriesApi.md#RegistriesRepositoriesFindByName) | **Get** /registries/{registryId}/repositories/{repositoryName} | Retrieve Repository|
|[**RegistriesRepositoriesGet**](RepositoriesApi.md#RegistriesRepositoriesGet) | **Get** /registries/{registryId}/repositories | Retrieve all Repositories|



## RegistriesRepositoriesDelete

```go
var result  = RegistriesRepositoriesDelete(ctx, registryId, repositoryName)
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
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    registryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the registry
    repositoryName := "my-service" // string | The name of the repository

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := containerregistry.NewAPIClient(configuration)
    resp, err := apiClient.RepositoriesApi.RegistriesRepositoriesDelete(context.Background(), registryId, repositoryName).Execute()
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
|**registryId** | **string** | The unique ID of the registry | |
|**repositoryName** | **string** | The name of the repository | |

### Other Parameters

Other parameters are passed through a pointer to an apiRegistriesRepositoriesDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined



## RegistriesRepositoriesFindByName

```go
var result RepositoryRead = RegistriesRepositoriesFindByName(ctx, registryId, repositoryName)
                      .Execute()
```

Retrieve Repository



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    containerregistry "github.com/ionos-cloud/sdk-go-bundle/products/containerregistry"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    registryId := "1e41a73c-59d0-5507-86dd-fa2fc2501cfd" // string | The ID (UUID) of the Registry.
    repositoryName := "my-service" // string | The Name of the Repository.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := containerregistry.NewAPIClient(configuration)
    resource, resp, err := apiClient.RepositoriesApi.RegistriesRepositoriesFindByName(context.Background(), registryId, repositoryName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.RegistriesRepositoriesFindByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `RegistriesRepositoriesFindByName`: RepositoryRead
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.RegistriesRepositoriesFindByName`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**registryId** | **string** | The ID (UUID) of the Registry. | |
|**repositoryName** | **string** | The Name of the Repository. | |

### Other Parameters

Other parameters are passed through a pointer to an apiRegistriesRepositoriesFindByNameRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**RepositoryRead**](../models/RepositoryRead.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## RegistriesRepositoriesGet

```go
var result RepositoryReadList = RegistriesRepositoriesGet(ctx, registryId)
                      .Offset(offset)
                      .Limit(limit)
                      .FilterName(filterName)
                      .FilterVulnerabilitySeverity(filterVulnerabilitySeverity)
                      .OrderBy(orderBy)
                      .Execute()
```

Retrieve all Repositories



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    containerregistry "github.com/ionos-cloud/sdk-go-bundle/products/containerregistry"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    registryId := "1e41a73c-59d0-5507-86dd-fa2fc2501cfd" // string | The ID (UUID) of the Registry.
    offset := int32(0) // int32 | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. (optional) (default to 0)
    limit := int32(100) // int32 | The maximum number of elements to return. Use together with offset for pagination. (optional) (default to 100)
    filterName := "filterName_example" // string | Filter resources by name. (optional)
    filterVulnerabilitySeverity := "filterVulnerabilitySeverity_example" // string | Filter resources by vulnerability severity. (optional)
    orderBy := "orderBy_example" // string | The field to order the results by. If not provided, the results will be ordered by the default field. (optional) (default to "-lastPush")

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := containerregistry.NewAPIClient(configuration)
    resource, resp, err := apiClient.RepositoriesApi.RegistriesRepositoriesGet(context.Background(), registryId).Offset(offset).Limit(limit).FilterName(filterName).FilterVulnerabilitySeverity(filterVulnerabilitySeverity).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.RegistriesRepositoriesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `RegistriesRepositoriesGet`: RepositoryReadList
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.RegistriesRepositoriesGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**registryId** | **string** | The ID (UUID) of the Registry. | |

### Other Parameters

Other parameters are passed through a pointer to an apiRegistriesRepositoriesGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **offset** | **int32** | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return. Use together with offset for pagination. | [default to 100]|
| **filterName** | **string** | Filter resources by name. | |
| **filterVulnerabilitySeverity** | **string** | Filter resources by vulnerability severity. | |
| **orderBy** | **string** | The field to order the results by. If not provided, the results will be ordered by the default field. | [default to &quot;-lastPush&quot;]|

### Return type

[**RepositoryReadList**](../models/RepositoryReadList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


