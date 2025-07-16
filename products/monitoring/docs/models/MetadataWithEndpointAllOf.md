# MetadataWithEndpointAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Key** | Pointer to **string** | The authentication key of the monitoring instance.  | [optional] [readonly] |
|**GrafanaEndpoint** | Pointer to **string** | The endpoint of the Grafana instance.  | [optional] [readonly] |
|**HttpEndpoint** | Pointer to **string** | The HTTP endpoint of the monitoring instance.  | [optional] [readonly] |

## Methods

### NewMetadataWithEndpointAllOf

`func NewMetadataWithEndpointAllOf() *MetadataWithEndpointAllOf`

NewMetadataWithEndpointAllOf instantiates a new MetadataWithEndpointAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataWithEndpointAllOfWithDefaults

`func NewMetadataWithEndpointAllOfWithDefaults() *MetadataWithEndpointAllOf`

NewMetadataWithEndpointAllOfWithDefaults instantiates a new MetadataWithEndpointAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *MetadataWithEndpointAllOf) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MetadataWithEndpointAllOf) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MetadataWithEndpointAllOf) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *MetadataWithEndpointAllOf) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetGrafanaEndpoint

`func (o *MetadataWithEndpointAllOf) GetGrafanaEndpoint() string`

GetGrafanaEndpoint returns the GrafanaEndpoint field if non-nil, zero value otherwise.

### GetGrafanaEndpointOk

`func (o *MetadataWithEndpointAllOf) GetGrafanaEndpointOk() (*string, bool)`

GetGrafanaEndpointOk returns a tuple with the GrafanaEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrafanaEndpoint

`func (o *MetadataWithEndpointAllOf) SetGrafanaEndpoint(v string)`

SetGrafanaEndpoint sets GrafanaEndpoint field to given value.

### HasGrafanaEndpoint

`func (o *MetadataWithEndpointAllOf) HasGrafanaEndpoint() bool`

HasGrafanaEndpoint returns a boolean if a field has been set.

### GetHttpEndpoint

`func (o *MetadataWithEndpointAllOf) GetHttpEndpoint() string`

GetHttpEndpoint returns the HttpEndpoint field if non-nil, zero value otherwise.

### GetHttpEndpointOk

`func (o *MetadataWithEndpointAllOf) GetHttpEndpointOk() (*string, bool)`

GetHttpEndpointOk returns a tuple with the HttpEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpEndpoint

`func (o *MetadataWithEndpointAllOf) SetHttpEndpoint(v string)`

SetHttpEndpoint sets HttpEndpoint field to given value.

### HasHttpEndpoint

`func (o *MetadataWithEndpointAllOf) HasHttpEndpoint() bool`

HasHttpEndpoint returns a boolean if a field has been set.


