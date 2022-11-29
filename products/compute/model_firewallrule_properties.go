/*
 * CLOUD API
 *
 * IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package compute

import (
	"encoding/json"
)

// FirewallruleProperties struct for FirewallruleProperties
type FirewallruleProperties struct {
	// The name of the  resource.
	Name *string `json:"name,omitempty"`
	// The protocol for the rule. Property cannot be modified after it is created (disallowed in update requests).
	Protocol *string `json:"protocol"`
	// Only traffic originating from the respective MAC address is allowed. Valid format: aa:bb:cc:dd:ee:ff. Value null allows traffic from any MAC address.
	SourceMac *string `json:"sourceMac,omitempty"`
	// Only traffic originating from the respective IPv4 address is allowed. Value null allows traffic from any IP address.
	SourceIp *string `json:"sourceIp,omitempty"`
	// If the target NIC has multiple IP addresses, only the traffic directed to the respective IP address of the NIC is allowed. Value null Value null allows traffic to any target IP address.
	TargetIp *string `json:"targetIp,omitempty"`
	// Defines the allowed code (from 0 to 254) if protocol ICMP is chosen. Value null allows all codes.
	IcmpCode *int32 `json:"icmpCode,omitempty"`
	// Defines the allowed type (from 0 to 254) if the protocol ICMP is chosen. Value null allows all types.
	IcmpType *int32 `json:"icmpType,omitempty"`
	// Defines the start range of the allowed port (from 1 to 65534) if protocol TCP or UDP is chosen. Leave portRangeStart and portRangeEnd value null to allow all ports.
	PortRangeStart *int32 `json:"portRangeStart,omitempty"`
	// Defines the end range of the allowed port (from 1 to 65534) if the protocol TCP or UDP is chosen. Leave portRangeStart and portRangeEnd null to allow all ports.
	PortRangeEnd *int32 `json:"portRangeEnd,omitempty"`
	// The type of the firewall rule. If not specified, the default INGRESS value is used.
	Type *string `json:"type,omitempty"`
}

// NewFirewallruleProperties instantiates a new FirewallruleProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirewallruleProperties(protocol string) *FirewallruleProperties {
	this := FirewallruleProperties{}

	this.Protocol = &protocol

	return &this
}

// NewFirewallrulePropertiesWithDefaults instantiates a new FirewallruleProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirewallrulePropertiesWithDefaults() *FirewallruleProperties {
	this := FirewallruleProperties{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FirewallruleProperties) GetName() *string {
	if o == nil {
		return nil
	}

	return o.Name

}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirewallruleProperties) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Name, true
}

// SetName sets field value
func (o *FirewallruleProperties) SetName(v string) {

	o.Name = &v

}

// HasName returns a boolean if a field has been set.
func (o *FirewallruleProperties) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// GetProtocol returns the Protocol field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FirewallruleProperties) GetProtocol() *string {
	if o == nil {
		return nil
	}

	return o.Protocol

}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirewallruleProperties) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Protocol, true
}

// SetProtocol sets field value
func (o *FirewallruleProperties) SetProtocol(v string) {

	o.Protocol = &v

}

// HasProtocol returns a boolean if a field has been set.
func (o *FirewallruleProperties) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// GetSourceMac returns the SourceMac field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FirewallruleProperties) GetSourceMac() *string {
	if o == nil {
		return nil
	}

	return o.SourceMac

}

// GetSourceMacOk returns a tuple with the SourceMac field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirewallruleProperties) GetSourceMacOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.SourceMac, true
}

// SetSourceMac sets field value
func (o *FirewallruleProperties) SetSourceMac(v string) {

	o.SourceMac = &v

}

// HasSourceMac returns a boolean if a field has been set.
func (o *FirewallruleProperties) HasSourceMac() bool {
	if o != nil && o.SourceMac != nil {
		return true
	}

	return false
}

// GetSourceIp returns the SourceIp field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FirewallruleProperties) GetSourceIp() *string {
	if o == nil {
		return nil
	}

	return o.SourceIp

}

// GetSourceIpOk returns a tuple with the SourceIp field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirewallruleProperties) GetSourceIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.SourceIp, true
}

// SetSourceIp sets field value
func (o *FirewallruleProperties) SetSourceIp(v string) {

	o.SourceIp = &v

}

// HasSourceIp returns a boolean if a field has been set.
func (o *FirewallruleProperties) HasSourceIp() bool {
	if o != nil && o.SourceIp != nil {
		return true
	}

	return false
}

// GetTargetIp returns the TargetIp field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FirewallruleProperties) GetTargetIp() *string {
	if o == nil {
		return nil
	}

	return o.TargetIp

}

// GetTargetIpOk returns a tuple with the TargetIp field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirewallruleProperties) GetTargetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.TargetIp, true
}

// SetTargetIp sets field value
func (o *FirewallruleProperties) SetTargetIp(v string) {

	o.TargetIp = &v

}

// HasTargetIp returns a boolean if a field has been set.
func (o *FirewallruleProperties) HasTargetIp() bool {
	if o != nil && o.TargetIp != nil {
		return true
	}

	return false
}

// GetIcmpCode returns the IcmpCode field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *FirewallruleProperties) GetIcmpCode() *int32 {
	if o == nil {
		return nil
	}

	return o.IcmpCode

}

// GetIcmpCodeOk returns a tuple with the IcmpCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirewallruleProperties) GetIcmpCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.IcmpCode, true
}

// SetIcmpCode sets field value
func (o *FirewallruleProperties) SetIcmpCode(v int32) {

	o.IcmpCode = &v

}

// HasIcmpCode returns a boolean if a field has been set.
func (o *FirewallruleProperties) HasIcmpCode() bool {
	if o != nil && o.IcmpCode != nil {
		return true
	}

	return false
}

// GetIcmpType returns the IcmpType field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *FirewallruleProperties) GetIcmpType() *int32 {
	if o == nil {
		return nil
	}

	return o.IcmpType

}

// GetIcmpTypeOk returns a tuple with the IcmpType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirewallruleProperties) GetIcmpTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.IcmpType, true
}

// SetIcmpType sets field value
func (o *FirewallruleProperties) SetIcmpType(v int32) {

	o.IcmpType = &v

}

// HasIcmpType returns a boolean if a field has been set.
func (o *FirewallruleProperties) HasIcmpType() bool {
	if o != nil && o.IcmpType != nil {
		return true
	}

	return false
}

// GetPortRangeStart returns the PortRangeStart field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *FirewallruleProperties) GetPortRangeStart() *int32 {
	if o == nil {
		return nil
	}

	return o.PortRangeStart

}

// GetPortRangeStartOk returns a tuple with the PortRangeStart field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirewallruleProperties) GetPortRangeStartOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.PortRangeStart, true
}

// SetPortRangeStart sets field value
func (o *FirewallruleProperties) SetPortRangeStart(v int32) {

	o.PortRangeStart = &v

}

// HasPortRangeStart returns a boolean if a field has been set.
func (o *FirewallruleProperties) HasPortRangeStart() bool {
	if o != nil && o.PortRangeStart != nil {
		return true
	}

	return false
}

// GetPortRangeEnd returns the PortRangeEnd field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *FirewallruleProperties) GetPortRangeEnd() *int32 {
	if o == nil {
		return nil
	}

	return o.PortRangeEnd

}

// GetPortRangeEndOk returns a tuple with the PortRangeEnd field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirewallruleProperties) GetPortRangeEndOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.PortRangeEnd, true
}

// SetPortRangeEnd sets field value
func (o *FirewallruleProperties) SetPortRangeEnd(v int32) {

	o.PortRangeEnd = &v

}

// HasPortRangeEnd returns a boolean if a field has been set.
func (o *FirewallruleProperties) HasPortRangeEnd() bool {
	if o != nil && o.PortRangeEnd != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FirewallruleProperties) GetType() *string {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirewallruleProperties) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *FirewallruleProperties) SetType(v string) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *FirewallruleProperties) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

func (o FirewallruleProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	if o.Protocol != nil {
		toSerialize["protocol"] = o.Protocol
	}

	toSerialize["sourceMac"] = o.SourceMac

	toSerialize["sourceIp"] = o.SourceIp

	toSerialize["targetIp"] = o.TargetIp

	toSerialize["icmpCode"] = o.IcmpCode

	toSerialize["icmpType"] = o.IcmpType

	if o.PortRangeStart != nil {
		toSerialize["portRangeStart"] = o.PortRangeStart
	}

	if o.PortRangeEnd != nil {
		toSerialize["portRangeEnd"] = o.PortRangeEnd
	}

	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	return json.Marshal(toSerialize)
}

type NullableFirewallruleProperties struct {
	value *FirewallruleProperties
	isSet bool
}

func (v NullableFirewallruleProperties) Get() *FirewallruleProperties {
	return v.value
}

func (v *NullableFirewallruleProperties) Set(val *FirewallruleProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableFirewallruleProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableFirewallruleProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirewallruleProperties(val *FirewallruleProperties) *NullableFirewallruleProperties {
	return &NullableFirewallruleProperties{value: val, isSet: true}
}

func (v NullableFirewallruleProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirewallruleProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
