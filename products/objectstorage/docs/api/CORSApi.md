# \CORSApi

All URIs are relative to *https://s3.eu-central-3.ionoscloud.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**DeleteBucketCors**](CORSApi.md#DeleteBucketCors) | **Delete** /{Bucket}?cors | DeleteBucketCors|
|[**GetBucketCors**](CORSApi.md#GetBucketCors) | **Get** /{Bucket}?cors | GetBucketCors|
|[**PutBucketCors**](CORSApi.md#PutBucketCors) | **Put** /{Bucket}?cors | PutBucketCors|



## DeleteBucketCors

```go
var result  = DeleteBucketCors(ctx, bucket)
                      .Execute()
```

DeleteBucketCors



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    objectstorage "github.com/ionos-cloud/sdk-go-bundle/objectstorage"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    bucket := "bucket_example" // string | 

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := objectstorage.NewAPIClient(configuration)
    resource, resp, err := apiClient.CORSApi.DeleteBucketCors(context.Background(), bucket).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CORSApi.DeleteBucketCors``: %v\n", err)
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

Other parameters are passed through a pointer to an apiDeleteBucketCorsRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"CORSApiService.DeleteBucketCors"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), shared.ContextOperationServerIndices, map[string]int{
    "CORSApiService.DeleteBucketCors": 2,
})
ctx = context.WithValue(context.Background(), shared.ContextOperationServerVariables, map[string]map[string]string{
    "CORSApiService.DeleteBucketCors": {
    "port": "8443",
},
})
```


## GetBucketCors

```go
var result GetBucketCorsOutput = GetBucketCors(ctx, bucket)
                      .Execute()
```

GetBucketCors



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    objectstorage "github.com/ionos-cloud/sdk-go-bundle/objectstorage"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    bucket := "bucket_example" // string | 

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := objectstorage.NewAPIClient(configuration)
    resource, resp, err := apiClient.CORSApi.GetBucketCors(context.Background(), bucket).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CORSApi.GetBucketCors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `GetBucketCors`: GetBucketCorsOutput
    fmt.Fprintf(os.Stdout, "Response from `CORSApi.GetBucketCors`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**bucket** | **string** |  | |

### Other Parameters

Other parameters are passed through a pointer to an apiGetBucketCorsRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**GetBucketCorsOutput**](../models/GetBucketCorsOutput.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"CORSApiService.GetBucketCors"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), shared.ContextOperationServerIndices, map[string]int{
    "CORSApiService.GetBucketCors": 2,
})
ctx = context.WithValue(context.Background(), shared.ContextOperationServerVariables, map[string]map[string]string{
    "CORSApiService.GetBucketCors": {
    "port": "8443",
},
})
```


## PutBucketCors

```go
var result  = PutBucketCors(ctx, bucket)
                      .PutBucketCorsRequest(putBucketCorsRequest)
                      .ContentMD5(contentMD5)
                      .Execute()
```

PutBucketCors



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    objectstorage "github.com/ionos-cloud/sdk-go-bundle/objectstorage"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    bucket := "bucket_example" // string | 
    putBucketCorsRequest := *openapiclient.NewPutBucketCorsRequest() // PutBucketCorsRequest | 
    contentMD5 := "contentMD5_example" // string |  (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := objectstorage.NewAPIClient(configuration)
    resource, resp, err := apiClient.CORSApi.PutBucketCors(context.Background(), bucket).PutBucketCorsRequest(putBucketCorsRequest).ContentMD5(contentMD5).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CORSApi.PutBucketCors``: %v\n", err)
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

Other parameters are passed through a pointer to an apiPutBucketCorsRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **putBucketCorsRequest** | [**PutBucketCorsRequest**](../models/PutBucketCorsRequest.md) |  | |
| **contentMD5** | **string** |  | |

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: application/xml
- **Accept**: Not defined


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"CORSApiService.PutBucketCors"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), shared.ContextOperationServerIndices, map[string]int{
    "CORSApiService.PutBucketCors": 2,
})
ctx = context.WithValue(context.Background(), shared.ContextOperationServerVariables, map[string]map[string]string{
    "CORSApiService.PutBucketCors": {
    "port": "8443",
},
})
```

