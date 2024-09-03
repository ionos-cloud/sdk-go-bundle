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

// checks if the GroupProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupProperties{}

// GroupProperties struct for GroupProperties
type GroupProperties struct {
	// The name of the resource.
	Name *string `json:"name,omitempty"`
	// Create data center privilege.
	CreateDataCenter *bool `json:"createDataCenter,omitempty"`
	// Create snapshot privilege.
	CreateSnapshot *bool `json:"createSnapshot,omitempty"`
	// Reserve IP block privilege.
	ReserveIp *bool `json:"reserveIp,omitempty"`
	// Activity log access privilege.
	AccessActivityLog *bool `json:"accessActivityLog,omitempty"`
	// User privilege to create a cross connect.
	CreatePcc *bool `json:"createPcc,omitempty"`
	// S3 privilege.
	S3Privilege *bool `json:"s3Privilege,omitempty"`
	// Create backup unit privilege.
	CreateBackupUnit *bool `json:"createBackupUnit,omitempty"`
	// Create internet access privilege.
	CreateInternetAccess *bool `json:"createInternetAccess,omitempty"`
	// Create Kubernetes cluster privilege.
	CreateK8sCluster *bool `json:"createK8sCluster,omitempty"`
	// Create Flow Logs privilege.
	CreateFlowLog *bool `json:"createFlowLog,omitempty"`
	// Privilege for a group to access and manage monitoring related functionality (access metrics, CRUD on alarms, alarm-actions etc) using Monotoring-as-a-Service (MaaS).
	AccessAndManageMonitoring *bool `json:"accessAndManageMonitoring,omitempty"`
	// Privilege for a group to access and manage certificates.
	AccessAndManageCertificates *bool `json:"accessAndManageCertificates,omitempty"`
	// Privilege for a group to manage DBaaS related functionality.
	ManageDBaaS *bool `json:"manageDBaaS,omitempty"`
	// Privilege for a group to access and manage dns records.
	AccessAndManageDns *bool `json:"accessAndManageDns,omitempty"`
	// Privilege for group accessing container registry related functionality.
	ManageRegistry *bool `json:"manageRegistry,omitempty"`
	// Privilege for a group to access and manage the Data Platform.
	ManageDataplatform *bool `json:"manageDataplatform,omitempty"`
}

// NewGroupProperties instantiates a new GroupProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupProperties() *GroupProperties {
	this := GroupProperties{}

	return &this
}

// NewGroupPropertiesWithDefaults instantiates a new GroupProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupPropertiesWithDefaults() *GroupProperties {
	this := GroupProperties{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GroupProperties) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupProperties) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GroupProperties) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GroupProperties) SetName(v string) {
	o.Name = &v
}

// GetCreateDataCenter returns the CreateDataCenter field value if set, zero value otherwise.
func (o *GroupProperties) GetCreateDataCenter() bool {
	if o == nil || IsNil(o.CreateDataCenter) {
		var ret bool
		return ret
	}
	return *o.CreateDataCenter
}

// GetCreateDataCenterOk returns a tuple with the CreateDataCenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupProperties) GetCreateDataCenterOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateDataCenter) {
		return nil, false
	}
	return o.CreateDataCenter, true
}

// HasCreateDataCenter returns a boolean if a field has been set.
func (o *GroupProperties) HasCreateDataCenter() bool {
	if o != nil && !IsNil(o.CreateDataCenter) {
		return true
	}

	return false
}

// SetCreateDataCenter gets a reference to the given bool and assigns it to the CreateDataCenter field.
func (o *GroupProperties) SetCreateDataCenter(v bool) {
	o.CreateDataCenter = &v
}

// GetCreateSnapshot returns the CreateSnapshot field value if set, zero value otherwise.
func (o *GroupProperties) GetCreateSnapshot() bool {
	if o == nil || IsNil(o.CreateSnapshot) {
		var ret bool
		return ret
	}
	return *o.CreateSnapshot
}

// GetCreateSnapshotOk returns a tuple with the CreateSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupProperties) GetCreateSnapshotOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateSnapshot) {
		return nil, false
	}
	return o.CreateSnapshot, true
}

