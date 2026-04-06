# InvoicesGet200Response

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] |
|**Invoices** | Pointer to [**[]InvoicesGet200ResponseInvoices**](InvoicesGet200ResponseInvoices.md) |  | [optional] |

## Methods

### NewInvoicesGet200Response

`func NewInvoicesGet200Response() *InvoicesGet200Response`

NewInvoicesGet200Response instantiates a new InvoicesGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoicesGet200ResponseWithDefaults

`func NewInvoicesGet200ResponseWithDefaults() *InvoicesGet200Response`

NewInvoicesGet200ResponseWithDefaults instantiates a new InvoicesGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *InvoicesGet200Response) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InvoicesGet200Response) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InvoicesGet200Response) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InvoicesGet200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetInvoices

`func (o *InvoicesGet200Response) GetInvoices() []InvoicesGet200ResponseInvoices`

GetInvoices returns the Invoices field if non-nil, zero value otherwise.

### GetInvoicesOk

`func (o *InvoicesGet200Response) GetInvoicesOk() (*[]InvoicesGet200ResponseInvoices, bool)`

GetInvoicesOk returns a tuple with the Invoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoices

`func (o *InvoicesGet200Response) SetInvoices(v []InvoicesGet200ResponseInvoices)`

SetInvoices sets Invoices field to given value.

### HasInvoices

`func (o *InvoicesGet200Response) HasInvoices() bool`

HasInvoices returns a boolean if a field has been set.


