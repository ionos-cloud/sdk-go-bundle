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

// checks if the ApplicationLoadBalancerProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationLoadBalancerProperties{}

// ApplicationLoadBalancerProperties struct for ApplicationLoadBalancerProperties
type ApplicationLoadBalancerProperties struct {
	// The Application Load Balancer name.
	Name string `json:"name"`
	// The ID of the listening (inbound) LAN.
	ListenerLan int32 `json:"listenerLan"`
	// Collection of the Application Load Balancer IP addresses. (Inbound and outbound) IPs of the 'listenerLan' are customer-reserved public IPs for the public load balancers, and private IPs for the private load balancers.
	Ips []string `json:"ips,omitempty"`
	// The ID of the balanced private target LAN (outbound).
	TargetLan int32 `json:"targetLan"`
	// Collection of private IP addresses with the subnet mask of the Application Load Balancer. IPs must contain valid a subnet mask. If no IP is provided, the system will generate an IP with /24 subnet.
	LbPrivateIps []string `json:"lbPrivateIps,omitempty"`
	// Turn logging on and off for this product. Default value is 'false'.
	CentralLogging *bool `json:"centralLogging,omitempty"`
	// Specifies the format of the logs.
	LoggingFormat *string `json:"loggingFormat,omitempty"`
}

// NewApplicationLoadBalancerProperties instantiates a new ApplicationLoadBalancerProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationLoadBalancerProperties(name string, listenerLan int32, targetLan int32) *ApplicationLoadBalancerProperties {
	this := ApplicationLoadBalancerProperties{}

	this.Name = name
	this.ListenerLan = listenerLan
	this.TargetLan = targetLan

	return &this
}

// NewApplicationLoadBalancerPropertiesWithDefaults instantiates a new ApplicationLoadBalancerProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationLoadBalancerPropertiesWithDefaults() *ApplicationLoadBalancerProperties {
	this := ApplicationLoadBalancerProperties{}
	return &this
}

// GetName returns the Name field value
func (o *ApplicationLoadBalancerProperties) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationLoadBalancerProperties) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplicationLoadBalancerProperties) SetName(v string) {
	o.Name = v
}

// GetListenerLan returns the ListenerLan field value
func (o *ApplicationLoadBalancerProperties) GetListenerLan() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ListenerLan
}

// GetListenerLanOk returns a tuple with the ListenerLan field value
// and a boolean to check if the value has been set.
func (o *ApplicationLoadBalancerProperties) GetListenerLanOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListenerLan, true
}

// SetListenerLan sets field value
func (o *ApplicationLoadBalancerProperties) SetListenerLan(v int32) {
	o.ListenerLan = v
}

// GetIps returns the Ips field value if set, zero value otherwise.
func (o *ApplicationLoadBalancerProperties) GetIps() []string {
	if o == nil || IsNil(o.Ips) {
		var ret []string
		return ret
	}
	return o.Ips
}

// GetIpsOk returns a tuple with the Ips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLoadBalancerProperties) GetIpsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ips) {
		return nil, false
	}
	return o.Ips, true
}

// HasIps returns a boolean if a field has been set.
func (o *ApplicationLoadBalancerProperties) HasIps() bool {
	if o != nil && !IsNil(o.Ips) {
		return true
	}

	return false
}

// SetIps gets a reference to the given []string and assigns it to the Ips field.
func (o *ApplicationLoadBalancerProperties) SetIps(v []string) {
	o.Ips = v
}

// GetTargetLan returns the TargetLan field value
func (o *ApplicationLoadBalancerProperties) GetTargetLan() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TargetLan
}

// GetTargetLanOk returns a tuple with the TargetLan field value
// and a boolean to check if the value has been set.
func (o *ApplicationLoadBalancerProperties) GetTargetLanOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetLan, true
}

// SetTargetLan sets field value
func (o *ApplicationLoadBalancerProperties) SetTargetLan(v int32) {
	o.TargetLan = v
}

// GetLbPrivateIps returns the LbPrivateIps field value if set, zero value otherwise.
func (o *ApplicationLoadBalancerProperties) GetLbPrivateIps() []string {
	if o == nil || IsNil(o.LbPrivateIps) {
		var ret []string
		return ret
	}
	return o.LbPrivateIps
}

