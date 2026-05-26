# GetByContractResponseHitsHitsSourceEventResources

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Action** | Pointer to **[]string** | An array of items containing a string describing the action that was performed on the resource. There are a large number of possible values returned here.  Some example values are  * in the security domain &#x60;sec.user.create&#x60; for a user    creation,  * in the IP domain &#x60;ip.ipblock.reserve&#x60; for reserving an    IP block, * in the virtual resources domain &#x60;vr.firewall.activate&#x60; for    a firewall activation or * in the snapshot domain &#x60;sn.snapshot.create&#x60; for creating   a snapshot.  | [optional] |
|**Id** | Pointer to **string** | Identifier of the given resource. | [optional] |
|**Type** | Pointer to **string** | Type of the given resource. Example values are  * &#x60;datacenter&#x60;, * &#x60;firewallrule&#x60; or * &#x60;backupUnit&#x60;.  | [optional] |

## Methods

### NewGetByContractResponseHitsHitsSourceEventResources

`func NewGetByContractResponseHitsHitsSourceEventResources() *GetByContractResponseHitsHitsSourceEventResources`

NewGetByContractResponseHitsHitsSourceEventResources instantiates a new GetByContractResponseHitsHitsSourceEventResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetByContractResponseHitsHitsSourceEventResourcesWithDefaults

`func NewGetByContractResponseHitsHitsSourceEventResourcesWithDefaults() *GetByContractResponseHitsHitsSourceEventResources`

NewGetByContractResponseHitsHitsSourceEventResourcesWithDefaults instantiates a new GetByContractResponseHitsHitsSourceEventResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *GetByContractResponseHitsHitsSourceEventResources) GetAction() []string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GetByContractResponseHitsHitsSourceEventResources) GetActionOk() (*[]string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GetByContractResponseHitsHitsSourceEventResources) SetAction(v []string)`

SetAction sets Action field to given value.

### HasAction

`func (o *GetByContractResponseHitsHitsSourceEventResources) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetId

`func (o *GetByContractResponseHitsHitsSourceEventResources) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetByContractResponseHitsHitsSourceEventResources) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetByContractResponseHitsHitsSourceEventResources) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetByContractResponseHitsHitsSourceEventResources) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetByContractResponseHitsHitsSourceEventResources) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetByContractResponseHitsHitsSourceEventResources) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetByContractResponseHitsHitsSourceEventResources) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetByContractResponseHitsHitsSourceEventResources) HasType() bool`

HasType returns a boolean if a field has been set.


