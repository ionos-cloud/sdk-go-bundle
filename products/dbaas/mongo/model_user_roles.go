/*
 * IONOS DBaaS MongoDB REST API
 *
 * With IONOS Cloud Database as a Service, you have the ability to quickly set up and manage a MongoDB database. You can also delete clusters, manage backups and users via the API.  MongoDB is an open source, cross-platform, document-oriented database program. Classified as a NoSQL database program, it uses JSON-like documents with optional schemas.  The MongoDB API allows you to create additional database clusters or modify existing ones. Both tools, the Data Center Designer (DCD) and the API use the same concepts consistently and are well suited for smooth and intuitive use.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongo

import (
	"encoding/json"
)

// checks if the UserRoles type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserRoles{}

// UserRoles a list of mongodb user role.
type UserRoles struct {
	Role     *string `json:"role,omitempty"`
	Database *string `json:"database,omitempty"`
}

// NewUserRoles instantiates a new UserRoles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserRoles() *UserRoles {
	this := UserRoles{}

	return &this
}

// NewUserRolesWithDefaults instantiates a new UserRoles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserRolesWithDefaults() *UserRoles {
	this := UserRoles{}
	return &this
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *UserRoles) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRoles) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *UserRoles) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *UserRoles) SetRole(v string) {
	o.Role = &v
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *UserRoles) GetDatabase() string {
	if o == nil || IsNil(o.Database) {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRoles) GetDatabaseOk() (*string, bool) {
	if o == nil || IsNil(o.Database) {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *UserRoles) HasDatabase() bool {
	if o != nil && !IsNil(o.Database) {
		return true
	}

	return false
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *UserRoles) SetDatabase(v string) {
	o.Database = &v
}

func (o UserRoles) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.Database) {
		toSerialize["database"] = o.Database
	}
	return toSerialize, nil
}

type NullableUserRoles struct {
	value *UserRoles
	isSet bool
}

func (v NullableUserRoles) Get() *UserRoles {
	return v.value
}

func (v *NullableUserRoles) Set(val *UserRoles) {
	v.value = val
	v.isSet = true
}

func (v NullableUserRoles) IsSet() bool {
	return v.isSet
}

func (v *NullableUserRoles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserRoles(val *UserRoles) *NullableUserRoles {
	return &NullableUserRoles{value: val, isSet: true}
}

func (v NullableUserRoles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserRoles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
