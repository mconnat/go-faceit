# Division

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | Pointer to [**Assets**](Assets.md) |  | [optional] 
**ConfigType** | Pointer to **string** | The type of the division. Can be nested or classic. Nested means that the division has tiers, classic is without tiers. | [optional] 
**LeaderboardConfig** | Pointer to [**LeaderboardConfig**](LeaderboardConfig.md) |  | [optional] 
**Leaderboards** | Pointer to **[]string** | The leaderboards of the division | [optional] 
**MaxElo** | Pointer to **int64** | Max ELO for a user to be placed in this division after placement matches | [optional] 
**MinElo** | Pointer to **int64** | Min ELO for a user to be placed in this division after placement matches | [optional] 
**Name** | Pointer to **string** | The name of the division. | [optional] 
**Tiers** | Pointer to [**[]Tier**](Tier.md) | The tiers of the division | [optional] 
**Type** | Pointer to **string** | The type of the division. | [optional] 

## Methods

### NewDivision

`func NewDivision() *Division`

NewDivision instantiates a new Division object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDivisionWithDefaults

`func NewDivisionWithDefaults() *Division`

NewDivisionWithDefaults instantiates a new Division object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *Division) GetAssets() Assets`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *Division) GetAssetsOk() (*Assets, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *Division) SetAssets(v Assets)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *Division) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetConfigType

`func (o *Division) GetConfigType() string`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *Division) GetConfigTypeOk() (*string, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *Division) SetConfigType(v string)`

SetConfigType sets ConfigType field to given value.

### HasConfigType

`func (o *Division) HasConfigType() bool`

HasConfigType returns a boolean if a field has been set.

### GetLeaderboardConfig

`func (o *Division) GetLeaderboardConfig() LeaderboardConfig`

GetLeaderboardConfig returns the LeaderboardConfig field if non-nil, zero value otherwise.

### GetLeaderboardConfigOk

`func (o *Division) GetLeaderboardConfigOk() (*LeaderboardConfig, bool)`

GetLeaderboardConfigOk returns a tuple with the LeaderboardConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaderboardConfig

`func (o *Division) SetLeaderboardConfig(v LeaderboardConfig)`

SetLeaderboardConfig sets LeaderboardConfig field to given value.

### HasLeaderboardConfig

`func (o *Division) HasLeaderboardConfig() bool`

HasLeaderboardConfig returns a boolean if a field has been set.

### GetLeaderboards

`func (o *Division) GetLeaderboards() []string`

GetLeaderboards returns the Leaderboards field if non-nil, zero value otherwise.

### GetLeaderboardsOk

`func (o *Division) GetLeaderboardsOk() (*[]string, bool)`

GetLeaderboardsOk returns a tuple with the Leaderboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaderboards

`func (o *Division) SetLeaderboards(v []string)`

SetLeaderboards sets Leaderboards field to given value.

### HasLeaderboards

`func (o *Division) HasLeaderboards() bool`

HasLeaderboards returns a boolean if a field has been set.

### GetMaxElo

`func (o *Division) GetMaxElo() int64`

GetMaxElo returns the MaxElo field if non-nil, zero value otherwise.

### GetMaxEloOk

`func (o *Division) GetMaxEloOk() (*int64, bool)`

GetMaxEloOk returns a tuple with the MaxElo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxElo

`func (o *Division) SetMaxElo(v int64)`

SetMaxElo sets MaxElo field to given value.

### HasMaxElo

`func (o *Division) HasMaxElo() bool`

HasMaxElo returns a boolean if a field has been set.

### GetMinElo

`func (o *Division) GetMinElo() int64`

GetMinElo returns the MinElo field if non-nil, zero value otherwise.

### GetMinEloOk

`func (o *Division) GetMinEloOk() (*int64, bool)`

GetMinEloOk returns a tuple with the MinElo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinElo

`func (o *Division) SetMinElo(v int64)`

SetMinElo sets MinElo field to given value.

### HasMinElo

`func (o *Division) HasMinElo() bool`

HasMinElo returns a boolean if a field has been set.

### GetName

`func (o *Division) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Division) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Division) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Division) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTiers

`func (o *Division) GetTiers() []Tier`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *Division) GetTiersOk() (*[]Tier, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *Division) SetTiers(v []Tier)`

SetTiers sets Tiers field to given value.

### HasTiers

`func (o *Division) HasTiers() bool`

HasTiers returns a boolean if a field has been set.

### GetType

`func (o *Division) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Division) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Division) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Division) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


