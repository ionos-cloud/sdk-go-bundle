# BucketPolicyStatementCondition

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**IpAddress** | Pointer to [**BucketPolicyStatementConditionIpAddress**](BucketPolicyStatementConditionIpAddress.md) |  | [optional] |
|**NotIpAddress** | Pointer to [**BucketPolicyStatementConditionIpAddress**](BucketPolicyStatementConditionIpAddress.md) |  | [optional] |
|**DateGreaterThan** | Pointer to [**BucketPolicyStatementConditionDateGreaterThan**](BucketPolicyStatementConditionDateGreaterThan.md) |  | [optional] |
|**DateLessThan** | Pointer to [**BucketPolicyStatementConditionDateLessThan**](BucketPolicyStatementConditionDateLessThan.md) |  | [optional] |

## Methods

### NewBucketPolicyStatementCondition

`func NewBucketPolicyStatementCondition() *BucketPolicyStatementCondition`

NewBucketPolicyStatementCondition instantiates a new BucketPolicyStatementCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketPolicyStatementConditionWithDefaults

`func NewBucketPolicyStatementConditionWithDefaults() *BucketPolicyStatementCondition`

NewBucketPolicyStatementConditionWithDefaults instantiates a new BucketPolicyStatementCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddress

`func (o *BucketPolicyStatementCondition) GetIpAddress() BucketPolicyStatementConditionIpAddress`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *BucketPolicyStatementCondition) GetIpAddressOk() (*BucketPolicyStatementConditionIpAddress, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *BucketPolicyStatementCondition) SetIpAddress(v BucketPolicyStatementConditionIpAddress)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *BucketPolicyStatementCondition) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetNotIpAddress

`func (o *BucketPolicyStatementCondition) GetNotIpAddress() BucketPolicyStatementConditionIpAddress`

GetNotIpAddress returns the NotIpAddress field if non-nil, zero value otherwise.

### GetNotIpAddressOk

`func (o *BucketPolicyStatementCondition) GetNotIpAddressOk() (*BucketPolicyStatementConditionIpAddress, bool)`

GetNotIpAddressOk returns a tuple with the NotIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotIpAddress

`func (o *BucketPolicyStatementCondition) SetNotIpAddress(v BucketPolicyStatementConditionIpAddress)`

SetNotIpAddress sets NotIpAddress field to given value.

### HasNotIpAddress

`func (o *BucketPolicyStatementCondition) HasNotIpAddress() bool`

HasNotIpAddress returns a boolean if a field has been set.

### GetDateGreaterThan

`func (o *BucketPolicyStatementCondition) GetDateGreaterThan() BucketPolicyStatementConditionDateGreaterThan`

GetDateGreaterThan returns the DateGreaterThan field if non-nil, zero value otherwise.

### GetDateGreaterThanOk

`func (o *BucketPolicyStatementCondition) GetDateGreaterThanOk() (*BucketPolicyStatementConditionDateGreaterThan, bool)`

GetDateGreaterThanOk returns a tuple with the DateGreaterThan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateGreaterThan

`func (o *BucketPolicyStatementCondition) SetDateGreaterThan(v BucketPolicyStatementConditionDateGreaterThan)`

SetDateGreaterThan sets DateGreaterThan field to given value.

### HasDateGreaterThan

`func (o *BucketPolicyStatementCondition) HasDateGreaterThan() bool`

HasDateGreaterThan returns a boolean if a field has been set.

### GetDateLessThan

`func (o *BucketPolicyStatementCondition) GetDateLessThan() BucketPolicyStatementConditionDateLessThan`

GetDateLessThan returns the DateLessThan field if non-nil, zero value otherwise.

### GetDateLessThanOk

`func (o *BucketPolicyStatementCondition) GetDateLessThanOk() (*BucketPolicyStatementConditionDateLessThan, bool)`

GetDateLessThanOk returns a tuple with the DateLessThan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateLessThan

`func (o *BucketPolicyStatementCondition) SetDateLessThan(v BucketPolicyStatementConditionDateLessThan)`

SetDateLessThan sets DateLessThan field to given value.

### HasDateLessThan

`func (o *BucketPolicyStatementCondition) HasDateLessThan() bool`

HasDateLessThan returns a boolean if a field has been set.


