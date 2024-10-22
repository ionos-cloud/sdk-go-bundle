/*
 * Container Registry service
 *
 * ## Overview Container Registry service enables IONOS clients to manage docker and OCI compliant registries for use by their managed Kubernetes clusters. Use a Container Registry to ensure you have a privately accessed registry to efficiently support image pulls. ## Changelog ### 1.1.0  - Added new endpoints for Repositories  - Added new endpoints for Artifacts  - Added new endpoints for Vulnerabilities  - Added registry vulnerabilityScanning feature ### 1.2.0  - Added registry `apiSubnetAllowList` ### 1.2.1  - Amended `apiSubnetAllowList` Regex
 *
 * API version: 1.2.1
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package containerregistry

import (
	"encoding/json"

	"time"
)

// checks if the ArtifactMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArtifactMetadata{}

// ArtifactMetadata struct for ArtifactMetadata
type ArtifactMetadata struct {
	// The ISO 8601 creation timestamp.
	CreatedDate *IonosTime `json:"createdDate,omitempty"`
	// Unique name of the identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`
	// Unique id of the identity that created the resource.
	CreatedByUserId *string `json:"createdByUserId,omitempty"`
	// The ISO 8601 modified timestamp.
	LastModifiedDate *IonosTime `json:"lastModifiedDate,omitempty"`
	// Unique name of the identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`
	// Unique id of the identity that last modified the resource.
	LastModifiedByUserId *string `json:"lastModifiedByUserId,omitempty"`
	// Unique name of the resource.
	ResourceURN *string `json:"resourceURN,omitempty"`
	// The date and time the artifact was last pushed
	LastPushedAt *IonosTime `json:"lastPushedAt"`
	// The date and time the artifact was last pulled
	LastPulledAt *IonosTime `json:"lastPulledAt,omitempty"`
	// The date and time the artifact was last scanned
	LastScannedAt *IonosTime `json:"lastScannedAt,omitempty"`
	// The number of times the artifact was pushed
	PushCount int64 `json:"pushCount"`
	// The number of times the artifact was pulled
	PullCount int64 `json:"pullCount"`
	// The CVSS vulnerability severity rating
	VulnMaxSeverity *string `json:"vulnMaxSeverity,omitempty"`
	// The total CVSS score of all vulnerabilities of the artifact
	VulnTotalScore *float32 `json:"vulnTotalScore,omitempty"`
	// The total number of vulnerabilities of the artifact
	VulnTotalCount *int64 `json:"vulnTotalCount,omitempty"`
	// The number of fixable vulnerabilities of the artifact
	VulnFixableCount *int64 `json:"vulnFixableCount,omitempty"`
}

// NewArtifactMetadata instantiates a new ArtifactMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtifactMetadata(lastPushedAt time.Time, pushCount int64, pullCount int64) *ArtifactMetadata {
	this := ArtifactMetadata{}

	this.LastPushedAt = &IonosTime{lastPushedAt}
	this.PushCount = pushCount
	this.PullCount = pullCount

	return &this
}

// NewArtifactMetadataWithDefaults instantiates a new ArtifactMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtifactMetadataWithDefaults() *ArtifactMetadata {
	this := ArtifactMetadata{}
	return &this
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *ArtifactMetadata) GetCreatedDate() time.Time {
	if o == nil || IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return o.CreatedDate.Time
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactMetadata) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return &o.CreatedDate.Time, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *ArtifactMetadata) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *ArtifactMetadata) SetCreatedDate(v time.Time) {
	o.CreatedDate = &IonosTime{v}
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ArtifactMetadata) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactMetadata) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ArtifactMetadata) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *ArtifactMetadata) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedByUserId returns the CreatedByUserId field value if set, zero value otherwise.
func (o *ArtifactMetadata) GetCreatedByUserId() string {
	if o == nil || IsNil(o.CreatedByUserId) {
		var ret string
		return ret
	}
	return *o.CreatedByUserId
}

// GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactMetadata) GetCreatedByUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedByUserId) {
		return nil, false
	}
	return o.CreatedByUserId, true
}

// HasCreatedByUserId returns a boolean if a field has been set.
func (o *ArtifactMetadata) HasCreatedByUserId() bool {
	if o != nil && !IsNil(o.CreatedByUserId) {
		return true
	}

	return false
}

// SetCreatedByUserId gets a reference to the given string and assigns it to the CreatedByUserId field.
func (o *ArtifactMetadata) SetCreatedByUserId(v string) {
	o.CreatedByUserId = &v
}

// GetLastModifiedDate returns the LastModifiedDate field value if set, zero value otherwise.
func (o *ArtifactMetadata) GetLastModifiedDate() time.Time {
	if o == nil || IsNil(o.LastModifiedDate) {
		var ret time.Time
		return ret
	}
	return o.LastModifiedDate.Time
}

// GetLastModifiedDateOk returns a tuple with the LastModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactMetadata) GetLastModifiedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModifiedDate) {
		return nil, false
	}
	return &o.LastModifiedDate.Time, true
}

// HasLastModifiedDate returns a boolean if a field has been set.
func (o *ArtifactMetadata) HasLastModifiedDate() bool {
	if o != nil && !IsNil(o.LastModifiedDate) {
		return true
	}

	return false
}

// SetLastModifiedDate gets a reference to the given time.Time and assigns it to the LastModifiedDate field.
func (o *ArtifactMetadata) SetLastModifiedDate(v time.Time) {
	o.LastModifiedDate = &IonosTime{v}
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise.
func (o *ArtifactMetadata) GetLastModifiedBy() string {
	if o == nil || IsNil(o.LastModifiedBy) {
		var ret string
		return ret
	}
	return *o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactMetadata) GetLastModifiedByOk() (*string, bool) {
	if o == nil || IsNil(o.LastModifiedBy) {
		return nil, false
	}
	return o.LastModifiedBy, true
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *ArtifactMetadata) HasLastModifiedBy() bool {
	if o != nil && !IsNil(o.LastModifiedBy) {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given string and assigns it to the LastModifiedBy field.
func (o *ArtifactMetadata) SetLastModifiedBy(v string) {
	o.LastModifiedBy = &v
}

// GetLastModifiedByUserId returns the LastModifiedByUserId field value if set, zero value otherwise.
func (o *ArtifactMetadata) GetLastModifiedByUserId() string {
	if o == nil || IsNil(o.LastModifiedByUserId) {
		var ret string
		return ret
	}
	return *o.LastModifiedByUserId
}

// GetLastModifiedByUserIdOk returns a tuple with the LastModifiedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactMetadata) GetLastModifiedByUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.LastModifiedByUserId) {
		return nil, false
	}
	return o.LastModifiedByUserId, true
}

// HasLastModifiedByUserId returns a boolean if a field has been set.
func (o *ArtifactMetadata) HasLastModifiedByUserId() bool {
	if o != nil && !IsNil(o.LastModifiedByUserId) {
		return true
	}

	return false
}

// SetLastModifiedByUserId gets a reference to the given string and assigns it to the LastModifiedByUserId field.
func (o *ArtifactMetadata) SetLastModifiedByUserId(v string) {
	o.LastModifiedByUserId = &v
}

// GetResourceURN returns the ResourceURN field value if set, zero value otherwise.
func (o *ArtifactMetadata) GetResourceURN() string {
	if o == nil || IsNil(o.ResourceURN) {
		var ret string
		return ret
	}
	return *o.ResourceURN
}

// GetResourceURNOk returns a tuple with the ResourceURN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactMetadata) GetResourceURNOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceURN) {
		return nil, false
	}
	return o.ResourceURN, true
}

// HasResourceURN returns a boolean if a field has been set.
func (o *ArtifactMetadata) HasResourceURN() bool {
	if o != nil && !IsNil(o.ResourceURN) {
		return true
	}

	return false
}

// SetResourceURN gets a reference to the given string and assigns it to the ResourceURN field.
func (o *ArtifactMetadata) SetResourceURN(v string) {
	o.ResourceURN = &v
}

// GetLastPushedAt returns the LastPushedAt field value
func (o *ArtifactMetadata) GetLastPushedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	if o.LastPushedAt == nil {
		var ret time.Time
		return ret
	}
	return o.LastPushedAt.Time
}

// GetLastPushedAtOk returns a tuple with the LastPushedAt field value
// and a boolean to check if the value has been set.
func (o *ArtifactMetadata) GetLastPushedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	if o.LastPushedAt == nil {
		return nil, true
	}
	return &o.LastPushedAt.Time, true
}

// SetLastPushedAt sets field value
func (o *ArtifactMetadata) SetLastPushedAt(v time.Time) {
	o.LastPushedAt = &IonosTime{v}
}

// GetLastPulledAt returns the LastPulledAt field value if set, zero value otherwise.
func (o *ArtifactMetadata) GetLastPulledAt() time.Time {
	if o == nil || IsNil(o.LastPulledAt) {
		var ret time.Time
		return ret
	}
	return o.LastPulledAt.Time
}

// GetLastPulledAtOk returns a tuple with the LastPulledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactMetadata) GetLastPulledAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastPulledAt) {
		return nil, false
	}
	return &o.LastPulledAt.Time, true
}

// HasLastPulledAt returns a boolean if a field has been set.
func (o *ArtifactMetadata) HasLastPulledAt() bool {
	if o != nil && !IsNil(o.LastPulledAt) {
		return true
	}

	return false
}

// SetLastPulledAt gets a reference to the given time.Time and assigns it to the LastPulledAt field.
func (o *ArtifactMetadata) SetLastPulledAt(v time.Time) {
	o.LastPulledAt = &IonosTime{v}
}

// GetLastScannedAt returns the LastScannedAt field value if set, zero value otherwise.
func (o *ArtifactMetadata) GetLastScannedAt() time.Time {
	if o == nil || IsNil(o.LastScannedAt) {
		var ret time.Time
		return ret
	}
	return o.LastScannedAt.Time
}

// GetLastScannedAtOk returns a tuple with the LastScannedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactMetadata) GetLastScannedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastScannedAt) {
		return nil, false
	}
	return &o.LastScannedAt.Time, true
}

// HasLastScannedAt returns a boolean if a field has been set.
func (o *ArtifactMetadata) HasLastScannedAt() bool {
	if o != nil && !IsNil(o.LastScannedAt) {
		return true
	}

	return false
}

// SetLastScannedAt gets a reference to the given time.Time and assigns it to the LastScannedAt field.
func (o *ArtifactMetadata) SetLastScannedAt(v time.Time) {
	o.LastScannedAt = &IonosTime{v}
}

// GetPushCount returns the PushCount field value
func (o *ArtifactMetadata) GetPushCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PushCount
}

// GetPushCountOk returns a tuple with the PushCount field value
// and a boolean to check if the value has been set.
func (o *ArtifactMetadata) GetPushCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PushCount, true
}

// SetPushCount sets field value
func (o *ArtifactMetadata) SetPushCount(v int64) {
	o.PushCount = v
}

// GetPullCount returns the PullCount field value
func (o *ArtifactMetadata) GetPullCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PullCount
}

// GetPullCountOk returns a tuple with the PullCount field value
// and a boolean to check if the value has been set.
func (o *ArtifactMetadata) GetPullCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PullCount, true
}

// SetPullCount sets field value
func (o *ArtifactMetadata) SetPullCount(v int64) {
	o.PullCount = v
}

// GetVulnMaxSeverity returns the VulnMaxSeverity field value if set, zero value otherwise.
func (o *ArtifactMetadata) GetVulnMaxSeverity() string {
	if o == nil || IsNil(o.VulnMaxSeverity) {
		var ret string
		return ret
	}
	return *o.VulnMaxSeverity
}

// GetVulnMaxSeverityOk returns a tuple with the VulnMaxSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactMetadata) GetVulnMaxSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.VulnMaxSeverity) {
		return nil, false
	}
	return o.VulnMaxSeverity, true
}

// HasVulnMaxSeverity returns a boolean if a field has been set.
func (o *ArtifactMetadata) HasVulnMaxSeverity() bool {
	if o != nil && !IsNil(o.VulnMaxSeverity) {
		return true
	}

	return false
}

// SetVulnMaxSeverity gets a reference to the given string and assigns it to the VulnMaxSeverity field.
func (o *ArtifactMetadata) SetVulnMaxSeverity(v string) {
	o.VulnMaxSeverity = &v
}

// GetVulnTotalScore returns the VulnTotalScore field value if set, zero value otherwise.
func (o *ArtifactMetadata) GetVulnTotalScore() float32 {
	if o == nil || IsNil(o.VulnTotalScore) {
		var ret float32
		return ret
	}
	return *o.VulnTotalScore
}

// GetVulnTotalScoreOk returns a tuple with the VulnTotalScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactMetadata) GetVulnTotalScoreOk() (*float32, bool) {
	if o == nil || IsNil(o.VulnTotalScore) {
		return nil, false
	}
	return o.VulnTotalScore, true
}

// HasVulnTotalScore returns a boolean if a field has been set.
func (o *ArtifactMetadata) HasVulnTotalScore() bool {
	if o != nil && !IsNil(o.VulnTotalScore) {
		return true
	}

	return false
}

// SetVulnTotalScore gets a reference to the given float32 and assigns it to the VulnTotalScore field.
func (o *ArtifactMetadata) SetVulnTotalScore(v float32) {
	o.VulnTotalScore = &v
}

// GetVulnTotalCount returns the VulnTotalCount field value if set, zero value otherwise.
func (o *ArtifactMetadata) GetVulnTotalCount() int64 {
	if o == nil || IsNil(o.VulnTotalCount) {
		var ret int64
		return ret
	}
	return *o.VulnTotalCount
}

// GetVulnTotalCountOk returns a tuple with the VulnTotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactMetadata) GetVulnTotalCountOk() (*int64, bool) {
	if o == nil || IsNil(o.VulnTotalCount) {
		return nil, false
	}
	return o.VulnTotalCount, true
}

// HasVulnTotalCount returns a boolean if a field has been set.
func (o *ArtifactMetadata) HasVulnTotalCount() bool {
	if o != nil && !IsNil(o.VulnTotalCount) {
		return true
	}

	return false
}

// SetVulnTotalCount gets a reference to the given int64 and assigns it to the VulnTotalCount field.
func (o *ArtifactMetadata) SetVulnTotalCount(v int64) {
	o.VulnTotalCount = &v
}

// GetVulnFixableCount returns the VulnFixableCount field value if set, zero value otherwise.
func (o *ArtifactMetadata) GetVulnFixableCount() int64 {
	if o == nil || IsNil(o.VulnFixableCount) {
		var ret int64
		return ret
	}
	return *o.VulnFixableCount
}

// GetVulnFixableCountOk returns a tuple with the VulnFixableCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactMetadata) GetVulnFixableCountOk() (*int64, bool) {
	if o == nil || IsNil(o.VulnFixableCount) {
		return nil, false
	}
	return o.VulnFixableCount, true
}

// HasVulnFixableCount returns a boolean if a field has been set.
func (o *ArtifactMetadata) HasVulnFixableCount() bool {
	if o != nil && !IsNil(o.VulnFixableCount) {
		return true
	}

	return false
}

// SetVulnFixableCount gets a reference to the given int64 and assigns it to the VulnFixableCount field.
func (o *ArtifactMetadata) SetVulnFixableCount(v int64) {
	o.VulnFixableCount = &v
}

func (o ArtifactMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	if !IsNil(o.ResourceURN) {
		toSerialize["resourceURN"] = o.ResourceURN
	}
	toSerialize["lastPushedAt"] = o.LastPushedAt
	if !IsNil(o.LastPulledAt) {
		toSerialize["lastPulledAt"] = o.LastPulledAt
	}
	if !IsNil(o.LastScannedAt) {
		toSerialize["lastScannedAt"] = o.LastScannedAt
	}
	toSerialize["pushCount"] = o.PushCount
	toSerialize["pullCount"] = o.PullCount
	if !IsNil(o.VulnMaxSeverity) {
		toSerialize["vulnMaxSeverity"] = o.VulnMaxSeverity
	}
	if !IsNil(o.VulnTotalScore) {
		toSerialize["vulnTotalScore"] = o.VulnTotalScore
	}
	if !IsNil(o.VulnTotalCount) {
		toSerialize["vulnTotalCount"] = o.VulnTotalCount
	}
	if !IsNil(o.VulnFixableCount) {
		toSerialize["vulnFixableCount"] = o.VulnFixableCount
	}
	return toSerialize, nil
}

type NullableArtifactMetadata struct {
	value *ArtifactMetadata
	isSet bool
}

func (v NullableArtifactMetadata) Get() *ArtifactMetadata {
	return v.value
}

func (v *NullableArtifactMetadata) Set(val *ArtifactMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableArtifactMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableArtifactMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtifactMetadata(val *ArtifactMetadata) *NullableArtifactMetadata {
	return &NullableArtifactMetadata{value: val, isSet: true}
}

func (v NullableArtifactMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtifactMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}