# RegistryProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**GarbageCollectionSchedule** | Pointer to [**NullableWeeklySchedule**](WeeklySchedule.md) |  | [optional] |
|**Hostname** | Pointer to **string** |  | [optional] |
|**Location** | **string** |  | |
|**Name** | **string** |  | |
|**StorageUsage** | Pointer to [**NullableStorageUsage**](StorageUsage.md) |  | [optional] |

## Methods

### NewRegistryProperties

`func NewRegistryProperties(location string, name string, ) *RegistryProperties`

NewRegistryProperties instantiates a new RegistryProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryPropertiesWithDefaults

`func NewRegistryPropertiesWithDefaults() *RegistryProperties`

NewRegistryPropertiesWithDefaults instantiates a new RegistryProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGarbageCollectionSchedule

`func (o *RegistryProperties) GetGarbageCollectionSchedule() WeeklySchedule`

GetGarbageCollectionSchedule returns the GarbageCollectionSchedule field if non-nil, zero value otherwise.

### GetGarbageCollectionScheduleOk

`func (o *RegistryProperties) GetGarbageCollectionScheduleOk() (*WeeklySchedule, bool)`

GetGarbageCollectionScheduleOk returns a tuple with the GarbageCollectionSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGarbageCollectionSchedule

`func (o *RegistryProperties) SetGarbageCollectionSchedule(v WeeklySchedule)`

SetGarbageCollectionSchedule sets GarbageCollectionSchedule field to given value.

### HasGarbageCollectionSchedule

`func (o *RegistryProperties) HasGarbageCollectionSchedule() bool`

HasGarbageCollectionSchedule returns a boolean if a field has been set.

### SetGarbageCollectionScheduleNil

`func (o *RegistryProperties) SetGarbageCollectionScheduleNil(b bool)`

 SetGarbageCollectionScheduleNil sets the value for GarbageCollectionSchedule to be an explicit nil

### UnsetGarbageCollectionSchedule
`func (o *RegistryProperties) UnsetGarbageCollectionSchedule()`

UnsetGarbageCollectionSchedule ensures that no value is present for GarbageCollectionSchedule, not even an explicit nil
### GetHostname

`func (o *RegistryProperties) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *RegistryProperties) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *RegistryProperties) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *RegistryProperties) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetLocation

`func (o *RegistryProperties) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *RegistryProperties) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *RegistryProperties) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetName

`func (o *RegistryProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegistryProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegistryProperties) SetName(v string)`

SetName sets Name field to given value.


### GetStorageUsage

`func (o *RegistryProperties) GetStorageUsage() StorageUsage`

GetStorageUsage returns the StorageUsage field if non-nil, zero value otherwise.

### GetStorageUsageOk

`func (o *RegistryProperties) GetStorageUsageOk() (*StorageUsage, bool)`

GetStorageUsageOk returns a tuple with the StorageUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUsage

`func (o *RegistryProperties) SetStorageUsage(v StorageUsage)`

SetStorageUsage sets StorageUsage field to given value.

### HasStorageUsage

`func (o *RegistryProperties) HasStorageUsage() bool`

HasStorageUsage returns a boolean if a field has been set.

### SetStorageUsageNil

`func (o *RegistryProperties) SetStorageUsageNil(b bool)`

 SetStorageUsageNil sets the value for StorageUsage to be an explicit nil

### UnsetStorageUsage
`func (o *RegistryProperties) UnsetStorageUsage()`

UnsetStorageUsage ensures that no value is present for StorageUsage, not even an explicit nil

