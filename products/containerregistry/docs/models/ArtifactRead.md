# ArtifactRead

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The digest of the artifact | |
|**Type** | **string** |  | |
|**Href** | **string** |  | |
|**Metadata** | [**ArtifactMetadata**](ArtifactMetadata.md) |  | |
|**Properties** | [**Artifact**](Artifact.md) |  | |

## Methods

### NewArtifactRead

`func NewArtifactRead(id string, type_ string, href string, metadata ArtifactMetadata, properties Artifact, ) *ArtifactRead`

NewArtifactRead instantiates a new ArtifactRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactReadWithDefaults

`func NewArtifactReadWithDefaults() *ArtifactRead`

NewArtifactReadWithDefaults instantiates a new ArtifactRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ArtifactRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArtifactRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArtifactRead) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ArtifactRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArtifactRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArtifactRead) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *ArtifactRead) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ArtifactRead) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ArtifactRead) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *ArtifactRead) GetMetadata() ArtifactMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ArtifactRead) GetMetadataOk() (*ArtifactMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ArtifactRead) SetMetadata(v ArtifactMetadata)`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *ArtifactRead) GetProperties() Artifact`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ArtifactRead) GetPropertiesOk() (*Artifact, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ArtifactRead) SetProperties(v Artifact)`

SetProperties sets Properties field to given value.



