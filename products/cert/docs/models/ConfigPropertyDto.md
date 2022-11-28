# ConfigPropertyDto

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The property id. | [optional] [readonly] |
|**Name** | **string** | The property name. | |
|**Value** | **string** | The property value. | |
|**Description** | **string** | The property description. | |

## Methods

### NewConfigPropertyDto

`func NewConfigPropertyDto(name string, value string, description string, ) *ConfigPropertyDto`

NewConfigPropertyDto instantiates a new ConfigPropertyDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigPropertyDtoWithDefaults

`func NewConfigPropertyDtoWithDefaults() *ConfigPropertyDto`

NewConfigPropertyDtoWithDefaults instantiates a new ConfigPropertyDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConfigPropertyDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigPropertyDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigPropertyDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConfigPropertyDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ConfigPropertyDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigPropertyDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigPropertyDto) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *ConfigPropertyDto) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ConfigPropertyDto) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ConfigPropertyDto) SetValue(v string)`

SetValue sets Value field to given value.


### GetDescription

`func (o *ConfigPropertyDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfigPropertyDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfigPropertyDto) SetDescription(v string)`

SetDescription sets Description field to given value.



