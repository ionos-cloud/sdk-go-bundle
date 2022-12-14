# PostRegistryProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**GarbageCollectionSchedule** | Pointer to [**NullableWeeklySchedule**](WeeklySchedule.md) |  | [optional] |
|**Location** | **string** |  | |
|**Name** | **string** |  | |

## Methods

### NewPostRegistryProperties

`func NewPostRegistryProperties(location string, name string, ) *PostRegistryProperties`

NewPostRegistryProperties instantiates a new PostRegistryProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostRegistryPropertiesWithDefaults

`func NewPostRegistryPropertiesWithDefaults() *PostRegistryProperties`

NewPostRegistryPropertiesWithDefaults instantiates a new PostRegistryProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGarbageCollectionSchedule

`func (o *PostRegistryProperties) GetGarbageCollectionSchedule() WeeklySchedule`

GetGarbageCollectionSchedule returns the GarbageCollectionSchedule field if non-nil, zero value otherwise.

### GetGarbageCollectionScheduleOk

`func (o *PostRegistryProperties) GetGarbageCollectionScheduleOk() (*WeeklySchedule, bool)`

GetGarbageCollectionScheduleOk returns a tuple with the GarbageCollectionSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGarbageCollectionSchedule

`func (o *PostRegistryProperties) SetGarbageCollectionSchedule(v WeeklySchedule)`

SetGarbageCollectionSchedule sets GarbageCollectionSchedule field to given value.

### HasGarbageCollectionSchedule

`func (o *PostRegistryProperties) HasGarbageCollectionSchedule() bool`

HasGarbageCollectionSchedule returns a boolean if a field has been set.

### SetGarbageCollectionScheduleNil

`func (o *PostRegistryProperties) SetGarbageCollectionScheduleNil(b bool)`

 SetGarbageCollectionScheduleNil sets the value for GarbageCollectionSchedule to be an explicit nil

### UnsetGarbageCollectionSchedule
`func (o *PostRegistryProperties) UnsetGarbageCollectionSchedule()`

UnsetGarbageCollectionSchedule ensures that no value is present for GarbageCollectionSchedule, not even an explicit nil
### GetLocation

`func (o *PostRegistryProperties) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PostRegistryProperties) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PostRegistryProperties) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetName

`func (o *PostRegistryProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostRegistryProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostRegistryProperties) SetName(v string)`

SetName sets Name field to given value.



