# UserProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Username** | **string** | The username of a given user. | |
|**Password** | Pointer to **string** | The password of a given user. | [optional] |
|**System** | Pointer to **bool** | Describes whether this user is a system user or not. A system user cannot be updated or deleted. | [optional] [readonly] |

## Methods

### NewUserProperties

`func NewUserProperties(username string, ) *UserProperties`

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

### HasPassword

`func (o *UserProperties) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSystem

`func (o *UserProperties) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *UserProperties) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *UserProperties) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *UserProperties) HasSystem() bool`

HasSystem returns a boolean if a field has been set.


