# \ClustersApi

All URIs are relative to *https://api.ionos.com/databases/postgresql*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**ClusterPostgresVersionsGet**](ClustersApi.md#ClusterPostgresVersionsGet) | **Get** /clusters/{clusterId}/postgresversions | List PostgreSQL versions|
|[**ClustersDelete**](ClustersApi.md#ClustersDelete) | **Delete** /clusters/{clusterId} | Delete a cluster|
|[**ClustersFindById**](ClustersApi.md#ClustersFindById) | **Get** /clusters/{clusterId} | Fetch a cluster|
|[**ClustersGet**](ClustersApi.md#ClustersGet) | **Get** /clusters | List clusters|
|[**ClustersPatch**](ClustersApi.md#ClustersPatch) | **Patch** /clusters/{clusterId} | Patch a cluster|
|[**ClustersPost**](ClustersApi.md#ClustersPost) | **Post** /clusters | Create a cluster|
|[**PostgresVersionsGet**](ClustersApi.md#PostgresVersionsGet) | **Get** /clusters/postgresversions | List PostgreSQL versions|



## ClusterPostgresVersionsGet

```go
var result PostgresVersionList = ClusterPostgresVersionsGet(ctx, clusterId)
                      .Execute()
```

List PostgreSQL versions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    psql "github.com/ionos-cloud/sdk-go-bundle/products/dbaas/psql"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    clusterId := "498ae72f-411f-11eb-9d07-046c59cc737e" // string | The unique ID of the cluster.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := psql.NewAPIClient(configuration)
    resource, resp, err := apiClient.ClustersApi.ClusterPostgresVersionsGet(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClusterPostgresVersionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClusterPostgresVersionsGet`: PostgresVersionList
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ClusterPostgresVersionsGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The unique ID of the cluster. | |

### Other Parameters

Other parameters are passed through a pointer to an apiClusterPostgresVersionsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**PostgresVersionList**](../models/PostgresVersionList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ClustersDelete

```go
var result ClusterResponse = ClustersDelete(ctx, clusterId)
                      .Execute()
```

Delete a cluster



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    psql "github.com/ionos-cloud/sdk-go-bundle/products/dbaas/psql"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    clusterId := "498ae72f-411f-11eb-9d07-046c59cc737e" // string | The unique ID of the cluster.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := psql.NewAPIClient(configuration)
    resp, err := apiClient.ClustersApi.ClustersDelete(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersDelete`: ClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ClustersDelete`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The unique ID of the cluster. | |

### Other Parameters

Other parameters are passed through a pointer to an apiClustersDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**ClusterResponse**](../models/ClusterResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ClustersFindById

```go
var result ClusterResponse = ClustersFindById(ctx, clusterId)
                      .Execute()
```

Fetch a cluster



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    psql "github.com/ionos-cloud/sdk-go-bundle/products/dbaas/psql"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    clusterId := "498ae72f-411f-11eb-9d07-046c59cc737e" // string | The unique ID of the cluster.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := psql.NewAPIClient(configuration)
    resource, resp, err := apiClient.ClustersApi.ClustersFindById(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersFindById`: ClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ClustersFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The unique ID of the cluster. | |

### Other Parameters

Other parameters are passed through a pointer to an apiClustersFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**ClusterResponse**](../models/ClusterResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ClustersGet

```go
var result ClusterList = ClustersGet(ctx)
                      .Limit(limit)
                      .Offset(offset)
                      .FilterName(filterName)
                      .Execute()
```

List clusters



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    psql "github.com/ionos-cloud/sdk-go-bundle/products/dbaas/psql"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    limit := int32(100) // int32 | The maximum number of elements to return. Use together with 'offset' for pagination. (optional) (default to 100)
    offset := int32(200) // int32 | The first element to return. Use together with 'limit' for pagination. (optional) (default to 0)
    filterName := "filterName_example" // string | Response filter to list only the PostgreSQL clusters that contain the specified name. The value is case insensitive and matched on the 'displayName' field.  (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := psql.NewAPIClient(configuration)
    resource, resp, err := apiClient.ClustersApi.ClustersGet(context.Background()).Limit(limit).Offset(offset).FilterName(filterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersGet`: ClusterList
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ClustersGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiClustersGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **limit** | **int32** | The maximum number of elements to return. Use together with &#39;offset&#39; for pagination. | [default to 100]|
| **offset** | **int32** | The first element to return. Use together with &#39;limit&#39; for pagination. | [default to 0]|
| **filterName** | **string** | Response filter to list only the PostgreSQL clusters that contain the specified name. The value is case insensitive and matched on the &#39;displayName&#39; field.  | |

### Return type

[**ClusterList**](../models/ClusterList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ClustersPatch

```go
var result ClusterResponse = ClustersPatch(ctx, clusterId)
                      .PatchClusterRequest(patchClusterRequest)
                      .Execute()
```

Patch a cluster



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    psql "github.com/ionos-cloud/sdk-go-bundle/products/dbaas/psql"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    clusterId := "498ae72f-411f-11eb-9d07-046c59cc737e" // string | The unique ID of the cluster.
    patchClusterRequest := *openapiclient.NewPatchClusterRequest() // PatchClusterRequest | Part of the cluster which should be modified.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := psql.NewAPIClient(configuration)
    resource, resp, err := apiClient.ClustersApi.ClustersPatch(context.Background(), clusterId).PatchClusterRequest(patchClusterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersPatch`: ClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ClustersPatch`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The unique ID of the cluster. | |

### Other Parameters

Other parameters are passed through a pointer to an apiClustersPatchRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **patchClusterRequest** | [**PatchClusterRequest**](../models/PatchClusterRequest.md) | Part of the cluster which should be modified. | |

### Return type

[**ClusterResponse**](../models/ClusterResponse.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## ClustersPost

```go
var result ClusterResponse = ClustersPost(ctx)
                      .CreateClusterRequest(createClusterRequest)
                      .Execute()
```

Create a cluster



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    psql "github.com/ionos-cloud/sdk-go-bundle/products/dbaas/psql"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    createClusterRequest := *openapiclient.NewCreateClusterRequest() // CreateClusterRequest | The cluster to be created.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := psql.NewAPIClient(configuration)
    resource, resp, err := apiClient.ClustersApi.ClustersPost(context.Background()).CreateClusterRequest(createClusterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersPost`: ClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ClustersPost`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiClustersPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **createClusterRequest** | [**CreateClusterRequest**](../models/CreateClusterRequest.md) | The cluster to be created. | |

### Return type

[**ClusterResponse**](../models/ClusterResponse.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## PostgresVersionsGet

```go
var result PostgresVersionList = PostgresVersionsGet(ctx)
                      .Execute()
```

List PostgreSQL versions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    psql "github.com/ionos-cloud/sdk-go-bundle/products/dbaas/psql"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := psql.NewAPIClient(configuration)
    resource, resp, err := apiClient.ClustersApi.PostgresVersionsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.PostgresVersionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `PostgresVersionsGet`: PostgresVersionList
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.PostgresVersionsGet`: %v\n", resource)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to an apiPostgresVersionsGetRequest struct via the builder pattern


### Return type

[**PostgresVersionList**](../models/PostgresVersionList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


