# Evn

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Metadata** | Pointer to [**EvnMetadata**](EvnMetadata.md) |  | [optional] |
|**Datacenters** | Pointer to [**[]EvnDatacenters**](EvnDatacenters.md) |  | [optional] |
|**EvnCSV** | Pointer to **[]string** |  | [optional] |

## Methods

### NewEvn

`func NewEvn() *Evn`

NewEvn instantiates a new Evn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvnWithDefaults

`func NewEvnWithDefaults() *Evn`

NewEvnWithDefaults instantiates a new Evn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *Evn) GetMetadata() EvnMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Evn) GetMetadataOk() (*EvnMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Evn) SetMetadata(v EvnMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Evn) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDatacenters

`func (o *Evn) GetDatacenters() []EvnDatacenters`

GetDatacenters returns the Datacenters field if non-nil, zero value otherwise.

### GetDatacentersOk

`func (o *Evn) GetDatacentersOk() (*[]EvnDatacenters, bool)`

GetDatacentersOk returns a tuple with the Datacenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenters

`func (o *Evn) SetDatacenters(v []EvnDatacenters)`

SetDatacenters sets Datacenters field to given value.

### HasDatacenters

`func (o *Evn) HasDatacenters() bool`

HasDatacenters returns a boolean if a field has been set.

### GetEvnCSV

`func (o *Evn) GetEvnCSV() []string`

GetEvnCSV returns the EvnCSV field if non-nil, zero value otherwise.

### GetEvnCSVOk

`func (o *Evn) GetEvnCSVOk() (*[]string, bool)`

GetEvnCSVOk returns a tuple with the EvnCSV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvnCSV

`func (o *Evn) SetEvnCSV(v []string)`

SetEvnCSV sets EvnCSV field to given value.

### HasEvnCSV

`func (o *Evn) HasEvnCSV() bool`

HasEvnCSV returns a boolean if a field has been set.


