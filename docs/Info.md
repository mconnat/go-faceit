# Info

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessibleTopology** | [**[]Topology**](Topology.md) | AccessibleTopolgoy is the topology this volume is actually accessible from. | [optional] [default to null]
**CapacityBytes** | **int64** | CapacityBytes is the capacity of the volume in bytes. A value of 0 indicates that the capacity is unknown. | [optional] [default to null]
**VolumeContext** | **map[string]string** | VolumeContext is the context originating from the CSI storage plugin when the Volume is created. | [optional] [default to null]
**VolumeID** | **string** | VolumeID is the ID of the Volume as seen by the CSI storage plugin. This is distinct from the Volume&#x27;s Swarm ID, which is the ID used by all of the Docker Engine to refer to the Volume. If this field is blank, then the Volume has not been successfully created yet. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

