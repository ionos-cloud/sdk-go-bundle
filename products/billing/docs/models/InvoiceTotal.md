# InvoiceTotal

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Quantity** | Pointer to **float32** |  | [optional] |
|**Unit** | Pointer to **string** |  | [optional] |

## Methods

### NewInvoiceTotal

`func NewInvoiceTotal() *InvoiceTotal`

NewInvoiceTotal instantiates a new InvoiceTotal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceTotalWithDefaults

`func NewInvoiceTotalWithDefaults() *InvoiceTotal`

NewInvoiceTotalWithDefaults instantiates a new InvoiceTotal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantity

`func (o *InvoiceTotal) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InvoiceTotal) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InvoiceTotal) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *InvoiceTotal) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetUnit

`func (o *InvoiceTotal) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *InvoiceTotal) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *InvoiceTotal) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *InvoiceTotal) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


