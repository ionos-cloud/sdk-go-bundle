# UserPassword

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Algorithm** | **string** |  | |
|**Hash** | **string** |  | |

## Methods

### NewUserPassword

`func NewUserPassword(algorithm string, hash string, ) *UserPassword`

NewUserPassword instantiates a new UserPassword object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPasswordWithDefaults

`func NewUserPasswordWithDefaults() *UserPassword`

NewUserPasswordWithDefaults instantiates a new UserPassword object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *UserPassword) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *UserPassword) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *UserPassword) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.


### GetHash

`func (o *UserPassword) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *UserPassword) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *UserPassword) SetHash(v string)`

SetHash sets Hash field to given value.



