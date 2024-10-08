# RepositoryMetadata

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
|**ArtifactCount** | **int64** |  | |
|**PullCount** | **int64** |  | |
|**PushCount** | **int64** |  | |
|**LastPulledAt** | Pointer to [**time.Time**](time.Time.md) |  | [optional] |
|**LastPushedAt** | Pointer to [**time.Time**](time.Time.md) |  | [optional] |
|**LastSeverity** | Pointer to **string** | The CVSS vulnerability severity rating | [optional] |

## Methods

### NewRepositoryMetadata

`func NewRepositoryMetadata(artifactCount int64, pullCount int64, pushCount int64, ) *RepositoryMetadata`

NewRepositoryMetadata instantiates a new RepositoryMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryMetadataWithDefaults

`func NewRepositoryMetadataWithDefaults() *RepositoryMetadata`

NewRepositoryMetadataWithDefaults instantiates a new RepositoryMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDate

`func (o *RepositoryMetadata) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *RepositoryMetadata) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *RepositoryMetadata) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *RepositoryMetadata) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetCreatedBy

`func (o *RepositoryMetadata) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RepositoryMetadata) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RepositoryMetadata) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *RepositoryMetadata) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *RepositoryMetadata) GetCreatedByUserId() string`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *RepositoryMetadata) GetCreatedByUserIdOk() (*string, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *RepositoryMetadata) SetCreatedByUserId(v string)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *RepositoryMetadata) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetLastModifiedDate

`func (o *RepositoryMetadata) GetLastModifiedDate() time.Time`

GetLastModifiedDate returns the LastModifiedDate field if non-nil, zero value otherwise.

### GetLastModifiedDateOk

`func (o *RepositoryMetadata) GetLastModifiedDateOk() (*time.Time, bool)`

GetLastModifiedDateOk returns a tuple with the LastModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDate

`func (o *RepositoryMetadata) SetLastModifiedDate(v time.Time)`

SetLastModifiedDate sets LastModifiedDate field to given value.

### HasLastModifiedDate

`func (o *RepositoryMetadata) HasLastModifiedDate() bool`

HasLastModifiedDate returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *RepositoryMetadata) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *RepositoryMetadata) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *RepositoryMetadata) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *RepositoryMetadata) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### GetLastModifiedByUserId

`func (o *RepositoryMetadata) GetLastModifiedByUserId() string`

GetLastModifiedByUserId returns the LastModifiedByUserId field if non-nil, zero value otherwise.

### GetLastModifiedByUserIdOk

`func (o *RepositoryMetadata) GetLastModifiedByUserIdOk() (*string, bool)`

GetLastModifiedByUserIdOk returns a tuple with the LastModifiedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedByUserId

`func (o *RepositoryMetadata) SetLastModifiedByUserId(v string)`

SetLastModifiedByUserId sets LastModifiedByUserId field to given value.

### HasLastModifiedByUserId

`func (o *RepositoryMetadata) HasLastModifiedByUserId() bool`

HasLastModifiedByUserId returns a boolean if a field has been set.

### GetResourceURN

`func (o *RepositoryMetadata) GetResourceURN() string`

GetResourceURN returns the ResourceURN field if non-nil, zero value otherwise.

### GetResourceURNOk

`func (o *RepositoryMetadata) GetResourceURNOk() (*string, bool)`

GetResourceURNOk returns a tuple with the ResourceURN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceURN

`func (o *RepositoryMetadata) SetResourceURN(v string)`

SetResourceURN sets ResourceURN field to given value.

### HasResourceURN

`func (o *RepositoryMetadata) HasResourceURN() bool`

HasResourceURN returns a boolean if a field has been set.

### GetArtifactCount

`func (o *RepositoryMetadata) GetArtifactCount() int64`

GetArtifactCount returns the ArtifactCount field if non-nil, zero value otherwise.

### GetArtifactCountOk

`func (o *RepositoryMetadata) GetArtifactCountOk() (*int64, bool)`

GetArtifactCountOk returns a tuple with the ArtifactCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactCount

`func (o *RepositoryMetadata) SetArtifactCount(v int64)`

SetArtifactCount sets ArtifactCount field to given value.


### GetPullCount

`func (o *RepositoryMetadata) GetPullCount() int64`

GetPullCount returns the PullCount field if non-nil, zero value otherwise.

### GetPullCountOk

`func (o *RepositoryMetadata) GetPullCountOk() (*int64, bool)`

GetPullCountOk returns a tuple with the PullCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullCount

`func (o *RepositoryMetadata) SetPullCount(v int64)`

SetPullCount sets PullCount field to given value.


### GetPushCount

`func (o *RepositoryMetadata) GetPushCount() int64`

GetPushCount returns the PushCount field if non-nil, zero value otherwise.

### GetPushCountOk

`func (o *RepositoryMetadata) GetPushCountOk() (*int64, bool)`

GetPushCountOk returns a tuple with the PushCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushCount

`func (o *RepositoryMetadata) SetPushCount(v int64)`

SetPushCount sets PushCount field to given value.


### GetLastPulledAt

`func (o *RepositoryMetadata) GetLastPulledAt() time.Time`

GetLastPulledAt returns the LastPulledAt field if non-nil, zero value otherwise.

### GetLastPulledAtOk

`func (o *RepositoryMetadata) GetLastPulledAtOk() (*time.Time, bool)`

GetLastPulledAtOk returns a tuple with the LastPulledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPulledAt

`func (o *RepositoryMetadata) SetLastPulledAt(v time.Time)`

SetLastPulledAt sets LastPulledAt field to given value.

### HasLastPulledAt

`func (o *RepositoryMetadata) HasLastPulledAt() bool`

HasLastPulledAt returns a boolean if a field has been set.

### GetLastPushedAt

`func (o *RepositoryMetadata) GetLastPushedAt() time.Time`

GetLastPushedAt returns the LastPushedAt field if non-nil, zero value otherwise.

### GetLastPushedAtOk

`func (o *RepositoryMetadata) GetLastPushedAtOk() (*time.Time, bool)`

GetLastPushedAtOk returns a tuple with the LastPushedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPushedAt

`func (o *RepositoryMetadata) SetLastPushedAt(v time.Time)`

SetLastPushedAt sets LastPushedAt field to given value.

### HasLastPushedAt

`func (o *RepositoryMetadata) HasLastPushedAt() bool`

HasLastPushedAt returns a boolean if a field has been set.

### GetLastSeverity

`func (o *RepositoryMetadata) GetLastSeverity() string`

GetLastSeverity returns the LastSeverity field if non-nil, zero value otherwise.

### GetLastSeverityOk

`func (o *RepositoryMetadata) GetLastSeverityOk() (*string, bool)`

GetLastSeverityOk returns a tuple with the LastSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeverity

`func (o *RepositoryMetadata) SetLastSeverity(v string)`

SetLastSeverity sets LastSeverity field to given value.

### HasLastSeverity

`func (o *RepositoryMetadata) HasLastSeverity() bool`

HasLastSeverity returns a boolean if a field has been set.


