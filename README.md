# fibrous

Comparing I/O (async or otherwise) code in various languages/frameworks/platforms.

Currently implemented:
* node.js
* Go

Soon:
* Python
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

## Running the benchmark

Run ```./bench.sh``` from the root of the repository, and wait for ```boom``` to complete the run.

# Results

Results are available and will be update in this gist:

https://gist.github.com/kidoman/1e39bcbf8e70099400ce