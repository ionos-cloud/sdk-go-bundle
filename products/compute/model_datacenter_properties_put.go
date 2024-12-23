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

// checks if the DatacenterPropertiesPut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatacenterPropertiesPut{}

// DatacenterPropertiesPut struct for DatacenterPropertiesPut
type DatacenterPropertiesPut struct {
	// The name of the  resource.
	Name *string `json:"name,omitempty"`
	// A description for the datacenter, such as staging, production.
	Description *string `json:"description,omitempty"`
	// The physical location where the datacenter will be created. This will be where all of your servers live. Property cannot be modified after datacenter creation (disallowed in update requests).
	Location string `json:"location"`
	// The version of the data center; incremented with every change.
	Version *int32 `json:"version,omitempty"`
	// List of features supported by the location where this data center is provisioned.
	Features []string `json:"features,omitempty"`
	// Boolean value representing if the data center requires extra protection, such as two-step verification.
	SecAuthProtection *bool `json:"secAuthProtection,omitempty"`
	// Array of features and CPU families available in a location
	CpuArchitecture []CpuArchitectureProperties `json:"cpuArchitecture,omitempty"`
	// This value is either 'null' or contains an automatically-assigned /56 IPv6 CIDR block if IPv6 is enabled on this virtual data center. It can neither be changed nor removed.
	Ipv6CidrBlock NullableString `json:"ipv6CidrBlock,omitempty"`
	// This will become the default security group for the datacenter, replacing the old one if already exists.  This security group must already exists prior to this request. Provide this field only if the `createDefaultSecurityGroup` field is missing. You cannot provide both of them
	DefaultSecurityGroupId *string `json:"defaultSecurityGroupId,omitempty"`
	// If this field is set on true and this datacenter has no default security group then a default security group, with predefined rules, will be created for this datacenter. Default value is false.  Provide this field only if the `defaultSecurityGroupId` field is missing. You cannot provide both of them
	CreateDefaultSecurityGroup *bool `json:"createDefaultSecurityGroup,omitempty"`
}

// NewDatacenterPropertiesPut instantiates a new DatacenterPropertiesPut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatacenterPropertiesPut(location string) *DatacenterPropertiesPut {
	this := DatacenterPropertiesPut{}

	this.Location = location

	return &this
}

// NewDatacenterPropertiesPutWithDefaults instantiates a new DatacenterPropertiesPut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatacenterPropertiesPutWithDefaults() *DatacenterPropertiesPut {
	this := DatacenterPropertiesPut{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DatacenterPropertiesPut) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatacenterPropertiesPut) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DatacenterPropertiesPut) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DatacenterPropertiesPut) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DatacenterPropertiesPut) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatacenterPropertiesPut) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DatacenterPropertiesPut) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DatacenterPropertiesPut) SetDescription(v string) {
	o.Description = &v
}

