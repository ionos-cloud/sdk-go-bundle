/*
 * IONOS DBaaS MongoDB REST API
 *
 * With IONOS Cloud Database as a Service, you have the ability to quickly set up and manage a MongoDB database. You can also delete clusters, manage backups and users via the API.   MongoDB is an open source, cross-platform, document-oriented database program. Classified as a NoSQL database program, it uses JSON-like documents with optional schemas.  The MongoDB API allows you to create additional database clusters or modify existing ones. Both tools, the Data Center Designer (DCD) and the API use the same concepts consistently and are well suited for smooth and intuitive use.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongo

import (
	"encoding/json"
)

// MaintenanceWindow A weekly window of 4 hours during which maintenance work can be performed.
type MaintenanceWindow struct {
	Time         *string       `json:"time"`
	DayOfTheWeek *DayOfTheWeek `json:"dayOfTheWeek"`
}

// NewMaintenanceWindow instantiates a new MaintenanceWindow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaintenanceWindow(time string, dayOfTheWeek DayOfTheWeek) *MaintenanceWindow {
	this := MaintenanceWindow{}

	this.Time = &time
	this.DayOfTheWeek = &dayOfTheWeek

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
// If the value is explicit nil, the zero value for string will be returned
func (o *MaintenanceWindow) GetTime() *string {
	if o == nil {
		return nil
	}

	return o.Time

}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MaintenanceWindow) GetTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Time, true
}

// SetTime sets field value
func (o *MaintenanceWindow) SetTime(v string) {

	o.Time = &v

}

// HasTime returns a boolean if a field has been set.
func (o *MaintenanceWindow) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// GetDayOfTheWeek returns the DayOfTheWeek field value
// If the value is explicit nil, the zero value for DayOfTheWeek will be returned
func (o *MaintenanceWindow) GetDayOfTheWeek() *DayOfTheWeek {
	if o == nil {
		return nil
	}

	return o.DayOfTheWeek

}

// GetDayOfTheWeekOk returns a tuple with the DayOfTheWeek field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MaintenanceWindow) GetDayOfTheWeekOk() (*DayOfTheWeek, bool) {
	if o == nil {
		return nil, false
	}

	return o.DayOfTheWeek, true
}

// SetDayOfTheWeek sets field value
func (o *MaintenanceWindow) SetDayOfTheWeek(v DayOfTheWeek) {

	o.DayOfTheWeek = &v

}

// HasDayOfTheWeek returns a boolean if a field has been set.
func (o *MaintenanceWindow) HasDayOfTheWeek() bool {
	if o != nil && o.DayOfTheWeek != nil {
		return true
	}

	return false
}

func (o MaintenanceWindow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}

	if o.DayOfTheWeek != nil {
		toSerialize["dayOfTheWeek"] = o.DayOfTheWeek
	}

	return json.Marshal(toSerialize)
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
