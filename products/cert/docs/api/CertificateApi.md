# \CertificateApi

All URIs are relative to *https://certificate-manager.de-fra.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**CertificatesDelete**](CertificateApi.md#CertificatesDelete) | **Delete** /certificates/{certificateId} | Delete Certificate|
|[**CertificatesFindById**](CertificateApi.md#CertificatesFindById) | **Get** /certificates/{certificateId} | Retrieve Certificate|
|[**CertificatesGet**](CertificateApi.md#CertificatesGet) | **Get** /certificates | Retrieve all Certificate|
|[**CertificatesPatch**](CertificateApi.md#CertificatesPatch) | **Patch** /certificates/{certificateId} | Updates Certificate|
|[**CertificatesPost**](CertificateApi.md#CertificatesPost) | **Post** /certificates | Create Certificate|



## CertificatesDelete

```go
var result  = CertificatesDelete(ctx, certificateId)
                      .Execute()
```

Delete Certificate



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
    certificateId := "cbee81a3-9389-57ba-bc50-393adcfca141" // string | The ID (UUID) of the Certificate.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := cert.NewAPIClient(configuration)
    resp, err := apiClient.CertificateApi.CertificatesDelete(context.Background(), certificateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificatesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**certificateId** | **string** | The ID (UUID) of the Certificate. | |

### Other Parameters

Other parameters are passed through a pointer to an apiCertificatesDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## CertificatesFindById

```go
var result CertificateRead = CertificatesFindById(ctx, certificateId)
                      .Execute()
```

Retrieve Certificate



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
    certificateId := "cbee81a3-9389-57ba-bc50-393adcfca141" // string | The ID (UUID) of the Certificate.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := cert.NewAPIClient(configuration)
    resource, resp, err := apiClient.CertificateApi.CertificatesFindById(context.Background(), certificateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificatesFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `CertificatesFindById`: CertificateRead
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.CertificatesFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**certificateId** | **string** | The ID (UUID) of the Certificate. | |

### Other Parameters

Other parameters are passed through a pointer to an apiCertificatesFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**CertificateRead**](../models/CertificateRead.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## CertificatesGet

```go
var result CertificateReadList = CertificatesGet(ctx)
                      .Offset(offset)
                      .Limit(limit)
                      .FilterCommonName(filterCommonName)
                      .FilterAutoCertificate(filterAutoCertificate)
                      .Execute()
```

Retrieve all Certificate



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
    offset := int32(0) // int32 | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. (optional) (default to 0)
    limit := int32(100) // int32 | The maximum number of elements to return. Use together with offset for pagination. (optional) (default to 100)
    filterCommonName := "www.example.com" // string | Filter by the common name (DNS).  (optional)
    filterAutoCertificate := "b471cd03-ef51-52c5-91a5-49195b0a04d4" // string | Filter by autoCertificateID.  (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := cert.NewAPIClient(configuration)
    resource, resp, err := apiClient.CertificateApi.CertificatesGet(context.Background()).Offset(offset).Limit(limit).FilterCommonName(filterCommonName).FilterAutoCertificate(filterAutoCertificate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificatesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `CertificatesGet`: CertificateReadList
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.CertificatesGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiCertificatesGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **offset** | **int32** | The first element (of the total list of elements) to include in the response. Use together with limit for pagination. | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return. Use together with offset for pagination. | [default to 100]|
| **filterCommonName** | **string** | Filter by the common name (DNS).  | |
| **filterAutoCertificate** | **string** | Filter by autoCertificateID.  | |

### Return type

[**CertificateReadList**](../models/CertificateReadList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## CertificatesPatch

```go
var result CertificateRead = CertificatesPatch(ctx, certificateId)
                      .CertificatePatch(certificatePatch)
                      .Execute()
```

Updates Certificate



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
    certificateId := "cbee81a3-9389-57ba-bc50-393adcfca141" // string | The ID (UUID) of the Certificate.
    certificatePatch := *openapiclient.NewCertificatePatch(*openapiclient.NewPatchName("My name")) // CertificatePatch | patch Certificate

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := cert.NewAPIClient(configuration)
    resource, resp, err := apiClient.CertificateApi.CertificatesPatch(context.Background(), certificateId).CertificatePatch(certificatePatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificatesPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `CertificatesPatch`: CertificateRead
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.CertificatesPatch`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**certificateId** | **string** | The ID (UUID) of the Certificate. | |

### Other Parameters

Other parameters are passed through a pointer to an apiCertificatesPatchRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **certificatePatch** | [**CertificatePatch**](../models/CertificatePatch.md) | patch Certificate | |

### Return type

[**CertificateRead**](../models/CertificateRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## CertificatesPost

```go
var result CertificateRead = CertificatesPost(ctx)
                      .CertificateCreate(certificateCreate)
                      .Execute()
```

Create Certificate



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
    certificateCreate := *openapiclient.NewCertificateCreate(*openapiclient.NewCertificate("My Certificate", "-----BEGIN CERTIFICATE-----
MIIE5TCCAs2gAwIBAgIBATANBgkqhkiG9w0BAQsFADA2MQswCQYDVQQGEwJSTzEK
-----END CERTIFICATE-----
", "-----BEGIN CERTIFICATE-----
MIIDoTCCAokCFDrAUWffdxWJVz2Axl9lp/4xiUteMA0GCSqGSIb3DQEBCwUAMIGG
-----END CERTIFICATE-----
", "-----BEGIN RSA PRIVATE KEY-----
MIIJKQIBAAKCAgEAzDehfqWBr+9q0pxwCDDRph7QSPiMbkDGaGKc+Fd2h3doT8Li
-----END RSA PRIVATE KEY-----
")) // CertificateCreate | Certificate to create.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := cert.NewAPIClient(configuration)
    resource, resp, err := apiClient.CertificateApi.CertificatesPost(context.Background()).CertificateCreate(certificateCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateApi.CertificatesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `CertificatesPost`: CertificateRead
    fmt.Fprintf(os.Stdout, "Response from `CertificateApi.CertificatesPost`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiCertificatesPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **certificateCreate** | [**CertificateCreate**](../models/CertificateCreate.md) | Certificate to create. | |

### Return type

[**CertificateRead**](../models/CertificateRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


