# Info

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessibleTopology** | Pointer to [**[]Topology**](Topology.md) | AccessibleTopolgoy is the topology this volume is actually accessible from. | [optional] 
**CapacityBytes** | Pointer to **int64** | CapacityBytes is the capacity of the volume in bytes. A value of 0 indicates that the capacity is unknown. | [optional] 
**VolumeContext** | Pointer to **map[string]string** | VolumeContext is the context originating from the CSI storage plugin when the Volume is created. | [optional] 
**VolumeID** | Pointer to **string** | VolumeID is the ID of the Volume as seen by the CSI storage plugin. This is distinct from the Volume&#39;s Swarm ID, which is the ID used by all of the Docker Engine to refer to the Volume. If this field is blank, then the Volume has not been successfully created yet. | [optional] 

## Methods

### NewInfo

`func NewInfo() *Info`

NewInfo instantiates a new Info object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfoWithDefaults

`func NewInfoWithDefaults() *Info`

NewInfoWithDefaults instantiates a new Info object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessibleTopology

`func (o *Info) GetAccessibleTopology() []Topology`

GetAccessibleTopology returns the AccessibleTopology field if non-nil, zero value otherwise.

### GetAccessibleTopologyOk

`func (o *Info) GetAccessibleTopologyOk() (*[]Topology, bool)`

GetAccessibleTopologyOk returns a tuple with the AccessibleTopology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibleTopology

`func (o *Info) SetAccessibleTopology(v []Topology)`

SetAccessibleTopology sets AccessibleTopology field to given value.

### HasAccessibleTopology

`func (o *Info) HasAccessibleTopology() bool`

HasAccessibleTopology returns a boolean if a field has been set.

### GetCapacityBytes

`func (o *Info) GetCapacityBytes() int64`

GetCapacityBytes returns the CapacityBytes field if non-nil, zero value otherwise.

### GetCapacityBytesOk

`func (o *Info) GetCapacityBytesOk() (*int64, bool)`

GetCapacityBytesOk returns a tuple with the CapacityBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityBytes

`func (o *Info) SetCapacityBytes(v int64)`

SetCapacityBytes sets CapacityBytes field to given value.

### HasCapacityBytes

`func (o *Info) HasCapacityBytes() bool`

HasCapacityBytes returns a boolean if a field has been set.

### GetVolumeContext

`func (o *Info) GetVolumeContext() map[string]string`

GetVolumeContext returns the VolumeContext field if non-nil, zero value otherwise.

### GetVolumeContextOk

`func (o *Info) GetVolumeContextOk() (*map[string]string, bool)`

GetVolumeContextOk returns a tuple with the VolumeContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeContext

`func (o *Info) SetVolumeContext(v map[string]string)`

SetVolumeContext sets VolumeContext field to given value.

### HasVolumeContext

`func (o *Info) HasVolumeContext() bool`

HasVolumeContext returns a boolean if a field has been set.

### GetVolumeID

`func (o *Info) GetVolumeID() string`

GetVolumeID returns the VolumeID field if non-nil, zero value otherwise.

### GetVolumeIDOk

`func (o *Info) GetVolumeIDOk() (*string, bool)`

GetVolumeIDOk returns a tuple with the VolumeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeID

`func (o *Info) SetVolumeID(v string)`

SetVolumeID sets VolumeID field to given value.

### HasVolumeID

`func (o *Info) HasVolumeID() bool`

HasVolumeID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


