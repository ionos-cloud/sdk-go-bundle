# PutRegistryInput

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Properties** | [**PostRegistryProperties**](PostRegistryProperties.md) |  | |

## Methods

### NewPutRegistryInput

`func NewPutRegistryInput(properties PostRegistryProperties, ) *PutRegistryInput`

NewPutRegistryInput instantiates a new PutRegistryInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutRegistryInputWithDefaults

`func NewPutRegistryInputWithDefaults() *PutRegistryInput`

NewPutRegistryInputWithDefaults instantiates a new PutRegistryInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *PutRegistryInput) GetProperties() PostRegistryProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PutRegistryInput) GetPropertiesOk() (*PostRegistryProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PutRegistryInput) SetProperties(v PostRegistryProperties)`

SetProperties sets Properties field to given value.



