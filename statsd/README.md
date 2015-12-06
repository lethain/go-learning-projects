Our first project will be focused on getting more comfortable with Go by writing a simple implementation
of [[ https://github.com/etsy/statsd | statsd ]] in Go. This will give exposure to writing a UDP server in Go,
as well as to the simple [Statsd Protocol](https://github.com/b/statsd_spec). The simple joy of this is that
you'll be able to hook use existing statsd client libraries to communicate with your new server.

For those who are not familiar with Statsd, it is generally used with [Graphite](http://graphite.wikidot.com/),
running locally on each of your servers to collect metrics and then periodically flushes collected metrics from
each server into Graphite. This greatly reduces the load on the Graphite server, allowing you to collect far more
metrics with fewer computing resources.

## Prerequisites

1. [Install Go on your laptop. ](https://golang.org/dl/)
2. [Do the Go Tour.](https://tour.golang.org/welcome/1)
3. [Review the Go Language Reference.](https://golang.org/ref/spec)

## Goals

1.  Write a Go server that runs on port 8125.
2.  It should use either TCP or UDP or both, but note that Python script below assumes UDP
    and you'll have to pass a different flag to it if you prefer TCP.
3.  Implement support for the the [Counter](https://github.com/b/statsd_spec) type (one of three types of metrics supported by Statsd).
4.  Buffer metrics and flush them to stdout or stderr every 10 seconds.
5.  (Bonus) Use the [flag](https://golang.org/pkg/flag/) module to allow changing port.
6.  (Bonus) Implement Gauge and Timer types.
7.  (Bonus) Synchronize flushing to avoid threading race conditions using either the [sync](https://golang.org/pkg/sync/) module
    or [channels](https://gobyexample.com/channels).

## References

1. [Statsd Protocol Specification ](https://github.com/b/statsd_spec)
2. [net module for UDP and TCP servers ](https://golang.org/pkg/net/)
3. [flag module for command line flags](https://golang.org/pkg/flag/)
4. [time module for periodic execution](https://golang.org/pkg/time/)

## Validation

Here is a simple example of testing your implementation in Python using the
[statsd](http://statsd.readthedocs.org/en/v3.2.1/reference.html) module.

First, install dependencies:

```
cd ~/somewhere
virtualenv env
. ./env/bin/activate
pip install statsd
python
```

Then from within the Python shell (see [stats.py](./stats.py)).

```
import statsd
sc = statsd.StatsClient(host='localhost', port=8125, prefix=None, maxudpsize=512)

for i in range(0, 5):
    sc.incr('counter', count=i, rate=1)
    sc.gauge("gauge", i, rate=1, delta=False)
    sc.timing("timer", i, rate=1)
```

## Reference Implementations

Here are a couple of reference implementations:

1. [statsd.go](./statsd.go) - simple solution which does not synchronize flushing
2. [channels.go](./channels.go) - solution which synchronizes flushing using channels
