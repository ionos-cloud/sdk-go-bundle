# KeyData

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Flags** | Pointer to **int32** | Represents the key&#39;s metadata and usage information. | [optional] |
|**PubKey** | Pointer to **string** | Represents the public key data in Base64 encoding. | [optional] |

## Methods

### NewKeyData

`func NewKeyData() *KeyData`

NewKeyData instantiates a new KeyData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyDataWithDefaults

`func NewKeyDataWithDefaults() *KeyData`

NewKeyDataWithDefaults instantiates a new KeyData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlags

`func (o *KeyData) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *KeyData) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *KeyData) SetFlags(v int32)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *KeyData) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetPubKey

`func (o *KeyData) GetPubKey() string`

GetPubKey returns the PubKey field if non-nil, zero value otherwise.

### GetPubKeyOk

`func (o *KeyData) GetPubKeyOk() (*string, bool)`

GetPubKeyOk returns a tuple with the PubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubKey

`func (o *KeyData) SetPubKey(v string)`

SetPubKey sets PubKey field to given value.

### HasPubKey

`func (o *KeyData) HasPubKey() bool`

HasPubKey returns a boolean if a field has been set.


