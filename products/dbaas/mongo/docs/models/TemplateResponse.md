# TemplateResponse

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The unique template ID. | [optional] |
|**Name** | Pointer to **string** | The name of the template. | [optional] |
|**Edition** | Pointer to **string** | The edition of the template (e.g. enterprise) | [optional] |
|**Cores** | Pointer to **int32** | The number of CPU cores. | [optional] |
|**Ram** | Pointer to **int32** | The amount of memory in GB. | [optional] |
|**StorageSize** | Pointer to **int32** | The amount of storage size in GB. | [optional] |

## Methods

### NewTemplateResponse

`func NewTemplateResponse() *TemplateResponse`

NewTemplateResponse instantiates a new TemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateResponseWithDefaults

`func NewTemplateResponseWithDefaults() *TemplateResponse`

NewTemplateResponseWithDefaults instantiates a new TemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplateResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TemplateResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TemplateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TemplateResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEdition

`func (o *TemplateResponse) GetEdition() string`

GetEdition returns the Edition field if non-nil, zero value otherwise.

### GetEditionOk

`func (o *TemplateResponse) GetEditionOk() (*string, bool)`

GetEditionOk returns a tuple with the Edition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdition

`func (o *TemplateResponse) SetEdition(v string)`

SetEdition sets Edition field to given value.

### HasEdition

`func (o *TemplateResponse) HasEdition() bool`

HasEdition returns a boolean if a field has been set.

### GetCores

`func (o *TemplateResponse) GetCores() int32`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *TemplateResponse) GetCoresOk() (*int32, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *TemplateResponse) SetCores(v int32)`

SetCores sets Cores field to given value.

### HasCores

`func (o *TemplateResponse) HasCores() bool`

HasCores returns a boolean if a field has been set.

### GetRam

`func (o *TemplateResponse) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *TemplateResponse) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *TemplateResponse) SetRam(v int32)`

SetRam sets Ram field to given value.

### HasRam

`func (o *TemplateResponse) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetStorageSize

`func (o *TemplateResponse) GetStorageSize() int32`

GetStorageSize returns the StorageSize field if non-nil, zero value otherwise.

### GetStorageSizeOk

`func (o *TemplateResponse) GetStorageSizeOk() (*int32, bool)`

GetStorageSizeOk returns a tuple with the StorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSize

`func (o *TemplateResponse) SetStorageSize(v int32)`

SetStorageSize sets StorageSize field to given value.

### HasStorageSize

`func (o *TemplateResponse) HasStorageSize() bool`

HasStorageSize returns a boolean if a field has been set.


