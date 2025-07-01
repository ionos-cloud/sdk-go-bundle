# PutObjectLockConfigurationRequest

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**ObjectLockConfiguration** | Pointer to [**PutObjectLockConfigurationRequestObjectLockConfiguration**](PutObjectLockConfigurationRequestObjectLockConfiguration.md) |  | [optional] |

## Methods

### NewPutObjectLockConfigurationRequest

`func NewPutObjectLockConfigurationRequest() *PutObjectLockConfigurationRequest`

NewPutObjectLockConfigurationRequest instantiates a new PutObjectLockConfigurationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutObjectLockConfigurationRequestWithDefaults

`func NewPutObjectLockConfigurationRequestWithDefaults() *PutObjectLockConfigurationRequest`

NewPutObjectLockConfigurationRequestWithDefaults instantiates a new PutObjectLockConfigurationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectLockConfiguration

`func (o *PutObjectLockConfigurationRequest) GetObjectLockConfiguration() PutObjectLockConfigurationRequestObjectLockConfiguration`

GetObjectLockConfiguration returns the ObjectLockConfiguration field if non-nil, zero value otherwise.

### GetObjectLockConfigurationOk

`func (o *PutObjectLockConfigurationRequest) GetObjectLockConfigurationOk() (*PutObjectLockConfigurationRequestObjectLockConfiguration, bool)`

GetObjectLockConfigurationOk returns a tuple with the ObjectLockConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectLockConfiguration

`func (o *PutObjectLockConfigurationRequest) SetObjectLockConfiguration(v PutObjectLockConfigurationRequestObjectLockConfiguration)`

SetObjectLockConfiguration sets ObjectLockConfiguration field to given value.

### HasObjectLockConfiguration

`func (o *PutObjectLockConfigurationRequest) HasObjectLockConfiguration() bool`

HasObjectLockConfiguration returns a boolean if a field has been set.


