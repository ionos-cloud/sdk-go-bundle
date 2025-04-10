# \EncryptionApi

All URIs are relative to *https://s3.eu-central-3.ionoscloud.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**DeleteBucketEncryption**](EncryptionApi.md#DeleteBucketEncryption) | **Delete** /{Bucket}?encryption | DeleteBucketEncryption|
|[**GetBucketEncryption**](EncryptionApi.md#GetBucketEncryption) | **Get** /{Bucket}?encryption | GetBucketEncryption|
|[**PutBucketEncryption**](EncryptionApi.md#PutBucketEncryption) | **Put** /{Bucket}?encryption | PutBucketEncryption|



## DeleteBucketEncryption

```go
var result  = DeleteBucketEncryption(ctx, bucket)
                      .Execute()
```

DeleteBucketEncryption



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
    resource, resp, err := apiClient.EncryptionApi.DeleteBucketEncryption(context.Background(), bucket).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EncryptionApi.DeleteBucketEncryption``: %v\n", err)
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

Other parameters are passed through a pointer to an apiDeleteBucketEncryptionRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"EncryptionApiService.DeleteBucketEncryption"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), shared.ContextOperationServerIndices, map[string]int{
    "EncryptionApiService.DeleteBucketEncryption": 2,
})
ctx = context.WithValue(context.Background(), shared.ContextOperationServerVariables, map[string]map[string]string{
    "EncryptionApiService.DeleteBucketEncryption": {
    "port": "8443",
},
})
```


## GetBucketEncryption

```go
var result ServerSideEncryptionConfiguration = GetBucketEncryption(ctx, bucket)
                      .Execute()
```

GetBucketEncryption



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
    resource, resp, err := apiClient.EncryptionApi.GetBucketEncryption(context.Background(), bucket).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EncryptionApi.GetBucketEncryption``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `GetBucketEncryption`: ServerSideEncryptionConfiguration
    fmt.Fprintf(os.Stdout, "Response from `EncryptionApi.GetBucketEncryption`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**bucket** | **string** |  | |

### Other Parameters

Other parameters are passed through a pointer to an apiGetBucketEncryptionRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**ServerSideEncryptionConfiguration**](../models/ServerSideEncryptionConfiguration.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"EncryptionApiService.GetBucketEncryption"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), shared.ContextOperationServerIndices, map[string]int{
    "EncryptionApiService.GetBucketEncryption": 2,
})
ctx = context.WithValue(context.Background(), shared.ContextOperationServerVariables, map[string]map[string]string{
    "EncryptionApiService.GetBucketEncryption": {
    "port": "8443",
},
})
```


## PutBucketEncryption

```go
var result  = PutBucketEncryption(ctx, bucket)
                      .PutBucketEncryptionRequest(putBucketEncryptionRequest)
                      .Execute()
```

PutBucketEncryption



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
    putBucketEncryptionRequest := *openapiclient.NewPutBucketEncryptionRequest() // PutBucketEncryptionRequest | 

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := objectstorage.NewAPIClient(configuration)
    resource, resp, err := apiClient.EncryptionApi.PutBucketEncryption(context.Background(), bucket).PutBucketEncryptionRequest(putBucketEncryptionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EncryptionApi.PutBucketEncryption``: %v\n", err)
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

Other parameters are passed through a pointer to an apiPutBucketEncryptionRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **putBucketEncryptionRequest** | [**PutBucketEncryptionRequest**](../models/PutBucketEncryptionRequest.md) |  | |

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: application/xml
- **Accept**: Not defined


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"EncryptionApiService.PutBucketEncryption"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), shared.ContextOperationServerIndices, map[string]int{
    "EncryptionApiService.PutBucketEncryption": 2,
})
ctx = context.WithValue(context.Background(), shared.ContextOperationServerVariables, map[string]map[string]string{
    "EncryptionApiService.PutBucketEncryption": {
    "port": "8443",
},
})
```

