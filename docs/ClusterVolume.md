# ClusterVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ID** | Pointer to **string** | ID is the Swarm ID of the volume. Because cluster volumes are Swarm objects, they have an ID, unlike non-cluster volumes, which only have a Name. This ID can be used to refer to the cluster volume. | [optional] 
**Info** | Pointer to [**Info**](Info.md) |  | [optional] 
**PublishStatus** | Pointer to [**[]PublishStatus**](PublishStatus.md) | PublishStatus contains the status of the volume as it pertains to its publishing on Nodes. | [optional] 
**Spec** | Pointer to [**ClusterVolumeSpec**](ClusterVolumeSpec.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Version** | Pointer to [**Version**](Version.md) |  | [optional] 

## Methods

### NewClusterVolume

`func NewClusterVolume() *ClusterVolume`

NewClusterVolume instantiates a new ClusterVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterVolumeWithDefaults

`func NewClusterVolumeWithDefaults() *ClusterVolume`

NewClusterVolumeWithDefaults instantiates a new ClusterVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ClusterVolume) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ClusterVolume) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ClusterVolume) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ClusterVolume) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetID

`func (o *ClusterVolume) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *ClusterVolume) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *ClusterVolume) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *ClusterVolume) HasID() bool`

HasID returns a boolean if a field has been set.

### GetInfo

`func (o *ClusterVolume) GetInfo() Info`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ClusterVolume) GetInfoOk() (*Info, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ClusterVolume) SetInfo(v Info)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ClusterVolume) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetPublishStatus

`func (o *ClusterVolume) GetPublishStatus() []PublishStatus`

GetPublishStatus returns the PublishStatus field if non-nil, zero value otherwise.

### GetPublishStatusOk

`func (o *ClusterVolume) GetPublishStatusOk() (*[]PublishStatus, bool)`

GetPublishStatusOk returns a tuple with the PublishStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishStatus

`func (o *ClusterVolume) SetPublishStatus(v []PublishStatus)`

SetPublishStatus sets PublishStatus field to given value.

### HasPublishStatus

`func (o *ClusterVolume) HasPublishStatus() bool`

HasPublishStatus returns a boolean if a field has been set.

### GetSpec

`func (o *ClusterVolume) GetSpec() ClusterVolumeSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ClusterVolume) GetSpecOk() (*ClusterVolumeSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ClusterVolume) SetSpec(v ClusterVolumeSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ClusterVolume) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ClusterVolume) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ClusterVolume) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ClusterVolume) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ClusterVolume) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *ClusterVolume) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ClusterVolume) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ClusterVolume) SetVersion(v Version)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ClusterVolume) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


