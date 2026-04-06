# Invoice

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Metadata** | Pointer to [**InvoiceMetadata**](InvoiceMetadata.md) |  | [optional] |
|**Datacenters** | Pointer to [**[]InvoiceDatacenters**](InvoiceDatacenters.md) |  | [optional] |
|**Total** | Pointer to [**InvoiceTotal**](InvoiceTotal.md) |  | [optional] |

## Methods

### NewInvoice

`func NewInvoice() *Invoice`

NewInvoice instantiates a new Invoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceWithDefaults

`func NewInvoiceWithDefaults() *Invoice`

NewInvoiceWithDefaults instantiates a new Invoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *Invoice) GetMetadata() InvoiceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Invoice) GetMetadataOk() (*InvoiceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Invoice) SetMetadata(v InvoiceMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Invoice) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDatacenters

`func (o *Invoice) GetDatacenters() []InvoiceDatacenters`

GetDatacenters returns the Datacenters field if non-nil, zero value otherwise.

### GetDatacentersOk

`func (o *Invoice) GetDatacentersOk() (*[]InvoiceDatacenters, bool)`

GetDatacentersOk returns a tuple with the Datacenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenters

`func (o *Invoice) SetDatacenters(v []InvoiceDatacenters)`

SetDatacenters sets Datacenters field to given value.

### HasDatacenters

`func (o *Invoice) HasDatacenters() bool`

HasDatacenters returns a boolean if a field has been set.

### GetTotal

`func (o *Invoice) GetTotal() InvoiceTotal`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Invoice) GetTotalOk() (*InvoiceTotal, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Invoice) SetTotal(v InvoiceTotal)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *Invoice) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


