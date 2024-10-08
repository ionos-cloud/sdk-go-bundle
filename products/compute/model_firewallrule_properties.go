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

// checks if the FirewallruleProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirewallruleProperties{}

// FirewallruleProperties struct for FirewallruleProperties
type FirewallruleProperties struct {
	// The name of the  resource.
	Name *string `json:"name,omitempty"`
	// The protocol for the rule. Property cannot be modified after it is created (disallowed in update requests).
	Protocol string `json:"protocol"`
	// Only traffic originating from the respective MAC address is allowed. Valid format: aa:bb:cc:dd:ee:ff. Value null allows traffic from any MAC address.
	SourceMac NullableString `json:"sourceMac,omitempty"`
	// The IP version for this rule. If sourceIp or targetIp are specified, you can omit this value - the IP version will then be deduced from the IP address(es) used; if you specify it anyway, it must match the specified IP address(es). If neither sourceIp nor targetIp are specified, this rule allows traffic only for the specified IP version. If neither sourceIp, targetIp nor ipVersion are specified, this rule will only allow IPv4 traffic.
	IpVersion NullableString `json:"ipVersion,omitempty"`
	// Only traffic originating from the respective IP address (or CIDR block) is allowed. Value null allows traffic from any IP address (according to the selected ipVersion).
	SourceIp NullableString `json:"sourceIp,omitempty"`
	// If the target NIC has multiple IP addresses, only the traffic directed to the respective IP address (or CIDR block) of the NIC is allowed. Value null allows traffic to any target IP address (according to the selected ipVersion).
	TargetIp NullableString `json:"targetIp,omitempty"`
	// Defines the allowed code (from 0 to 254) if protocol ICMP or ICMPv6 is chosen. Value null allows all codes.
	IcmpCode NullableInt32 `json:"icmpCode,omitempty"`
	// Defines the allowed type (from 0 to 254) if the protocol ICMP or ICMPv6 is chosen. Value null allows all types.
	IcmpType NullableInt32 `json:"icmpType,omitempty"`
	// Defines the start range of the allowed port (from 1 to 65535) if protocol TCP or UDP is chosen. Leave portRangeStart and portRangeEnd value null to allow all ports.
	PortRangeStart *int32 `json:"portRangeStart,omitempty"`
	// Defines the end range of the allowed port (from 1 to 65535) if the protocol TCP or UDP is chosen. Leave portRangeStart and portRangeEnd null to allow all ports.
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

	this.Protocol = protocol

	return &this
}

// NewFirewallrulePropertiesWithDefaults instantiates a new FirewallruleProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirewallrulePropertiesWithDefaults() *FirewallruleProperties {
	this := FirewallruleProperties{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FirewallruleProperties) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallruleProperties) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FirewallruleProperties) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FirewallruleProperties) SetName(v string) {
	o.Name = &v
}

// GetProtocol returns the Protocol field value
func (o *FirewallruleProperties) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *FirewallruleProperties) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *FirewallruleProperties) SetProtocol(v string) {
	o.Protocol = v
}

// GetSourceMac returns the SourceMac field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirewallruleProperties) GetSourceMac() string {
	if o == nil || IsNil(o.SourceMac.Get()) {
		var ret string
		return ret
	}
	return *o.SourceMac.Get()
}

// GetSourceMacOk returns a tuple with the SourceMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirewallruleProperties) GetSourceMacOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceMac.Get(), o.SourceMac.IsSet()
}

// HasSourceMac returns a boolean if a field has been set.
func (o *FirewallruleProperties) HasSourceMac() bool {
	if o != nil && o.SourceMac.IsSet() {
		return true
	}

	return false
}

// SetSourceMac gets a reference to the given NullableString and assigns it to the SourceMac field.
func (o *FirewallruleProperties) SetSourceMac(v string) {
	o.SourceMac.Set(&v)
}

// SetSourceMacNil sets the value for SourceMac to be an explicit nil
func (o *FirewallruleProperties) SetSourceMacNil() {
	o.SourceMac.Set(nil)
}

// UnsetSourceMac ensures that no value is present for SourceMac, not even an explicit nil
func (o *FirewallruleProperties) UnsetSourceMac() {
	o.SourceMac.Unset()
}

// GetIpVersion returns the IpVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirewallruleProperties) GetIpVersion() string {
	if o == nil || IsNil(o.IpVersion.Get()) {
		var ret string
		return ret
	}
	return *o.IpVersion.Get()
}

