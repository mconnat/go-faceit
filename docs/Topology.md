# Topology

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Segments** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewTopology

`func NewTopology() *Topology`

NewTopology instantiates a new Topology object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyWithDefaults

`func NewTopologyWithDefaults() *Topology`

NewTopologyWithDefaults instantiates a new Topology object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegments

`func (o *Topology) GetSegments() map[string]string`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *Topology) GetSegmentsOk() (*map[string]string, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *Topology) SetSegments(v map[string]string)`

SetSegments sets Segments field to given value.

### HasSegments

`func (o *Topology) HasSegments() bool`

HasSegments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


