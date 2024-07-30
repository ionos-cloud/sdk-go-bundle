# BackupRetentionProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**SnapshotRetentionDays** | Pointer to **int32** | Number of days to keep recent snapshots.  | [optional] [default to 2]|
|**DailySnapshotRetentionDays** | Pointer to **int32** | Number of days to retain daily snapshots.  | [optional] [default to 0]|
|**WeeklySnapshotRetentionWeeks** | Pointer to **int32** | Number of weeks to retain weekly snapshots..  | [optional] [default to 2]|
|**MonthlySnapshotRetentionMonths** | Pointer to **int32** | Number of months to retain monthly snapshots.  | [optional] [default to 1]|

## Methods

### NewBackupRetentionProperties

`func NewBackupRetentionProperties() *BackupRetentionProperties`

NewBackupRetentionProperties instantiates a new BackupRetentionProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRetentionPropertiesWithDefaults

`func NewBackupRetentionPropertiesWithDefaults() *BackupRetentionProperties`

NewBackupRetentionPropertiesWithDefaults instantiates a new BackupRetentionProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshotRetentionDays

`func (o *BackupRetentionProperties) GetSnapshotRetentionDays() int32`

GetSnapshotRetentionDays returns the SnapshotRetentionDays field if non-nil, zero value otherwise.

### GetSnapshotRetentionDaysOk

`func (o *BackupRetentionProperties) GetSnapshotRetentionDaysOk() (*int32, bool)`

GetSnapshotRetentionDaysOk returns a tuple with the SnapshotRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotRetentionDays

`func (o *BackupRetentionProperties) SetSnapshotRetentionDays(v int32)`

SetSnapshotRetentionDays sets SnapshotRetentionDays field to given value.

### HasSnapshotRetentionDays

`func (o *BackupRetentionProperties) HasSnapshotRetentionDays() bool`

HasSnapshotRetentionDays returns a boolean if a field has been set.

### GetDailySnapshotRetentionDays

`func (o *BackupRetentionProperties) GetDailySnapshotRetentionDays() int32`

GetDailySnapshotRetentionDays returns the DailySnapshotRetentionDays field if non-nil, zero value otherwise.

### GetDailySnapshotRetentionDaysOk

`func (o *BackupRetentionProperties) GetDailySnapshotRetentionDaysOk() (*int32, bool)`

GetDailySnapshotRetentionDaysOk returns a tuple with the DailySnapshotRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailySnapshotRetentionDays

`func (o *BackupRetentionProperties) SetDailySnapshotRetentionDays(v int32)`

SetDailySnapshotRetentionDays sets DailySnapshotRetentionDays field to given value.

### HasDailySnapshotRetentionDays

`func (o *BackupRetentionProperties) HasDailySnapshotRetentionDays() bool`

HasDailySnapshotRetentionDays returns a boolean if a field has been set.

### GetWeeklySnapshotRetentionWeeks

`func (o *BackupRetentionProperties) GetWeeklySnapshotRetentionWeeks() int32`

GetWeeklySnapshotRetentionWeeks returns the WeeklySnapshotRetentionWeeks field if non-nil, zero value otherwise.

### GetWeeklySnapshotRetentionWeeksOk

`func (o *BackupRetentionProperties) GetWeeklySnapshotRetentionWeeksOk() (*int32, bool)`

GetWeeklySnapshotRetentionWeeksOk returns a tuple with the WeeklySnapshotRetentionWeeks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeklySnapshotRetentionWeeks

`func (o *BackupRetentionProperties) SetWeeklySnapshotRetentionWeeks(v int32)`

SetWeeklySnapshotRetentionWeeks sets WeeklySnapshotRetentionWeeks field to given value.

### HasWeeklySnapshotRetentionWeeks

`func (o *BackupRetentionProperties) HasWeeklySnapshotRetentionWeeks() bool`

HasWeeklySnapshotRetentionWeeks returns a boolean if a field has been set.

### GetMonthlySnapshotRetentionMonths

`func (o *BackupRetentionProperties) GetMonthlySnapshotRetentionMonths() int32`

GetMonthlySnapshotRetentionMonths returns the MonthlySnapshotRetentionMonths field if non-nil, zero value otherwise.

### GetMonthlySnapshotRetentionMonthsOk

`func (o *BackupRetentionProperties) GetMonthlySnapshotRetentionMonthsOk() (*int32, bool)`

GetMonthlySnapshotRetentionMonthsOk returns a tuple with the MonthlySnapshotRetentionMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlySnapshotRetentionMonths

`func (o *BackupRetentionProperties) SetMonthlySnapshotRetentionMonths(v int32)`

SetMonthlySnapshotRetentionMonths sets MonthlySnapshotRetentionMonths field to given value.

### HasMonthlySnapshotRetentionMonths

`func (o *BackupRetentionProperties) HasMonthlySnapshotRetentionMonths() bool`

HasMonthlySnapshotRetentionMonths returns a boolean if a field has been set.


