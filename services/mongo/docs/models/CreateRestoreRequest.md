# CreateRestoreRequest

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**SnapshotId** | **string** | The unique ID of the snapshot you want to restore. | |

## Methods

### NewCreateRestoreRequest

`func NewCreateRestoreRequest(snapshotId string, ) *CreateRestoreRequest`

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



