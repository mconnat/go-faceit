# ClusterVolumeSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessMode** | [***AccessMode**](AccessMode.md) |  | [optional] [default to null]
**AccessibilityRequirements** | [***TopologyRequirement**](TopologyRequirement.md) |  | [optional] [default to null]
**Availability** | [***Availability**](Availability.md) |  | [optional] [default to null]
**CapacityRange** | [***CapacityRange**](CapacityRange.md) |  | [optional] [default to null]
**Group** | **string** | Group defines the volume group of this volume. Volumes belonging to the same group can be referred to by group name when creating Services. Referring to a volume by group instructs swarm to treat volumes in that group interchangeably for the purpose of scheduling. Volumes with an empty string for a group technically all belong to the same, emptystring group. | [optional] [default to null]
**Secrets** | [**[]Secret**](Secret.md) | Secrets defines Swarm Secrets that are passed to the CSI storage plugin when operating on this volume. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


