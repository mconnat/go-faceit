# ClusterVolume

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**ID** | **string** | ID is the Swarm ID of the volume. Because cluster volumes are Swarm objects, they have an ID, unlike non-cluster volumes, which only have a Name. This ID can be used to refer to the cluster volume. | [optional] [default to null]
**Info** | [***Info**](Info.md) |  | [optional] [default to null]
**PublishStatus** | [**[]PublishStatus**](PublishStatus.md) | PublishStatus contains the status of the volume as it pertains to its publishing on Nodes. | [optional] [default to null]
**Spec** | [***ClusterVolumeSpec**](ClusterVolumeSpec.md) |  | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Version** | [***Version**](Version.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

