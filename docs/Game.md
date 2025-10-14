# Game

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | Pointer to [**GameAssets**](GameAssets.md) |  | [optional] 
**GameId** | Pointer to **string** |  | [optional] 
**LongLabel** | Pointer to **string** |  | [optional] 
**Order** | Pointer to **int64** |  | [optional] 
**ParentGameId** | Pointer to **string** |  | [optional] 
**Platforms** | Pointer to **[]string** |  | [optional] 
**Regions** | Pointer to **[]string** |  | [optional] 
**ShortLabel** | Pointer to **string** |  | [optional] 

## Methods

### NewGame

`func NewGame() *Game`

NewGame instantiates a new Game object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameWithDefaults

`func NewGameWithDefaults() *Game`

NewGameWithDefaults instantiates a new Game object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *Game) GetAssets() GameAssets`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *Game) GetAssetsOk() (*GameAssets, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *Game) SetAssets(v GameAssets)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *Game) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetGameId

`func (o *Game) GetGameId() string`

GetGameId returns the GameId field if non-nil, zero value otherwise.

### GetGameIdOk

`func (o *Game) GetGameIdOk() (*string, bool)`

GetGameIdOk returns a tuple with the GameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameId

`func (o *Game) SetGameId(v string)`

SetGameId sets GameId field to given value.

### HasGameId

`func (o *Game) HasGameId() bool`

HasGameId returns a boolean if a field has been set.

### GetLongLabel

`func (o *Game) GetLongLabel() string`

GetLongLabel returns the LongLabel field if non-nil, zero value otherwise.

### GetLongLabelOk

`func (o *Game) GetLongLabelOk() (*string, bool)`

GetLongLabelOk returns a tuple with the LongLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongLabel

`func (o *Game) SetLongLabel(v string)`

SetLongLabel sets LongLabel field to given value.

### HasLongLabel

`func (o *Game) HasLongLabel() bool`

HasLongLabel returns a boolean if a field has been set.

### GetOrder

`func (o *Game) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *Game) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *Game) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *Game) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetParentGameId

`func (o *Game) GetParentGameId() string`

GetParentGameId returns the ParentGameId field if non-nil, zero value otherwise.

### GetParentGameIdOk

`func (o *Game) GetParentGameIdOk() (*string, bool)`

GetParentGameIdOk returns a tuple with the ParentGameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentGameId

`func (o *Game) SetParentGameId(v string)`

SetParentGameId sets ParentGameId field to given value.

### HasParentGameId

`func (o *Game) HasParentGameId() bool`

HasParentGameId returns a boolean if a field has been set.

### GetPlatforms

`func (o *Game) GetPlatforms() []string`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *Game) GetPlatformsOk() (*[]string, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *Game) SetPlatforms(v []string)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *Game) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### GetRegions

`func (o *Game) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *Game) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *Game) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *Game) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetShortLabel

`func (o *Game) GetShortLabel() string`

GetShortLabel returns the ShortLabel field if non-nil, zero value otherwise.

### GetShortLabelOk

`func (o *Game) GetShortLabelOk() (*string, bool)`

GetShortLabelOk returns a tuple with the ShortLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortLabel

`func (o *Game) SetShortLabel(v string)`

SetShortLabel sets ShortLabel field to given value.

### HasShortLabel

`func (o *Game) HasShortLabel() bool`

HasShortLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


