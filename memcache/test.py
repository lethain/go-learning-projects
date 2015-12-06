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
print "incrementing: ", mc.incr("key")
print "after incr should be 1: ", mc.get("key")
print "incrementing: ", mc.incr("key")
print "after incr should be 2: ", mc.get("key")
print "decrementing: ", mc.decr("key")
print "after decr should be 1", mc.get("key")
