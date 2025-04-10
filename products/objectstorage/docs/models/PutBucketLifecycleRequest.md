# PutBucketLifecycleRequest

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Rules** | Pointer to [**[]Rule**](Rule.md) | Container for a lifecycle rules. | [optional] |

## Methods

### NewPutBucketLifecycleRequest

`func NewPutBucketLifecycleRequest() *PutBucketLifecycleRequest`

NewPutBucketLifecycleRequest instantiates a new PutBucketLifecycleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutBucketLifecycleRequestWithDefaults

`func NewPutBucketLifecycleRequestWithDefaults() *PutBucketLifecycleRequest`

NewPutBucketLifecycleRequestWithDefaults instantiates a new PutBucketLifecycleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *PutBucketLifecycleRequest) GetRules() []Rule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *PutBucketLifecycleRequest) GetRulesOk() (*[]Rule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *PutBucketLifecycleRequest) SetRules(v []Rule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *PutBucketLifecycleRequest) HasRules() bool`

HasRules returns a boolean if a field has been set.


