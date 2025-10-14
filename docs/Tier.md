# Tier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the tier | [optional] 
**PointsTarget** | Pointer to **int64** | The target points for the tier | [optional] 
**Rank** | Pointer to **int64** | The rank of the tier | [optional] 

## Methods

### NewTier

`func NewTier() *Tier`

NewTier instantiates a new Tier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTierWithDefaults

`func NewTierWithDefaults() *Tier`

NewTierWithDefaults instantiates a new Tier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Tier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Tier) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Tier) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Tier) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPointsTarget

`func (o *Tier) GetPointsTarget() int64`

GetPointsTarget returns the PointsTarget field if non-nil, zero value otherwise.

### GetPointsTargetOk

`func (o *Tier) GetPointsTargetOk() (*int64, bool)`

GetPointsTargetOk returns a tuple with the PointsTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointsTarget

`func (o *Tier) SetPointsTarget(v int64)`

SetPointsTarget sets PointsTarget field to given value.

### HasPointsTarget

`func (o *Tier) HasPointsTarget() bool`

HasPointsTarget returns a boolean if a field has been set.

### GetRank

`func (o *Tier) GetRank() int64`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *Tier) GetRankOk() (*int64, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *Tier) SetRank(v int64)`

SetRank sets Rank field to given value.

### HasRank

`func (o *Tier) HasRank() bool`

HasRank returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


