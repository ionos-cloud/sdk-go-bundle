# \RegistriesApi

All URIs are relative to *https://api.ionos.com/containerregistries*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**RegistriesDelete**](RegistriesApi.md#RegistriesDelete) | **Delete** /registries/{registryId} | Delete registry|
|[**RegistriesFindById**](RegistriesApi.md#RegistriesFindById) | **Get** /registries/{registryId} | Get a registry|
|[**RegistriesGet**](RegistriesApi.md#RegistriesGet) | **Get** /registries | List all container registries|
|[**RegistriesPatch**](RegistriesApi.md#RegistriesPatch) | **Patch** /registries/{registryId} | Update the properties of a registry|
|[**RegistriesPost**](RegistriesApi.md#RegistriesPost) | **Post** /registries | Create container registry|
|[**RegistriesPut**](RegistriesApi.md#RegistriesPut) | **Put** /registries/{registryId} | Create or replace a container registry|



## RegistriesDelete

```go
var result  = RegistriesDelete(ctx, registryId)
                      .Execute()
```

Delete registry

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    containerregistry "github.com/ionos-cloud/sdk-go-bundle/services/containerregistry"
    "github.com/ionos-cloud/sdk-go-bundle/common"
)

func main() {
    registryId := TODO // string | The unique ID of the registry

    configuration := common.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := containerregistry.NewAPIClient(configuration)
    resource, resp, err := apiClient.RegistriesApi.RegistriesDelete(context.Background(), registryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistriesApi.RegistriesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**registryId** | [**string**](.md) | The unique ID of the registry | |

### Other Parameters

Other parameters are passed through a pointer to an apiRegistriesDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## RegistriesFindById

```go
var result RegistryResponse = RegistriesFindById(ctx, registryId)
                      .Execute()
```

Get a registry



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    containerregistry "github.com/ionos-cloud/sdk-go-bundle/services/containerregistry"
    "github.com/ionos-cloud/sdk-go-bundle/common"
)

func main() {
    registryId := TODO // string | The unique ID of the registry

    configuration := common.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := containerregistry.NewAPIClient(configuration)
    resource, resp, err := apiClient.RegistriesApi.RegistriesFindById(context.Background(), registryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistriesApi.RegistriesFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `RegistriesFindById`: RegistryResponse
    fmt.Fprintf(os.Stdout, "Response from `RegistriesApi.RegistriesFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**registryId** | [**string**](.md) | The unique ID of the registry | |

### Other Parameters

Other parameters are passed through a pointer to an apiRegistriesFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**RegistryResponse**](RegistryResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## RegistriesGet

```go
var result RegistriesResponse = RegistriesGet(ctx)
                      .FilterName(filterName)
                      .Limit(limit)
                      .NextPageToken(nextPageToken)
                      .Execute()
```

List all container registries



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    containerregistry "github.com/ionos-cloud/sdk-go-bundle/services/containerregistry"
    "github.com/ionos-cloud/sdk-go-bundle/common"
)

func main() {
    filterName := "my-registry" // string | The registry name to search for (optional)
    limit := "limit_example" // string | The maximum number of elements to return (used together with nextPageToken for pagination) (optional) (default to "100")
    nextPageToken := "nextPageToken_example" // string | The next page from the complete list of elements (used together with limit for pagination) (optional)

    configuration := common.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := containerregistry.NewAPIClient(configuration)
    resource, resp, err := apiClient.RegistriesApi.RegistriesGet(context.Background()).FilterName(filterName).Limit(limit).NextPageToken(nextPageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistriesApi.RegistriesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `RegistriesGet`: RegistriesResponse
    fmt.Fprintf(os.Stdout, "Response from `RegistriesApi.RegistriesGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiRegistriesGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **filterName** | **string** | The registry name to search for | |
| **limit** | **string** | The maximum number of elements to return (used together with nextPageToken for pagination) | [default to &quot;100&quot;]|
| **nextPageToken** | **string** | The next page from the complete list of elements (used together with limit for pagination) | |

### Return type

[**RegistriesResponse**](RegistriesResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## RegistriesPatch

```go
var result RegistryResponse = RegistriesPatch(ctx, registryId)
                      .PatchRegistryInput(patchRegistryInput)
                      .Execute()
```

Update the properties of a registry



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    containerregistry "github.com/ionos-cloud/sdk-go-bundle/services/containerregistry"
    "github.com/ionos-cloud/sdk-go-bundle/common"
)

func main() {
    registryId := TODO // string | The unique ID of the registry
    patchRegistryInput := *openapiclient.NewPatchRegistryInput() // PatchRegistryInput | 

    configuration := common.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := containerregistry.NewAPIClient(configuration)
    resource, resp, err := apiClient.RegistriesApi.RegistriesPatch(context.Background(), registryId).PatchRegistryInput(patchRegistryInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistriesApi.RegistriesPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `RegistriesPatch`: RegistryResponse
    fmt.Fprintf(os.Stdout, "Response from `RegistriesApi.RegistriesPatch`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**registryId** | [**string**](.md) | The unique ID of the registry | |

### Other Parameters

Other parameters are passed through a pointer to an apiRegistriesPatchRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **patchRegistryInput** | [**PatchRegistryInput**](PatchRegistryInput.md) |  | |

### Return type

[**RegistryResponse**](RegistryResponse.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## RegistriesPost

```go
var result PostRegistryOutput = RegistriesPost(ctx)
                      .PostRegistryInput(postRegistryInput)
                      .Execute()
```

Create container registry



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    containerregistry "github.com/ionos-cloud/sdk-go-bundle/services/containerregistry"
    "github.com/ionos-cloud/sdk-go-bundle/common"
)

func main() {
    postRegistryInput := *openapiclient.NewPostRegistryInput(*openapiclient.NewPostRegistryProperties("de/txl", "my-registry")) // PostRegistryInput | 

    configuration := common.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := containerregistry.NewAPIClient(configuration)
    resource, resp, err := apiClient.RegistriesApi.RegistriesPost(context.Background()).PostRegistryInput(postRegistryInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistriesApi.RegistriesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `RegistriesPost`: PostRegistryOutput
    fmt.Fprintf(os.Stdout, "Response from `RegistriesApi.RegistriesPost`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiRegistriesPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **postRegistryInput** | [**PostRegistryInput**](PostRegistryInput.md) |  | |

### Return type

[**PostRegistryOutput**](PostRegistryOutput.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## RegistriesPut

```go
var result PutRegistryOutput = RegistriesPut(ctx, registryId)
                      .PutRegistryInput(putRegistryInput)
                      .Execute()
```

Create or replace a container registry



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    containerregistry "github.com/ionos-cloud/sdk-go-bundle/services/containerregistry"
    "github.com/ionos-cloud/sdk-go-bundle/common"
)

func main() {
    registryId := TODO // string | The unique ID of the registry
    putRegistryInput := *openapiclient.NewPutRegistryInput(*openapiclient.NewPostRegistryProperties("de/txl", "my-registry")) // PutRegistryInput | 

    configuration := common.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := containerregistry.NewAPIClient(configuration)
    resource, resp, err := apiClient.RegistriesApi.RegistriesPut(context.Background(), registryId).PutRegistryInput(putRegistryInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistriesApi.RegistriesPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `RegistriesPut`: PutRegistryOutput
    fmt.Fprintf(os.Stdout, "Response from `RegistriesApi.RegistriesPut`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**registryId** | [**string**](.md) | The unique ID of the registry | |

### Other Parameters

Other parameters are passed through a pointer to an apiRegistriesPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **putRegistryInput** | [**PutRegistryInput**](PutRegistryInput.md) |  | |

### Return type

[**PutRegistryOutput**](PutRegistryOutput.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


