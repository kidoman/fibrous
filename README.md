# fibrous

Comparing I/O (async or otherwise) code in various languages/frameworks/platforms.

Currently implemented:
* node.js
* Go
* Python (peer review pending)
* Ruby (peer review pending)

Soon:
* Scala
* Java
* Clojure

# How to run

I am currently testing the performance of the respective solutions using ```wrk```.

This requires you to have a valid Go installation and a properly setup ```GOPATH```.

Also, The test currently also requires a ```Redis``` server to be available at ```localhost:6379```.

## Go

CD into the ```go``` folder and run either ```./start_martini.sh``` or ```./start_stdlib.sh``` to start the server to benchmark.

## Nodejs

CD into the ```nodejs``` folder and run either ```./start_stdlib.sh```, ```./start_callback.sh```, ```./start_promise.sh``` or ```./start_fiber.sh``` to start the server to benchmark.

## Python

CD into the ```python``` folder and run either ```./start_twisted.sh``` or ```./start_tornado.sh``` to start the server to benchmark.

**Note:** If you are on Mac OS X, and the compile fails because of ```cc```, run the command with prefix ```CC=gcc-4.2```. You will need to install ```gcc-4.2``` by doing ```brew install apple-gcc42```.

## Ruby

CD into the ```ruby``` folder and run ```./start_sinatra.sh``` to start the server to benchmark.

## Running the benchmark

Run ```./bench.sh``` from the root of the repository, and wait for ```wrk``` to complete the run.

# Results

We run:

```
wrk -c100 -d20s --latency http://127.0.0.1:3000/users/1
```

and record the result. 20 seconds of requests at a concurrency of 100.

**Disclaimer:** The code probably could use some cleanup/fixes/optimizations. So the numbers below are only representative until the community gets a chance to look them over.

## Summary

All the following tests are run on the same hardware. The Redis instance is local, so is the test runner. All the servers are geared to use as many cores as they want.

* Golang (stdlib) 134566.47 requests/s *
* nodejs (stdlib) 54510.15 requests/s
* Golang (martini) 51330.49 requests/s
* nodejs (callbacks) 36106.59 requests/s
* nodejs (fibers) 27371.92 requests/s
* nodejs (promises) 22664.96 requests/s

\* Redis was pegged at 100%, and became the bottleneck

## Detailed Results

### node.js (Stdlib)

```
Running 20s test @ http://127.0.0.1:3000/users/1
  2 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.85ms    1.63ms  27.29ms   86.34%
    Req/Sec    29.17k     5.77k   44.44k    60.83%
  Latency Distribution
     50%    1.56ms
     75%    2.51ms
     90%    3.24ms
     99%    7.78ms
  1090190 requests in 20.00s, 152.83MB read
Requests/sec:  54510.15
Transfer/sec:      7.64MB
```

### node.js (Callbacks)

```
Running 20s test @ http://127.0.0.1:3000/users/1
  2 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.69ms    2.75ms  38.54ms   84.83%
    Req/Sec    19.63k     3.63k   29.89k    60.01%
  Latency Distribution
     50%    1.77ms
     75%    4.08ms
     90%    6.68ms
     99%   10.84ms
  722079 requests in 20.00s, 139.79MB read
Requests/sec:  36106.59
Transfer/sec:      6.99MB
```

### node.js (Promises)

```
Running 20s test @ http://127.0.0.1:3000/users/1
  2 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     4.49ms    3.92ms  28.70ms   68.95%
    Req/Sec    11.78k     2.51k   17.27k    58.65%
  Latency Distribution
     50%    3.90ms
     75%    6.06ms
     90%    8.78ms
     99%   17.15ms
  453271 requests in 20.00s, 87.75MB read
Requests/sec:  22664.96
Transfer/sec:      4.39MB
```

### node.js (Fibers)

```
Running 20s test @ http://127.0.0.1:3000/users/1
  2 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     3.70ms    4.01ms  48.38ms   87.00%
    Req/Sec    14.12k     1.72k   19.68k    67.04%
  Latency Distribution
     50%    3.00ms
     75%    5.09ms
     90%    7.95ms
     99%   18.76ms
  547399 requests in 20.00s, 105.97MB read
Requests/sec:  27371.92
Transfer/sec:      5.30MB
```

### Go (stdlib)

```
Running 20s test @ http://127.0.0.1:3000/users/1
  2 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   742.08us  834.30us  14.64ms   87.19%
    Req/Sec    71.39k     9.82k  132.70k    81.53%
  Latency Distribution
     50%  373.00us
     75%    1.04ms
     90%    1.78ms
     99%    3.81ms
  2691299 requests in 20.00s, 369.59MB read
Requests/sec: 134566.47
Transfer/sec:     18.48MB
```

### Go (martini)

```
Running 20s test @ http://127.0.0.1:3000/users/1
  2 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.96ms    1.96ms  26.09ms   87.17%
    Req/Sec    27.13k     3.43k   49.22k    82.44%
  Latency Distribution
     50%    1.27ms
     75%    2.47ms
     90%    4.42ms
     99%    9.51ms
  1026622 requests in 20.00s, 145.88MB read
Requests/sec:  51330.49
Transfer/sec:      7.29MB
```
