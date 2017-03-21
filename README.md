# Twiliolo

Golang API wrapper for Twilio API [WIP]

[![Build Status](https://travis-ci.org/Genesor/twiliolo.svg?branch=master)](https://travis-ci.org/Genesor/twiliolo)


# Installation

``` bash
go get github.com/Genesor/twiliolo
```

# Documentation

[GoDoc](http://godoc.org/github.com/Genesor/twiliolo)

# Usage

## Get an Incoming phone number with its Sid

``` go
package main

import (
  "fmt"
  "github.com/Genesor/twiliolo"
)

func main() {
  client := twiliolo.NewClient("ACCOUNT_SID", "AUTH_TOKEN")
  
  number, err := twiliolo.GetIncomingPhoneNumber(client, "NUMBER_SID")
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(number.FriendlyName)
  }
}
```
