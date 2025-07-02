# \VersioningApi

All URIs are relative to *https://s3.eu-central-1.ionoscloud.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**GetBucketVersioning**](VersioningApi.md#GetBucketVersioning) | **Get** /{Bucket}?versioning | GetBucketVersioning|
|[**PutBucketVersioning**](VersioningApi.md#PutBucketVersioning) | **Put** /{Bucket}?versioning | PutBucketVersioning|



## GetBucketVersioning

```go
var result GetBucketVersioningOutput = GetBucketVersioning(ctx, bucket)
                      .Versioning(versioning)
                      .Execute()
```

GetBucketVersioning



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
    versioning := true // bool | 

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := userobjectstorage.NewAPIClient(configuration)
    resource, resp, err := apiClient.VersioningApi.GetBucketVersioning(context.Background(), bucket).Versioning(versioning).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersioningApi.GetBucketVersioning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `GetBucketVersioning`: GetBucketVersioningOutput
    fmt.Fprintf(os.Stdout, "Response from `VersioningApi.GetBucketVersioning`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**bucket** | **string** |  | |

### Other Parameters

Other parameters are passed through a pointer to an apiGetBucketVersioningRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **versioning** | **bool** |  | |

### Return type

[**GetBucketVersioningOutput**](../models/GetBucketVersioningOutput.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"VersioningApiService.GetBucketVersioning"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), shared.ContextOperationServerIndices, map[string]int{
    "VersioningApiService.GetBucketVersioning": 2,
})
ctx = context.WithValue(context.Background(), shared.ContextOperationServerVariables, map[string]map[string]string{
    "VersioningApiService.GetBucketVersioning": {
    "port": "8443",
},
})
```


## PutBucketVersioning

```go
var result  = PutBucketVersioning(ctx, bucket)
                      .Versioning(versioning)
                      .PutBucketVersioningRequest(putBucketVersioningRequest)
                      .ContentMD5(contentMD5)
                      .Execute()
```

PutBucketVersioning



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
    versioning := true // bool | 
    putBucketVersioningRequest := *openapiclient.NewPutBucketVersioningRequest(*openapiclient.NewPutBucketVersioningRequestVersioningConfiguration()) // PutBucketVersioningRequest | 
    contentMD5 := "contentMD5_example" // string |  (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := userobjectstorage.NewAPIClient(configuration)
    resource, resp, err := apiClient.VersioningApi.PutBucketVersioning(context.Background(), bucket).Versioning(versioning).PutBucketVersioningRequest(putBucketVersioningRequest).ContentMD5(contentMD5).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersioningApi.PutBucketVersioning``: %v\n", err)
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

Other parameters are passed through a pointer to an apiPutBucketVersioningRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **versioning** | **bool** |  | |
| **putBucketVersioningRequest** | [**PutBucketVersioningRequest**](../models/PutBucketVersioningRequest.md) |  | |
| **contentMD5** | **string** |  | |

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: application/xml
- **Accept**: application/xml


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"VersioningApiService.PutBucketVersioning"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), shared.ContextOperationServerIndices, map[string]int{
    "VersioningApiService.PutBucketVersioning": 2,
})
ctx = context.WithValue(context.Background(), shared.ContextOperationServerVariables, map[string]map[string]string{
    "VersioningApiService.PutBucketVersioning": {
    "port": "8443",
},
})
```

