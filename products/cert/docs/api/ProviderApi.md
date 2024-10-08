# \ProviderApi

All URIs are relative to *https://certificate-manager.de-fra.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**ProvidersDelete**](ProviderApi.md#ProvidersDelete) | **Delete** /providers/{providerId} | Delete Provider|
|[**ProvidersFindById**](ProviderApi.md#ProvidersFindById) | **Get** /providers/{providerId} | Retrieve Provider|
|[**ProvidersGet**](ProviderApi.md#ProvidersGet) | **Get** /providers | Retrieve all Provider|
|[**ProvidersPatch**](ProviderApi.md#ProvidersPatch) | **Patch** /providers/{providerId} | Updates Provider|
|[**ProvidersPost**](ProviderApi.md#ProvidersPost) | **Post** /providers | Create Provider|



## ProvidersDelete

```go
var result  = ProvidersDelete(ctx, providerId)
                      .Execute()
```

Delete Provider



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    cert "github.com/ionos-cloud/sdk-go-bundle/products/cert"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    providerId := "74edc770-5cc6-5976-ac99-013ddb4af403" // string | The ID (UUID) of the Provider.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := cert.NewAPIClient(configuration)
    resp, err := apiClient.ProviderApi.ProvidersDelete(context.Background(), providerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderApi.ProvidersDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**providerId** | **string** | The ID (UUID) of the Provider. | |

### Other Parameters

Other parameters are passed through a pointer to an apiProvidersDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ProvidersFindById

```go
var result ProviderRead = ProvidersFindById(ctx, providerId)
                      .Execute()
```

Retrieve Provider



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    cert "github.com/ionos-cloud/sdk-go-bundle/products/cert"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    providerId := "74edc770-5cc6-5976-ac99-013ddb4af403" // string | The ID (UUID) of the Provider.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := cert.NewAPIClient(configuration)
    resource, resp, err := apiClient.ProviderApi.ProvidersFindById(context.Background(), providerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderApi.ProvidersFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ProvidersFindById`: ProviderRead
    fmt.Fprintf(os.Stdout, "Response from `ProviderApi.ProvidersFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**providerId** | **string** | The ID (UUID) of the Provider. | |

### Other Parameters

Other parameters are passed through a pointer to an apiProvidersFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**ProviderRead**](../models/ProviderRead.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ProvidersGet

```go
var result ProviderReadList = ProvidersGet(ctx)
                      .Offset(offset)
                      .Limit(limit)
                      .Execute()
```

Retrieve all Provider



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    cert "github.com/ionos-cloud/sdk-go-bundle/products/cert"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    offset := int32(0) // int32 | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. (optional) (default to 0)
    limit := int32(100) // int32 | The maximum number of elements to return. Use together with offset for pagination. (optional) (default to 100)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := cert.NewAPIClient(configuration)
    resource, resp, err := apiClient.ProviderApi.ProvidersGet(context.Background()).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderApi.ProvidersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ProvidersGet`: ProviderReadList
    fmt.Fprintf(os.Stdout, "Response from `ProviderApi.ProvidersGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiProvidersGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **offset** | **int32** | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return. Use together with offset for pagination. | [default to 100]|

### Return type

[**ProviderReadList**](../models/ProviderReadList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ProvidersPatch

```go
var result ProviderRead = ProvidersPatch(ctx, providerId)
                      .ProviderPatch(providerPatch)
                      .Execute()
```

Updates Provider



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    cert "github.com/ionos-cloud/sdk-go-bundle/products/cert"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    providerId := "74edc770-5cc6-5976-ac99-013ddb4af403" // string | The ID (UUID) of the Provider.
    providerPatch := *openapiclient.NewProviderPatch(*openapiclient.NewPatchName("My name")) // ProviderPatch | patch Provider

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := cert.NewAPIClient(configuration)
    resource, resp, err := apiClient.ProviderApi.ProvidersPatch(context.Background(), providerId).ProviderPatch(providerPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderApi.ProvidersPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ProvidersPatch`: ProviderRead
    fmt.Fprintf(os.Stdout, "Response from `ProviderApi.ProvidersPatch`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**providerId** | **string** | The ID (UUID) of the Provider. | |

### Other Parameters

Other parameters are passed through a pointer to an apiProvidersPatchRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **providerPatch** | [**ProviderPatch**](../models/ProviderPatch.md) | patch Provider | |

### Return type

[**ProviderRead**](../models/ProviderRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## ProvidersPost

```go
var result ProviderRead = ProvidersPost(ctx)
                      .ProviderCreate(providerCreate)
                      .Execute()
```

Create Provider



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    cert "github.com/ionos-cloud/sdk-go-bundle/products/cert"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    providerCreate := *openapiclient.NewProviderCreate(*openapiclient.NewProvider("Let's Encrypt", "user@example.com", "https://acme-v02.api.letsencrypt.org/directory")) // ProviderCreate | Provider to create.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := cert.NewAPIClient(configuration)
    resource, resp, err := apiClient.ProviderApi.ProvidersPost(context.Background()).ProviderCreate(providerCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderApi.ProvidersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ProvidersPost`: ProviderRead
    fmt.Fprintf(os.Stdout, "Response from `ProviderApi.ProvidersPost`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiProvidersPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **providerCreate** | [**ProviderCreate**](../models/ProviderCreate.md) | Provider to create. | |

### Return type

[**ProviderRead**](../models/ProviderRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


