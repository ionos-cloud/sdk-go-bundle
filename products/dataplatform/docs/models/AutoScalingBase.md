# AutoScalingBase

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**MinNodeCount** | **int32** | The minimum number of nodes in a Node Pool | |
|**MaxNodeCount** | **int32** | The maximum number of nodes in a Node Pool | |

## Methods

### NewAutoScalingBase

`func NewAutoScalingBase(minNodeCount int32, maxNodeCount int32, ) *AutoScalingBase`

NewAutoScalingBase instantiates a new AutoScalingBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoScalingBaseWithDefaults

`func NewAutoScalingBaseWithDefaults() *AutoScalingBase`

NewAutoScalingBaseWithDefaults instantiates a new AutoScalingBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinNodeCount

`func (o *AutoScalingBase) GetMinNodeCount() int32`

GetMinNodeCount returns the MinNodeCount field if non-nil, zero value otherwise.

### GetMinNodeCountOk

`func (o *AutoScalingBase) GetMinNodeCountOk() (*int32, bool)`

GetMinNodeCountOk returns a tuple with the MinNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNodeCount

`func (o *AutoScalingBase) SetMinNodeCount(v int32)`

SetMinNodeCount sets MinNodeCount field to given value.


### GetMaxNodeCount

`func (o *AutoScalingBase) GetMaxNodeCount() int32`

GetMaxNodeCount returns the MaxNodeCount field if non-nil, zero value otherwise.

### GetMaxNodeCountOk

`func (o *AutoScalingBase) GetMaxNodeCountOk() (*int32, bool)`

GetMaxNodeCountOk returns a tuple with the MaxNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNodeCount

`func (o *AutoScalingBase) SetMaxNodeCount(v int32)`

SetMaxNodeCount sets MaxNodeCount field to given value.



