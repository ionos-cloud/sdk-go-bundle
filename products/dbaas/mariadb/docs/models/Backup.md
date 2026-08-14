# Backup

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**ClusterId** | Pointer to **string** | The unique identifier of the cluster this backup belongs to. | [optional] |
|**ClusterName** | Pointer to **string** | The name of the MariaDB cluster this backup belongs to. | [optional] [readonly] |
|**MariadbClusterVersion** | Pointer to **string** | The MariaDB version of the cluster at backup time. | [optional] |
|**EarliestRecoveryTargetTime** | Pointer to [**time.Time**](time.Time.md) | The earliest point in time to which the cluster can be restored from this backup. | [optional] |
|**LatestRecoveryTargetTime** | Pointer to [**NullableTime**](time.Time.md) | The latest possible point in time to which the cluster can be restored. If the backup can be restored up to the current time, this field will be null.  | [optional] |
|**Location** | Pointer to **string** | The Object Storage location where the backup will be created. The BackupLocations provides a list of supported locations.  | [optional] |

## Methods

### NewBackup

`func NewBackup() *Backup`

NewBackup instantiates a new Backup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupWithDefaults

`func NewBackupWithDefaults() *Backup`

NewBackupWithDefaults instantiates a new Backup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *Backup) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *Backup) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *Backup) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *Backup) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetClusterName

`func (o *Backup) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *Backup) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *Backup) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *Backup) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetMariadbClusterVersion

`func (o *Backup) GetMariadbClusterVersion() string`

GetMariadbClusterVersion returns the MariadbClusterVersion field if non-nil, zero value otherwise.

### GetMariadbClusterVersionOk

`func (o *Backup) GetMariadbClusterVersionOk() (*string, bool)`

GetMariadbClusterVersionOk returns a tuple with the MariadbClusterVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMariadbClusterVersion

`func (o *Backup) SetMariadbClusterVersion(v string)`

SetMariadbClusterVersion sets MariadbClusterVersion field to given value.

### HasMariadbClusterVersion

`func (o *Backup) HasMariadbClusterVersion() bool`

HasMariadbClusterVersion returns a boolean if a field has been set.

### GetEarliestRecoveryTargetTime

`func (o *Backup) GetEarliestRecoveryTargetTime() time.Time`

GetEarliestRecoveryTargetTime returns the EarliestRecoveryTargetTime field if non-nil, zero value otherwise.

### GetEarliestRecoveryTargetTimeOk

`func (o *Backup) GetEarliestRecoveryTargetTimeOk() (*time.Time, bool)`

GetEarliestRecoveryTargetTimeOk returns a tuple with the EarliestRecoveryTargetTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestRecoveryTargetTime

`func (o *Backup) SetEarliestRecoveryTargetTime(v time.Time)`

SetEarliestRecoveryTargetTime sets EarliestRecoveryTargetTime field to given value.

### HasEarliestRecoveryTargetTime

`func (o *Backup) HasEarliestRecoveryTargetTime() bool`

HasEarliestRecoveryTargetTime returns a boolean if a field has been set.

### GetLatestRecoveryTargetTime

`func (o *Backup) GetLatestRecoveryTargetTime() time.Time`

GetLatestRecoveryTargetTime returns the LatestRecoveryTargetTime field if non-nil, zero value otherwise.

### GetLatestRecoveryTargetTimeOk

`func (o *Backup) GetLatestRecoveryTargetTimeOk() (*time.Time, bool)`

GetLatestRecoveryTargetTimeOk returns a tuple with the LatestRecoveryTargetTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRecoveryTargetTime

`func (o *Backup) SetLatestRecoveryTargetTime(v time.Time)`

SetLatestRecoveryTargetTime sets LatestRecoveryTargetTime field to given value.

### HasLatestRecoveryTargetTime

`func (o *Backup) HasLatestRecoveryTargetTime() bool`

HasLatestRecoveryTargetTime returns a boolean if a field has been set.

### SetLatestRecoveryTargetTimeNil

`func (o *Backup) SetLatestRecoveryTargetTimeNil(b bool)`

 SetLatestRecoveryTargetTimeNil sets the value for LatestRecoveryTargetTime to be an explicit nil

### UnsetLatestRecoveryTargetTime
`func (o *Backup) UnsetLatestRecoveryTargetTime()`

UnsetLatestRecoveryTargetTime ensures that no value is present for LatestRecoveryTargetTime, not even an explicit nil
### GetLocation

`func (o *Backup) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Backup) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Backup) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Backup) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


