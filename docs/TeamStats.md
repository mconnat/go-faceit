# TeamStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GameId** | Pointer to **string** |  | [optional] 
**Lifetime** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Segments** | Pointer to **[]map[string]map[string]interface{}** |  | [optional] 
**TeamId** | Pointer to **string** |  | [optional] 

## Methods

### NewTeamStats

`func NewTeamStats() *TeamStats`

NewTeamStats instantiates a new TeamStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamStatsWithDefaults

`func NewTeamStatsWithDefaults() *TeamStats`

NewTeamStatsWithDefaults instantiates a new TeamStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGameId

`func (o *TeamStats) GetGameId() string`

GetGameId returns the GameId field if non-nil, zero value otherwise.

### GetGameIdOk

`func (o *TeamStats) GetGameIdOk() (*string, bool)`

GetGameIdOk returns a tuple with the GameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameId

`func (o *TeamStats) SetGameId(v string)`

SetGameId sets GameId field to given value.

### HasGameId

`func (o *TeamStats) HasGameId() bool`

HasGameId returns a boolean if a field has been set.

### GetLifetime

`func (o *TeamStats) GetLifetime() map[string]map[string]interface{}`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *TeamStats) GetLifetimeOk() (*map[string]map[string]interface{}, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *TeamStats) SetLifetime(v map[string]map[string]interface{})`

SetLifetime sets Lifetime field to given value.

### HasLifetime

`func (o *TeamStats) HasLifetime() bool`

HasLifetime returns a boolean if a field has been set.

### GetSegments

`func (o *TeamStats) GetSegments() []map[string]map[string]interface{}`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *TeamStats) GetSegmentsOk() (*[]map[string]map[string]interface{}, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *TeamStats) SetSegments(v []map[string]map[string]interface{})`

SetSegments sets Segments field to given value.

### HasSegments

`func (o *TeamStats) HasSegments() bool`

HasSegments returns a boolean if a field has been set.

### GetTeamId

`func (o *TeamStats) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *TeamStats) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *TeamStats) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *TeamStats) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


