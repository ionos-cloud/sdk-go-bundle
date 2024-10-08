# KubernetesNodePool

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The resource&#39;s unique identifier. | [optional] [readonly] |
|**Type** | Pointer to **string** | The object type. | [optional] [readonly] |
|**Href** | Pointer to **string** | The URL to the object representation (absolute path). | [optional] [readonly] |
|**Metadata** | Pointer to [**DatacenterElementMetadata**](DatacenterElementMetadata.md) |  | [optional] |
|**Properties** | [**KubernetesNodePoolProperties**](KubernetesNodePoolProperties.md) |  | |

## Methods

### NewKubernetesNodePool

`func NewKubernetesNodePool(properties KubernetesNodePoolProperties, ) *KubernetesNodePool`

NewKubernetesNodePool instantiates a new KubernetesNodePool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNodePoolWithDefaults

`func NewKubernetesNodePoolWithDefaults() *KubernetesNodePool`

NewKubernetesNodePoolWithDefaults instantiates a new KubernetesNodePool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesNodePool) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesNodePool) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesNodePool) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KubernetesNodePool) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *KubernetesNodePool) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KubernetesNodePool) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KubernetesNodePool) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KubernetesNodePool) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *KubernetesNodePool) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *KubernetesNodePool) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *KubernetesNodePool) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *KubernetesNodePool) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetMetadata

`func (o *KubernetesNodePool) GetMetadata() DatacenterElementMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *KubernetesNodePool) GetMetadataOk() (*DatacenterElementMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *KubernetesNodePool) SetMetadata(v DatacenterElementMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *KubernetesNodePool) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *KubernetesNodePool) GetProperties() KubernetesNodePoolProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *KubernetesNodePool) GetPropertiesOk() (*KubernetesNodePoolProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *KubernetesNodePool) SetProperties(v KubernetesNodePoolProperties)`

SetProperties sets Properties field to given value.



