# StorageClass

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Description** | **string** | Explains the motivation for the storage class | |
|**Durability** | **string** | The durability of the storage class | |
|**Availability** | **string** | The availability of the storage class | |

## Methods

### NewStorageClass

`func NewStorageClass(description string, durability string, availability string, ) *StorageClass`

NewStorageClass instantiates a new StorageClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageClassWithDefaults

`func NewStorageClassWithDefaults() *StorageClass`

NewStorageClassWithDefaults instantiates a new StorageClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *StorageClass) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StorageClass) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StorageClass) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDurability

`func (o *StorageClass) GetDurability() string`

GetDurability returns the Durability field if non-nil, zero value otherwise.

### GetDurabilityOk

`func (o *StorageClass) GetDurabilityOk() (*string, bool)`

GetDurabilityOk returns a tuple with the Durability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurability

`func (o *StorageClass) SetDurability(v string)`

SetDurability sets Durability field to given value.


### GetAvailability

`func (o *StorageClass) GetAvailability() string`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *StorageClass) GetAvailabilityOk() (*string, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *StorageClass) SetAvailability(v string)`

SetAvailability sets Availability field to given value.



