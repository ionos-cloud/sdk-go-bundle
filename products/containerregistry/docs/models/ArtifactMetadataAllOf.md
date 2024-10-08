# ArtifactMetadataAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
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

### NewArtifactMetadataAllOf

`func NewArtifactMetadataAllOf(lastPushedAt time.Time, pushCount int64, pullCount int64, ) *ArtifactMetadataAllOf`

NewArtifactMetadataAllOf instantiates a new ArtifactMetadataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactMetadataAllOfWithDefaults

`func NewArtifactMetadataAllOfWithDefaults() *ArtifactMetadataAllOf`

NewArtifactMetadataAllOfWithDefaults instantiates a new ArtifactMetadataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastPushedAt

`func (o *ArtifactMetadataAllOf) GetLastPushedAt() time.Time`

GetLastPushedAt returns the LastPushedAt field if non-nil, zero value otherwise.

### GetLastPushedAtOk

`func (o *ArtifactMetadataAllOf) GetLastPushedAtOk() (*time.Time, bool)`

GetLastPushedAtOk returns a tuple with the LastPushedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPushedAt

`func (o *ArtifactMetadataAllOf) SetLastPushedAt(v time.Time)`

SetLastPushedAt sets LastPushedAt field to given value.


### GetLastPulledAt

`func (o *ArtifactMetadataAllOf) GetLastPulledAt() time.Time`

GetLastPulledAt returns the LastPulledAt field if non-nil, zero value otherwise.

### GetLastPulledAtOk

`func (o *ArtifactMetadataAllOf) GetLastPulledAtOk() (*time.Time, bool)`

GetLastPulledAtOk returns a tuple with the LastPulledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPulledAt

`func (o *ArtifactMetadataAllOf) SetLastPulledAt(v time.Time)`

SetLastPulledAt sets LastPulledAt field to given value.

### HasLastPulledAt

`func (o *ArtifactMetadataAllOf) HasLastPulledAt() bool`

HasLastPulledAt returns a boolean if a field has been set.

### GetLastScannedAt

`func (o *ArtifactMetadataAllOf) GetLastScannedAt() time.Time`

GetLastScannedAt returns the LastScannedAt field if non-nil, zero value otherwise.

### GetLastScannedAtOk

`func (o *ArtifactMetadataAllOf) GetLastScannedAtOk() (*time.Time, bool)`

GetLastScannedAtOk returns a tuple with the LastScannedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastScannedAt

`func (o *ArtifactMetadataAllOf) SetLastScannedAt(v time.Time)`

SetLastScannedAt sets LastScannedAt field to given value.

### HasLastScannedAt

`func (o *ArtifactMetadataAllOf) HasLastScannedAt() bool`

HasLastScannedAt returns a boolean if a field has been set.

### GetPushCount

`func (o *ArtifactMetadataAllOf) GetPushCount() int64`

GetPushCount returns the PushCount field if non-nil, zero value otherwise.

### GetPushCountOk

`func (o *ArtifactMetadataAllOf) GetPushCountOk() (*int64, bool)`

GetPushCountOk returns a tuple with the PushCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushCount

`func (o *ArtifactMetadataAllOf) SetPushCount(v int64)`

SetPushCount sets PushCount field to given value.


### GetPullCount

`func (o *ArtifactMetadataAllOf) GetPullCount() int64`

GetPullCount returns the PullCount field if non-nil, zero value otherwise.

### GetPullCountOk

`func (o *ArtifactMetadataAllOf) GetPullCountOk() (*int64, bool)`

GetPullCountOk returns a tuple with the PullCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullCount

`func (o *ArtifactMetadataAllOf) SetPullCount(v int64)`

SetPullCount sets PullCount field to given value.


### GetVulnMaxSeverity

`func (o *ArtifactMetadataAllOf) GetVulnMaxSeverity() string`

GetVulnMaxSeverity returns the VulnMaxSeverity field if non-nil, zero value otherwise.

### GetVulnMaxSeverityOk

`func (o *ArtifactMetadataAllOf) GetVulnMaxSeverityOk() (*string, bool)`

GetVulnMaxSeverityOk returns a tuple with the VulnMaxSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnMaxSeverity

`func (o *ArtifactMetadataAllOf) SetVulnMaxSeverity(v string)`

SetVulnMaxSeverity sets VulnMaxSeverity field to given value.

### HasVulnMaxSeverity

`func (o *ArtifactMetadataAllOf) HasVulnMaxSeverity() bool`

HasVulnMaxSeverity returns a boolean if a field has been set.

### GetVulnTotalScore

`func (o *ArtifactMetadataAllOf) GetVulnTotalScore() float32`

GetVulnTotalScore returns the VulnTotalScore field if non-nil, zero value otherwise.

### GetVulnTotalScoreOk

`func (o *ArtifactMetadataAllOf) GetVulnTotalScoreOk() (*float32, bool)`

GetVulnTotalScoreOk returns a tuple with the VulnTotalScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnTotalScore

`func (o *ArtifactMetadataAllOf) SetVulnTotalScore(v float32)`

SetVulnTotalScore sets VulnTotalScore field to given value.

### HasVulnTotalScore

`func (o *ArtifactMetadataAllOf) HasVulnTotalScore() bool`

HasVulnTotalScore returns a boolean if a field has been set.

### GetVulnTotalCount

`func (o *ArtifactMetadataAllOf) GetVulnTotalCount() int64`

GetVulnTotalCount returns the VulnTotalCount field if non-nil, zero value otherwise.

### GetVulnTotalCountOk

`func (o *ArtifactMetadataAllOf) GetVulnTotalCountOk() (*int64, bool)`

GetVulnTotalCountOk returns a tuple with the VulnTotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnTotalCount

`func (o *ArtifactMetadataAllOf) SetVulnTotalCount(v int64)`

SetVulnTotalCount sets VulnTotalCount field to given value.

### HasVulnTotalCount

`func (o *ArtifactMetadataAllOf) HasVulnTotalCount() bool`

HasVulnTotalCount returns a boolean if a field has been set.

### GetVulnFixableCount

`func (o *ArtifactMetadataAllOf) GetVulnFixableCount() int64`

GetVulnFixableCount returns the VulnFixableCount field if non-nil, zero value otherwise.

### GetVulnFixableCountOk

`func (o *ArtifactMetadataAllOf) GetVulnFixableCountOk() (*int64, bool)`

GetVulnFixableCountOk returns a tuple with the VulnFixableCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnFixableCount

`func (o *ArtifactMetadataAllOf) SetVulnFixableCount(v int64)`

SetVulnFixableCount sets VulnFixableCount field to given value.

### HasVulnFixableCount

`func (o *ArtifactMetadataAllOf) HasVulnFixableCount() bool`

HasVulnFixableCount returns a boolean if a field has been set.


