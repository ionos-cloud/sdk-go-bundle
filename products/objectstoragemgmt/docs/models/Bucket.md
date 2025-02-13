# Bucket

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Region** | **string** | The region where the bucket is located | |
|**Website** | **string** | The website URL for the bucket | |

## Methods

### NewBucket

`func NewBucket(region string, website string, ) *Bucket`

NewBucket instantiates a new Bucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketWithDefaults

`func NewBucketWithDefaults() *Bucket`

NewBucketWithDefaults instantiates a new Bucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *Bucket) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Bucket) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Bucket) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetWebsite

`func (o *Bucket) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *Bucket) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *Bucket) SetWebsite(v string)`

SetWebsite sets Website field to given value.



