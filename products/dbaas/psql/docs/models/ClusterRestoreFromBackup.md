# ClusterRestoreFromBackup

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**SourceBackupId** | **string** | UUID for the backup to get data from. | |
|**RecoveryTargetDatetime** | [**time.Time**](time.Time.md) | Providing this value as an ISO 8601 timestamp causes the system to restore the cluster up to the specified time.  | |

## Methods

### NewClusterRestoreFromBackup

`func NewClusterRestoreFromBackup(sourceBackupId string, recoveryTargetDatetime time.Time, ) *ClusterRestoreFromBackup`

NewClusterRestoreFromBackup instantiates a new ClusterRestoreFromBackup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterRestoreFromBackupWithDefaults

`func NewClusterRestoreFromBackupWithDefaults() *ClusterRestoreFromBackup`

NewClusterRestoreFromBackupWithDefaults instantiates a new ClusterRestoreFromBackup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceBackupId

`func (o *ClusterRestoreFromBackup) GetSourceBackupId() string`

GetSourceBackupId returns the SourceBackupId field if non-nil, zero value otherwise.

### GetSourceBackupIdOk

`func (o *ClusterRestoreFromBackup) GetSourceBackupIdOk() (*string, bool)`

GetSourceBackupIdOk returns a tuple with the SourceBackupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceBackupId

`func (o *ClusterRestoreFromBackup) SetSourceBackupId(v string)`

SetSourceBackupId sets SourceBackupId field to given value.


### GetRecoveryTargetDatetime

`func (o *ClusterRestoreFromBackup) GetRecoveryTargetDatetime() time.Time`

GetRecoveryTargetDatetime returns the RecoveryTargetDatetime field if non-nil, zero value otherwise.

### GetRecoveryTargetDatetimeOk

`func (o *ClusterRestoreFromBackup) GetRecoveryTargetDatetimeOk() (*time.Time, bool)`

GetRecoveryTargetDatetimeOk returns a tuple with the RecoveryTargetDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTargetDatetime

`func (o *ClusterRestoreFromBackup) SetRecoveryTargetDatetime(v time.Time)`

SetRecoveryTargetDatetime sets RecoveryTargetDatetime field to given value.



