package main

import (
	"net/rpc"

	plugin "github.com/dkiser/go-plugin-example/plugin"
	gplugin "github.com/hashicorp/go-plugin"
)

// Here is a real implementation of Greeter
type GreeterHello struct{}

func (GreeterHello) Greet() string { return "Hello!" }

// This is the implementation of plugin.Plugin so we can serve/consume this
//
// This has two methods: Server must return an RPC server for this plugin
// type. We construct a GreeterRPCServer for this.
//
// Client must return an implementation of our interface that communicates
// over an RPC client. We return GreeterRPC for this.
//
// Ignore MuxBroker. That is used to create more multiplexed streams on our
// plugin connection and is a more advanced use case.
type GreeterPlugin struct{}

func (GreeterPlugin) Server(*gplugin.MuxBroker) (interface{}, error) {
	return &plugin.GreeterRPCServer{Impl: new(GreeterHello)}, nil
}

func (GreeterPlugin) Client(b *gplugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &plugin.GreeterRPC{Client: c}, nil
}

func main() {
	// We're a plugin! Serve the plugin. We set the handshake config
	// so that the host and our plugin can verify they can talk to each other.
	// Then we set the plugin map to say what plugins we're serving.
	gplugin.Serve(&gplugin.ServeConfig{
		HandshakeConfig: plugin.HandshakeConfig,
		Plugins:         pluginMap,
	})
}

// pluginMap is the map of plugins we can dispense.
var pluginMap = map[string]gplugin.Plugin{
	"hello": new(GreeterPlugin),
}
