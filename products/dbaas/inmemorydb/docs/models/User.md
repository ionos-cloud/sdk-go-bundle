# User

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Username** | **string** | The username for the initial In-Memory DB user. Some system usernames are restricted (e.g. &#x60;\&quot;admin\&quot;&#x60;, &#x60;\&quot;standby\&quot;&#x60;).  | |
|**Password** | Pointer to [**UserPassword**](UserPassword.md) |  | [optional] |

## Methods

### NewUser

`func NewUser(username string, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *User) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *User) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *User) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *User) GetPassword() UserPassword`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *User) GetPasswordOk() (*UserPassword, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *User) SetPassword(v UserPassword)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *User) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


