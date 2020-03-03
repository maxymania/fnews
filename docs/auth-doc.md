---
layout: default
title: auth.conf
---

# {{page.title}}
Configuration of the FNews service news storage backend

Example config:

```
enable: true

no-auth {
	no-posting: true
	no-reading: false
}

method cass {
	pwd-hash: bcrypt
	dbname: fnews
	host: host1
	host: host2
}

```

The `enable` key is used to specify, whether or not Authentication is enabled for this service or not.

The `no-auth` Entry has the following keys:
* __*no-posting*__

  Must be `true` or `false`. If `true`, an anonymous user with no account shall not be capable of posting
  Articles to this news-server.
* __*no-reading*__

  Must be `true` or `false`. If `true`, an anonymous user with no account shall not be capable of accessing
  any newsgroup on this news-server. This logically includes posting.

## The `method`-Entry

```
method <name> {
	# body
}
```

`<name>` can be one of:

* __*cass*__

  Uses Apache Cassandra as datasource.

A `method` Entry has the following keys:

* __*pwd-hash*__

  The name of the algorithm used for hashing. Commonly supported is __*plain*__, __*sha2*__/__*sha256*__
  and __*bcrypt*__.

* __*dbname*__

  The name of the datasource. Only useful if `<name>` is __*cass*__.

* __*host*__

  The host names of the datasource. Supports multiple key-value pairs. Only useful if `<name>` is __*cass*__.

* __*user*__

  The username of the datasource. Only useful if `<name>` is __*cass*__ and a username/password authentication is required.

* __*pass*__

  The password of the datasource. Only useful if `<name>` is __*cass*__ and a username/password authentication is required.

