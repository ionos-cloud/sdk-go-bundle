# PatchUserProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Password** | Pointer to **string** |  | [optional] |
|**Roles** | Pointer to [**[]UserRoles**](UserRoles.md) |  | [optional] |

## Methods

### NewPatchUserProperties

`func NewPatchUserProperties() *PatchUserProperties`

NewPatchUserProperties instantiates a new PatchUserProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchUserPropertiesWithDefaults

`func NewPatchUserPropertiesWithDefaults() *PatchUserProperties`

NewPatchUserPropertiesWithDefaults instantiates a new PatchUserProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *PatchUserProperties) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PatchUserProperties) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PatchUserProperties) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PatchUserProperties) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRoles

`func (o *PatchUserProperties) GetRoles() []UserRoles`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *PatchUserProperties) GetRolesOk() (*[]UserRoles, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *PatchUserProperties) SetRoles(v []UserRoles)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *PatchUserProperties) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


