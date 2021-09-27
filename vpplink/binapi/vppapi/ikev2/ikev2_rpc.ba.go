// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package ikev2

import (
	"context"
	"fmt"
	"io"

	api "git.fd.io/govpp.git/api"
	vlib "github.com/projectcalico/vpp-dataplane/vpplink/binapi/vppapi/vlib"
)

// RPCService defines RPC service ikev2.
type RPCService interface {
	Ikev2ChildSaDump(ctx context.Context, in *Ikev2ChildSaDump) (RPCService_Ikev2ChildSaDumpClient, error)
	Ikev2InitiateDelChildSa(ctx context.Context, in *Ikev2InitiateDelChildSa) (*Ikev2InitiateDelChildSaReply, error)
	Ikev2InitiateDelIkeSa(ctx context.Context, in *Ikev2InitiateDelIkeSa) (*Ikev2InitiateDelIkeSaReply, error)
	Ikev2InitiateRekeyChildSa(ctx context.Context, in *Ikev2InitiateRekeyChildSa) (*Ikev2InitiateRekeyChildSaReply, error)
	Ikev2InitiateSaInit(ctx context.Context, in *Ikev2InitiateSaInit) (*Ikev2InitiateSaInitReply, error)
	Ikev2NonceGet(ctx context.Context, in *Ikev2NonceGet) (*Ikev2NonceGetReply, error)
	Ikev2PluginGetVersion(ctx context.Context, in *Ikev2PluginGetVersion) (*Ikev2PluginGetVersionReply, error)
	Ikev2ProfileAddDel(ctx context.Context, in *Ikev2ProfileAddDel) (*Ikev2ProfileAddDelReply, error)
	Ikev2ProfileDisableNatt(ctx context.Context, in *Ikev2ProfileDisableNatt) (*Ikev2ProfileDisableNattReply, error)
	Ikev2ProfileDump(ctx context.Context, in *Ikev2ProfileDump) (RPCService_Ikev2ProfileDumpClient, error)
	Ikev2ProfileSetAuth(ctx context.Context, in *Ikev2ProfileSetAuth) (*Ikev2ProfileSetAuthReply, error)
	Ikev2ProfileSetID(ctx context.Context, in *Ikev2ProfileSetID) (*Ikev2ProfileSetIDReply, error)
	Ikev2ProfileSetIpsecUDPPort(ctx context.Context, in *Ikev2ProfileSetIpsecUDPPort) (*Ikev2ProfileSetIpsecUDPPortReply, error)
	Ikev2ProfileSetLiveness(ctx context.Context, in *Ikev2ProfileSetLiveness) (*Ikev2ProfileSetLivenessReply, error)
	Ikev2ProfileSetTs(ctx context.Context, in *Ikev2ProfileSetTs) (*Ikev2ProfileSetTsReply, error)
	Ikev2ProfileSetUDPEncap(ctx context.Context, in *Ikev2ProfileSetUDPEncap) (*Ikev2ProfileSetUDPEncapReply, error)
	Ikev2SaDump(ctx context.Context, in *Ikev2SaDump) (RPCService_Ikev2SaDumpClient, error)
	Ikev2SetEspTransforms(ctx context.Context, in *Ikev2SetEspTransforms) (*Ikev2SetEspTransformsReply, error)
	Ikev2SetIkeTransforms(ctx context.Context, in *Ikev2SetIkeTransforms) (*Ikev2SetIkeTransformsReply, error)
	Ikev2SetLocalKey(ctx context.Context, in *Ikev2SetLocalKey) (*Ikev2SetLocalKeyReply, error)
	Ikev2SetResponder(ctx context.Context, in *Ikev2SetResponder) (*Ikev2SetResponderReply, error)
	Ikev2SetResponderHostname(ctx context.Context, in *Ikev2SetResponderHostname) (*Ikev2SetResponderHostnameReply, error)
	Ikev2SetSaLifetime(ctx context.Context, in *Ikev2SetSaLifetime) (*Ikev2SetSaLifetimeReply, error)
	Ikev2SetTunnelInterface(ctx context.Context, in *Ikev2SetTunnelInterface) (*Ikev2SetTunnelInterfaceReply, error)
	Ikev2TrafficSelectorDump(ctx context.Context, in *Ikev2TrafficSelectorDump) (RPCService_Ikev2TrafficSelectorDumpClient, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) Ikev2ChildSaDump(ctx context.Context, in *Ikev2ChildSaDump) (RPCService_Ikev2ChildSaDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_Ikev2ChildSaDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vlib.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_Ikev2ChildSaDumpClient interface {
	Recv() (*Ikev2ChildSaDetails, error)
	api.Stream
}

type serviceClient_Ikev2ChildSaDumpClient struct {
	api.Stream
}

func (c *serviceClient_Ikev2ChildSaDumpClient) Recv() (*Ikev2ChildSaDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *Ikev2ChildSaDetails:
		return m, nil
	case *vlib.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) Ikev2InitiateDelChildSa(ctx context.Context, in *Ikev2InitiateDelChildSa) (*Ikev2InitiateDelChildSaReply, error) {
	out := new(Ikev2InitiateDelChildSaReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Ikev2InitiateDelIkeSa(ctx context.Context, in *Ikev2InitiateDelIkeSa) (*Ikev2InitiateDelIkeSaReply, error) {
	out := new(Ikev2InitiateDelIkeSaReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Ikev2InitiateRekeyChildSa(ctx context.Context, in *Ikev2InitiateRekeyChildSa) (*Ikev2InitiateRekeyChildSaReply, error) {
	out := new(Ikev2InitiateRekeyChildSaReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Ikev2InitiateSaInit(ctx context.Context, in *Ikev2InitiateSaInit) (*Ikev2InitiateSaInitReply, error) {
	out := new(Ikev2InitiateSaInitReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Ikev2NonceGet(ctx context.Context, in *Ikev2NonceGet) (*Ikev2NonceGetReply, error) {
	out := new(Ikev2NonceGetReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Ikev2PluginGetVersion(ctx context.Context, in *Ikev2PluginGetVersion) (*Ikev2PluginGetVersionReply, error) {
	out := new(Ikev2PluginGetVersionReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Ikev2ProfileAddDel(ctx context.Context, in *Ikev2ProfileAddDel) (*Ikev2ProfileAddDelReply, error) {
	out := new(Ikev2ProfileAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Ikev2ProfileDisableNatt(ctx context.Context, in *Ikev2ProfileDisableNatt) (*Ikev2ProfileDisableNattReply, error) {
	out := new(Ikev2ProfileDisableNattReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Ikev2ProfileDump(ctx context.Context, in *Ikev2ProfileDump) (RPCService_Ikev2ProfileDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_Ikev2ProfileDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vlib.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_Ikev2ProfileDumpClient interface {
	Recv() (*Ikev2ProfileDetails, error)
	api.Stream
}

type serviceClient_Ikev2ProfileDumpClient struct {
	api.Stream
}

func (c *serviceClient_Ikev2ProfileDumpClient) Recv() (*Ikev2ProfileDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *Ikev2ProfileDetails:
		return m, nil
	case *vlib.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) Ikev2ProfileSetAuth(ctx context.Context, in *Ikev2ProfileSetAuth) (*Ikev2ProfileSetAuthReply, error) {
	out := new(Ikev2ProfileSetAuthReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Ikev2ProfileSetID(ctx context.Context, in *Ikev2ProfileSetID) (*Ikev2ProfileSetIDReply, error) {
	out := new(Ikev2ProfileSetIDReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Ikev2ProfileSetIpsecUDPPort(ctx context.Context, in *Ikev2ProfileSetIpsecUDPPort) (*Ikev2ProfileSetIpsecUDPPortReply, error) {
	out := new(Ikev2ProfileSetIpsecUDPPortReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Ikev2ProfileSetLiveness(ctx context.Context, in *Ikev2ProfileSetLiveness) (*Ikev2ProfileSetLivenessReply, error) {
	out := new(Ikev2ProfileSetLivenessReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Ikev2ProfileSetTs(ctx context.Context, in *Ikev2ProfileSetTs) (*Ikev2ProfileSetTsReply, error) {
	out := new(Ikev2ProfileSetTsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Ikev2ProfileSetUDPEncap(ctx context.Context, in *Ikev2ProfileSetUDPEncap) (*Ikev2ProfileSetUDPEncapReply, error) {
	out := new(Ikev2ProfileSetUDPEncapReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Ikev2SaDump(ctx context.Context, in *Ikev2SaDump) (RPCService_Ikev2SaDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_Ikev2SaDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vlib.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_Ikev2SaDumpClient interface {
	Recv() (*Ikev2SaDetails, error)
	api.Stream
}

type serviceClient_Ikev2SaDumpClient struct {
	api.Stream
}

func (c *serviceClient_Ikev2SaDumpClient) Recv() (*Ikev2SaDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *Ikev2SaDetails:
		return m, nil
	case *vlib.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) Ikev2SetEspTransforms(ctx context.Context, in *Ikev2SetEspTransforms) (*Ikev2SetEspTransformsReply, error) {
	out := new(Ikev2SetEspTransformsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Ikev2SetIkeTransforms(ctx context.Context, in *Ikev2SetIkeTransforms) (*Ikev2SetIkeTransformsReply, error) {
	out := new(Ikev2SetIkeTransformsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Ikev2SetLocalKey(ctx context.Context, in *Ikev2SetLocalKey) (*Ikev2SetLocalKeyReply, error) {
	out := new(Ikev2SetLocalKeyReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Ikev2SetResponder(ctx context.Context, in *Ikev2SetResponder) (*Ikev2SetResponderReply, error) {
	out := new(Ikev2SetResponderReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Ikev2SetResponderHostname(ctx context.Context, in *Ikev2SetResponderHostname) (*Ikev2SetResponderHostnameReply, error) {
	out := new(Ikev2SetResponderHostnameReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Ikev2SetSaLifetime(ctx context.Context, in *Ikev2SetSaLifetime) (*Ikev2SetSaLifetimeReply, error) {
	out := new(Ikev2SetSaLifetimeReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Ikev2SetTunnelInterface(ctx context.Context, in *Ikev2SetTunnelInterface) (*Ikev2SetTunnelInterfaceReply, error) {
	out := new(Ikev2SetTunnelInterfaceReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Ikev2TrafficSelectorDump(ctx context.Context, in *Ikev2TrafficSelectorDump) (RPCService_Ikev2TrafficSelectorDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_Ikev2TrafficSelectorDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vlib.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_Ikev2TrafficSelectorDumpClient interface {
	Recv() (*Ikev2TrafficSelectorDetails, error)
	api.Stream
}

type serviceClient_Ikev2TrafficSelectorDumpClient struct {
	api.Stream
}

func (c *serviceClient_Ikev2TrafficSelectorDumpClient) Recv() (*Ikev2TrafficSelectorDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *Ikev2TrafficSelectorDetails:
		return m, nil
	case *vlib.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}
