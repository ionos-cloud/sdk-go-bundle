# PatchRegistryInput

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**GarbageCollectionSchedule** | Pointer to [**NullableWeeklySchedule**](WeeklySchedule.md) |  | [optional] |

## Methods

### NewPatchRegistryInput

`func NewPatchRegistryInput() *PatchRegistryInput`

NewPatchRegistryInput instantiates a new PatchRegistryInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchRegistryInputWithDefaults

`func NewPatchRegistryInputWithDefaults() *PatchRegistryInput`

NewPatchRegistryInputWithDefaults instantiates a new PatchRegistryInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGarbageCollectionSchedule

`func (o *PatchRegistryInput) GetGarbageCollectionSchedule() WeeklySchedule`

GetGarbageCollectionSchedule returns the GarbageCollectionSchedule field if non-nil, zero value otherwise.

### GetGarbageCollectionScheduleOk

`func (o *PatchRegistryInput) GetGarbageCollectionScheduleOk() (*WeeklySchedule, bool)`

GetGarbageCollectionScheduleOk returns a tuple with the GarbageCollectionSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGarbageCollectionSchedule

`func (o *PatchRegistryInput) SetGarbageCollectionSchedule(v WeeklySchedule)`

SetGarbageCollectionSchedule sets GarbageCollectionSchedule field to given value.

### HasGarbageCollectionSchedule

`func (o *PatchRegistryInput) HasGarbageCollectionSchedule() bool`

HasGarbageCollectionSchedule returns a boolean if a field has been set.

### SetGarbageCollectionScheduleNil

`func (o *PatchRegistryInput) SetGarbageCollectionScheduleNil(b bool)`

 SetGarbageCollectionScheduleNil sets the value for GarbageCollectionSchedule to be an explicit nil

### UnsetGarbageCollectionSchedule
`func (o *PatchRegistryInput) UnsetGarbageCollectionSchedule()`

UnsetGarbageCollectionSchedule ensures that no value is present for GarbageCollectionSchedule, not even an explicit nil

