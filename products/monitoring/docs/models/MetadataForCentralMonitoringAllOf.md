# MetadataForCentralMonitoringAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**GrafanaEndpoint** | Pointer to **string** | The endpoint of the Grafana instance.  | [optional] [readonly] |
|**Products** | Pointer to **[]string** | Products with central monitoring enabled. This is a comma-separated list of product names.  | [optional] [readonly] |

## Methods

### NewMetadataForCentralMonitoringAllOf

`func NewMetadataForCentralMonitoringAllOf() *MetadataForCentralMonitoringAllOf`

NewMetadataForCentralMonitoringAllOf instantiates a new MetadataForCentralMonitoringAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataForCentralMonitoringAllOfWithDefaults

`func NewMetadataForCentralMonitoringAllOfWithDefaults() *MetadataForCentralMonitoringAllOf`

NewMetadataForCentralMonitoringAllOfWithDefaults instantiates a new MetadataForCentralMonitoringAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrafanaEndpoint

`func (o *MetadataForCentralMonitoringAllOf) GetGrafanaEndpoint() string`

GetGrafanaEndpoint returns the GrafanaEndpoint field if non-nil, zero value otherwise.

### GetGrafanaEndpointOk

`func (o *MetadataForCentralMonitoringAllOf) GetGrafanaEndpointOk() (*string, bool)`

GetGrafanaEndpointOk returns a tuple with the GrafanaEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrafanaEndpoint

`func (o *MetadataForCentralMonitoringAllOf) SetGrafanaEndpoint(v string)`

SetGrafanaEndpoint sets GrafanaEndpoint field to given value.

### HasGrafanaEndpoint

`func (o *MetadataForCentralMonitoringAllOf) HasGrafanaEndpoint() bool`

HasGrafanaEndpoint returns a boolean if a field has been set.

### GetProducts

`func (o *MetadataForCentralMonitoringAllOf) GetProducts() []string`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *MetadataForCentralMonitoringAllOf) GetProductsOk() (*[]string, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *MetadataForCentralMonitoringAllOf) SetProducts(v []string)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *MetadataForCentralMonitoringAllOf) HasProducts() bool`

HasProducts returns a boolean if a field has been set.


