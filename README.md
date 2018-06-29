Wake on LAN command.

Usage:

```
wakeonlan IPADDR:PORT MAC
```

where:

*IPADDR* is destination IP addres, used by OS to determine the output network
interface. Usualy directed broadcast (eg: 192.168.1.255 in case of
192.168.1.0/24 network).  

*PORT* is destination UDP port (doesn't matter, 9 is safe because reserved for
*discard* service).

*MAC* is destination phisical address. Network interface with this addres will
wake up its machine.

Examples:

```
wakeonlan 192.168.1.255:9 00-11-22-33-44-55
wakeonlan 255.255.255.255:9 00:11:22:33:44:55
```
