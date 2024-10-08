# ArtifactMetadata

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**CreatedDate** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 creation timestamp. | [optional] [readonly] |
|**CreatedBy** | Pointer to **string** | Unique name of the identity that created the resource. | [optional] [readonly] |
|**CreatedByUserId** | Pointer to **string** | Unique id of the identity that created the resource. | [optional] [readonly] |
|**LastModifiedDate** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 modified timestamp. | [optional] [readonly] |
|**LastModifiedBy** | Pointer to **string** | Unique name of the identity that last modified the resource. | [optional] [readonly] |
|**LastModifiedByUserId** | Pointer to **string** | Unique id of the identity that last modified the resource. | [optional] [readonly] |
|**ResourceURN** | Pointer to **string** | Unique name of the resource. | [optional] [readonly] |
|**LastPushedAt** | [**time.Time**](time.Time.md) | The date and time the artifact was last pushed | |
|**LastPulledAt** | Pointer to [**time.Time**](time.Time.md) | The date and time the artifact was last pulled | [optional] |
|**LastScannedAt** | Pointer to [**time.Time**](time.Time.md) | The date and time the artifact was last scanned | [optional] |
|**PushCount** | **int64** | The number of times the artifact was pushed | |
|**PullCount** | **int64** | The number of times the artifact was pulled | |
|**VulnMaxSeverity** | Pointer to **string** | The CVSS vulnerability severity rating | [optional] |
|**VulnTotalScore** | Pointer to **float32** | The total CVSS score of all vulnerabilities of the artifact | [optional] |
|**VulnTotalCount** | Pointer to **int64** | The total number of vulnerabilities of the artifact | [optional] |
|**VulnFixableCount** | Pointer to **int64** | The number of fixable vulnerabilities of the artifact | [optional] |

## Methods

### NewArtifactMetadata

`func NewArtifactMetadata(lastPushedAt time.Time, pushCount int64, pullCount int64, ) *ArtifactMetadata`

NewArtifactMetadata instantiates a new ArtifactMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactMetadataWithDefaults

`func NewArtifactMetadataWithDefaults() *ArtifactMetadata`

NewArtifactMetadataWithDefaults instantiates a new ArtifactMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDate

`func (o *ArtifactMetadata) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ArtifactMetadata) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ArtifactMetadata) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *ArtifactMetadata) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ArtifactMetadata) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ArtifactMetadata) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ArtifactMetadata) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ArtifactMetadata) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *ArtifactMetadata) GetCreatedByUserId() string`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *ArtifactMetadata) GetCreatedByUserIdOk() (*string, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *ArtifactMetadata) SetCreatedByUserId(v string)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *ArtifactMetadata) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetLastModifiedDate

`func (o *ArtifactMetadata) GetLastModifiedDate() time.Time`

GetLastModifiedDate returns the LastModifiedDate field if non-nil, zero value otherwise.

### GetLastModifiedDateOk

`func (o *ArtifactMetadata) GetLastModifiedDateOk() (*time.Time, bool)`

GetLastModifiedDateOk returns a tuple with the LastModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDate

`func (o *ArtifactMetadata) SetLastModifiedDate(v time.Time)`

SetLastModifiedDate sets LastModifiedDate field to given value.

### HasLastModifiedDate

`func (o *ArtifactMetadata) HasLastModifiedDate() bool`

HasLastModifiedDate returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *ArtifactMetadata) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *ArtifactMetadata) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *ArtifactMetadata) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *ArtifactMetadata) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### GetLastModifiedByUserId

`func (o *ArtifactMetadata) GetLastModifiedByUserId() string`

GetLastModifiedByUserId returns the LastModifiedByUserId field if non-nil, zero value otherwise.

### GetLastModifiedByUserIdOk

`func (o *ArtifactMetadata) GetLastModifiedByUserIdOk() (*string, bool)`

GetLastModifiedByUserIdOk returns a tuple with the LastModifiedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedByUserId

`func (o *ArtifactMetadata) SetLastModifiedByUserId(v string)`

SetLastModifiedByUserId sets LastModifiedByUserId field to given value.

### HasLastModifiedByUserId

`func (o *ArtifactMetadata) HasLastModifiedByUserId() bool`

HasLastModifiedByUserId returns a boolean if a field has been set.

### GetResourceURN

`func (o *ArtifactMetadata) GetResourceURN() string`

GetResourceURN returns the ResourceURN field if non-nil, zero value otherwise.

### GetResourceURNOk

`func (o *ArtifactMetadata) GetResourceURNOk() (*string, bool)`

GetResourceURNOk returns a tuple with the ResourceURN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceURN

`func (o *ArtifactMetadata) SetResourceURN(v string)`

SetResourceURN sets ResourceURN field to given value.

### HasResourceURN

`func (o *ArtifactMetadata) HasResourceURN() bool`

HasResourceURN returns a boolean if a field has been set.

### GetLastPushedAt

`func (o *ArtifactMetadata) GetLastPushedAt() time.Time`

GetLastPushedAt returns the LastPushedAt field if non-nil, zero value otherwise.

### GetLastPushedAtOk

`func (o *ArtifactMetadata) GetLastPushedAtOk() (*time.Time, bool)`

GetLastPushedAtOk returns a tuple with the LastPushedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPushedAt

`func (o *ArtifactMetadata) SetLastPushedAt(v time.Time)`

