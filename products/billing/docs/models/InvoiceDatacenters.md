# InvoiceDatacenters

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** |  | [optional] |
|**Name** | Pointer to **string** |  | [optional] |
|**Location** | Pointer to **string** |  | [optional] |
|**ProductGroup** | Pointer to **NullableString** | Product group of the data center | [optional] |
|**Meters** | Pointer to [**[]InvoiceMeter**](InvoiceMeter.md) |  | [optional] |
|**Rebate** | Pointer to [**InvoiceDatacentersRebate**](InvoiceDatacentersRebate.md) |  | [optional] |

## Methods

### NewInvoiceDatacenters

`func NewInvoiceDatacenters() *InvoiceDatacenters`

NewInvoiceDatacenters instantiates a new InvoiceDatacenters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceDatacentersWithDefaults

`func NewInvoiceDatacentersWithDefaults() *InvoiceDatacenters`

NewInvoiceDatacentersWithDefaults instantiates a new InvoiceDatacenters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceDatacenters) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceDatacenters) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceDatacenters) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InvoiceDatacenters) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InvoiceDatacenters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InvoiceDatacenters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InvoiceDatacenters) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InvoiceDatacenters) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocation

`func (o *InvoiceDatacenters) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *InvoiceDatacenters) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *InvoiceDatacenters) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *InvoiceDatacenters) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetProductGroup

`func (o *InvoiceDatacenters) GetProductGroup() string`

GetProductGroup returns the ProductGroup field if non-nil, zero value otherwise.

### GetProductGroupOk

`func (o *InvoiceDatacenters) GetProductGroupOk() (*string, bool)`

GetProductGroupOk returns a tuple with the ProductGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductGroup

`func (o *InvoiceDatacenters) SetProductGroup(v string)`

SetProductGroup sets ProductGroup field to given value.

### HasProductGroup

`func (o *InvoiceDatacenters) HasProductGroup() bool`

HasProductGroup returns a boolean if a field has been set.

### SetProductGroupNil

`func (o *InvoiceDatacenters) SetProductGroupNil(b bool)`

 SetProductGroupNil sets the value for ProductGroup to be an explicit nil

### UnsetProductGroup
`func (o *InvoiceDatacenters) UnsetProductGroup()`

UnsetProductGroup ensures that no value is present for ProductGroup, not even an explicit nil
### GetMeters

`func (o *InvoiceDatacenters) GetMeters() []InvoiceMeter`

GetMeters returns the Meters field if non-nil, zero value otherwise.

### GetMetersOk

`func (o *InvoiceDatacenters) GetMetersOk() (*[]InvoiceMeter, bool)`

GetMetersOk returns a tuple with the Meters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeters

`func (o *InvoiceDatacenters) SetMeters(v []InvoiceMeter)`

SetMeters sets Meters field to given value.

### HasMeters

`func (o *InvoiceDatacenters) HasMeters() bool`

HasMeters returns a boolean if a field has been set.

### GetRebate

`func (o *InvoiceDatacenters) GetRebate() InvoiceDatacentersRebate`

GetRebate returns the Rebate field if non-nil, zero value otherwise.

### GetRebateOk

`func (o *InvoiceDatacenters) GetRebateOk() (*InvoiceDatacentersRebate, bool)`

GetRebateOk returns a tuple with the Rebate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebate

`func (o *InvoiceDatacenters) SetRebate(v InvoiceDatacentersRebate)`

SetRebate sets Rebate field to given value.

### HasRebate

`func (o *InvoiceDatacenters) HasRebate() bool`

HasRebate returns a boolean if a field has been set.


