# epilog

[![Build Status](https://travis-ci.org/go-st/epilog.svg?branch=master)](https://travis-ci.org/go-st/epilog)

Epilog is a simple, fast and reliable logger for GO. 

Epilog natively supports base `ILogger` interface from package `github.com/go-st/logger`. 
So, if you use `ILogger` interface in your application - 
migration to Epilog will be very easy.

Inspired by Monolog.

## Install
Using go get:
`go get github.com/go-st/epilog`

Using [glide](https://github.com/masterminds/glide): 
`glide get github.com/go-st/epilog`

Checkout latest version on [release page](https://github.com/go-st/epilog/releases).

## Usage

```
logger := New("MyLogger", DefaultHandler)
logger.Debug("hello debug")
logger.Info("hello info")

// Output:
// [2017-01-18T16:00:00.000000+00:00] (DEBUG) hello debug
// [2017-01-18T16:00:00.000000+00:00] (INFO) hello info
```

Extended usage examples see in [example_test.go](https://github.com/go-st/epilog/blob/master/example_test.go)
