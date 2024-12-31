# goTreiberStack

A basic treiber stack without locks to test its performance.

### Tests:

```
$ go test -v -race --cover ./...
Testing goTreiberStack
go test -v -race --cover ./...
=== RUN   TestStack
--- PASS: TestStack (0.00s)
=== RUN   TestRaceConditions
--- PASS: TestRaceConditions (0.20s)
=== RUN   TestThroughput
    stack_test.go:185: Throughput: 372960.40 ops/sec with 10 goroutines
--- PASS: TestThroughput (5.00s)
PASS
coverage: 100.0% of statements
ok  	github.com/el10savio/goTreiberStack/treiber	6.497s	coverage: 100.0% of statements
```

### Benchmarks:

```
$ go test -bench=. -count=15 ./...
Running benchmarks on goTreiberStack
go test -bench=. -count=15 ./...
goos: darwin
goarch: arm64
pkg: github.com/el10savio/goTreiberStack/treiber
cpu: Apple M4
BenchmarkPushSingleThreaded-10    	32955111	        32.42 ns/op
BenchmarkPushSingleThreaded-10    	45448807	        31.72 ns/op
BenchmarkPushSingleThreaded-10    	45159326	        31.68 ns/op
BenchmarkPushSingleThreaded-10    	45013502	        32.25 ns/op
BenchmarkPushSingleThreaded-10    	44336817	        32.13 ns/op
BenchmarkPushSingleThreaded-10    	44277447	        31.89 ns/op
BenchmarkPushSingleThreaded-10    	39060380	        31.16 ns/op
BenchmarkPushSingleThreaded-10    	45878938	        32.33 ns/op
BenchmarkPushSingleThreaded-10    	37849417	        31.78 ns/op
BenchmarkPushSingleThreaded-10    	40479056	        31.37 ns/op
BenchmarkPushSingleThreaded-10    	45418632	        32.27 ns/op
BenchmarkPushSingleThreaded-10    	41082816	        33.33 ns/op
BenchmarkPushSingleThreaded-10    	42684182	        32.72 ns/op
BenchmarkPushSingleThreaded-10    	45386494	        31.63 ns/op
BenchmarkPushSingleThreaded-10    	45661738	        31.34 ns/op
BenchmarkPopSingleThreaded-10     	205196066	        12.06 ns/op
BenchmarkPopSingleThreaded-10     	260270870	        11.24 ns/op
BenchmarkPopSingleThreaded-10     	179655577	        19.18 ns/op
BenchmarkPopSingleThreaded-10     	271590244	        10.30 ns/op
BenchmarkPopSingleThreaded-10     	257164164	        10.92 ns/op
BenchmarkPopSingleThreaded-10     	260224543	        10.20 ns/op
BenchmarkPopSingleThreaded-10     	259611459	        10.55 ns/op
BenchmarkPopSingleThreaded-10     	212966084	         7.593 ns/op
BenchmarkPopSingleThreaded-10     	267173428	        11.73 ns/op
BenchmarkPopSingleThreaded-10     	247361774	        14.21 ns/op
BenchmarkPopSingleThreaded-10     	261509314	        11.46 ns/op
BenchmarkPopSingleThreaded-10     	253057358	        11.26 ns/op
BenchmarkPopSingleThreaded-10     	254886474	         4.690 ns/op
BenchmarkPopSingleThreaded-10     	272164592	         6.562 ns/op
BenchmarkPopSingleThreaded-10     	258323348	         4.964 ns/op
BenchmarkConcurrentPush/goroutines-2-10         	21076588	        53.39 ns/op
BenchmarkConcurrentPush/goroutines-2-10         	17617081	        60.34 ns/op
BenchmarkConcurrentPush/goroutines-2-10         	22062067	        61.36 ns/op
BenchmarkConcurrentPush/goroutines-2-10         	20399821	        60.94 ns/op
BenchmarkConcurrentPush/goroutines-2-10         	20015136	        55.00 ns/op
BenchmarkConcurrentPush/goroutines-2-10         	18005581	        58.75 ns/op
BenchmarkConcurrentPush/goroutines-2-10         	22309948	        59.58 ns/op
BenchmarkConcurrentPush/goroutines-2-10         	19558706	        59.95 ns/op
BenchmarkConcurrentPush/goroutines-2-10         	20636808	        57.51 ns/op
BenchmarkConcurrentPush/goroutines-2-10         	17851176	        64.41 ns/op
BenchmarkConcurrentPush/goroutines-2-10         	20703001	        67.69 ns/op
BenchmarkConcurrentPush/goroutines-2-10         	19315881	        64.94 ns/op
BenchmarkConcurrentPush/goroutines-2-10         	19142482	        62.19 ns/op
BenchmarkConcurrentPush/goroutines-2-10         	20433153	        63.17 ns/op
BenchmarkConcurrentPush/goroutines-2-10         	20597866	        50.16 ns/op
BenchmarkConcurrentPush/goroutines-4-10         	13788334	        90.41 ns/op
BenchmarkConcurrentPush/goroutines-4-10         	13922007	        80.52 ns/op
BenchmarkConcurrentPush/goroutines-4-10         	13983392	        85.87 ns/op
BenchmarkConcurrentPush/goroutines-4-10         	13650892	        90.97 ns/op
BenchmarkConcurrentPush/goroutines-4-10         	13611782	        88.07 ns/op
BenchmarkConcurrentPush/goroutines-4-10         	12470647	        83.53 ns/op
BenchmarkConcurrentPush/goroutines-4-10         	13216731	        92.77 ns/op
BenchmarkConcurrentPush/goroutines-4-10         	12981058	        87.46 ns/op
BenchmarkConcurrentPush/goroutines-4-10         	14435887	        82.56 ns/op
BenchmarkConcurrentPush/goroutines-4-10         	12532603	        85.37 ns/op
BenchmarkConcurrentPush/goroutines-4-10         	13863229	        89.78 ns/op
BenchmarkConcurrentPush/goroutines-4-10         	13475278	        82.00 ns/op
BenchmarkConcurrentPush/goroutines-4-10         	14069038	        91.23 ns/op
BenchmarkConcurrentPush/goroutines-4-10         	13747665	        85.69 ns/op
BenchmarkConcurrentPush/goroutines-4-10         	11660646	        86.55 ns/op
BenchmarkConcurrentPush/goroutines-8-10         	 8909335	       132.9 ns/op
BenchmarkConcurrentPush/goroutines-8-10         	 8720209	       120.3 ns/op
BenchmarkConcurrentPush/goroutines-8-10         	 8657022	       170.4 ns/op
BenchmarkConcurrentPush/goroutines-8-10         	 5350971	       244.9 ns/op
BenchmarkConcurrentPush/goroutines-8-10         	 5683262	       245.9 ns/op
BenchmarkConcurrentPush/goroutines-8-10         	 5408887	       247.4 ns/op
BenchmarkConcurrentPush/goroutines-8-10         	 5157391	       245.3 ns/op
BenchmarkConcurrentPush/goroutines-8-10         	 5128448	       252.5 ns/op
BenchmarkConcurrentPush/goroutines-8-10         	 5262987	       251.6 ns/op
BenchmarkConcurrentPush/goroutines-8-10         	 6005272	       218.2 ns/op
BenchmarkConcurrentPush/goroutines-8-10         	 5466292	       249.7 ns/op
BenchmarkConcurrentPush/goroutines-8-10         	 6712852	       248.9 ns/op
BenchmarkConcurrentPush/goroutines-8-10         	 5867862	       194.1 ns/op
BenchmarkConcurrentPush/goroutines-8-10         	 5713899	       238.1 ns/op
BenchmarkConcurrentPush/goroutines-8-10         	 5647350	       215.4 ns/op
BenchmarkConcurrentPush/goroutines-16-10        	 8033808	       143.4 ns/op
BenchmarkConcurrentPush/goroutines-16-10        	 8007826	       156.6 ns/op
BenchmarkConcurrentPush/goroutines-16-10        	 7450542	       179.6 ns/op
BenchmarkConcurrentPush/goroutines-16-10        	 7244997	       179.9 ns/op
BenchmarkConcurrentPush/goroutines-16-10        	 6892216	       180.7 ns/op
BenchmarkConcurrentPush/goroutines-16-10        	 6799303	       185.1 ns/op
BenchmarkConcurrentPush/goroutines-16-10        	 6364639	       182.4 ns/op
BenchmarkConcurrentPush/goroutines-16-10        	 7397026	       150.6 ns/op
BenchmarkConcurrentPush/goroutines-16-10        	 7237086	       158.8 ns/op
BenchmarkConcurrentPush/goroutines-16-10        	 8213888	       167.5 ns/op
BenchmarkConcurrentPush/goroutines-16-10        	 6685999	       170.0 ns/op
BenchmarkConcurrentPush/goroutines-16-10        	 7380842	       171.1 ns/op
BenchmarkConcurrentPush/goroutines-16-10        	 7338676	       162.2 ns/op
BenchmarkConcurrentPush/goroutines-16-10        	11435130	       100.5 ns/op
BenchmarkConcurrentPush/goroutines-16-10        	11729277	       105.4 ns/op
BenchmarkConcurrentPushPop/goroutines-2-10      	29897476	        39.00 ns/op
BenchmarkConcurrentPushPop/goroutines-2-10      	30175330	        37.82 ns/op
BenchmarkConcurrentPushPop/goroutines-2-10      	34923738	        38.88 ns/op
BenchmarkConcurrentPushPop/goroutines-2-10      	30646024	        38.30 ns/op
BenchmarkConcurrentPushPop/goroutines-2-10      	31504400	        36.97 ns/op
BenchmarkConcurrentPushPop/goroutines-2-10      	35738971	        38.58 ns/op
BenchmarkConcurrentPushPop/goroutines-2-10      	31074531	        38.42 ns/op
BenchmarkConcurrentPushPop/goroutines-2-10      	31054828	        37.85 ns/op
BenchmarkConcurrentPushPop/goroutines-2-10      	32616044	        38.11 ns/op
BenchmarkConcurrentPushPop/goroutines-2-10      	36323002	        37.57 ns/op
BenchmarkConcurrentPushPop/goroutines-2-10      	29506714	        39.20 ns/op
BenchmarkConcurrentPushPop/goroutines-2-10      	35299957	        36.78 ns/op
BenchmarkConcurrentPushPop/goroutines-2-10      	33090516	        37.85 ns/op
BenchmarkConcurrentPushPop/goroutines-2-10      	32955224	        40.59 ns/op
BenchmarkConcurrentPushPop/goroutines-2-10      	29778337	        42.06 ns/op
BenchmarkConcurrentPushPop/goroutines-4-10      	19620092	        64.98 ns/op
BenchmarkConcurrentPushPop/goroutines-4-10      	19347296	        65.25 ns/op
BenchmarkConcurrentPushPop/goroutines-4-10      	19772424	        65.65 ns/op
BenchmarkConcurrentPushPop/goroutines-4-10      	19825480	        65.34 ns/op
BenchmarkConcurrentPushPop/goroutines-4-10      	18847365	        66.04 ns/op
BenchmarkConcurrentPushPop/goroutines-4-10      	18797026	        66.27 ns/op
BenchmarkConcurrentPushPop/goroutines-4-10      	19780464	        65.71 ns/op
BenchmarkConcurrentPushPop/goroutines-4-10      	19248994	        65.26 ns/op
BenchmarkConcurrentPushPop/goroutines-4-10      	18303598	        67.89 ns/op
BenchmarkConcurrentPushPop/goroutines-4-10      	17684116	        65.80 ns/op
BenchmarkConcurrentPushPop/goroutines-4-10      	21014074	        66.35 ns/op
BenchmarkConcurrentPushPop/goroutines-4-10      	19389172	        65.22 ns/op
BenchmarkConcurrentPushPop/goroutines-4-10      	20674891	        81.04 ns/op
BenchmarkConcurrentPushPop/goroutines-4-10      	12046886	        92.99 ns/op
BenchmarkConcurrentPushPop/goroutines-4-10      	17256895	        87.09 ns/op
BenchmarkConcurrentPushPop/goroutines-8-10      	 8133920	       179.2 ns/op
BenchmarkConcurrentPushPop/goroutines-8-10      	 6868760	       181.5 ns/op
BenchmarkConcurrentPushPop/goroutines-8-10      	 6720931	       180.1 ns/op
BenchmarkConcurrentPushPop/goroutines-8-10      	 8267024	       182.7 ns/op
BenchmarkConcurrentPushPop/goroutines-8-10      	 7222530	       181.8 ns/op
BenchmarkConcurrentPushPop/goroutines-8-10      	 7119036	       182.8 ns/op
BenchmarkConcurrentPushPop/goroutines-8-10      	 7022289	       176.9 ns/op
BenchmarkConcurrentPushPop/goroutines-8-10      	 6998655	       163.2 ns/op
BenchmarkConcurrentPushPop/goroutines-8-10      	 6856105	       181.4 ns/op
BenchmarkConcurrentPushPop/goroutines-8-10      	 7235562	       184.0 ns/op
BenchmarkConcurrentPushPop/goroutines-8-10      	 7380487	       177.8 ns/op
BenchmarkConcurrentPushPop/goroutines-8-10      	 6915183	       181.5 ns/op
BenchmarkConcurrentPushPop/goroutines-8-10      	 6769935	       188.3 ns/op
BenchmarkConcurrentPushPop/goroutines-8-10      	 6136521	       193.4 ns/op
BenchmarkConcurrentPushPop/goroutines-8-10      	 6288847	       193.2 ns/op
BenchmarkConcurrentPushPop/goroutines-16-10     	 5568824	       223.3 ns/op
BenchmarkConcurrentPushPop/goroutines-16-10     	 5412572	       222.4 ns/op
BenchmarkConcurrentPushPop/goroutines-16-10     	 5591935	       220.1 ns/op
BenchmarkConcurrentPushPop/goroutines-16-10     	 5469765	       223.9 ns/op
BenchmarkConcurrentPushPop/goroutines-16-10     	 5431999	       223.2 ns/op
BenchmarkConcurrentPushPop/goroutines-16-10     	 5558011	       225.5 ns/op
BenchmarkConcurrentPushPop/goroutines-16-10     	 5737201	       224.5 ns/op
BenchmarkConcurrentPushPop/goroutines-16-10     	 5504365	       221.9 ns/op
BenchmarkConcurrentPushPop/goroutines-16-10     	 5562378	       224.5 ns/op
BenchmarkConcurrentPushPop/goroutines-16-10     	 5437063	       224.0 ns/op
BenchmarkConcurrentPushPop/goroutines-16-10     	 5567288	       220.8 ns/op
BenchmarkConcurrentPushPop/goroutines-16-10     	 5437704	       225.8 ns/op
BenchmarkConcurrentPushPop/goroutines-16-10     	 5381047	       221.2 ns/op
BenchmarkConcurrentPushPop/goroutines-16-10     	 5559006	       226.3 ns/op
BenchmarkConcurrentPushPop/goroutines-16-10     	 5335094	       224.7 ns/op
BenchmarkLatency-10                             	16480185	        71.91 ns/op	        41.00 p50-ns	        42.00 p90-ns	        42.00 p99-ns
BenchmarkLatency-10                             	16921934	        71.78 ns/op	        41.00 p50-ns	        42.00 p90-ns	        42.00 p99-ns
BenchmarkLatency-10                             	16770306	        71.84 ns/op	        41.00 p50-ns	        42.00 p90-ns	        42.00 p99-ns
BenchmarkLatency-10                             	16927911	        71.77 ns/op	        41.00 p50-ns	        42.00 p90-ns	        42.00 p99-ns
BenchmarkLatency-10                             	16907431	        71.67 ns/op	        41.00 p50-ns	        42.00 p90-ns	        42.00 p99-ns
BenchmarkLatency-10                             	16940797	        71.39 ns/op	        41.00 p50-ns	        42.00 p90-ns	        42.00 p99-ns
BenchmarkLatency-10                             	17034565	        71.38 ns/op	        41.00 p50-ns	        42.00 p90-ns	        42.00 p99-ns
BenchmarkLatency-10                             	16931893	        71.22 ns/op	        41.00 p50-ns	        42.00 p90-ns	        42.00 p99-ns
BenchmarkLatency-10                             	16993756	        71.12 ns/op	        41.00 p50-ns	        42.00 p90-ns	        42.00 p99-ns
BenchmarkLatency-10                             	16929364	        71.65 ns/op	        41.00 p50-ns	        42.00 p90-ns	        42.00 p99-ns
BenchmarkLatency-10                             	16874700	        70.92 ns/op	        41.00 p50-ns	        42.00 p90-ns	        42.00 p99-ns
BenchmarkLatency-10                             	16994548	        71.38 ns/op	        41.00 p50-ns	        42.00 p90-ns	        42.00 p99-ns
BenchmarkLatency-10                             	17035552	        71.67 ns/op	        41.00 p50-ns	        42.00 p90-ns	        42.00 p99-ns
BenchmarkLatency-10                             	16863978	        71.11 ns/op	        41.00 p50-ns	        42.00 p90-ns	        42.00 p99-ns
BenchmarkLatency-10                             	16996514	        71.24 ns/op	        41.00 p50-ns	        42.00 p90-ns	        42.00 p99-ns
PASS
ok  	github.com/el10savio/goTreiberStack/treiber	527.911s
```
