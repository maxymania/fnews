---
layout: default
title: fnews.conf
---

# {{page.title}}
Configuration of the FNews service news storage frontend

## Configuring the Listener

In order to make FNews to accept connections from a certain port, we need to specify it. For this, we specify a `listen` Entry.

```
listen :119 {
	ip-version: 4
}

listen :63119 {
	ip-version: 6
}
```

This can be written as this as well:

```
listen {
	listen: :119
	ip-version: 4
}

listen {
	listen: :63119
	ip-version: 6
}
```

A `listen` Entry has the following keys:

* __*listen*__

  The `<host>:<port>` Pair to bind against. `<host>` can be omitted.

* __*ip-version*__

  Can be used to force the Listener into *IPv4-Only* (`ip-version: 4`) or *IPv6-Only* (`ip-version: 6`) mode.
  If neither *IPv4-Only* nor *IPv6-Only* is desired, use any other value (must be a number, or server won't start) or
  don't specify this key.

