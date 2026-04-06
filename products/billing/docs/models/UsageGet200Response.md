# UsageGet200Response

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**StartDate** | Pointer to **string** |  | [optional] |
|**EndDate** | Pointer to **string** |  | [optional] |
|**Datacenters** | Pointer to [**[]UsageDataCenter**](UsageDataCenter.md) |  | [optional] |
|**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] |

## Methods

### NewUsageGet200Response

`func NewUsageGet200Response() *UsageGet200Response`

NewUsageGet200Response instantiates a new UsageGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageGet200ResponseWithDefaults

`func NewUsageGet200ResponseWithDefaults() *UsageGet200Response`

NewUsageGet200ResponseWithDefaults instantiates a new UsageGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *UsageGet200Response) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UsageGet200Response) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UsageGet200Response) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *UsageGet200Response) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *UsageGet200Response) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *UsageGet200Response) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *UsageGet200Response) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *UsageGet200Response) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetDatacenters

`func (o *UsageGet200Response) GetDatacenters() []UsageDataCenter`

GetDatacenters returns the Datacenters field if non-nil, zero value otherwise.

### GetDatacentersOk

`func (o *UsageGet200Response) GetDatacentersOk() (*[]UsageDataCenter, bool)`

GetDatacentersOk returns a tuple with the Datacenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenters

`func (o *UsageGet200Response) SetDatacenters(v []UsageDataCenter)`

SetDatacenters sets Datacenters field to given value.

### HasDatacenters

`func (o *UsageGet200Response) HasDatacenters() bool`

HasDatacenters returns a boolean if a field has been set.

### GetMetadata

`func (o *UsageGet200Response) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UsageGet200Response) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UsageGet200Response) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UsageGet200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


