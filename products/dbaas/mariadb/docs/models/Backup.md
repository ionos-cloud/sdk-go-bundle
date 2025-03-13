# Backup

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**ClusterId** | Pointer to **string** | The unique ID of the cluster that was backed up. | [optional] |
|**Location** | Pointer to **string** | The S3 location where the backups will be stored. | [optional] |
|**EarliestRecoveryTargetTime** | Pointer to [**time.Time**](time.Time.md) | The oldest available timestamp to which you can restore. | [optional] |
|**Size** | Pointer to **int32** | Size of all base backups in MiB. This is at least the sum of all base backup sizes.  | [optional] |
|**BaseBackups** | Pointer to [**[]BaseBackup**](BaseBackup.md) |  | [optional] |

## Methods

### NewBackup

`func NewBackup() *Backup`

NewBackup instantiates a new Backup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupWithDefaults

`func NewBackupWithDefaults() *Backup`

NewBackupWithDefaults instantiates a new Backup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *Backup) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *Backup) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *Backup) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *Backup) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetLocation

`func (o *Backup) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Backup) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Backup) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Backup) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetEarliestRecoveryTargetTime

`func (o *Backup) GetEarliestRecoveryTargetTime() time.Time`

GetEarliestRecoveryTargetTime returns the EarliestRecoveryTargetTime field if non-nil, zero value otherwise.

### GetEarliestRecoveryTargetTimeOk

`func (o *Backup) GetEarliestRecoveryTargetTimeOk() (*time.Time, bool)`

GetEarliestRecoveryTargetTimeOk returns a tuple with the EarliestRecoveryTargetTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestRecoveryTargetTime

`func (o *Backup) SetEarliestRecoveryTargetTime(v time.Time)`

SetEarliestRecoveryTargetTime sets EarliestRecoveryTargetTime field to given value.

### HasEarliestRecoveryTargetTime

`func (o *Backup) HasEarliestRecoveryTargetTime() bool`

HasEarliestRecoveryTargetTime returns a boolean if a field has been set.

### GetSize

`func (o *Backup) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Backup) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Backup) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *Backup) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetBaseBackups

`func (o *Backup) GetBaseBackups() []BaseBackup`

GetBaseBackups returns the BaseBackups field if non-nil, zero value otherwise.

### GetBaseBackupsOk

`func (o *Backup) GetBaseBackupsOk() (*[]BaseBackup, bool)`

GetBaseBackupsOk returns a tuple with the BaseBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseBackups

`func (o *Backup) SetBaseBackups(v []BaseBackup)`

SetBaseBackups sets BaseBackups field to given value.

### HasBaseBackups

`func (o *Backup) HasBaseBackups() bool`

HasBaseBackups returns a boolean if a field has been set.


