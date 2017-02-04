package plugin

import (
	plugin "github.com/hashicorp/go-plugin"
)

// <plugin type>-<plugin id/name>
// e.g. greeter-foo would be a plugin of type "greeter", with an id of "foo"

// handshakeConfigs are used to just do a basic handshake between
// a plugin and host. If the handshake fails, a user friendly error is shown.
// This prevents users from executing bad plugins or executing a plugin
// directory. It is a UX feature, not a security feature.
var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

// // PluginMap should be used by clients for the map of plugins.
// // This maps types of plugins to interface implementations.
// var PluginMap = map[string]plugin.Plugin{
// 	"foo":   &GreeterPlugin{},
// 	"hello": &GreeterPlugin{},
// }
