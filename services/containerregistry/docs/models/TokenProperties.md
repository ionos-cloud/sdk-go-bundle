# TokenProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Credentials** | [**NullableCredentials**](Credentials.md) |  | |
|**ExpiryDate** | Pointer to [**NullableTime**](time.Time.md) |  | [optional] |
|**Name** | **string** |  | |
|**Scopes** | Pointer to [**[]Scope**](Scope.md) |  | [optional] |
|**Status** | Pointer to **string** |  | [optional] |

## Methods

### NewTokenProperties

`func NewTokenProperties(credentials NullableCredentials, name string, ) *TokenProperties`

NewTokenProperties instantiates a new TokenProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenPropertiesWithDefaults

`func NewTokenPropertiesWithDefaults() *TokenProperties`

NewTokenPropertiesWithDefaults instantiates a new TokenProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *TokenProperties) GetCredentials() Credentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *TokenProperties) GetCredentialsOk() (*Credentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *TokenProperties) SetCredentials(v Credentials)`

SetCredentials sets Credentials field to given value.


### SetCredentialsNil

`func (o *TokenProperties) SetCredentialsNil(b bool)`

 SetCredentialsNil sets the value for Credentials to be an explicit nil

### UnsetCredentials
`func (o *TokenProperties) UnsetCredentials()`

UnsetCredentials ensures that no value is present for Credentials, not even an explicit nil
### GetExpiryDate

`func (o *TokenProperties) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *TokenProperties) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *TokenProperties) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *TokenProperties) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDateNil

`func (o *TokenProperties) SetExpiryDateNil(b bool)`

 SetExpiryDateNil sets the value for ExpiryDate to be an explicit nil

### UnsetExpiryDate
`func (o *TokenProperties) UnsetExpiryDate()`

UnsetExpiryDate ensures that no value is present for ExpiryDate, not even an explicit nil
### GetName

`func (o *TokenProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenProperties) SetName(v string)`

SetName sets Name field to given value.


### GetScopes

`func (o *TokenProperties) GetScopes() []Scope`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *TokenProperties) GetScopesOk() (*[]Scope, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *TokenProperties) SetScopes(v []Scope)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *TokenProperties) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *TokenProperties) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *TokenProperties) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetStatus

`func (o *TokenProperties) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TokenProperties) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TokenProperties) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TokenProperties) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


