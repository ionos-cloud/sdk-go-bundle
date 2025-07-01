# \LoggingApi

All URIs are relative to *https://s3.eu-central-1.ionoscloud.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**GetBucketLogging**](LoggingApi.md#GetBucketLogging) | **Get** /{Bucket}?logging | GetBucketLogging|
|[**PutBucketLogging**](LoggingApi.md#PutBucketLogging) | **Put** /{Bucket}?logging | PutBucketLogging|



## GetBucketLogging

```go
var result GetBucketLogging200Response = GetBucketLogging(ctx, bucket)
                      .Logging(logging)
                      .Execute()
```

GetBucketLogging



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    userobjectstorage "github.com/ionos-cloud/sdk-go-bundle/products/userobjectstorage"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    bucket := "bucket_example" // string | The bucket name for which to get the logging information.
    logging := true // bool | 

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := userobjectstorage.NewAPIClient(configuration)
    resource, resp, err := apiClient.LoggingApi.GetBucketLogging(context.Background(), bucket).Logging(logging).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingApi.GetBucketLogging``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `GetBucketLogging`: GetBucketLogging200Response
    fmt.Fprintf(os.Stdout, "Response from `LoggingApi.GetBucketLogging`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**bucket** | **string** | The bucket name for which to get the logging information. | |

### Other Parameters

Other parameters are passed through a pointer to an apiGetBucketLoggingRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **logging** | **bool** |  | |

### Return type

[**GetBucketLogging200Response**](../models/GetBucketLogging200Response.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml



## PutBucketLogging

```go
var result  = PutBucketLogging(ctx, bucket)
                      .Logging(logging)
                      .PutBucketLoggingRequest(putBucketLoggingRequest)
                      .ContentMD5(contentMD5)
                      .Execute()
```

PutBucketLogging



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    userobjectstorage "github.com/ionos-cloud/sdk-go-bundle/products/userobjectstorage"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    bucket := "bucket_example" // string | The name of the bucket for which to set the logging parameters.
    logging := true // bool | 
    putBucketLoggingRequest := *openapiclient.NewPutBucketLoggingRequest(*openapiclient.NewPutBucketLoggingRequestBucketLoggingStatus()) // PutBucketLoggingRequest | 
    contentMD5 := TODO // string |  (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := userobjectstorage.NewAPIClient(configuration)
    resource, resp, err := apiClient.LoggingApi.PutBucketLogging(context.Background(), bucket).Logging(logging).PutBucketLoggingRequest(putBucketLoggingRequest).ContentMD5(contentMD5).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoggingApi.PutBucketLogging``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**bucket** | **string** | The name of the bucket for which to set the logging parameters. | |

### Other Parameters

Other parameters are passed through a pointer to an apiPutBucketLoggingRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **logging** | **bool** |  | |
| **putBucketLoggingRequest** | [**PutBucketLoggingRequest**](../models/PutBucketLoggingRequest.md) |  | |
| **contentMD5** | [**string**](../models/.md) |  | |

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: application/xml
- **Accept**: application/xml


