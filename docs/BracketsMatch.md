# BracketsMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FaceitUrl** | Pointer to **string** |  | [optional] 
**MatchId** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **int64** |  | [optional] 
**Results** | Pointer to [**MatchResult**](MatchResult.md) |  | [optional] 
**Round** | Pointer to **int64** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Teams** | Pointer to [**map[string]BracketsFaction**](BracketsFaction.md) |  | [optional] 

## Methods

### NewBracketsMatch

`func NewBracketsMatch() *BracketsMatch`

NewBracketsMatch instantiates a new BracketsMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBracketsMatchWithDefaults

`func NewBracketsMatchWithDefaults() *BracketsMatch`

NewBracketsMatchWithDefaults instantiates a new BracketsMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFaceitUrl

`func (o *BracketsMatch) GetFaceitUrl() string`

GetFaceitUrl returns the FaceitUrl field if non-nil, zero value otherwise.

### GetFaceitUrlOk

`func (o *BracketsMatch) GetFaceitUrlOk() (*string, bool)`

GetFaceitUrlOk returns a tuple with the FaceitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceitUrl

`func (o *BracketsMatch) SetFaceitUrl(v string)`

SetFaceitUrl sets FaceitUrl field to given value.

### HasFaceitUrl

`func (o *BracketsMatch) HasFaceitUrl() bool`

HasFaceitUrl returns a boolean if a field has been set.

### GetMatchId

`func (o *BracketsMatch) GetMatchId() string`

GetMatchId returns the MatchId field if non-nil, zero value otherwise.

### GetMatchIdOk

`func (o *BracketsMatch) GetMatchIdOk() (*string, bool)`

GetMatchIdOk returns a tuple with the MatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchId

`func (o *BracketsMatch) SetMatchId(v string)`

SetMatchId sets MatchId field to given value.

### HasMatchId

`func (o *BracketsMatch) HasMatchId() bool`

HasMatchId returns a boolean if a field has been set.

### GetPosition

`func (o *BracketsMatch) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *BracketsMatch) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *BracketsMatch) SetPosition(v int64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *BracketsMatch) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetResults

`func (o *BracketsMatch) GetResults() MatchResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BracketsMatch) GetResultsOk() (*MatchResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BracketsMatch) SetResults(v MatchResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *BracketsMatch) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetRound

`func (o *BracketsMatch) GetRound() int64`

GetRound returns the Round field if non-nil, zero value otherwise.

### GetRoundOk

`func (o *BracketsMatch) GetRoundOk() (*int64, bool)`

GetRoundOk returns a tuple with the Round field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRound

`func (o *BracketsMatch) SetRound(v int64)`

SetRound sets Round field to given value.

### HasRound

`func (o *BracketsMatch) HasRound() bool`

HasRound returns a boolean if a field has been set.

### GetState

`func (o *BracketsMatch) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BracketsMatch) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BracketsMatch) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *BracketsMatch) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTeams

`func (o *BracketsMatch) GetTeams() map[string]BracketsFaction`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *BracketsMatch) GetTeamsOk() (*map[string]BracketsFaction, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *BracketsMatch) SetTeams(v map[string]BracketsFaction)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *BracketsMatch) HasTeams() bool`

HasTeams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


