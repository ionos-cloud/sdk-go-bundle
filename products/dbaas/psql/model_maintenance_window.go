/*
 * IONOS DBaaS PostgreSQL REST API
 *
 * An enterprise-grade Database is provided as a Service (DBaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.  The API allows you to create additional PostgreSQL database clusters or modify existing ones. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psql

import (
	"encoding/json"
)

// checks if the MaintenanceWindow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MaintenanceWindow{}

// MaintenanceWindow A weekly 4 hour-long window, during which maintenance might occur.
type MaintenanceWindow struct {
	// Start of the maintenance window in UTC time.
	Time         string       `json:"time"`
	DayOfTheWeek DayOfTheWeek `json:"dayOfTheWeek"`
}

// NewMaintenanceWindow instantiates a new MaintenanceWindow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaintenanceWindow(time string, dayOfTheWeek DayOfTheWeek) *MaintenanceWindow {
	this := MaintenanceWindow{}

	this.Time = time
	this.DayOfTheWeek = dayOfTheWeek

	return &this
}

// NewMaintenanceWindowWithDefaults instantiates a new MaintenanceWindow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaintenanceWindowWithDefaults() *MaintenanceWindow {
	this := MaintenanceWindow{}
	return &this
}

// GetTime returns the Time field value
func (o *MaintenanceWindow) GetTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *MaintenanceWindow) GetTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *MaintenanceWindow) SetTime(v string) {
	o.Time = v
}

// GetDayOfTheWeek returns the DayOfTheWeek field value
func (o *MaintenanceWindow) GetDayOfTheWeek() DayOfTheWeek {
	if o == nil {
		var ret DayOfTheWeek
		return ret
	}

	return o.DayOfTheWeek
}

// GetDayOfTheWeekOk returns a tuple with the DayOfTheWeek field value
// and a boolean to check if the value has been set.
func (o *MaintenanceWindow) GetDayOfTheWeekOk() (*DayOfTheWeek, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DayOfTheWeek, true
}

// SetDayOfTheWeek sets field value
func (o *MaintenanceWindow) SetDayOfTheWeek(v DayOfTheWeek) {
	o.DayOfTheWeek = v
}

func (o MaintenanceWindow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["time"] = o.Time
	toSerialize["dayOfTheWeek"] = o.DayOfTheWeek
	return toSerialize, nil
}

type NullableMaintenanceWindow struct {
	value *MaintenanceWindow
	isSet bool
}

func (v NullableMaintenanceWindow) Get() *MaintenanceWindow {
	return v.value
}

func (v *NullableMaintenanceWindow) Set(val *MaintenanceWindow) {
	v.value = val
	v.isSet = true
}

func (v NullableMaintenanceWindow) IsSet() bool {
	return v.isSet
}

func (v *NullableMaintenanceWindow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaintenanceWindow(val *MaintenanceWindow) *NullableMaintenanceWindow {
	return &NullableMaintenanceWindow{value: val, isSet: true}
}

func (v NullableMaintenanceWindow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaintenanceWindow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
