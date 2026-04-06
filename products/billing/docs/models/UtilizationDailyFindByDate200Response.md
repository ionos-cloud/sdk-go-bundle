# UtilizationDailyFindByDate200Response

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**StartDate** | Pointer to **string** |  | [optional] |
|**EndDate** | Pointer to **string** |  | [optional] |
|**Datacenters** | Pointer to [**[]UtilizationDataCenter**](UtilizationDataCenter.md) |  | [optional] |
|**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] |

## Methods

### NewUtilizationDailyFindByDate200Response

`func NewUtilizationDailyFindByDate200Response() *UtilizationDailyFindByDate200Response`

NewUtilizationDailyFindByDate200Response instantiates a new UtilizationDailyFindByDate200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtilizationDailyFindByDate200ResponseWithDefaults

`func NewUtilizationDailyFindByDate200ResponseWithDefaults() *UtilizationDailyFindByDate200Response`

NewUtilizationDailyFindByDate200ResponseWithDefaults instantiates a new UtilizationDailyFindByDate200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *UtilizationDailyFindByDate200Response) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UtilizationDailyFindByDate200Response) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UtilizationDailyFindByDate200Response) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *UtilizationDailyFindByDate200Response) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *UtilizationDailyFindByDate200Response) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *UtilizationDailyFindByDate200Response) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *UtilizationDailyFindByDate200Response) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *UtilizationDailyFindByDate200Response) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetDatacenters

`func (o *UtilizationDailyFindByDate200Response) GetDatacenters() []UtilizationDataCenter`

GetDatacenters returns the Datacenters field if non-nil, zero value otherwise.

### GetDatacentersOk

`func (o *UtilizationDailyFindByDate200Response) GetDatacentersOk() (*[]UtilizationDataCenter, bool)`

GetDatacentersOk returns a tuple with the Datacenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenters

`func (o *UtilizationDailyFindByDate200Response) SetDatacenters(v []UtilizationDataCenter)`

SetDatacenters sets Datacenters field to given value.

### HasDatacenters

`func (o *UtilizationDailyFindByDate200Response) HasDatacenters() bool`

HasDatacenters returns a boolean if a field has been set.

### GetMetadata

`func (o *UtilizationDailyFindByDate200Response) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UtilizationDailyFindByDate200Response) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UtilizationDailyFindByDate200Response) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UtilizationDailyFindByDate200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


