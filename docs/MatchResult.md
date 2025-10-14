# MatchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Score** | Pointer to **map[string]int64** |  | [optional] 
**Winner** | Pointer to **string** |  | [optional] 

## Methods

### NewMatchResult

`func NewMatchResult() *MatchResult`

NewMatchResult instantiates a new MatchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatchResultWithDefaults

`func NewMatchResultWithDefaults() *MatchResult`

NewMatchResultWithDefaults instantiates a new MatchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScore

`func (o *MatchResult) GetScore() map[string]int64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *MatchResult) GetScoreOk() (*map[string]int64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *MatchResult) SetScore(v map[string]int64)`

SetScore sets Score field to given value.

### HasScore

`func (o *MatchResult) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetWinner

`func (o *MatchResult) GetWinner() string`

GetWinner returns the Winner field if non-nil, zero value otherwise.

### GetWinnerOk

`func (o *MatchResult) GetWinnerOk() (*string, bool)`

GetWinnerOk returns a tuple with the Winner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinner

`func (o *MatchResult) SetWinner(v string)`

SetWinner sets Winner field to given value.

### HasWinner

`func (o *MatchResult) HasWinner() bool`

HasWinner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


