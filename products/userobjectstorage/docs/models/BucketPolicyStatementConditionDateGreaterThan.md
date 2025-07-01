# BucketPolicyStatementConditionDateGreaterThan

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**AwsCurrentTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] |
|**AwsEpochTime** | Pointer to **int32** |  | [optional] |

## Methods

### NewBucketPolicyStatementConditionDateGreaterThan

`func NewBucketPolicyStatementConditionDateGreaterThan() *BucketPolicyStatementConditionDateGreaterThan`

NewBucketPolicyStatementConditionDateGreaterThan instantiates a new BucketPolicyStatementConditionDateGreaterThan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketPolicyStatementConditionDateGreaterThanWithDefaults

`func NewBucketPolicyStatementConditionDateGreaterThanWithDefaults() *BucketPolicyStatementConditionDateGreaterThan`

NewBucketPolicyStatementConditionDateGreaterThanWithDefaults instantiates a new BucketPolicyStatementConditionDateGreaterThan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsCurrentTime

`func (o *BucketPolicyStatementConditionDateGreaterThan) GetAwsCurrentTime() time.Time`

GetAwsCurrentTime returns the AwsCurrentTime field if non-nil, zero value otherwise.

### GetAwsCurrentTimeOk

`func (o *BucketPolicyStatementConditionDateGreaterThan) GetAwsCurrentTimeOk() (*time.Time, bool)`

GetAwsCurrentTimeOk returns a tuple with the AwsCurrentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsCurrentTime

`func (o *BucketPolicyStatementConditionDateGreaterThan) SetAwsCurrentTime(v time.Time)`

SetAwsCurrentTime sets AwsCurrentTime field to given value.

### HasAwsCurrentTime

`func (o *BucketPolicyStatementConditionDateGreaterThan) HasAwsCurrentTime() bool`

HasAwsCurrentTime returns a boolean if a field has been set.

### GetAwsEpochTime

`func (o *BucketPolicyStatementConditionDateGreaterThan) GetAwsEpochTime() int32`

GetAwsEpochTime returns the AwsEpochTime field if non-nil, zero value otherwise.

### GetAwsEpochTimeOk

`func (o *BucketPolicyStatementConditionDateGreaterThan) GetAwsEpochTimeOk() (*int32, bool)`

GetAwsEpochTimeOk returns a tuple with the AwsEpochTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsEpochTime

`func (o *BucketPolicyStatementConditionDateGreaterThan) SetAwsEpochTime(v int32)`

SetAwsEpochTime sets AwsEpochTime field to given value.

### HasAwsEpochTime

`func (o *BucketPolicyStatementConditionDateGreaterThan) HasAwsEpochTime() bool`

HasAwsEpochTime returns a boolean if a field has been set.


