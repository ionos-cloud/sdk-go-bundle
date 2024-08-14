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

// checks if the UserProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserProperties{}

// UserProperties struct for UserProperties
type UserProperties struct {
	// The first name of the user.
	Firstname *string `json:"firstname,omitempty"`
	// The last name of the user.
	Lastname *string `json:"lastname,omitempty"`
	// The email address of the user.
	Email *string `json:"email,omitempty"`
	// Indicates if the user has admin rights.
	Administrator *bool `json:"administrator,omitempty"`
	// Indicates if secure authentication should be forced on the user.
	ForceSecAuth *bool `json:"forceSecAuth,omitempty"`
	// Indicates if secure authentication is active for the user.
	SecAuthActive *bool `json:"secAuthActive,omitempty"`
	// Canonical (S3) ID of the user for a given identity.
	S3CanonicalUserId *string `json:"s3CanonicalUserId,omitempty"`
	// Indicates if the user is active.
	Active *bool `json:"active,omitempty"`
}

// NewUserProperties instantiates a new UserProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserProperties() *UserProperties {
	this := UserProperties{}

	return &this
}

// NewUserPropertiesWithDefaults instantiates a new UserProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPropertiesWithDefaults() *UserProperties {
	this := UserProperties{}
	return &this
}

// GetFirstname returns the Firstname field value if set, zero value otherwise.
func (o *UserProperties) GetFirstname() string {
	if o == nil || IsNil(o.Firstname) {
		var ret string
		return ret
	}
	return *o.Firstname
}

// GetFirstnameOk returns a tuple with the Firstname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProperties) GetFirstnameOk() (*string, bool) {
	if o == nil || IsNil(o.Firstname) {
		return nil, false
	}
	return o.Firstname, true
}

// HasFirstname returns a boolean if a field has been set.
func (o *UserProperties) HasFirstname() bool {
	if o != nil && !IsNil(o.Firstname) {
		return true
	}

	return false
}

// SetFirstname gets a reference to the given string and assigns it to the Firstname field.
func (o *UserProperties) SetFirstname(v string) {
	o.Firstname = &v
}

// GetLastname returns the Lastname field value if set, zero value otherwise.
func (o *UserProperties) GetLastname() string {
	if o == nil || IsNil(o.Lastname) {
		var ret string
		return ret
	}
	return *o.Lastname
}

// GetLastnameOk returns a tuple with the Lastname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProperties) GetLastnameOk() (*string, bool) {
	if o == nil || IsNil(o.Lastname) {
		return nil, false
	}
	return o.Lastname, true
}

// HasLastname returns a boolean if a field has been set.
func (o *UserProperties) HasLastname() bool {
	if o != nil && !IsNil(o.Lastname) {
		return true
	}

	return false
}

// SetLastname gets a reference to the given string and assigns it to the Lastname field.
func (o *UserProperties) SetLastname(v string) {
	o.Lastname = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserProperties) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProperties) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserProperties) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserProperties) SetEmail(v string) {
	o.Email = &v
}

// GetAdministrator returns the Administrator field value if set, zero value otherwise.
func (o *UserProperties) GetAdministrator() bool {
	if o == nil || IsNil(o.Administrator) {
		var ret bool
		return ret
	}
	return *o.Administrator
}

// GetAdministratorOk returns a tuple with the Administrator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProperties) GetAdministratorOk() (*bool, bool) {
	if o == nil || IsNil(o.Administrator) {
		return nil, false
	}
	return o.Administrator, true
}

// HasAdministrator returns a boolean if a field has been set.
func (o *UserProperties) HasAdministrator() bool {
	if o != nil && !IsNil(o.Administrator) {
		return true
	}

	return false
}

// SetAdministrator gets a reference to the given bool and assigns it to the Administrator field.
func (o *UserProperties) SetAdministrator(v bool) {
	o.Administrator = &v
}

