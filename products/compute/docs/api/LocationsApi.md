# \LocationsApi

All URIs are relative to *https://api.ionos.com/cloudapi/v6*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**LocationsFindByRegionId**](LocationsApi.md#LocationsFindByRegionId) | **Get** /locations/{regionId} | Get Locations within a Region|
|[**LocationsFindByRegionIdAndId**](LocationsApi.md#LocationsFindByRegionIdAndId) | **Get** /locations/{regionId}/{locationId} | Get Location by ID|
|[**LocationsGet**](LocationsApi.md#LocationsGet) | **Get** /locations | Get Locations|



## LocationsFindByRegionId

```go
var result Locations = LocationsFindByRegionId(ctx, regionId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Get Locations within a Region



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
    regionId := "regionId_example" // string | The unique ID of the region.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.LocationsApi.LocationsFindByRegionId(context.Background(), regionId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationsApi.LocationsFindByRegionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `LocationsFindByRegionId`: Locations
    fmt.Fprintf(os.Stdout, "Response from `LocationsApi.LocationsFindByRegionId`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**regionId** | **string** | The unique ID of the region. | |

### Other Parameters

Other parameters are passed through a pointer to an apiLocationsFindByRegionIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**Locations**](../models/Locations.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## LocationsFindByRegionIdAndId

```go
var result Location = LocationsFindByRegionIdAndId(ctx, regionId, locationId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Get Location by ID



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
    regionId := "regionId_example" // string | The unique ID of the region.
    locationId := "locationId_example" // string | The unique ID of the location.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.LocationsApi.LocationsFindByRegionIdAndId(context.Background(), regionId, locationId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationsApi.LocationsFindByRegionIdAndId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `LocationsFindByRegionIdAndId`: Location
    fmt.Fprintf(os.Stdout, "Response from `LocationsApi.LocationsFindByRegionIdAndId`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**regionId** | **string** | The unique ID of the region. | |
|**locationId** | **string** | The unique ID of the location. | |

### Other Parameters

Other parameters are passed through a pointer to an apiLocationsFindByRegionIdAndIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**Location**](../models/Location.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## LocationsGet

```go
var result Locations = LocationsGet(ctx)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Get Locations



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
    resource, resp, err := apiClient.LocationsApi.LocationsGet(context.Background()).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationsApi.LocationsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `LocationsGet`: Locations
    fmt.Fprintf(os.Stdout, "Response from `LocationsApi.LocationsGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiLocationsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**Locations**](../models/Locations.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


