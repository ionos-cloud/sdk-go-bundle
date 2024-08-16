# DistributionUpdate

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The ID (UUID) of the Distribution. | |
|**Metadata** | Pointer to **map[string]interface{}** | Metadata | [optional] |
|**Properties** | [**DistributionProperties**](DistributionProperties.md) |  | |

## Methods

### NewDistributionUpdate

`func NewDistributionUpdate(id string, properties DistributionProperties, ) *DistributionUpdate`

NewDistributionUpdate instantiates a new DistributionUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDistributionUpdateWithDefaults

`func NewDistributionUpdateWithDefaults() *DistributionUpdate`

NewDistributionUpdateWithDefaults instantiates a new DistributionUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DistributionUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DistributionUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DistributionUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *DistributionUpdate) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DistributionUpdate) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DistributionUpdate) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DistributionUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *DistributionUpdate) GetProperties() DistributionProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *DistributionUpdate) GetPropertiesOk() (*DistributionProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *DistributionUpdate) SetProperties(v DistributionProperties)`

SetProperties sets Properties field to given value.



