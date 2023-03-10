// Code generated by vpplink DO NOT EDIT.
// Copyright (C) 2020 Cisco Systems Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package generated

import (
	"fmt"
	"io"
	"net"

	types "github.com/calico-vpp/vpplink/api/v0"
	"github.com/projectcalico/vpp-dataplane/v3/vpplink/generated/bindings/interface_types"
	"github.com/projectcalico/vpp-dataplane/v3/vpplink/generated/bindings/ip_types"
	"github.com/projectcalico/vpp-dataplane/v3/vpplink/generated/bindings/wireguard"
)

func (v *Vpp) ListWireguardTunnels() ([]*types.WireguardTunnel, error) {
	tunnels, err := v.listWireguardTunnels(interface_types.InterfaceIndex(types.InvalidInterface))
	return tunnels, err
}

func (v *Vpp) GetWireguardTunnel(swIfIndex uint32) (*types.WireguardTunnel, error) {
	tunnels, err := v.listWireguardTunnels(interface_types.InterfaceIndex(swIfIndex))
	if err != nil {
		return nil, err
	}
	if len(tunnels) != 1 {
		return nil, fmt.Errorf("found %d Wireguard tunnels for swIfIndex %d", len(tunnels), swIfIndex)
	}
	return tunnels[0], nil
}

func (v *Vpp) listWireguardTunnels(swIfIndex interface_types.InterfaceIndex) ([]*types.WireguardTunnel, error) {
	client := wireguard.NewServiceClient(v.conn)

	stream, err := client.WireguardInterfaceDump(v.ctx, &wireguard.WireguardInterfaceDump{
		ShowPrivateKey: false,
		SwIfIndex:      swIfIndex,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list Wireguard tunnels: %w", err)
	}
	var tunnels []*types.WireguardTunnel
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to list Wireguard tunnels: %w", err)
		}
		tunnels = append(tunnels, &types.WireguardTunnel{
			Port:      response.Interface.Port,
			Addr:      response.Interface.SrcIP.ToIP(),
			SwIfIndex: uint32(response.Interface.SwIfIndex),
			PublicKey: response.Interface.PublicKey,
		})
	}
	return tunnels, nil
}

func (v *Vpp) AddWireguardTunnel(tunnel *types.WireguardTunnel, generateKey bool) (uint32, error) {
	client := wireguard.NewServiceClient(v.conn)

	response, err := client.WireguardInterfaceCreate(v.ctx, &wireguard.WireguardInterfaceCreate{
		GenerateKey: generateKey,
		Interface: wireguard.WireguardInterface{
			UserInstance: ^uint32(0),
			SwIfIndex:    interface_types.InterfaceIndex(types.InvalidInterface),
			Port:         tunnel.Port,
			SrcIP:        ip_types.NewAddress(tunnel.Addr),
			PrivateKey:   tunnel.PrivateKey,
			PublicKey:    tunnel.PublicKey,
		},
	})
	if err != nil {
		return InvalidSwIfIndex, fmt.Errorf("failed to add Wireguard Tunnel: %w", err)
	}
	tunnel.SwIfIndex = uint32(response.SwIfIndex)
	return uint32(response.SwIfIndex), nil
}

func (v *Vpp) DelWireguardTunnel(tunnel *types.WireguardTunnel) error {
	client := wireguard.NewServiceClient(v.conn)

	_, err := client.WireguardInterfaceDelete(v.ctx, &wireguard.WireguardInterfaceDelete{
		SwIfIndex: interface_types.InterfaceIndex(tunnel.SwIfIndex),
	})
	if err != nil {
		return fmt.Errorf("failed to delete Wireguard Tunnel %s: %w", tunnel, err)
	}
	return nil
}

func (v *Vpp) ListWireguardPeers() ([]*types.WireguardPeer, error) {
	client := wireguard.NewServiceClient(v.conn)

	stream, err := client.WireguardPeersDump(v.ctx, &wireguard.WireguardPeersDump{})
	if err != nil {
		return nil, fmt.Errorf("failed to list Wireguard peers: %w", err)
	}
	var tunnels []*types.WireguardPeer
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to list Wireguard peers: %w", err)
		}
		allowedIps := make([]net.IPNet, 0)
		for _, aip := range response.Peer.AllowedIps {
			allowedIps = append(allowedIps, *aip.ToIPNet())
		}
		tunnels = append(tunnels, &types.WireguardPeer{
			Port:                response.Peer.Port,
			PersistentKeepalive: int(response.Peer.PersistentKeepalive),
			TableID:             response.Peer.TableID,
			Addr:                response.Peer.Endpoint.ToIP(),
			SwIfIndex:           uint32(response.Peer.SwIfIndex),
			PublicKey:           response.Peer.PublicKey,
			AllowedIps:          allowedIps,
		})
	}
	return tunnels, nil
}

func (v *Vpp) AddWireguardPeer(peer *types.WireguardPeer) (uint32, error) {
	allowedIps := make([]ip_types.Prefix, 0)
	for _, aip := range peer.AllowedIps {
		allowedIps = append(allowedIps, ToVppPrefix(&aip))
	}
	ka := uint16(peer.PersistentKeepalive)
	if ka == 0 {
		ka = 1 /* default to 1 */
	}

	client := wireguard.NewServiceClient(v.conn)

	response, err := client.WireguardPeerAdd(v.ctx, &wireguard.WireguardPeerAdd{
		Peer: wireguard.WireguardPeer{
			PublicKey:           peer.PublicKey,
			Port:                peer.Port,
			PersistentKeepalive: ka,
			TableID:             peer.TableID,
			Endpoint:            ip_types.NewAddress(peer.Addr),
			SwIfIndex:           interface_types.InterfaceIndex(peer.SwIfIndex),
			AllowedIps:          allowedIps,
		},
	})
	if err != nil {
		return InvalidSwIfIndex, fmt.Errorf("failed to add Wireguard Peer: %w", err)
	}
	peer.Index = uint32(response.PeerIndex)
	return uint32(response.PeerIndex), nil
}

func (v *Vpp) DelWireguardPeer(peer *types.WireguardPeer) error {
	client := wireguard.NewServiceClient(v.conn)

	_, err := client.WireguardPeerRemove(v.ctx, &wireguard.WireguardPeerRemove{
		PeerIndex: uint32(peer.Index),
	})
	if err != nil {
		return fmt.Errorf("failed to delete Wireguard Peer %s: %w", peer, err)
	}
	return nil
}