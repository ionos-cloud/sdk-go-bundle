# UtilizationGet200Response

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**StartDate** | Pointer to **string** |  | [optional] |
|**EndDate** | Pointer to **string** |  | [optional] |
|**Datacenters** | Pointer to [**[]UtilizationDataCenter**](UtilizationDataCenter.md) |  | [optional] |
|**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] |

## Methods

### NewUtilizationGet200Response

`func NewUtilizationGet200Response() *UtilizationGet200Response`

NewUtilizationGet200Response instantiates a new UtilizationGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtilizationGet200ResponseWithDefaults

`func NewUtilizationGet200ResponseWithDefaults() *UtilizationGet200Response`

NewUtilizationGet200ResponseWithDefaults instantiates a new UtilizationGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *UtilizationGet200Response) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UtilizationGet200Response) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UtilizationGet200Response) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *UtilizationGet200Response) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *UtilizationGet200Response) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *UtilizationGet200Response) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *UtilizationGet200Response) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *UtilizationGet200Response) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetDatacenters

`func (o *UtilizationGet200Response) GetDatacenters() []UtilizationDataCenter`

GetDatacenters returns the Datacenters field if non-nil, zero value otherwise.

### GetDatacentersOk

`func (o *UtilizationGet200Response) GetDatacentersOk() (*[]UtilizationDataCenter, bool)`

GetDatacentersOk returns a tuple with the Datacenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenters

`func (o *UtilizationGet200Response) SetDatacenters(v []UtilizationDataCenter)`

SetDatacenters sets Datacenters field to given value.

### HasDatacenters

`func (o *UtilizationGet200Response) HasDatacenters() bool`

HasDatacenters returns a boolean if a field has been set.

### GetMetadata

`func (o *UtilizationGet200Response) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UtilizationGet200Response) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UtilizationGet200Response) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UtilizationGet200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


