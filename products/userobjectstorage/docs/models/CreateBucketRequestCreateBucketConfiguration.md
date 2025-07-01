# CreateBucketRequestCreateBucketConfiguration

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**LocationConstraint** | Pointer to **string** | Specifies the Region where the bucket will be created. Please refer to the &lt;a href&#x3D;\&quot;#section/Endpoints\&quot;&gt;list of available regions&lt;/a&gt;.  | [optional] |

## Methods

### NewCreateBucketRequestCreateBucketConfiguration

`func NewCreateBucketRequestCreateBucketConfiguration() *CreateBucketRequestCreateBucketConfiguration`

NewCreateBucketRequestCreateBucketConfiguration instantiates a new CreateBucketRequestCreateBucketConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBucketRequestCreateBucketConfigurationWithDefaults

`func NewCreateBucketRequestCreateBucketConfigurationWithDefaults() *CreateBucketRequestCreateBucketConfiguration`

NewCreateBucketRequestCreateBucketConfigurationWithDefaults instantiates a new CreateBucketRequestCreateBucketConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationConstraint

`func (o *CreateBucketRequestCreateBucketConfiguration) GetLocationConstraint() string`

GetLocationConstraint returns the LocationConstraint field if non-nil, zero value otherwise.

### GetLocationConstraintOk

`func (o *CreateBucketRequestCreateBucketConfiguration) GetLocationConstraintOk() (*string, bool)`

GetLocationConstraintOk returns a tuple with the LocationConstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationConstraint

`func (o *CreateBucketRequestCreateBucketConfiguration) SetLocationConstraint(v string)`

SetLocationConstraint sets LocationConstraint field to given value.

### HasLocationConstraint

`func (o *CreateBucketRequestCreateBucketConfiguration) HasLocationConstraint() bool`

HasLocationConstraint returns a boolean if a field has been set.


