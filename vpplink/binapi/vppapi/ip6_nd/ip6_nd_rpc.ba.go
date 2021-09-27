// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package ip6_nd

import (
	"context"
	"fmt"
	"io"

	api "git.fd.io/govpp.git/api"
	vlib "github.com/projectcalico/vpp-dataplane/vpplink/binapi/vppapi/vlib"
)

// RPCService defines RPC service ip6_nd.
type RPCService interface {
	IP6ndProxyAddDel(ctx context.Context, in *IP6ndProxyAddDel) (*IP6ndProxyAddDelReply, error)
	IP6ndProxyDump(ctx context.Context, in *IP6ndProxyDump) (RPCService_IP6ndProxyDumpClient, error)
	IP6ndProxyEnableDisable(ctx context.Context, in *IP6ndProxyEnableDisable) (*IP6ndProxyEnableDisableReply, error)
	IP6ndSendRouterSolicitation(ctx context.Context, in *IP6ndSendRouterSolicitation) (*IP6ndSendRouterSolicitationReply, error)
	SwInterfaceIP6ndRaConfig(ctx context.Context, in *SwInterfaceIP6ndRaConfig) (*SwInterfaceIP6ndRaConfigReply, error)
	SwInterfaceIP6ndRaPrefix(ctx context.Context, in *SwInterfaceIP6ndRaPrefix) (*SwInterfaceIP6ndRaPrefixReply, error)
	WantIP6RaEvents(ctx context.Context, in *WantIP6RaEvents) (*WantIP6RaEventsReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) IP6ndProxyAddDel(ctx context.Context, in *IP6ndProxyAddDel) (*IP6ndProxyAddDelReply, error) {
	out := new(IP6ndProxyAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) IP6ndProxyDump(ctx context.Context, in *IP6ndProxyDump) (RPCService_IP6ndProxyDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_IP6ndProxyDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vlib.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_IP6ndProxyDumpClient interface {
	Recv() (*IP6ndProxyDetails, error)
	api.Stream
}

type serviceClient_IP6ndProxyDumpClient struct {
	api.Stream
}

func (c *serviceClient_IP6ndProxyDumpClient) Recv() (*IP6ndProxyDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *IP6ndProxyDetails:
		return m, nil
	case *vlib.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) IP6ndProxyEnableDisable(ctx context.Context, in *IP6ndProxyEnableDisable) (*IP6ndProxyEnableDisableReply, error) {
	out := new(IP6ndProxyEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) IP6ndSendRouterSolicitation(ctx context.Context, in *IP6ndSendRouterSolicitation) (*IP6ndSendRouterSolicitationReply, error) {
	out := new(IP6ndSendRouterSolicitationReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SwInterfaceIP6ndRaConfig(ctx context.Context, in *SwInterfaceIP6ndRaConfig) (*SwInterfaceIP6ndRaConfigReply, error) {
	out := new(SwInterfaceIP6ndRaConfigReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SwInterfaceIP6ndRaPrefix(ctx context.Context, in *SwInterfaceIP6ndRaPrefix) (*SwInterfaceIP6ndRaPrefixReply, error) {
	out := new(SwInterfaceIP6ndRaPrefixReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) WantIP6RaEvents(ctx context.Context, in *WantIP6RaEvents) (*WantIP6RaEventsReply, error) {
	out := new(WantIP6RaEventsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}
