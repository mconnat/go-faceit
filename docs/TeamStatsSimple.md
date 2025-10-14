# TeamStatsSimple

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Players** | Pointer to [**[]PlayerStatsSimple**](PlayerStatsSimple.md) |  | [optional] 
**Premade** | Pointer to **map[string]interface{}** |  | [optional] 
**TeamId** | Pointer to **map[string]interface{}** |  | [optional] 
**TeamStats** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewTeamStatsSimple

`func NewTeamStatsSimple() *TeamStatsSimple`

NewTeamStatsSimple instantiates a new TeamStatsSimple object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamStatsSimpleWithDefaults

`func NewTeamStatsSimpleWithDefaults() *TeamStatsSimple`

NewTeamStatsSimpleWithDefaults instantiates a new TeamStatsSimple object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlayers

`func (o *TeamStatsSimple) GetPlayers() []PlayerStatsSimple`

GetPlayers returns the Players field if non-nil, zero value otherwise.

### GetPlayersOk

`func (o *TeamStatsSimple) GetPlayersOk() (*[]PlayerStatsSimple, bool)`

GetPlayersOk returns a tuple with the Players field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayers

`func (o *TeamStatsSimple) SetPlayers(v []PlayerStatsSimple)`

SetPlayers sets Players field to given value.

### HasPlayers

`func (o *TeamStatsSimple) HasPlayers() bool`

HasPlayers returns a boolean if a field has been set.

### GetPremade

`func (o *TeamStatsSimple) GetPremade() map[string]interface{}`

GetPremade returns the Premade field if non-nil, zero value otherwise.

### GetPremadeOk

`func (o *TeamStatsSimple) GetPremadeOk() (*map[string]interface{}, bool)`

GetPremadeOk returns a tuple with the Premade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremade

`func (o *TeamStatsSimple) SetPremade(v map[string]interface{})`

SetPremade sets Premade field to given value.

### HasPremade

`func (o *TeamStatsSimple) HasPremade() bool`

HasPremade returns a boolean if a field has been set.

### GetTeamId

`func (o *TeamStatsSimple) GetTeamId() map[string]interface{}`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *TeamStatsSimple) GetTeamIdOk() (*map[string]interface{}, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *TeamStatsSimple) SetTeamId(v map[string]interface{})`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *TeamStatsSimple) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetTeamStats

`func (o *TeamStatsSimple) GetTeamStats() map[string]map[string]interface{}`

GetTeamStats returns the TeamStats field if non-nil, zero value otherwise.

### GetTeamStatsOk

`func (o *TeamStatsSimple) GetTeamStatsOk() (*map[string]map[string]interface{}, bool)`

GetTeamStatsOk returns a tuple with the TeamStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamStats

`func (o *TeamStatsSimple) SetTeamStats(v map[string]map[string]interface{})`

SetTeamStats sets TeamStats field to given value.

### HasTeamStats

`func (o *TeamStatsSimple) HasTeamStats() bool`

HasTeamStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


