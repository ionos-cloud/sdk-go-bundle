# Tokens

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Tokens** | Pointer to [**[]Token**](Token.md) | Array of items in that collection. | [optional] [readonly] |

## Methods

### NewTokens

`func NewTokens() *Tokens`

NewTokens instantiates a new Tokens object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokensWithDefaults

`func NewTokensWithDefaults() *Tokens`

NewTokensWithDefaults instantiates a new Tokens object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokens

`func (o *Tokens) GetTokens() []Token`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *Tokens) GetTokensOk() (*[]Token, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *Tokens) SetTokens(v []Token)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *Tokens) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


