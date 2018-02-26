---
layout: default
title: DSCP, aka. TrafficClass for IPv6 and TOS for IPv4
---
# DSCP, aka. TrafficClass for IPv6 and TOS for IPv4


These days the *TrafficClass* field of IPv6 and the *TOS* field of IPv4 are defined by [RFC 2474](https://tools.ietf.org/html/rfc2474) to hold a *DSCP* value.

        0   1   2   3   4   5   6   7
      +---+---+---+---+---+---+---+---+
      |         DSCP          |  CU   |
      +---+---+---+---+---+---+---+---+

        DSCP: differentiated services codepoint
        CU:   currently unused


The **CU** field is used for **ECN** (Explicit congestion notification) these days. But this should not be set by the user (should be zero).

## Possible *DSCP* values

Name|Binary|Hex|Decimal|Description
-|-|-|-|-
CS0|000 000|0x00|0|Best Efford
CS1|001 000|0x08|8|Priority
CS2|010 000|0x10|16|Immediate
CS3|011 000|0x18|24|Flash - mainly used for voice signaling
CS4|100 000|0x20|32|Flash Override
CS5|101 000|0x28|40|Critical - mainly used for voice RTP
CS6|110 000|0x30|48|Internetwork Control
CS7|111 000|0x38|56|Network Control
AF11|001 010|0x0a|10|Assured Forwarding
AF12|001 100|0x0c|12|Assured Forwarding
AF13|001 110|0x0e|14|Assured Forwarding
AF21|010 010|0x12|18|Assured Forwarding
AF22|010 100|0x14|20|Assured Forwarding
AF23|010 110|0x16|22|Assured Forwarding
AF31|011 010|0x1a|26|Assured Forwarding
AF32|011 100|0x1c|28|Assured Forwarding
AF33|011 110|0x1e|30|Assured Forwarding
AF41|100 010|0x22|34|Assured Forwarding
AF42|100 100|0x24|36|Assured Forwarding
AF43|100 110|0x26|38|Assured Forwarding
EF|101 110|0x2e|46|Expedited forwarding

### Assured Forwarding

The Assured Forwarding constants Are a combination of a Class Selector __CS*x*__ and a __Drop Propability__.
The first three bits, '*xxx*000' of the CS*x* constant are used and concatenated with a three bit drop propablility *yyy* tho that we get 'xxx*yyy*'.

Drop P. / Class |CS1|CS2|CS3|CS4
-|-|-|-|-
**Low drop Propability** | AF11 | AF21 | AF31 | AF41
**Med drop Propability** | AF12 | AF22 | AF32 | AF42
**High drop Propability** | AF13 | AF23 | AF33 | AF43


## Further Reading
- [Cisco: Implementing Quality of Service Policies with DSCP](https://www.cisco.com/c/en/us/support/docs/quality-of-service-qos/qos-packet-marking/10103-dscpvalues.html)
- [Cisco: DSCP and Precedence Values](https://www.cisco.com/c/en/us/td/docs/switches/datacenter/nexus1000/sw/4_0/qos/configuration/guide/nexus1000v_qos/qos_6dscp_val.pdf)
- [Differentiated services](https://en.wikipedia.org/wiki/Differentiated_services)