// GetLocation returns the Location field value
func (o *DatacenterPropertiesPut) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *DatacenterPropertiesPut) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *DatacenterPropertiesPut) SetLocation(v string) {
	o.Location = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DatacenterPropertiesPut) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatacenterPropertiesPut) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DatacenterPropertiesPut) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *DatacenterPropertiesPut) SetVersion(v int32) {
	o.Version = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *DatacenterPropertiesPut) GetFeatures() []string {
	if o == nil || IsNil(o.Features) {
		var ret []string
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatacenterPropertiesPut) GetFeaturesOk() ([]string, bool) {
	if o == nil || IsNil(o.Features) {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *DatacenterPropertiesPut) HasFeatures() bool {
	if o != nil && !IsNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []string and assigns it to the Features field.
func (o *DatacenterPropertiesPut) SetFeatures(v []string) {
	o.Features = v
}

// GetSecAuthProtection returns the SecAuthProtection field value if set, zero value otherwise.
func (o *DatacenterPropertiesPut) GetSecAuthProtection() bool {
	if o == nil || IsNil(o.SecAuthProtection) {
		var ret bool
		return ret
	}
	return *o.SecAuthProtection
}

// GetSecAuthProtectionOk returns a tuple with the SecAuthProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatacenterPropertiesPut) GetSecAuthProtectionOk() (*bool, bool) {
	if o == nil || IsNil(o.SecAuthProtection) {
		return nil, false
	}
	return o.SecAuthProtection, true
}

// HasSecAuthProtection returns a boolean if a field has been set.
func (o *DatacenterPropertiesPut) HasSecAuthProtection() bool {
	if o != nil && !IsNil(o.SecAuthProtection) {
		return true
	}

	return false
}

// SetSecAuthProtection gets a reference to the given bool and assigns it to the SecAuthProtection field.
func (o *DatacenterPropertiesPut) SetSecAuthProtection(v bool) {
	o.SecAuthProtection = &v
}

// GetCpuArchitecture returns the CpuArchitecture field value if set, zero value otherwise.
func (o *DatacenterPropertiesPut) GetCpuArchitecture() []CpuArchitectureProperties {
	if o == nil || IsNil(o.CpuArchitecture) {
		var ret []CpuArchitectureProperties
		return ret
	}
	return o.CpuArchitecture
}

// GetCpuArchitectureOk returns a tuple with the CpuArchitecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatacenterPropertiesPut) GetCpuArchitectureOk() ([]CpuArchitectureProperties, bool) {
	if o == nil || IsNil(o.CpuArchitecture) {
		return nil, false
	}
	return o.CpuArchitecture, true
}

// HasCpuArchitecture returns a boolean if a field has been set.
func (o *DatacenterPropertiesPut) HasCpuArchitecture() bool {
	if o != nil && !IsNil(o.CpuArchitecture) {
		return true
	}

	return false
}

// SetCpuArchitecture gets a reference to the given []CpuArchitectureProperties and assigns it to the CpuArchitecture field.
func (o *DatacenterPropertiesPut) SetCpuArchitecture(v []CpuArchitectureProperties) {
	o.CpuArchitecture = v
}

// GetIpv6CidrBlock returns the Ipv6CidrBlock field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatacenterPropertiesPut) GetIpv6CidrBlock() string {
	if o == nil || IsNil(o.Ipv6CidrBlock.Get()) {
		var ret string
		return ret
	}
	return *o.Ipv6CidrBlock.Get()
}

// GetIpv6CidrBlockOk returns a tuple with the Ipv6CidrBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatacenterPropertiesPut) GetIpv6CidrBlockOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ipv6CidrBlock.Get(), o.Ipv6CidrBlock.IsSet()
}

// HasIpv6CidrBlock returns a boolean if a field has been set.
func (o *DatacenterPropertiesPut) HasIpv6CidrBlock() bool {
	if o != nil && o.Ipv6CidrBlock.IsSet() {
		return true
	}

	return false
}

// SetIpv6CidrBlock gets a reference to the given NullableString and assigns it to the Ipv6CidrBlock field.
func (o *DatacenterPropertiesPut) SetIpv6CidrBlock(v string) {
	o.Ipv6CidrBlock.Set(&v)
}

// SetIpv6CidrBlockNil sets the value for Ipv6CidrBlock to be an explicit nil
func (o *DatacenterPropertiesPut) SetIpv6CidrBlockNil() {
	o.Ipv6CidrBlock.Set(nil)
}

// UnsetIpv6CidrBlock ensures that no value is present for Ipv6CidrBlock, not even an explicit nil
func (o *DatacenterPropertiesPut) UnsetIpv6CidrBlock() {
	o.Ipv6CidrBlock.Unset()
}

// GetDefaultSecurityGroupId returns the DefaultSecurityGroupId field value if set, zero value otherwise.
func (o *DatacenterPropertiesPut) GetDefaultSecurityGroupId() string {
	if o == nil || IsNil(o.DefaultSecurityGroupId) {
		var ret string
		return ret
	}
	return *o.DefaultSecurityGroupId
}

