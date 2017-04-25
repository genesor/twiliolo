# Twiliolo

Golang API wrapper for Twilio API [WIP]

[![Build Status](https://travis-ci.org/genesor/twiliolo.svg?branch=master)](https://travis-ci.org/genesor/twiliolo)
[![Go Report Card](https://goreportcard.com/badge/github.com/genesor/twiliolo)](https://goreportcard.com/report/github.com/genesor/twiliolo)
[![GoDoc](https://godoc.org/github.com/genesor/twiliolo?status.svg)](https://godoc.org/github.com/genesor/twiliolo)



# Installation

``` bash
go get github.com/genesor/twiliolo
```

# Documentation

[GoDoc](http://godoc.org/github.com/genesor/twiliolo)

# Usage

## Get an Incoming phone number with its Sid

``` go
package main

import (
  "fmt"
  "net/http"
  "github.com/genesor/twiliolo"
)

func main() {
  client := twiliolo.NewClient("ACCOUNT_SID", "AUTH_TOKEN", &http.Client{})

  number, err := client.IncomingPhoneNumber.Get("NUMBER_SID")
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(number.FriendlyName)
  }
}
```
