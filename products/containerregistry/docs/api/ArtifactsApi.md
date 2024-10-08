# \ArtifactsApi

All URIs are relative to *https://api.ionos.com/containerregistries*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**RegistriesArtifactsGet**](ArtifactsApi.md#RegistriesArtifactsGet) | **Get** /registries/{registryId}/artifacts | Retrieve all Artifacts by Registry|
|[**RegistriesRepositoriesArtifactsFindByDigest**](ArtifactsApi.md#RegistriesRepositoriesArtifactsFindByDigest) | **Get** /registries/{registryId}/repositories/{repositoryName}/artifacts/{digest} | Retrieve Artifact|
|[**RegistriesRepositoriesArtifactsGet**](ArtifactsApi.md#RegistriesRepositoriesArtifactsGet) | **Get** /registries/{registryId}/repositories/{repositoryName}/artifacts | Retrieve all Artifacts by Repository|
|[**RegistriesRepositoriesArtifactsVulnerabilitiesGet**](ArtifactsApi.md#RegistriesRepositoriesArtifactsVulnerabilitiesGet) | **Get** /registries/{registryId}/repositories/{repositoryName}/artifacts/{digest}/vulnerabilities | Retrieve all Vulnerabilities|



## RegistriesArtifactsGet

```go
var result RegistryArtifactsReadList = RegistriesArtifactsGet(ctx, registryId)
                      .Offset(offset)
                      .Limit(limit)
                      .FilterVulnerabilityId(filterVulnerabilityId)
                      .OrderBy(orderBy)
                      .Execute()
```

Retrieve all Artifacts by Registry



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
    filterVulnerabilityId := "filterVulnerabilityId_example" // string | Filter resources by vulnerabilityId. (optional)
    orderBy := "orderBy_example" // string | The field to order the results by. If not provided, the results will be ordered by the default field. (optional) (default to "-pullCount")

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := containerregistry.NewAPIClient(configuration)
    resource, resp, err := apiClient.ArtifactsApi.RegistriesArtifactsGet(context.Background(), registryId).Offset(offset).Limit(limit).FilterVulnerabilityId(filterVulnerabilityId).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsApi.RegistriesArtifactsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `RegistriesArtifactsGet`: RegistryArtifactsReadList
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsApi.RegistriesArtifactsGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**registryId** | **string** | The ID (UUID) of the Registry. | |

### Other Parameters

Other parameters are passed through a pointer to an apiRegistriesArtifactsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **offset** | **int32** | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return. Use together with offset for pagination. | [default to 100]|
| **filterVulnerabilityId** | **string** | Filter resources by vulnerabilityId. | |
| **orderBy** | **string** | The field to order the results by. If not provided, the results will be ordered by the default field. | [default to &quot;-pullCount&quot;]|

### Return type

[**RegistryArtifactsReadList**](../models/RegistryArtifactsReadList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## RegistriesRepositoriesArtifactsFindByDigest

```go
var result ArtifactRead = RegistriesRepositoriesArtifactsFindByDigest(ctx, registryId, repositoryName, digest)
                      .Execute()
```

Retrieve Artifact



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
    digest := "sha256:12345678901234567890123456789012" // string | The Digest of the Artifact.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := containerregistry.NewAPIClient(configuration)
    resource, resp, err := apiClient.ArtifactsApi.RegistriesRepositoriesArtifactsFindByDigest(context.Background(), registryId, repositoryName, digest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsApi.RegistriesRepositoriesArtifactsFindByDigest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `RegistriesRepositoriesArtifactsFindByDigest`: ArtifactRead
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsApi.RegistriesRepositoriesArtifactsFindByDigest`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**registryId** | **string** | The ID (UUID) of the Registry. | |
|**repositoryName** | **string** | The Name of the Repository. | |
|**digest** | **string** | The Digest of the Artifact. | |

### Other Parameters

Other parameters are passed through a pointer to an apiRegistriesRepositoriesArtifactsFindByDigestRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**ArtifactRead**](../models/ArtifactRead.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## RegistriesRepositoriesArtifactsGet

```go
var result ArtifactReadList = RegistriesRepositoriesArtifactsGet(ctx, registryId, repositoryName)
                      .Offset(offset)
                      .Limit(limit)
                      .OrderBy(orderBy)
                      .Execute()
```

Retrieve all Artifacts by Repository



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
    offset := int32(0) // int32 | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. (optional) (default to 0)
    limit := int32(100) // int32 | The maximum number of elements to return. Use together with offset for pagination. (optional) (default to 100)
    orderBy := "orderBy_example" // string | The field to order the results by. If not provided, the results will be ordered by the default field. (optional) (default to "-lastPush")

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := containerregistry.NewAPIClient(configuration)
    resource, resp, err := apiClient.ArtifactsApi.RegistriesRepositoriesArtifactsGet(context.Background(), registryId, repositoryName).Offset(offset).Limit(limit).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsApi.RegistriesRepositoriesArtifactsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `RegistriesRepositoriesArtifactsGet`: ArtifactReadList
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsApi.RegistriesRepositoriesArtifactsGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**registryId** | **string** | The ID (UUID) of the Registry. | |
|**repositoryName** | **string** | The Name of the Repository. | |

### Other Parameters

Other parameters are passed through a pointer to an apiRegistriesRepositoriesArtifactsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **offset** | **int32** | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return. Use together with offset for pagination. | [default to 100]|
| **orderBy** | **string** | The field to order the results by. If not provided, the results will be ordered by the default field. | [default to &quot;-lastPush&quot;]|

### Return type

[**ArtifactReadList**](../models/ArtifactReadList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## RegistriesRepositoriesArtifactsVulnerabilitiesGet

```go
var result ArtifactVulnerabilityReadList = RegistriesRepositoriesArtifactsVulnerabilitiesGet(ctx, registryId, repositoryName, digest)
                      .Offset(offset)
                      .Limit(limit)
                      .FilterSeverity(filterSeverity)
                      .FilterFixable(filterFixable)
                      .OrderBy(orderBy)
                      .Execute()
```

Retrieve all Vulnerabilities



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
    digest := "sha256:12345678901234567890123456789012" // string | The Digest of the Artifact.
    offset := int32(0) // int32 | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. (optional) (default to 0)
    limit := int32(100) // int32 | The maximum number of elements to return. Use together with offset for pagination. (optional) (default to 100)
    filterSeverity := "filterSeverity_example" // string | Filter resources by vulnerability severity. (optional)
    filterFixable := true // bool | Filter resources by fixable (i.e. remediation action is available) (optional)
    orderBy := "orderBy_example" // string | The field to order the results by. If not provided, the results will be ordered by the default field. (optional) (default to "-score")

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := containerregistry.NewAPIClient(configuration)
    resource, resp, err := apiClient.ArtifactsApi.RegistriesRepositoriesArtifactsVulnerabilitiesGet(context.Background(), registryId, repositoryName, digest).Offset(offset).Limit(limit).FilterSeverity(filterSeverity).FilterFixable(filterFixable).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsApi.RegistriesRepositoriesArtifactsVulnerabilitiesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `RegistriesRepositoriesArtifactsVulnerabilitiesGet`: ArtifactVulnerabilityReadList
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsApi.RegistriesRepositoriesArtifactsVulnerabilitiesGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**registryId** | **string** | The ID (UUID) of the Registry. | |
|**repositoryName** | **string** | The Name of the Repository. | |
|**digest** | **string** | The Digest of the Artifact. | |

### Other Parameters

Other parameters are passed through a pointer to an apiRegistriesRepositoriesArtifactsVulnerabilitiesGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **offset** | **int32** | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return. Use together with offset for pagination. | [default to 100]|
| **filterSeverity** | **string** | Filter resources by vulnerability severity. | |
| **filterFixable** | **bool** | Filter resources by fixable (i.e. remediation action is available) | |
| **orderBy** | **string** | The field to order the results by. If not provided, the results will be ordered by the default field. | [default to &quot;-score&quot;]|

### Return type

[**ArtifactVulnerabilityReadList**](../models/ArtifactVulnerabilityReadList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


