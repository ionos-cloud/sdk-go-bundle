# ProductsGet200Response

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] |
|**Liability** | Pointer to **string** |  | [optional] |
|**Products** | Pointer to [**[]Product**](Product.md) |  | [optional] |

## Methods

### NewProductsGet200Response

`func NewProductsGet200Response() *ProductsGet200Response`

NewProductsGet200Response instantiates a new ProductsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductsGet200ResponseWithDefaults

`func NewProductsGet200ResponseWithDefaults() *ProductsGet200Response`

NewProductsGet200ResponseWithDefaults instantiates a new ProductsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *ProductsGet200Response) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ProductsGet200Response) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ProductsGet200Response) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ProductsGet200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetLiability

`func (o *ProductsGet200Response) GetLiability() string`

GetLiability returns the Liability field if non-nil, zero value otherwise.

### GetLiabilityOk

`func (o *ProductsGet200Response) GetLiabilityOk() (*string, bool)`

GetLiabilityOk returns a tuple with the Liability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiability

`func (o *ProductsGet200Response) SetLiability(v string)`

SetLiability sets Liability field to given value.

### HasLiability

`func (o *ProductsGet200Response) HasLiability() bool`

HasLiability returns a boolean if a field has been set.

### GetProducts

`func (o *ProductsGet200Response) GetProducts() []Product`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *ProductsGet200Response) GetProductsOk() (*[]Product, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *ProductsGet200Response) SetProducts(v []Product)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *ProductsGet200Response) HasProducts() bool`

HasProducts returns a boolean if a field has been set.


