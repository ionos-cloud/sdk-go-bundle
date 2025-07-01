# PutBucketVersioningRequest

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**VersioningConfiguration** | [**PutBucketVersioningRequestVersioningConfiguration**](PutBucketVersioningRequestVersioningConfiguration.md) |  | |

## Methods

### NewPutBucketVersioningRequest

`func NewPutBucketVersioningRequest(versioningConfiguration PutBucketVersioningRequestVersioningConfiguration, ) *PutBucketVersioningRequest`

NewPutBucketVersioningRequest instantiates a new PutBucketVersioningRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutBucketVersioningRequestWithDefaults

`func NewPutBucketVersioningRequestWithDefaults() *PutBucketVersioningRequest`

NewPutBucketVersioningRequestWithDefaults instantiates a new PutBucketVersioningRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersioningConfiguration

`func (o *PutBucketVersioningRequest) GetVersioningConfiguration() PutBucketVersioningRequestVersioningConfiguration`

GetVersioningConfiguration returns the VersioningConfiguration field if non-nil, zero value otherwise.

### GetVersioningConfigurationOk

`func (o *PutBucketVersioningRequest) GetVersioningConfigurationOk() (*PutBucketVersioningRequestVersioningConfiguration, bool)`

GetVersioningConfigurationOk returns a tuple with the VersioningConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersioningConfiguration

`func (o *PutBucketVersioningRequest) SetVersioningConfiguration(v PutBucketVersioningRequestVersioningConfiguration)`

SetVersioningConfiguration sets VersioningConfiguration field to given value.



