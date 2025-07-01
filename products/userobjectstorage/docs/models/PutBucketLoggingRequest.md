# PutBucketLoggingRequest

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**BucketLoggingStatus** | [**PutBucketLoggingRequestBucketLoggingStatus**](PutBucketLoggingRequestBucketLoggingStatus.md) |  | |

## Methods

### NewPutBucketLoggingRequest

`func NewPutBucketLoggingRequest(bucketLoggingStatus PutBucketLoggingRequestBucketLoggingStatus, ) *PutBucketLoggingRequest`

NewPutBucketLoggingRequest instantiates a new PutBucketLoggingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutBucketLoggingRequestWithDefaults

`func NewPutBucketLoggingRequestWithDefaults() *PutBucketLoggingRequest`

NewPutBucketLoggingRequestWithDefaults instantiates a new PutBucketLoggingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketLoggingStatus

`func (o *PutBucketLoggingRequest) GetBucketLoggingStatus() PutBucketLoggingRequestBucketLoggingStatus`

GetBucketLoggingStatus returns the BucketLoggingStatus field if non-nil, zero value otherwise.

### GetBucketLoggingStatusOk

`func (o *PutBucketLoggingRequest) GetBucketLoggingStatusOk() (*PutBucketLoggingRequestBucketLoggingStatus, bool)`

GetBucketLoggingStatusOk returns a tuple with the BucketLoggingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketLoggingStatus

`func (o *PutBucketLoggingRequest) SetBucketLoggingStatus(v PutBucketLoggingRequestBucketLoggingStatus)`

SetBucketLoggingStatus sets BucketLoggingStatus field to given value.



