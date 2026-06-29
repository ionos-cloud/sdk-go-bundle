# SnapshotConfiguration

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Location** | **string** | The Object Storage location where snapshots will be stored. For added data safety, use a different location than the cluster. A list of supported locations is provided by the SnapshotLocations endpoint.  | |
|**RetentionDays** | **int32** | The number of days snapshots are retained before being automatically deleted. Reducing this value causes the platform to purge any existing snapshots that fall outside the new retention window. Pre-existing clusters that were created before this field was introduced default to 7 days.  | [default to 7]|
|**SnapshotHours** | **[]int32** | Hours of the day (UTC) at which snapshots are scheduled to be taken. Each value must be between 0 and 23. At least one hour must be specified.  | |

## Methods

### NewSnapshotConfiguration

`func NewSnapshotConfiguration(location string, retentionDays int32, snapshotHours []int32, ) *SnapshotConfiguration`

NewSnapshotConfiguration instantiates a new SnapshotConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotConfigurationWithDefaults

`func NewSnapshotConfigurationWithDefaults() *SnapshotConfiguration`

NewSnapshotConfigurationWithDefaults instantiates a new SnapshotConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *SnapshotConfiguration) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SnapshotConfiguration) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SnapshotConfiguration) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetRetentionDays

`func (o *SnapshotConfiguration) GetRetentionDays() int32`

GetRetentionDays returns the RetentionDays field if non-nil, zero value otherwise.

### GetRetentionDaysOk

`func (o *SnapshotConfiguration) GetRetentionDaysOk() (*int32, bool)`

GetRetentionDaysOk returns a tuple with the RetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionDays

`func (o *SnapshotConfiguration) SetRetentionDays(v int32)`

SetRetentionDays sets RetentionDays field to given value.


### GetSnapshotHours

`func (o *SnapshotConfiguration) GetSnapshotHours() []int32`

GetSnapshotHours returns the SnapshotHours field if non-nil, zero value otherwise.

### GetSnapshotHoursOk

`func (o *SnapshotConfiguration) GetSnapshotHoursOk() (*[]int32, bool)`

GetSnapshotHoursOk returns a tuple with the SnapshotHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotHours

`func (o *SnapshotConfiguration) SetSnapshotHours(v []int32)`

SetSnapshotHours sets SnapshotHours field to given value.



