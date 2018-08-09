# Implementing a Transparent Sidecar in Network Service Mesh

The current model of inserting sidecars into Kubernetes Pods works as a network proxy, and needs some modifications of IPTables to ensure that the traffic is directed to the sidecar in both ingress and egress directions. In addtion the proxy only inspects traffic that it has protocol handlers for. 

Being a network proxy introduces a degree of network latency that for many applications should be avoided.

An alternative approach is proposed using Network Service Mesh to insert an transparent network proxy into the Pod to replace the connection based sidecar proxy. This approach has the following advantages:

1. Provides visibility into all traffic in both ingress and egress direction. This is important for audiences such as SecOps who generatl want to see all traffic.
2. Provides low latency as connections are not terminated.
3. Acceleration of the packet path with SR-IOV and DPDK can be enabled.
4. Easy to install - no complicated IPTable rules, and works with any applications.

##Approach

The current sidecar deployment model is shown in the diagram below:

```
┌──────────────────────┬──────────┬─────────────────────┐
│                      │  Pod-0   │                     │
│                      └──────────┘                     │
│                                                       │
│   ┌────────┐   ┌─────┐    ┌─────┐                     │
│   │ pause  │   │App1 │    │App2 │                     │
│   └───┬────┘   └──┬──┘    └──┬──┘                     │
│       │           │          │     ┌────────────┐     │
│ ──────┴─────┬─────┴────┬─────┴─────┤ localhost  ├─────│
│             │          │           └────────────┘     │
│             │          │                              │
│         ┌───┴───┐      │                              │
│         │Sidecar│      │                              │
│         └───────┘      │                              │
│                        │                              │
│                     ┌──┴──┐                           │
└─────────────────────┤eth0 ├───────────────────────────┘
                      └─────┘
```

The sidecar and all applications are listening on localhost and/or host IP. Traffic needs to be steered to the Sidecar using IPTable rules. This can be automated in most cases but not all.

An alternative approach is to insert a transparent sidecar that does NOT proxy traffic but rather transparently forwards traffic between two interfaces of a container that implements PACKET_MMAP. This deployment model is shown below:

```
┌─────────────────────┬──────────┬──────────────────────┐
│                     │  Pod-0   │                      │
│                     └──────────┘                      │
│                                                       │
│  ┌────────┐   ┌─────┐    ┌─────┐                      │
│  │ pause  │   │App1 │    │App2 │                      │
│  └───┬────┘   └──┬──┘    └──┬──┘                      │
│      │           │          │     ┌────────────┐      │
│──────┴───────────┴────┬─────┴─────┤ localhost  ├───── │
│                       │           └────────────┘      │
│                       │                               │
│                    ┌──┴──┐                            │
│                  ┌─┤eth0 ├─┐                          │
│                  │ └──┬──┘ │                          │
│                  │    │    │                          │
│                  │ ┌──┴──┐ │                          │
│                  │ │ CNF │ │                          │
│                  │ └──┬──┘ │                          │
│                  │    │    │                          │
│                  │ ┌──┴──┐ │                          │
│                  └─┤veth2├─┘                          │
│                    └──┬──┘                            │
│                       │                               │
│                    ┌──┴──┐                            │
└────────────────────┤veth1├────────────────────────────┘
                     └─────┘
```

Here the eth0 interface is preserved with its IP and MAC Addresses, so Kubernetes works with no changes. There is a new Linux Network Namespace introduced to run the CNF Container and this container is connected by veth interfaces to the existing applications and the network bridge of the Node.

The CNF is implements a PACKET_MMAP interface that simply forwards packets (with no changes) from veth2 to eth0 and vice-versa, this includes and ICMP/ARP traffic, so App1 and App2 are discoverable as normal. There is no need to modify IPTables or any other part of the Pod infrastructure.

##Approach
To insert the new transparent sidecar and redo the wiring for the Pod, access to the Pod is required. The approach taken is to implement a L2 Network Insertion plugin using Network Service Mesh.

Extending the network service mesh to add a new DaemonSet that will insert a L2 Network into the Pod makes the implementation fairly easy. The CNF container is implemented as a Sidecar that is inserted by a NSM DaemonSet that inserts the new data-plane.


