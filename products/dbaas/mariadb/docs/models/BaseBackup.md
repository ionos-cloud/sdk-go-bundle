# BaseBackup

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Created** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 creation timestamp. | [optional] |
|**Size** | Pointer to **int32** | The size of the backup in MiB. This is the size of the binary backup file that was stored.  | [optional] |

## Methods

### NewBaseBackup

`func NewBaseBackup() *BaseBackup`

NewBaseBackup instantiates a new BaseBackup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseBackupWithDefaults

`func NewBaseBackupWithDefaults() *BaseBackup`

NewBaseBackupWithDefaults instantiates a new BaseBackup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *BaseBackup) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *BaseBackup) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *BaseBackup) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *BaseBackup) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetSize

`func (o *BaseBackup) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BaseBackup) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BaseBackup) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *BaseBackup) HasSize() bool`

HasSize returns a boolean if a field has been set.


