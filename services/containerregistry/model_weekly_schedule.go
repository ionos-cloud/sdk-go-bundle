/*
 * Container Registry service
 *
 * Container Registry service enables IONOS clients to manage docker and OCI compliant registries for use by their managed Kubernetes clusters. Use a Container Registry to ensure you have a privately accessed registry to efficiently support image pulls.
 *
 * API version: 1.0
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package containerregistry

import (
	"encoding/json"
)

// WeeklySchedule struct for WeeklySchedule
type WeeklySchedule struct {
	Days *[]Day `json:"days"`
	// UTC time of day e.g. 01:00:00 - as defined by partial-time - RFC3339
	Time *string `json:"time"`
}

// NewWeeklySchedule instantiates a new WeeklySchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeeklySchedule(days []Day, time string) *WeeklySchedule {
	this := WeeklySchedule{}

	this.Days = &days
	this.Time = &time

	return &this
}

// NewWeeklyScheduleWithDefaults instantiates a new WeeklySchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeeklyScheduleWithDefaults() *WeeklySchedule {
	this := WeeklySchedule{}
	return &this
}

// GetDays returns the Days field value
// If the value is explicit nil, the zero value for []Day will be returned
func (o *WeeklySchedule) GetDays() *[]Day {
	if o == nil {
		return nil
	}

	return o.Days

}

// GetDaysOk returns a tuple with the Days field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WeeklySchedule) GetDaysOk() (*[]Day, bool) {
	if o == nil {
		return nil, false
	}

	return o.Days, true
}

// SetDays sets field value
func (o *WeeklySchedule) SetDays(v []Day) {

	o.Days = &v

}

// HasDays returns a boolean if a field has been set.
func (o *WeeklySchedule) HasDays() bool {
	if o != nil && o.Days != nil {
		return true
	}

	return false
}

// GetTime returns the Time field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WeeklySchedule) GetTime() *string {
	if o == nil {
		return nil
	}

	return o.Time

}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WeeklySchedule) GetTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Time, true
}

// SetTime sets field value
func (o *WeeklySchedule) SetTime(v string) {

	o.Time = &v

}

// HasTime returns a boolean if a field has been set.
func (o *WeeklySchedule) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

func (o WeeklySchedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["days"] = o.Days

	if o.Time != nil {
		toSerialize["time"] = o.Time
	}

	return json.Marshal(toSerialize)
}

type NullableWeeklySchedule struct {
	value *WeeklySchedule
	isSet bool
}

func (v NullableWeeklySchedule) Get() *WeeklySchedule {
	return v.value
}

func (v *NullableWeeklySchedule) Set(val *WeeklySchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableWeeklySchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableWeeklySchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeeklySchedule(val *WeeklySchedule) *NullableWeeklySchedule {
	return &NullableWeeklySchedule{value: val, isSet: true}
}

func (v NullableWeeklySchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeeklySchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
