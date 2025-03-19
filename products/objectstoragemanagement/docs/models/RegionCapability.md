# RegionCapability

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Iam** | Pointer to **bool** | Indicates if IAM policy based access is supported | [optional] |
|**S3select** | Pointer to **bool** | Indicates if S3 Select is supported | [optional] |

## Methods

### NewRegionCapability

`func NewRegionCapability() *RegionCapability`

NewRegionCapability instantiates a new RegionCapability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionCapabilityWithDefaults

`func NewRegionCapabilityWithDefaults() *RegionCapability`

NewRegionCapabilityWithDefaults instantiates a new RegionCapability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIam

`func (o *RegionCapability) GetIam() bool`

GetIam returns the Iam field if non-nil, zero value otherwise.

### GetIamOk

`func (o *RegionCapability) GetIamOk() (*bool, bool)`

GetIamOk returns a tuple with the Iam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIam

`func (o *RegionCapability) SetIam(v bool)`

SetIam sets Iam field to given value.

### HasIam

`func (o *RegionCapability) HasIam() bool`

HasIam returns a boolean if a field has been set.

### GetS3select

`func (o *RegionCapability) GetS3select() bool`

GetS3select returns the S3select field if non-nil, zero value otherwise.

### GetS3selectOk

`func (o *RegionCapability) GetS3selectOk() (*bool, bool)`

GetS3selectOk returns a tuple with the S3select field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3select

`func (o *RegionCapability) SetS3select(v bool)`

SetS3select sets S3select field to given value.

### HasS3select

`func (o *RegionCapability) HasS3select() bool`

HasS3select returns a boolean if a field has been set.


