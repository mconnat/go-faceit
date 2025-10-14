# AccessMode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockVolume** | Pointer to **map[string]interface{}** | Intentionally empty. | [optional] 
**MountVolume** | Pointer to [**TypeMount**](TypeMount.md) |  | [optional] 
**Scope** | Pointer to **string** | Scope defines the Scope of a Cluster Volume. This is how many nodes a Volume can be accessed simultaneously on. | [optional] 
**Sharing** | Pointer to **string** | SharingMode defines the Sharing of a Cluster Volume. This is how Tasks using a Volume at the same time can use it. | [optional] 

## Methods

### NewAccessMode

`func NewAccessMode() *AccessMode`

NewAccessMode instantiates a new AccessMode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessModeWithDefaults

`func NewAccessModeWithDefaults() *AccessMode`

NewAccessModeWithDefaults instantiates a new AccessMode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockVolume

`func (o *AccessMode) GetBlockVolume() map[string]interface{}`

GetBlockVolume returns the BlockVolume field if non-nil, zero value otherwise.

### GetBlockVolumeOk

`func (o *AccessMode) GetBlockVolumeOk() (*map[string]interface{}, bool)`

GetBlockVolumeOk returns a tuple with the BlockVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockVolume

`func (o *AccessMode) SetBlockVolume(v map[string]interface{})`

SetBlockVolume sets BlockVolume field to given value.

### HasBlockVolume

`func (o *AccessMode) HasBlockVolume() bool`

HasBlockVolume returns a boolean if a field has been set.

### GetMountVolume

`func (o *AccessMode) GetMountVolume() TypeMount`

GetMountVolume returns the MountVolume field if non-nil, zero value otherwise.

### GetMountVolumeOk

`func (o *AccessMode) GetMountVolumeOk() (*TypeMount, bool)`

GetMountVolumeOk returns a tuple with the MountVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountVolume

`func (o *AccessMode) SetMountVolume(v TypeMount)`

SetMountVolume sets MountVolume field to given value.

### HasMountVolume

`func (o *AccessMode) HasMountVolume() bool`

HasMountVolume returns a boolean if a field has been set.

### GetScope

`func (o *AccessMode) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AccessMode) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AccessMode) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *AccessMode) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSharing

`func (o *AccessMode) GetSharing() string`

GetSharing returns the Sharing field if non-nil, zero value otherwise.

### GetSharingOk

`func (o *AccessMode) GetSharingOk() (*string, bool)`

GetSharingOk returns a tuple with the Sharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharing

`func (o *AccessMode) SetSharing(v string)`

SetSharing sets Sharing field to given value.

### HasSharing

`func (o *AccessMode) HasSharing() bool`

HasSharing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


