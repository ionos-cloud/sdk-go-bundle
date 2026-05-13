# BucketTagSet

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Tag** | Pointer to [**[]Tag**](Tag.md) |  | [optional] |

## Methods

### NewBucketTagSet

`func NewBucketTagSet() *BucketTagSet`

NewBucketTagSet instantiates a new BucketTagSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketTagSetWithDefaults

`func NewBucketTagSetWithDefaults() *BucketTagSet`

NewBucketTagSetWithDefaults instantiates a new BucketTagSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTag

`func (o *BucketTagSet) GetTag() []Tag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *BucketTagSet) GetTagOk() (*[]Tag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *BucketTagSet) SetTag(v []Tag)`

SetTag sets Tag field to given value.

### HasTag

`func (o *BucketTagSet) HasTag() bool`

HasTag returns a boolean if a field has been set.


