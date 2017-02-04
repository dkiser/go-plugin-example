package plugin

import "net/rpc"
import gplugin "github.com/hashicorp/go-plugin"

// Greeter is the interface that we're exposing as a plugin.
type Greeter interface {
	Greet() string
}

// Here is an implementation that talks over RPC
type GreeterRPC struct {
	Client *rpc.Client
}

func (g *GreeterRPC) Greet() string {
	var resp string
	err := g.Client.Call("Plugin.Greet", new(interface{}), &resp)
	if err != nil {
		// You usually want your interfaces to return errors. If they don't,
		// there isn't much other choice here.
		panic(err)
	}

	return resp
}

// Here is the RPC server that GreeterRPC talks to, conforming to
// the requirements of net/rpc
type GreeterRPCServer struct {
	// This is the real implementation
	Impl Greeter
}

func (s *GreeterRPCServer) Greet(args interface{}, resp *string) error {
	*resp = s.Impl.Greet()
	return nil
}

// Dummy implementation of a plugin.Plugin interface for use in PluginMap.
// At runtime, a real implementation from a plugin implementation overwrides
// this.
type GreeterPlugin struct{}

func (GreeterPlugin) Server(*gplugin.MuxBroker) (interface{}, error) {
	return &GreeterRPCServer{}, nil
	//return interface{}, nil
}

func (GreeterPlugin) Client(b *gplugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &GreeterRPC{Client: c}, nil
	//return interface{}, nil
}
