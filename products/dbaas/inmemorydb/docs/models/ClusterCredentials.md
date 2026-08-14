# ClusterCredentials

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Username** | **string** | The username for the In-Memory DB user. Must be 2–16 characters and may only contain alphanumeric characters ([A-Za-z0-9]) and underscores (_). Restricted usernames (for example, admin, standby) are not allowed.  | |
|**Password** | [**HashedPassword**](HashedPassword.md) |  | |

## Methods

### NewClusterCredentials

`func NewClusterCredentials(username string, password HashedPassword, ) *ClusterCredentials`

NewClusterCredentials instantiates a new ClusterCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCredentialsWithDefaults

`func NewClusterCredentialsWithDefaults() *ClusterCredentials`

NewClusterCredentialsWithDefaults instantiates a new ClusterCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *ClusterCredentials) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ClusterCredentials) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ClusterCredentials) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *ClusterCredentials) GetPassword() HashedPassword`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ClusterCredentials) GetPasswordOk() (*HashedPassword, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ClusterCredentials) SetPassword(v HashedPassword)`

SetPassword sets Password field to given value.



