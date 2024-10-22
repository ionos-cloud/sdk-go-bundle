# \DistributionsApi

All URIs are relative to *https://cdn.de-fra.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**DistributionsDelete**](DistributionsApi.md#DistributionsDelete) | **Delete** /distributions/{distributionId} | Delete Distribution|
|[**DistributionsFindById**](DistributionsApi.md#DistributionsFindById) | **Get** /distributions/{distributionId} | Retrieve Distribution|
|[**DistributionsGet**](DistributionsApi.md#DistributionsGet) | **Get** /distributions | Retrieve all Distributions|
|[**DistributionsPost**](DistributionsApi.md#DistributionsPost) | **Post** /distributions | Create Distribution|
|[**DistributionsPut**](DistributionsApi.md#DistributionsPut) | **Put** /distributions/{distributionId} | Ensure Distribution|



## DistributionsDelete

```go
var result  = DistributionsDelete(ctx, distributionId)
                      .Execute()
```

Delete Distribution



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    cdn "github.com/ionos-cloud/sdk-go-bundle/products/cdn"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    distributionId := "9ba15778-16c4-543c-8775-e52acf4853f5" // string | The ID (UUID) of the Distribution.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := cdn.NewAPIClient(configuration)
    resp, err := apiClient.DistributionsApi.DistributionsDelete(context.Background(), distributionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributionsApi.DistributionsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**distributionId** | **string** | The ID (UUID) of the Distribution. | |

### Other Parameters

Other parameters are passed through a pointer to an apiDistributionsDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DistributionsFindById

```go
var result Distribution = DistributionsFindById(ctx, distributionId)
                      .Execute()
```

Retrieve Distribution



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    cdn "github.com/ionos-cloud/sdk-go-bundle/products/cdn"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    distributionId := "9ba15778-16c4-543c-8775-e52acf4853f5" // string | The ID (UUID) of the Distribution.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := cdn.NewAPIClient(configuration)
    resource, resp, err := apiClient.DistributionsApi.DistributionsFindById(context.Background(), distributionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributionsApi.DistributionsFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `DistributionsFindById`: Distribution
    fmt.Fprintf(os.Stdout, "Response from `DistributionsApi.DistributionsFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**distributionId** | **string** | The ID (UUID) of the Distribution. | |

### Other Parameters

Other parameters are passed through a pointer to an apiDistributionsFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**Distribution**](../models/Distribution.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DistributionsGet

```go
var result Distributions = DistributionsGet(ctx)
                      .Offset(offset)
                      .Limit(limit)
                      .FilterDomain(filterDomain)
                      .FilterState(filterState)
                      .Execute()
```

Retrieve all Distributions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    cdn "github.com/ionos-cloud/sdk-go-bundle/products/cdn"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    offset := int32(0) // int32 | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. (optional) (default to 0)
    limit := int32(100) // int32 | The maximum number of elements to return. Use together with offset for pagination. (optional) (default to 100)
    filterDomain := "example.com" // string | filters resources by domain. (optional)
    filterState := "AVAILABLE" // string | filters resources by state. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := cdn.NewAPIClient(configuration)
    resource, resp, err := apiClient.DistributionsApi.DistributionsGet(context.Background()).Offset(offset).Limit(limit).FilterDomain(filterDomain).FilterState(filterState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributionsApi.DistributionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `DistributionsGet`: Distributions
    fmt.Fprintf(os.Stdout, "Response from `DistributionsApi.DistributionsGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiDistributionsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **offset** | **int32** | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return. Use together with offset for pagination. | [default to 100]|
| **filterDomain** | **string** | filters resources by domain. | |
| **filterState** | **string** | filters resources by state. | |

### Return type

[**Distributions**](../models/Distributions.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DistributionsPost

```go
var result Distribution = DistributionsPost(ctx)
                      .DistributionCreate(distributionCreate)
                      .Execute()
```

Create Distribution



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    cdn "github.com/ionos-cloud/sdk-go-bundle/products/cdn"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    distributionCreate := *openapiclient.NewDistributionCreate(*openapiclient.NewDistributionProperties("example.com", []openapiclient.RoutingRule{*openapiclient.NewRoutingRule("http/https", "/api", *openapiclient.NewUpstream("server.example.com", true, true, "RateLimitClass_example", "SniMode_example"))})) // DistributionCreate | Distribution to create.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := cdn.NewAPIClient(configuration)
    resource, resp, err := apiClient.DistributionsApi.DistributionsPost(context.Background()).DistributionCreate(distributionCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributionsApi.DistributionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `DistributionsPost`: Distribution
    fmt.Fprintf(os.Stdout, "Response from `DistributionsApi.DistributionsPost`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiDistributionsPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **distributionCreate** | [**DistributionCreate**](../models/DistributionCreate.md) | Distribution to create. | |

### Return type

[**Distribution**](../models/Distribution.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DistributionsPut

```go
var result Distribution = DistributionsPut(ctx, distributionId)
                      .DistributionUpdate(distributionUpdate)
                      .Execute()
```

Ensure Distribution



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    cdn "github.com/ionos-cloud/sdk-go-bundle/products/cdn"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    distributionId := "9ba15778-16c4-543c-8775-e52acf4853f5" // string | The ID (UUID) of the Distribution.
    distributionUpdate := *openapiclient.NewDistributionUpdate("9ba15778-16c4-543c-8775-e52acf4853f5", *openapiclient.NewDistributionProperties("example.com", []openapiclient.RoutingRule{*openapiclient.NewRoutingRule("http/https", "/api", *openapiclient.NewUpstream("server.example.com", true, true, "RateLimitClass_example", "SniMode_example"))})) // DistributionUpdate | update Distribution

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := cdn.NewAPIClient(configuration)
    resource, resp, err := apiClient.DistributionsApi.DistributionsPut(context.Background(), distributionId).DistributionUpdate(distributionUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributionsApi.DistributionsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `DistributionsPut`: Distribution
    fmt.Fprintf(os.Stdout, "Response from `DistributionsApi.DistributionsPut`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**distributionId** | **string** | The ID (UUID) of the Distribution. | |

### Other Parameters

Other parameters are passed through a pointer to an apiDistributionsPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **distributionUpdate** | [**DistributionUpdate**](../models/DistributionUpdate.md) | update Distribution | |

### Return type

[**Distribution**](../models/Distribution.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


