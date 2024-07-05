# UsersPatchRequest

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Properties** | [**PatchUserProperties**](PatchUserProperties.md) |  | |

## Methods

### NewUsersPatchRequest

`func NewUsersPatchRequest(properties PatchUserProperties, ) *UsersPatchRequest`

NewUsersPatchRequest instantiates a new UsersPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersPatchRequestWithDefaults

`func NewUsersPatchRequestWithDefaults() *UsersPatchRequest`

NewUsersPatchRequestWithDefaults instantiates a new UsersPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *UsersPatchRequest) GetProperties() PatchUserProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *UsersPatchRequest) GetPropertiesOk() (*PatchUserProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *UsersPatchRequest) SetProperties(v PatchUserProperties)`

SetProperties sets Properties field to given value.



