# NodePoolResponseData

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The unique ID of the resource. Must conform to the UUID format.  | |
|**Type** | Pointer to **string** | The type of the resource. | [optional] |
|**Href** | Pointer to **string** | URL to the object representation (absolute path). | [optional] |
|**Metadata** | [**Metadata**](Metadata.md) |  | |
|**Properties** | [**NodePool**](NodePool.md) |  | |

## Methods

### NewNodePoolResponseData

`func NewNodePoolResponseData(id string, metadata Metadata, properties NodePool, ) *NodePoolResponseData`

NewNodePoolResponseData instantiates a new NodePoolResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodePoolResponseDataWithDefaults

`func NewNodePoolResponseDataWithDefaults() *NodePoolResponseData`

NewNodePoolResponseDataWithDefaults instantiates a new NodePoolResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NodePoolResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodePoolResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodePoolResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *NodePoolResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NodePoolResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NodePoolResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NodePoolResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *NodePoolResponseData) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *NodePoolResponseData) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *NodePoolResponseData) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *NodePoolResponseData) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetMetadata

`func (o *NodePoolResponseData) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NodePoolResponseData) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NodePoolResponseData) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *NodePoolResponseData) GetProperties() NodePool`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *NodePoolResponseData) GetPropertiesOk() (*NodePool, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *NodePoolResponseData) SetProperties(v NodePool)`

SetProperties sets Properties field to given value.



