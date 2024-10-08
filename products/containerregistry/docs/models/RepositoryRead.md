# RepositoryRead

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** |  | |
|**Type** | **string** |  | |
|**Href** | **string** |  | |
|**Metadata** | [**RepositoryMetadata**](RepositoryMetadata.md) |  | |
|**Properties** | [**Repository**](Repository.md) |  | |

## Methods

### NewRepositoryRead

`func NewRepositoryRead(id string, type_ string, href string, metadata RepositoryMetadata, properties Repository, ) *RepositoryRead`

NewRepositoryRead instantiates a new RepositoryRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryReadWithDefaults

`func NewRepositoryReadWithDefaults() *RepositoryRead`

NewRepositoryReadWithDefaults instantiates a new RepositoryRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RepositoryRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RepositoryRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RepositoryRead) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *RepositoryRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RepositoryRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RepositoryRead) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *RepositoryRead) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RepositoryRead) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RepositoryRead) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *RepositoryRead) GetMetadata() RepositoryMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RepositoryRead) GetMetadataOk() (*RepositoryMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RepositoryRead) SetMetadata(v RepositoryMetadata)`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *RepositoryRead) GetProperties() Repository`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *RepositoryRead) GetPropertiesOk() (*Repository, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *RepositoryRead) SetProperties(v Repository)`

SetProperties sets Properties field to given value.



