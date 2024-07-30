# WireguardPeerRead

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The ID (UUID) of the WireguardPeer. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the WireguardPeer. | |
|**Metadata** | [**WireguardPeerMetadata**](WireguardPeerMetadata.md) |  | |
|**Properties** | [**WireguardPeer**](WireguardPeer.md) |  | |

## Methods

### NewWireguardPeerRead

`func NewWireguardPeerRead(id string, type_ string, href string, metadata WireguardPeerMetadata, properties WireguardPeer, ) *WireguardPeerRead`

NewWireguardPeerRead instantiates a new WireguardPeerRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireguardPeerReadWithDefaults

`func NewWireguardPeerReadWithDefaults() *WireguardPeerRead`

NewWireguardPeerReadWithDefaults instantiates a new WireguardPeerRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WireguardPeerRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WireguardPeerRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WireguardPeerRead) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *WireguardPeerRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WireguardPeerRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WireguardPeerRead) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *WireguardPeerRead) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *WireguardPeerRead) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *WireguardPeerRead) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *WireguardPeerRead) GetMetadata() WireguardPeerMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WireguardPeerRead) GetMetadataOk() (*WireguardPeerMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WireguardPeerRead) SetMetadata(v WireguardPeerMetadata)`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *WireguardPeerRead) GetProperties() WireguardPeer`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WireguardPeerRead) GetPropertiesOk() (*WireguardPeer, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WireguardPeerRead) SetProperties(v WireguardPeer)`

SetProperties sets Properties field to given value.



