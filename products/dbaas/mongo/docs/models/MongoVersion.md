# MongoVersion

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Version** | Pointer to **string** | The MongoDB version for the cluster. Use GET /versions to retrieve the list of supported versions and their available upgrade paths.  | [optional] |
|**Status** | Pointer to **string** | The support status of the version. | [optional] |
|**Comment** | Pointer to **string** | Additional information about the version status. | [optional] |
|**CanUpgradeTo** | Pointer to **[]string** | List of versions this version can be upgraded to. Only versions listed here are valid upgrade targets when updating the version field via PUT. Downgrades are not supported. | [optional] |

## Methods

### NewMongoVersion

`func NewMongoVersion() *MongoVersion`

NewMongoVersion instantiates a new MongoVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMongoVersionWithDefaults

`func NewMongoVersionWithDefaults() *MongoVersion`

NewMongoVersionWithDefaults instantiates a new MongoVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *MongoVersion) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MongoVersion) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MongoVersion) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MongoVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetStatus

`func (o *MongoVersion) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MongoVersion) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MongoVersion) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MongoVersion) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetComment

`func (o *MongoVersion) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *MongoVersion) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *MongoVersion) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *MongoVersion) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCanUpgradeTo

`func (o *MongoVersion) GetCanUpgradeTo() []string`

GetCanUpgradeTo returns the CanUpgradeTo field if non-nil, zero value otherwise.

### GetCanUpgradeToOk

`func (o *MongoVersion) GetCanUpgradeToOk() (*[]string, bool)`

GetCanUpgradeToOk returns a tuple with the CanUpgradeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUpgradeTo

`func (o *MongoVersion) SetCanUpgradeTo(v []string)`

SetCanUpgradeTo sets CanUpgradeTo field to given value.

### HasCanUpgradeTo

`func (o *MongoVersion) HasCanUpgradeTo() bool`

HasCanUpgradeTo returns a boolean if a field has been set.


