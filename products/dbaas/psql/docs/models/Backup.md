# Backup

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**ClusterId** | Pointer to **string** | The unique identifier of the cluster this backup belongs to. | [optional] |
|**PostgresClusterVersion** | Pointer to **string** | The PostgreSQL version of the cluster at backup time. | [optional] |
|**IsActive** | Pointer to **bool** | Whether this is the currently active backup for the cluster. | [optional] |
|**EarliestRecoveryTargetTime** | Pointer to [**time.Time**](time.Time.md) | The earliest point in time to which the cluster can be restored from this backup. | [optional] |
|**LatestRecoveryTargetTime** | Pointer to [**time.Time**](time.Time.md) | The latest possible point in time to which the cluster can be restored. If the backup can be restored up to the current time, this field will be null.  | [optional] |
|**Location** | Pointer to **string** | The Object Storage location where the backup will be created. The BackupLocations provides a list of supported locations.  | [optional] |

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

### GetPostgresClusterVersion

`func (o *Backup) GetPostgresClusterVersion() string`

GetPostgresClusterVersion returns the PostgresClusterVersion field if non-nil, zero value otherwise.

### GetPostgresClusterVersionOk

`func (o *Backup) GetPostgresClusterVersionOk() (*string, bool)`

GetPostgresClusterVersionOk returns a tuple with the PostgresClusterVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgresClusterVersion

`func (o *Backup) SetPostgresClusterVersion(v string)`

SetPostgresClusterVersion sets PostgresClusterVersion field to given value.

### HasPostgresClusterVersion

`func (o *Backup) HasPostgresClusterVersion() bool`

HasPostgresClusterVersion returns a boolean if a field has been set.

### GetIsActive

`func (o *Backup) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *Backup) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *Backup) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *Backup) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

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

### GetLatestRecoveryTargetTime

`func (o *Backup) GetLatestRecoveryTargetTime() time.Time`

GetLatestRecoveryTargetTime returns the LatestRecoveryTargetTime field if non-nil, zero value otherwise.

### GetLatestRecoveryTargetTimeOk

`func (o *Backup) GetLatestRecoveryTargetTimeOk() (*time.Time, bool)`

GetLatestRecoveryTargetTimeOk returns a tuple with the LatestRecoveryTargetTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRecoveryTargetTime

`func (o *Backup) SetLatestRecoveryTargetTime(v time.Time)`

SetLatestRecoveryTargetTime sets LatestRecoveryTargetTime field to given value.

### HasLatestRecoveryTargetTime

`func (o *Backup) HasLatestRecoveryTargetTime() bool`

HasLatestRecoveryTargetTime returns a boolean if a field has been set.

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


