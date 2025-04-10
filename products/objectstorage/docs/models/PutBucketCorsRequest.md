# PutBucketCorsRequest

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**CORSRules** | Pointer to [**[]CORSRule**](CORSRule.md) | A set of origins and methods (cross-origin access that you want to allow). You can add up to 100 rules to the configuration. | [optional] |

## Methods

### NewPutBucketCorsRequest

`func NewPutBucketCorsRequest() *PutBucketCorsRequest`

NewPutBucketCorsRequest instantiates a new PutBucketCorsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutBucketCorsRequestWithDefaults

`func NewPutBucketCorsRequestWithDefaults() *PutBucketCorsRequest`

NewPutBucketCorsRequestWithDefaults instantiates a new PutBucketCorsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCORSRules

`func (o *PutBucketCorsRequest) GetCORSRules() []CORSRule`

GetCORSRules returns the CORSRules field if non-nil, zero value otherwise.

### GetCORSRulesOk

`func (o *PutBucketCorsRequest) GetCORSRulesOk() (*[]CORSRule, bool)`

GetCORSRulesOk returns a tuple with the CORSRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCORSRules

`func (o *PutBucketCorsRequest) SetCORSRules(v []CORSRule)`

SetCORSRules sets CORSRules field to given value.

### HasCORSRules

`func (o *PutBucketCorsRequest) HasCORSRules() bool`

HasCORSRules returns a boolean if a field has been set.


