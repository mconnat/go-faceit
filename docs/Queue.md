# Queue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminTool** | Pointer to **bool** |  | [optional] 
**AnticheatRequired** | Pointer to **bool** |  | [optional] 
**CalculateElo** | Pointer to **bool** |  | [optional] 
**CaptainSelection** | Pointer to **string** |  | [optional] 
**CheckIn** | Pointer to [**CheckIn**](CheckIn.md) |  | [optional] 
**EntityId** | Pointer to **string** |  | [optional] 
**EntityType** | Pointer to **string** |  | [optional] 
**FbiManagement** | Pointer to **bool** |  | [optional] 
**Game** | Pointer to **string** |  | [optional] 
**GroupSimilar** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**JoinType** | Pointer to [**JoinType**](JoinType.md) |  | [optional] 
**LastModified** | Pointer to **time.Time** |  | [optional] 
**MaxSkill** | Pointer to **int64** |  | [optional] 
**MinSkill** | Pointer to **int64** |  | [optional] 
**NoOfPlayers** | Pointer to **int64** |  | [optional] 
**Open** | Pointer to **bool** |  | [optional] 
**OrganizerId** | Pointer to **string** |  | [optional] 
**Paused** | Pointer to **bool** |  | [optional] 
**QueueAlgorithm** | Pointer to [**QueueAlgorithm**](QueueAlgorithm.md) |  | [optional] 
**QueueName** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**VerifiedMatching** | Pointer to **bool** |  | [optional] 

## Methods

### NewQueue

`func NewQueue() *Queue`

NewQueue instantiates a new Queue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueWithDefaults

`func NewQueueWithDefaults() *Queue`

NewQueueWithDefaults instantiates a new Queue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminTool

`func (o *Queue) GetAdminTool() bool`

GetAdminTool returns the AdminTool field if non-nil, zero value otherwise.

### GetAdminToolOk

`func (o *Queue) GetAdminToolOk() (*bool, bool)`

GetAdminToolOk returns a tuple with the AdminTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminTool

`func (o *Queue) SetAdminTool(v bool)`

SetAdminTool sets AdminTool field to given value.

### HasAdminTool

`func (o *Queue) HasAdminTool() bool`

HasAdminTool returns a boolean if a field has been set.

### GetAnticheatRequired

`func (o *Queue) GetAnticheatRequired() bool`

GetAnticheatRequired returns the AnticheatRequired field if non-nil, zero value otherwise.

### GetAnticheatRequiredOk

`func (o *Queue) GetAnticheatRequiredOk() (*bool, bool)`

GetAnticheatRequiredOk returns a tuple with the AnticheatRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnticheatRequired

`func (o *Queue) SetAnticheatRequired(v bool)`

SetAnticheatRequired sets AnticheatRequired field to given value.

### HasAnticheatRequired

`func (o *Queue) HasAnticheatRequired() bool`

HasAnticheatRequired returns a boolean if a field has been set.

### GetCalculateElo

`func (o *Queue) GetCalculateElo() bool`

GetCalculateElo returns the CalculateElo field if non-nil, zero value otherwise.

### GetCalculateEloOk

`func (o *Queue) GetCalculateEloOk() (*bool, bool)`

GetCalculateEloOk returns a tuple with the CalculateElo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculateElo

`func (o *Queue) SetCalculateElo(v bool)`

SetCalculateElo sets CalculateElo field to given value.

### HasCalculateElo

`func (o *Queue) HasCalculateElo() bool`

HasCalculateElo returns a boolean if a field has been set.

### GetCaptainSelection

`func (o *Queue) GetCaptainSelection() string`

GetCaptainSelection returns the CaptainSelection field if non-nil, zero value otherwise.

### GetCaptainSelectionOk

`func (o *Queue) GetCaptainSelectionOk() (*string, bool)`

GetCaptainSelectionOk returns a tuple with the CaptainSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptainSelection

`func (o *Queue) SetCaptainSelection(v string)`

SetCaptainSelection sets CaptainSelection field to given value.

### HasCaptainSelection

`func (o *Queue) HasCaptainSelection() bool`

HasCaptainSelection returns a boolean if a field has been set.