// HasCreateSnapshot returns a boolean if a field has been set.
func (o *GroupProperties) HasCreateSnapshot() bool {
	if o != nil && !IsNil(o.CreateSnapshot) {
		return true
	}

	return false
}

// SetCreateSnapshot gets a reference to the given bool and assigns it to the CreateSnapshot field.
func (o *GroupProperties) SetCreateSnapshot(v bool) {
	o.CreateSnapshot = &v
}

// GetReserveIp returns the ReserveIp field value if set, zero value otherwise.
func (o *GroupProperties) GetReserveIp() bool {
	if o == nil || IsNil(o.ReserveIp) {
		var ret bool
		return ret
	}
	return *o.ReserveIp
}

// GetReserveIpOk returns a tuple with the ReserveIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupProperties) GetReserveIpOk() (*bool, bool) {
	if o == nil || IsNil(o.ReserveIp) {
		return nil, false
	}
	return o.ReserveIp, true
}

// HasReserveIp returns a boolean if a field has been set.
func (o *GroupProperties) HasReserveIp() bool {
	if o != nil && !IsNil(o.ReserveIp) {
		return true
	}

	return false
}

// SetReserveIp gets a reference to the given bool and assigns it to the ReserveIp field.
func (o *GroupProperties) SetReserveIp(v bool) {
	o.ReserveIp = &v
}

// GetAccessActivityLog returns the AccessActivityLog field value if set, zero value otherwise.
func (o *GroupProperties) GetAccessActivityLog() bool {
	if o == nil || IsNil(o.AccessActivityLog) {
		var ret bool
		return ret
	}
	return *o.AccessActivityLog
}

// GetAccessActivityLogOk returns a tuple with the AccessActivityLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupProperties) GetAccessActivityLogOk() (*bool, bool) {
	if o == nil || IsNil(o.AccessActivityLog) {
		return nil, false
	}
	return o.AccessActivityLog, true
}

// HasAccessActivityLog returns a boolean if a field has been set.
func (o *GroupProperties) HasAccessActivityLog() bool {
	if o != nil && !IsNil(o.AccessActivityLog) {
		return true
	}

	return false
}

// SetAccessActivityLog gets a reference to the given bool and assigns it to the AccessActivityLog field.
func (o *GroupProperties) SetAccessActivityLog(v bool) {
	o.AccessActivityLog = &v
}

// GetCreatePcc returns the CreatePcc field value if set, zero value otherwise.
func (o *GroupProperties) GetCreatePcc() bool {
	if o == nil || IsNil(o.CreatePcc) {
		var ret bool
		return ret
	}
	return *o.CreatePcc
}

// GetCreatePccOk returns a tuple with the CreatePcc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupProperties) GetCreatePccOk() (*bool, bool) {
	if o == nil || IsNil(o.CreatePcc) {
		return nil, false
	}
	return o.CreatePcc, true
}

// HasCreatePcc returns a boolean if a field has been set.
func (o *GroupProperties) HasCreatePcc() bool {
	if o != nil && !IsNil(o.CreatePcc) {
		return true
	}

	return false
}

// SetCreatePcc gets a reference to the given bool and assigns it to the CreatePcc field.
func (o *GroupProperties) SetCreatePcc(v bool) {
	o.CreatePcc = &v
}

// GetS3Privilege returns the S3Privilege field value if set, zero value otherwise.
func (o *GroupProperties) GetS3Privilege() bool {
	if o == nil || IsNil(o.S3Privilege) {
		var ret bool
		return ret
	}
	return *o.S3Privilege
}

// GetS3PrivilegeOk returns a tuple with the S3Privilege field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupProperties) GetS3PrivilegeOk() (*bool, bool) {
	if o == nil || IsNil(o.S3Privilege) {
		return nil, false
	}
	return o.S3Privilege, true
}

// HasS3Privilege returns a boolean if a field has been set.
func (o *GroupProperties) HasS3Privilege() bool {
	if o != nil && !IsNil(o.S3Privilege) {
		return true
	}

	return false
}

// SetS3Privilege gets a reference to the given bool and assigns it to the S3Privilege field.
func (o *GroupProperties) SetS3Privilege(v bool) {
	o.S3Privilege = &v
}

