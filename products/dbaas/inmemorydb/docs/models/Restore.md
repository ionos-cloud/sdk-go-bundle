# Restore

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**ReplicasetId** | **string** | The ID of the replica set the restore was applied on. | |
|**DisplayName** | Pointer to **string** | The human readable name of your snapshot. | [optional] |
|**Description** | Pointer to **string** | A description of the snapshot. | [optional] |

## Methods

### NewRestore

`func NewRestore(replicasetId string, ) *Restore`

NewRestore instantiates a new Restore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreWithDefaults

`func NewRestoreWithDefaults() *Restore`

NewRestoreWithDefaults instantiates a new Restore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReplicasetId

`func (o *Restore) GetReplicasetId() string`

GetReplicasetId returns the ReplicasetId field if non-nil, zero value otherwise.

### GetReplicasetIdOk

`func (o *Restore) GetReplicasetIdOk() (*string, bool)`

GetReplicasetIdOk returns a tuple with the ReplicasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicasetId

`func (o *Restore) SetReplicasetId(v string)`

SetReplicasetId sets ReplicasetId field to given value.


### GetDisplayName

`func (o *Restore) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Restore) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Restore) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Restore) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *Restore) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Restore) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Restore) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Restore) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


