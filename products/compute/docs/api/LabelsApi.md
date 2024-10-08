# \LabelsApi

All URIs are relative to *https://api.ionos.com/cloudapi/v6*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**DatacentersLabelsDelete**](LabelsApi.md#DatacentersLabelsDelete) | **Delete** /datacenters/{datacenterId}/labels/{key} | Delete data center labels|
|[**DatacentersLabelsFindByKey**](LabelsApi.md#DatacentersLabelsFindByKey) | **Get** /datacenters/{datacenterId}/labels/{key} | Retrieve data center labels|
|[**DatacentersLabelsGet**](LabelsApi.md#DatacentersLabelsGet) | **Get** /datacenters/{datacenterId}/labels | List data center labels|
|[**DatacentersLabelsPost**](LabelsApi.md#DatacentersLabelsPost) | **Post** /datacenters/{datacenterId}/labels | Create a Data Center Label|
|[**DatacentersLabelsPut**](LabelsApi.md#DatacentersLabelsPut) | **Put** /datacenters/{datacenterId}/labels/{key} | Modify a Data Center Label by Key|
|[**DatacentersServersLabelsDelete**](LabelsApi.md#DatacentersServersLabelsDelete) | **Delete** /datacenters/{datacenterId}/servers/{serverId}/labels/{key} | Delete server labels|
|[**DatacentersServersLabelsFindByKey**](LabelsApi.md#DatacentersServersLabelsFindByKey) | **Get** /datacenters/{datacenterId}/servers/{serverId}/labels/{key} | Retrieve server labels|
|[**DatacentersServersLabelsGet**](LabelsApi.md#DatacentersServersLabelsGet) | **Get** /datacenters/{datacenterId}/servers/{serverId}/labels | List server labels|
|[**DatacentersServersLabelsPost**](LabelsApi.md#DatacentersServersLabelsPost) | **Post** /datacenters/{datacenterId}/servers/{serverId}/labels | Create a Server Label|
|[**DatacentersServersLabelsPut**](LabelsApi.md#DatacentersServersLabelsPut) | **Put** /datacenters/{datacenterId}/servers/{serverId}/labels/{key} | Modify a Server Label|
|[**DatacentersVolumesLabelsDelete**](LabelsApi.md#DatacentersVolumesLabelsDelete) | **Delete** /datacenters/{datacenterId}/volumes/{volumeId}/labels/{key} | Delete volume labels|
|[**DatacentersVolumesLabelsFindByKey**](LabelsApi.md#DatacentersVolumesLabelsFindByKey) | **Get** /datacenters/{datacenterId}/volumes/{volumeId}/labels/{key} | Retrieve volume labels|
|[**DatacentersVolumesLabelsGet**](LabelsApi.md#DatacentersVolumesLabelsGet) | **Get** /datacenters/{datacenterId}/volumes/{volumeId}/labels | List volume labels|
|[**DatacentersVolumesLabelsPost**](LabelsApi.md#DatacentersVolumesLabelsPost) | **Post** /datacenters/{datacenterId}/volumes/{volumeId}/labels | Create a Volume Label|
|[**DatacentersVolumesLabelsPut**](LabelsApi.md#DatacentersVolumesLabelsPut) | **Put** /datacenters/{datacenterId}/volumes/{volumeId}/labels/{key} | Modify a Volume Label|
|[**ImagesLabelsDelete**](LabelsApi.md#ImagesLabelsDelete) | **Delete** /images/{imageId}/labels/{key} | Delete image label|
|[**ImagesLabelsFindByKey**](LabelsApi.md#ImagesLabelsFindByKey) | **Get** /images/{imageId}/labels/{key} | Retrieve image labels|
|[**ImagesLabelsGet**](LabelsApi.md#ImagesLabelsGet) | **Get** /images/{imageId}/labels | List image labels|
|[**ImagesLabelsPost**](LabelsApi.md#ImagesLabelsPost) | **Post** /images/{imageId}/labels | Create an Image Label|
|[**ImagesLabelsPut**](LabelsApi.md#ImagesLabelsPut) | **Put** /images/{imageId}/labels/{key} | Modify an Image Label by Key|
|[**IpblocksLabelsDelete**](LabelsApi.md#IpblocksLabelsDelete) | **Delete** /ipblocks/{ipblockId}/labels/{key} | Delete IP block labels|
|[**IpblocksLabelsFindByKey**](LabelsApi.md#IpblocksLabelsFindByKey) | **Get** /ipblocks/{ipblockId}/labels/{key} | Retrieve IP block labels|
|[**IpblocksLabelsGet**](LabelsApi.md#IpblocksLabelsGet) | **Get** /ipblocks/{ipblockId}/labels | List IP block labels|
|[**IpblocksLabelsPost**](LabelsApi.md#IpblocksLabelsPost) | **Post** /ipblocks/{ipblockId}/labels | Create IP block labels|
|[**IpblocksLabelsPut**](LabelsApi.md#IpblocksLabelsPut) | **Put** /ipblocks/{ipblockId}/labels/{key} | Modify a IP Block Label by ID|
|[**LabelsFindByUrn**](LabelsApi.md#LabelsFindByUrn) | **Get** /labels/{labelurn} | Retrieve labels by URN|
|[**LabelsGet**](LabelsApi.md#LabelsGet) | **Get** /labels | List labels |
|[**SnapshotsLabelsDelete**](LabelsApi.md#SnapshotsLabelsDelete) | **Delete** /snapshots/{snapshotId}/labels/{key} | Delete snapshot labels|
|[**SnapshotsLabelsFindByKey**](LabelsApi.md#SnapshotsLabelsFindByKey) | **Get** /snapshots/{snapshotId}/labels/{key} | Retrieve snapshot labels|
|[**SnapshotsLabelsGet**](LabelsApi.md#SnapshotsLabelsGet) | **Get** /snapshots/{snapshotId}/labels | List snapshot labels|
|[**SnapshotsLabelsPost**](LabelsApi.md#SnapshotsLabelsPost) | **Post** /snapshots/{snapshotId}/labels | Create a Snapshot Label|
|[**SnapshotsLabelsPut**](LabelsApi.md#SnapshotsLabelsPut) | **Put** /snapshots/{snapshotId}/labels/{key} | Modify a Snapshot Label by ID|



## DatacentersLabelsDelete

```go
var result  = DatacentersLabelsDelete(ctx, datacenterId, key)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Delete data center labels



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    compute "github.com/ionos-cloud/sdk-go-bundle/products/compute"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    key := "key_example" // string | The label key
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resp, err := apiClient.LabelsApi.DatacentersLabelsDelete(context.Background(), datacenterId, key).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.DatacentersLabelsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**key** | **string** | The label key | |

### Other Parameters

Other parameters are passed through a pointer to an apiDatacentersLabelsDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersLabelsFindByKey

```go
var result LabelResource = DatacentersLabelsFindByKey(ctx, datacenterId, key)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve data center labels



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    compute "github.com/ionos-cloud/sdk-go-bundle/products/compute"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    key := "key_example" // string | The label key
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.LabelsApi.DatacentersLabelsFindByKey(context.Background(), datacenterId, key).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.DatacentersLabelsFindByKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `DatacentersLabelsFindByKey`: LabelResource
    fmt.Fprintf(os.Stdout, "Response from `LabelsApi.DatacentersLabelsFindByKey`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**key** | **string** | The label key | |

### Other Parameters

Other parameters are passed through a pointer to an apiDatacentersLabelsFindByKeyRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**LabelResource**](../models/LabelResource.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersLabelsGet

```go
var result LabelResources = DatacentersLabelsGet(ctx, datacenterId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

List data center labels



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    compute "github.com/ionos-cloud/sdk-go-bundle/products/compute"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.LabelsApi.DatacentersLabelsGet(context.Background(), datacenterId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.DatacentersLabelsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `DatacentersLabelsGet`: LabelResources
    fmt.Fprintf(os.Stdout, "Response from `LabelsApi.DatacentersLabelsGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |

### Other Parameters

Other parameters are passed through a pointer to an apiDatacentersLabelsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**LabelResources**](../models/LabelResources.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersLabelsPost

```go
var result LabelResource = DatacentersLabelsPost(ctx, datacenterId)
                      .Label(label)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Create a Data Center Label



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    compute "github.com/ionos-cloud/sdk-go-bundle/products/compute"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    label := *openapiclient.NewLabelResource(*openapiclient.NewLabelResourceProperties()) // LabelResource | The label to create.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.LabelsApi.DatacentersLabelsPost(context.Background(), datacenterId).Label(label).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.DatacentersLabelsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `DatacentersLabelsPost`: LabelResource
    fmt.Fprintf(os.Stdout, "Response from `LabelsApi.DatacentersLabelsPost`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |

### Other Parameters

Other parameters are passed through a pointer to an apiDatacentersLabelsPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **label** | [**LabelResource**](../models/LabelResource.md) | The label to create. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**LabelResource**](../models/LabelResource.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersLabelsPut

```go
var result LabelResource = DatacentersLabelsPut(ctx, datacenterId, key)
                      .Label(label)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Modify a Data Center Label by Key



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    compute "github.com/ionos-cloud/sdk-go-bundle/products/compute"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    key := "key_example" // string | The label key
    label := *openapiclient.NewLabelResource(*openapiclient.NewLabelResourceProperties()) // LabelResource | The modified label
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.LabelsApi.DatacentersLabelsPut(context.Background(), datacenterId, key).Label(label).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.DatacentersLabelsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `DatacentersLabelsPut`: LabelResource
    fmt.Fprintf(os.Stdout, "Response from `LabelsApi.DatacentersLabelsPut`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**key** | **string** | The label key | |

### Other Parameters

Other parameters are passed through a pointer to an apiDatacentersLabelsPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **label** | [**LabelResource**](../models/LabelResource.md) | The modified label | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**LabelResource**](../models/LabelResource.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersServersLabelsDelete

```go
var result  = DatacentersServersLabelsDelete(ctx, datacenterId, serverId, key)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Delete server labels



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    compute "github.com/ionos-cloud/sdk-go-bundle/products/compute"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    serverId := "serverId_example" // string | The unique ID of the server.
    key := "key_example" // string | The label key
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resp, err := apiClient.LabelsApi.DatacentersServersLabelsDelete(context.Background(), datacenterId, serverId, key).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.DatacentersServersLabelsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**serverId** | **string** | The unique ID of the server. | |
|**key** | **string** | The label key | |

### Other Parameters

Other parameters are passed through a pointer to an apiDatacentersServersLabelsDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersServersLabelsFindByKey

```go
var result LabelResource = DatacentersServersLabelsFindByKey(ctx, datacenterId, serverId, key)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve server labels



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    compute "github.com/ionos-cloud/sdk-go-bundle/products/compute"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    serverId := "serverId_example" // string | The unique ID of the server.
    key := "key_example" // string | The label key
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.LabelsApi.DatacentersServersLabelsFindByKey(context.Background(), datacenterId, serverId, key).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.DatacentersServersLabelsFindByKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `DatacentersServersLabelsFindByKey`: LabelResource
    fmt.Fprintf(os.Stdout, "Response from `LabelsApi.DatacentersServersLabelsFindByKey`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**serverId** | **string** | The unique ID of the server. | |
|**key** | **string** | The label key | |

### Other Parameters

Other parameters are passed through a pointer to an apiDatacentersServersLabelsFindByKeyRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**LabelResource**](../models/LabelResource.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersServersLabelsGet

```go
var result LabelResources = DatacentersServersLabelsGet(ctx, datacenterId, serverId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

List server labels



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    compute "github.com/ionos-cloud/sdk-go-bundle/products/compute"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    serverId := "serverId_example" // string | The unique ID of the server.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.LabelsApi.DatacentersServersLabelsGet(context.Background(), datacenterId, serverId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.DatacentersServersLabelsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `DatacentersServersLabelsGet`: LabelResources
    fmt.Fprintf(os.Stdout, "Response from `LabelsApi.DatacentersServersLabelsGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**serverId** | **string** | The unique ID of the server. | |

### Other Parameters

Other parameters are passed through a pointer to an apiDatacentersServersLabelsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**LabelResources**](../models/LabelResources.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersServersLabelsPost

```go
var result LabelResource = DatacentersServersLabelsPost(ctx, datacenterId, serverId)
                      .Label(label)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Create a Server Label



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    compute "github.com/ionos-cloud/sdk-go-bundle/products/compute"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    serverId := "serverId_example" // string | The unique ID of the server.
    label := *openapiclient.NewLabelResource(*openapiclient.NewLabelResourceProperties()) // LabelResource | The label to create.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.LabelsApi.DatacentersServersLabelsPost(context.Background(), datacenterId, serverId).Label(label).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.DatacentersServersLabelsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `DatacentersServersLabelsPost`: LabelResource
    fmt.Fprintf(os.Stdout, "Response from `LabelsApi.DatacentersServersLabelsPost`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**serverId** | **string** | The unique ID of the server. | |

### Other Parameters

Other parameters are passed through a pointer to an apiDatacentersServersLabelsPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **label** | [**LabelResource**](../models/LabelResource.md) | The label to create. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**LabelResource**](../models/LabelResource.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersServersLabelsPut

```go
var result LabelResource = DatacentersServersLabelsPut(ctx, datacenterId, serverId, key)
                      .Label(label)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Modify a Server Label



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    compute "github.com/ionos-cloud/sdk-go-bundle/products/compute"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    serverId := "serverId_example" // string | The unique ID of the server.
    key := "key_example" // string | The label key
    label := *openapiclient.NewLabelResource(*openapiclient.NewLabelResourceProperties()) // LabelResource | The modified label
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.LabelsApi.DatacentersServersLabelsPut(context.Background(), datacenterId, serverId, key).Label(label).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.DatacentersServersLabelsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `DatacentersServersLabelsPut`: LabelResource
    fmt.Fprintf(os.Stdout, "Response from `LabelsApi.DatacentersServersLabelsPut`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**serverId** | **string** | The unique ID of the server. | |
|**key** | **string** | The label key | |

### Other Parameters

Other parameters are passed through a pointer to an apiDatacentersServersLabelsPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **label** | [**LabelResource**](../models/LabelResource.md) | The modified label | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**LabelResource**](../models/LabelResource.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersVolumesLabelsDelete

```go
var result  = DatacentersVolumesLabelsDelete(ctx, datacenterId, volumeId, key)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Delete volume labels



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    compute "github.com/ionos-cloud/sdk-go-bundle/products/compute"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    volumeId := "volumeId_example" // string | The unique ID of the volume.
    key := "key_example" // string | The label key
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resp, err := apiClient.LabelsApi.DatacentersVolumesLabelsDelete(context.Background(), datacenterId, volumeId, key).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.DatacentersVolumesLabelsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**volumeId** | **string** | The unique ID of the volume. | |
|**key** | **string** | The label key | |

### Other Parameters

Other parameters are passed through a pointer to an apiDatacentersVolumesLabelsDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersVolumesLabelsFindByKey

```go
var result LabelResource = DatacentersVolumesLabelsFindByKey(ctx, datacenterId, volumeId, key)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve volume labels



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    compute "github.com/ionos-cloud/sdk-go-bundle/products/compute"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    volumeId := "volumeId_example" // string | The unique ID of the volume.
    key := "key_example" // string | The label key
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.LabelsApi.DatacentersVolumesLabelsFindByKey(context.Background(), datacenterId, volumeId, key).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.DatacentersVolumesLabelsFindByKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `DatacentersVolumesLabelsFindByKey`: LabelResource
    fmt.Fprintf(os.Stdout, "Response from `LabelsApi.DatacentersVolumesLabelsFindByKey`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**volumeId** | **string** | The unique ID of the volume. | |
|**key** | **string** | The label key | |

### Other Parameters

Other parameters are passed through a pointer to an apiDatacentersVolumesLabelsFindByKeyRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**LabelResource**](../models/LabelResource.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersVolumesLabelsGet

```go
var result LabelResources = DatacentersVolumesLabelsGet(ctx, datacenterId, volumeId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

List volume labels



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    compute "github.com/ionos-cloud/sdk-go-bundle/products/compute"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    volumeId := "volumeId_example" // string | The unique ID of the volume.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.LabelsApi.DatacentersVolumesLabelsGet(context.Background(), datacenterId, volumeId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.DatacentersVolumesLabelsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `DatacentersVolumesLabelsGet`: LabelResources
    fmt.Fprintf(os.Stdout, "Response from `LabelsApi.DatacentersVolumesLabelsGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**volumeId** | **string** | The unique ID of the volume. | |

### Other Parameters

Other parameters are passed through a pointer to an apiDatacentersVolumesLabelsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**LabelResources**](../models/LabelResources.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersVolumesLabelsPost

```go
var result LabelResource = DatacentersVolumesLabelsPost(ctx, datacenterId, volumeId)
                      .Label(label)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Create a Volume Label



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    compute "github.com/ionos-cloud/sdk-go-bundle/products/compute"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    volumeId := "volumeId_example" // string | The unique ID of the volume.
    label := *openapiclient.NewLabelResource(*openapiclient.NewLabelResourceProperties()) // LabelResource | The label to create.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.LabelsApi.DatacentersVolumesLabelsPost(context.Background(), datacenterId, volumeId).Label(label).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.DatacentersVolumesLabelsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `DatacentersVolumesLabelsPost`: LabelResource
    fmt.Fprintf(os.Stdout, "Response from `LabelsApi.DatacentersVolumesLabelsPost`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**volumeId** | **string** | The unique ID of the volume. | |

### Other Parameters

Other parameters are passed through a pointer to an apiDatacentersVolumesLabelsPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **label** | [**LabelResource**](../models/LabelResource.md) | The label to create. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**LabelResource**](../models/LabelResource.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersVolumesLabelsPut

```go
var result LabelResource = DatacentersVolumesLabelsPut(ctx, datacenterId, volumeId, key)
                      .Label(label)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Modify a Volume Label



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    compute "github.com/ionos-cloud/sdk-go-bundle/products/compute"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    volumeId := "volumeId_example" // string | The unique ID of the volume.
    key := "key_example" // string | The label key
    label := *openapiclient.NewLabelResource(*openapiclient.NewLabelResourceProperties()) // LabelResource | The modified label
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.LabelsApi.DatacentersVolumesLabelsPut(context.Background(), datacenterId, volumeId, key).Label(label).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.DatacentersVolumesLabelsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `DatacentersVolumesLabelsPut`: LabelResource
    fmt.Fprintf(os.Stdout, "Response from `LabelsApi.DatacentersVolumesLabelsPut`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**volumeId** | **string** | The unique ID of the volume. | |
|**key** | **string** | The label key | |

### Other Parameters

Other parameters are passed through a pointer to an apiDatacentersVolumesLabelsPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **label** | [**LabelResource**](../models/LabelResource.md) | The modified label | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**LabelResource**](../models/LabelResource.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## ImagesLabelsDelete

```go
var result  = ImagesLabelsDelete(ctx, imageId, key)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Delete image label



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    compute "github.com/ionos-cloud/sdk-go-bundle/products/compute"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    imageId := "imageId_example" // string | The unique ID of the image.
    key := "key_example" // string | The label key
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resp, err := apiClient.LabelsApi.ImagesLabelsDelete(context.Background(), imageId, key).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.ImagesLabelsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**imageId** | **string** | The unique ID of the image. | |
|**key** | **string** | The label key | |

### Other Parameters

Other parameters are passed through a pointer to an apiImagesLabelsDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ImagesLabelsFindByKey

```go
var result LabelResource = ImagesLabelsFindByKey(ctx, imageId, key)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve image labels



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    compute "github.com/ionos-cloud/sdk-go-bundle/products/compute"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    imageId := "imageId_example" // string | The unique ID of the image.
    key := "key_example" // string | The label key
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.LabelsApi.ImagesLabelsFindByKey(context.Background(), imageId, key).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.ImagesLabelsFindByKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ImagesLabelsFindByKey`: LabelResource
    fmt.Fprintf(os.Stdout, "Response from `LabelsApi.ImagesLabelsFindByKey`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**imageId** | **string** | The unique ID of the image. | |
|**key** | **string** | The label key | |

### Other Parameters

Other parameters are passed through a pointer to an apiImagesLabelsFindByKeyRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**LabelResource**](../models/LabelResource.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ImagesLabelsGet

```go
var result LabelResources = ImagesLabelsGet(ctx, imageId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

List image labels



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    compute "github.com/ionos-cloud/sdk-go-bundle/products/compute"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    imageId := "imageId_example" // string | The unique ID of the image.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.LabelsApi.ImagesLabelsGet(context.Background(), imageId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.ImagesLabelsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ImagesLabelsGet`: LabelResources
    fmt.Fprintf(os.Stdout, "Response from `LabelsApi.ImagesLabelsGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**imageId** | **string** | The unique ID of the image. | |

### Other Parameters

Other parameters are passed through a pointer to an apiImagesLabelsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**LabelResources**](../models/LabelResources.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## ImagesLabelsPost

```go
var result LabelResource = ImagesLabelsPost(ctx, imageId)
                      .Label(label)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Create an Image Label



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    compute "github.com/ionos-cloud/sdk-go-bundle/products/compute"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    imageId := "imageId_example" // string | The unique ID of the image
    label := *openapiclient.NewLabelResource(*openapiclient.NewLabelResourceProperties()) // LabelResource | The label to create.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.LabelsApi.ImagesLabelsPost(context.Background(), imageId).Label(label).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.ImagesLabelsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ImagesLabelsPost`: LabelResource
    fmt.Fprintf(os.Stdout, "Response from `LabelsApi.ImagesLabelsPost`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**imageId** | **string** | The unique ID of the image | |

### Other Parameters

Other parameters are passed through a pointer to an apiImagesLabelsPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **label** | [**LabelResource**](../models/LabelResource.md) | The label to create. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**LabelResource**](../models/LabelResource.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## ImagesLabelsPut

```go
var result LabelResource = ImagesLabelsPut(ctx, imageId, key)
                      .Label(label)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Modify an Image Label by Key



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    compute "github.com/ionos-cloud/sdk-go-bundle/products/compute"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    imageId := "imageId_example" // string | The unique ID of the image.
    key := "key_example" // string | The label key
    label := *openapiclient.NewLabelResource(*openapiclient.NewLabelResourceProperties()) // LabelResource | The modified label
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.LabelsApi.ImagesLabelsPut(context.Background(), imageId, key).Label(label).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.ImagesLabelsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ImagesLabelsPut`: LabelResource
    fmt.Fprintf(os.Stdout, "Response from `LabelsApi.ImagesLabelsPut`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**imageId** | **string** | The unique ID of the image. | |
|**key** | **string** | The label key | |

### Other Parameters

Other parameters are passed through a pointer to an apiImagesLabelsPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **label** | [**LabelResource**](../models/LabelResource.md) | The modified label | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**LabelResource**](../models/LabelResource.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## IpblocksLabelsDelete

```go
var result  = IpblocksLabelsDelete(ctx, ipblockId, key)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Delete IP block labels



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    compute "github.com/ionos-cloud/sdk-go-bundle/products/compute"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    ipblockId := "ipblockId_example" // string | The unique ID of the IP block.
    key := "key_example" // string | The label key
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resp, err := apiClient.LabelsApi.IpblocksLabelsDelete(context.Background(), ipblockId, key).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.IpblocksLabelsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**ipblockId** | **string** | The unique ID of the IP block. | |
|**key** | **string** | The label key | |

### Other Parameters

Other parameters are passed through a pointer to an apiIpblocksLabelsDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## IpblocksLabelsFindByKey

```go
var result LabelResource = IpblocksLabelsFindByKey(ctx, ipblockId, key)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve IP block labels



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    compute "github.com/ionos-cloud/sdk-go-bundle/products/compute"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    ipblockId := "ipblockId_example" // string | The unique ID of the IP block.
    key := "key_example" // string | The label key
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.LabelsApi.IpblocksLabelsFindByKey(context.Background(), ipblockId, key).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.IpblocksLabelsFindByKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `IpblocksLabelsFindByKey`: LabelResource
    fmt.Fprintf(os.Stdout, "Response from `LabelsApi.IpblocksLabelsFindByKey`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**ipblockId** | **string** | The unique ID of the IP block. | |
|**key** | **string** | The label key | |

### Other Parameters

Other parameters are passed through a pointer to an apiIpblocksLabelsFindByKeyRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**LabelResource**](../models/LabelResource.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## IpblocksLabelsGet

```go
var result LabelResources = IpblocksLabelsGet(ctx, ipblockId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

List IP block labels



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    compute "github.com/ionos-cloud/sdk-go-bundle/products/compute"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    ipblockId := "ipblockId_example" // string | The unique ID of the IP block.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.LabelsApi.IpblocksLabelsGet(context.Background(), ipblockId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.IpblocksLabelsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `IpblocksLabelsGet`: LabelResources
    fmt.Fprintf(os.Stdout, "Response from `LabelsApi.IpblocksLabelsGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**ipblockId** | **string** | The unique ID of the IP block. | |

### Other Parameters

Other parameters are passed through a pointer to an apiIpblocksLabelsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**LabelResources**](../models/LabelResources.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## IpblocksLabelsPost

```go
var result LabelResource = IpblocksLabelsPost(ctx, ipblockId)
                      .Label(label)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Create IP block labels



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    compute "github.com/ionos-cloud/sdk-go-bundle/products/compute"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    ipblockId := "ipblockId_example" // string | The unique ID of the IP block.
    label := *openapiclient.NewLabelResource(*openapiclient.NewLabelResourceProperties()) // LabelResource | The label to create.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.LabelsApi.IpblocksLabelsPost(context.Background(), ipblockId).Label(label).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.IpblocksLabelsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `IpblocksLabelsPost`: LabelResource
    fmt.Fprintf(os.Stdout, "Response from `LabelsApi.IpblocksLabelsPost`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**ipblockId** | **string** | The unique ID of the IP block. | |

### Other Parameters

Other parameters are passed through a pointer to an apiIpblocksLabelsPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **label** | [**LabelResource**](../models/LabelResource.md) | The label to create. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**LabelResource**](../models/LabelResource.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## IpblocksLabelsPut

```go
var result LabelResource = IpblocksLabelsPut(ctx, ipblockId, key)
                      .Label(label)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Modify a IP Block Label by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    compute "github.com/ionos-cloud/sdk-go-bundle/products/compute"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    ipblockId := "ipblockId_example" // string | The unique ID of the IP block.
    key := "key_example" // string | The label key
    label := *openapiclient.NewLabelResource(*openapiclient.NewLabelResourceProperties()) // LabelResource | The modified label
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.LabelsApi.IpblocksLabelsPut(context.Background(), ipblockId, key).Label(label).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.IpblocksLabelsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `IpblocksLabelsPut`: LabelResource
    fmt.Fprintf(os.Stdout, "Response from `LabelsApi.IpblocksLabelsPut`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**ipblockId** | **string** | The unique ID of the IP block. | |
|**key** | **string** | The label key | |

### Other Parameters

Other parameters are passed through a pointer to an apiIpblocksLabelsPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **label** | [**LabelResource**](../models/LabelResource.md) | The modified label | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**LabelResource**](../models/LabelResource.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## LabelsFindByUrn

```go
var result Label = LabelsFindByUrn(ctx, labelurn)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve labels by URN



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    compute "github.com/ionos-cloud/sdk-go-bundle/products/compute"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    labelurn := "labelurn_example" // string | The label URN; URN is unique for each label, and consists of:  urn:label:<resource_type>:<resource_uuid>:<key><key>
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.LabelsApi.LabelsFindByUrn(context.Background(), labelurn).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.LabelsFindByUrn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `LabelsFindByUrn`: Label
    fmt.Fprintf(os.Stdout, "Response from `LabelsApi.LabelsFindByUrn`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**labelurn** | **string** | The label URN; URN is unique for each label, and consists of:  urn:label:&lt;resource_type&gt;:&lt;resource_uuid&gt;:&lt;key&gt;&lt;key&gt; | |

### Other Parameters

Other parameters are passed through a pointer to an apiLabelsFindByUrnRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**Label**](../models/Label.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## LabelsGet

```go
var result Labels = LabelsGet(ctx)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

List labels 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    compute "github.com/ionos-cloud/sdk-go-bundle/products/compute"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.LabelsApi.LabelsGet(context.Background()).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.LabelsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `LabelsGet`: Labels
    fmt.Fprintf(os.Stdout, "Response from `LabelsApi.LabelsGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiLabelsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**Labels**](../models/Labels.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## SnapshotsLabelsDelete

```go
var result  = SnapshotsLabelsDelete(ctx, snapshotId, key)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Delete snapshot labels



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    compute "github.com/ionos-cloud/sdk-go-bundle/products/compute"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    snapshotId := "snapshotId_example" // string | The unique ID of the snapshot.
    key := "key_example" // string | The label key
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resp, err := apiClient.LabelsApi.SnapshotsLabelsDelete(context.Background(), snapshotId, key).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.SnapshotsLabelsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**snapshotId** | **string** | The unique ID of the snapshot. | |
|**key** | **string** | The label key | |

### Other Parameters

Other parameters are passed through a pointer to an apiSnapshotsLabelsDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## SnapshotsLabelsFindByKey

```go
var result LabelResource = SnapshotsLabelsFindByKey(ctx, snapshotId, key)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Retrieve snapshot labels



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    compute "github.com/ionos-cloud/sdk-go-bundle/products/compute"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    snapshotId := "snapshotId_example" // string | The unique ID of the snapshot.
    key := "key_example" // string | The label key
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.LabelsApi.SnapshotsLabelsFindByKey(context.Background(), snapshotId, key).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.SnapshotsLabelsFindByKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `SnapshotsLabelsFindByKey`: LabelResource
    fmt.Fprintf(os.Stdout, "Response from `LabelsApi.SnapshotsLabelsFindByKey`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**snapshotId** | **string** | The unique ID of the snapshot. | |
|**key** | **string** | The label key | |

### Other Parameters

Other parameters are passed through a pointer to an apiSnapshotsLabelsFindByKeyRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**LabelResource**](../models/LabelResource.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## SnapshotsLabelsGet

```go
var result LabelResources = SnapshotsLabelsGet(ctx, snapshotId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

List snapshot labels



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    compute "github.com/ionos-cloud/sdk-go-bundle/products/compute"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    snapshotId := "snapshotId_example" // string | The unique ID of the snapshot.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.LabelsApi.SnapshotsLabelsGet(context.Background(), snapshotId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.SnapshotsLabelsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `SnapshotsLabelsGet`: LabelResources
    fmt.Fprintf(os.Stdout, "Response from `LabelsApi.SnapshotsLabelsGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**snapshotId** | **string** | The unique ID of the snapshot. | |

### Other Parameters

Other parameters are passed through a pointer to an apiSnapshotsLabelsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**LabelResources**](../models/LabelResources.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## SnapshotsLabelsPost

```go
var result LabelResource = SnapshotsLabelsPost(ctx, snapshotId)
                      .Label(label)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Create a Snapshot Label



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    compute "github.com/ionos-cloud/sdk-go-bundle/products/compute"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    snapshotId := "snapshotId_example" // string | The unique ID of the snapshot.
    label := *openapiclient.NewLabelResource(*openapiclient.NewLabelResourceProperties()) // LabelResource | The label to create.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.LabelsApi.SnapshotsLabelsPost(context.Background(), snapshotId).Label(label).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.SnapshotsLabelsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `SnapshotsLabelsPost`: LabelResource
    fmt.Fprintf(os.Stdout, "Response from `LabelsApi.SnapshotsLabelsPost`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**snapshotId** | **string** | The unique ID of the snapshot. | |

### Other Parameters

Other parameters are passed through a pointer to an apiSnapshotsLabelsPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **label** | [**LabelResource**](../models/LabelResource.md) | The label to create. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**LabelResource**](../models/LabelResource.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## SnapshotsLabelsPut

```go
var result LabelResource = SnapshotsLabelsPut(ctx, snapshotId, key)
                      .Label(label)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Modify a Snapshot Label by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    compute "github.com/ionos-cloud/sdk-go-bundle/products/compute"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    snapshotId := "snapshotId_example" // string | The unique ID of the snapshot.
    key := "key_example" // string | The label key
    label := *openapiclient.NewLabelResource(*openapiclient.NewLabelResourceProperties()) // LabelResource | The modified label
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.LabelsApi.SnapshotsLabelsPut(context.Background(), snapshotId, key).Label(label).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsApi.SnapshotsLabelsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `SnapshotsLabelsPut`: LabelResource
    fmt.Fprintf(os.Stdout, "Response from `LabelsApi.SnapshotsLabelsPut`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**snapshotId** | **string** | The unique ID of the snapshot. | |
|**key** | **string** | The label key | |

### Other Parameters

Other parameters are passed through a pointer to an apiSnapshotsLabelsPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **label** | [**LabelResource**](../models/LabelResource.md) | The modified label | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**LabelResource**](../models/LabelResource.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


