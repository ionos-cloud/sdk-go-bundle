# GetBucketPolicyStatusOutput

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**PolicyStatus** | Pointer to [**PolicyStatus**](PolicyStatus.md) |  | [optional] |

## Methods

### NewGetBucketPolicyStatusOutput

`func NewGetBucketPolicyStatusOutput() *GetBucketPolicyStatusOutput`

NewGetBucketPolicyStatusOutput instantiates a new GetBucketPolicyStatusOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBucketPolicyStatusOutputWithDefaults

`func NewGetBucketPolicyStatusOutputWithDefaults() *GetBucketPolicyStatusOutput`

NewGetBucketPolicyStatusOutputWithDefaults instantiates a new GetBucketPolicyStatusOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyStatus

`func (o *GetBucketPolicyStatusOutput) GetPolicyStatus() PolicyStatus`

GetPolicyStatus returns the PolicyStatus field if non-nil, zero value otherwise.

### GetPolicyStatusOk

`func (o *GetBucketPolicyStatusOutput) GetPolicyStatusOk() (*PolicyStatus, bool)`

GetPolicyStatusOk returns a tuple with the PolicyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyStatus

`func (o *GetBucketPolicyStatusOutput) SetPolicyStatus(v PolicyStatus)`

SetPolicyStatus sets PolicyStatus field to given value.

### HasPolicyStatus

`func (o *GetBucketPolicyStatusOutput) HasPolicyStatus() bool`

HasPolicyStatus returns a boolean if a field has been set.


