# ProviderRead

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The ID (UUID) of the Provider. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the Provider. | |
|**Metadata** | [**MetadataWithStatus**](MetadataWithStatus.md) |  | |
|**Properties** | [**Provider**](Provider.md) |  | |

## Methods

### NewProviderRead

`func NewProviderRead(id string, type_ string, href string, metadata MetadataWithStatus, properties Provider, ) *ProviderRead`

NewProviderRead instantiates a new ProviderRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderReadWithDefaults

`func NewProviderReadWithDefaults() *ProviderRead`

NewProviderReadWithDefaults instantiates a new ProviderRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProviderRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProviderRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProviderRead) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ProviderRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProviderRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProviderRead) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *ProviderRead) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ProviderRead) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ProviderRead) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *ProviderRead) GetMetadata() MetadataWithStatus`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ProviderRead) GetMetadataOk() (*MetadataWithStatus, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ProviderRead) SetMetadata(v MetadataWithStatus)`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *ProviderRead) GetProperties() Provider`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ProviderRead) GetPropertiesOk() (*Provider, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ProviderRead) SetProperties(v Provider)`

SetProperties sets Properties field to given value.



