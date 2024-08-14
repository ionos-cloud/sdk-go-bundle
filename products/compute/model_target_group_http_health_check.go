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

// checks if the TargetGroupHttpHealthCheck type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TargetGroupHttpHealthCheck{}

// TargetGroupHttpHealthCheck struct for TargetGroupHttpHealthCheck
type TargetGroupHttpHealthCheck struct {
	// The destination URL for HTTP the health check; the default is '/'.
	Path *string `json:"path,omitempty"`
	// The method used for the health check request.
	Method *string `json:"method,omitempty"`
	// Specify the target's response type to match ALB's request.
	MatchType string `json:"matchType"`
	// The response returned by the request. It can be a status code or a response body depending on the definition of 'matchType'.
	Response string `json:"response"`
	// Specifies whether to use a regular expression to parse the response body; the default value is 'FALSE'.  By using regular expressions, you can flexibly customize the expected response from a healthy server.
	Regex *bool `json:"regex,omitempty"`
	// Specifies whether to negate an individual entry; the default value is 'FALSE'.
	Negate *bool `json:"negate,omitempty"`
}

// NewTargetGroupHttpHealthCheck instantiates a new TargetGroupHttpHealthCheck object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetGroupHttpHealthCheck(matchType string, response string) *TargetGroupHttpHealthCheck {
	this := TargetGroupHttpHealthCheck{}

	this.MatchType = matchType
	this.Response = response

	return &this
}

// NewTargetGroupHttpHealthCheckWithDefaults instantiates a new TargetGroupHttpHealthCheck object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetGroupHttpHealthCheckWithDefaults() *TargetGroupHttpHealthCheck {
	this := TargetGroupHttpHealthCheck{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *TargetGroupHttpHealthCheck) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetGroupHttpHealthCheck) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *TargetGroupHttpHealthCheck) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *TargetGroupHttpHealthCheck) SetPath(v string) {
	o.Path = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *TargetGroupHttpHealthCheck) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetGroupHttpHealthCheck) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *TargetGroupHttpHealthCheck) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *TargetGroupHttpHealthCheck) SetMethod(v string) {
	o.Method = &v
}

// GetMatchType returns the MatchType field value
func (o *TargetGroupHttpHealthCheck) GetMatchType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value
// and a boolean to check if the value has been set.
func (o *TargetGroupHttpHealthCheck) GetMatchTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchType, true
}

// SetMatchType sets field value
func (o *TargetGroupHttpHealthCheck) SetMatchType(v string) {
	o.MatchType = v
}

// GetResponse returns the Response field value
func (o *TargetGroupHttpHealthCheck) GetResponse() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Response
}

// GetResponseOk returns a tuple with the Response field value
// and a boolean to check if the value has been set.
func (o *TargetGroupHttpHealthCheck) GetResponseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Response, true
}

// SetResponse sets field value
func (o *TargetGroupHttpHealthCheck) SetResponse(v string) {
	o.Response = v
}

// GetRegex returns the Regex field value if set, zero value otherwise.
func (o *TargetGroupHttpHealthCheck) GetRegex() bool {
	if o == nil || IsNil(o.Regex) {
		var ret bool
		return ret
	}
	return *o.Regex
}

// GetRegexOk returns a tuple with the Regex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetGroupHttpHealthCheck) GetRegexOk() (*bool, bool) {
	if o == nil || IsNil(o.Regex) {
		return nil, false
	}
	return o.Regex, true
}

// HasRegex returns a boolean if a field has been set.
func (o *TargetGroupHttpHealthCheck) HasRegex() bool {
	if o != nil && !IsNil(o.Regex) {
		return true
	}

	return false
}

// SetRegex gets a reference to the given bool and assigns it to the Regex field.
func (o *TargetGroupHttpHealthCheck) SetRegex(v bool) {
	o.Regex = &v
}

// GetNegate returns the Negate field value if set, zero value otherwise.
func (o *TargetGroupHttpHealthCheck) GetNegate() bool {
	if o == nil || IsNil(o.Negate) {
		var ret bool
		return ret
	}
	return *o.Negate
}

// GetNegateOk returns a tuple with the Negate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetGroupHttpHealthCheck) GetNegateOk() (*bool, bool) {
	if o == nil || IsNil(o.Negate) {
		return nil, false
	}
	return o.Negate, true
}

// HasNegate returns a boolean if a field has been set.
func (o *TargetGroupHttpHealthCheck) HasNegate() bool {
	if o != nil && !IsNil(o.Negate) {
		return true
	}

	return false
}

// SetNegate gets a reference to the given bool and assigns it to the Negate field.
func (o *TargetGroupHttpHealthCheck) SetNegate(v bool) {
	o.Negate = &v
}

func (o TargetGroupHttpHealthCheck) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TargetGroupHttpHealthCheck) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsZero(o.MatchType) {
		toSerialize["matchType"] = o.MatchType
	}
	if !IsZero(o.Response) {
		toSerialize["response"] = o.Response
	}
	if !IsNil(o.Regex) {
		toSerialize["regex"] = o.Regex
	}
	if !IsNil(o.Negate) {
		toSerialize["negate"] = o.Negate
	}
	return toSerialize, nil
}

type NullableTargetGroupHttpHealthCheck struct {
	value *TargetGroupHttpHealthCheck
	isSet bool
}

func (v NullableTargetGroupHttpHealthCheck) Get() *TargetGroupHttpHealthCheck {
	return v.value
}

func (v *NullableTargetGroupHttpHealthCheck) Set(val *TargetGroupHttpHealthCheck) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetGroupHttpHealthCheck) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetGroupHttpHealthCheck) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetGroupHttpHealthCheck(val *TargetGroupHttpHealthCheck) *NullableTargetGroupHttpHealthCheck {
	return &NullableTargetGroupHttpHealthCheck{value: val, isSet: true}
}

func (v NullableTargetGroupHttpHealthCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetGroupHttpHealthCheck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
