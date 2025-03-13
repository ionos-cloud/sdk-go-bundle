# RestoreRequest

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**BackupId** | **string** | The unique ID of the resource. | |
|**RecoveryTargetTime** | Pointer to [**time.Time**](time.Time.md) | The timestamp to which the cluster should be restored. If empty, the backup will be applied to the latest timestamp.  This value must be supplied as ISO 8601 timestamp, the backup will be replayed up until the given timestamp. If empty, the backup will be applied completely.  Must be within the earliestRecoveryTargetTime and now.  The earliestRecoveryTargetTime can be looked up in the backup object.  | [optional] |

## Methods

### NewRestoreRequest

`func NewRestoreRequest(backupId string, ) *RestoreRequest`

NewRestoreRequest instantiates a new RestoreRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreRequestWithDefaults

`func NewRestoreRequestWithDefaults() *RestoreRequest`

NewRestoreRequestWithDefaults instantiates a new RestoreRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupId

`func (o *RestoreRequest) GetBackupId() string`

GetBackupId returns the BackupId field if non-nil, zero value otherwise.

### GetBackupIdOk

`func (o *RestoreRequest) GetBackupIdOk() (*string, bool)`

GetBackupIdOk returns a tuple with the BackupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupId

`func (o *RestoreRequest) SetBackupId(v string)`

SetBackupId sets BackupId field to given value.


### GetRecoveryTargetTime

`func (o *RestoreRequest) GetRecoveryTargetTime() time.Time`

GetRecoveryTargetTime returns the RecoveryTargetTime field if non-nil, zero value otherwise.

### GetRecoveryTargetTimeOk

`func (o *RestoreRequest) GetRecoveryTargetTimeOk() (*time.Time, bool)`

GetRecoveryTargetTimeOk returns a tuple with the RecoveryTargetTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTargetTime

`func (o *RestoreRequest) SetRecoveryTargetTime(v time.Time)`

SetRecoveryTargetTime sets RecoveryTargetTime field to given value.

### HasRecoveryTargetTime

`func (o *RestoreRequest) HasRecoveryTargetTime() bool`

HasRecoveryTargetTime returns a boolean if a field has been set.


