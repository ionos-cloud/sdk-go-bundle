# RouteRead

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The ID (UUID) of the Route. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the Route. | |
|**Metadata** | [**MetadataWithEndpoint**](MetadataWithEndpoint.md) |  | |
|**Properties** | [**Route**](Route.md) |  | |

## Methods

### NewRouteRead

`func NewRouteRead(id string, type_ string, href string, metadata MetadataWithEndpoint, properties Route, ) *RouteRead`

NewRouteRead instantiates a new RouteRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteReadWithDefaults

`func NewRouteReadWithDefaults() *RouteRead`

NewRouteReadWithDefaults instantiates a new RouteRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RouteRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RouteRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RouteRead) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *RouteRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RouteRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RouteRead) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *RouteRead) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RouteRead) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RouteRead) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *RouteRead) GetMetadata() MetadataWithEndpoint`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RouteRead) GetMetadataOk() (*MetadataWithEndpoint, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RouteRead) SetMetadata(v MetadataWithEndpoint)`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *RouteRead) GetProperties() Route`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *RouteRead) GetPropertiesOk() (*Route, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *RouteRead) SetProperties(v Route)`

SetProperties sets Properties field to given value.