// GetCreateBackupUnit returns the CreateBackupUnit field value if set, zero value otherwise.
func (o *GroupProperties) GetCreateBackupUnit() bool {
	if o == nil || IsNil(o.CreateBackupUnit) {
		var ret bool
		return ret
	}
	return *o.CreateBackupUnit
}

// GetCreateBackupUnitOk returns a tuple with the CreateBackupUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupProperties) GetCreateBackupUnitOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateBackupUnit) {
		return nil, false
	}
	return o.CreateBackupUnit, true
}

// HasCreateBackupUnit returns a boolean if a field has been set.
func (o *GroupProperties) HasCreateBackupUnit() bool {
	if o != nil && !IsNil(o.CreateBackupUnit) {
		return true
	}

	return false
}

// SetCreateBackupUnit gets a reference to the given bool and assigns it to the CreateBackupUnit field.
func (o *GroupProperties) SetCreateBackupUnit(v bool) {
	o.CreateBackupUnit = &v
}

// GetCreateInternetAccess returns the CreateInternetAccess field value if set, zero value otherwise.
func (o *GroupProperties) GetCreateInternetAccess() bool {
	if o == nil || IsNil(o.CreateInternetAccess) {
		var ret bool
		return ret
	}
	return *o.CreateInternetAccess
}

// GetCreateInternetAccessOk returns a tuple with the CreateInternetAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupProperties) GetCreateInternetAccessOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateInternetAccess) {
		return nil, false
	}
	return o.CreateInternetAccess, true
}

// HasCreateInternetAccess returns a boolean if a field has been set.
func (o *GroupProperties) HasCreateInternetAccess() bool {
	if o != nil && !IsNil(o.CreateInternetAccess) {
		return true
	}

	return false
}

// SetCreateInternetAccess gets a reference to the given bool and assigns it to the CreateInternetAccess field.
func (o *GroupProperties) SetCreateInternetAccess(v bool) {
	o.CreateInternetAccess = &v
}

// GetCreateK8sCluster returns the CreateK8sCluster field value if set, zero value otherwise.
func (o *GroupProperties) GetCreateK8sCluster() bool {
	if o == nil || IsNil(o.CreateK8sCluster) {
		var ret bool
		return ret
	}
	return *o.CreateK8sCluster
}

// GetCreateK8sClusterOk returns a tuple with the CreateK8sCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupProperties) GetCreateK8sClusterOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateK8sCluster) {
		return nil, false
	}
	return o.CreateK8sCluster, true
}

// HasCreateK8sCluster returns a boolean if a field has been set.
func (o *GroupProperties) HasCreateK8sCluster() bool {
	if o != nil && !IsNil(o.CreateK8sCluster) {
		return true
	}

	return false
}

// SetCreateK8sCluster gets a reference to the given bool and assigns it to the CreateK8sCluster field.
func (o *GroupProperties) SetCreateK8sCluster(v bool) {
	o.CreateK8sCluster = &v
}

// GetCreateFlowLog returns the CreateFlowLog field value if set, zero value otherwise.
func (o *GroupProperties) GetCreateFlowLog() bool {
	if o == nil || IsNil(o.CreateFlowLog) {
		var ret bool
		return ret
	}
	return *o.CreateFlowLog
}

// GetCreateFlowLogOk returns a tuple with the CreateFlowLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupProperties) GetCreateFlowLogOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateFlowLog) {
		return nil, false
	}
	return o.CreateFlowLog, true
}

// HasCreateFlowLog returns a boolean if a field has been set.
func (o *GroupProperties) HasCreateFlowLog() bool {
	if o != nil && !IsNil(o.CreateFlowLog) {
		return true
	}

	return false
}

// SetCreateFlowLog gets a reference to the given bool and assigns it to the CreateFlowLog field.
func (o *GroupProperties) SetCreateFlowLog(v bool) {
	o.CreateFlowLog = &v
}

// GetAccessAndManageMonitoring returns the AccessAndManageMonitoring field value if set, zero value otherwise.
func (o *GroupProperties) GetAccessAndManageMonitoring() bool {
	if o == nil || IsNil(o.AccessAndManageMonitoring) {
		var ret bool
		return ret
	}
	return *o.AccessAndManageMonitoring
}

