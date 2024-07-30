# TemplateProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | Pointer to **string** | The name of the template. | [optional] |
|**Edition** | Pointer to **string** | The edition of the template (e.g. enterprise) | [optional] |
|**Cores** | Pointer to **int32** | The number of CPU cores. | [optional] |
|**Ram** | Pointer to **int32** | The amount of memory in MB. | [optional] |
|**StorageSize** | Pointer to **int32** | The amount of storage size in GB. | [optional] |

## Methods

### NewTemplateProperties

`func NewTemplateProperties() *TemplateProperties`

NewTemplateProperties instantiates a new TemplateProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplatePropertiesWithDefaults

`func NewTemplatePropertiesWithDefaults() *TemplateProperties`

NewTemplatePropertiesWithDefaults instantiates a new TemplateProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TemplateProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateProperties) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TemplateProperties) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEdition

`func (o *TemplateProperties) GetEdition() string`

GetEdition returns the Edition field if non-nil, zero value otherwise.

### GetEditionOk

`func (o *TemplateProperties) GetEditionOk() (*string, bool)`

GetEditionOk returns a tuple with the Edition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdition

`func (o *TemplateProperties) SetEdition(v string)`

SetEdition sets Edition field to given value.

### HasEdition

`func (o *TemplateProperties) HasEdition() bool`

HasEdition returns a boolean if a field has been set.

### GetCores

`func (o *TemplateProperties) GetCores() int32`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *TemplateProperties) GetCoresOk() (*int32, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *TemplateProperties) SetCores(v int32)`

SetCores sets Cores field to given value.

### HasCores

`func (o *TemplateProperties) HasCores() bool`

HasCores returns a boolean if a field has been set.

### GetRam

`func (o *TemplateProperties) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *TemplateProperties) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *TemplateProperties) SetRam(v int32)`

SetRam sets Ram field to given value.

### HasRam

`func (o *TemplateProperties) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetStorageSize

`func (o *TemplateProperties) GetStorageSize() int32`

GetStorageSize returns the StorageSize field if non-nil, zero value otherwise.

### GetStorageSizeOk

`func (o *TemplateProperties) GetStorageSizeOk() (*int32, bool)`

GetStorageSizeOk returns a tuple with the StorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSize

`func (o *TemplateProperties) SetStorageSize(v int32)`

SetStorageSize sets StorageSize field to given value.

### HasStorageSize

`func (o *TemplateProperties) HasStorageSize() bool`

HasStorageSize returns a boolean if a field has been set.


