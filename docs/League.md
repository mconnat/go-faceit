# League

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Divisions** | Pointer to [**[]Division**](Division.md) | The divisions of the league. | [optional] 
**Game** | Pointer to **string** | The game of the league. | [optional] 
**Id** | Pointer to **string** | The id of the league. | [optional] 
**MinMatches** | Pointer to **int64** | The minimum matches of the league. | [optional] 
**Region** | Pointer to **string** | The region of the league. | [optional] 
**Season** | Pointer to [**Season**](Season.md) |  | [optional] 

## Methods

### NewLeague

`func NewLeague() *League`

NewLeague instantiates a new League object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeagueWithDefaults

`func NewLeagueWithDefaults() *League`

NewLeagueWithDefaults instantiates a new League object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDivisions

`func (o *League) GetDivisions() []Division`

GetDivisions returns the Divisions field if non-nil, zero value otherwise.

### GetDivisionsOk

`func (o *League) GetDivisionsOk() (*[]Division, bool)`

GetDivisionsOk returns a tuple with the Divisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisions

`func (o *League) SetDivisions(v []Division)`

SetDivisions sets Divisions field to given value.

### HasDivisions

`func (o *League) HasDivisions() bool`

HasDivisions returns a boolean if a field has been set.

### GetGame

`func (o *League) GetGame() string`

GetGame returns the Game field if non-nil, zero value otherwise.

### GetGameOk

`func (o *League) GetGameOk() (*string, bool)`

GetGameOk returns a tuple with the Game field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGame

`func (o *League) SetGame(v string)`

SetGame sets Game field to given value.

### HasGame

`func (o *League) HasGame() bool`

HasGame returns a boolean if a field has been set.

### GetId

`func (o *League) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *League) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *League) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *League) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMinMatches

`func (o *League) GetMinMatches() int64`

GetMinMatches returns the MinMatches field if non-nil, zero value otherwise.

### GetMinMatchesOk

`func (o *League) GetMinMatchesOk() (*int64, bool)`

GetMinMatchesOk returns a tuple with the MinMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMatches

`func (o *League) SetMinMatches(v int64)`

SetMinMatches sets MinMatches field to given value.

### HasMinMatches

`func (o *League) HasMinMatches() bool`

HasMinMatches returns a boolean if a field has been set.

### GetRegion

`func (o *League) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *League) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *League) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *League) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSeason

`func (o *League) GetSeason() Season`

GetSeason returns the Season field if non-nil, zero value otherwise.

### GetSeasonOk

`func (o *League) GetSeasonOk() (*Season, bool)`

GetSeasonOk returns a tuple with the Season field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeason

`func (o *League) SetSeason(v Season)`

SetSeason sets Season field to given value.

### HasSeason

`func (o *League) HasSeason() bool`

HasSeason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


