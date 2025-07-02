# \BucketsApi

All URIs are relative to *https://s3.eu-central-1.ionoscloud.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**CreateBucket**](BucketsApi.md#CreateBucket) | **Put** /{Bucket} | CreateBucket|
|[**DeleteBucket**](BucketsApi.md#DeleteBucket) | **Delete** /{Bucket} | DeleteBucket|
|[**GetBucketLocation**](BucketsApi.md#GetBucketLocation) | **Get** /{Bucket}?location | GetBucketLocation|
|[**HeadBucket**](BucketsApi.md#HeadBucket) | **Head** /{Bucket} | HeadBucket|
|[**ListBuckets**](BucketsApi.md#ListBuckets) | **Get** / | ListBuckets|



## CreateBucket

```go
var result  = CreateBucket(ctx, bucket)
                      .CreateBucketRequest(createBucketRequest)
                      .XAmzBucketObjectLockEnabled(xAmzBucketObjectLockEnabled)
                      .Execute()
```

CreateBucket



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    userobjectstorage "github.com/ionos-cloud/sdk-go-bundle/userobjectstorage"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    bucket := "bucket_example" // string | 
    createBucketRequest := *openapiclient.NewCreateBucketRequest() // CreateBucketRequest | 
    xAmzBucketObjectLockEnabled := true // bool | Specifies whether you want Object Lock enabled for the new bucket. After bucket creation, you must apply the [Object Lock configuration](#tag/Object-Lock/operation/PutObjectLockConfiguration). (optional) (default to false)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := userobjectstorage.NewAPIClient(configuration)
    resource, resp, err := apiClient.BucketsApi.CreateBucket(context.Background(), bucket).CreateBucketRequest(createBucketRequest).XAmzBucketObjectLockEnabled(xAmzBucketObjectLockEnabled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketsApi.CreateBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**bucket** | **string** |  | |

### Other Parameters

Other parameters are passed through a pointer to an apiCreateBucketRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **createBucketRequest** | [**CreateBucketRequest**](../models/CreateBucketRequest.md) |  | |
| **xAmzBucketObjectLockEnabled** | **bool** | Specifies whether you want Object Lock enabled for the new bucket. After bucket creation, you must apply the [Object Lock configuration](#tag/Object-Lock/operation/PutObjectLockConfiguration). | [default to false]|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: application/xml
- **Accept**: application/xml


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"BucketsApiService.CreateBucket"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), shared.ContextOperationServerIndices, map[string]int{
    "BucketsApiService.CreateBucket": 2,
})
ctx = context.WithValue(context.Background(), shared.ContextOperationServerVariables, map[string]map[string]string{
    "BucketsApiService.CreateBucket": {
    "port": "8443",
},
})
```


## DeleteBucket

```go
var result  = DeleteBucket(ctx, bucket)
                      .Execute()
```

DeleteBucket



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    userobjectstorage "github.com/ionos-cloud/sdk-go-bundle/userobjectstorage"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    bucket := "bucket_example" // string | 

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := userobjectstorage.NewAPIClient(configuration)
    resp, err := apiClient.BucketsApi.DeleteBucket(context.Background(), bucket).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketsApi.DeleteBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**bucket** | **string** |  | |

### Other Parameters

Other parameters are passed through a pointer to an apiDeleteBucketRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"BucketsApiService.DeleteBucket"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), shared.ContextOperationServerIndices, map[string]int{
    "BucketsApiService.DeleteBucket": 2,
})
ctx = context.WithValue(context.Background(), shared.ContextOperationServerVariables, map[string]map[string]string{
    "BucketsApiService.DeleteBucket": {
    "port": "8443",
},
})
```


## GetBucketLocation

```go
var result GetBucketLocation200Response = GetBucketLocation(ctx, bucket)
                      .Location(location)
                      .Execute()
```

GetBucketLocation



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    userobjectstorage "github.com/ionos-cloud/sdk-go-bundle/userobjectstorage"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    bucket := "bucket_example" // string | 
    location := true // bool | 

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := userobjectstorage.NewAPIClient(configuration)
    resource, resp, err := apiClient.BucketsApi.GetBucketLocation(context.Background(), bucket).Location(location).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketsApi.GetBucketLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `GetBucketLocation`: GetBucketLocation200Response
    fmt.Fprintf(os.Stdout, "Response from `BucketsApi.GetBucketLocation`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**bucket** | **string** |  | |

### Other Parameters

Other parameters are passed through a pointer to an apiGetBucketLocationRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **location** | **bool** |  | |

### Return type

[**GetBucketLocation200Response**](../models/GetBucketLocation200Response.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"BucketsApiService.GetBucketLocation"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), shared.ContextOperationServerIndices, map[string]int{
    "BucketsApiService.GetBucketLocation": 2,
})
ctx = context.WithValue(context.Background(), shared.ContextOperationServerVariables, map[string]map[string]string{
    "BucketsApiService.GetBucketLocation": {
    "port": "8443",
},
})
```


## HeadBucket

```go
var result  = HeadBucket(ctx, bucket)
                      .Execute()
```

HeadBucket



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    userobjectstorage "github.com/ionos-cloud/sdk-go-bundle/userobjectstorage"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    bucket := "bucket_example" // string | 

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := userobjectstorage.NewAPIClient(configuration)
    resource, resp, err := apiClient.BucketsApi.HeadBucket(context.Background(), bucket).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketsApi.HeadBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**bucket** | **string** |  | |

### Other Parameters

Other parameters are passed through a pointer to an apiHeadBucketRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"BucketsApiService.HeadBucket"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), shared.ContextOperationServerIndices, map[string]int{
    "BucketsApiService.HeadBucket": 2,
})
ctx = context.WithValue(context.Background(), shared.ContextOperationServerVariables, map[string]map[string]string{
    "BucketsApiService.HeadBucket": {
    "port": "8443",
},
})
```


## ListBuckets

```go
var result ListAllMyBucketsResult = ListBuckets(ctx)
                      .Execute()
```

ListBuckets



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    userobjectstorage "github.com/ionos-cloud/sdk-go-bundle/userobjectstorage"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := userobjectstorage.NewAPIClient(configuration)
    resource, resp, err := apiClient.BucketsApi.ListBuckets(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketsApi.ListBuckets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ListBuckets`: ListAllMyBucketsResult
    fmt.Fprintf(os.Stdout, "Response from `BucketsApi.ListBuckets`: %v\n", resource)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to an apiListBucketsRequest struct via the builder pattern


### Return type

[**ListAllMyBucketsResult**](../models/ListAllMyBucketsResult.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"BucketsApiService.ListBuckets"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), shared.ContextOperationServerIndices, map[string]int{
    "BucketsApiService.ListBuckets": 2,
})
ctx = context.WithValue(context.Background(), shared.ContextOperationServerVariables, map[string]map[string]string{
    "BucketsApiService.ListBuckets": {
    "port": "8443",
},
})
```

