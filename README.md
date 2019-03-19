# Granitic 2 tutorial source code

Completed examples and starting points for the tutorials at [granitic.io/tutorials](http://www.granitic.io/tutorials)

## Requirements

You must have Go 1.11 or later and have your `GOPATH` environment variable correctly set.

You must have followed the [Granitic installation instructions](http://www.granitic.io/getting-started-installing-granitic)
and (if you want to use YAML for configuration) the [Granitic YAML installation instructions](https://github.com/graniticio/granitic-yaml/blob/master/README.md)

### Note for Go 1.11 and 1.12 users

If you clone this repo to a location inside your GOPATH, you will need to the environment variable

`GO111MODULE=on`

## Running the JSON tutorials

Open a terminal in the folder where you have checked out this repo.

```go
cd json/001/recordstore
go mod download
grnc-bind && go build && ./recordstore
```

## Running the JSON tutorials

Open a terminal in the folder where you have checked out this repo.

```go
cd yaml/001/recordstore
go mod download
grnc-yaml-bind && go build && ./recordstore
```

## Which folder should I use?

Each numbered folder contains the _completed_ code for a tutorial. For example, if I wanted to followed tutorial 3, I 
would start using the code for tutorial 2.
