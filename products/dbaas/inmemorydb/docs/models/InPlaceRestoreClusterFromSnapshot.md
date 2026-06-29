# InPlaceRestoreClusterFromSnapshot

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**RecoveryTargetDatetime** | [**time.Time**](time.Time.md) | Provide an ISO 8601 timestamp to restore from the most recent snapshot taken at or before that time. In-Memory DB does not provide continuous point-in-time recovery; the nearest preceding snapshot is used. This field is required for in-place restore.  | |

## Methods

### NewInPlaceRestoreClusterFromSnapshot

`func NewInPlaceRestoreClusterFromSnapshot(recoveryTargetDatetime time.Time, ) *InPlaceRestoreClusterFromSnapshot`

NewInPlaceRestoreClusterFromSnapshot instantiates a new InPlaceRestoreClusterFromSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInPlaceRestoreClusterFromSnapshotWithDefaults

`func NewInPlaceRestoreClusterFromSnapshotWithDefaults() *InPlaceRestoreClusterFromSnapshot`

NewInPlaceRestoreClusterFromSnapshotWithDefaults instantiates a new InPlaceRestoreClusterFromSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecoveryTargetDatetime

`func (o *InPlaceRestoreClusterFromSnapshot) GetRecoveryTargetDatetime() time.Time`

GetRecoveryTargetDatetime returns the RecoveryTargetDatetime field if non-nil, zero value otherwise.

### GetRecoveryTargetDatetimeOk

`func (o *InPlaceRestoreClusterFromSnapshot) GetRecoveryTargetDatetimeOk() (*time.Time, bool)`

GetRecoveryTargetDatetimeOk returns a tuple with the RecoveryTargetDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTargetDatetime

`func (o *InPlaceRestoreClusterFromSnapshot) SetRecoveryTargetDatetime(v time.Time)`

SetRecoveryTargetDatetime sets RecoveryTargetDatetime field to given value.



