# Brackets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Game** | Pointer to **string** |  | [optional] 
**Matches** | Pointer to [**[]BracketsMatch**](BracketsMatch.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Rounds** | Pointer to [**[]BracketsRound**](BracketsRound.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewBrackets

`func NewBrackets() *Brackets`

NewBrackets instantiates a new Brackets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBracketsWithDefaults

`func NewBracketsWithDefaults() *Brackets`

NewBracketsWithDefaults instantiates a new Brackets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGame

`func (o *Brackets) GetGame() string`

GetGame returns the Game field if non-nil, zero value otherwise.

### GetGameOk

`func (o *Brackets) GetGameOk() (*string, bool)`

GetGameOk returns a tuple with the Game field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGame

`func (o *Brackets) SetGame(v string)`

SetGame sets Game field to given value.

### HasGame

`func (o *Brackets) HasGame() bool`

HasGame returns a boolean if a field has been set.

### GetMatches

`func (o *Brackets) GetMatches() []BracketsMatch`

GetMatches returns the Matches field if non-nil, zero value otherwise.

### GetMatchesOk

`func (o *Brackets) GetMatchesOk() (*[]BracketsMatch, bool)`

GetMatchesOk returns a tuple with the Matches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatches

`func (o *Brackets) SetMatches(v []BracketsMatch)`

SetMatches sets Matches field to given value.

### HasMatches

`func (o *Brackets) HasMatches() bool`

HasMatches returns a boolean if a field has been set.

### GetName

`func (o *Brackets) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Brackets) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Brackets) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Brackets) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRounds

`func (o *Brackets) GetRounds() []BracketsRound`

GetRounds returns the Rounds field if non-nil, zero value otherwise.

### GetRoundsOk

`func (o *Brackets) GetRoundsOk() (*[]BracketsRound, bool)`

GetRoundsOk returns a tuple with the Rounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRounds

`func (o *Brackets) SetRounds(v []BracketsRound)`

SetRounds sets Rounds field to given value.

### HasRounds

`func (o *Brackets) HasRounds() bool`

HasRounds returns a boolean if a field has been set.

### GetStatus

`func (o *Brackets) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Brackets) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Brackets) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Brackets) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


