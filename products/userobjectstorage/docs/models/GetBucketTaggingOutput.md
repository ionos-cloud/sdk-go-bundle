# GetBucketTaggingOutput

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**TagSet** | [**BucketTagSet**](BucketTagSet.md) |  | |

## Methods

### NewGetBucketTaggingOutput

`func NewGetBucketTaggingOutput(tagSet BucketTagSet, ) *GetBucketTaggingOutput`

NewGetBucketTaggingOutput instantiates a new GetBucketTaggingOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBucketTaggingOutputWithDefaults

`func NewGetBucketTaggingOutputWithDefaults() *GetBucketTaggingOutput`

NewGetBucketTaggingOutputWithDefaults instantiates a new GetBucketTaggingOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagSet

`func (o *GetBucketTaggingOutput) GetTagSet() BucketTagSet`

GetTagSet returns the TagSet field if non-nil, zero value otherwise.

### GetTagSetOk

`func (o *GetBucketTaggingOutput) GetTagSetOk() (*BucketTagSet, bool)`

GetTagSetOk returns a tuple with the TagSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagSet

`func (o *GetBucketTaggingOutput) SetTagSet(v BucketTagSet)`

SetTagSet sets TagSet field to given value.



