
# Goal

1. Create a TCP server on port 9000.
2. Make it support HTTP/2 requests.

# References

1. [HTTP/2 Draft Specifications](https://github.com/http2/http2-spec) - define the HTTP/2 protocol.


# Validation

Install [hyper](https://hyper.readthedocs.org/en/latest/quickstart.html), a Python HTTP2 client.

```
cd ~/somewhere
virtualenv env
. ./env/bin/activate
pip install hyper
```

Then interact with your service via either the command-line:

```
hyper http://localhost:9000/
```

or via Python script:

```
from hyper import HTTPConnection
c = HTTPConnection('http2bin.org')
first = c.request('GET', '/get', headers={'key': 'value'})
second = c.request('POST', '/post', body=b'hello')
third = c.request('GET', '/ip')
print "1st", c.get_response(second)
print "2nd", c.get_response(first)
print "3rd", c.get_response(third)
```