// GetIpVersionOk returns a tuple with the IpVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirewallruleProperties) GetIpVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpVersion.Get(), o.IpVersion.IsSet()
}

// HasIpVersion returns a boolean if a field has been set.
func (o *FirewallruleProperties) HasIpVersion() bool {
	if o != nil && o.IpVersion.IsSet() {
		return true
	}

	return false
}

// SetIpVersion gets a reference to the given NullableString and assigns it to the IpVersion field.
func (o *FirewallruleProperties) SetIpVersion(v string) {
	o.IpVersion.Set(&v)
}

// SetIpVersionNil sets the value for IpVersion to be an explicit nil
func (o *FirewallruleProperties) SetIpVersionNil() {
	o.IpVersion.Set(nil)
}

// UnsetIpVersion ensures that no value is present for IpVersion, not even an explicit nil
func (o *FirewallruleProperties) UnsetIpVersion() {
	o.IpVersion.Unset()
}

// GetSourceIp returns the SourceIp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirewallruleProperties) GetSourceIp() string {
	if o == nil || IsNil(o.SourceIp.Get()) {
		var ret string
		return ret
	}
	return *o.SourceIp.Get()
}

// GetSourceIpOk returns a tuple with the SourceIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirewallruleProperties) GetSourceIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceIp.Get(), o.SourceIp.IsSet()
}

// HasSourceIp returns a boolean if a field has been set.
func (o *FirewallruleProperties) HasSourceIp() bool {
	if o != nil && o.SourceIp.IsSet() {
		return true
	}

	return false
}

// SetSourceIp gets a reference to the given NullableString and assigns it to the SourceIp field.
func (o *FirewallruleProperties) SetSourceIp(v string) {
	o.SourceIp.Set(&v)
}

// SetSourceIpNil sets the value for SourceIp to be an explicit nil
func (o *FirewallruleProperties) SetSourceIpNil() {
	o.SourceIp.Set(nil)
}

// UnsetSourceIp ensures that no value is present for SourceIp, not even an explicit nil
func (o *FirewallruleProperties) UnsetSourceIp() {
	o.SourceIp.Unset()
}

// GetTargetIp returns the TargetIp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirewallruleProperties) GetTargetIp() string {
	if o == nil || IsNil(o.TargetIp.Get()) {
		var ret string
		return ret
	}
	return *o.TargetIp.Get()
}

// GetTargetIpOk returns a tuple with the TargetIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirewallruleProperties) GetTargetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetIp.Get(), o.TargetIp.IsSet()
}

// HasTargetIp returns a boolean if a field has been set.
func (o *FirewallruleProperties) HasTargetIp() bool {
	if o != nil && o.TargetIp.IsSet() {
		return true
	}

	return false
}

// SetTargetIp gets a reference to the given NullableString and assigns it to the TargetIp field.
func (o *FirewallruleProperties) SetTargetIp(v string) {
	o.TargetIp.Set(&v)
}

// SetTargetIpNil sets the value for TargetIp to be an explicit nil
func (o *FirewallruleProperties) SetTargetIpNil() {
	o.TargetIp.Set(nil)
}

// UnsetTargetIp ensures that no value is present for TargetIp, not even an explicit nil
func (o *FirewallruleProperties) UnsetTargetIp() {
	o.TargetIp.Unset()
}

// GetIcmpCode returns the IcmpCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirewallruleProperties) GetIcmpCode() int32 {
	if o == nil || IsNil(o.IcmpCode.Get()) {
		var ret int32
		return ret
	}
	return *o.IcmpCode.Get()
}

// GetIcmpCodeOk returns a tuple with the IcmpCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirewallruleProperties) GetIcmpCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.IcmpCode.Get(), o.IcmpCode.IsSet()
}

// HasIcmpCode returns a boolean if a field has been set.
func (o *FirewallruleProperties) HasIcmpCode() bool {
	if o != nil && o.IcmpCode.IsSet() {
		return true
	}

	return false
}

// SetIcmpCode gets a reference to the given NullableInt32 and assigns it to the IcmpCode field.
func (o *FirewallruleProperties) SetIcmpCode(v int32) {
	o.IcmpCode.Set(&v)
}

// SetIcmpCodeNil sets the value for IcmpCode to be an explicit nil
func (o *FirewallruleProperties) SetIcmpCodeNil() {
	o.IcmpCode.Set(nil)
}

// UnsetIcmpCode ensures that no value is present for IcmpCode, not even an explicit nil
func (o *FirewallruleProperties) UnsetIcmpCode() {
	o.IcmpCode.Unset()
}

