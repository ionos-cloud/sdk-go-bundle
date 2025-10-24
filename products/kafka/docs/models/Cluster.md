# Cluster

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** | The name of your Kafka cluster. Must be 63 characters or less and must begin and end with an alphanumeric character (&#x60;[a-z0-9A-Z]&#x60;) with dashes (&#x60;-&#x60;), underscores (&#x60;_&#x60;), dots (&#x60;.&#x60;), and alphanumerics between.  | |
|**Version** | **string** | The version of Kafka. Allowed values are &#x60;3.7.0&#x60; and &#x60;3.8.0&#x60;. Deprecation Warning: &#x60;3.7.0&#x60; will be removed end of May 2025.  | |
|**Size** | **string** | The size of your Kafka cluster. The size of the Kafka cluster is given in T-shirt sizes. Valid values are: &#x60;XS&#x60;, &#x60;S&#x60;, &#x60;M&#x60;, &#x60;L&#x60;, and &#x60;XL&#x60;.  | |
|**Connections** | [**[]KafkaClusterConnection**](KafkaClusterConnection.md) |  | |

## Methods

### NewCluster

`func NewCluster(name string, version string, size string, connections []KafkaClusterConnection, ) *Cluster`

NewCluster instantiates a new Cluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterWithDefaults

`func NewClusterWithDefaults() *Cluster`

NewClusterWithDefaults instantiates a new Cluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Cluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Cluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Cluster) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *Cluster) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Cluster) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Cluster) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetSize

`func (o *Cluster) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Cluster) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Cluster) SetSize(v string)`

SetSize sets Size field to given value.


### GetConnections

`func (o *Cluster) GetConnections() []KafkaClusterConnection`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *Cluster) GetConnectionsOk() (*[]KafkaClusterConnection, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *Cluster) SetConnections(v []KafkaClusterConnection)`

SetConnections sets Connections field to given value.



