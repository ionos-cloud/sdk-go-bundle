# MongoUser

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Username** | **string** | The username of the initial MongoDB user. Must be between 2 and 16 characters and must include only alphanumeric characters (&#x60;[A-Za-z0-9]&#x60;) and underscores (&#x60;_&#x60;). Some usernames are reserved for platform use.  | |
|**Password** | Pointer to **string** | The password for the initial MongoDB user. Password is required whenever credentials are supplied, on create and on update; never returned in responses. Must be between 8 and 256 characters long. For a strong password we recommend that it also meets the following criteria, though these are not enforced: - Contains at least one lowercase letter. - Contains at least one uppercase letter. - Contains at least one digit (0-9). - Contains at least one special character from the set: @$!%*?&amp;  | [optional] |
|**Database** | **string** | The name of the initial database to be created. Must be 63 characters or less and must include only alphanumeric characters (&#x60;[A-Za-z0-9]&#x60;) and underscores (&#x60;_&#x60;).  | |

## Methods

### NewMongoUser

`func NewMongoUser(username string, database string, ) *MongoUser`

NewMongoUser instantiates a new MongoUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMongoUserWithDefaults

`func NewMongoUserWithDefaults() *MongoUser`

NewMongoUserWithDefaults instantiates a new MongoUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *MongoUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *MongoUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *MongoUser) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *MongoUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MongoUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MongoUser) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *MongoUser) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetDatabase

`func (o *MongoUser) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *MongoUser) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *MongoUser) SetDatabase(v string)`

SetDatabase sets Database field to given value.



