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

* __*hop-limit*__

  Specifies the Hop-Limit (aka. TTL in IPv4) of the outgoing IP packets. The hop limit, must
  be `1 <= value <= 255` when specifying. Otherwise, it will be ignored.

* __*dscp*__

  Specifies the DSCP bits of the Traffic-Class (IPv6) or TOS (IPv4) field.
  It must be eighter a decimal or hexadecimal integer, or a name of a [DSCP constant](../docs_info/traffic_class).
  Valid values are `0 <= value <= 63`. If the value is `0` or `CF0`, and might be overridden by the Operating System
  or the network equipment.

