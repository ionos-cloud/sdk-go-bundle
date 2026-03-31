# PostgresVersion

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Version** | Pointer to **string** | The PostgreSQL version for the cluster. | [optional] |
|**Status** | Pointer to **string** | The support status of the version. | [optional] |
|**Comment** | Pointer to **string** | Additional information about the version status. | [optional] |
|**CanUpgradeTo** | Pointer to **[]string** | List of versions that this version can be upgraded to. | [optional] |

## Methods

### NewPostgresVersion

`func NewPostgresVersion() *PostgresVersion`

NewPostgresVersion instantiates a new PostgresVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresVersionWithDefaults

`func NewPostgresVersionWithDefaults() *PostgresVersion`

NewPostgresVersionWithDefaults instantiates a new PostgresVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *PostgresVersion) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PostgresVersion) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PostgresVersion) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PostgresVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetStatus

`func (o *PostgresVersion) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PostgresVersion) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PostgresVersion) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PostgresVersion) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetComment

`func (o *PostgresVersion) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PostgresVersion) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PostgresVersion) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PostgresVersion) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCanUpgradeTo

`func (o *PostgresVersion) GetCanUpgradeTo() []string`

GetCanUpgradeTo returns the CanUpgradeTo field if non-nil, zero value otherwise.

### GetCanUpgradeToOk

`func (o *PostgresVersion) GetCanUpgradeToOk() (*[]string, bool)`

GetCanUpgradeToOk returns a tuple with the CanUpgradeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUpgradeTo

`func (o *PostgresVersion) SetCanUpgradeTo(v []string)`

SetCanUpgradeTo sets CanUpgradeTo field to given value.

### HasCanUpgradeTo

`func (o *PostgresVersion) HasCanUpgradeTo() bool`

HasCanUpgradeTo returns a boolean if a field has been set.


