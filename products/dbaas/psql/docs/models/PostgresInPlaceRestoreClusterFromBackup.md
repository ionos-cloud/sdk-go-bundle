# PostgresInPlaceRestoreClusterFromBackup

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**RecoveryTargetDatetime** | [**time.Time**](time.Time.md) | Providing this value as an ISO 8601 timestamp causes the system to restore the cluster up to the specified time.  | |

## Methods

### NewPostgresInPlaceRestoreClusterFromBackup

`func NewPostgresInPlaceRestoreClusterFromBackup(recoveryTargetDatetime time.Time, ) *PostgresInPlaceRestoreClusterFromBackup`

NewPostgresInPlaceRestoreClusterFromBackup instantiates a new PostgresInPlaceRestoreClusterFromBackup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresInPlaceRestoreClusterFromBackupWithDefaults

`func NewPostgresInPlaceRestoreClusterFromBackupWithDefaults() *PostgresInPlaceRestoreClusterFromBackup`

NewPostgresInPlaceRestoreClusterFromBackupWithDefaults instantiates a new PostgresInPlaceRestoreClusterFromBackup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecoveryTargetDatetime

`func (o *PostgresInPlaceRestoreClusterFromBackup) GetRecoveryTargetDatetime() time.Time`

GetRecoveryTargetDatetime returns the RecoveryTargetDatetime field if non-nil, zero value otherwise.

### GetRecoveryTargetDatetimeOk

`func (o *PostgresInPlaceRestoreClusterFromBackup) GetRecoveryTargetDatetimeOk() (*time.Time, bool)`

GetRecoveryTargetDatetimeOk returns a tuple with the RecoveryTargetDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTargetDatetime

`func (o *PostgresInPlaceRestoreClusterFromBackup) SetRecoveryTargetDatetime(v time.Time)`

SetRecoveryTargetDatetime sets RecoveryTargetDatetime field to given value.



