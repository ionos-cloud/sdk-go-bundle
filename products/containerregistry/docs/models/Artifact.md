# Artifact

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**RepositoryName** | **string** |  | |
|**Digest** | **string** | The digest of the artifact | |
|**Tags** | Pointer to **[]string** | The tags of an artifact | [optional] |
|**MediaType** | **string** | The media type of the artifact | |

## Methods

### NewArtifact

`func NewArtifact(repositoryName string, digest string, mediaType string, ) *Artifact`

NewArtifact instantiates a new Artifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactWithDefaults

`func NewArtifactWithDefaults() *Artifact`

NewArtifactWithDefaults instantiates a new Artifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepositoryName

`func (o *Artifact) GetRepositoryName() string`

GetRepositoryName returns the RepositoryName field if non-nil, zero value otherwise.

### GetRepositoryNameOk

`func (o *Artifact) GetRepositoryNameOk() (*string, bool)`

GetRepositoryNameOk returns a tuple with the RepositoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryName

`func (o *Artifact) SetRepositoryName(v string)`

SetRepositoryName sets RepositoryName field to given value.


### GetDigest

`func (o *Artifact) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *Artifact) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *Artifact) SetDigest(v string)`

SetDigest sets Digest field to given value.


### GetTags

`func (o *Artifact) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Artifact) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Artifact) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Artifact) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMediaType

`func (o *Artifact) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *Artifact) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *Artifact) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.



