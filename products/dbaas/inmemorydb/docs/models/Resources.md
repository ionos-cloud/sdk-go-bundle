# Resources

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Cores** | **int32** | The number of CPU cores per instance. | |
|**Ram** | **int32** | The amount of memory per instance in gigabytes (GB). | |
|**Storage** | Pointer to **int32** | The size of the storage in GB. The size is derived from the amount of RAM and the persistence mode and is not configurable.  | [optional] [readonly] |

## Methods

### NewResources

`func NewResources(cores int32, ram int32, ) *Resources`

NewResources instantiates a new Resources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcesWithDefaults

`func NewResourcesWithDefaults() *Resources`

NewResourcesWithDefaults instantiates a new Resources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCores

`func (o *Resources) GetCores() int32`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *Resources) GetCoresOk() (*int32, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *Resources) SetCores(v int32)`

SetCores sets Cores field to given value.


### GetRam

`func (o *Resources) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *Resources) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *Resources) SetRam(v int32)`

SetRam sets Ram field to given value.


### GetStorage

`func (o *Resources) GetStorage() int32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *Resources) GetStorageOk() (*int32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *Resources) SetStorage(v int32)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *Resources) HasStorage() bool`

HasStorage returns a boolean if a field has been set.


