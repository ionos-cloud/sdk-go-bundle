# UserReadAccess

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The ID (UUID) of the User. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the User. | |
|**Metadata** | [**UserAccessMetadata**](UserAccessMetadata.md) |  | |
|**Properties** | [**User**](User.md) |  | |

## Methods

### NewUserReadAccess

`func NewUserReadAccess(id string, type_ string, href string, metadata UserAccessMetadata, properties User, ) *UserReadAccess`

NewUserReadAccess instantiates a new UserReadAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserReadAccessWithDefaults

`func NewUserReadAccessWithDefaults() *UserReadAccess`

NewUserReadAccessWithDefaults instantiates a new UserReadAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserReadAccess) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserReadAccess) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserReadAccess) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *UserReadAccess) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserReadAccess) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserReadAccess) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *UserReadAccess) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *UserReadAccess) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *UserReadAccess) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *UserReadAccess) GetMetadata() UserAccessMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UserReadAccess) GetMetadataOk() (*UserAccessMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UserReadAccess) SetMetadata(v UserAccessMetadata)`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *UserReadAccess) GetProperties() User`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *UserReadAccess) GetPropertiesOk() (*User, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *UserReadAccess) SetProperties(v User)`

SetProperties sets Properties field to given value.



