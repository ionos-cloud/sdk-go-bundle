# InvoiceMeter

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**MeterId** | Pointer to **string** |  | [optional] |
|**MeterDesc** | Pointer to **string** |  | [optional] |
|**ProductGroup** | Pointer to **NullableString** | Product group of the price item | [optional] |
|**Quantity** | Pointer to [**InvoiceMeterQuantity**](InvoiceMeterQuantity.md) |  | [optional] |
|**Rate** | Pointer to [**InvoiceTotal**](InvoiceTotal.md) |  | [optional] |
|**Amount** | Pointer to [**InvoiceTotal**](InvoiceTotal.md) |  | [optional] |

## Methods

### NewInvoiceMeter

`func NewInvoiceMeter() *InvoiceMeter`

NewInvoiceMeter instantiates a new InvoiceMeter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceMeterWithDefaults

`func NewInvoiceMeterWithDefaults() *InvoiceMeter`

NewInvoiceMeterWithDefaults instantiates a new InvoiceMeter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeterId

`func (o *InvoiceMeter) GetMeterId() string`

GetMeterId returns the MeterId field if non-nil, zero value otherwise.

### GetMeterIdOk

`func (o *InvoiceMeter) GetMeterIdOk() (*string, bool)`

GetMeterIdOk returns a tuple with the MeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterId

`func (o *InvoiceMeter) SetMeterId(v string)`

SetMeterId sets MeterId field to given value.

### HasMeterId

`func (o *InvoiceMeter) HasMeterId() bool`

HasMeterId returns a boolean if a field has been set.

### GetMeterDesc

`func (o *InvoiceMeter) GetMeterDesc() string`

GetMeterDesc returns the MeterDesc field if non-nil, zero value otherwise.

### GetMeterDescOk

`func (o *InvoiceMeter) GetMeterDescOk() (*string, bool)`

GetMeterDescOk returns a tuple with the MeterDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterDesc

`func (o *InvoiceMeter) SetMeterDesc(v string)`

SetMeterDesc sets MeterDesc field to given value.

### HasMeterDesc

`func (o *InvoiceMeter) HasMeterDesc() bool`

HasMeterDesc returns a boolean if a field has been set.

### GetProductGroup

`func (o *InvoiceMeter) GetProductGroup() string`

GetProductGroup returns the ProductGroup field if non-nil, zero value otherwise.

### GetProductGroupOk

`func (o *InvoiceMeter) GetProductGroupOk() (*string, bool)`

GetProductGroupOk returns a tuple with the ProductGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductGroup

`func (o *InvoiceMeter) SetProductGroup(v string)`

SetProductGroup sets ProductGroup field to given value.

### HasProductGroup

`func (o *InvoiceMeter) HasProductGroup() bool`

HasProductGroup returns a boolean if a field has been set.

### SetProductGroupNil

`func (o *InvoiceMeter) SetProductGroupNil(b bool)`

 SetProductGroupNil sets the value for ProductGroup to be an explicit nil

### UnsetProductGroup
`func (o *InvoiceMeter) UnsetProductGroup()`

UnsetProductGroup ensures that no value is present for ProductGroup, not even an explicit nil
### GetQuantity

`func (o *InvoiceMeter) GetQuantity() InvoiceMeterQuantity`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InvoiceMeter) GetQuantityOk() (*InvoiceMeterQuantity, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InvoiceMeter) SetQuantity(v InvoiceMeterQuantity)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *InvoiceMeter) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetRate

`func (o *InvoiceMeter) GetRate() InvoiceTotal`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *InvoiceMeter) GetRateOk() (*InvoiceTotal, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *InvoiceMeter) SetRate(v InvoiceTotal)`

SetRate sets Rate field to given value.

### HasRate

`func (o *InvoiceMeter) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetAmount

`func (o *InvoiceMeter) GetAmount() InvoiceTotal`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InvoiceMeter) GetAmountOk() (*InvoiceTotal, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InvoiceMeter) SetAmount(v InvoiceTotal)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *InvoiceMeter) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


