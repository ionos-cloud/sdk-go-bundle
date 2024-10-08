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

	"time"
)

// checks if the UserMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserMetadata{}

// UserMetadata struct for UserMetadata
type UserMetadata struct {
	// Resource's Entity Tag as defined in http://www.w3.org/Protocols/rfc2616/rfc2616-sec3.html#sec3.11  Entity Tag is also added as an 'ETag response header to requests which don't use 'depth' parameter.
	Etag *string `json:"etag,omitempty"`
	// The time the user was created.
	CreatedDate *IonosTime `json:"createdDate,omitempty"`
	// The user who has created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`
	// The unique ID of the user who created the resource.
	CreatedByUserId *string `json:"createdByUserId,omitempty"`
	// The last time the resource was modified.
	LastModifiedDate *IonosTime `json:"lastModifiedDate,omitempty"`
	// The user who last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`
	// The unique ID of the user who last modified the resource.
	LastModifiedByUserId *string `json:"lastModifiedByUserId,omitempty"`
	// The time of the last login by the user.
	LastLogin *IonosTime `json:"lastLogin,omitempty"`
}

// NewUserMetadata instantiates a new UserMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserMetadata() *UserMetadata {
	this := UserMetadata{}

	return &this
}

// NewUserMetadataWithDefaults instantiates a new UserMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserMetadataWithDefaults() *UserMetadata {
	this := UserMetadata{}
	return &this
}

// GetEtag returns the Etag field value if set, zero value otherwise.
func (o *UserMetadata) GetEtag() string {
	if o == nil || IsNil(o.Etag) {
		var ret string
		return ret
	}
	return *o.Etag
}

// GetEtagOk returns a tuple with the Etag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserMetadata) GetEtagOk() (*string, bool) {
	if o == nil || IsNil(o.Etag) {
		return nil, false
	}
	return o.Etag, true
}

// HasEtag returns a boolean if a field has been set.
func (o *UserMetadata) HasEtag() bool {
	if o != nil && !IsNil(o.Etag) {
		return true
	}

	return false
}

// SetEtag gets a reference to the given string and assigns it to the Etag field.
func (o *UserMetadata) SetEtag(v string) {
	o.Etag = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *UserMetadata) GetCreatedDate() time.Time {
	if o == nil || IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return o.CreatedDate.Time
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserMetadata) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return &o.CreatedDate.Time, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *UserMetadata) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *UserMetadata) SetCreatedDate(v time.Time) {
	o.CreatedDate = &IonosTime{v}
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *UserMetadata) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserMetadata) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *UserMetadata) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *UserMetadata) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedByUserId returns the CreatedByUserId field value if set, zero value otherwise.
func (o *UserMetadata) GetCreatedByUserId() string {
	if o == nil || IsNil(o.CreatedByUserId) {
		var ret string
		return ret
	}
	return *o.CreatedByUserId
}

// GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserMetadata) GetCreatedByUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedByUserId) {
		return nil, false
	}
	return o.CreatedByUserId, true
}

// HasCreatedByUserId returns a boolean if a field has been set.
func (o *UserMetadata) HasCreatedByUserId() bool {
	if o != nil && !IsNil(o.CreatedByUserId) {
		return true
	}

	return false
}

// SetCreatedByUserId gets a reference to the given string and assigns it to the CreatedByUserId field.
func (o *UserMetadata) SetCreatedByUserId(v string) {
	o.CreatedByUserId = &v
}

// GetLastModifiedDate returns the LastModifiedDate field value if set, zero value otherwise.
func (o *UserMetadata) GetLastModifiedDate() time.Time {
	if o == nil || IsNil(o.LastModifiedDate) {
		var ret time.Time
		return ret
	}
	return o.LastModifiedDate.Time
}

// GetLastModifiedDateOk returns a tuple with the LastModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserMetadata) GetLastModifiedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModifiedDate) {
		return nil, false
	}
	return &o.LastModifiedDate.Time, true
}

// HasLastModifiedDate returns a boolean if a field has been set.
func (o *UserMetadata) HasLastModifiedDate() bool {
	if o != nil && !IsNil(o.LastModifiedDate) {
		return true
	}

	return false
}

