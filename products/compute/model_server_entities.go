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

// checks if the ServerEntities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerEntities{}

// ServerEntities struct for ServerEntities
type ServerEntities struct {
	Cdroms  *Cdroms          `json:"cdroms,omitempty"`
	Volumes *AttachedVolumes `json:"volumes,omitempty"`
	Nics    *Nics            `json:"nics,omitempty"`
}

// NewServerEntities instantiates a new ServerEntities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerEntities() *ServerEntities {
	this := ServerEntities{}

	return &this
}

// NewServerEntitiesWithDefaults instantiates a new ServerEntities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerEntitiesWithDefaults() *ServerEntities {
	this := ServerEntities{}
	return &this
}

// GetCdroms returns the Cdroms field value if set, zero value otherwise.
func (o *ServerEntities) GetCdroms() Cdroms {
	if o == nil || IsNil(o.Cdroms) {
		var ret Cdroms
		return ret
	}
	return *o.Cdroms
}

// GetCdromsOk returns a tuple with the Cdroms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerEntities) GetCdromsOk() (*Cdroms, bool) {
	if o == nil || IsNil(o.Cdroms) {
		return nil, false
	}
	return o.Cdroms, true
}

// HasCdroms returns a boolean if a field has been set.
func (o *ServerEntities) HasCdroms() bool {
	if o != nil && !IsNil(o.Cdroms) {
		return true
	}

	return false
}

// SetCdroms gets a reference to the given Cdroms and assigns it to the Cdroms field.
func (o *ServerEntities) SetCdroms(v Cdroms) {
	o.Cdroms = &v
}

// GetVolumes returns the Volumes field value if set, zero value otherwise.
func (o *ServerEntities) GetVolumes() AttachedVolumes {
	if o == nil || IsNil(o.Volumes) {
		var ret AttachedVolumes
		return ret
	}
	return *o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerEntities) GetVolumesOk() (*AttachedVolumes, bool) {
	if o == nil || IsNil(o.Volumes) {
		return nil, false
	}
	return o.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *ServerEntities) HasVolumes() bool {
	if o != nil && !IsNil(o.Volumes) {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given AttachedVolumes and assigns it to the Volumes field.
func (o *ServerEntities) SetVolumes(v AttachedVolumes) {
	o.Volumes = &v
}

// GetNics returns the Nics field value if set, zero value otherwise.
func (o *ServerEntities) GetNics() Nics {
	if o == nil || IsNil(o.Nics) {
		var ret Nics
		return ret
	}
	return *o.Nics
}

// GetNicsOk returns a tuple with the Nics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerEntities) GetNicsOk() (*Nics, bool) {
	if o == nil || IsNil(o.Nics) {
		return nil, false
	}
	return o.Nics, true
}

// HasNics returns a boolean if a field has been set.
func (o *ServerEntities) HasNics() bool {
	if o != nil && !IsNil(o.Nics) {
		return true
	}

	return false
}

// SetNics gets a reference to the given Nics and assigns it to the Nics field.
func (o *ServerEntities) SetNics(v Nics) {
	o.Nics = &v
}

func (o ServerEntities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cdroms) {
		toSerialize["cdroms"] = o.Cdroms
	}
	if !IsNil(o.Volumes) {
		toSerialize["volumes"] = o.Volumes
	}
	if !IsNil(o.Nics) {
		toSerialize["nics"] = o.Nics
	}
	return toSerialize, nil
}

type NullableServerEntities struct {
	value *ServerEntities
	isSet bool
}

func (v NullableServerEntities) Get() *ServerEntities {
	return v.value
}

func (v *NullableServerEntities) Set(val *ServerEntities) {
	v.value = val
	v.isSet = true
}

func (v NullableServerEntities) IsSet() bool {
	return v.isSet
}

func (v *NullableServerEntities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerEntities(val *ServerEntities) *NullableServerEntities {
	return &NullableServerEntities{value: val, isSet: true}
}

func (v NullableServerEntities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerEntities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
