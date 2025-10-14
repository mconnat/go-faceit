# ClanSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Avatar** | Pointer to **string** | The clan&#39;s avatar url | [optional] 
**Game** | Pointer to **string** | The game of the clan | [optional] 
**Id** | Pointer to **string** | The id of the clan | [optional] 
**Join** | Pointer to **string** | The clan&#39;s join type | [optional] 
**LastMatchFinished** | Pointer to **time.Time** | The time the clan&#39;s last match finished | [optional] 
**MatchesCount24h** | Pointer to **int64** | The clan&#39;s matches count in the last 24 hours | [optional] 
**MaxSkillLevel** | Pointer to **int64** | The clan&#39;s maximum skill level | [optional] 
**MembersCount** | Pointer to **int64** | The clan&#39;s members count | [optional] 
**MembersCount24h** | Pointer to **int64** | The clan&#39;s members count in the last 24 hours | [optional] 
**MinSkillLevel** | Pointer to **int64** | The clan&#39;s minimum skill level | [optional] 
**Name** | Pointer to **string** | The name of the clan | [optional] 
**OrganizerId** | Pointer to **string** | The clan&#39;s organizer id | [optional] 
**Region** | Pointer to **string** | The region of the clan | [optional] 
**Type** | Pointer to **string** | The type of the clan | [optional] 

## Methods

### NewClanSearch

`func NewClanSearch() *ClanSearch`

NewClanSearch instantiates a new ClanSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClanSearchWithDefaults

`func NewClanSearchWithDefaults() *ClanSearch`

NewClanSearchWithDefaults instantiates a new ClanSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatar

`func (o *ClanSearch) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *ClanSearch) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *ClanSearch) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *ClanSearch) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetGame

`func (o *ClanSearch) GetGame() string`

GetGame returns the Game field if non-nil, zero value otherwise.

### GetGameOk

`func (o *ClanSearch) GetGameOk() (*string, bool)`

GetGameOk returns a tuple with the Game field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGame

`func (o *ClanSearch) SetGame(v string)`

SetGame sets Game field to given value.

### HasGame

`func (o *ClanSearch) HasGame() bool`

HasGame returns a boolean if a field has been set.

### GetId

`func (o *ClanSearch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClanSearch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClanSearch) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClanSearch) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJoin

`func (o *ClanSearch) GetJoin() string`

GetJoin returns the Join field if non-nil, zero value otherwise.

### GetJoinOk

`func (o *ClanSearch) GetJoinOk() (*string, bool)`

GetJoinOk returns a tuple with the Join field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoin

`func (o *ClanSearch) SetJoin(v string)`

SetJoin sets Join field to given value.

### HasJoin

`func (o *ClanSearch) HasJoin() bool`

HasJoin returns a boolean if a field has been set.

### GetLastMatchFinished

`func (o *ClanSearch) GetLastMatchFinished() time.Time`

GetLastMatchFinished returns the LastMatchFinished field if non-nil, zero value otherwise.

### GetLastMatchFinishedOk

`func (o *ClanSearch) GetLastMatchFinishedOk() (*time.Time, bool)`

GetLastMatchFinishedOk returns a tuple with the LastMatchFinished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMatchFinished

`func (o *ClanSearch) SetLastMatchFinished(v time.Time)`

SetLastMatchFinished sets LastMatchFinished field to given value.

### HasLastMatchFinished

`func (o *ClanSearch) HasLastMatchFinished() bool`

HasLastMatchFinished returns a boolean if a field has been set.

### GetMatchesCount24h

`func (o *ClanSearch) GetMatchesCount24h() int64`

GetMatchesCount24h returns the MatchesCount24h field if non-nil, zero value otherwise.

### GetMatchesCount24hOk

`func (o *ClanSearch) GetMatchesCount24hOk() (*int64, bool)`

GetMatchesCount24hOk returns a tuple with the MatchesCount24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchesCount24h

