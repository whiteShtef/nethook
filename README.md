# nethook
A simple client-server utility that keeps track of IP-s.

## What for?

I use this thingy to keep track of my home machine IP when going on vacation.

## Usage
### Client
```bash
$ go run main.go --logtostderr --destination 0.0.0.0:8001 --name jarjarbinks
I0727 10:07:23.325832    9015 main.go:82] Starting in client mode.
I0727 10:07:23.325835    9015 main.go:83] Reporting name jarjarbinks to 0.0.0.0:8001 every 10 seconds.
Scheduled for  2018-07-27 10:07:43.32584299 +0200 CEST
I0727 10:07:33.327143    9015 main.go:31] Reporting name jarjarbinks to 0.0.0.0:8001
I0727 10:07:33.327659    9015 main.go:40] Report successful
Scheduled for  2018-07-27 10:07:53.32584299 +0200 CEST
I0727 10:07:43.327143    9015 main.go:31] Reporting name jarjarbinks to 0.0.0.0:8001
I0727 10:07:43.327659    9015 main.go:40] Report successful
...
```
### Server
```bash
$ go run main.go --logtostderr --server --server_port  8001
I0727 10:06:13.174606    9121 main.go:69] Starting in server mode. Listening at 8001
I0727 10:07:34.079341    9121 main.go:64] Processed report from jarjarbinks at 127.0.0.1
I0727 10:07:44.080089    9121 main.go:64] Processed report from jarjarbinks at 127.0.0.1

```

```bash
$ cat reports.csv
jarjarbinks,127.0.0.1
```
