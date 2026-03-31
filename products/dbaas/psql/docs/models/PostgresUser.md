# PostgresUser

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Username** | **string** | The username of the master database user. Must be 16 characters or less and must include only alphanumeric characters (&#x60;[A-Za-z0-9_]&#x60;) and underscores (&#x60;_&#x60;).  | |
|**Password** | **string** | The password for the master database user. Must meet the following requirements: - At least 8 characters long. - Contains at least one lowercase letter. - Contains at least one uppercase letter. - Contains at least one digit (0-9). - Contains at least one special character from the set: @$!%*?&amp;  | |
|**Database** | **string** | The name of the initial database to be created. Must be 63 characters or less and must include only alphanumeric characters (&#x60;[a-z0-9A-Z]&#x60;) and underscores (&#x60;_&#x60;).  | |

## Methods

### NewPostgresUser

`func NewPostgresUser(username string, password string, database string, ) *PostgresUser`

NewPostgresUser instantiates a new PostgresUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresUserWithDefaults

`func NewPostgresUserWithDefaults() *PostgresUser`

NewPostgresUserWithDefaults instantiates a new PostgresUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *PostgresUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PostgresUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PostgresUser) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *PostgresUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PostgresUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PostgresUser) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetDatabase

`func (o *PostgresUser) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *PostgresUser) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *PostgresUser) SetDatabase(v string)`

SetDatabase sets Database field to given value.



