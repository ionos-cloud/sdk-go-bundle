# MariadbVersionRead

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The ID (UUID) of the MariadbVersion. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the MariadbVersion. | |
|**Metadata** | **map[string]interface{}** |  | [readonly] |
|**Properties** | [**MariadbVersion**](MariadbVersion.md) |  | |

## Methods

### NewMariadbVersionRead

`func NewMariadbVersionRead(id string, type_ string, href string, metadata map[string]interface{}, properties MariadbVersion, ) *MariadbVersionRead`

NewMariadbVersionRead instantiates a new MariadbVersionRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMariadbVersionReadWithDefaults

`func NewMariadbVersionReadWithDefaults() *MariadbVersionRead`

NewMariadbVersionReadWithDefaults instantiates a new MariadbVersionRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MariadbVersionRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MariadbVersionRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MariadbVersionRead) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *MariadbVersionRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MariadbVersionRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MariadbVersionRead) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *MariadbVersionRead) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *MariadbVersionRead) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *MariadbVersionRead) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *MariadbVersionRead) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MariadbVersionRead) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MariadbVersionRead) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *MariadbVersionRead) GetProperties() MariadbVersion`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *MariadbVersionRead) GetPropertiesOk() (*MariadbVersion, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *MariadbVersionRead) SetProperties(v MariadbVersion)`

SetProperties sets Properties field to given value.



