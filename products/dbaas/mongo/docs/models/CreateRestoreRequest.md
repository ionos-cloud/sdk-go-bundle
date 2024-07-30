# CreateRestoreRequest

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**SnapshotId** | Pointer to **string** | The unique ID of the snapshot you want to restore. | [optional] |
|**RecoveryTargetTime** | Pointer to [**time.Time**](time.Time.md) | If this value is supplied as ISO 8601 timestamp, the backup will be replayed up until the given timestamp.  | [optional] |

## Methods

### NewCreateRestoreRequest

`func NewCreateRestoreRequest() *CreateRestoreRequest`

NewCreateRestoreRequest instantiates a new CreateRestoreRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRestoreRequestWithDefaults

`func NewCreateRestoreRequestWithDefaults() *CreateRestoreRequest`

NewCreateRestoreRequestWithDefaults instantiates a new CreateRestoreRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshotId

`func (o *CreateRestoreRequest) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *CreateRestoreRequest) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *CreateRestoreRequest) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *CreateRestoreRequest) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### GetRecoveryTargetTime

`func (o *CreateRestoreRequest) GetRecoveryTargetTime() time.Time`

GetRecoveryTargetTime returns the RecoveryTargetTime field if non-nil, zero value otherwise.

### GetRecoveryTargetTimeOk

`func (o *CreateRestoreRequest) GetRecoveryTargetTimeOk() (*time.Time, bool)`

GetRecoveryTargetTimeOk returns a tuple with the RecoveryTargetTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTargetTime

`func (o *CreateRestoreRequest) SetRecoveryTargetTime(v time.Time)`

SetRecoveryTargetTime sets RecoveryTargetTime field to given value.

### HasRecoveryTargetTime

`func (o *CreateRestoreRequest) HasRecoveryTargetTime() bool`

HasRecoveryTargetTime returns a boolean if a field has been set.


