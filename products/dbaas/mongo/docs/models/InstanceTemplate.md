# InstanceTemplate

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | Pointer to **string** | The name of the template. | [optional] |
|**Edition** | Pointer to [**MongoEdition**](MongoEdition.md) |  | [optional] |
|**Ram** | Pointer to **int32** | The amount of memory per instance in gigabytes (GB). | [optional] |
|**Cores** | Pointer to **int32** | The number of CPU cores per instance. | [optional] |
|**StorageSize** | Pointer to **int32** | The amount of storage per instance in gigabytes (GB). | [optional] |

## Methods

### NewInstanceTemplate

`func NewInstanceTemplate() *InstanceTemplate`

NewInstanceTemplate instantiates a new InstanceTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceTemplateWithDefaults

`func NewInstanceTemplateWithDefaults() *InstanceTemplate`

NewInstanceTemplateWithDefaults instantiates a new InstanceTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InstanceTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEdition

`func (o *InstanceTemplate) GetEdition() MongoEdition`

GetEdition returns the Edition field if non-nil, zero value otherwise.

### GetEditionOk

`func (o *InstanceTemplate) GetEditionOk() (*MongoEdition, bool)`

GetEditionOk returns a tuple with the Edition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdition

`func (o *InstanceTemplate) SetEdition(v MongoEdition)`

SetEdition sets Edition field to given value.

### HasEdition

`func (o *InstanceTemplate) HasEdition() bool`

HasEdition returns a boolean if a field has been set.

### GetRam

`func (o *InstanceTemplate) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *InstanceTemplate) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *InstanceTemplate) SetRam(v int32)`

SetRam sets Ram field to given value.

### HasRam

`func (o *InstanceTemplate) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetCores

`func (o *InstanceTemplate) GetCores() int32`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *InstanceTemplate) GetCoresOk() (*int32, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *InstanceTemplate) SetCores(v int32)`

SetCores sets Cores field to given value.

### HasCores

`func (o *InstanceTemplate) HasCores() bool`

HasCores returns a boolean if a field has been set.

### GetStorageSize

`func (o *InstanceTemplate) GetStorageSize() int32`

GetStorageSize returns the StorageSize field if non-nil, zero value otherwise.

### GetStorageSizeOk

`func (o *InstanceTemplate) GetStorageSizeOk() (*int32, bool)`

GetStorageSizeOk returns a tuple with the StorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSize

`func (o *InstanceTemplate) SetStorageSize(v int32)`

SetStorageSize sets StorageSize field to given value.

### HasStorageSize

`func (o *InstanceTemplate) HasStorageSize() bool`

HasStorageSize returns a boolean if a field has been set.


