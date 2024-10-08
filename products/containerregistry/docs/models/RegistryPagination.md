# RegistryPagination

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Limit** | **int32** | The maximum number of elements to return (used together with pagination.token for pagination) | |
|**Token** | **string** | An opaque token used to iterate the set of results (used together with limit for pagination) | |

## Methods

### NewRegistryPagination

`func NewRegistryPagination(limit int32, token string, ) *RegistryPagination`

NewRegistryPagination instantiates a new RegistryPagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryPaginationWithDefaults

`func NewRegistryPaginationWithDefaults() *RegistryPagination`

NewRegistryPaginationWithDefaults instantiates a new RegistryPagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *RegistryPagination) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *RegistryPagination) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *RegistryPagination) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetToken

`func (o *RegistryPagination) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RegistryPagination) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RegistryPagination) SetToken(v string)`

SetToken sets Token field to given value.



