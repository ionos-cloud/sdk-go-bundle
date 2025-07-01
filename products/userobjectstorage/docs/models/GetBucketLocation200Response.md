# GetBucketLocation200Response

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**LocationConstraint** | Pointer to **string** | Specifies the Region where the bucket resides. | [optional] |

## Methods

### NewGetBucketLocation200Response

`func NewGetBucketLocation200Response() *GetBucketLocation200Response`

NewGetBucketLocation200Response instantiates a new GetBucketLocation200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBucketLocation200ResponseWithDefaults

`func NewGetBucketLocation200ResponseWithDefaults() *GetBucketLocation200Response`

NewGetBucketLocation200ResponseWithDefaults instantiates a new GetBucketLocation200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationConstraint

`func (o *GetBucketLocation200Response) GetLocationConstraint() string`

GetLocationConstraint returns the LocationConstraint field if non-nil, zero value otherwise.

### GetLocationConstraintOk

`func (o *GetBucketLocation200Response) GetLocationConstraintOk() (*string, bool)`

GetLocationConstraintOk returns a tuple with the LocationConstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationConstraint

`func (o *GetBucketLocation200Response) SetLocationConstraint(v string)`

SetLocationConstraint sets LocationConstraint field to given value.

### HasLocationConstraint

`func (o *GetBucketLocation200Response) HasLocationConstraint() bool`

HasLocationConstraint returns a boolean if a field has been set.


