# AccessKey

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Description** | **string** | Description of the Access key. | |
|**AccessKey** | **string** | Access key metadata is a string of 92 characters. | [readonly] |
|**SecretKey** | **string** | The secret key of the Access key. | [readonly] |
|**CanonicalUserId** | Pointer to **string** | The canonical user ID which is valid for user-owned buckets. | [optional] [readonly] |
|**ContractUserId** | Pointer to **string** | The contract user ID which is valid for contract-owned buckets. | [optional] [readonly] |

## Methods

### NewAccessKey

`func NewAccessKey(description string, accessKey string, secretKey string, ) *AccessKey`

NewAccessKey instantiates a new AccessKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessKeyWithDefaults

`func NewAccessKeyWithDefaults() *AccessKey`

NewAccessKeyWithDefaults instantiates a new AccessKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AccessKey) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccessKey) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccessKey) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetAccessKey

`func (o *AccessKey) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *AccessKey) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *AccessKey) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### GetSecretKey

`func (o *AccessKey) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *AccessKey) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *AccessKey) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.


### GetCanonicalUserId

`func (o *AccessKey) GetCanonicalUserId() string`

GetCanonicalUserId returns the CanonicalUserId field if non-nil, zero value otherwise.

### GetCanonicalUserIdOk

`func (o *AccessKey) GetCanonicalUserIdOk() (*string, bool)`

GetCanonicalUserIdOk returns a tuple with the CanonicalUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalUserId

`func (o *AccessKey) SetCanonicalUserId(v string)`

SetCanonicalUserId sets CanonicalUserId field to given value.

### HasCanonicalUserId

`func (o *AccessKey) HasCanonicalUserId() bool`

HasCanonicalUserId returns a boolean if a field has been set.

### GetContractUserId

`func (o *AccessKey) GetContractUserId() string`

GetContractUserId returns the ContractUserId field if non-nil, zero value otherwise.

### GetContractUserIdOk

`func (o *AccessKey) GetContractUserIdOk() (*string, bool)`

GetContractUserIdOk returns a tuple with the ContractUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractUserId

`func (o *AccessKey) SetContractUserId(v string)`

SetContractUserId sets ContractUserId field to given value.

### HasContractUserId

`func (o *AccessKey) HasContractUserId() bool`

HasContractUserId returns a boolean if a field has been set.


