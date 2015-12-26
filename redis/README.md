[Redis[(http://redis.io) is a widely used datastore, offering a rich array of
datstructures (lists, sets, sorted sets, ...), fast performance and awkward
high availability strategies.

It also offers a very simple protocol (far simpler than the Memcache
protocol), making it a good learning project.

## Prerequisites

1. [Install Go on your laptop. ](https://golang.org/dl/)
2. [Do the Go Tour.](https://tour.golang.org/welcome/1)
3. [Review the Go Language Reference.](https://golang.org/ref/spec)

## Goals

1. Write a TCP server in Go that runs on port `6379`.
2. It should support `get`, `set` and `del` operations.
3. It should synchronize data access for safe concurrent read/writes (look into the [sync](https://golang.org/pkg/sync/) package).
4. Bonus: periodically synchronize data to disc (take a look at the [gob](https://golang.org/pkg/encoding/gob/) package).
5. Bonus: [support more commands](http://redis.io/commands), especially consider how to
    support the blocking pop commands (`BRPOP` and `BLPOP`), which are often used to
    implement simple queuing systems on top of Redis.

## References

1. [Redis protocol](http://redis.io/topics/protocol) - specification for the Redis protocol.
3. [net](https://golang.org/pkg/net/) - Go module for UDP and TCP servers.
4. [sync](https://golang.org/pkg/sync/) - Go module for synchronizing data via locks.
4. [gob](https://golang.org/pkg/encoding/gob/) - Go module, similar to Python's pickle, for serializing Go objects.
5. [channels](https://gobyexample.com/channels) - Go language feature for safe data passing between goroutines.
6. [Python Redis client](https://pypi.python.org/pypi/redis) - a Python client to use for validating your implementation.

## Validation

One easy way to test our server is to use [py-redis](https://pypi.python.org/pypi/redis).

To install:

```
cd ~/somewhere
virtualenv env
. ./env/bin/activate
pip install redis
```

Then run this script:

```
import redis
r = redis.StrictRedis(host='localhost', port=6379, db=0)
print "GET (should be None)\t", r.get('foo')
print "SET                 \t", r.set('foo', 'bar')
print "GET (should be bar) \t", r.get('foo')
```
