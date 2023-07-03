# MetadataWithStateNameservers

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**LastModifiedDate** | Pointer to [**time.Time**](time.Time.md) | The date of the last change formatted as yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSS&#39;Z&#39;. | [optional] [readonly] |
|**CreatedDate** | Pointer to [**time.Time**](time.Time.md) | The date of creation of the zone formatted as yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSS&#39;Z&#39;. | [optional] [readonly] |
|**State** | [**ProvisioningState**](ProvisioningState.md) |  | |
|**Nameservers** | **[]string** | The list of nameservers associated to the zone | |

## Methods

### NewMetadataWithStateNameservers

`func NewMetadataWithStateNameservers(state ProvisioningState, nameservers []string, ) *MetadataWithStateNameservers`

NewMetadataWithStateNameservers instantiates a new MetadataWithStateNameservers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataWithStateNameserversWithDefaults

`func NewMetadataWithStateNameserversWithDefaults() *MetadataWithStateNameservers`

NewMetadataWithStateNameserversWithDefaults instantiates a new MetadataWithStateNameservers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastModifiedDate

`func (o *MetadataWithStateNameservers) GetLastModifiedDate() time.Time`

GetLastModifiedDate returns the LastModifiedDate field if non-nil, zero value otherwise.

### GetLastModifiedDateOk

`func (o *MetadataWithStateNameservers) GetLastModifiedDateOk() (*time.Time, bool)`

GetLastModifiedDateOk returns a tuple with the LastModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDate

`func (o *MetadataWithStateNameservers) SetLastModifiedDate(v time.Time)`

SetLastModifiedDate sets LastModifiedDate field to given value.

### HasLastModifiedDate

`func (o *MetadataWithStateNameservers) HasLastModifiedDate() bool`

HasLastModifiedDate returns a boolean if a field has been set.

### GetCreatedDate

`func (o *MetadataWithStateNameservers) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *MetadataWithStateNameservers) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *MetadataWithStateNameservers) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *MetadataWithStateNameservers) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetState

`func (o *MetadataWithStateNameservers) GetState() ProvisioningState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MetadataWithStateNameservers) GetStateOk() (*ProvisioningState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MetadataWithStateNameservers) SetState(v ProvisioningState)`

SetState sets State field to given value.


### GetNameservers

`func (o *MetadataWithStateNameservers) GetNameservers() []string`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *MetadataWithStateNameservers) GetNameserversOk() (*[]string, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *MetadataWithStateNameservers) SetNameservers(v []string)`

SetNameservers sets Nameservers field to given value.



