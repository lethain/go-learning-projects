import memcache
mc = memcache.Client(['127.0.0.1:11211'], debug=0)

# test incr/decr first, which is where I started out implementaiton
keys = ('a', 'b')
for i in range(0, 10):
    for k in keys:
        print "incr %s by %s\t=>\t%s\tGet %s" % (k, i, mc.incr(k, i), mc.get(k))
for i in range(0, 10):
    for k in keys:
        print "decr %s by %s\t=>\t%s\tGet %s" % (k, i, mc.decr(k, i), mc.get(k))
for k in keys:
    mc.set(k, 0)
    print "Zeroing out %s\t=>\t%s" % (k, mc.get(k))
print ""

mc.set("some_key", "Some value")
print "set to 'Some value': ", mc.get("some_key")
mc.set("another_key", 3)
print "set to 3: ", mc.get("another_key")
mc.delete("another_key")
print "deleted", mc.get("another_key")
mc.set("key", "1")
print "set", mc.get("key")

