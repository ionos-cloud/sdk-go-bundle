# VersionsGet200Response

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The ID of a list of resources.  | [optional] |
|**Type** | Pointer to **string** |  | [optional] |
|**Href** | Pointer to **string** |  | [optional] |
|**Items** | **[]string** |  | |

## Methods

### NewVersionsGet200Response

`func NewVersionsGet200Response(items []string, ) *VersionsGet200Response`

NewVersionsGet200Response instantiates a new VersionsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionsGet200ResponseWithDefaults

`func NewVersionsGet200ResponseWithDefaults() *VersionsGet200Response`

NewVersionsGet200ResponseWithDefaults instantiates a new VersionsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VersionsGet200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VersionsGet200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VersionsGet200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VersionsGet200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *VersionsGet200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VersionsGet200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VersionsGet200Response) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VersionsGet200Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *VersionsGet200Response) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *VersionsGet200Response) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *VersionsGet200Response) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *VersionsGet200Response) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetItems

`func (o *VersionsGet200Response) GetItems() []string`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *VersionsGet200Response) GetItemsOk() (*[]string, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *VersionsGet200Response) SetItems(v []string)`

SetItems sets Items field to given value.



