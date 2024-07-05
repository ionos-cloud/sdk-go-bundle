# \UsersApi

All URIs are relative to *https://api.ionos.com/databases/postgresql*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**UsersDelete**](UsersApi.md#UsersDelete) | **Delete** /clusters/{clusterId}/users/{username} | Delete user|
|[**UsersGet**](UsersApi.md#UsersGet) | **Get** /clusters/{clusterId}/users/{username} | Get user|
|[**UsersList**](UsersApi.md#UsersList) | **Get** /clusters/{clusterId}/users | List users|
|[**UsersPatch**](UsersApi.md#UsersPatch) | **Patch** /clusters/{clusterId}/users/{username} | Patch user|
|[**UsersPost**](UsersApi.md#UsersPost) | **Post** /clusters/{clusterId}/users | Create a user|



## UsersDelete

```go
var result  = UsersDelete(ctx, clusterId, username)
                      .Execute()
```

Delete user



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
    username := "benjamin" // string | The authentication username.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := psql.NewAPIClient(configuration)
    resp, err := apiClient.UsersApi.UsersDelete(context.Background(), clusterId, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The unique ID of the cluster. | |
|**username** | **string** | The authentication username. | |

### Other Parameters

Other parameters are passed through a pointer to an apiUsersDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## UsersGet

```go
var result UserResource = UsersGet(ctx, clusterId, username)
                      .Execute()
```

Get user



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
    username := "benjamin" // string | The authentication username.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := psql.NewAPIClient(configuration)
    resource, resp, err := apiClient.UsersApi.UsersGet(context.Background(), clusterId, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `UsersGet`: UserResource
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The unique ID of the cluster. | |
|**username** | **string** | The authentication username. | |

### Other Parameters

Other parameters are passed through a pointer to an apiUsersGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**UserResource**](../models/UserResource.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## UsersList

```go
var result UserList = UsersList(ctx, clusterId)
                      .Limit(limit)
                      .Offset(offset)
                      .System(system)
                      .Execute()
```

List users



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
    limit := int32(100) // int32 | The maximum number of elements to return. Use together with 'offset' for pagination. (optional) (default to 100)
    offset := int32(200) // int32 | The first element to return. Use together with 'limit' for pagination. (optional) (default to 0)
    system := true // bool | If set to 'true' all users, including system users are returned. System users cannot be deleted or updated. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := psql.NewAPIClient(configuration)
    resource, resp, err := apiClient.UsersApi.UsersList(context.Background(), clusterId).Limit(limit).Offset(offset).System(system).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `UsersList`: UserList
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersList`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The unique ID of the cluster. | |

### Other Parameters

Other parameters are passed through a pointer to an apiUsersListRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **limit** | **int32** | The maximum number of elements to return. Use together with &#39;offset&#39; for pagination. | [default to 100]|
| **offset** | **int32** | The first element to return. Use together with &#39;limit&#39; for pagination. | [default to 0]|
| **system** | **bool** | If set to &#39;true&#39; all users, including system users are returned. System users cannot be deleted or updated. | |

### Return type

[**UserList**](../models/UserList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## UsersPatch

```go
var result UserResource = UsersPatch(ctx, clusterId, username)
                      .UsersPatchRequest(usersPatchRequest)
                      .Execute()
```

Patch user



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
    username := "benjamin" // string | The authentication username.
    usersPatchRequest := *openapiclient.NewUsersPatchRequest(*openapiclient.NewPatchUserProperties()) // UsersPatchRequest | Patch containing all properties that should be updated

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := psql.NewAPIClient(configuration)
    resource, resp, err := apiClient.UsersApi.UsersPatch(context.Background(), clusterId, username).UsersPatchRequest(usersPatchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `UsersPatch`: UserResource
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersPatch`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The unique ID of the cluster. | |
|**username** | **string** | The authentication username. | |

### Other Parameters

Other parameters are passed through a pointer to an apiUsersPatchRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **usersPatchRequest** | [**UsersPatchRequest**](../models/UsersPatchRequest.md) | Patch containing all properties that should be updated | |

### Return type

[**UserResource**](../models/UserResource.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## UsersPost

```go
var result UserResource = UsersPost(ctx, clusterId)
                      .User(user)
                      .Execute()
```

Create a user



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
    user := *openapiclient.NewUser(*openapiclient.NewUserProperties("benjamin")) // User | 

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := psql.NewAPIClient(configuration)
    resource, resp, err := apiClient.UsersApi.UsersPost(context.Background(), clusterId).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `UsersPost`: UserResource
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersPost`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The unique ID of the cluster. | |

### Other Parameters

Other parameters are passed through a pointer to an apiUsersPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **user** | [**User**](../models/User.md) |  | |

### Return type

[**UserResource**](../models/UserResource.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


