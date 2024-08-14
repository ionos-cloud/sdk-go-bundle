/*
 * IONOS DBaaS MongoDB REST API
 *
 * With IONOS Cloud Database as a Service, you have the ability to quickly set up and manage a MongoDB database. You can also delete clusters, manage backups and users via the API.  MongoDB is an open source, cross-platform, document-oriented database program. Classified as a NoSQL database program, it uses JSON-like documents with optional schemas.  The MongoDB API allows you to create additional database clusters or modify existing ones. Both tools, the Data Center Designer (DCD) and the API use the same concepts consistently and are well suited for smooth and intuitive use.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongo

import (
	"encoding/json"
)

// checks if the BackupRetentionProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupRetentionProperties{}

// BackupRetentionProperties Backup retention related properties.
type BackupRetentionProperties struct {
	// Number of days to keep recent snapshots.
	SnapshotRetentionDays *int32 `json:"snapshotRetentionDays,omitempty"`
	// Number of days to retain daily snapshots.
	DailySnapshotRetentionDays *int32 `json:"dailySnapshotRetentionDays,omitempty"`
	// Number of weeks to retain weekly snapshots..
	WeeklySnapshotRetentionWeeks *int32 `json:"weeklySnapshotRetentionWeeks,omitempty"`
	// Number of months to retain monthly snapshots.
	MonthlySnapshotRetentionMonths *int32 `json:"monthlySnapshotRetentionMonths,omitempty"`
}

// NewBackupRetentionProperties instantiates a new BackupRetentionProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupRetentionProperties() *BackupRetentionProperties {
	this := BackupRetentionProperties{}

	var snapshotRetentionDays int32 = 2
	this.SnapshotRetentionDays = &snapshotRetentionDays
	var dailySnapshotRetentionDays int32 = 0
	this.DailySnapshotRetentionDays = &dailySnapshotRetentionDays
	var weeklySnapshotRetentionWeeks int32 = 2
	this.WeeklySnapshotRetentionWeeks = &weeklySnapshotRetentionWeeks
	var monthlySnapshotRetentionMonths int32 = 1
	this.MonthlySnapshotRetentionMonths = &monthlySnapshotRetentionMonths

	return &this
}

// NewBackupRetentionPropertiesWithDefaults instantiates a new BackupRetentionProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupRetentionPropertiesWithDefaults() *BackupRetentionProperties {
	this := BackupRetentionProperties{}
	var snapshotRetentionDays int32 = 2
	this.SnapshotRetentionDays = &snapshotRetentionDays
	var dailySnapshotRetentionDays int32 = 0
	this.DailySnapshotRetentionDays = &dailySnapshotRetentionDays
	var weeklySnapshotRetentionWeeks int32 = 2
	this.WeeklySnapshotRetentionWeeks = &weeklySnapshotRetentionWeeks
	var monthlySnapshotRetentionMonths int32 = 1
	this.MonthlySnapshotRetentionMonths = &monthlySnapshotRetentionMonths
	return &this
}

// GetSnapshotRetentionDays returns the SnapshotRetentionDays field value if set, zero value otherwise.
func (o *BackupRetentionProperties) GetSnapshotRetentionDays() int32 {
	if o == nil || IsNil(o.SnapshotRetentionDays) {
		var ret int32
		return ret
	}
	return *o.SnapshotRetentionDays
}

// GetSnapshotRetentionDaysOk returns a tuple with the SnapshotRetentionDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRetentionProperties) GetSnapshotRetentionDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.SnapshotRetentionDays) {
		return nil, false
	}
	return o.SnapshotRetentionDays, true
}

// HasSnapshotRetentionDays returns a boolean if a field has been set.
func (o *BackupRetentionProperties) HasSnapshotRetentionDays() bool {
	if o != nil && !IsNil(o.SnapshotRetentionDays) {
		return true
	}

	return false
}

// SetSnapshotRetentionDays gets a reference to the given int32 and assigns it to the SnapshotRetentionDays field.
func (o *BackupRetentionProperties) SetSnapshotRetentionDays(v int32) {
	o.SnapshotRetentionDays = &v
}

// GetDailySnapshotRetentionDays returns the DailySnapshotRetentionDays field value if set, zero value otherwise.
func (o *BackupRetentionProperties) GetDailySnapshotRetentionDays() int32 {
	if o == nil || IsNil(o.DailySnapshotRetentionDays) {
		var ret int32
		return ret
	}
	return *o.DailySnapshotRetentionDays
}

// GetDailySnapshotRetentionDaysOk returns a tuple with the DailySnapshotRetentionDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRetentionProperties) GetDailySnapshotRetentionDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.DailySnapshotRetentionDays) {
		return nil, false
	}
	return o.DailySnapshotRetentionDays, true
}

// HasDailySnapshotRetentionDays returns a boolean if a field has been set.
func (o *BackupRetentionProperties) HasDailySnapshotRetentionDays() bool {
	if o != nil && !IsNil(o.DailySnapshotRetentionDays) {
		return true
	}

	return false
}

// SetDailySnapshotRetentionDays gets a reference to the given int32 and assigns it to the DailySnapshotRetentionDays field.
func (o *BackupRetentionProperties) SetDailySnapshotRetentionDays(v int32) {
	o.DailySnapshotRetentionDays = &v
}

// GetWeeklySnapshotRetentionWeeks returns the WeeklySnapshotRetentionWeeks field value if set, zero value otherwise.
func (o *BackupRetentionProperties) GetWeeklySnapshotRetentionWeeks() int32 {
	if o == nil || IsNil(o.WeeklySnapshotRetentionWeeks) {
		var ret int32
		return ret
	}
	return *o.WeeklySnapshotRetentionWeeks
}

// GetWeeklySnapshotRetentionWeeksOk returns a tuple with the WeeklySnapshotRetentionWeeks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRetentionProperties) GetWeeklySnapshotRetentionWeeksOk() (*int32, bool) {
	if o == nil || IsNil(o.WeeklySnapshotRetentionWeeks) {
		return nil, false
	}
	return o.WeeklySnapshotRetentionWeeks, true
}

// HasWeeklySnapshotRetentionWeeks returns a boolean if a field has been set.
func (o *BackupRetentionProperties) HasWeeklySnapshotRetentionWeeks() bool {
	if o != nil && !IsNil(o.WeeklySnapshotRetentionWeeks) {
		return true
	}

	return false
}

// SetWeeklySnapshotRetentionWeeks gets a reference to the given int32 and assigns it to the WeeklySnapshotRetentionWeeks field.
func (o *BackupRetentionProperties) SetWeeklySnapshotRetentionWeeks(v int32) {
	o.WeeklySnapshotRetentionWeeks = &v
}

// GetMonthlySnapshotRetentionMonths returns the MonthlySnapshotRetentionMonths field value if set, zero value otherwise.
func (o *BackupRetentionProperties) GetMonthlySnapshotRetentionMonths() int32 {
	if o == nil || IsNil(o.MonthlySnapshotRetentionMonths) {
		var ret int32
		return ret
	}
	return *o.MonthlySnapshotRetentionMonths
}

// GetMonthlySnapshotRetentionMonthsOk returns a tuple with the MonthlySnapshotRetentionMonths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRetentionProperties) GetMonthlySnapshotRetentionMonthsOk() (*int32, bool) {
	if o == nil || IsNil(o.MonthlySnapshotRetentionMonths) {
		return nil, false
	}
	return o.MonthlySnapshotRetentionMonths, true
}

// HasMonthlySnapshotRetentionMonths returns a boolean if a field has been set.
func (o *BackupRetentionProperties) HasMonthlySnapshotRetentionMonths() bool {
	if o != nil && !IsNil(o.MonthlySnapshotRetentionMonths) {
		return true
	}

	return false
}

// SetMonthlySnapshotRetentionMonths gets a reference to the given int32 and assigns it to the MonthlySnapshotRetentionMonths field.
func (o *BackupRetentionProperties) SetMonthlySnapshotRetentionMonths(v int32) {
	o.MonthlySnapshotRetentionMonths = &v
}

func (o BackupRetentionProperties) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackupRetentionProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SnapshotRetentionDays) {
		toSerialize["snapshotRetentionDays"] = o.SnapshotRetentionDays
	}
	if !IsNil(o.DailySnapshotRetentionDays) {
		toSerialize["dailySnapshotRetentionDays"] = o.DailySnapshotRetentionDays
	}
	if !IsNil(o.WeeklySnapshotRetentionWeeks) {
		toSerialize["weeklySnapshotRetentionWeeks"] = o.WeeklySnapshotRetentionWeeks
	}
	if !IsNil(o.MonthlySnapshotRetentionMonths) {
		toSerialize["monthlySnapshotRetentionMonths"] = o.MonthlySnapshotRetentionMonths
	}
	return toSerialize, nil
}

type NullableBackupRetentionProperties struct {
	value *BackupRetentionProperties
	isSet bool
}

func (v NullableBackupRetentionProperties) Get() *BackupRetentionProperties {
	return v.value
}

func (v *NullableBackupRetentionProperties) Set(val *BackupRetentionProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupRetentionProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupRetentionProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupRetentionProperties(val *BackupRetentionProperties) *NullableBackupRetentionProperties {
	return &NullableBackupRetentionProperties{value: val, isSet: true}
}

func (v NullableBackupRetentionProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupRetentionProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}