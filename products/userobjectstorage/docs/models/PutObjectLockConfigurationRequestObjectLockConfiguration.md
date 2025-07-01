# PutObjectLockConfigurationRequestObjectLockConfiguration

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**ObjectLockEnabled** | Pointer to **string** |  | [optional] |
|**Rule** | Pointer to [**PutObjectLockConfigurationRequestObjectLockConfigurationRule**](PutObjectLockConfigurationRequestObjectLockConfigurationRule.md) |  | [optional] |

## Methods

### NewPutObjectLockConfigurationRequestObjectLockConfiguration

`func NewPutObjectLockConfigurationRequestObjectLockConfiguration() *PutObjectLockConfigurationRequestObjectLockConfiguration`

NewPutObjectLockConfigurationRequestObjectLockConfiguration instantiates a new PutObjectLockConfigurationRequestObjectLockConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutObjectLockConfigurationRequestObjectLockConfigurationWithDefaults

`func NewPutObjectLockConfigurationRequestObjectLockConfigurationWithDefaults() *PutObjectLockConfigurationRequestObjectLockConfiguration`

NewPutObjectLockConfigurationRequestObjectLockConfigurationWithDefaults instantiates a new PutObjectLockConfigurationRequestObjectLockConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectLockEnabled

`func (o *PutObjectLockConfigurationRequestObjectLockConfiguration) GetObjectLockEnabled() string`

GetObjectLockEnabled returns the ObjectLockEnabled field if non-nil, zero value otherwise.

### GetObjectLockEnabledOk

`func (o *PutObjectLockConfigurationRequestObjectLockConfiguration) GetObjectLockEnabledOk() (*string, bool)`

GetObjectLockEnabledOk returns a tuple with the ObjectLockEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectLockEnabled

`func (o *PutObjectLockConfigurationRequestObjectLockConfiguration) SetObjectLockEnabled(v string)`

SetObjectLockEnabled sets ObjectLockEnabled field to given value.

### HasObjectLockEnabled

`func (o *PutObjectLockConfigurationRequestObjectLockConfiguration) HasObjectLockEnabled() bool`

HasObjectLockEnabled returns a boolean if a field has been set.

### GetRule

`func (o *PutObjectLockConfigurationRequestObjectLockConfiguration) GetRule() PutObjectLockConfigurationRequestObjectLockConfigurationRule`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *PutObjectLockConfigurationRequestObjectLockConfiguration) GetRuleOk() (*PutObjectLockConfigurationRequestObjectLockConfigurationRule, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *PutObjectLockConfigurationRequestObjectLockConfiguration) SetRule(v PutObjectLockConfigurationRequestObjectLockConfigurationRule)`

SetRule sets Rule field to given value.

### HasRule

`func (o *PutObjectLockConfigurationRequestObjectLockConfiguration) HasRule() bool`

HasRule returns a boolean if a field has been set.


