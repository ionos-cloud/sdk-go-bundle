# PutBucketEncryptionRequest

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Rules** | Pointer to [**[]ServerSideEncryptionRule**](ServerSideEncryptionRule.md) |  | [optional] |

## Methods

### NewPutBucketEncryptionRequest

`func NewPutBucketEncryptionRequest() *PutBucketEncryptionRequest`

NewPutBucketEncryptionRequest instantiates a new PutBucketEncryptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutBucketEncryptionRequestWithDefaults

`func NewPutBucketEncryptionRequestWithDefaults() *PutBucketEncryptionRequest`

NewPutBucketEncryptionRequestWithDefaults instantiates a new PutBucketEncryptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *PutBucketEncryptionRequest) GetRules() []ServerSideEncryptionRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *PutBucketEncryptionRequest) GetRulesOk() (*[]ServerSideEncryptionRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *PutBucketEncryptionRequest) SetRules(v []ServerSideEncryptionRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *PutBucketEncryptionRequest) HasRules() bool`

HasRules returns a boolean if a field has been set.


