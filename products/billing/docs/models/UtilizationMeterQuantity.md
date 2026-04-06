# UtilizationMeterQuantity

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Quantity** | Pointer to **float32** |  | [optional] |
|**Unit** | Pointer to **string** |  | [optional] |

## Methods

### NewUtilizationMeterQuantity

`func NewUtilizationMeterQuantity() *UtilizationMeterQuantity`

NewUtilizationMeterQuantity instantiates a new UtilizationMeterQuantity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtilizationMeterQuantityWithDefaults

`func NewUtilizationMeterQuantityWithDefaults() *UtilizationMeterQuantity`

NewUtilizationMeterQuantityWithDefaults instantiates a new UtilizationMeterQuantity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantity

`func (o *UtilizationMeterQuantity) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UtilizationMeterQuantity) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UtilizationMeterQuantity) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UtilizationMeterQuantity) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetUnit

`func (o *UtilizationMeterQuantity) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *UtilizationMeterQuantity) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *UtilizationMeterQuantity) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *UtilizationMeterQuantity) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


