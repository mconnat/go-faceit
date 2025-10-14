# PlayerStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GameId** | Pointer to **string** |  | [optional] 
**Lifetime** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**PlayerId** | Pointer to **string** |  | [optional] 
**Segments** | Pointer to **[]map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewPlayerStats

`func NewPlayerStats() *PlayerStats`

NewPlayerStats instantiates a new PlayerStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerStatsWithDefaults

`func NewPlayerStatsWithDefaults() *PlayerStats`

NewPlayerStatsWithDefaults instantiates a new PlayerStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGameId

`func (o *PlayerStats) GetGameId() string`

GetGameId returns the GameId field if non-nil, zero value otherwise.

### GetGameIdOk

`func (o *PlayerStats) GetGameIdOk() (*string, bool)`

GetGameIdOk returns a tuple with the GameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameId

`func (o *PlayerStats) SetGameId(v string)`

SetGameId sets GameId field to given value.

### HasGameId

`func (o *PlayerStats) HasGameId() bool`

HasGameId returns a boolean if a field has been set.

### GetLifetime

`func (o *PlayerStats) GetLifetime() map[string]map[string]interface{}`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *PlayerStats) GetLifetimeOk() (*map[string]map[string]interface{}, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *PlayerStats) SetLifetime(v map[string]map[string]interface{})`

SetLifetime sets Lifetime field to given value.

### HasLifetime

`func (o *PlayerStats) HasLifetime() bool`

HasLifetime returns a boolean if a field has been set.

### GetPlayerId

`func (o *PlayerStats) GetPlayerId() string`

GetPlayerId returns the PlayerId field if non-nil, zero value otherwise.

### GetPlayerIdOk

`func (o *PlayerStats) GetPlayerIdOk() (*string, bool)`

GetPlayerIdOk returns a tuple with the PlayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerId

`func (o *PlayerStats) SetPlayerId(v string)`

SetPlayerId sets PlayerId field to given value.

### HasPlayerId

`func (o *PlayerStats) HasPlayerId() bool`

HasPlayerId returns a boolean if a field has been set.

### GetSegments

`func (o *PlayerStats) GetSegments() []map[string]map[string]interface{}`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *PlayerStats) GetSegmentsOk() (*[]map[string]map[string]interface{}, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *PlayerStats) SetSegments(v []map[string]map[string]interface{})`

SetSegments sets Segments field to given value.

### HasSegments

`func (o *PlayerStats) HasSegments() bool`

HasSegments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


