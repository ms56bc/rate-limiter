# Rate-Limiter

Api rate limiter with configurable eviction algorithm
Rate limiter implementations:

- Token bucket
- Fixed window counters
- Sliding window log
Let’s look at how each of them work and compare them in terms of accuracy and memory usage. 
This will give us some context for the final approach we took to build the rate limiter that spared us from the spammers.

