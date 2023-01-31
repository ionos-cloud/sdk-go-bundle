# UserProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Username** | **string** |  | |
|**Password** | **string** |  | |
|**Roles** | Pointer to [**[]UserRoles**](UserRoles.md) |  | [optional] |

## Methods

### NewUserProperties

`func NewUserProperties(username string, password string, ) *UserProperties`

NewUserProperties instantiates a new UserProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPropertiesWithDefaults

`func NewUserPropertiesWithDefaults() *UserProperties`

NewUserPropertiesWithDefaults instantiates a new UserProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UserProperties) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserProperties) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserProperties) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *UserProperties) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserProperties) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserProperties) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetRoles

`func (o *UserProperties) GetRoles() []UserRoles`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserProperties) GetRolesOk() (*[]UserRoles, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserProperties) SetRoles(v []UserRoles)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UserProperties) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


