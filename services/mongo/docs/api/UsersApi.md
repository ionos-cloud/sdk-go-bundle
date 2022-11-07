# \UsersApi

All URIs are relative to *https://api.ionos.com/databases/mongodb*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**ClustersUsersDelete**](UsersApi.md#ClustersUsersDelete) | **Delete** /clusters/{clusterId}/users/{database}/{username} | Delete a MongoDB User by ID|
|[**ClustersUsersFindById**](UsersApi.md#ClustersUsersFindById) | **Get** /clusters/{clusterId}/users/{database}/{username} | Get a MongoDB User by ID|
|[**ClustersUsersGet**](UsersApi.md#ClustersUsersGet) | **Get** /clusters/{clusterId}/users | Get a Cluster Users|
|[**ClustersUsersPost**](UsersApi.md#ClustersUsersPost) | **Post** /clusters/{clusterId}/users | Create MongoDB User|



## ClustersUsersDelete

```go
var result User = ClustersUsersDelete(ctx, clusterId, database, username)
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

    mongo "github.com/ionos-cloud/sdk-go-bundle/services/mongo"
    "github.com/ionos-cloud/sdk-go-bundle/common"
)

func main() {
    clusterId := "clusterId_example" // string | The unique ID of the cluster.
    database := "database_example" // string | The authentication database.
    username := "username_example" // string | The authentication username.

    configuration := common.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := mongo.NewAPIClient(configuration)
    resource, resp, err := apiClient.UsersApi.ClustersUsersDelete(context.Background(), clusterId, database, username).Execute()
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
|**database** | **string** | The authentication database. | |
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
var result User = ClustersUsersFindById(ctx, clusterId, database, username)
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

    mongo "github.com/ionos-cloud/sdk-go-bundle/services/mongo"
    "github.com/ionos-cloud/sdk-go-bundle/common"
)

func main() {
    clusterId := "clusterId_example" // string | The unique ID of the cluster.
    database := "database_example" // string | The authentication database.
    username := "username_example" // string | The authentication username.

    configuration := common.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := mongo.NewAPIClient(configuration)
    resource, resp, err := apiClient.UsersApi.ClustersUsersFindById(context.Background(), clusterId, database, username).Execute()
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
|**database** | **string** | The authentication database. | |
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

Get a Cluster Users



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    mongo "github.com/ionos-cloud/sdk-go-bundle/services/mongo"
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

    mongo "github.com/ionos-cloud/sdk-go-bundle/services/mongo"
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


