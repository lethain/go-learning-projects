import statsd, random

sc = statsd.StatsClient(host='localhost', port=8125, prefix=None, maxudpsize=512)


metrics = [str(x) for x in range(0, 100)]

for i in range(0, 5):
    metric = random.choice(metrics)
    sc.incr('%s.counter' % metric, count=i, rate=1)
    sc.gauge("%s.gauge" % metric, i, rate=1, delta=False)
    sc.timing("%s.timer" % metric, i, rate=1)
