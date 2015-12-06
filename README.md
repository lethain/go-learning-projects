This repository contains a variety of projects intended for practicing the [Go](https://golang.org)
programming language. Each of these projects is intended to provide experience with Go, while also providing
exposure to the implementation details of an interesting piece of technology you probably already use,
such as Statsd, Memcache or Redis.


The intended way to use these projects is:

1. Select one of the below projects.
2. Read it's README.
3. Do not look at the reference implementations.
4. Take a stab at implementing it yourself.
5. Take a look at the reference implementations.
6. File issues if you have suggestions for improvements.

Caveat: these projects are absolutely not trying to be production grade solutions!
These are learning projects, and I would never recommend that you replace Statsd with
the output of the below Statsd project, etc.

Projects:

1. [Writing a Statsd Server in Go.](./statsd/)
2. [Writing Memcached in Go.](./memcache/)


Future project ideas:

1. Writing a Redis server in Go.
2. Writing a Memcache server in Go.
3. Writing a PostgreSQL server in Go.
4. Consuming MySQL binlog in Go.
5. Write a [Dynamo](https://www.google.com/url?sa=t&rct=j&q=&esrc=s&source=web&cd=1&cad=rja&uact=8&ved=0ahUKEwj2g9uE58fJAhXjjIMKHUspAT8QFggcMAA&url=http%3A%2F%2Fwww.allthingsdistributed.com%2Ffiles%2Famazon-dynamo-sosp2007.pdf&usg=AFQjCNHhJccl0_0I9x7tkWizMx6NjcuUkQ&sig2=MxsX4LhM7QJRYg4GPcdGeA&bvm=bv.108538919,d.amc) inspired key-value store in Go.