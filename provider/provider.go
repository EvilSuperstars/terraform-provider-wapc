package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-plugin"
	"github.com/EvilSuperstars/terraform-provider-wapc/tfplugin5"
	"google.golang.org/grpc"
)

// Provider entrypoint.
func Serve() {
	handshake := plugin.HandshakeConfig{
		ProtocolVersion: 5,
		// The magic cookie values should NEVER be changed.
		MagicCookieKey:   "TF_PLUGIN_MAGIC_COOKIE",
		MagicCookieValue: "d602bf8f470bc67ca7faa0386276bbdd4330efaf76d1a219cb4d6991ca9872b2",
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshake,
		GRPCServer:      plugin.DefaultGRPCServer,
		Plugins: plugin.PluginSet{
			"provider": &grpcPlugin{
				providerServer: NewGRPCProviderServer(),
			},
		},
	})
}

type grpcPlugin struct {
	plugin.Plugin
	providerServer tfplugin5.ProviderServer
}

//
// GRPCPlugin interface methods.
//

// GRPCServer should register this plugin for serving with the given gRPC server.
func (p *grpcPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	tfplugin5.RegisterProviderServer(s, p.providerServer)

	return nil
}

func (p *grpcPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return nil, fmt.Errorf("This is a plugin - it cannot implement GRPCClient")
}
