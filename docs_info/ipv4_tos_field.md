---
layout: default
title: IPv4 Type Of Service (TOS)
---

# IPv4 Type Of Service (TOS)

The type of service (ToS) field in the IPv4 header has had various purposes over the years, and has been defined in different ways by five RFCs. The modern redefinition of the ToS field is a six-bit Differentiated Services Code Point (DSCP) field[2] and a two-bit Explicit Congestion Notification (ECN) field.[3] While Differentiated Services is somewhat backwards compatible with ToS, ECN is not.

## Prior to `RFC 2474`

Partially taken from [www.lartc.org / howto / lartc.qdisc.classless.html](htt:://www.lartc.org/howto/lartc.qdisc.classless.html)

The TOS octet/byte is structured as:
```
   0     1     2     3     4     5     6     7
+-----+-----+-----+-----+-----+-----+-----+-----+
|                 |                       |     |
|   PRECEDENCE    |          TOS          | MBZ |
|                 |                       |     |
+-----+-----+-----+-----+-----+-----+-----+-----+
```
### MBZ
Must be zero

### TOS

The four TOS bits (the 'TOS field') are defined as:

Binary | Decimcal |  Meaning
-|-|-
1000 |  8  |       Minimize delay (md)
0100 |  4  |       Maximize throughput (mt)
0010 |  2  |       Maximize reliability (mr)
0001 |  1  |       Minimize monetary cost (mmc)
0000 |  0  |       Normal Service

### PRECEDENCE

Binary|Decimal | Meaning
- |- | -
000 | 0 | Best Effort
001 |1 | Priority
010 |2 | Immediate
011 |3 | Flash - mainly used for voice signaling
100 |4 | Flash Override
101 |5 | Critical - mainly used for voice RTP
110 |6 | Internetwork Control
111 |7 | Network Control

## Since `RFC 2474`

Take a look into [RFC 2474](https://tools.ietf.org/html/rfc2474)

        0   1   2   3   4   5   6   7
      +---+---+---+---+---+---+---+---+
      |         DSCP          |  CU   |
      +---+---+---+---+---+---+---+---+

        DSCP: differentiated services codepoint
        CU:   currently unused

This format is also the format used by IPv6 Headers, so DSCP values can be used for both IPv4 and IPv6 connections.

The **CU** field is used for ECN (Explicit congestion notification) these days. But this should not be set by the user.

