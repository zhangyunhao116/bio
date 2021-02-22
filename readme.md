# bio

*Better io utils for Go*



### FixedReadAll

FixedReadAll reads from r into a preallocated slice until an error or EOF.

About 30% faster and save 80% memory compared to `io.ReadAll`.

```
name                          time/op
All/FixedReadAll-_1_MB-16     2.50ms ± 1%
All/io.ReadAll-___1_MB-16     3.18ms ± 0%
All/FixedReadAll-_5_MB-16     11.4ms ± 4%
All/io.ReadAll-___5_MB-16     15.2ms ± 2%
All/FixedReadAll-_10_MB-16    19.8ms ± 3%
All/io.ReadAll-___10_MB-16    28.1ms ± 4%
All/FixedReadAll-_100_MB-16    172ms ± 9%
All/io.ReadAll-___100_MB-16    212ms ± 3%
All/FixedReadAll-_1024_MB-16   2.31s ±12%
All/io.ReadAll-___1024_MB-16   2.12s ± 4%

name                          alloc/op
All/FixedReadAll-_1_MB-16     1.06MB ± 0%
All/io.ReadAll-___1_MB-16     5.86MB ± 0%
All/FixedReadAll-_5_MB-16     5.25MB ± 0%
All/io.ReadAll-___5_MB-16     28.8MB ± 0%
All/FixedReadAll-_10_MB-16    10.5MB ± 0%
All/io.ReadAll-___10_MB-16    56.6MB ± 0%
All/FixedReadAll-_100_MB-16    105MB ± 0%
All/io.ReadAll-___100_MB-16    529MB ± 0%
All/FixedReadAll-_1024_MB-16  1.07GB ± 0%
All/io.ReadAll-___1024_MB-16  6.17GB ± 0%
```

