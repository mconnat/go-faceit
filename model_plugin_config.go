/*
 * Data API
 *
 * # This API provide access to FACEIT's data
 *
 * API version: 4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package faceit

type PluginConfigLinux struct {
	// allow all devices
	AllowAllDevices bool `json:"AllowAllDevices"`
	// capabilities
	Capabilities []string `json:"Capabilities"`
	// devices
	Devices []PluginDevice `json:"Devices"`
}
type PluginConfig struct {
	Args *PluginConfigArgs `json:"Args"`
	// description
	Description string `json:"Description"`
	// Docker Version used to create the plugin
	DockerVersion string `json:"DockerVersion,omitempty"`
	// documentation
	Documentation string `json:"Documentation"`
	// entrypoint
	Entrypoint []string `json:"Entrypoint"`
	// env
	Env        []PluginEnv            `json:"Env"`
	Interface_ *PluginConfigInterface `json:"Interface"`
	// ipc host
	IpcHost bool               `json:"IpcHost"`
	Linux   *PluginConfigLinux `json:"Linux"`
	// mounts
	Mounts  []PluginMount        `json:"Mounts"`
	Network *PluginConfigNetwork `json:"Network"`
	// pid host
	PidHost bool `json:"PidHost"`
	// propagated mount
	PropagatedMount string            `json:"PropagatedMount"`
	User            *PluginConfigUser `json:"User,omitempty"`
	// work dir
	WorkDir string              `json:"WorkDir"`
	Rootfs  *PluginConfigRootfs `json:"rootfs,omitempty"`
}
