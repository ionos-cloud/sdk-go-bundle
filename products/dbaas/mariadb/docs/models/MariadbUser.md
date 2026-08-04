# MariadbUser

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Username** | **string** | The username of the initial MariaDB user. Must be 16 characters or less and must include only alphanumeric characters (&#x60;[A-Za-z0-9_]&#x60;) and underscores (&#x60;_&#x60;). Some usernames are reserved for platform use (for example &#x60;mariadb&#x60;, &#x60;admin&#x60;, &#x60;standby&#x60;).  | |
|**Password** | **string** | The password for the initial MariaDB user. Must be between 10 and 256 characters long. For a strong password we recommend that it also meets the following criteria, though these are not enforced: - Contains at least one lowercase letter. - Contains at least one uppercase letter. - Contains at least one digit (0-9). - Contains at least one special character from the set: @$!%*?&amp;  | |
|**Database** | **string** | The name of the initial database to be created. Must be 63 characters or less and must include only alphanumeric characters (&#x60;[a-z0-9A-Z]&#x60;) and underscores (&#x60;_&#x60;).  | |

## Methods

### NewMariadbUser

`func NewMariadbUser(username string, password string, database string, ) *MariadbUser`

NewMariadbUser instantiates a new MariadbUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMariadbUserWithDefaults

`func NewMariadbUserWithDefaults() *MariadbUser`

NewMariadbUserWithDefaults instantiates a new MariadbUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *MariadbUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *MariadbUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *MariadbUser) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *MariadbUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MariadbUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MariadbUser) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetDatabase

`func (o *MariadbUser) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *MariadbUser) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *MariadbUser) SetDatabase(v string)`

SetDatabase sets Database field to given value.



