# MetadataWithStateNameserversAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**State** | [**ProvisioningState**](ProvisioningState.md) |  | |
|**Nameservers** | **[]string** | The list of nameservers associated to the zone | |

## Methods

### NewMetadataWithStateNameserversAllOf

`func NewMetadataWithStateNameserversAllOf(state ProvisioningState, nameservers []string, ) *MetadataWithStateNameserversAllOf`

NewMetadataWithStateNameserversAllOf instantiates a new MetadataWithStateNameserversAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataWithStateNameserversAllOfWithDefaults

`func NewMetadataWithStateNameserversAllOfWithDefaults() *MetadataWithStateNameserversAllOf`

NewMetadataWithStateNameserversAllOfWithDefaults instantiates a new MetadataWithStateNameserversAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *MetadataWithStateNameserversAllOf) GetState() ProvisioningState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MetadataWithStateNameserversAllOf) GetStateOk() (*ProvisioningState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MetadataWithStateNameserversAllOf) SetState(v ProvisioningState)`

SetState sets State field to given value.


### GetNameservers

`func (o *MetadataWithStateNameserversAllOf) GetNameservers() []string`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *MetadataWithStateNameserversAllOf) GetNameserversOk() (*[]string, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *MetadataWithStateNameserversAllOf) SetNameservers(v []string)`

SetNameservers sets Nameservers field to given value.