### GetCheckIn

`func (o *Queue) GetCheckIn() CheckIn`

GetCheckIn returns the CheckIn field if non-nil, zero value otherwise.

### GetCheckInOk

`func (o *Queue) GetCheckInOk() (*CheckIn, bool)`

GetCheckInOk returns a tuple with the CheckIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckIn

`func (o *Queue) SetCheckIn(v CheckIn)`

SetCheckIn sets CheckIn field to given value.

### HasCheckIn

`func (o *Queue) HasCheckIn() bool`

HasCheckIn returns a boolean if a field has been set.

### GetEntityId

`func (o *Queue) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *Queue) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *Queue) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *Queue) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetEntityType

`func (o *Queue) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *Queue) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *Queue) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *Queue) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetFbiManagement

`func (o *Queue) GetFbiManagement() bool`

GetFbiManagement returns the FbiManagement field if non-nil, zero value otherwise.

### GetFbiManagementOk

`func (o *Queue) GetFbiManagementOk() (*bool, bool)`

GetFbiManagementOk returns a tuple with the FbiManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFbiManagement

`func (o *Queue) SetFbiManagement(v bool)`

SetFbiManagement sets FbiManagement field to given value.

### HasFbiManagement

`func (o *Queue) HasFbiManagement() bool`

HasFbiManagement returns a boolean if a field has been set.

### GetGame

`func (o *Queue) GetGame() string`

GetGame returns the Game field if non-nil, zero value otherwise.

### GetGameOk

`func (o *Queue) GetGameOk() (*string, bool)`

GetGameOk returns a tuple with the Game field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGame

`func (o *Queue) SetGame(v string)`

SetGame sets Game field to given value.

### HasGame

`func (o *Queue) HasGame() bool`

HasGame returns a boolean if a field has been set.

### GetGroupSimilar

`func (o *Queue) GetGroupSimilar() bool`

GetGroupSimilar returns the GroupSimilar field if non-nil, zero value otherwise.

### GetGroupSimilarOk

`func (o *Queue) GetGroupSimilarOk() (*bool, bool)`

GetGroupSimilarOk returns a tuple with the GroupSimilar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSimilar

`func (o *Queue) SetGroupSimilar(v bool)`

SetGroupSimilar sets GroupSimilar field to given value.

### HasGroupSimilar

`func (o *Queue) HasGroupSimilar() bool`

HasGroupSimilar returns a boolean if a field has been set.

### GetId

`func (o *Queue) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Queue) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Queue) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Queue) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJoinType

`func (o *Queue) GetJoinType() JoinType`

GetJoinType returns the JoinType field if non-nil, zero value otherwise.

### GetJoinTypeOk

`func (o *Queue) GetJoinTypeOk() (*JoinType, bool)`

GetJoinTypeOk returns a tuple with the JoinType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinType

`func (o *Queue) SetJoinType(v JoinType)`

SetJoinType sets JoinType field to given value.

### HasJoinType

`func (o *Queue) HasJoinType() bool`

HasJoinType returns a boolean if a field has been set.

### GetLastModified

`func (o *Queue) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *Queue) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *Queue) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *Queue) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetMaxSkill

`func (o *Queue) GetMaxSkill() int64`

GetMaxSkill returns the MaxSkill field if non-nil, zero value otherwise.

### GetMaxSkillOk

`func (o *Queue) GetMaxSkillOk() (*int64, bool)`

GetMaxSkillOk returns a tuple with the MaxSkill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSkill

`func (o *Queue) SetMaxSkill(v int64)`

SetMaxSkill sets MaxSkill field to given value.

### HasMaxSkill

`func (o *Queue) HasMaxSkill() bool`

HasMaxSkill returns a boolean if a field has been set.

### GetMinSkill

`func (o *Queue) GetMinSkill() int64`

GetMinSkill returns the MinSkill field if non-nil, zero value otherwise.

### GetMinSkillOk

`func (o *Queue) GetMinSkillOk() (*int64, bool)`

GetMinSkillOk returns a tuple with the MinSkill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSkill

`func (o *Queue) SetMinSkill(v int64)`

