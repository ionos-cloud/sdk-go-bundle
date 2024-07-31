# CentralLoggingResponse

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Metadata** | Pointer to [**CentralLoggingResponseMetadata**](CentralLoggingResponseMetadata.md) |  | [optional] |
|**Properties** | Pointer to [**CentralLoggingResponseProperties**](CentralLoggingResponseProperties.md) |  | [optional] |

## Methods

### NewCentralLoggingResponse

`func NewCentralLoggingResponse() *CentralLoggingResponse`

NewCentralLoggingResponse instantiates a new CentralLoggingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCentralLoggingResponseWithDefaults

`func NewCentralLoggingResponseWithDefaults() *CentralLoggingResponse`

NewCentralLoggingResponseWithDefaults instantiates a new CentralLoggingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *CentralLoggingResponse) GetMetadata() CentralLoggingResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CentralLoggingResponse) GetMetadataOk() (*CentralLoggingResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CentralLoggingResponse) SetMetadata(v CentralLoggingResponseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CentralLoggingResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *CentralLoggingResponse) GetProperties() CentralLoggingResponseProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CentralLoggingResponse) GetPropertiesOk() (*CentralLoggingResponseProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CentralLoggingResponse) SetProperties(v CentralLoggingResponseProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CentralLoggingResponse) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


