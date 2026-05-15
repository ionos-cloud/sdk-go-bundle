# ClusterBackup

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The unique ID of the resource. | [optional] |
|**ClusterId** | Pointer to **string** | The unique ID of the cluster. | [optional] |
|**Version** | Pointer to **string** | The PostgreSQL version this backup was created from. | [optional] |
|**IsActive** | Pointer to **bool** | Whether a cluster currently backs up data to this backup. | [optional] |
|**EarliestRecoveryTargetTime** | Pointer to [**time.Time**](time.Time.md) | The oldest available timestamp to which you can restore. | [optional] |
|**Size** | Pointer to **int32** | Size of all base backups including the wal size in MB. | [optional] |
|**Location** | Pointer to **string** | The S3 location where the backups will be stored. | [optional] |

## Methods

### NewClusterBackup

`func NewClusterBackup() *ClusterBackup`

NewClusterBackup instantiates a new ClusterBackup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterBackupWithDefaults

`func NewClusterBackupWithDefaults() *ClusterBackup`

NewClusterBackupWithDefaults instantiates a new ClusterBackup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterBackup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterBackup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterBackup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterBackup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClusterId

`func (o *ClusterBackup) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ClusterBackup) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ClusterBackup) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *ClusterBackup) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetVersion

`func (o *ClusterBackup) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ClusterBackup) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ClusterBackup) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ClusterBackup) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetIsActive

`func (o *ClusterBackup) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ClusterBackup) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ClusterBackup) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ClusterBackup) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetEarliestRecoveryTargetTime

`func (o *ClusterBackup) GetEarliestRecoveryTargetTime() time.Time`

GetEarliestRecoveryTargetTime returns the EarliestRecoveryTargetTime field if non-nil, zero value otherwise.

### GetEarliestRecoveryTargetTimeOk

`func (o *ClusterBackup) GetEarliestRecoveryTargetTimeOk() (*time.Time, bool)`

GetEarliestRecoveryTargetTimeOk returns a tuple with the EarliestRecoveryTargetTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestRecoveryTargetTime

`func (o *ClusterBackup) SetEarliestRecoveryTargetTime(v time.Time)`

SetEarliestRecoveryTargetTime sets EarliestRecoveryTargetTime field to given value.

### HasEarliestRecoveryTargetTime

`func (o *ClusterBackup) HasEarliestRecoveryTargetTime() bool`

HasEarliestRecoveryTargetTime returns a boolean if a field has been set.

### GetSize

`func (o *ClusterBackup) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ClusterBackup) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ClusterBackup) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ClusterBackup) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetLocation

`func (o *ClusterBackup) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ClusterBackup) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ClusterBackup) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ClusterBackup) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


