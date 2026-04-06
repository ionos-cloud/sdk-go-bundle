# Metadata

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**ContractId** | Pointer to **string** |  | [optional] |
|**CustomerId** | Pointer to **string** |  | [optional] |
|**Reference** | Pointer to **NullableString** |  | [optional] |

## Methods

### NewMetadata

`func NewMetadata() *Metadata`

NewMetadata instantiates a new Metadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataWithDefaults

`func NewMetadataWithDefaults() *Metadata`

NewMetadataWithDefaults instantiates a new Metadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractId

`func (o *Metadata) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *Metadata) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *Metadata) SetContractId(v string)`

SetContractId sets ContractId field to given value.

### HasContractId

`func (o *Metadata) HasContractId() bool`

HasContractId returns a boolean if a field has been set.

### GetCustomerId

`func (o *Metadata) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *Metadata) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *Metadata) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *Metadata) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetReference

`func (o *Metadata) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *Metadata) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *Metadata) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *Metadata) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *Metadata) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *Metadata) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil

