# ClusterRead

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The ID (UUID) of the Cluster. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the Cluster. | |
|**Metadata** | [**ClusterMetadata**](ClusterMetadata.md) |  | |
|**Properties** | [**Cluster**](Cluster.md) |  | |

## Methods

### NewClusterRead

`func NewClusterRead(id string, type_ string, href string, metadata ClusterMetadata, properties Cluster, ) *ClusterRead`

NewClusterRead instantiates a new ClusterRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterReadWithDefaults

`func NewClusterReadWithDefaults() *ClusterRead`

NewClusterReadWithDefaults instantiates a new ClusterRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterRead) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ClusterRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterRead) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *ClusterRead) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ClusterRead) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ClusterRead) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *ClusterRead) GetMetadata() ClusterMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ClusterRead) GetMetadataOk() (*ClusterMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ClusterRead) SetMetadata(v ClusterMetadata)`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *ClusterRead) GetProperties() Cluster`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ClusterRead) GetPropertiesOk() (*Cluster, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ClusterRead) SetProperties(v Cluster)`

SetProperties sets Properties field to given value.



