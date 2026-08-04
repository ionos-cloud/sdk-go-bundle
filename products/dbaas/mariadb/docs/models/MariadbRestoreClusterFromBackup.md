# MariadbRestoreClusterFromBackup

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**SourceBackupId** | Pointer to **string** | UUID of the backup to restore from. Required for restore on cluster creation; not valid for in-place restore, where the source is inferred from the cluster&#39;s own backups.  | [optional] |
|**RecoveryTargetDatetime** | Pointer to [**time.Time**](time.Time.md) | Providing this value as an ISO 8601 timestamp causes the system to replay the backups up to the specified time. Optional on cluster creation (the backup is applied in its entirety if omitted); required for in-place restore.  | [optional] |

## Methods

### NewMariadbRestoreClusterFromBackup

`func NewMariadbRestoreClusterFromBackup() *MariadbRestoreClusterFromBackup`

NewMariadbRestoreClusterFromBackup instantiates a new MariadbRestoreClusterFromBackup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMariadbRestoreClusterFromBackupWithDefaults

`func NewMariadbRestoreClusterFromBackupWithDefaults() *MariadbRestoreClusterFromBackup`

NewMariadbRestoreClusterFromBackupWithDefaults instantiates a new MariadbRestoreClusterFromBackup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceBackupId

`func (o *MariadbRestoreClusterFromBackup) GetSourceBackupId() string`

GetSourceBackupId returns the SourceBackupId field if non-nil, zero value otherwise.

### GetSourceBackupIdOk

`func (o *MariadbRestoreClusterFromBackup) GetSourceBackupIdOk() (*string, bool)`

GetSourceBackupIdOk returns a tuple with the SourceBackupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceBackupId

`func (o *MariadbRestoreClusterFromBackup) SetSourceBackupId(v string)`

SetSourceBackupId sets SourceBackupId field to given value.

### HasSourceBackupId

`func (o *MariadbRestoreClusterFromBackup) HasSourceBackupId() bool`

HasSourceBackupId returns a boolean if a field has been set.

### GetRecoveryTargetDatetime

`func (o *MariadbRestoreClusterFromBackup) GetRecoveryTargetDatetime() time.Time`

GetRecoveryTargetDatetime returns the RecoveryTargetDatetime field if non-nil, zero value otherwise.

### GetRecoveryTargetDatetimeOk

`func (o *MariadbRestoreClusterFromBackup) GetRecoveryTargetDatetimeOk() (*time.Time, bool)`

GetRecoveryTargetDatetimeOk returns a tuple with the RecoveryTargetDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTargetDatetime

`func (o *MariadbRestoreClusterFromBackup) SetRecoveryTargetDatetime(v time.Time)`

SetRecoveryTargetDatetime sets RecoveryTargetDatetime field to given value.

### HasRecoveryTargetDatetime

`func (o *MariadbRestoreClusterFromBackup) HasRecoveryTargetDatetime() bool`

HasRecoveryTargetDatetime returns a boolean if a field has been set.


