# ReplicaSetRead

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The ID (UUID) of the ReplicaSet. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the ReplicaSet. | |
|**Metadata** | [**ReplicaSetMetadata**](ReplicaSetMetadata.md) |  | |
|**Properties** | [**ReplicaSet**](ReplicaSet.md) |  | |

## Methods

### NewReplicaSetRead

`func NewReplicaSetRead(id string, type_ string, href string, metadata ReplicaSetMetadata, properties ReplicaSet, ) *ReplicaSetRead`

NewReplicaSetRead instantiates a new ReplicaSetRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicaSetReadWithDefaults

`func NewReplicaSetReadWithDefaults() *ReplicaSetRead`

NewReplicaSetReadWithDefaults instantiates a new ReplicaSetRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReplicaSetRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReplicaSetRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReplicaSetRead) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ReplicaSetRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReplicaSetRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReplicaSetRead) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *ReplicaSetRead) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ReplicaSetRead) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ReplicaSetRead) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *ReplicaSetRead) GetMetadata() ReplicaSetMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ReplicaSetRead) GetMetadataOk() (*ReplicaSetMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ReplicaSetRead) SetMetadata(v ReplicaSetMetadata)`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *ReplicaSetRead) GetProperties() ReplicaSet`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ReplicaSetRead) GetPropertiesOk() (*ReplicaSet, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ReplicaSetRead) SetProperties(v ReplicaSet)`

SetProperties sets Properties field to given value.



