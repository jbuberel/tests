# Benchmarking strings.Contains

To replicate the testing, start by cloning the two Go repos:

```
$ mkdir go15 go15strings
$ git clone https://go.googlesource.com/go go15
$ git clone https://github.com/mesb/go go15strings
```

Then build each Go tool:
```
$ cd ~/go15/src
$ ./make.bash
$ cd ~/go15strings/src
$ ./make.bash
```

Make sure you don't have `$GOROOT` set in your environment:

```
$ unset $GOROOT
```

Next, install this repo and run the benchmark command against each Go tool:

```
$ mkdir ~/stringstest
$ cd ~/stringstest
$ export GOPATH=`pwd`
$ go get github.com/jbuberel/tests
$ cd src/github.com/jbuberel/tests
$ ~/go15/bin/go test -bench=. -benchtime=1m
[...output...]
$ ~/go15strings/bin/go test -bench=. -benchtime=1m
[...output...]
```
