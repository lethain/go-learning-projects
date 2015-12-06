[Memcached](http://memcached.org/) is one of the widest spread caching solution,
and like Statsd or Redis, it also operates with a custom protocol between its clients
and servers.

This project looks at learning about the Memcached protocol while implementing a TCP server
in Go.


## Prerequisites

1. [Install Go on your laptop. ](https://golang.org/dl/)
2. [Do the Go Tour.](https://tour.golang.org/welcome/1)
3. [Review the Go Language Reference.](https://golang.org/ref/spec)

## Goals

1. Write a TCP server in Go that runs on port 11211.
2. It should support `set`, `get`, `incr` and `decr` commands.
3. It should synchronize data access for safe concurrent access.
4. (Bonus) Add Redis-like feature of periodically snapshotting data to disk.
5.  (Bonus) Allow [long-polling](https://en.wikipedia.org/wiki/Push_technology#Long_polling) style
    subscription to modifications on a key.

## References

1. [memcached protocol](https://github.com/memcached/memcached/blob/master/doc/protocol.txt) - specification for the Memcached protocol.
2. [python-memcached](https://github.com/linsomniac/python-memcached) - a Python client for connecting to Memcached.
3. [net](https://golang.org/pkg/net/) - Go module for UDP and TCP servers.
4. [sync](https://golang.org/pkg/sync/) - Go module for synchronizing data via locks.
5. [channels](https://gobyexample.com/channels) - Go language feature for safe data passing between goroutines.

## Validation

One way to test your server is using the [python-memcached](https://github.com/linsomniac/python-memcached) module,
which is a pure Python implementation without any external dependencies.

To install:

```
cd ~/somewhere
virtualenv env
. ./env/bin/activate
pip install python-memcached
```

Then to use:

```
import memcache
mc = memcache.Client(['127.0.0.1:11211'], debug=0)

mc.set("some_key", "Some value")
print "set to 'Some value': ", mc.get("some_key")
mc.set("another_key", 3)
print "set to 3: ", mc.get("another_key")
mc.delete("another_key")
print "deleted", mc.get("another_key")
mc.set("key", "1")
print "set", mc.get("key")
mc.incr("key")
print "incr", mc.get("key")
mc.decr("key")
print "decr", mc.get("key")
```
## Implementations

After you've spent some time on it, here are some reference implementations:

1. [memcached.go](./memcached.go) - implementation of the non-bonus componenits using `sync.RWMutex` for consistency.

