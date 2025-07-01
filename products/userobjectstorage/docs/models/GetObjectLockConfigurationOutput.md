# GetObjectLockConfigurationOutput

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**ObjectLockConfiguration** | Pointer to [**GetObjectLockConfigurationOutputObjectLockConfiguration**](GetObjectLockConfigurationOutputObjectLockConfiguration.md) |  | [optional] |

## Methods

### NewGetObjectLockConfigurationOutput

`func NewGetObjectLockConfigurationOutput() *GetObjectLockConfigurationOutput`

NewGetObjectLockConfigurationOutput instantiates a new GetObjectLockConfigurationOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetObjectLockConfigurationOutputWithDefaults

`func NewGetObjectLockConfigurationOutputWithDefaults() *GetObjectLockConfigurationOutput`

NewGetObjectLockConfigurationOutputWithDefaults instantiates a new GetObjectLockConfigurationOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectLockConfiguration

`func (o *GetObjectLockConfigurationOutput) GetObjectLockConfiguration() GetObjectLockConfigurationOutputObjectLockConfiguration`

GetObjectLockConfiguration returns the ObjectLockConfiguration field if non-nil, zero value otherwise.

### GetObjectLockConfigurationOk

`func (o *GetObjectLockConfigurationOutput) GetObjectLockConfigurationOk() (*GetObjectLockConfigurationOutputObjectLockConfiguration, bool)`

GetObjectLockConfigurationOk returns a tuple with the ObjectLockConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectLockConfiguration

`func (o *GetObjectLockConfigurationOutput) SetObjectLockConfiguration(v GetObjectLockConfigurationOutputObjectLockConfiguration)`

SetObjectLockConfiguration sets ObjectLockConfiguration field to given value.

### HasObjectLockConfiguration

`func (o *GetObjectLockConfigurationOutput) HasObjectLockConfiguration() bool`

HasObjectLockConfiguration returns a boolean if a field has been set.


