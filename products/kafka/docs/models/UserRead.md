# UserRead

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The ID (UUID) of the User. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the User. | |
|**Metadata** | [**ResourceMetadata**](ResourceMetadata.md) |  | |
|**Properties** | [**User**](User.md) |  | |

## Methods

### NewUserRead

`func NewUserRead(id string, type_ string, href string, metadata ResourceMetadata, properties User, ) *UserRead`

NewUserRead instantiates a new UserRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserReadWithDefaults

`func NewUserReadWithDefaults() *UserRead`

NewUserReadWithDefaults instantiates a new UserRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserRead) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *UserRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserRead) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *UserRead) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *UserRead) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *UserRead) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *UserRead) GetMetadata() ResourceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UserRead) GetMetadataOk() (*ResourceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UserRead) SetMetadata(v ResourceMetadata)`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *UserRead) GetProperties() User`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *UserRead) GetPropertiesOk() (*User, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *UserRead) SetProperties(v User)`

SetProperties sets Properties field to given value.



