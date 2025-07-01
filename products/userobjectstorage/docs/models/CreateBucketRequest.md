# CreateBucketRequest

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**CreateBucketConfiguration** | Pointer to [**CreateBucketRequestCreateBucketConfiguration**](CreateBucketRequestCreateBucketConfiguration.md) |  | [optional] |

## Methods

### NewCreateBucketRequest

`func NewCreateBucketRequest() *CreateBucketRequest`

NewCreateBucketRequest instantiates a new CreateBucketRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBucketRequestWithDefaults

`func NewCreateBucketRequestWithDefaults() *CreateBucketRequest`

NewCreateBucketRequestWithDefaults instantiates a new CreateBucketRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateBucketConfiguration

`func (o *CreateBucketRequest) GetCreateBucketConfiguration() CreateBucketRequestCreateBucketConfiguration`

GetCreateBucketConfiguration returns the CreateBucketConfiguration field if non-nil, zero value otherwise.

### GetCreateBucketConfigurationOk

`func (o *CreateBucketRequest) GetCreateBucketConfigurationOk() (*CreateBucketRequestCreateBucketConfiguration, bool)`

GetCreateBucketConfigurationOk returns a tuple with the CreateBucketConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBucketConfiguration

`func (o *CreateBucketRequest) SetCreateBucketConfiguration(v CreateBucketRequestCreateBucketConfiguration)`

SetCreateBucketConfiguration sets CreateBucketConfiguration field to given value.

### HasCreateBucketConfiguration

`func (o *CreateBucketRequest) HasCreateBucketConfiguration() bool`

HasCreateBucketConfiguration returns a boolean if a field has been set.


