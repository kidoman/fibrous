# fibrous

Comparing I/O (async or otherwise) code in various languages/frameworks/platforms.

Currently implemented:
* node.js
* Go
* Python

Soon:
* Scala
* Java
* Clojure

# How to run

I am currently testing the performance of the respective solutions using ```boom```.

This requires you to have a valid Go installation and a properly setup ```GOPATH```.

Also, The test currently also requires a ```Redis``` server to be available at ```localhost:6379```.

## Go

CD into the ```go``` folder and run either ```./start_martini.sh``` or ```./start_stdlib.sh``` to start the server to benchmark.

## Nodejs

CD into the ```nodejs``` folder and run either ```./start_callback.sh```, ```./start_promise.sh``` or ```./start_fiber.sh``` to start the server to benchmark.

## Python

CD into the ```python``` folder and run either ```./start_twisted.sh``` or ```./start_tornado.sh``` to start the server to benchmark.

**Note:** If you are on Mac OS X, and the compile fails because of ```cc```, run the command with prefix ```CC=gcc-4.2```. You will need to install ```gcc-4.2``` by doing ```brew install apple-gcc42```.

## Running the benchmark

Run ```./bench.sh``` from the root of the repository, and wait for ```boom``` to complete the run.

# Results

We run:

```
boom -n 100000 -c 100 http://localhost:3000/users/1
```

and record the result. 100_000 requests at a concurrency of 100.

**Disclaimer:** The code probably could use some cleanup/fixes/optimizations. So the numbers below are only representative until the community gets a chance to look them over.

## Summary

All the following tests are run on the same hardware. The Redis instance is local, so is the test runner. All the servers are effectively single threaded (GOMAXPROCS = 1 and only a single instance of the node app was started.)

* Golang (stdlib) 31066 req/s
* Golang (martini) 12126 req/s
* nodejs (callbacks) 8077 req/s
* nodejs (fibers) 5060 req/s
* nodejs (promises) 4963 req/s
* CPython (tornado) 2739 req/s
* CPython (twisted) 2716 req/s

## Detailed Results

### node.js (Callbacks)

```
Summary:
  Total:  12.3812 secs.
  Slowest:  0.1213 secs.
  Fastest:  0.0033 secs.
  Average:  0.0124 secs.
  Requests/sec: **8076.7825**
  Total Data Received:  2800000 bytes.
  Response Size per Request:  28 bytes.

Status code distribution:
  [200] 100000 responses
```

### node.js (Promises)

```
Summary:
  Total:  20.1474 secs.
  Slowest:  0.1079 secs.
  Fastest:  0.0043 secs.
  Average:  0.0201 secs.
  Requests/sec: **4963.4166**
  Total Data Received:  2800000 bytes.
  Response Size per Request:  28 bytes.

Status code distribution:
  [200] 100000 responses
```

### node.js (Fibers)

```
Summary:
  Total:  19.7636 secs.
  Slowest:  0.1668 secs.
  Fastest:  0.0062 secs.
  Average:  0.0197 secs.
  Requests/sec: **5059.7968**
  Total Data Received:  2800000 bytes.
  Response Size per Request:  28 bytes.

Status code distribution:
  [200] 100000 responses
```

### Go (stdlib)

```
Summary:
  Total:  3.2190 secs.
  Slowest:  0.1164 secs.
  Fastest:  0.0001 secs.
  Average:  0.0032 secs.
  Requests/sec: **31065.6375**
  Total Data Received:  2700000 bytes.
  Response Size per Request:  27 bytes.

Status code distribution:
  [200] 100000 responses
```

### Go (martini)

```
Summary:
  Total:  8.2470 secs.
  Slowest:  0.1317 secs.
  Fastest:  0.0002 secs.
  Average:  0.0082 secs.
  Requests/sec: **12125.5973**
  Total Data Received:  2600000 bytes.
  Response Size per Request:  26 bytes.

Status code distribution:
  [200] 100000 responses
```

### CPython (twisted)

```
Summary:
  Total:  36.8215 secs.
  Slowest:  0.1481 secs.
  Fastest:  0.0048 secs.
  Average:  0.0368 secs.
  Requests/sec: 2715.8060
  Total Data Received:  3100000 bytes.
  Response Size per Request:  31 bytes.

Status code distribution:
  [200] 100000 responses
```

### CPython (tornado)

```
Summary:
  Total:  36.5037 secs.
  Slowest:  0.1242 secs.
  Fastest:  0.0034 secs.
  Average:  0.0365 secs.
  Requests/sec: 2739.4510
  Total Data Received:  3100000 bytes.
  Response Size per Request:  31 bytes.

Status code distribution:
  [200] 100000 responses
```
