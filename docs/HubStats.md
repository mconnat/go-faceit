# HubStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GameId** | Pointer to **string** |  | [optional] 
**Players** | Pointer to [**[]StatsCompetitionPlayer**](StatsCompetitionPlayer.md) |  | [optional] 

## Methods

### NewHubStats

`func NewHubStats() *HubStats`

NewHubStats instantiates a new HubStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHubStatsWithDefaults

`func NewHubStatsWithDefaults() *HubStats`

NewHubStatsWithDefaults instantiates a new HubStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGameId

`func (o *HubStats) GetGameId() string`

GetGameId returns the GameId field if non-nil, zero value otherwise.

### GetGameIdOk

`func (o *HubStats) GetGameIdOk() (*string, bool)`

GetGameIdOk returns a tuple with the GameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameId

`func (o *HubStats) SetGameId(v string)`

SetGameId sets GameId field to given value.

### HasGameId

`func (o *HubStats) HasGameId() bool`

HasGameId returns a boolean if a field has been set.

### GetPlayers

`func (o *HubStats) GetPlayers() []StatsCompetitionPlayer`

GetPlayers returns the Players field if non-nil, zero value otherwise.

### GetPlayersOk

`func (o *HubStats) GetPlayersOk() (*[]StatsCompetitionPlayer, bool)`

GetPlayersOk returns a tuple with the Players field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayers

`func (o *HubStats) SetPlayers(v []StatsCompetitionPlayer)`

SetPlayers sets Players field to given value.

### HasPlayers

`func (o *HubStats) HasPlayers() bool`

HasPlayers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


