/*
 * CLOUD API
 *
 *  IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package compute

import (
	"encoding/json"
)

// checks if the LocationProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocationProperties{}

// LocationProperties struct for LocationProperties
type LocationProperties struct {
	// The location name.
	Name *string `json:"name,omitempty"`
	// A list of available features in the location.
	Features []string `json:"features,omitempty"`
	// A list of image aliases available in the location.
	ImageAliases []string `json:"imageAliases,omitempty"`
	// A list of available CPU types and related resources available in the location.
	CpuArchitecture []CpuArchitectureProperties `json:"cpuArchitecture,omitempty"`
}

// NewLocationProperties instantiates a new LocationProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationProperties() *LocationProperties {
	this := LocationProperties{}

	return &this
}

// NewLocationPropertiesWithDefaults instantiates a new LocationProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationPropertiesWithDefaults() *LocationProperties {
	this := LocationProperties{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LocationProperties) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationProperties) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LocationProperties) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LocationProperties) SetName(v string) {
	o.Name = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *LocationProperties) GetFeatures() []string {
	if o == nil || IsNil(o.Features) {
		var ret []string
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationProperties) GetFeaturesOk() ([]string, bool) {
	if o == nil || IsNil(o.Features) {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *LocationProperties) HasFeatures() bool {
	if o != nil && !IsNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []string and assigns it to the Features field.
func (o *LocationProperties) SetFeatures(v []string) {
	o.Features = v
}

// GetImageAliases returns the ImageAliases field value if set, zero value otherwise.
func (o *LocationProperties) GetImageAliases() []string {
	if o == nil || IsNil(o.ImageAliases) {
		var ret []string
		return ret
	}
	return o.ImageAliases
}

// GetImageAliasesOk returns a tuple with the ImageAliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationProperties) GetImageAliasesOk() ([]string, bool) {
	if o == nil || IsNil(o.ImageAliases) {
		return nil, false
	}
	return o.ImageAliases, true
}

// HasImageAliases returns a boolean if a field has been set.
func (o *LocationProperties) HasImageAliases() bool {
	if o != nil && !IsNil(o.ImageAliases) {
		return true
	}

	return false
}

// SetImageAliases gets a reference to the given []string and assigns it to the ImageAliases field.
func (o *LocationProperties) SetImageAliases(v []string) {
	o.ImageAliases = v
}

// GetCpuArchitecture returns the CpuArchitecture field value if set, zero value otherwise.
func (o *LocationProperties) GetCpuArchitecture() []CpuArchitectureProperties {
	if o == nil || IsNil(o.CpuArchitecture) {
		var ret []CpuArchitectureProperties
		return ret
	}
	return o.CpuArchitecture
}

// GetCpuArchitectureOk returns a tuple with the CpuArchitecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationProperties) GetCpuArchitectureOk() ([]CpuArchitectureProperties, bool) {
	if o == nil || IsNil(o.CpuArchitecture) {
		return nil, false
	}
	return o.CpuArchitecture, true
}

// HasCpuArchitecture returns a boolean if a field has been set.
func (o *LocationProperties) HasCpuArchitecture() bool {
	if o != nil && !IsNil(o.CpuArchitecture) {
		return true
	}

	return false
}

// SetCpuArchitecture gets a reference to the given []CpuArchitectureProperties and assigns it to the CpuArchitecture field.
func (o *LocationProperties) SetCpuArchitecture(v []CpuArchitectureProperties) {
	o.CpuArchitecture = v
}

func (o LocationProperties) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocationProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	if !IsNil(o.ImageAliases) {
		toSerialize["imageAliases"] = o.ImageAliases
	}
	if !IsNil(o.CpuArchitecture) {
		toSerialize["cpuArchitecture"] = o.CpuArchitecture
	}
	return toSerialize, nil
}

type NullableLocationProperties struct {
	value *LocationProperties
	isSet bool
}

func (v NullableLocationProperties) Get() *LocationProperties {
	return v.value
}

func (v *NullableLocationProperties) Set(val *LocationProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationProperties(val *LocationProperties) *NullableLocationProperties {
	return &NullableLocationProperties{value: val, isSet: true}
}

func (v NullableLocationProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
