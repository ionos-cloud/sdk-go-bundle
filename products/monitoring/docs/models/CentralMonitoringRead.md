# CentralMonitoringRead

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The ID (UUID) of the CentralMonitoring. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the CentralMonitoring. | |
|**Metadata** | [**MetadataForCentralMonitoring**](MetadataForCentralMonitoring.md) |  | |
|**Properties** | [**CentralMonitoring**](CentralMonitoring.md) |  | |

## Methods

### NewCentralMonitoringRead

`func NewCentralMonitoringRead(id string, type_ string, href string, metadata MetadataForCentralMonitoring, properties CentralMonitoring, ) *CentralMonitoringRead`

NewCentralMonitoringRead instantiates a new CentralMonitoringRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCentralMonitoringReadWithDefaults

`func NewCentralMonitoringReadWithDefaults() *CentralMonitoringRead`

NewCentralMonitoringReadWithDefaults instantiates a new CentralMonitoringRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CentralMonitoringRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CentralMonitoringRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CentralMonitoringRead) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *CentralMonitoringRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CentralMonitoringRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CentralMonitoringRead) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *CentralMonitoringRead) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CentralMonitoringRead) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CentralMonitoringRead) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *CentralMonitoringRead) GetMetadata() MetadataForCentralMonitoring`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CentralMonitoringRead) GetMetadataOk() (*MetadataForCentralMonitoring, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CentralMonitoringRead) SetMetadata(v MetadataForCentralMonitoring)`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *CentralMonitoringRead) GetProperties() CentralMonitoring`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CentralMonitoringRead) GetPropertiesOk() (*CentralMonitoring, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CentralMonitoringRead) SetProperties(v CentralMonitoring)`

SetProperties sets Properties field to given value.



