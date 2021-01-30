# Funktional

Is an implementation of basic functional programming helpers in go.

It is a project implemented in a couple of hours just for fun, so I would not treat it as a production-ready solution at the moment.
It is still work in progress.

**It leverages golang v2**, so it cannot be used for enterprise.

As go2 has limitations, proper unit-testing is not implemented yet, just a bunch of tests in `test.go` file.

## Helper functions

* Map
* Filter
* Reduce
* Chunk
* Fill
* FirstIndex
* _To be continued..._

## Prepare your environment

Prepare your local environment for go2 according to this article - 
[How to install go2 on MacOS](https://blog.bullgare.com/2021/01/install-go2-beta-on-macos/).

## Run

```sh
make run
```

It will run all the tests and also will output an example for each helper function.