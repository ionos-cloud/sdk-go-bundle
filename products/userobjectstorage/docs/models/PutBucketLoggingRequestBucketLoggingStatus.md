# PutBucketLoggingRequestBucketLoggingStatus

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**TargetBucket** | Pointer to **string** | Specifies the bucket where you want IONOS Object Storage to store server access logs. You can have your logs delivered to any bucket that you own, including the same bucket that is being logged. You can also configure multiple buckets to deliver their logs to the same target bucket. In this case, you should choose a different &lt;code&gt;TargetPrefix&lt;/code&gt; for each source bucket so that the delivered log files can be distinguished by key. | [optional] |
|**TargetPrefix** | Pointer to **string** | A prefix for all log object keys. If you store log files from multiple IONOS Object Storage buckets in a single bucket, you can use a prefix to distinguish which log files came from which bucket. | [optional] |

## Methods

### NewPutBucketLoggingRequestBucketLoggingStatus

`func NewPutBucketLoggingRequestBucketLoggingStatus() *PutBucketLoggingRequestBucketLoggingStatus`

NewPutBucketLoggingRequestBucketLoggingStatus instantiates a new PutBucketLoggingRequestBucketLoggingStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutBucketLoggingRequestBucketLoggingStatusWithDefaults

`func NewPutBucketLoggingRequestBucketLoggingStatusWithDefaults() *PutBucketLoggingRequestBucketLoggingStatus`

NewPutBucketLoggingRequestBucketLoggingStatusWithDefaults instantiates a new PutBucketLoggingRequestBucketLoggingStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetBucket

`func (o *PutBucketLoggingRequestBucketLoggingStatus) GetTargetBucket() string`

GetTargetBucket returns the TargetBucket field if non-nil, zero value otherwise.

### GetTargetBucketOk

`func (o *PutBucketLoggingRequestBucketLoggingStatus) GetTargetBucketOk() (*string, bool)`

GetTargetBucketOk returns a tuple with the TargetBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetBucket

`func (o *PutBucketLoggingRequestBucketLoggingStatus) SetTargetBucket(v string)`

SetTargetBucket sets TargetBucket field to given value.

### HasTargetBucket

`func (o *PutBucketLoggingRequestBucketLoggingStatus) HasTargetBucket() bool`

HasTargetBucket returns a boolean if a field has been set.

### GetTargetPrefix

`func (o *PutBucketLoggingRequestBucketLoggingStatus) GetTargetPrefix() string`

GetTargetPrefix returns the TargetPrefix field if non-nil, zero value otherwise.

### GetTargetPrefixOk

`func (o *PutBucketLoggingRequestBucketLoggingStatus) GetTargetPrefixOk() (*string, bool)`

GetTargetPrefixOk returns a tuple with the TargetPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPrefix

`func (o *PutBucketLoggingRequestBucketLoggingStatus) SetTargetPrefix(v string)`

SetTargetPrefix sets TargetPrefix field to given value.

### HasTargetPrefix

`func (o *PutBucketLoggingRequestBucketLoggingStatus) HasTargetPrefix() bool`

HasTargetPrefix returns a boolean if a field has been set.


