---
layout: default
title: backend.conf
---

# {{page.title}}
Configuration of the FNews service news storage backend

## Configuring the BucketStore-Access

In order for FNews to work, we need a BucketStore. To connect to it, we specify a `bucket` Entry.

```
bucket {
	address: localhost:62092
	proto: kcphttp
	data-shards: 10
	parity-shards: 3
	turbo: true
}
```

A `bucket` Entry has the following keys:

* __*proto*__

  The Protocol that is used for communication. Must be one of __*http*__, __*oohttp*__ or __*kcphttp*__.
  With __*http*__, the server implements Plain HTTP/1.1. With __*oohttp*__, the server implements an
  out-of-order protocol, that encapsulates HTTP/1.1 requests. With __*kcphttp*__, the server implements
  an out-of-order protocol, that encapsulates HTTP/1.1 requests and is built upon KCP, which in turn
  is built upon UDP, thus offering superior performance.

* __*address*__

  The `<host>:<port>` Pair of the Bucket-Server.

* __*data-shards*__

  The number of data shards in KCP's reed-solomon encoding. Only useful if __*proto*__ is __*kcphttp*__.

* __*parity-shards*__

  The number of parity shards in KCP's reed-solomon encoding. Only useful if __*proto*__ is __*kcphttp*__.

* __*turbo*__

  If __*true*__, KCP is switched into, what i call, Turbo-Mode, which reduces KCPs CPU-overhead.

## Configuring the Database

FNews requires a Database to function. For this, a  `db` Entry.

```
db {
	driver: postgres
	data-source: "user=usr password=pwd dbname=mydatabase sslmode=disable"
	schema: v2
}
```

The `db`Entry has the following keys:

* __*driver*__

  The database driver, to connect to the database. Currently, only __*postgres*__ (PostgreSQL) is supported.

* __*data-source*__

  The database connection string.
  The format for connection strings for the __*postgres*__ driver: [lib/pq](https://godoc.org/github.com/lib/pq)

* __*schema*__

  The domain-model/table-model to be instantiated and to be used. Must be __*v1*__ or __*v2*__.
  __*v2*__ is default and __*v1*__ is deprecated.

