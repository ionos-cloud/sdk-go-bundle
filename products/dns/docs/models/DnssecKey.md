# DnssecKey

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**KeyTag** | Pointer to **int32** |  | [optional] |
|**DigestAlgorithmMnemonic** | Pointer to **string** | A string that denotes the digest algorithm. This value must conform to the guidelines in [RFC-8624 Section 3.3](https://datatracker.ietf.org/doc/html/rfc8624#section-3.3).  | [optional] |
|**Digest** | Pointer to **string** |  | [optional] |
|**KeyData** | Pointer to [**KeyData**](KeyData.md) |  | [optional] |
|**ComposedKeyData** | Pointer to **string** | Represents the composed value of the The RDATA for a DNSKEY. The format should be the following: Flags | Protocol | Algorithm | Public Key The values must conform to the guidelines in [RFC-4034 Section 2.1](https://www.rfc-editor.org/rfc/rfc4034#section-2.1).  | [optional] |

## Methods

### NewDnssecKey

`func NewDnssecKey() *DnssecKey`

NewDnssecKey instantiates a new DnssecKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnssecKeyWithDefaults

`func NewDnssecKeyWithDefaults() *DnssecKey`

NewDnssecKeyWithDefaults instantiates a new DnssecKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyTag

`func (o *DnssecKey) GetKeyTag() int32`

GetKeyTag returns the KeyTag field if non-nil, zero value otherwise.

### GetKeyTagOk

`func (o *DnssecKey) GetKeyTagOk() (*int32, bool)`

GetKeyTagOk returns a tuple with the KeyTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyTag

`func (o *DnssecKey) SetKeyTag(v int32)`

SetKeyTag sets KeyTag field to given value.

### HasKeyTag

`func (o *DnssecKey) HasKeyTag() bool`

HasKeyTag returns a boolean if a field has been set.

### GetDigestAlgorithmMnemonic

`func (o *DnssecKey) GetDigestAlgorithmMnemonic() string`

GetDigestAlgorithmMnemonic returns the DigestAlgorithmMnemonic field if non-nil, zero value otherwise.

### GetDigestAlgorithmMnemonicOk

`func (o *DnssecKey) GetDigestAlgorithmMnemonicOk() (*string, bool)`

GetDigestAlgorithmMnemonicOk returns a tuple with the DigestAlgorithmMnemonic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestAlgorithmMnemonic

`func (o *DnssecKey) SetDigestAlgorithmMnemonic(v string)`

SetDigestAlgorithmMnemonic sets DigestAlgorithmMnemonic field to given value.

### HasDigestAlgorithmMnemonic

`func (o *DnssecKey) HasDigestAlgorithmMnemonic() bool`

HasDigestAlgorithmMnemonic returns a boolean if a field has been set.

### GetDigest

`func (o *DnssecKey) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *DnssecKey) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *DnssecKey) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *DnssecKey) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetKeyData

`func (o *DnssecKey) GetKeyData() KeyData`

GetKeyData returns the KeyData field if non-nil, zero value otherwise.

### GetKeyDataOk

`func (o *DnssecKey) GetKeyDataOk() (*KeyData, bool)`

GetKeyDataOk returns a tuple with the KeyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyData

`func (o *DnssecKey) SetKeyData(v KeyData)`

SetKeyData sets KeyData field to given value.

### HasKeyData

`func (o *DnssecKey) HasKeyData() bool`

HasKeyData returns a boolean if a field has been set.

### GetComposedKeyData

`func (o *DnssecKey) GetComposedKeyData() string`

GetComposedKeyData returns the ComposedKeyData field if non-nil, zero value otherwise.

### GetComposedKeyDataOk

`func (o *DnssecKey) GetComposedKeyDataOk() (*string, bool)`

GetComposedKeyDataOk returns a tuple with the ComposedKeyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComposedKeyData

`func (o *DnssecKey) SetComposedKeyData(v string)`

SetComposedKeyData sets ComposedKeyData field to given value.

### HasComposedKeyData

`func (o *DnssecKey) HasComposedKeyData() bool`

HasComposedKeyData returns a boolean if a field has been set.


