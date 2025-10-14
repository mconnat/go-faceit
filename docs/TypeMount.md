# TypeMount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FsType** | Pointer to **string** | FsType specifies the filesystem type for the mount volume. Optional. | [optional] 
**MountFlags** | Pointer to **[]string** | MountFlags defines flags to pass when mounting the volume. Optional. | [optional] 

## Methods

### NewTypeMount

`func NewTypeMount() *TypeMount`

NewTypeMount instantiates a new TypeMount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypeMountWithDefaults

`func NewTypeMountWithDefaults() *TypeMount`

NewTypeMountWithDefaults instantiates a new TypeMount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFsType

`func (o *TypeMount) GetFsType() string`

GetFsType returns the FsType field if non-nil, zero value otherwise.

### GetFsTypeOk

`func (o *TypeMount) GetFsTypeOk() (*string, bool)`

GetFsTypeOk returns a tuple with the FsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsType

`func (o *TypeMount) SetFsType(v string)`

SetFsType sets FsType field to given value.

### HasFsType

`func (o *TypeMount) HasFsType() bool`

HasFsType returns a boolean if a field has been set.

### GetMountFlags

`func (o *TypeMount) GetMountFlags() []string`

GetMountFlags returns the MountFlags field if non-nil, zero value otherwise.

### GetMountFlagsOk

`func (o *TypeMount) GetMountFlagsOk() (*[]string, bool)`

GetMountFlagsOk returns a tuple with the MountFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountFlags

`func (o *TypeMount) SetMountFlags(v []string)`

SetMountFlags sets MountFlags field to given value.

### HasMountFlags

`func (o *TypeMount) HasMountFlags() bool`

HasMountFlags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


