# ClusterVolumeSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessMode** | Pointer to [**AccessMode**](AccessMode.md) |  | [optional] 
**AccessibilityRequirements** | Pointer to [**TopologyRequirement**](TopologyRequirement.md) |  | [optional] 
**Availability** | Pointer to **string** |  | [optional] 
**CapacityRange** | Pointer to [**CapacityRange**](CapacityRange.md) |  | [optional] 
**Group** | Pointer to **string** | Group defines the volume group of this volume. Volumes belonging to the same group can be referred to by group name when creating Services. Referring to a volume by group instructs swarm to treat volumes in that group interchangeably for the purpose of scheduling. Volumes with an empty string for a group technically all belong to the same, emptystring group. | [optional] 
**Secrets** | Pointer to [**[]Secret**](Secret.md) | Secrets defines Swarm Secrets that are passed to the CSI storage plugin when operating on this volume. | [optional] 

## Methods

### NewClusterVolumeSpec

`func NewClusterVolumeSpec() *ClusterVolumeSpec`

NewClusterVolumeSpec instantiates a new ClusterVolumeSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterVolumeSpecWithDefaults

`func NewClusterVolumeSpecWithDefaults() *ClusterVolumeSpec`

NewClusterVolumeSpecWithDefaults instantiates a new ClusterVolumeSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessMode

`func (o *ClusterVolumeSpec) GetAccessMode() AccessMode`

GetAccessMode returns the AccessMode field if non-nil, zero value otherwise.

### GetAccessModeOk

`func (o *ClusterVolumeSpec) GetAccessModeOk() (*AccessMode, bool)`

GetAccessModeOk returns a tuple with the AccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMode

`func (o *ClusterVolumeSpec) SetAccessMode(v AccessMode)`

SetAccessMode sets AccessMode field to given value.

### HasAccessMode

`func (o *ClusterVolumeSpec) HasAccessMode() bool`

HasAccessMode returns a boolean if a field has been set.

### GetAccessibilityRequirements

`func (o *ClusterVolumeSpec) GetAccessibilityRequirements() TopologyRequirement`

GetAccessibilityRequirements returns the AccessibilityRequirements field if non-nil, zero value otherwise.

### GetAccessibilityRequirementsOk

`func (o *ClusterVolumeSpec) GetAccessibilityRequirementsOk() (*TopologyRequirement, bool)`

GetAccessibilityRequirementsOk returns a tuple with the AccessibilityRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibilityRequirements

`func (o *ClusterVolumeSpec) SetAccessibilityRequirements(v TopologyRequirement)`

SetAccessibilityRequirements sets AccessibilityRequirements field to given value.

### HasAccessibilityRequirements

`func (o *ClusterVolumeSpec) HasAccessibilityRequirements() bool`

HasAccessibilityRequirements returns a boolean if a field has been set.

### GetAvailability

`func (o *ClusterVolumeSpec) GetAvailability() string`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *ClusterVolumeSpec) GetAvailabilityOk() (*string, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *ClusterVolumeSpec) SetAvailability(v string)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *ClusterVolumeSpec) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetCapacityRange

`func (o *ClusterVolumeSpec) GetCapacityRange() CapacityRange`

GetCapacityRange returns the CapacityRange field if non-nil, zero value otherwise.

### GetCapacityRangeOk

`func (o *ClusterVolumeSpec) GetCapacityRangeOk() (*CapacityRange, bool)`

GetCapacityRangeOk returns a tuple with the CapacityRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityRange

`func (o *ClusterVolumeSpec) SetCapacityRange(v CapacityRange)`

SetCapacityRange sets CapacityRange field to given value.

### HasCapacityRange

`func (o *ClusterVolumeSpec) HasCapacityRange() bool`

HasCapacityRange returns a boolean if a field has been set.

### GetGroup

`func (o *ClusterVolumeSpec) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ClusterVolumeSpec) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ClusterVolumeSpec) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ClusterVolumeSpec) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetSecrets

`func (o *ClusterVolumeSpec) GetSecrets() []Secret`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *ClusterVolumeSpec) GetSecretsOk() (*[]Secret, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *ClusterVolumeSpec) SetSecrets(v []Secret)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *ClusterVolumeSpec) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


