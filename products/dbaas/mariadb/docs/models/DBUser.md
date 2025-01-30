# DBUser

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Username** | **string** | The username for the initial MariaDB user. Some system usernames are restricted (e.g. \&quot;mariadb\&quot;, \&quot;admin\&quot;, \&quot;standby\&quot;).  The username should be compliant with the following rules: - Must not exceed 16 characters - Must start with a letter - Must contain only letters, numbers, or underscores  | |
|**Password** | **string** | The password for a MariaDB user. | |

## Methods

### NewDBUser

`func NewDBUser(username string, password string, ) *DBUser`

NewDBUser instantiates a new DBUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDBUserWithDefaults

`func NewDBUserWithDefaults() *DBUser`

NewDBUserWithDefaults instantiates a new DBUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *DBUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DBUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DBUser) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *DBUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DBUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DBUser) SetPassword(v string)`

SetPassword sets Password field to given value.



