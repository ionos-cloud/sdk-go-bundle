# MaintenanceWindow

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Time** | **string** | Start of the maintenance window in UTC time. | |
|**DayOfTheWeek** | [**DayOfTheWeek**](DayOfTheWeek.md) |  | |

## Methods

### NewMaintenanceWindow

`func NewMaintenanceWindow(time string, dayOfTheWeek DayOfTheWeek, ) *MaintenanceWindow`

NewMaintenanceWindow instantiates a new MaintenanceWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintenanceWindowWithDefaults

`func NewMaintenanceWindowWithDefaults() *MaintenanceWindow`

NewMaintenanceWindowWithDefaults instantiates a new MaintenanceWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *MaintenanceWindow) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *MaintenanceWindow) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *MaintenanceWindow) SetTime(v string)`

SetTime sets Time field to given value.


### GetDayOfTheWeek

`func (o *MaintenanceWindow) GetDayOfTheWeek() DayOfTheWeek`

GetDayOfTheWeek returns the DayOfTheWeek field if non-nil, zero value otherwise.

### GetDayOfTheWeekOk

`func (o *MaintenanceWindow) GetDayOfTheWeekOk() (*DayOfTheWeek, bool)`

GetDayOfTheWeekOk returns a tuple with the DayOfTheWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfTheWeek

`func (o *MaintenanceWindow) SetDayOfTheWeek(v DayOfTheWeek)`

SetDayOfTheWeek sets DayOfTheWeek field to given value.



