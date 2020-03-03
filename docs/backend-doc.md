---
layout: default
title: backend.conf
---

# {{page.title}}
Configuration of the FNews service news storage backend

## The article storage

The configuration file must contain a `store` entry in order to be valid.
This entry specifies, how, and where FNews stores xover, header and body of an article.

```
store <methodname> {
	# body
}
```

`<methodname>` is the name of a storage method for articles. Currently available storage methods are:

* __*cass*__

  Uses Apache Cassandra as storage backend. Requires a contained `keyspace`-Entry, see [below](#keyspace).

### The article-to-group assignment.

The configuration file must contain a `groupindex` entry in order to be valid.
This entry specifies, how, and where FNews assigns articles to `newsgroup+number` pairs.

```
groupindex <methodname> {
	# body
}
```

`<methodname>` is the name of a storage method for article-to-group assignments. Currently available storage methods are:

* __*cass*__

  Uses Apache Cassandra as storage backend. Requires a contained `keyspace`-Entry, see [below](#keyspace).

### The newsgroups+description+status storage.

The configuration file must contain a `grouplist` entry in order to be valid.
This entry specifies, how, and where FNews stores the list of all newsgroups
including their attributes like description and status.

```
grouplist <methodname> {
	# body
}
```

`<methodname>` is the name of a storage method for newsgroups+description+status records. Currently available storage methods are:

* __*cass*__

  Uses Apache Cassandra as storage backend. Requires a contained `keyspace`-Entry, see [below](#keyspace).

### The article-to-group assignment.

The configuration file must contain a `grouphead` entry in order to be valid.
This entry specifies, how, and where FNews allocates article-numbers for each posted article.

```
grouphead <methodname> {
	# body
}
```

`<methodname>` is the name of a storage method for articles. Currently available storage methods are:

* __*postgres*__

  Uses PostgreSQL as storage backend. The Entry must contain a `dburl` key, specifying the connection url
  It usually looks like: `'user=usr password=pwd dbname=mydatabase sslmode=disable'`
  See also: [lib/pq](https://godoc.org/github.com/lib/pq)

## Specifying the article retention

On usenet-servers, articles have a certain retention, after which they are purged from the server.
By default FNews specifies a retention of `30` days.

The `retention`-Entry is optional in {{page.title}}, and is useful to specify another Retention.

```
retention {
	# `incremental' must be one of `true' or `false'
	#    `true'  : Iterate through all `element'-Entries and apply those which match.
	#    `false' : Iterate through the `element'-Entries; if the entry matches, then
	#              apply the entry and terminate the loop.
	incremental: false
	element {
		# Expire after N days.
		expire-after: 200
		
		# Specifies a filter, an article must match to get the specified expiration time.
		where {
			# Requirement 1: The article must be posted to a newsgroup
			# that starts with `de.' ...
			newsgroups: de.*
			# ... but not with `de.ctrl.'!
			except: de.ctrl.*
			
			# Requirement 2: The article must not be cross-posted to a newsgroup,
			# that starts with `alt.binaries'!
			exclude: alt.binaries.*
			
			# Requirement 3: The article must be between 100 bytes and 10000 bytes long!
			size: 100,10000
			
			# Requirement 4: The article must be between 5 bytes and 100 lines long!
			lines: 5,100
		}
	}
	
	# This is an example, how to specify a default retention.
	# Note that `incremental' is `false'
	element {
		expire-after: 50
	}
}
```

## Specifying Cassandra Keyspaces.

The `keyspace`<a name="keyspace">-</a>Entry must be contained by any `store`,
`groupindex`, `grouplist` or `grouphead`-record with the `<methodname>`=__*cass*__.

```
keyspace <name> {
	#body
}
```

A `keyspace`-Entry has the following keys:

* __*cluster*__

  (Optional) Used to differentiate between different connection pools. By default, the
  __*cass*__ storage method create one connection-pool for every distinct `keyspace`-`<name>`
  which can cause a conflict, if two Cassandra-clusters are used.

* __*host*__

  The host-name which can be used

* __*user*__

  Username. Useful if the Cassandra-cluster requires an authentication.

* __*pass*__

  Password. Useful if the Cassandra-cluster requires an authentication.

```
store cass {
	keyspace auth_keyspace {
		host: node1.cassandra.local
		host: node2.cassandra.local
		host: node3.cassandra.local
		user: CassandraUser
		pass: Secret123
	}
}
```

