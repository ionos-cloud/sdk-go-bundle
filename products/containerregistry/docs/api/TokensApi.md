# \TokensApi

All URIs are relative to *https://api.ionos.com/containerregistries*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**RegistriesTokensDelete**](TokensApi.md#RegistriesTokensDelete) | **Delete** /registries/{registryId}/tokens/{tokenId} | Delete token|
|[**RegistriesTokensFindById**](TokensApi.md#RegistriesTokensFindById) | **Get** /registries/{registryId}/tokens/{tokenId} | Get token information|
|[**RegistriesTokensGet**](TokensApi.md#RegistriesTokensGet) | **Get** /registries/{registryId}/tokens | List all tokens for the container registry|
|[**RegistriesTokensPatch**](TokensApi.md#RegistriesTokensPatch) | **Patch** /registries/{registryId}/tokens/{tokenId} | Update token|
|[**RegistriesTokensPost**](TokensApi.md#RegistriesTokensPost) | **Post** /registries/{registryId}/tokens | Create token|
|[**RegistriesTokensPut**](TokensApi.md#RegistriesTokensPut) | **Put** /registries/{registryId}/tokens/{tokenId} | Create or replace token|



## RegistriesTokensDelete

```go
var result  = RegistriesTokensDelete(ctx, registryId, tokenId)
                      .Execute()
```

Delete token

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
    tokenId := TODO // string | The unique ID of the token

    configuration := common.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := containerregistry.NewAPIClient(configuration)
    resource, resp, err := apiClient.TokensApi.RegistriesTokensDelete(context.Background(), registryId, tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.RegistriesTokensDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**registryId** | [**string**](.md) | The unique ID of the registry | |
|**tokenId** | [**string**](.md) | The unique ID of the token | |

### Other Parameters

Other parameters are passed through a pointer to an apiRegistriesTokensDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## RegistriesTokensFindById

```go
var result TokenResponse = RegistriesTokensFindById(ctx, registryId, tokenId)
                      .Execute()
```

Get token information



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
    tokenId := TODO // string | The unique ID of the token

    configuration := common.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := containerregistry.NewAPIClient(configuration)
    resource, resp, err := apiClient.TokensApi.RegistriesTokensFindById(context.Background(), registryId, tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.RegistriesTokensFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `RegistriesTokensFindById`: TokenResponse
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.RegistriesTokensFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**registryId** | [**string**](.md) | The unique ID of the registry | |
|**tokenId** | [**string**](.md) | The unique ID of the token | |

### Other Parameters

Other parameters are passed through a pointer to an apiRegistriesTokensFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**TokenResponse**](TokenResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## RegistriesTokensGet

```go
var result TokensResponse = RegistriesTokensGet(ctx, registryId)
                      .Offset(offset)
                      .Limit(limit)
                      .Execute()
```

List all tokens for the container registry

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
    offset := "offset_example" // string | The first element (from the complete list of the elements) to include in the response (used together with limit for pagination) (optional) (default to "0")
    limit := "limit_example" // string | The maximum number of elements to return (used together with offset for pagination) (optional) (default to "100")

    configuration := common.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := containerregistry.NewAPIClient(configuration)
    resource, resp, err := apiClient.TokensApi.RegistriesTokensGet(context.Background(), registryId).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.RegistriesTokensGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `RegistriesTokensGet`: TokensResponse
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.RegistriesTokensGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**registryId** | [**string**](.md) | The unique ID of the registry | |

### Other Parameters

Other parameters are passed through a pointer to an apiRegistriesTokensGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **offset** | **string** | The first element (from the complete list of the elements) to include in the response (used together with limit for pagination) | [default to &quot;0&quot;]|
| **limit** | **string** | The maximum number of elements to return (used together with offset for pagination) | [default to &quot;100&quot;]|

### Return type

[**TokensResponse**](TokensResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## RegistriesTokensPatch

```go
var result TokenResponse = RegistriesTokensPatch(ctx, registryId, tokenId)
                      .PatchTokenInput(patchTokenInput)
                      .Execute()
```

Update token



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
    tokenId := TODO // string | The unique ID of the token
    patchTokenInput := *openapiclient.NewPatchTokenInput() // PatchTokenInput | 

    configuration := common.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := containerregistry.NewAPIClient(configuration)
    resource, resp, err := apiClient.TokensApi.RegistriesTokensPatch(context.Background(), registryId, tokenId).PatchTokenInput(patchTokenInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.RegistriesTokensPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `RegistriesTokensPatch`: TokenResponse
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.RegistriesTokensPatch`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**registryId** | [**string**](.md) | The unique ID of the registry | |
|**tokenId** | [**string**](.md) | The unique ID of the token | |

### Other Parameters

Other parameters are passed through a pointer to an apiRegistriesTokensPatchRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **patchTokenInput** | [**PatchTokenInput**](PatchTokenInput.md) |  | |

### Return type

[**TokenResponse**](TokenResponse.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## RegistriesTokensPost

```go
var result PostTokenOutput = RegistriesTokensPost(ctx, registryId)
                      .PostTokenInput(postTokenInput)
                      .Execute()
```

Create token



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
    postTokenInput := *openapiclient.NewPostTokenInput(*openapiclient.NewPostTokenProperties("push-token")) // PostTokenInput | 

    configuration := common.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := containerregistry.NewAPIClient(configuration)
    resource, resp, err := apiClient.TokensApi.RegistriesTokensPost(context.Background(), registryId).PostTokenInput(postTokenInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.RegistriesTokensPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `RegistriesTokensPost`: PostTokenOutput
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.RegistriesTokensPost`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**registryId** | [**string**](.md) | The unique ID of the registry | |

### Other Parameters

Other parameters are passed through a pointer to an apiRegistriesTokensPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **postTokenInput** | [**PostTokenInput**](PostTokenInput.md) |  | |

### Return type

[**PostTokenOutput**](PostTokenOutput.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## RegistriesTokensPut

```go
var result PutTokenOutput = RegistriesTokensPut(ctx, registryId, tokenId)
                      .PutTokenInput(putTokenInput)
                      .Execute()
```

Create or replace token



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
    tokenId := "tokenId_example" // string | The unique ID of the token
    putTokenInput := *openapiclient.NewPutTokenInput(*openapiclient.NewPostTokenProperties("push-token")) // PutTokenInput | 

    configuration := common.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := containerregistry.NewAPIClient(configuration)
    resource, resp, err := apiClient.TokensApi.RegistriesTokensPut(context.Background(), registryId, tokenId).PutTokenInput(putTokenInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.RegistriesTokensPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `RegistriesTokensPut`: PutTokenOutput
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.RegistriesTokensPut`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**registryId** | [**string**](.md) | The unique ID of the registry | |
|**tokenId** | **string** | The unique ID of the token | |

### Other Parameters

Other parameters are passed through a pointer to an apiRegistriesTokensPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **putTokenInput** | [**PutTokenInput**](PutTokenInput.md) |  | |

### Return type

[**PutTokenOutput**](PutTokenOutput.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


