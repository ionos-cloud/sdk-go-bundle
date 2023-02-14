# \CertificatesApi

All URIs are relative to *https://api.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**CertificatesDelete**](CertificatesApi.md#CertificatesDelete) | **Delete** /certificatemanager/certificates/{certificateId} | Delete a Certificate by ID|
|[**CertificatesGet**](CertificatesApi.md#CertificatesGet) | **Get** /certificatemanager/certificates | Get Certificates|
|[**CertificatesGetById**](CertificatesApi.md#CertificatesGetById) | **Get** /certificatemanager/certificates/{certificateId} | Get a Certificate by ID|
|[**CertificatesPatch**](CertificatesApi.md#CertificatesPatch) | **Patch** /certificatemanager/certificates/{certificateId} | Update a Certificate Name by ID|
|[**CertificatesPost**](CertificatesApi.md#CertificatesPost) | **Post** /certificatemanager/certificates | Add a New Certificate|



## CertificatesDelete

```go
var result  = CertificatesDelete(ctx, certificateId)
                      .Execute()
```

Delete a Certificate by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    cert "github.com/ionos-cloud/sdk-go-bundle/products/cert"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    certificateId := "certificateId_example" // string | 

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := cert.NewAPIClient(configuration)
    resource, resp, err := apiClient.CertificatesApi.CertificatesDelete(context.Background(), certificateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesApi.CertificatesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**certificateId** | **string** |  | |

### Other Parameters

Other parameters are passed through a pointer to an apiCertificatesDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined



## CertificatesGet

```go
var result CertificateCollectionDto = CertificatesGet(ctx)
                      .Offset(offset)
                      .Limit(limit)
                      .Execute()
```

Get Certificates



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    cert "github.com/ionos-cloud/sdk-go-bundle/products/cert"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    offset := "offset_example" // string | 'Limit' and 'Offset' are optional; you can use these filter parameters to retrieve only part of the results obtained by a request.  Offset is the first element (from the complete list of elements) to be included in the response. (optional)
    limit := "limit_example" // string | 'Limit' and 'Offset' are optional; you can use these filter parameters to retrieve only part of the results of a query.  If both 'Offset' and 'Limit'are specified, the offset lines are skipped before counting the returned limit lines. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := cert.NewAPIClient(configuration)
    resource, resp, err := apiClient.CertificatesApi.CertificatesGet(context.Background()).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesApi.CertificatesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `CertificatesGet`: CertificateCollectionDto
    fmt.Fprintf(os.Stdout, "Response from `CertificatesApi.CertificatesGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiCertificatesGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **offset** | **string** | &#39;Limit&#39; and &#39;Offset&#39; are optional; you can use these filter parameters to retrieve only part of the results obtained by a request.  Offset is the first element (from the complete list of elements) to be included in the response. | |
| **limit** | **string** | &#39;Limit&#39; and &#39;Offset&#39; are optional; you can use these filter parameters to retrieve only part of the results of a query.  If both &#39;Offset&#39; and &#39;Limit&#39;are specified, the offset lines are skipped before counting the returned limit lines. | |

### Return type

[**CertificateCollectionDto**](../models/CertificateCollectionDto.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## CertificatesGetById

```go
var result CertificateDto = CertificatesGetById(ctx, certificateId)
                      .Execute()
```

Get a Certificate by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    cert "github.com/ionos-cloud/sdk-go-bundle/products/cert"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    certificateId := "certificateId_example" // string | 

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := cert.NewAPIClient(configuration)
    resource, resp, err := apiClient.CertificatesApi.CertificatesGetById(context.Background(), certificateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesApi.CertificatesGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `CertificatesGetById`: CertificateDto
    fmt.Fprintf(os.Stdout, "Response from `CertificatesApi.CertificatesGetById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**certificateId** | **string** |  | |

### Other Parameters

Other parameters are passed through a pointer to an apiCertificatesGetByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**CertificateDto**](../models/CertificateDto.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## CertificatesPatch

```go
var result CertificateDto = CertificatesPatch(ctx, certificateId)
                      .CertificatePatchDto(certificatePatchDto)
                      .Execute()
```

Update a Certificate Name by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    cert "github.com/ionos-cloud/sdk-go-bundle/products/cert"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    certificateId := "certificateId_example" // string | 
    certificatePatchDto := *openapiclient.NewCertificatePatchDto(*openapiclient.NewCertificatePatchPropertiesDto("My Certificate")) // CertificatePatchDto | 

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := cert.NewAPIClient(configuration)
    resource, resp, err := apiClient.CertificatesApi.CertificatesPatch(context.Background(), certificateId).CertificatePatchDto(certificatePatchDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesApi.CertificatesPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `CertificatesPatch`: CertificateDto
    fmt.Fprintf(os.Stdout, "Response from `CertificatesApi.CertificatesPatch`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**certificateId** | **string** |  | |

### Other Parameters

Other parameters are passed through a pointer to an apiCertificatesPatchRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **certificatePatchDto** | [**CertificatePatchDto**](../models/CertificatePatchDto.md) |  | |

### Return type

[**CertificateDto**](../models/CertificateDto.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## CertificatesPost

```go
var result CertificateDto = CertificatesPost(ctx)
                      .CertificatePostDto(certificatePostDto)
                      .Execute()
```

Add a New Certificate



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    cert "github.com/ionos-cloud/sdk-go-bundle/products/cert"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    certificatePostDto := *openapiclient.NewCertificatePostDto(*openapiclient.NewCertificatePostPropertiesDto("My Certificate", "-----BEGIN CERTIFICATE-----MIIE5TCCAs2gAwIBAgIBATANBgkqhkiG9w0BAQsFADA2MQswCQYDVQQGEwJSTzEK-----END CERTIFICATE-----", "-----BEGIN CERTIFICATE-----MIIDoTCCAokCFDrAUWffdxWJVz2Axl9lp/4xiUteMA0GCSqGSIb3DQEBCwUAMIGG-----END CERTIFICATE-----", "-----BEGIN RSA PRIVATE KEY-----MIIJKQIBAAKCAgEAzDehfqWBr+9q0pxwCDDRph7QSPiMbkDGaGKc+Fd2h3doT8Li-----END RSA PRIVATE KEY-----")) // CertificatePostDto | 

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := cert.NewAPIClient(configuration)
    resource, resp, err := apiClient.CertificatesApi.CertificatesPost(context.Background()).CertificatePostDto(certificatePostDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificatesApi.CertificatesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `CertificatesPost`: CertificateDto
    fmt.Fprintf(os.Stdout, "Response from `CertificatesApi.CertificatesPost`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiCertificatesPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **certificatePostDto** | [**CertificatePostDto**](../models/CertificatePostDto.md) |  | |

### Return type

[**CertificateDto**](../models/CertificateDto.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