SetLastPushedAt sets LastPushedAt field to given value.


### GetLastPulledAt

`func (o *ArtifactMetadata) GetLastPulledAt() time.Time`

GetLastPulledAt returns the LastPulledAt field if non-nil, zero value otherwise.

### GetLastPulledAtOk

`func (o *ArtifactMetadata) GetLastPulledAtOk() (*time.Time, bool)`

GetLastPulledAtOk returns a tuple with the LastPulledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPulledAt

`func (o *ArtifactMetadata) SetLastPulledAt(v time.Time)`

SetLastPulledAt sets LastPulledAt field to given value.

### HasLastPulledAt

`func (o *ArtifactMetadata) HasLastPulledAt() bool`

HasLastPulledAt returns a boolean if a field has been set.

### GetLastScannedAt

`func (o *ArtifactMetadata) GetLastScannedAt() time.Time`

GetLastScannedAt returns the LastScannedAt field if non-nil, zero value otherwise.

### GetLastScannedAtOk

`func (o *ArtifactMetadata) GetLastScannedAtOk() (*time.Time, bool)`

GetLastScannedAtOk returns a tuple with the LastScannedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastScannedAt

`func (o *ArtifactMetadata) SetLastScannedAt(v time.Time)`

SetLastScannedAt sets LastScannedAt field to given value.

### HasLastScannedAt

`func (o *ArtifactMetadata) HasLastScannedAt() bool`

HasLastScannedAt returns a boolean if a field has been set.

### GetPushCount

`func (o *ArtifactMetadata) GetPushCount() int64`

GetPushCount returns the PushCount field if non-nil, zero value otherwise.

### GetPushCountOk

`func (o *ArtifactMetadata) GetPushCountOk() (*int64, bool)`

GetPushCountOk returns a tuple with the PushCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushCount

`func (o *ArtifactMetadata) SetPushCount(v int64)`

SetPushCount sets PushCount field to given value.


### GetPullCount

`func (o *ArtifactMetadata) GetPullCount() int64`

GetPullCount returns the PullCount field if non-nil, zero value otherwise.

### GetPullCountOk

`func (o *ArtifactMetadata) GetPullCountOk() (*int64, bool)`

GetPullCountOk returns a tuple with the PullCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullCount

`func (o *ArtifactMetadata) SetPullCount(v int64)`

SetPullCount sets PullCount field to given value.


### GetVulnMaxSeverity

`func (o *ArtifactMetadata) GetVulnMaxSeverity() string`

GetVulnMaxSeverity returns the VulnMaxSeverity field if non-nil, zero value otherwise.

### GetVulnMaxSeverityOk

`func (o *ArtifactMetadata) GetVulnMaxSeverityOk() (*string, bool)`

GetVulnMaxSeverityOk returns a tuple with the VulnMaxSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnMaxSeverity

`func (o *ArtifactMetadata) SetVulnMaxSeverity(v string)`

SetVulnMaxSeverity sets VulnMaxSeverity field to given value.

### HasVulnMaxSeverity

`func (o *ArtifactMetadata) HasVulnMaxSeverity() bool`

HasVulnMaxSeverity returns a boolean if a field has been set.

### GetVulnTotalScore

`func (o *ArtifactMetadata) GetVulnTotalScore() float32`

GetVulnTotalScore returns the VulnTotalScore field if non-nil, zero value otherwise.

### GetVulnTotalScoreOk

`func (o *ArtifactMetadata) GetVulnTotalScoreOk() (*float32, bool)`

GetVulnTotalScoreOk returns a tuple with the VulnTotalScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnTotalScore

`func (o *ArtifactMetadata) SetVulnTotalScore(v float32)`

SetVulnTotalScore sets VulnTotalScore field to given value.

### HasVulnTotalScore

`func (o *ArtifactMetadata) HasVulnTotalScore() bool`

HasVulnTotalScore returns a boolean if a field has been set.

### GetVulnTotalCount

`func (o *ArtifactMetadata) GetVulnTotalCount() int64`

GetVulnTotalCount returns the VulnTotalCount field if non-nil, zero value otherwise.

### GetVulnTotalCountOk

`func (o *ArtifactMetadata) GetVulnTotalCountOk() (*int64, bool)`

GetVulnTotalCountOk returns a tuple with the VulnTotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnTotalCount

`func (o *ArtifactMetadata) SetVulnTotalCount(v int64)`

SetVulnTotalCount sets VulnTotalCount field to given value.

### HasVulnTotalCount

`func (o *ArtifactMetadata) HasVulnTotalCount() bool`

HasVulnTotalCount returns a boolean if a field has been set.

### GetVulnFixableCount

`func (o *ArtifactMetadata) GetVulnFixableCount() int64`

GetVulnFixableCount returns the VulnFixableCount field if non-nil, zero value otherwise.

### GetVulnFixableCountOk

`func (o *ArtifactMetadata) GetVulnFixableCountOk() (*int64, bool)`

GetVulnFixableCountOk returns a tuple with the VulnFixableCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnFixableCount

`func (o *ArtifactMetadata) SetVulnFixableCount(v int64)`

SetVulnFixableCount sets VulnFixableCount field to given value.

### HasVulnFixableCount

`func (o *ArtifactMetadata) HasVulnFixableCount() bool`

HasVulnFixableCount returns a boolean if a field has been set.


