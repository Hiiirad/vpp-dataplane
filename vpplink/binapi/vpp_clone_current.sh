#!/bin/bash
set -e

function git_cherry_pick ()
{
	refs=$1
	git fetch "https://gerrit.fd.io/r/vpp" ${refs}
	git cherry-pick FETCH_HEAD
	git commit --amend -m "gerrit:${refs#refs/changes/*/} $(git log -1 --pretty=%B)"
}

if [ -z "$1" ]; then
	echo "Missing VPP path"
	exit 1
fi

VPP_COMMIT=078d258034cef5b4ca74d9deb37b2684cc77d060
VPP_DIR="$1"

if [ ! -d $1/.git ]; then
	rm -rf $1
	git clone "https://gerrit.fd.io/r/vpp" $1
	cd $1
	git reset --hard ${VPP_COMMIT}
else
	cd $1
	git fetch "https://gerrit.fd.io/r/vpp" && git reset --hard ${VPP_COMMIT}
fi

git_cherry_pick refs/changes/68/32468/3 # 32468: buffers: fix buffer linearization | https://gerrit.fd.io/r/c/vpp/+/32468
git_cherry_pick refs/changes/33/32833/1 # 32833: ipsec: disable linearization | https://gerrit.fd.io/r/c/vpp/+/32833
git_cherry_pick refs/changes/86/29386/9 # 29386: virtio: DRAFT: multi tx support | https://gerrit.fd.io/r/c/vpp/+/29386
git_cherry_pick refs/changes/21/31321/11 # 31321: devices: add support for pseudo header checksum | https://gerrit.fd.io/r/c/vpp/+/31321

# Revert for now
git fetch "https://gerrit.fd.io/r/vpp" refs/changes/68/32568/6 && git revert --no-edit FETCH_HEAD
git_cherry_pick refs/changes/69/31869/13 # 31869: gso: do not try gro on small packets | https://gerrit.fd.io/r/c/vpp/+/31869
git_cherry_pick refs/changes/76/32476/3 # 32476: gso: handle push flag in gro | https://gerrit.fd.io/r/c/vpp/+/32476

git_cherry_pick refs/changes/82/32482/1 # 32482: virtio: compute cksums in output no offload | https://gerrit.fd.io/r/c/vpp/+/32482
git_cherry_pick refs/changes/83/32483/1 # 32483: virtio: Still init unused txq | https://gerrit.fd.io/r/c/vpp/+/32483

git_cherry_pick refs/changes/71/32871/1 # 32871: devices: Add queues params in create_if | https://gerrit.fd.io/r/c/vpp/+/32871

# IPv6 ND patch (temporary)
git_cherry_pick refs/changes/68/31868/1 # 31868: ip6-nd: silent the source and target checks on given interface | https://gerrit.fd.io/r/c/vpp/+/31868

# --------------- Cnat patches ---------------
git_cherry_pick refs/changes/88/31588/1 # 31588: cnat: [WIP] no k8s maglev from pods | https://gerrit.fd.io/r/c/vpp/+/31588
# --------------- Cnat patches ---------------

# ------------- Policies patches -------------
git_cherry_pick refs/changes/83/28083/20 # 28083: acl: acl-plugin custom policies |  https://gerrit.fd.io/r/c/vpp/+/28083
git_cherry_pick refs/changes/13/28513/24 # 25813: capo: Calico Policies plugin | https://gerrit.fd.io/r/c/vpp/+/28513
# ------------- Policies patches -------------

git_cherry_pick refs/changes/64/33164/3  # 33164: dpdk: enable ena interrupt support in dpdk version 21.05 | https://gerrit.fd.io/r/c/vpp/+/33164

# -------------------- NSM -------------------
git_cherry_pick refs/changes/49/33449/9  # 33449 ip: source address selection
git_cherry_pick refs/changes/95/33495/2  # 33495 arp: source address selection
git_cherry_pick refs/changes/13/32413/6  # 32413 wireguard: move adjacency processing from wireguard_peer to wireguard_interface
git_cherry_pick refs/changes/43/32443/7  # 32443 wireguard: use the same udp-port for multi-tunnel
git_cherry_pick refs/changes/09/32009/5  # 32009 wireguard: add ipv6 support
git_cherry_pick refs/changes/85/32685/4  # 32685 wireguard: add events for peer
git_cherry_pick refs/changes/68/33568/1  # 33568 ip: check if interface has link-local address (addition)
git_cherry_pick refs/changes/58/33558/1  # 33558 ip: check if interface has link-local address
# -------------------- NSM -------------------

git_cherry_pick refs/changes/71/32271/6  # 32271 memif: add support for ns abstract sockets
