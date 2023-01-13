# godnsbl(modded) [![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

originally [mrichman/godnsbl](https://github.com/mrichman/godnsbl), added to specify target dns server

Package godnsbl lets you perform RBL (Real-time Blackhole List - https://en.wikipedia.org/wiki/DNSBL)
lookups using Go.

The command-line tool in `cmd` demonstrates the use of [goroutines](https://tour.golang.org/concurrency/1) to perform concurrent lookups.

To test:

```
git clone https://github.com/kitkatayama/godnsbl
cd godnsbl/cmd/godnsbl
go run main.go b.barracudacentral.org 127.0.0.2
```

The output will be a JSON-formatted list of results with the following fields:

```
{
  "rbl": "b.barracudacentral.org",
  "address": "127.0.0.2",
  "listed": true,
  "text": "http://www.barracudanetworks.com/reputation/?pr=1\u0026ip=127.0.0.2",
  "error": false,
  "error_type": null
}
```