// GetForceSecAuth returns the ForceSecAuth field value if set, zero value otherwise.
func (o *UserProperties) GetForceSecAuth() bool {
	if o == nil || IsNil(o.ForceSecAuth) {
		var ret bool
		return ret
	}
	return *o.ForceSecAuth
}

// GetForceSecAuthOk returns a tuple with the ForceSecAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProperties) GetForceSecAuthOk() (*bool, bool) {
	if o == nil || IsNil(o.ForceSecAuth) {
		return nil, false
	}
	return o.ForceSecAuth, true
}

// HasForceSecAuth returns a boolean if a field has been set.
func (o *UserProperties) HasForceSecAuth() bool {
	if o != nil && !IsNil(o.ForceSecAuth) {
		return true
	}

	return false
}

// SetForceSecAuth gets a reference to the given bool and assigns it to the ForceSecAuth field.
func (o *UserProperties) SetForceSecAuth(v bool) {
	o.ForceSecAuth = &v
}

// GetSecAuthActive returns the SecAuthActive field value if set, zero value otherwise.
func (o *UserProperties) GetSecAuthActive() bool {
	if o == nil || IsNil(o.SecAuthActive) {
		var ret bool
		return ret
	}
	return *o.SecAuthActive
}

// GetSecAuthActiveOk returns a tuple with the SecAuthActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProperties) GetSecAuthActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.SecAuthActive) {
		return nil, false
	}
	return o.SecAuthActive, true
}

// HasSecAuthActive returns a boolean if a field has been set.
func (o *UserProperties) HasSecAuthActive() bool {
	if o != nil && !IsNil(o.SecAuthActive) {
		return true
	}

	return false
}

// SetSecAuthActive gets a reference to the given bool and assigns it to the SecAuthActive field.
func (o *UserProperties) SetSecAuthActive(v bool) {
	o.SecAuthActive = &v
}

// GetS3CanonicalUserId returns the S3CanonicalUserId field value if set, zero value otherwise.
func (o *UserProperties) GetS3CanonicalUserId() string {
	if o == nil || IsNil(o.S3CanonicalUserId) {
		var ret string
		return ret
	}
	return *o.S3CanonicalUserId
}

// GetS3CanonicalUserIdOk returns a tuple with the S3CanonicalUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProperties) GetS3CanonicalUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.S3CanonicalUserId) {
		return nil, false
	}
	return o.S3CanonicalUserId, true
}

// HasS3CanonicalUserId returns a boolean if a field has been set.
func (o *UserProperties) HasS3CanonicalUserId() bool {
	if o != nil && !IsNil(o.S3CanonicalUserId) {
		return true
	}

	return false
}

// SetS3CanonicalUserId gets a reference to the given string and assigns it to the S3CanonicalUserId field.
func (o *UserProperties) SetS3CanonicalUserId(v string) {
	o.S3CanonicalUserId = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *UserProperties) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProperties) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *UserProperties) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *UserProperties) SetActive(v bool) {
	o.Active = &v
}

func (o UserProperties) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Firstname) {
		toSerialize["firstname"] = o.Firstname
	}
	if !IsNil(o.Lastname) {
		toSerialize["lastname"] = o.Lastname
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Administrator) {
		toSerialize["administrator"] = o.Administrator
	}
	if !IsNil(o.ForceSecAuth) {
		toSerialize["forceSecAuth"] = o.ForceSecAuth
	}
	if !IsNil(o.SecAuthActive) {
		toSerialize["secAuthActive"] = o.SecAuthActive
	}
	if !IsNil(o.S3CanonicalUserId) {
		toSerialize["s3CanonicalUserId"] = o.S3CanonicalUserId
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	return toSerialize, nil
}

type NullableUserProperties struct {
	value *UserProperties
	isSet bool
}

func (v NullableUserProperties) Get() *UserProperties {
	return v.value
}

func (v *NullableUserProperties) Set(val *UserProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableUserProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableUserProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserProperties(val *UserProperties) *NullableUserProperties {
	return &NullableUserProperties{value: val, isSet: true}
}

func (v NullableUserProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
