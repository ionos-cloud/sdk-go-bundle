# Product

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**MeterId** | Pointer to **string** | Product code (just active will be provided) | [optional] |
|**MeterDesc** | Pointer to **string** | Product description (human readable) | [optional] |
|**Deprecated** | Pointer to **bool** | The flag specifies whether the product is marked as deprecated | [optional] |
|**Unit** | Pointer to **string** | The unit applied for the product during computation | [optional] |
|**Type** | Pointer to **string** | type of the billing (consumption-based or commitment-based) | [optional] |
|**Term** | Pointer to **NullableInt32** | length of the commitment period (only provided for commitment items) in months | [optional] |
|**UnitCost** | Pointer to [**NullableProductUnitCost**](ProductUnitCost.md) |  | [optional] |

## Methods

### NewProduct

`func NewProduct() *Product`

NewProduct instantiates a new Product object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductWithDefaults

`func NewProductWithDefaults() *Product`

NewProductWithDefaults instantiates a new Product object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeterId

`func (o *Product) GetMeterId() string`

GetMeterId returns the MeterId field if non-nil, zero value otherwise.

### GetMeterIdOk

`func (o *Product) GetMeterIdOk() (*string, bool)`

GetMeterIdOk returns a tuple with the MeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterId

`func (o *Product) SetMeterId(v string)`

SetMeterId sets MeterId field to given value.

### HasMeterId

`func (o *Product) HasMeterId() bool`

HasMeterId returns a boolean if a field has been set.

### GetMeterDesc

`func (o *Product) GetMeterDesc() string`

GetMeterDesc returns the MeterDesc field if non-nil, zero value otherwise.

### GetMeterDescOk

`func (o *Product) GetMeterDescOk() (*string, bool)`

GetMeterDescOk returns a tuple with the MeterDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterDesc

`func (o *Product) SetMeterDesc(v string)`

SetMeterDesc sets MeterDesc field to given value.

### HasMeterDesc

`func (o *Product) HasMeterDesc() bool`

HasMeterDesc returns a boolean if a field has been set.

### GetDeprecated

`func (o *Product) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *Product) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *Product) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *Product) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetUnit

`func (o *Product) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Product) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Product) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *Product) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetType

`func (o *Product) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Product) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Product) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Product) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTerm

`func (o *Product) GetTerm() int32`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *Product) GetTermOk() (*int32, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *Product) SetTerm(v int32)`

SetTerm sets Term field to given value.

### HasTerm

`func (o *Product) HasTerm() bool`

HasTerm returns a boolean if a field has been set.

### SetTermNil

`func (o *Product) SetTermNil(b bool)`

 SetTermNil sets the value for Term to be an explicit nil

### UnsetTerm
`func (o *Product) UnsetTerm()`

UnsetTerm ensures that no value is present for Term, not even an explicit nil
### GetUnitCost

`func (o *Product) GetUnitCost() ProductUnitCost`

GetUnitCost returns the UnitCost field if non-nil, zero value otherwise.

### GetUnitCostOk

`func (o *Product) GetUnitCostOk() (*ProductUnitCost, bool)`

GetUnitCostOk returns a tuple with the UnitCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitCost

`func (o *Product) SetUnitCost(v ProductUnitCost)`

SetUnitCost sets UnitCost field to given value.

### HasUnitCost

`func (o *Product) HasUnitCost() bool`

HasUnitCost returns a boolean if a field has been set.

### SetUnitCostNil

`func (o *Product) SetUnitCostNil(b bool)`

 SetUnitCostNil sets the value for UnitCost to be an explicit nil

### UnsetUnitCost
`func (o *Product) UnsetUnitCost()`

UnsetUnitCost ensures that no value is present for UnitCost, not even an explicit nil

