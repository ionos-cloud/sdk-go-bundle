# ESPEncryption

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**DiffieHellmanGroup** | Pointer to **string** | The Diffie-Hellman Group to use for IPSec Encryption.\\ Options:   - 15-MODP3072   - 16-MODP4096   - 19-ECP256   - 20-ECP384   - 21-ECP521   - 28-ECP256BP   - 29-ECP384BP   - 30-ECP512BP  | [optional] |
|**EncryptionAlgorithm** | Pointer to **string** | The encryption algorithm to use for IPSec Encryption.\\ Options: - AES128-CTR - AES256-CTR - AES128-GCM-16 - AES256-GCM-16 - AES128-GCM-12 - AES256-GCM-12 - AES128-CCM-12 - AES256-CCM-12 - AES128 - AES256  | [optional] |
|**IntegrityAlgorithm** | Pointer to **string** | The integrity algorithm to use for IPSec Encryption.\\ Options: - SHA256 - SHA384 - SHA512 - AES-XCBC  | [optional] |
|**Lifetime** | Pointer to **int32** | The phase lifetime in seconds. | [optional] [default to 3600]|

## Methods

### NewESPEncryption

`func NewESPEncryption() *ESPEncryption`

NewESPEncryption instantiates a new ESPEncryption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewESPEncryptionWithDefaults

`func NewESPEncryptionWithDefaults() *ESPEncryption`

NewESPEncryptionWithDefaults instantiates a new ESPEncryption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiffieHellmanGroup

`func (o *ESPEncryption) GetDiffieHellmanGroup() string`

GetDiffieHellmanGroup returns the DiffieHellmanGroup field if non-nil, zero value otherwise.

### GetDiffieHellmanGroupOk

`func (o *ESPEncryption) GetDiffieHellmanGroupOk() (*string, bool)`

GetDiffieHellmanGroupOk returns a tuple with the DiffieHellmanGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffieHellmanGroup

`func (o *ESPEncryption) SetDiffieHellmanGroup(v string)`

SetDiffieHellmanGroup sets DiffieHellmanGroup field to given value.

### HasDiffieHellmanGroup

`func (o *ESPEncryption) HasDiffieHellmanGroup() bool`

HasDiffieHellmanGroup returns a boolean if a field has been set.

### GetEncryptionAlgorithm

`func (o *ESPEncryption) GetEncryptionAlgorithm() string`

GetEncryptionAlgorithm returns the EncryptionAlgorithm field if non-nil, zero value otherwise.

### GetEncryptionAlgorithmOk

`func (o *ESPEncryption) GetEncryptionAlgorithmOk() (*string, bool)`

GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAlgorithm

`func (o *ESPEncryption) SetEncryptionAlgorithm(v string)`

SetEncryptionAlgorithm sets EncryptionAlgorithm field to given value.

### HasEncryptionAlgorithm

`func (o *ESPEncryption) HasEncryptionAlgorithm() bool`

HasEncryptionAlgorithm returns a boolean if a field has been set.

### GetIntegrityAlgorithm

`func (o *ESPEncryption) GetIntegrityAlgorithm() string`

GetIntegrityAlgorithm returns the IntegrityAlgorithm field if non-nil, zero value otherwise.

### GetIntegrityAlgorithmOk

`func (o *ESPEncryption) GetIntegrityAlgorithmOk() (*string, bool)`

GetIntegrityAlgorithmOk returns a tuple with the IntegrityAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrityAlgorithm

`func (o *ESPEncryption) SetIntegrityAlgorithm(v string)`

SetIntegrityAlgorithm sets IntegrityAlgorithm field to given value.

### HasIntegrityAlgorithm

`func (o *ESPEncryption) HasIntegrityAlgorithm() bool`

HasIntegrityAlgorithm returns a boolean if a field has been set.

### GetLifetime

`func (o *ESPEncryption) GetLifetime() int32`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *ESPEncryption) GetLifetimeOk() (*int32, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *ESPEncryption) SetLifetime(v int32)`

SetLifetime sets Lifetime field to given value.

### HasLifetime

`func (o *ESPEncryption) HasLifetime() bool`

HasLifetime returns a boolean if a field has been set.


