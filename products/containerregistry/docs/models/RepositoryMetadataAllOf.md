# RepositoryMetadataAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**ArtifactCount** | **int64** |  | |
|**PullCount** | **int64** |  | |
|**PushCount** | **int64** |  | |
|**LastPulledAt** | Pointer to [**time.Time**](time.Time.md) |  | [optional] |
|**LastPushedAt** | Pointer to [**time.Time**](time.Time.md) |  | [optional] |
|**LastSeverity** | Pointer to **string** | The CVSS vulnerability severity rating | [optional] |

## Methods

### NewRepositoryMetadataAllOf

`func NewRepositoryMetadataAllOf(artifactCount int64, pullCount int64, pushCount int64, ) *RepositoryMetadataAllOf`

NewRepositoryMetadataAllOf instantiates a new RepositoryMetadataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryMetadataAllOfWithDefaults

`func NewRepositoryMetadataAllOfWithDefaults() *RepositoryMetadataAllOf`

NewRepositoryMetadataAllOfWithDefaults instantiates a new RepositoryMetadataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactCount

`func (o *RepositoryMetadataAllOf) GetArtifactCount() int64`

GetArtifactCount returns the ArtifactCount field if non-nil, zero value otherwise.

### GetArtifactCountOk

`func (o *RepositoryMetadataAllOf) GetArtifactCountOk() (*int64, bool)`

GetArtifactCountOk returns a tuple with the ArtifactCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactCount

`func (o *RepositoryMetadataAllOf) SetArtifactCount(v int64)`

SetArtifactCount sets ArtifactCount field to given value.


### GetPullCount

`func (o *RepositoryMetadataAllOf) GetPullCount() int64`

GetPullCount returns the PullCount field if non-nil, zero value otherwise.

### GetPullCountOk

`func (o *RepositoryMetadataAllOf) GetPullCountOk() (*int64, bool)`

GetPullCountOk returns a tuple with the PullCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullCount

`func (o *RepositoryMetadataAllOf) SetPullCount(v int64)`

SetPullCount sets PullCount field to given value.


### GetPushCount

`func (o *RepositoryMetadataAllOf) GetPushCount() int64`

GetPushCount returns the PushCount field if non-nil, zero value otherwise.

### GetPushCountOk

`func (o *RepositoryMetadataAllOf) GetPushCountOk() (*int64, bool)`

GetPushCountOk returns a tuple with the PushCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushCount

`func (o *RepositoryMetadataAllOf) SetPushCount(v int64)`

SetPushCount sets PushCount field to given value.


### GetLastPulledAt

`func (o *RepositoryMetadataAllOf) GetLastPulledAt() time.Time`

GetLastPulledAt returns the LastPulledAt field if non-nil, zero value otherwise.

### GetLastPulledAtOk

`func (o *RepositoryMetadataAllOf) GetLastPulledAtOk() (*time.Time, bool)`

GetLastPulledAtOk returns a tuple with the LastPulledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPulledAt

`func (o *RepositoryMetadataAllOf) SetLastPulledAt(v time.Time)`

SetLastPulledAt sets LastPulledAt field to given value.

### HasLastPulledAt

`func (o *RepositoryMetadataAllOf) HasLastPulledAt() bool`

HasLastPulledAt returns a boolean if a field has been set.

### GetLastPushedAt

`func (o *RepositoryMetadataAllOf) GetLastPushedAt() time.Time`

GetLastPushedAt returns the LastPushedAt field if non-nil, zero value otherwise.

### GetLastPushedAtOk

`func (o *RepositoryMetadataAllOf) GetLastPushedAtOk() (*time.Time, bool)`

GetLastPushedAtOk returns a tuple with the LastPushedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPushedAt

`func (o *RepositoryMetadataAllOf) SetLastPushedAt(v time.Time)`

SetLastPushedAt sets LastPushedAt field to given value.

### HasLastPushedAt

`func (o *RepositoryMetadataAllOf) HasLastPushedAt() bool`

HasLastPushedAt returns a boolean if a field has been set.

### GetLastSeverity

`func (o *RepositoryMetadataAllOf) GetLastSeverity() string`

GetLastSeverity returns the LastSeverity field if non-nil, zero value otherwise.

### GetLastSeverityOk

`func (o *RepositoryMetadataAllOf) GetLastSeverityOk() (*string, bool)`

GetLastSeverityOk returns a tuple with the LastSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeverity

`func (o *RepositoryMetadataAllOf) SetLastSeverity(v string)`

SetLastSeverity sets LastSeverity field to given value.

### HasLastSeverity

`func (o *RepositoryMetadataAllOf) HasLastSeverity() bool`

HasLastSeverity returns a boolean if a field has been set.


