# User

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Type** | Pointer to [**ResourceType**](ResourceType.md) |  | [optional] |
|**Metadata** | Pointer to [**UserMetadata**](UserMetadata.md) |  | [optional] |
|**Properties** | Pointer to [**UserProperties**](UserProperties.md) |  | [optional] |

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *User) GetType() ResourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *User) GetTypeOk() (*ResourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *User) SetType(v ResourceType)`

SetType sets Type field to given value.

### HasType

`func (o *User) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMetadata

`func (o *User) GetMetadata() UserMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *User) GetMetadataOk() (*UserMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *User) SetMetadata(v UserMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *User) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *User) GetProperties() UserProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *User) GetPropertiesOk() (*UserProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *User) SetProperties(v UserProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *User) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


