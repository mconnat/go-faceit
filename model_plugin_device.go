/*
 * Data API
 *
 * # This API provide access to FACEIT's data
 *
 * API version: 4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package go-faceit

// PluginDevice plugin device
type PluginDevice struct {
	// description
	Description string `json:"Description"`
	// name
	Name string `json:"Name"`
	// path
	Path string `json:"Path"`
	// settable
	Settable []string `json:"Settable"`
}
