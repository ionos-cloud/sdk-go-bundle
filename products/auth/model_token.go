/*
 * Auth API
 *
 * Use the Auth API to manage tokens for secure access to IONOS Cloud  APIs (Auth API, Cloud API, Reseller API, Activity Log API, and others).
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package auth

import (
	"encoding/json"
)

// Token struct for Token
type Token struct {
	// Identifier of the token.
	Id *string `json:"id"`
	// Hypertext REFerence for the specified token.
	Href *string `json:"href"`
	// The date the token was generated.
	CreatedDate *string `json:"createdDate"`
	// The date the token will expire.
	ExpirationDate *string `json:"expirationDate"`
}

// NewToken instantiates a new Token object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToken(id string, href string, createdDate string, expirationDate string) *Token {
	this := Token{}

	this.Id = &id
	this.Href = &href
	this.CreatedDate = &createdDate
	this.ExpirationDate = &expirationDate

	return &this
}

// NewTokenWithDefaults instantiates a new Token object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenWithDefaults() *Token {
	this := Token{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Token) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Token) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *Token) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *Token) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Token) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Token) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Href, true
}

// SetHref sets field value
func (o *Token) SetHref(v string) {

	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *Token) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// GetCreatedDate returns the CreatedDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Token) GetCreatedDate() *string {
	if o == nil {
		return nil
	}

	return o.CreatedDate

}

// GetCreatedDateOk returns a tuple with the CreatedDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Token) GetCreatedDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.CreatedDate, true
}

// SetCreatedDate sets field value
func (o *Token) SetCreatedDate(v string) {

	o.CreatedDate = &v

}

// HasCreatedDate returns a boolean if a field has been set.
func (o *Token) HasCreatedDate() bool {
	if o != nil && o.CreatedDate != nil {
		return true
	}

	return false
}

// GetExpirationDate returns the ExpirationDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Token) GetExpirationDate() *string {
	if o == nil {
		return nil
	}

	return o.ExpirationDate

}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Token) GetExpirationDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.ExpirationDate, true
}

// SetExpirationDate sets field value
func (o *Token) SetExpirationDate(v string) {

	o.ExpirationDate = &v

}

// HasExpirationDate returns a boolean if a field has been set.
func (o *Token) HasExpirationDate() bool {
	if o != nil && o.ExpirationDate != nil {
		return true
	}

	return false
}

func (o Token) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	if o.Href != nil {
		toSerialize["href"] = o.Href
	}

	if o.CreatedDate != nil {
		toSerialize["createdDate"] = o.CreatedDate
	}

	if o.ExpirationDate != nil {
		toSerialize["expirationDate"] = o.ExpirationDate
	}

	return json.Marshal(toSerialize)
}

type NullableToken struct {
	value *Token
	isSet bool
}

func (v NullableToken) Get() *Token {
	return v.value
}

func (v *NullableToken) Set(val *Token) {
	v.value = val
	v.isSet = true
}

func (v NullableToken) IsSet() bool {
	return v.isSet
}

func (v *NullableToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToken(val *Token) *NullableToken {
	return &NullableToken{value: val, isSet: true}
}

func (v NullableToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
