# InvoicesGet200ResponseInvoices

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** |  | [optional] |
|**Date** | Pointer to **string** |  | [optional] |
|**Amount** | Pointer to **NullableFloat32** |  | [optional] |
|**Unit** | Pointer to **NullableString** |  | [optional] |

## Methods

### NewInvoicesGet200ResponseInvoices

`func NewInvoicesGet200ResponseInvoices() *InvoicesGet200ResponseInvoices`

NewInvoicesGet200ResponseInvoices instantiates a new InvoicesGet200ResponseInvoices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoicesGet200ResponseInvoicesWithDefaults

`func NewInvoicesGet200ResponseInvoicesWithDefaults() *InvoicesGet200ResponseInvoices`

NewInvoicesGet200ResponseInvoicesWithDefaults instantiates a new InvoicesGet200ResponseInvoices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoicesGet200ResponseInvoices) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoicesGet200ResponseInvoices) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoicesGet200ResponseInvoices) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InvoicesGet200ResponseInvoices) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDate

`func (o *InvoicesGet200ResponseInvoices) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *InvoicesGet200ResponseInvoices) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *InvoicesGet200ResponseInvoices) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *InvoicesGet200ResponseInvoices) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetAmount

`func (o *InvoicesGet200ResponseInvoices) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InvoicesGet200ResponseInvoices) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InvoicesGet200ResponseInvoices) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *InvoicesGet200ResponseInvoices) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *InvoicesGet200ResponseInvoices) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *InvoicesGet200ResponseInvoices) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetUnit

`func (o *InvoicesGet200ResponseInvoices) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *InvoicesGet200ResponseInvoices) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *InvoicesGet200ResponseInvoices) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *InvoicesGet200ResponseInvoices) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnitNil

`func (o *InvoicesGet200ResponseInvoices) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *InvoicesGet200ResponseInvoices) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil

