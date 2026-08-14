# MariadbVersion

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Version** | Pointer to **string** | The MariaDB version for the cluster. Use GET /versions to retrieve the list of supported versions. To upgrade, provide a version listed in canUpgradeTo for the current version. Downgrades are not supported.  | [optional] |
|**Status** | Pointer to **string** | The support status of the version. | [optional] |
|**Comment** | Pointer to **string** | Additional information about the version status. | [optional] |
|**CanUpgradeTo** | Pointer to **[]string** | List of versions that a cluster running this version can be upgraded to. Only versions in this list are accepted when updating the version field via PUT.  | [optional] |

## Methods

### NewMariadbVersion

`func NewMariadbVersion() *MariadbVersion`

NewMariadbVersion instantiates a new MariadbVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMariadbVersionWithDefaults

`func NewMariadbVersionWithDefaults() *MariadbVersion`

NewMariadbVersionWithDefaults instantiates a new MariadbVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *MariadbVersion) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MariadbVersion) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MariadbVersion) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MariadbVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetStatus

`func (o *MariadbVersion) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MariadbVersion) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MariadbVersion) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MariadbVersion) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetComment

`func (o *MariadbVersion) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *MariadbVersion) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *MariadbVersion) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *MariadbVersion) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCanUpgradeTo

`func (o *MariadbVersion) GetCanUpgradeTo() []string`

GetCanUpgradeTo returns the CanUpgradeTo field if non-nil, zero value otherwise.

### GetCanUpgradeToOk

`func (o *MariadbVersion) GetCanUpgradeToOk() (*[]string, bool)`

GetCanUpgradeToOk returns a tuple with the CanUpgradeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUpgradeTo

`func (o *MariadbVersion) SetCanUpgradeTo(v []string)`

SetCanUpgradeTo sets CanUpgradeTo field to given value.

### HasCanUpgradeTo

`func (o *MariadbVersion) HasCanUpgradeTo() bool`

HasCanUpgradeTo returns a boolean if a field has been set.


