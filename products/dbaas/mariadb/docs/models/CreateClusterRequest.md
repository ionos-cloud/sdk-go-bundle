# CreateClusterRequest

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Properties** | Pointer to [**CreateClusterProperties**](CreateClusterProperties.md) |  | [optional] |

## Methods

### NewCreateClusterRequest

`func NewCreateClusterRequest() *CreateClusterRequest`

NewCreateClusterRequest instantiates a new CreateClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClusterRequestWithDefaults

`func NewCreateClusterRequestWithDefaults() *CreateClusterRequest`

NewCreateClusterRequestWithDefaults instantiates a new CreateClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *CreateClusterRequest) GetProperties() CreateClusterProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CreateClusterRequest) GetPropertiesOk() (*CreateClusterProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CreateClusterRequest) SetProperties(v CreateClusterProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CreateClusterRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


