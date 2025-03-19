# Region

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Version** | **int32** | The version of the region properties | |
|**Endpoint** | **string** | The endpoint URL for the region | |
|**Website** | **string** | The website URL for the region | |
|**Capability** | [**RegionCapability**](RegionCapability.md) |  | |
|**StorageClasses** | Pointer to **[]string** | The available classes in the region | [optional] |
|**Location** | **string** | The data center location of the region as per [Get Location](/docs/cloud/v6/#tag/Locations/operation/locationsGet). *Can&#39;t be used as &#x60;LocationConstraint&#x60; on bucket creation.*  | |

## Methods

### NewRegion

`func NewRegion(version int32, endpoint string, website string, capability RegionCapability, location string, ) *Region`

NewRegion instantiates a new Region object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionWithDefaults

`func NewRegionWithDefaults() *Region`

NewRegionWithDefaults instantiates a new Region object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *Region) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Region) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Region) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetEndpoint

`func (o *Region) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *Region) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *Region) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetWebsite

`func (o *Region) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *Region) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *Region) SetWebsite(v string)`

SetWebsite sets Website field to given value.


### GetCapability

`func (o *Region) GetCapability() RegionCapability`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *Region) GetCapabilityOk() (*RegionCapability, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *Region) SetCapability(v RegionCapability)`

SetCapability sets Capability field to given value.


### GetStorageClasses

`func (o *Region) GetStorageClasses() []string`

GetStorageClasses returns the StorageClasses field if non-nil, zero value otherwise.

### GetStorageClassesOk

`func (o *Region) GetStorageClassesOk() (*[]string, bool)`

GetStorageClassesOk returns a tuple with the StorageClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClasses

`func (o *Region) SetStorageClasses(v []string)`

SetStorageClasses sets StorageClasses field to given value.

### HasStorageClasses

`func (o *Region) HasStorageClasses() bool`

HasStorageClasses returns a boolean if a field has been set.

### GetLocation

`func (o *Region) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Region) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Region) SetLocation(v string)`

SetLocation sets Location field to given value.



