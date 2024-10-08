# ProviderExternalAccountBinding

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**KeyId** | Pointer to **string** | The key ID of the external account binding. | [optional] |
|**KeySecret** | Pointer to **string** | The secret of the external account binding. | [optional] |

## Methods

### NewProviderExternalAccountBinding

`func NewProviderExternalAccountBinding() *ProviderExternalAccountBinding`

NewProviderExternalAccountBinding instantiates a new ProviderExternalAccountBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderExternalAccountBindingWithDefaults

`func NewProviderExternalAccountBindingWithDefaults() *ProviderExternalAccountBinding`

NewProviderExternalAccountBindingWithDefaults instantiates a new ProviderExternalAccountBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyId

`func (o *ProviderExternalAccountBinding) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *ProviderExternalAccountBinding) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *ProviderExternalAccountBinding) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *ProviderExternalAccountBinding) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.

### GetKeySecret

`func (o *ProviderExternalAccountBinding) GetKeySecret() string`

GetKeySecret returns the KeySecret field if non-nil, zero value otherwise.

### GetKeySecretOk

`func (o *ProviderExternalAccountBinding) GetKeySecretOk() (*string, bool)`

GetKeySecretOk returns a tuple with the KeySecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySecret

`func (o *ProviderExternalAccountBinding) SetKeySecret(v string)`

SetKeySecret sets KeySecret field to given value.

### HasKeySecret

`func (o *ProviderExternalAccountBinding) HasKeySecret() bool`

HasKeySecret returns a boolean if a field has been set.