// SetLastModifiedDate gets a reference to the given time.Time and assigns it to the LastModifiedDate field.
func (o *UserMetadata) SetLastModifiedDate(v time.Time) {
	o.LastModifiedDate = &IonosTime{v}
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise.
func (o *UserMetadata) GetLastModifiedBy() string {
	if o == nil || IsNil(o.LastModifiedBy) {
		var ret string
		return ret
	}
	return *o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserMetadata) GetLastModifiedByOk() (*string, bool) {
	if o == nil || IsNil(o.LastModifiedBy) {
		return nil, false
	}
	return o.LastModifiedBy, true
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *UserMetadata) HasLastModifiedBy() bool {
	if o != nil && !IsNil(o.LastModifiedBy) {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given string and assigns it to the LastModifiedBy field.
func (o *UserMetadata) SetLastModifiedBy(v string) {
	o.LastModifiedBy = &v
}

// GetLastModifiedByUserId returns the LastModifiedByUserId field value if set, zero value otherwise.
func (o *UserMetadata) GetLastModifiedByUserId() string {
	if o == nil || IsNil(o.LastModifiedByUserId) {
		var ret string
		return ret
	}
	return *o.LastModifiedByUserId
}

// GetLastModifiedByUserIdOk returns a tuple with the LastModifiedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserMetadata) GetLastModifiedByUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.LastModifiedByUserId) {
		return nil, false
	}
	return o.LastModifiedByUserId, true
}

// HasLastModifiedByUserId returns a boolean if a field has been set.
func (o *UserMetadata) HasLastModifiedByUserId() bool {
	if o != nil && !IsNil(o.LastModifiedByUserId) {
		return true
	}

	return false
}

// SetLastModifiedByUserId gets a reference to the given string and assigns it to the LastModifiedByUserId field.
func (o *UserMetadata) SetLastModifiedByUserId(v string) {
	o.LastModifiedByUserId = &v
}

// GetLastLogin returns the LastLogin field value if set, zero value otherwise.
func (o *UserMetadata) GetLastLogin() time.Time {
	if o == nil || IsNil(o.LastLogin) {
		var ret time.Time
		return ret
	}
	return o.LastLogin.Time
}

// GetLastLoginOk returns a tuple with the LastLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserMetadata) GetLastLoginOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastLogin) {
		return nil, false
	}
	return &o.LastLogin.Time, true
}

// HasLastLogin returns a boolean if a field has been set.
func (o *UserMetadata) HasLastLogin() bool {
	if o != nil && !IsNil(o.LastLogin) {
		return true
	}

	return false
}

// SetLastLogin gets a reference to the given time.Time and assigns it to the LastLogin field.
func (o *UserMetadata) SetLastLogin(v time.Time) {
	o.LastLogin = &IonosTime{v}
}

func (o UserMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Etag) {
		toSerialize["etag"] = o.Etag
	}
	if !IsNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.CreatedByUserId) {
		toSerialize["createdByUserId"] = o.CreatedByUserId
	}
	if !IsNil(o.LastModifiedDate) {
		toSerialize["lastModifiedDate"] = o.LastModifiedDate
	}
	if !IsNil(o.LastModifiedBy) {
		toSerialize["lastModifiedBy"] = o.LastModifiedBy
	}
	if !IsNil(o.LastModifiedByUserId) {
		toSerialize["lastModifiedByUserId"] = o.LastModifiedByUserId
	}
	if !IsNil(o.LastLogin) {
		toSerialize["lastLogin"] = o.LastLogin
	}
	return toSerialize, nil
}

type NullableUserMetadata struct {
	value *UserMetadata
	isSet bool
}

func (v NullableUserMetadata) Get() *UserMetadata {
	return v.value
}

func (v *NullableUserMetadata) Set(val *UserMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableUserMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableUserMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserMetadata(val *UserMetadata) *NullableUserMetadata {
	return &NullableUserMetadata{value: val, isSet: true}
}

func (v NullableUserMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
