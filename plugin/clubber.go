package plugin

import "net/rpc"
import gplugin "github.com/hashicorp/go-plugin"

// Clubber is the interface that we're exposing as a plugin.
type Clubber interface {
	FistPump() string
}

// Here is an implementation that talks over RPC
type ClubberRPC struct {
	Client *rpc.Client
}

func (g *ClubberRPC) FistPump() string {
	var resp string
	err := g.Client.Call("Plugin.FistPump", new(interface{}), &resp)
	if err != nil {
		// You usually want your interfaces to return errors. If they don't,
		// there isn't much other choice here.
		panic(err)
	}

	return resp
}

// Here is the RPC server that ClubberRPC talks to, conforming to
// the requirements of net/rpc
type ClubberRPCServer struct {
	// This is the real implementation
	Impl Clubber
}

func (s *ClubberRPCServer) FistPump(args interface{}, resp *string) error {
	*resp = s.Impl.FistPump()
	return nil
}

// Dummy implementation of a plugin.Plugin interface for use in PluginMap.
// At runtime, a real implementation from a plugin implementation overwrides
// this.
type ClubberPlugin struct{}

func (ClubberPlugin) Server(*gplugin.MuxBroker) (interface{}, error) {
	return &ClubberRPCServer{}, nil
}

func (ClubberPlugin) Client(b *gplugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &ClubberRPC{Client: c}, nil
}