`func (o *ClanSearch) SetMatchesCount24h(v int64)`

SetMatchesCount24h sets MatchesCount24h field to given value.

### HasMatchesCount24h

`func (o *ClanSearch) HasMatchesCount24h() bool`

HasMatchesCount24h returns a boolean if a field has been set.

### GetMaxSkillLevel

`func (o *ClanSearch) GetMaxSkillLevel() int64`

GetMaxSkillLevel returns the MaxSkillLevel field if non-nil, zero value otherwise.

### GetMaxSkillLevelOk

`func (o *ClanSearch) GetMaxSkillLevelOk() (*int64, bool)`

GetMaxSkillLevelOk returns a tuple with the MaxSkillLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSkillLevel

`func (o *ClanSearch) SetMaxSkillLevel(v int64)`

SetMaxSkillLevel sets MaxSkillLevel field to given value.

### HasMaxSkillLevel

`func (o *ClanSearch) HasMaxSkillLevel() bool`

HasMaxSkillLevel returns a boolean if a field has been set.

### GetMembersCount

`func (o *ClanSearch) GetMembersCount() int64`

GetMembersCount returns the MembersCount field if non-nil, zero value otherwise.

### GetMembersCountOk

`func (o *ClanSearch) GetMembersCountOk() (*int64, bool)`

GetMembersCountOk returns a tuple with the MembersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCount

`func (o *ClanSearch) SetMembersCount(v int64)`

SetMembersCount sets MembersCount field to given value.

### HasMembersCount

`func (o *ClanSearch) HasMembersCount() bool`

HasMembersCount returns a boolean if a field has been set.

### GetMembersCount24h

`func (o *ClanSearch) GetMembersCount24h() int64`

GetMembersCount24h returns the MembersCount24h field if non-nil, zero value otherwise.

### GetMembersCount24hOk

`func (o *ClanSearch) GetMembersCount24hOk() (*int64, bool)`

GetMembersCount24hOk returns a tuple with the MembersCount24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCount24h

`func (o *ClanSearch) SetMembersCount24h(v int64)`

SetMembersCount24h sets MembersCount24h field to given value.

### HasMembersCount24h

`func (o *ClanSearch) HasMembersCount24h() bool`

HasMembersCount24h returns a boolean if a field has been set.

### GetMinSkillLevel

`func (o *ClanSearch) GetMinSkillLevel() int64`

GetMinSkillLevel returns the MinSkillLevel field if non-nil, zero value otherwise.

### GetMinSkillLevelOk

`func (o *ClanSearch) GetMinSkillLevelOk() (*int64, bool)`

GetMinSkillLevelOk returns a tuple with the MinSkillLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSkillLevel

`func (o *ClanSearch) SetMinSkillLevel(v int64)`

SetMinSkillLevel sets MinSkillLevel field to given value.

### HasMinSkillLevel

`func (o *ClanSearch) HasMinSkillLevel() bool`

HasMinSkillLevel returns a boolean if a field has been set.

### GetName

`func (o *ClanSearch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClanSearch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClanSearch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClanSearch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganizerId

`func (o *ClanSearch) GetOrganizerId() string`

GetOrganizerId returns the OrganizerId field if non-nil, zero value otherwise.

### GetOrganizerIdOk

`func (o *ClanSearch) GetOrganizerIdOk() (*string, bool)`

GetOrganizerIdOk returns a tuple with the OrganizerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizerId

`func (o *ClanSearch) SetOrganizerId(v string)`

SetOrganizerId sets OrganizerId field to given value.

### HasOrganizerId

`func (o *ClanSearch) HasOrganizerId() bool`

HasOrganizerId returns a boolean if a field has been set.

### GetRegion

`func (o *ClanSearch) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ClanSearch) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ClanSearch) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ClanSearch) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetType

`func (o *ClanSearch) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClanSearch) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClanSearch) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClanSearch) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


