# InvoiceMetadata

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**InvoiceId** | Pointer to **string** |  | [optional] |
|**ContractId** | Pointer to **string** |  | [optional] |
|**CustomerId** | Pointer to **string** |  | [optional] |
|**CreatedDate** | Pointer to **string** |  | [optional] |
|**StartDate** | Pointer to **string** |  | [optional] |
|**EndDate** | Pointer to **string** |  | [optional] |
|**PostingPeriod** | Pointer to **string** |  | [optional] |
|**FinallyPosted** | Pointer to **bool** |  | [optional] |
|**Reference** | Pointer to **NullableString** |  | [optional] |
|**ResellerRef** | Pointer to **NullableString** | Reference text or number for the subcontract set by the reseller | [optional] |

## Methods

### NewInvoiceMetadata

`func NewInvoiceMetadata() *InvoiceMetadata`

NewInvoiceMetadata instantiates a new InvoiceMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceMetadataWithDefaults

`func NewInvoiceMetadataWithDefaults() *InvoiceMetadata`

NewInvoiceMetadataWithDefaults instantiates a new InvoiceMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceId

`func (o *InvoiceMetadata) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *InvoiceMetadata) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *InvoiceMetadata) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *InvoiceMetadata) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetContractId

`func (o *InvoiceMetadata) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *InvoiceMetadata) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *InvoiceMetadata) SetContractId(v string)`

SetContractId sets ContractId field to given value.

### HasContractId

`func (o *InvoiceMetadata) HasContractId() bool`

HasContractId returns a boolean if a field has been set.

### GetCustomerId

`func (o *InvoiceMetadata) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *InvoiceMetadata) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *InvoiceMetadata) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *InvoiceMetadata) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetCreatedDate

`func (o *InvoiceMetadata) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *InvoiceMetadata) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *InvoiceMetadata) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *InvoiceMetadata) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetStartDate

`func (o *InvoiceMetadata) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InvoiceMetadata) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InvoiceMetadata) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *InvoiceMetadata) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *InvoiceMetadata) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *InvoiceMetadata) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *InvoiceMetadata) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *InvoiceMetadata) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetPostingPeriod

`func (o *InvoiceMetadata) GetPostingPeriod() string`

GetPostingPeriod returns the PostingPeriod field if non-nil, zero value otherwise.

### GetPostingPeriodOk

`func (o *InvoiceMetadata) GetPostingPeriodOk() (*string, bool)`

GetPostingPeriodOk returns a tuple with the PostingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostingPeriod

`func (o *InvoiceMetadata) SetPostingPeriod(v string)`

SetPostingPeriod sets PostingPeriod field to given value.

### HasPostingPeriod

`func (o *InvoiceMetadata) HasPostingPeriod() bool`

HasPostingPeriod returns a boolean if a field has been set.

### GetFinallyPosted

`func (o *InvoiceMetadata) GetFinallyPosted() bool`

GetFinallyPosted returns the FinallyPosted field if non-nil, zero value otherwise.

### GetFinallyPostedOk

`func (o *InvoiceMetadata) GetFinallyPostedOk() (*bool, bool)`

GetFinallyPostedOk returns a tuple with the FinallyPosted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinallyPosted

`func (o *InvoiceMetadata) SetFinallyPosted(v bool)`

SetFinallyPosted sets FinallyPosted field to given value.

### HasFinallyPosted

`func (o *InvoiceMetadata) HasFinallyPosted() bool`

HasFinallyPosted returns a boolean if a field has been set.

### GetReference

`func (o *InvoiceMetadata) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *InvoiceMetadata) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *InvoiceMetadata) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *InvoiceMetadata) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *InvoiceMetadata) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *InvoiceMetadata) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetResellerRef

`func (o *InvoiceMetadata) GetResellerRef() string`

GetResellerRef returns the ResellerRef field if non-nil, zero value otherwise.

### GetResellerRefOk

`func (o *InvoiceMetadata) GetResellerRefOk() (*string, bool)`

GetResellerRefOk returns a tuple with the ResellerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellerRef

`func (o *InvoiceMetadata) SetResellerRef(v string)`

SetResellerRef sets ResellerRef field to given value.

### HasResellerRef

`func (o *InvoiceMetadata) HasResellerRef() bool`

HasResellerRef returns a boolean if a field has been set.

### SetResellerRefNil

`func (o *InvoiceMetadata) SetResellerRefNil(b bool)`

 SetResellerRefNil sets the value for ResellerRef to be an explicit nil

### UnsetResellerRef
`func (o *InvoiceMetadata) UnsetResellerRef()`

UnsetResellerRef ensures that no value is present for ResellerRef, not even an explicit nil

