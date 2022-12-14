# PostTokenProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**ExpiryDate** | Pointer to [**NullableTime**](time.Time.md) |  | [optional] |
|**Name** | **string** |  | |
|**Scopes** | Pointer to [**[]Scope**](Scope.md) |  | [optional] |
|**Status** | Pointer to **string** |  | [optional] |

## Methods

### NewPostTokenProperties

`func NewPostTokenProperties(name string, ) *PostTokenProperties`

NewPostTokenProperties instantiates a new PostTokenProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostTokenPropertiesWithDefaults

`func NewPostTokenPropertiesWithDefaults() *PostTokenProperties`

NewPostTokenPropertiesWithDefaults instantiates a new PostTokenProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiryDate

`func (o *PostTokenProperties) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *PostTokenProperties) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *PostTokenProperties) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *PostTokenProperties) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDateNil

`func (o *PostTokenProperties) SetExpiryDateNil(b bool)`

 SetExpiryDateNil sets the value for ExpiryDate to be an explicit nil

### UnsetExpiryDate
`func (o *PostTokenProperties) UnsetExpiryDate()`

UnsetExpiryDate ensures that no value is present for ExpiryDate, not even an explicit nil
### GetName

`func (o *PostTokenProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostTokenProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostTokenProperties) SetName(v string)`

SetName sets Name field to given value.


### GetScopes

`func (o *PostTokenProperties) GetScopes() []Scope`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *PostTokenProperties) GetScopesOk() (*[]Scope, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *PostTokenProperties) SetScopes(v []Scope)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *PostTokenProperties) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *PostTokenProperties) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *PostTokenProperties) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetStatus

`func (o *PostTokenProperties) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PostTokenProperties) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PostTokenProperties) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PostTokenProperties) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


