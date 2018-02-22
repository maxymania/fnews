---
layout: default
title: bucketstore.conf
---

# {{page.title}}

Configuration of the BucketStore service

## Defining buckets

For Every Bucket you want to add, you need to add an Entry like this:

```
backend {
	type: dayfilemulti
	path: "G:/Path/To/Spool"
	config {
		MaxSpace: 1000000000000
		MaxFiles: 1000
	}
}
```

The `backend` Entry has the following keys:

* __*type*__

  The storage method, that is used to manage the spool folder. It must be **_dayfile_** or **_dayfilemulti_**.
  The storage method **_dayfilemulti_** is recommended NTFS filesystems larger than 16TB since NTFS has a
  maximum file size of 16 TB.

* __*path*__

  The Path where the spool is located. You can subsititute any `\ ` character with `/`, it will be converted automatically.

The `config` Entry within the `backend` Entry is required and should contain the following keys:

* __*MaxSpace*__

  The Maximum storage space to be consumed (approximately).

* __*MaxFiles*__

  The maximum number of file descriptors generated (approximately).


## Defining Protocol and Port

There are two types of entries to define the network Protocol: `node` Entries and `server` Entries.
There must be __exactly one__ `node` Entry but there can be __zero or more__ `server` Entries.
Both  `node` and `server` Entries have the same format and the same meaning.

The only technical thing, that sets the `node` Entry apart from the `server` Entries is, that the port
specified in the `node` Entry is used for Node-to-Node communication, if the software runs in Clustering mode.

```
node {
	proto: kcphttp
	port: 62092
	data-shards: 10
	parity-shards: 3
	turbo: true
}

server {
	proto: kcphttp
	port: 62093
	data-shards: 10
	parity-shards: 3
	turbo: true
}
```

A `node` or `server` Entry has the following keys:

* __*proto*__

  The Protocol that is used for communication. Must be one of __*http*__, __*oohttp*__ or __*kcphttp*__.
  With __*http*__, the server implements Plain HTTP/1.1. With __*oohttp*__, the server implements an
  out-of-order protocol, that encapsulates HTTP/1.1 requests. With __*kcphttp*__, the server implements
  an out-of-order protocol, that encapsulates HTTP/1.1 requests and is built upon KCP, which in turn
  is built upon UDP, thus offering superior performance.

* __*port*__

  The TCP- or UDP-port to be used.

* __*data-shards*__

  The number of data shards in KCP's reed-solomon encoding. Only useful if __*proto*__ is __*kcphttp*__.

* __*parity-shards*__

  The number of parity shards in KCP's reed-solomon encoding. Only useful if __*proto*__ is __*kcphttp*__.

* __*turbo*__

  If __*true*__, KCP is switched into, what i call, Turbo-Mode, which reduces KCPs CPU-overhead.



## Clustering

The Cluster Support of BucketStore is built upon [Hashicorp's Memberlist library](github.com/hashicorp/memberlist).
If you want to enable clustering, you need to add a `cluster` Entry:

```
cluster {
	specs: LAN
	node-name: MyHost123
	bind-addr: 100.133.28.5
	bind-port: 7946
	advertise-addr: 100.133.28.5
	advertise-port: 7946
	join-ip-prefer: 4
	
	join: 100.133.28.10:7946
	join: 100.133.28.11:7946
	join: 100.133.28.12:7946
	join: foo-host.local:7946
}
```

The `cluster` Entry has the following keys:

* __*specs*__

  Required. The configuration specs for the Memberlist. Must be one of __*LAN*__, __*WAN*__ or __*Local*__.
  Their lower-case counterparts work as well. __*LAN*__ uses the hostname as the node name, and otherwise
  sets very conservative values that are sane for most LAN environments.
  __*WAN*__ works like __*LAN*__ but sets a configuration that is optimized for most WAN environments
  __*Local*__ works like __*LAN*__ as well, but sets a configuration that is optimized for a local
  loopback environments.

* __*node-name*__

  If set, it overrides the node-name.

* __*bind-addr*__

  Overrides the IP-address to bind against. Default is `0.0.0.0`

* __*bind-port*__

  Overrides the TCP- and UDP-Port to bind against. Default is `7946`.

* __*advertise-addr*__

  Configuration related to what IP-address to advertise to other cluster members. Used for nat traversal.

* __*advertise-port*__

  Configuration related to what TCP- and UDP-port to advertise to other cluster members. Used for nat traversal.
  Default is 7946.

* __*join-ip-prefer*__

  Specifies, which IP you prefer when resolving Host names for Cluster joins. Should be __*4*__ or __*6*__.
  Everything else should be

* __*join*__

  Specifies `<host>:<port>` pairs of other cluster Members. Useful to join a cluster. `<host>` is
  either an IP-address or a host name to be resolved with DNS, NETBIOS or other name-services.


