/*
 * Kafka as a Service API
 *
 * An managed Apache Kafka cluster is designed to be highly fault-tolerant and scalable, allowing large volumes of data to be ingested, stored, and processed in real-time. By distributing data across multiple brokers, Kafka achieves high throughput and low latency, making it suitable for applications requiring real-time data processing and analytics.
 *
 * API version: 1.7.1
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafka

import (
	"encoding/json"
)

// checks if the ClusterReadListAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterReadListAllOf{}

// ClusterReadListAllOf struct for ClusterReadListAllOf
type ClusterReadListAllOf struct {
	// ID of the list of Cluster resources.
	Id string `json:"id"`
	// The type of the resource.
	Type string `json:"type"`
	// The URL of the list of Cluster resources.
	Href string `json:"href"`
	// The list of Cluster resources.
	Items []ClusterRead `json:"items,omitempty"`
}

// NewClusterReadListAllOf instantiates a new ClusterReadListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterReadListAllOf(id string, type_ string, href string) *ClusterReadListAllOf {
	this := ClusterReadListAllOf{}

	this.Id = id
	this.Type = type_
	this.Href = href

	return &this
}

// NewClusterReadListAllOfWithDefaults instantiates a new ClusterReadListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterReadListAllOfWithDefaults() *ClusterReadListAllOf {
	this := ClusterReadListAllOf{}
	return &this
}

// GetId returns the Id field value
func (o *ClusterReadListAllOf) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ClusterReadListAllOf) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ClusterReadListAllOf) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *ClusterReadListAllOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ClusterReadListAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ClusterReadListAllOf) SetType(v string) {
	o.Type = v
}

// GetHref returns the Href field value
func (o *ClusterReadListAllOf) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *ClusterReadListAllOf) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *ClusterReadListAllOf) SetHref(v string) {
	o.Href = v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ClusterReadListAllOf) GetItems() []ClusterRead {
	if o == nil || IsNil(o.Items) {
		var ret []ClusterRead
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterReadListAllOf) GetItemsOk() ([]ClusterRead, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ClusterReadListAllOf) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ClusterRead and assigns it to the Items field.
func (o *ClusterReadListAllOf) SetItems(v []ClusterRead) {
	o.Items = v
}

func (o ClusterReadListAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["href"] = o.Href
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableClusterReadListAllOf struct {
	value *ClusterReadListAllOf
	isSet bool
}

func (v NullableClusterReadListAllOf) Get() *ClusterReadListAllOf {
	return v.value
}

func (v *NullableClusterReadListAllOf) Set(val *ClusterReadListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterReadListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterReadListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterReadListAllOf(val *ClusterReadListAllOf) *NullableClusterReadListAllOf {
	return &NullableClusterReadListAllOf{value: val, isSet: true}
}

func (v NullableClusterReadListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterReadListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
