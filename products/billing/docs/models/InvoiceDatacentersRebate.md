# InvoiceDatacentersRebate

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Amount** | Pointer to [**InvoiceTotal**](InvoiceTotal.md) |  | [optional] |

## Methods

### NewInvoiceDatacentersRebate

`func NewInvoiceDatacentersRebate() *InvoiceDatacentersRebate`

NewInvoiceDatacentersRebate instantiates a new InvoiceDatacentersRebate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceDatacentersRebateWithDefaults

`func NewInvoiceDatacentersRebateWithDefaults() *InvoiceDatacentersRebate`

NewInvoiceDatacentersRebateWithDefaults instantiates a new InvoiceDatacentersRebate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *InvoiceDatacentersRebate) GetAmount() InvoiceTotal`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InvoiceDatacentersRebate) GetAmountOk() (*InvoiceTotal, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InvoiceDatacentersRebate) SetAmount(v InvoiceTotal)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *InvoiceDatacentersRebate) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


