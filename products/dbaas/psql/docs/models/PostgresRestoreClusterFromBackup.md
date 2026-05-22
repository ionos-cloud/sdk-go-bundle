# PostgresRestoreClusterFromBackup

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**SourceBackupId** | **string** | UUID for the backup to get data from. | |
|**RecoveryTargetDatetime** | Pointer to [**time.Time**](time.Time.md) | Providing this value as an ISO 8601 timestamp causes the system to replay the backups up to the specified time. If omitted on cluster creation, the system applies the backup in its entirety.  | [optional] |

## Methods

### NewPostgresRestoreClusterFromBackup

`func NewPostgresRestoreClusterFromBackup(sourceBackupId string, ) *PostgresRestoreClusterFromBackup`

NewPostgresRestoreClusterFromBackup instantiates a new PostgresRestoreClusterFromBackup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresRestoreClusterFromBackupWithDefaults

`func NewPostgresRestoreClusterFromBackupWithDefaults() *PostgresRestoreClusterFromBackup`

NewPostgresRestoreClusterFromBackupWithDefaults instantiates a new PostgresRestoreClusterFromBackup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceBackupId

`func (o *PostgresRestoreClusterFromBackup) GetSourceBackupId() string`

GetSourceBackupId returns the SourceBackupId field if non-nil, zero value otherwise.

### GetSourceBackupIdOk

`func (o *PostgresRestoreClusterFromBackup) GetSourceBackupIdOk() (*string, bool)`

GetSourceBackupIdOk returns a tuple with the SourceBackupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceBackupId

`func (o *PostgresRestoreClusterFromBackup) SetSourceBackupId(v string)`

SetSourceBackupId sets SourceBackupId field to given value.


### GetRecoveryTargetDatetime

`func (o *PostgresRestoreClusterFromBackup) GetRecoveryTargetDatetime() time.Time`

GetRecoveryTargetDatetime returns the RecoveryTargetDatetime field if non-nil, zero value otherwise.

### GetRecoveryTargetDatetimeOk

`func (o *PostgresRestoreClusterFromBackup) GetRecoveryTargetDatetimeOk() (*time.Time, bool)`

GetRecoveryTargetDatetimeOk returns a tuple with the RecoveryTargetDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTargetDatetime

`func (o *PostgresRestoreClusterFromBackup) SetRecoveryTargetDatetime(v time.Time)`

SetRecoveryTargetDatetime sets RecoveryTargetDatetime field to given value.

### HasRecoveryTargetDatetime

`func (o *PostgresRestoreClusterFromBackup) HasRecoveryTargetDatetime() bool`

HasRecoveryTargetDatetime returns a boolean if a field has been set.


