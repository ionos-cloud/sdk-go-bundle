# ClusterBackup

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Location** | **string** | The Object Storage location where the backup will be created. The BackupLocations provides a list of supported locations.  | |
|**RetentionDays** | **int32** | Configures how many days cluster backups are retained. | |

## Methods

### NewClusterBackup

`func NewClusterBackup(location string, retentionDays int32, ) *ClusterBackup`

NewClusterBackup instantiates a new ClusterBackup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterBackupWithDefaults

`func NewClusterBackupWithDefaults() *ClusterBackup`

NewClusterBackupWithDefaults instantiates a new ClusterBackup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetRetentionDays

`func (o *ClusterBackup) GetRetentionDays() int32`

GetRetentionDays returns the RetentionDays field if non-nil, zero value otherwise.

### GetRetentionDaysOk

`func (o *ClusterBackup) GetRetentionDaysOk() (*int32, bool)`

GetRetentionDaysOk returns a tuple with the RetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionDays

`func (o *ClusterBackup) SetRetentionDays(v int32)`

SetRetentionDays sets RetentionDays field to given value.