// GetDefaultSecurityGroupIdOk returns a tuple with the DefaultSecurityGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatacenterPropertiesPut) GetDefaultSecurityGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultSecurityGroupId) {
		return nil, false
	}
	return o.DefaultSecurityGroupId, true
}

// HasDefaultSecurityGroupId returns a boolean if a field has been set.
func (o *DatacenterPropertiesPut) HasDefaultSecurityGroupId() bool {
	if o != nil && !IsNil(o.DefaultSecurityGroupId) {
		return true
	}

	return false
}

// SetDefaultSecurityGroupId gets a reference to the given string and assigns it to the DefaultSecurityGroupId field.
func (o *DatacenterPropertiesPut) SetDefaultSecurityGroupId(v string) {
	o.DefaultSecurityGroupId = &v
}

// GetCreateDefaultSecurityGroup returns the CreateDefaultSecurityGroup field value if set, zero value otherwise.
func (o *DatacenterPropertiesPut) GetCreateDefaultSecurityGroup() bool {
	if o == nil || IsNil(o.CreateDefaultSecurityGroup) {
		var ret bool
		return ret
	}
	return *o.CreateDefaultSecurityGroup
}

// GetCreateDefaultSecurityGroupOk returns a tuple with the CreateDefaultSecurityGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatacenterPropertiesPut) GetCreateDefaultSecurityGroupOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateDefaultSecurityGroup) {
		return nil, false
	}
	return o.CreateDefaultSecurityGroup, true
}

// HasCreateDefaultSecurityGroup returns a boolean if a field has been set.
func (o *DatacenterPropertiesPut) HasCreateDefaultSecurityGroup() bool {
	if o != nil && !IsNil(o.CreateDefaultSecurityGroup) {
		return true
	}

	return false
}

// SetCreateDefaultSecurityGroup gets a reference to the given bool and assigns it to the CreateDefaultSecurityGroup field.
func (o *DatacenterPropertiesPut) SetCreateDefaultSecurityGroup(v bool) {
	o.CreateDefaultSecurityGroup = &v
}

func (o DatacenterPropertiesPut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["location"] = o.Location
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	if !IsNil(o.SecAuthProtection) {
		toSerialize["secAuthProtection"] = o.SecAuthProtection
	}
	if !IsNil(o.CpuArchitecture) {
		toSerialize["cpuArchitecture"] = o.CpuArchitecture
	}
	if o.Ipv6CidrBlock.IsSet() {
		toSerialize["ipv6CidrBlock"] = o.Ipv6CidrBlock.Get()
	}
	if !IsNil(o.DefaultSecurityGroupId) {
		toSerialize["defaultSecurityGroupId"] = o.DefaultSecurityGroupId
	}
	if !IsNil(o.CreateDefaultSecurityGroup) {
		toSerialize["createDefaultSecurityGroup"] = o.CreateDefaultSecurityGroup
	}
	return toSerialize, nil
}

type NullableDatacenterPropertiesPut struct {
	value *DatacenterPropertiesPut
	isSet bool
}

func (v NullableDatacenterPropertiesPut) Get() *DatacenterPropertiesPut {
	return v.value
}

func (v *NullableDatacenterPropertiesPut) Set(val *DatacenterPropertiesPut) {
	v.value = val
	v.isSet = true
}

func (v NullableDatacenterPropertiesPut) IsSet() bool {
	return v.isSet
}

func (v *NullableDatacenterPropertiesPut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatacenterPropertiesPut(val *DatacenterPropertiesPut) *NullableDatacenterPropertiesPut {
	return &NullableDatacenterPropertiesPut{value: val, isSet: true}
}

func (v NullableDatacenterPropertiesPut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatacenterPropertiesPut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
