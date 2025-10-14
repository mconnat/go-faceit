# Faction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Avatar** | Pointer to **string** |  | [optional] 
**FactionId** | Pointer to **string** |  | [optional] 
**Leader** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Roster** | Pointer to [**[]Roster**](Roster.md) |  | [optional] 
**RosterV1** | Pointer to **map[string]interface{}** |  | [optional] 
**Stats** | Pointer to [**Stats**](Stats.md) |  | [optional] 
**Substituted** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewFaction

`func NewFaction() *Faction`

NewFaction instantiates a new Faction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFactionWithDefaults

`func NewFactionWithDefaults() *Faction`

NewFactionWithDefaults instantiates a new Faction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatar

`func (o *Faction) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *Faction) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *Faction) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *Faction) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetFactionId

`func (o *Faction) GetFactionId() string`

GetFactionId returns the FactionId field if non-nil, zero value otherwise.

### GetFactionIdOk

`func (o *Faction) GetFactionIdOk() (*string, bool)`

GetFactionIdOk returns a tuple with the FactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactionId

`func (o *Faction) SetFactionId(v string)`

SetFactionId sets FactionId field to given value.

### HasFactionId

`func (o *Faction) HasFactionId() bool`

HasFactionId returns a boolean if a field has been set.

### GetLeader

`func (o *Faction) GetLeader() string`

GetLeader returns the Leader field if non-nil, zero value otherwise.

### GetLeaderOk

`func (o *Faction) GetLeaderOk() (*string, bool)`

GetLeaderOk returns a tuple with the Leader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeader

`func (o *Faction) SetLeader(v string)`

SetLeader sets Leader field to given value.

### HasLeader

`func (o *Faction) HasLeader() bool`

HasLeader returns a boolean if a field has been set.

### GetName

`func (o *Faction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Faction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Faction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Faction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoster

`func (o *Faction) GetRoster() []Roster`

GetRoster returns the Roster field if non-nil, zero value otherwise.

### GetRosterOk

`func (o *Faction) GetRosterOk() (*[]Roster, bool)`

GetRosterOk returns a tuple with the Roster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoster

`func (o *Faction) SetRoster(v []Roster)`

SetRoster sets Roster field to given value.

### HasRoster

`func (o *Faction) HasRoster() bool`

HasRoster returns a boolean if a field has been set.

### GetRosterV1

`func (o *Faction) GetRosterV1() map[string]interface{}`

GetRosterV1 returns the RosterV1 field if non-nil, zero value otherwise.

### GetRosterV1Ok

`func (o *Faction) GetRosterV1Ok() (*map[string]interface{}, bool)`

GetRosterV1Ok returns a tuple with the RosterV1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRosterV1

`func (o *Faction) SetRosterV1(v map[string]interface{})`

SetRosterV1 sets RosterV1 field to given value.

### HasRosterV1

`func (o *Faction) HasRosterV1() bool`

HasRosterV1 returns a boolean if a field has been set.

### GetStats

`func (o *Faction) GetStats() Stats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *Faction) GetStatsOk() (*Stats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *Faction) SetStats(v Stats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *Faction) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetSubstituted

`func (o *Faction) GetSubstituted() bool`

GetSubstituted returns the Substituted field if non-nil, zero value otherwise.

### GetSubstitutedOk

`func (o *Faction) GetSubstitutedOk() (*bool, bool)`

GetSubstitutedOk returns a tuple with the Substituted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstituted

`func (o *Faction) SetSubstituted(v bool)`

SetSubstituted sets Substituted field to given value.

### HasSubstituted

`func (o *Faction) HasSubstituted() bool`

HasSubstituted returns a boolean if a field has been set.

### GetType

`func (o *Faction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Faction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Faction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Faction) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


