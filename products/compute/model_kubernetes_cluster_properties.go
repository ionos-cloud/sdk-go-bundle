/*
 * CLOUD API
 *
 *  IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package compute

import (
	"encoding/json"
)

// checks if the KubernetesClusterProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesClusterProperties{}

// KubernetesClusterProperties struct for KubernetesClusterProperties
type KubernetesClusterProperties struct {
	// A Kubernetes cluster name. Valid Kubernetes cluster name must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between.
	Name string `json:"name"`
	// The Kubernetes version the cluster is running. This imposes restrictions on what Kubernetes versions can be run in a cluster's nodepools. Additionally, not all Kubernetes versions are viable upgrade targets for all prior versions.
	K8sVersion        *string                      `json:"k8sVersion,omitempty"`
	MaintenanceWindow *KubernetesMaintenanceWindow `json:"maintenanceWindow,omitempty"`
	// List of available versions for upgrading the cluster
	AvailableUpgradeVersions []string `json:"availableUpgradeVersions,omitempty"`
	// List of versions that may be used for node pools under this cluster
	ViableNodePoolVersions []string `json:"viableNodePoolVersions,omitempty"`
	// The indicator if the cluster is public or private. Be aware that setting it to false is currently in beta phase.
	Public *bool `json:"public,omitempty"`
	// The location of the cluster if the cluster is private. This property is immutable. The location must be enabled for your contract or you must have a Datacenter within that location. This attribute is mandatory if the cluster is private.
	Location *string `json:"location,omitempty"`
	// The nat gateway IP of the cluster if the cluster is private. This property is immutable. Must be a reserved IP in the same location as the cluster's location. This attribute is mandatory if the cluster is private.
	NatGatewayIp *string `json:"natGatewayIp,omitempty"`
	// The node subnet of the cluster, if the cluster is private. This property is optional and immutable. Must be a valid CIDR notation for an IPv4 network prefix of 16 bits length.
	NodeSubnet *string `json:"nodeSubnet,omitempty"`
	// Access to the K8s API server is restricted to these CIDRs. Traffic, internal to the cluster, is not affected by this restriction. If no allowlist is specified, access is not restricted. If an IP without subnet mask is provided, the default value is used: 32 for IPv4 and 128 for IPv6.
	ApiSubnetAllowList []string `json:"apiSubnetAllowList,omitempty"`
	// List of Object storage buckets configured for K8s usage. For now it contains only one bucket used to store K8s API audit logs
	S3Buckets []S3Bucket `json:"s3Buckets,omitempty"`
}

// NewKubernetesClusterProperties instantiates a new KubernetesClusterProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesClusterProperties(name string) *KubernetesClusterProperties {
	this := KubernetesClusterProperties{}

	this.Name = name
	var public bool = true
	this.Public = &public

	return &this
}

// NewKubernetesClusterPropertiesWithDefaults instantiates a new KubernetesClusterProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesClusterPropertiesWithDefaults() *KubernetesClusterProperties {
	this := KubernetesClusterProperties{}
	var public bool = true
	this.Public = &public
	return &this
}

// GetName returns the Name field value
func (o *KubernetesClusterProperties) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KubernetesClusterProperties) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KubernetesClusterProperties) SetName(v string) {
	o.Name = v
}

// GetK8sVersion returns the K8sVersion field value if set, zero value otherwise.
func (o *KubernetesClusterProperties) GetK8sVersion() string {
	if o == nil || IsNil(o.K8sVersion) {
		var ret string
		return ret
	}
	return *o.K8sVersion
}

// GetK8sVersionOk returns a tuple with the K8sVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterProperties) GetK8sVersionOk() (*string, bool) {
	if o == nil || IsNil(o.K8sVersion) {
		return nil, false
	}
	return o.K8sVersion, true
}

// HasK8sVersion returns a boolean if a field has been set.
func (o *KubernetesClusterProperties) HasK8sVersion() bool {
	if o != nil && !IsNil(o.K8sVersion) {
		return true
	}

	return false
}

// SetK8sVersion gets a reference to the given string and assigns it to the K8sVersion field.
func (o *KubernetesClusterProperties) SetK8sVersion(v string) {
	o.K8sVersion = &v
}

// GetMaintenanceWindow returns the MaintenanceWindow field value if set, zero value otherwise.
func (o *KubernetesClusterProperties) GetMaintenanceWindow() KubernetesMaintenanceWindow {
	if o == nil || IsNil(o.MaintenanceWindow) {
		var ret KubernetesMaintenanceWindow
		return ret
	}
	return *o.MaintenanceWindow
}

// GetMaintenanceWindowOk returns a tuple with the MaintenanceWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterProperties) GetMaintenanceWindowOk() (*KubernetesMaintenanceWindow, bool) {
	if o == nil || IsNil(o.MaintenanceWindow) {
		return nil, false
	}
	return o.MaintenanceWindow, true
}

// HasMaintenanceWindow returns a boolean if a field has been set.
func (o *KubernetesClusterProperties) HasMaintenanceWindow() bool {
	if o != nil && !IsNil(o.MaintenanceWindow) {
		return true
	}

	return false
}

// SetMaintenanceWindow gets a reference to the given KubernetesMaintenanceWindow and assigns it to the MaintenanceWindow field.
func (o *KubernetesClusterProperties) SetMaintenanceWindow(v KubernetesMaintenanceWindow) {
	o.MaintenanceWindow = &v
}

// GetAvailableUpgradeVersions returns the AvailableUpgradeVersions field value if set, zero value otherwise.
func (o *KubernetesClusterProperties) GetAvailableUpgradeVersions() []string {
	if o == nil || IsNil(o.AvailableUpgradeVersions) {
		var ret []string
		return ret
	}
	return o.AvailableUpgradeVersions
}

// GetAvailableUpgradeVersionsOk returns a tuple with the AvailableUpgradeVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterProperties) GetAvailableUpgradeVersionsOk() ([]string, bool) {
	if o == nil || IsNil(o.AvailableUpgradeVersions) {
		return nil, false
	}
	return o.AvailableUpgradeVersions, true
}

// HasAvailableUpgradeVersions returns a boolean if a field has been set.
func (o *KubernetesClusterProperties) HasAvailableUpgradeVersions() bool {
	if o != nil && !IsNil(o.AvailableUpgradeVersions) {
		return true
	}

	return false
}

// SetAvailableUpgradeVersions gets a reference to the given []string and assigns it to the AvailableUpgradeVersions field.
func (o *KubernetesClusterProperties) SetAvailableUpgradeVersions(v []string) {
	o.AvailableUpgradeVersions = v
}

// GetViableNodePoolVersions returns the ViableNodePoolVersions field value if set, zero value otherwise.
func (o *KubernetesClusterProperties) GetViableNodePoolVersions() []string {
	if o == nil || IsNil(o.ViableNodePoolVersions) {
		var ret []string
		return ret
	}
	return o.ViableNodePoolVersions
}

// GetViableNodePoolVersionsOk returns a tuple with the ViableNodePoolVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterProperties) GetViableNodePoolVersionsOk() ([]string, bool) {
	if o == nil || IsNil(o.ViableNodePoolVersions) {
		return nil, false
	}
	return o.ViableNodePoolVersions, true
}

// HasViableNodePoolVersions returns a boolean if a field has been set.
func (o *KubernetesClusterProperties) HasViableNodePoolVersions() bool {
	if o != nil && !IsNil(o.ViableNodePoolVersions) {
		return true
	}

	return false
}

// SetViableNodePoolVersions gets a reference to the given []string and assigns it to the ViableNodePoolVersions field.
func (o *KubernetesClusterProperties) SetViableNodePoolVersions(v []string) {
	o.ViableNodePoolVersions = v
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *KubernetesClusterProperties) GetPublic() bool {
	if o == nil || IsNil(o.Public) {
		var ret bool
		return ret
	}
	return *o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterProperties) GetPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.Public) {
		return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *KubernetesClusterProperties) HasPublic() bool {
	if o != nil && !IsNil(o.Public) {
		return true
	}

	return false
}

// SetPublic gets a reference to the given bool and assigns it to the Public field.
func (o *KubernetesClusterProperties) SetPublic(v bool) {
	o.Public = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *KubernetesClusterProperties) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterProperties) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *KubernetesClusterProperties) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *KubernetesClusterProperties) SetLocation(v string) {
	o.Location = &v
}

// GetNatGatewayIp returns the NatGatewayIp field value if set, zero value otherwise.
func (o *KubernetesClusterProperties) GetNatGatewayIp() string {
	if o == nil || IsNil(o.NatGatewayIp) {
		var ret string
		return ret
	}
	return *o.NatGatewayIp
}

// GetNatGatewayIpOk returns a tuple with the NatGatewayIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterProperties) GetNatGatewayIpOk() (*string, bool) {
	if o == nil || IsNil(o.NatGatewayIp) {
		return nil, false
	}
	return o.NatGatewayIp, true
}

// HasNatGatewayIp returns a boolean if a field has been set.
func (o *KubernetesClusterProperties) HasNatGatewayIp() bool {
	if o != nil && !IsNil(o.NatGatewayIp) {
		return true
	}

	return false
}

// SetNatGatewayIp gets a reference to the given string and assigns it to the NatGatewayIp field.
func (o *KubernetesClusterProperties) SetNatGatewayIp(v string) {
	o.NatGatewayIp = &v
}

// GetNodeSubnet returns the NodeSubnet field value if set, zero value otherwise.
func (o *KubernetesClusterProperties) GetNodeSubnet() string {
	if o == nil || IsNil(o.NodeSubnet) {
		var ret string
		return ret
	}
	return *o.NodeSubnet
}

// GetNodeSubnetOk returns a tuple with the NodeSubnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterProperties) GetNodeSubnetOk() (*string, bool) {
	if o == nil || IsNil(o.NodeSubnet) {
		return nil, false
	}
	return o.NodeSubnet, true
}

// HasNodeSubnet returns a boolean if a field has been set.
func (o *KubernetesClusterProperties) HasNodeSubnet() bool {
	if o != nil && !IsNil(o.NodeSubnet) {
		return true
	}

	return false
}

// SetNodeSubnet gets a reference to the given string and assigns it to the NodeSubnet field.
func (o *KubernetesClusterProperties) SetNodeSubnet(v string) {
	o.NodeSubnet = &v
}

// GetApiSubnetAllowList returns the ApiSubnetAllowList field value if set, zero value otherwise.
func (o *KubernetesClusterProperties) GetApiSubnetAllowList() []string {
	if o == nil || IsNil(o.ApiSubnetAllowList) {
		var ret []string
		return ret
	}
	return o.ApiSubnetAllowList
}

// GetApiSubnetAllowListOk returns a tuple with the ApiSubnetAllowList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterProperties) GetApiSubnetAllowListOk() ([]string, bool) {
	if o == nil || IsNil(o.ApiSubnetAllowList) {
		return nil, false
	}
	return o.ApiSubnetAllowList, true
}

// HasApiSubnetAllowList returns a boolean if a field has been set.
func (o *KubernetesClusterProperties) HasApiSubnetAllowList() bool {
	if o != nil && !IsNil(o.ApiSubnetAllowList) {
		return true
	}

	return false
}

// SetApiSubnetAllowList gets a reference to the given []string and assigns it to the ApiSubnetAllowList field.
func (o *KubernetesClusterProperties) SetApiSubnetAllowList(v []string) {
	o.ApiSubnetAllowList = v
}

// GetS3Buckets returns the S3Buckets field value if set, zero value otherwise.
func (o *KubernetesClusterProperties) GetS3Buckets() []S3Bucket {
	if o == nil || IsNil(o.S3Buckets) {
		var ret []S3Bucket
		return ret
	}
	return o.S3Buckets
}

// GetS3BucketsOk returns a tuple with the S3Buckets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterProperties) GetS3BucketsOk() ([]S3Bucket, bool) {
	if o == nil || IsNil(o.S3Buckets) {
		return nil, false
	}
	return o.S3Buckets, true
}

// HasS3Buckets returns a boolean if a field has been set.
func (o *KubernetesClusterProperties) HasS3Buckets() bool {
	if o != nil && !IsNil(o.S3Buckets) {
		return true
	}

	return false
}

// SetS3Buckets gets a reference to the given []S3Bucket and assigns it to the S3Buckets field.
func (o *KubernetesClusterProperties) SetS3Buckets(v []S3Bucket) {
	o.S3Buckets = v
}

func (o KubernetesClusterProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.K8sVersion) {
		toSerialize["k8sVersion"] = o.K8sVersion
	}
	if !IsNil(o.MaintenanceWindow) {
		toSerialize["maintenanceWindow"] = o.MaintenanceWindow
	}
	if !IsNil(o.AvailableUpgradeVersions) {
		toSerialize["availableUpgradeVersions"] = o.AvailableUpgradeVersions
	}
	if !IsNil(o.ViableNodePoolVersions) {
		toSerialize["viableNodePoolVersions"] = o.ViableNodePoolVersions
	}
	if !IsNil(o.Public) {
		toSerialize["public"] = o.Public
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.NatGatewayIp) {
		toSerialize["natGatewayIp"] = o.NatGatewayIp
	}
	if !IsNil(o.NodeSubnet) {
		toSerialize["nodeSubnet"] = o.NodeSubnet
	}
	if !IsNil(o.ApiSubnetAllowList) {
		toSerialize["apiSubnetAllowList"] = o.ApiSubnetAllowList
	}
	if !IsNil(o.S3Buckets) {
		toSerialize["s3Buckets"] = o.S3Buckets
	}
	return toSerialize, nil
}

type NullableKubernetesClusterProperties struct {
	value *KubernetesClusterProperties
	isSet bool
}

func (v NullableKubernetesClusterProperties) Get() *KubernetesClusterProperties {
	return v.value
}

func (v *NullableKubernetesClusterProperties) Set(val *KubernetesClusterProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesClusterProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesClusterProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesClusterProperties(val *KubernetesClusterProperties) *NullableKubernetesClusterProperties {
	return &NullableKubernetesClusterProperties{value: val, isSet: true}
}

func (v NullableKubernetesClusterProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesClusterProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
