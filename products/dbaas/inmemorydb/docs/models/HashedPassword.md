# HashedPassword

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Algorithm** | **string** |  | |
|**Hash** | **string** |  | |

## Methods

### NewHashedPassword

`func NewHashedPassword(algorithm string, hash string, ) *HashedPassword`

NewHashedPassword instantiates a new HashedPassword object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHashedPasswordWithDefaults

`func NewHashedPasswordWithDefaults() *HashedPassword`

NewHashedPasswordWithDefaults instantiates a new HashedPassword object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *HashedPassword) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *HashedPassword) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *HashedPassword) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.


### GetHash

`func (o *HashedPassword) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *HashedPassword) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *HashedPassword) SetHash(v string)`

SetHash sets Hash field to given value.



