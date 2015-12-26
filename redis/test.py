import redis
r = redis.StrictRedis(host='localhost', port=6379, db=0)
print "GET (should be None)\t", r.get('foo')
print "SET                 \t", r.set('foo', 'bar')
print "GET (should be bar) \t", r.get('foo')
print "DELETE (should be 1)\t", r.delete('foo')
print "GET (should be None)\t", r.get('foo')
print "DELETE (should be 0)\t", r.delete('foo')
