# PublishStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeID** | Pointer to **string** | NodeID is the ID of the swarm node this Volume is published to. | [optional] 
**PublishContext** | Pointer to **map[string]string** | PublishContext is the PublishContext returned by the CSI plugin when a volume is published. | [optional] 
**State** | Pointer to **string** | PublishState represents the state of a Volume as it pertains to its use on a particular Node. | [optional] 

## Methods

### NewPublishStatus

`func NewPublishStatus() *PublishStatus`

NewPublishStatus instantiates a new PublishStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublishStatusWithDefaults

`func NewPublishStatusWithDefaults() *PublishStatus`

NewPublishStatusWithDefaults instantiates a new PublishStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeID

`func (o *PublishStatus) GetNodeID() string`

GetNodeID returns the NodeID field if non-nil, zero value otherwise.

### GetNodeIDOk

`func (o *PublishStatus) GetNodeIDOk() (*string, bool)`

GetNodeIDOk returns a tuple with the NodeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeID

`func (o *PublishStatus) SetNodeID(v string)`

SetNodeID sets NodeID field to given value.

### HasNodeID

`func (o *PublishStatus) HasNodeID() bool`

HasNodeID returns a boolean if a field has been set.

### GetPublishContext

`func (o *PublishStatus) GetPublishContext() map[string]string`

GetPublishContext returns the PublishContext field if non-nil, zero value otherwise.

### GetPublishContextOk

`func (o *PublishStatus) GetPublishContextOk() (*map[string]string, bool)`

GetPublishContextOk returns a tuple with the PublishContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishContext

`func (o *PublishStatus) SetPublishContext(v map[string]string)`

SetPublishContext sets PublishContext field to given value.

### HasPublishContext

`func (o *PublishStatus) HasPublishContext() bool`

HasPublishContext returns a boolean if a field has been set.

### GetState

`func (o *PublishStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PublishStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PublishStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PublishStatus) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


