# RegionRead

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The Region of the Region. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the Region. | |
|**Metadata** | **map[string]interface{}** |  | |
|**Properties** | [**Region**](Region.md) |  | |

## Methods

### NewRegionRead

`func NewRegionRead(id string, type_ string, href string, metadata map[string]interface{}, properties Region, ) *RegionRead`

NewRegionRead instantiates a new RegionRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionReadWithDefaults

`func NewRegionReadWithDefaults() *RegionRead`

NewRegionReadWithDefaults instantiates a new RegionRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RegionRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegionRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegionRead) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *RegionRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RegionRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RegionRead) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *RegionRead) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RegionRead) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RegionRead) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *RegionRead) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RegionRead) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RegionRead) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *RegionRead) GetProperties() Region`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *RegionRead) GetPropertiesOk() (*Region, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *RegionRead) SetProperties(v Region)`

SetProperties sets Properties field to given value.



