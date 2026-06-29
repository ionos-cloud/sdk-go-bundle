# InstanceConfiguration

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Count** | **int32** | The total number of instances in the cluster. A value of 1 creates a standalone instance. Values 2-5 create a replicated setup with one primary and n-1 passive secondaries. Replication is asynchronous.  | |
|**Cores** | **int32** | The number of dedicated CPU cores per instance. | |
|**Ram** | **int32** | The amount of RAM per instance in gigabytes (GB). RAM cannot be downgraded. Because storage size is automatically derived from RAM, reducing RAM would require shrinking the disk, which is not supported.  | |

## Methods

### NewInstanceConfiguration

`func NewInstanceConfiguration(count int32, cores int32, ram int32, ) *InstanceConfiguration`

NewInstanceConfiguration instantiates a new InstanceConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceConfigurationWithDefaults

`func NewInstanceConfigurationWithDefaults() *InstanceConfiguration`

NewInstanceConfigurationWithDefaults instantiates a new InstanceConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *InstanceConfiguration) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *InstanceConfiguration) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *InstanceConfiguration) SetCount(v int32)`

SetCount sets Count field to given value.


### GetCores

`func (o *InstanceConfiguration) GetCores() int32`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *InstanceConfiguration) GetCoresOk() (*int32, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *InstanceConfiguration) SetCores(v int32)`

SetCores sets Cores field to given value.


### GetRam

`func (o *InstanceConfiguration) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *InstanceConfiguration) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *InstanceConfiguration) SetRam(v int32)`

SetRam sets Ram field to given value.



