# WeeklySchedule

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Days** | [**[]Day**](Day.md) |  | |
|**Time** | **string** | UTC time of day e.g. 01:00:00 - as defined by partial-time - RFC3339 | |

## Methods

### NewWeeklySchedule

`func NewWeeklySchedule(days []Day, time string, ) *WeeklySchedule`

NewWeeklySchedule instantiates a new WeeklySchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWeeklyScheduleWithDefaults

`func NewWeeklyScheduleWithDefaults() *WeeklySchedule`

NewWeeklyScheduleWithDefaults instantiates a new WeeklySchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *WeeklySchedule) GetDays() []Day`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *WeeklySchedule) GetDaysOk() (*[]Day, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *WeeklySchedule) SetDays(v []Day)`

SetDays sets Days field to given value.


### GetTime

`func (o *WeeklySchedule) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *WeeklySchedule) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *WeeklySchedule) SetTime(v string)`

SetTime sets Time field to given value.



