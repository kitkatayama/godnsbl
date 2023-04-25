# godnsbl(modded) [![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

originally [mrichman/godnsbl](https://github.com/mrichman/godnsbl), added to specify target dns server

Package godnsbl lets you look up a specified network with specified [DNSBL](https://en.wikipedia.org/wiki/DNSBL) server using Go.

The command-line tool in `cmd` demonstrates the use of [goroutines](https://tour.golang.org/concurrency/1) to perform concurrent lookups.

To test:

```
git clone https://github.com/kitkatayama/godnsbl
cd godnsbl/cmd/godnsbl
go run main.go b.barracudacentral.org 127.0.0.2/32
```

The output will be a JSON-formatted list of results with the following fields:

```
$ go run main.go b.barracudacentral.org 127.0.0.2/31 | jq .
[
  {
    "rbl": "b.barracudacentral.org",
    "address": "127.0.0.2",
    "a": "127.0.0.2",
    "text": "http://www.barracudanetworks.com/reputation/?pr=1&ip=127.0.0.2",
    "error": false,
    "error_type": null
  },
  {
    "rbl": "b.barracudacentral.org",
    "address": "127.0.0.3",
    "a": "",
    "text": "",
    "error": true,
    "error_type": {
      "Err": "no such host",
      "Name": "3.0.0.127.b.barracudacentral.org",
      "Server": "127.0.0.53:53",
      "IsTimeout": false,
      "IsTemporary": false,
      "IsNotFound": true
    }
  }
]
```

Note: This repository is not compatible with upstream.