// GetAccessAndManageMonitoringOk returns a tuple with the AccessAndManageMonitoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupProperties) GetAccessAndManageMonitoringOk() (*bool, bool) {
	if o == nil || IsNil(o.AccessAndManageMonitoring) {
		return nil, false
	}
	return o.AccessAndManageMonitoring, true
}

// HasAccessAndManageMonitoring returns a boolean if a field has been set.
func (o *GroupProperties) HasAccessAndManageMonitoring() bool {
	if o != nil && !IsNil(o.AccessAndManageMonitoring) {
		return true
	}

	return false
}

// SetAccessAndManageMonitoring gets a reference to the given bool and assigns it to the AccessAndManageMonitoring field.
func (o *GroupProperties) SetAccessAndManageMonitoring(v bool) {
	o.AccessAndManageMonitoring = &v
}

// GetAccessAndManageCertificates returns the AccessAndManageCertificates field value if set, zero value otherwise.
func (o *GroupProperties) GetAccessAndManageCertificates() bool {
	if o == nil || IsNil(o.AccessAndManageCertificates) {
		var ret bool
		return ret
	}
	return *o.AccessAndManageCertificates
}

// GetAccessAndManageCertificatesOk returns a tuple with the AccessAndManageCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupProperties) GetAccessAndManageCertificatesOk() (*bool, bool) {
	if o == nil || IsNil(o.AccessAndManageCertificates) {
		return nil, false
	}
	return o.AccessAndManageCertificates, true
}

// HasAccessAndManageCertificates returns a boolean if a field has been set.
func (o *GroupProperties) HasAccessAndManageCertificates() bool {
	if o != nil && !IsNil(o.AccessAndManageCertificates) {
		return true
	}

	return false
}

// SetAccessAndManageCertificates gets a reference to the given bool and assigns it to the AccessAndManageCertificates field.
func (o *GroupProperties) SetAccessAndManageCertificates(v bool) {
	o.AccessAndManageCertificates = &v
}

// GetManageDBaaS returns the ManageDBaaS field value if set, zero value otherwise.
func (o *GroupProperties) GetManageDBaaS() bool {
	if o == nil || IsNil(o.ManageDBaaS) {
		var ret bool
		return ret
	}
	return *o.ManageDBaaS
}

// GetManageDBaaSOk returns a tuple with the ManageDBaaS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupProperties) GetManageDBaaSOk() (*bool, bool) {
	if o == nil || IsNil(o.ManageDBaaS) {
		return nil, false
	}
	return o.ManageDBaaS, true
}

// HasManageDBaaS returns a boolean if a field has been set.
func (o *GroupProperties) HasManageDBaaS() bool {
	if o != nil && !IsNil(o.ManageDBaaS) {
		return true
	}

	return false
}

// SetManageDBaaS gets a reference to the given bool and assigns it to the ManageDBaaS field.
func (o *GroupProperties) SetManageDBaaS(v bool) {
	o.ManageDBaaS = &v
}

// GetAccessAndManageDns returns the AccessAndManageDns field value if set, zero value otherwise.
func (o *GroupProperties) GetAccessAndManageDns() bool {
	if o == nil || IsNil(o.AccessAndManageDns) {
		var ret bool
		return ret
	}
	return *o.AccessAndManageDns
}

// GetAccessAndManageDnsOk returns a tuple with the AccessAndManageDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupProperties) GetAccessAndManageDnsOk() (*bool, bool) {
	if o == nil || IsNil(o.AccessAndManageDns) {
		return nil, false
	}
	return o.AccessAndManageDns, true
}

// HasAccessAndManageDns returns a boolean if a field has been set.
func (o *GroupProperties) HasAccessAndManageDns() bool {
	if o != nil && !IsNil(o.AccessAndManageDns) {
		return true
	}

	return false
}

// SetAccessAndManageDns gets a reference to the given bool and assigns it to the AccessAndManageDns field.
func (o *GroupProperties) SetAccessAndManageDns(v bool) {
	o.AccessAndManageDns = &v
}

// GetManageRegistry returns the ManageRegistry field value if set, zero value otherwise.
func (o *GroupProperties) GetManageRegistry() bool {
	if o == nil || IsNil(o.ManageRegistry) {
		var ret bool
		return ret
	}
	return *o.ManageRegistry
}