// GetIcmpType returns the IcmpType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirewallruleProperties) GetIcmpType() int32 {
	if o == nil || IsNil(o.IcmpType.Get()) {
		var ret int32
		return ret
	}
	return *o.IcmpType.Get()
}

// GetIcmpTypeOk returns a tuple with the IcmpType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirewallruleProperties) GetIcmpTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.IcmpType.Get(), o.IcmpType.IsSet()
}

// HasIcmpType returns a boolean if a field has been set.
func (o *FirewallruleProperties) HasIcmpType() bool {
	if o != nil && o.IcmpType.IsSet() {
		return true
	}

	return false
}

// SetIcmpType gets a reference to the given NullableInt32 and assigns it to the IcmpType field.
func (o *FirewallruleProperties) SetIcmpType(v int32) {
	o.IcmpType.Set(&v)
}

// SetIcmpTypeNil sets the value for IcmpType to be an explicit nil
func (o *FirewallruleProperties) SetIcmpTypeNil() {
	o.IcmpType.Set(nil)
}

// UnsetIcmpType ensures that no value is present for IcmpType, not even an explicit nil
func (o *FirewallruleProperties) UnsetIcmpType() {
	o.IcmpType.Unset()
}

// GetPortRangeStart returns the PortRangeStart field value if set, zero value otherwise.
func (o *FirewallruleProperties) GetPortRangeStart() int32 {
	if o == nil || IsNil(o.PortRangeStart) {
		var ret int32
		return ret
	}
	return *o.PortRangeStart
}

// GetPortRangeStartOk returns a tuple with the PortRangeStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallruleProperties) GetPortRangeStartOk() (*int32, bool) {
	if o == nil || IsNil(o.PortRangeStart) {
		return nil, false
	}
	return o.PortRangeStart, true
}

// HasPortRangeStart returns a boolean if a field has been set.
func (o *FirewallruleProperties) HasPortRangeStart() bool {
	if o != nil && !IsNil(o.PortRangeStart) {
		return true
	}

	return false
}

// SetPortRangeStart gets a reference to the given int32 and assigns it to the PortRangeStart field.
func (o *FirewallruleProperties) SetPortRangeStart(v int32) {
	o.PortRangeStart = &v
}

// GetPortRangeEnd returns the PortRangeEnd field value if set, zero value otherwise.
func (o *FirewallruleProperties) GetPortRangeEnd() int32 {
	if o == nil || IsNil(o.PortRangeEnd) {
		var ret int32
		return ret
	}
	return *o.PortRangeEnd
}

// GetPortRangeEndOk returns a tuple with the PortRangeEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallruleProperties) GetPortRangeEndOk() (*int32, bool) {
	if o == nil || IsNil(o.PortRangeEnd) {
		return nil, false
	}
	return o.PortRangeEnd, true
}

// HasPortRangeEnd returns a boolean if a field has been set.
func (o *FirewallruleProperties) HasPortRangeEnd() bool {
	if o != nil && !IsNil(o.PortRangeEnd) {
		return true
	}

	return false
}

// SetPortRangeEnd gets a reference to the given int32 and assigns it to the PortRangeEnd field.
func (o *FirewallruleProperties) SetPortRangeEnd(v int32) {
	o.PortRangeEnd = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FirewallruleProperties) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallruleProperties) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FirewallruleProperties) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FirewallruleProperties) SetType(v string) {
	o.Type = &v
}

func (o FirewallruleProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["protocol"] = o.Protocol
	if o.SourceMac.IsSet() {
		toSerialize["sourceMac"] = o.SourceMac.Get()
	}
	if o.IpVersion.IsSet() {
		toSerialize["ipVersion"] = o.IpVersion.Get()
	}
	if o.SourceIp.IsSet() {
		toSerialize["sourceIp"] = o.SourceIp.Get()
	}
	if o.TargetIp.IsSet() {
		toSerialize["targetIp"] = o.TargetIp.Get()
	}
	if o.IcmpCode.IsSet() {
		toSerialize["icmpCode"] = o.IcmpCode.Get()
	}
	if o.IcmpType.IsSet() {
		toSerialize["icmpType"] = o.IcmpType.Get()
	}
	if !IsNil(o.PortRangeStart) {
		toSerialize["portRangeStart"] = o.PortRangeStart
	}
	if !IsNil(o.PortRangeEnd) {
		toSerialize["portRangeEnd"] = o.PortRangeEnd
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
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