// GetLbPrivateIpsOk returns a tuple with the LbPrivateIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLoadBalancerProperties) GetLbPrivateIpsOk() ([]string, bool) {
	if o == nil || IsNil(o.LbPrivateIps) {
		return nil, false
	}
	return o.LbPrivateIps, true
}

// HasLbPrivateIps returns a boolean if a field has been set.
func (o *ApplicationLoadBalancerProperties) HasLbPrivateIps() bool {
	if o != nil && !IsNil(o.LbPrivateIps) {
		return true
	}

	return false
}

// SetLbPrivateIps gets a reference to the given []string and assigns it to the LbPrivateIps field.
func (o *ApplicationLoadBalancerProperties) SetLbPrivateIps(v []string) {
	o.LbPrivateIps = v
}

// GetCentralLogging returns the CentralLogging field value if set, zero value otherwise.
func (o *ApplicationLoadBalancerProperties) GetCentralLogging() bool {
	if o == nil || IsNil(o.CentralLogging) {
		var ret bool
		return ret
	}
	return *o.CentralLogging
}

// GetCentralLoggingOk returns a tuple with the CentralLogging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLoadBalancerProperties) GetCentralLoggingOk() (*bool, bool) {
	if o == nil || IsNil(o.CentralLogging) {
		return nil, false
	}
	return o.CentralLogging, true
}

// HasCentralLogging returns a boolean if a field has been set.
func (o *ApplicationLoadBalancerProperties) HasCentralLogging() bool {
	if o != nil && !IsNil(o.CentralLogging) {
		return true
	}

	return false
}

// SetCentralLogging gets a reference to the given bool and assigns it to the CentralLogging field.
func (o *ApplicationLoadBalancerProperties) SetCentralLogging(v bool) {
	o.CentralLogging = &v
}

// GetLoggingFormat returns the LoggingFormat field value if set, zero value otherwise.
func (o *ApplicationLoadBalancerProperties) GetLoggingFormat() string {
	if o == nil || IsNil(o.LoggingFormat) {
		var ret string
		return ret
	}
	return *o.LoggingFormat
}

// GetLoggingFormatOk returns a tuple with the LoggingFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLoadBalancerProperties) GetLoggingFormatOk() (*string, bool) {
	if o == nil || IsNil(o.LoggingFormat) {
		return nil, false
	}
	return o.LoggingFormat, true
}

// HasLoggingFormat returns a boolean if a field has been set.
func (o *ApplicationLoadBalancerProperties) HasLoggingFormat() bool {
	if o != nil && !IsNil(o.LoggingFormat) {
		return true
	}

	return false
}

// SetLoggingFormat gets a reference to the given string and assigns it to the LoggingFormat field.
func (o *ApplicationLoadBalancerProperties) SetLoggingFormat(v string) {
	o.LoggingFormat = &v
}

func (o ApplicationLoadBalancerProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["listenerLan"] = o.ListenerLan
	if !IsNil(o.Ips) {
		toSerialize["ips"] = o.Ips
	}
	toSerialize["targetLan"] = o.TargetLan
	if !IsNil(o.LbPrivateIps) {
		toSerialize["lbPrivateIps"] = o.LbPrivateIps
	}
	if !IsNil(o.CentralLogging) {
		toSerialize["centralLogging"] = o.CentralLogging
	}
	if !IsNil(o.LoggingFormat) {
		toSerialize["loggingFormat"] = o.LoggingFormat
	}
	return toSerialize, nil
}

type NullableApplicationLoadBalancerProperties struct {
	value *ApplicationLoadBalancerProperties
	isSet bool
}

func (v NullableApplicationLoadBalancerProperties) Get() *ApplicationLoadBalancerProperties {
	return v.value
}

func (v *NullableApplicationLoadBalancerProperties) Set(val *ApplicationLoadBalancerProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationLoadBalancerProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationLoadBalancerProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationLoadBalancerProperties(val *ApplicationLoadBalancerProperties) *NullableApplicationLoadBalancerProperties {
	return &NullableApplicationLoadBalancerProperties{value: val, isSet: true}
}

func (v NullableApplicationLoadBalancerProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationLoadBalancerProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