// GetManageRegistryOk returns a tuple with the ManageRegistry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupProperties) GetManageRegistryOk() (*bool, bool) {
	if o == nil || IsNil(o.ManageRegistry) {
		return nil, false
	}
	return o.ManageRegistry, true
}

// HasManageRegistry returns a boolean if a field has been set.
func (o *GroupProperties) HasManageRegistry() bool {
	if o != nil && !IsNil(o.ManageRegistry) {
		return true
	}

	return false
}

// SetManageRegistry gets a reference to the given bool and assigns it to the ManageRegistry field.
func (o *GroupProperties) SetManageRegistry(v bool) {
	o.ManageRegistry = &v
}

// GetManageDataplatform returns the ManageDataplatform field value if set, zero value otherwise.
func (o *GroupProperties) GetManageDataplatform() bool {
	if o == nil || IsNil(o.ManageDataplatform) {
		var ret bool
		return ret
	}
	return *o.ManageDataplatform
}

// GetManageDataplatformOk returns a tuple with the ManageDataplatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupProperties) GetManageDataplatformOk() (*bool, bool) {
	if o == nil || IsNil(o.ManageDataplatform) {
		return nil, false
	}
	return o.ManageDataplatform, true
}

// HasManageDataplatform returns a boolean if a field has been set.
func (o *GroupProperties) HasManageDataplatform() bool {
	if o != nil && !IsNil(o.ManageDataplatform) {
		return true
	}

	return false
}

// SetManageDataplatform gets a reference to the given bool and assigns it to the ManageDataplatform field.
func (o *GroupProperties) SetManageDataplatform(v bool) {
	o.ManageDataplatform = &v
}

func (o GroupProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.CreateDataCenter) {
		toSerialize["createDataCenter"] = o.CreateDataCenter
	}
	if !IsNil(o.CreateSnapshot) {
		toSerialize["createSnapshot"] = o.CreateSnapshot
	}
	if !IsNil(o.ReserveIp) {
		toSerialize["reserveIp"] = o.ReserveIp
	}
	if !IsNil(o.AccessActivityLog) {
		toSerialize["accessActivityLog"] = o.AccessActivityLog
	}
	if !IsNil(o.CreatePcc) {
		toSerialize["createPcc"] = o.CreatePcc
	}
	if !IsNil(o.S3Privilege) {
		toSerialize["s3Privilege"] = o.S3Privilege
	}
	if !IsNil(o.CreateBackupUnit) {
		toSerialize["createBackupUnit"] = o.CreateBackupUnit
	}
	if !IsNil(o.CreateInternetAccess) {
		toSerialize["createInternetAccess"] = o.CreateInternetAccess
	}
	if !IsNil(o.CreateK8sCluster) {
		toSerialize["createK8sCluster"] = o.CreateK8sCluster
	}
	if !IsNil(o.CreateFlowLog) {
		toSerialize["createFlowLog"] = o.CreateFlowLog
	}
	if !IsNil(o.AccessAndManageMonitoring) {
		toSerialize["accessAndManageMonitoring"] = o.AccessAndManageMonitoring
	}
	if !IsNil(o.AccessAndManageCertificates) {
		toSerialize["accessAndManageCertificates"] = o.AccessAndManageCertificates
	}
	if !IsNil(o.ManageDBaaS) {
		toSerialize["manageDBaaS"] = o.ManageDBaaS
	}
	if !IsNil(o.AccessAndManageDns) {
		toSerialize["accessAndManageDns"] = o.AccessAndManageDns
	}
	if !IsNil(o.ManageRegistry) {
		toSerialize["manageRegistry"] = o.ManageRegistry
	}
	if !IsNil(o.ManageDataplatform) {
		toSerialize["manageDataplatform"] = o.ManageDataplatform
	}
	return toSerialize, nil
}

type NullableGroupProperties struct {
	value *GroupProperties
	isSet bool
}

func (v NullableGroupProperties) Get() *GroupProperties {
	return v.value
}

func (v *NullableGroupProperties) Set(val *GroupProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupProperties(val *GroupProperties) *NullableGroupProperties {
	return &NullableGroupProperties{value: val, isSet: true}
}

func (v NullableGroupProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
