# \UsersApi

All URIs are relative to *https://api.ionos.com/databases/mongodb*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**ClustersUsersDelete**](UsersApi.md#ClustersUsersDelete) | **Delete** /clusters/{clusterId}/users/{username} | Delete a MongoDB User by ID|
|[**ClustersUsersFindById**](UsersApi.md#ClustersUsersFindById) | **Get** /clusters/{clusterId}/users/{username} | Get a MongoDB User by ID|
|[**ClustersUsersGet**](UsersApi.md#ClustersUsersGet) | **Get** /clusters/{clusterId}/users | Get all Cluster Users|
|[**ClustersUsersPatch**](UsersApi.md#ClustersUsersPatch) | **Patch** /clusters/{clusterId}/users/{username} | Patch a MongoDB User by ID|
|[**ClustersUsersPost**](UsersApi.md#ClustersUsersPost) | **Post** /clusters/{clusterId}/users | Create MongoDB User|



## ClustersUsersDelete

```go
var result User = ClustersUsersDelete(ctx, clusterId, username)
                      .Execute()
```

Delete a MongoDB User by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    mongo "github.com/ionos-cloud/sdk-go-bundle/products/dbaas/mongo"
    "github.com/ionos-cloud/sdk-go-bundle/common"
)

func main() {
    clusterId := "clusterId_example" // string | The unique ID of the cluster.
    username := "username_example" // string | The authentication username.

    configuration := common.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := mongo.NewAPIClient(configuration)
    resource, resp, err := apiClient.UsersApi.ClustersUsersDelete(context.Background(), clusterId, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ClustersUsersDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersUsersDelete`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.ClustersUsersDelete`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The unique ID of the cluster. | |
|**username** | **string** | The authentication username. | |

### Other Parameters

Other parameters are passed through a pointer to an apiClustersUsersDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**User**](User.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ClustersUsersFindById

```go
var result User = ClustersUsersFindById(ctx, clusterId, username)
                      .Execute()
```

Get a MongoDB User by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    mongo "github.com/ionos-cloud/sdk-go-bundle/products/dbaas/mongo"
    "github.com/ionos-cloud/sdk-go-bundle/common"
)

func main() {
    clusterId := "clusterId_example" // string | The unique ID of the cluster.
    username := "username_example" // string | The authentication username.

    configuration := common.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := mongo.NewAPIClient(configuration)
    resource, resp, err := apiClient.UsersApi.ClustersUsersFindById(context.Background(), clusterId, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ClustersUsersFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersUsersFindById`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.ClustersUsersFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The unique ID of the cluster. | |
|**username** | **string** | The authentication username. | |

### Other Parameters

Other parameters are passed through a pointer to an apiClustersUsersFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**User**](User.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ClustersUsersGet

```go
var result UsersList = ClustersUsersGet(ctx, clusterId)
                      .Execute()
```

Get all Cluster Users



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    mongo "github.com/ionos-cloud/sdk-go-bundle/products/dbaas/mongo"
    "github.com/ionos-cloud/sdk-go-bundle/common"
)

func main() {
    clusterId := "clusterId_example" // string | The unique ID of the cluster.

    configuration := common.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := mongo.NewAPIClient(configuration)
    resource, resp, err := apiClient.UsersApi.ClustersUsersGet(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ClustersUsersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersUsersGet`: UsersList
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.ClustersUsersGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The unique ID of the cluster. | |

### Other Parameters

Other parameters are passed through a pointer to an apiClustersUsersGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**UsersList**](UsersList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ClustersUsersPatch

```go
var result User = ClustersUsersPatch(ctx, clusterId, username)
                      .PatchUserRequest(patchUserRequest)
                      .Execute()
```

Patch a MongoDB User by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    mongo "github.com/ionos-cloud/sdk-go-bundle/products/dbaas/mongo"
    "github.com/ionos-cloud/sdk-go-bundle/common"
)

func main() {
    clusterId := "clusterId_example" // string | The unique ID of the cluster.
    username := "username_example" // string | The authentication username.
    patchUserRequest := *openapiclient.NewPatchUserRequest() // PatchUserRequest | Part of the MongoDB user which should be modified.

    configuration := common.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := mongo.NewAPIClient(configuration)
    resource, resp, err := apiClient.UsersApi.ClustersUsersPatch(context.Background(), clusterId, username).PatchUserRequest(patchUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ClustersUsersPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersUsersPatch`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.ClustersUsersPatch`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The unique ID of the cluster. | |
|**username** | **string** | The authentication username. | |

### Other Parameters

Other parameters are passed through a pointer to an apiClustersUsersPatchRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **patchUserRequest** | [**PatchUserRequest**](PatchUserRequest.md) | Part of the MongoDB user which should be modified. | |

### Return type

[**User**](User.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## ClustersUsersPost

```go
var result User = ClustersUsersPost(ctx, clusterId)
                      .User(user)
                      .Execute()
```

Create MongoDB User



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    mongo "github.com/ionos-cloud/sdk-go-bundle/products/dbaas/mongo"
    "github.com/ionos-cloud/sdk-go-bundle/common"
)

func main() {
    clusterId := "clusterId_example" // string | The unique ID of the cluster.
    user := *openapiclient.NewUser() // User | The user to be created.

    configuration := common.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := mongo.NewAPIClient(configuration)
    resource, resp, err := apiClient.UsersApi.ClustersUsersPost(context.Background(), clusterId).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ClustersUsersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersUsersPost`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.ClustersUsersPost`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The unique ID of the cluster. | |

### Other Parameters

Other parameters are passed through a pointer to an apiClustersUsersPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **user** | [**User**](User.md) | The user to be created. | |

### Return type

[**User**](User.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


