# AutoCertificateReadListAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | ID of the list of AutoCertificate resources. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the list of AutoCertificate resources. | |
|**Items** | Pointer to [**[]AutoCertificateRead**](AutoCertificateRead.md) | The list of AutoCertificate resources. | [optional] |

## Methods

### NewAutoCertificateReadListAllOf

`func NewAutoCertificateReadListAllOf(id string, type_ string, href string, ) *AutoCertificateReadListAllOf`

NewAutoCertificateReadListAllOf instantiates a new AutoCertificateReadListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoCertificateReadListAllOfWithDefaults

`func NewAutoCertificateReadListAllOfWithDefaults() *AutoCertificateReadListAllOf`

NewAutoCertificateReadListAllOfWithDefaults instantiates a new AutoCertificateReadListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutoCertificateReadListAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoCertificateReadListAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoCertificateReadListAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *AutoCertificateReadListAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AutoCertificateReadListAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AutoCertificateReadListAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *AutoCertificateReadListAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AutoCertificateReadListAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AutoCertificateReadListAllOf) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *AutoCertificateReadListAllOf) GetItems() []AutoCertificateRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AutoCertificateReadListAllOf) GetItemsOk() (*[]AutoCertificateRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AutoCertificateReadListAllOf) SetItems(v []AutoCertificateRead)`

SetItems sets Items field to given value.

### HasItems

`func (o *AutoCertificateReadListAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


