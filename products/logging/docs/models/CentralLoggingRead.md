# CentralLoggingRead

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The ID (UUID) of the CentralLogging. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the CentralLogging. | |
|**Metadata** | [**MetadataForCentralLogging**](MetadataForCentralLogging.md) |  | |
|**Properties** | [**CentralLogging**](CentralLogging.md) |  | |

## Methods

### NewCentralLoggingRead

`func NewCentralLoggingRead(id string, type_ string, href string, metadata MetadataForCentralLogging, properties CentralLogging, ) *CentralLoggingRead`

NewCentralLoggingRead instantiates a new CentralLoggingRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCentralLoggingReadWithDefaults

`func NewCentralLoggingReadWithDefaults() *CentralLoggingRead`

NewCentralLoggingReadWithDefaults instantiates a new CentralLoggingRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CentralLoggingRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CentralLoggingRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CentralLoggingRead) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *CentralLoggingRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CentralLoggingRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CentralLoggingRead) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *CentralLoggingRead) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CentralLoggingRead) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CentralLoggingRead) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *CentralLoggingRead) GetMetadata() MetadataForCentralLogging`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CentralLoggingRead) GetMetadataOk() (*MetadataForCentralLogging, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CentralLoggingRead) SetMetadata(v MetadataForCentralLogging)`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *CentralLoggingRead) GetProperties() CentralLogging`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CentralLoggingRead) GetPropertiesOk() (*CentralLogging, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CentralLoggingRead) SetProperties(v CentralLogging)`

SetProperties sets Properties field to given value.



