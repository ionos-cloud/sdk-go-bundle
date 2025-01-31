# ClusterCreate

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Metadata** | Pointer to **map[string]interface{}** | Metadata | [optional] |
|**Properties** | [**Cluster**](Cluster.md) |  | |

## Methods

### NewClusterCreate

`func NewClusterCreate(properties Cluster, ) *ClusterCreate`

NewClusterCreate instantiates a new ClusterCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCreateWithDefaults

`func NewClusterCreateWithDefaults() *ClusterCreate`

NewClusterCreateWithDefaults instantiates a new ClusterCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *ClusterCreate) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ClusterCreate) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ClusterCreate) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ClusterCreate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *ClusterCreate) GetProperties() Cluster`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ClusterCreate) GetPropertiesOk() (*Cluster, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ClusterCreate) SetProperties(v Cluster)`

SetProperties sets Properties field to given value.



