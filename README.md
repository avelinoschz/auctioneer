# auctioneer

## About

Auctioneer is simple auto-auction system developed in Go.

## Table Of Contents

- [auctioneer](#auctioneer)
  - [About](#about)
  - [Table Of Contents](#table-of-contents)
    - [How it works?](#how-it-works)
  - [Pre-requisites](#pre-requisites)
  - [Running the project](#running-the-project)
    - [Installed locally](#installed-locally)
    - [Dockerized](#dockerized)

### How it works?

An `auctioneer` instance is needed to be created and configured.

For now the only configuration needed is the `maxRounds` that the auction will perform.

```go
a := auctioneer.NewAuctioneer(auctioneer.WithMaxRounds(99))
```

Then, a slice of `bidder` type instances is needed. These are the params for the creation:

```go
auctioneer.NewBidder(auctioneer.BidderParams{
    Name:       "funky name",
    InitialBid: 2500,
    MaxBid:     3500,
    Increment:  250,
})
```

> [!NOTE] Next version will support a currency type, to avoid using `float64`.

With the participants generated, what's left to do, is perform the auction and get the winner:

```go
winner, err := a.Auction(bidders)
```

## Pre-requisites

The program runs using only the standard library of Go. So the only pre-requisite is to have installed the [Go](https://go.dev/doc/install) language.

The project in particular uses:

- Go 1.22.2 >=

## Running the project

You can build and run the project manually with the `go` tool and commands. The preferred way is using either the provided `Makefile` or `Docker`.

### Installed locally

```shell
# builds the binary and run the program.
make run-go
```

There are some other useful commands like

```shell
# runs the unit tests of the project.
make run-test

# runs the tests too, and generate the cover file ands open it in the brower.
make open-cover 
```

### Dockerized

If you don't have `Go` installed and preferred only to use Docker, there are some make targets too.

```shell
# ensures docker is installed and builds the image.
make build-docker

# runs the previously generated docker image and remove the container after.
make run-docker
```
