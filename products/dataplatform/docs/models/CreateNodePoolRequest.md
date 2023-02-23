# CreateNodePoolRequest

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Properties** | [**CreateNodePoolProperties**](CreateNodePoolProperties.md) |  | |

## Methods

### NewCreateNodePoolRequest

`func NewCreateNodePoolRequest(properties CreateNodePoolProperties, ) *CreateNodePoolRequest`

NewCreateNodePoolRequest instantiates a new CreateNodePoolRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNodePoolRequestWithDefaults

`func NewCreateNodePoolRequestWithDefaults() *CreateNodePoolRequest`

NewCreateNodePoolRequestWithDefaults instantiates a new CreateNodePoolRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *CreateNodePoolRequest) GetProperties() CreateNodePoolProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CreateNodePoolRequest) GetPropertiesOk() (*CreateNodePoolProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CreateNodePoolRequest) SetProperties(v CreateNodePoolProperties)`

SetProperties sets Properties field to given value.



