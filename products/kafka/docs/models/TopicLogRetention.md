# TopicLogRetention

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**RetentionTime** | Pointer to **int32** | This configuration controls the maximum time we will retain a log before we will discard old log  segments to free up space.  This represents an SLA on how soon consumers must read their data. If set to -1,  no time limit is applied.  | [optional] [default to 604800000]|
|**SegmentBytes** | Pointer to **int32** | This configuration controls the segment file size for the log. Retention and cleaning is always done a file at a time so a larger segment size means fewer files but less granular control over retention.  | [optional] [default to 1073741824]|

## Methods

### NewTopicLogRetention

`func NewTopicLogRetention() *TopicLogRetention`

NewTopicLogRetention instantiates a new TopicLogRetention object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopicLogRetentionWithDefaults

`func NewTopicLogRetentionWithDefaults() *TopicLogRetention`

NewTopicLogRetentionWithDefaults instantiates a new TopicLogRetention object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetentionTime

`func (o *TopicLogRetention) GetRetentionTime() int32`

GetRetentionTime returns the RetentionTime field if non-nil, zero value otherwise.

### GetRetentionTimeOk

`func (o *TopicLogRetention) GetRetentionTimeOk() (*int32, bool)`

GetRetentionTimeOk returns a tuple with the RetentionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionTime

`func (o *TopicLogRetention) SetRetentionTime(v int32)`

SetRetentionTime sets RetentionTime field to given value.

### HasRetentionTime

`func (o *TopicLogRetention) HasRetentionTime() bool`

HasRetentionTime returns a boolean if a field has been set.

### GetSegmentBytes

`func (o *TopicLogRetention) GetSegmentBytes() int32`

GetSegmentBytes returns the SegmentBytes field if non-nil, zero value otherwise.

### GetSegmentBytesOk

`func (o *TopicLogRetention) GetSegmentBytesOk() (*int32, bool)`

GetSegmentBytesOk returns a tuple with the SegmentBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentBytes

`func (o *TopicLogRetention) SetSegmentBytes(v int32)`

SetSegmentBytes sets SegmentBytes field to given value.

### HasSegmentBytes

`func (o *TopicLogRetention) HasSegmentBytes() bool`

HasSegmentBytes returns a boolean if a field has been set.


