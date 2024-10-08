# PatchTokenInput

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) |  | [optional] |
|**Scopes** | Pointer to [**[]Scope**](Scope.md) |  | [optional] |
|**Status** | Pointer to **string** |  | [optional] |

## Methods

### NewPatchTokenInput

`func NewPatchTokenInput() *PatchTokenInput`

NewPatchTokenInput instantiates a new PatchTokenInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchTokenInputWithDefaults

`func NewPatchTokenInputWithDefaults() *PatchTokenInput`

NewPatchTokenInputWithDefaults instantiates a new PatchTokenInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiryDate

`func (o *PatchTokenInput) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *PatchTokenInput) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *PatchTokenInput) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *PatchTokenInput) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetScopes

`func (o *PatchTokenInput) GetScopes() []Scope`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *PatchTokenInput) GetScopesOk() (*[]Scope, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *PatchTokenInput) SetScopes(v []Scope)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *PatchTokenInput) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetStatus

`func (o *PatchTokenInput) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchTokenInput) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchTokenInput) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchTokenInput) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