SetMinSkill sets MinSkill field to given value.

### HasMinSkill

`func (o *Queue) HasMinSkill() bool`

HasMinSkill returns a boolean if a field has been set.

### GetNoOfPlayers

`func (o *Queue) GetNoOfPlayers() int64`

GetNoOfPlayers returns the NoOfPlayers field if non-nil, zero value otherwise.

### GetNoOfPlayersOk

`func (o *Queue) GetNoOfPlayersOk() (*int64, bool)`

GetNoOfPlayersOk returns a tuple with the NoOfPlayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfPlayers

`func (o *Queue) SetNoOfPlayers(v int64)`

SetNoOfPlayers sets NoOfPlayers field to given value.

### HasNoOfPlayers

`func (o *Queue) HasNoOfPlayers() bool`

HasNoOfPlayers returns a boolean if a field has been set.

### GetOpen

`func (o *Queue) GetOpen() bool`

GetOpen returns the Open field if non-nil, zero value otherwise.

### GetOpenOk

`func (o *Queue) GetOpenOk() (*bool, bool)`

GetOpenOk returns a tuple with the Open field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen

`func (o *Queue) SetOpen(v bool)`

SetOpen sets Open field to given value.

### HasOpen

`func (o *Queue) HasOpen() bool`

HasOpen returns a boolean if a field has been set.

### GetOrganizerId

`func (o *Queue) GetOrganizerId() string`

GetOrganizerId returns the OrganizerId field if non-nil, zero value otherwise.

### GetOrganizerIdOk

`func (o *Queue) GetOrganizerIdOk() (*string, bool)`

GetOrganizerIdOk returns a tuple with the OrganizerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizerId

`func (o *Queue) SetOrganizerId(v string)`

SetOrganizerId sets OrganizerId field to given value.

### HasOrganizerId

`func (o *Queue) HasOrganizerId() bool`

HasOrganizerId returns a boolean if a field has been set.

### GetPaused

`func (o *Queue) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *Queue) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *Queue) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *Queue) HasPaused() bool`

HasPaused returns a boolean if a field has been set.

### GetQueueAlgorithm

`func (o *Queue) GetQueueAlgorithm() QueueAlgorithm`

GetQueueAlgorithm returns the QueueAlgorithm field if non-nil, zero value otherwise.

### GetQueueAlgorithmOk

`func (o *Queue) GetQueueAlgorithmOk() (*QueueAlgorithm, bool)`

GetQueueAlgorithmOk returns a tuple with the QueueAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueAlgorithm

`func (o *Queue) SetQueueAlgorithm(v QueueAlgorithm)`

SetQueueAlgorithm sets QueueAlgorithm field to given value.

### HasQueueAlgorithm

`func (o *Queue) HasQueueAlgorithm() bool`

HasQueueAlgorithm returns a boolean if a field has been set.

### GetQueueName

`func (o *Queue) GetQueueName() string`

GetQueueName returns the QueueName field if non-nil, zero value otherwise.

### GetQueueNameOk

`func (o *Queue) GetQueueNameOk() (*string, bool)`

GetQueueNameOk returns a tuple with the QueueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueName

`func (o *Queue) SetQueueName(v string)`

SetQueueName sets QueueName field to given value.

### HasQueueName

`func (o *Queue) HasQueueName() bool`

HasQueueName returns a boolean if a field has been set.

### GetRegion

`func (o *Queue) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Queue) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Queue) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Queue) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetState

`func (o *Queue) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Queue) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Queue) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Queue) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVerifiedMatching

`func (o *Queue) GetVerifiedMatching() bool`

GetVerifiedMatching returns the VerifiedMatching field if non-nil, zero value otherwise.

### GetVerifiedMatchingOk

`func (o *Queue) GetVerifiedMatchingOk() (*bool, bool)`

GetVerifiedMatchingOk returns a tuple with the VerifiedMatching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedMatching

`func (o *Queue) SetVerifiedMatching(v bool)`

SetVerifiedMatching sets VerifiedMatching field to given value.

### HasVerifiedMatching

`func (o *Queue) HasVerifiedMatching() bool`

HasVerifiedMatching returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


