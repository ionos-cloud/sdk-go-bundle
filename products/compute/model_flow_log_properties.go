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

// checks if the FlowLogProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlowLogProperties{}

// FlowLogProperties struct for FlowLogProperties
type FlowLogProperties struct {
	// The resource name.
	Name string `json:"name"`
	// Specifies the traffic action pattern.
	Action string `json:"action"`
	// Specifies the traffic direction pattern.
	Direction string `json:"direction"`
	// The S3 bucket name of an existing IONOS Cloud S3 bucket.
	Bucket string `json:"bucket"`
}

// NewFlowLogProperties instantiates a new FlowLogProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlowLogProperties(name string, action string, direction string, bucket string) *FlowLogProperties {
	this := FlowLogProperties{}

	this.Name = name
	this.Action = action
	this.Direction = direction
	this.Bucket = bucket

	return &this
}

// NewFlowLogPropertiesWithDefaults instantiates a new FlowLogProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowLogPropertiesWithDefaults() *FlowLogProperties {
	this := FlowLogProperties{}
	return &this
}

// GetName returns the Name field value
func (o *FlowLogProperties) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FlowLogProperties) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FlowLogProperties) SetName(v string) {
	o.Name = v
}

// GetAction returns the Action field value
func (o *FlowLogProperties) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *FlowLogProperties) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *FlowLogProperties) SetAction(v string) {
	o.Action = v
}

// GetDirection returns the Direction field value
func (o *FlowLogProperties) GetDirection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value
// and a boolean to check if the value has been set.
func (o *FlowLogProperties) GetDirectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Direction, true
}

// SetDirection sets field value
func (o *FlowLogProperties) SetDirection(v string) {
	o.Direction = v
}

// GetBucket returns the Bucket field value
func (o *FlowLogProperties) GetBucket() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value
// and a boolean to check if the value has been set.
func (o *FlowLogProperties) GetBucketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bucket, true
}

// SetBucket sets field value
func (o *FlowLogProperties) SetBucket(v string) {
	o.Bucket = v
}

func (o FlowLogProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["action"] = o.Action
	toSerialize["direction"] = o.Direction
	toSerialize["bucket"] = o.Bucket
	return toSerialize, nil
}

type NullableFlowLogProperties struct {
	value *FlowLogProperties
	isSet bool
}

func (v NullableFlowLogProperties) Get() *FlowLogProperties {
	return v.value
}

func (v *NullableFlowLogProperties) Set(val *FlowLogProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowLogProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowLogProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowLogProperties(val *FlowLogProperties) *NullableFlowLogProperties {
	return &NullableFlowLogProperties{value: val, isSet: true}
}

func (v NullableFlowLogProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowLogProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
