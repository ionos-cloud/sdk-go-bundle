# Jwt

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Token** | Pointer to **string** | JSON Web Token (JWT) Base64url strings separated by dots | [optional] |

## Methods

### NewJwt

`func NewJwt() *Jwt`

NewJwt instantiates a new Jwt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJwtWithDefaults

`func NewJwtWithDefaults() *Jwt`

NewJwtWithDefaults instantiates a new Jwt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *Jwt) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Jwt) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Jwt) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *Jwt) HasToken() bool`

HasToken returns a boolean if a field has been set.


