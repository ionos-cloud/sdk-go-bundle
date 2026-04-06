# UsageMeter

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**MeterId** | Pointer to **string** | Usage meter product code | [optional] |
|**MeterDesc** | Pointer to **string** |  | [optional] |
|**Quantity** | Pointer to [**UsageMeterQuantity**](UsageMeterQuantity.md) |  | [optional] |

## Methods

### NewUsageMeter

`func NewUsageMeter() *UsageMeter`

NewUsageMeter instantiates a new UsageMeter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageMeterWithDefaults

`func NewUsageMeterWithDefaults() *UsageMeter`

NewUsageMeterWithDefaults instantiates a new UsageMeter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeterId

`func (o *UsageMeter) GetMeterId() string`

GetMeterId returns the MeterId field if non-nil, zero value otherwise.

### GetMeterIdOk

`func (o *UsageMeter) GetMeterIdOk() (*string, bool)`

GetMeterIdOk returns a tuple with the MeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterId

`func (o *UsageMeter) SetMeterId(v string)`

SetMeterId sets MeterId field to given value.

### HasMeterId

`func (o *UsageMeter) HasMeterId() bool`

HasMeterId returns a boolean if a field has been set.

### GetMeterDesc

`func (o *UsageMeter) GetMeterDesc() string`

GetMeterDesc returns the MeterDesc field if non-nil, zero value otherwise.

### GetMeterDescOk

`func (o *UsageMeter) GetMeterDescOk() (*string, bool)`

GetMeterDescOk returns a tuple with the MeterDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterDesc

`func (o *UsageMeter) SetMeterDesc(v string)`

SetMeterDesc sets MeterDesc field to given value.

### HasMeterDesc

`func (o *UsageMeter) HasMeterDesc() bool`

HasMeterDesc returns a boolean if a field has been set.

### GetQuantity

`func (o *UsageMeter) GetQuantity() UsageMeterQuantity`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UsageMeter) GetQuantityOk() (*UsageMeterQuantity, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UsageMeter) SetQuantity(v UsageMeterQuantity)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UsageMeter) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


