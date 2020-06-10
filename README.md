## connector-ipfs (uplink v1.0.5)

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/ac7bbc539c0a45a5a2140b3e6b7c823d)](https://app.codacy.com/gh/storj-thirdparty/connector-IPFS?utm_source=github.com&utm_medium=referral&utm_content=storj-thirdparty/connector-IPFS&utm_campaign=Badge_Grade_Dashboard)
[![Go Report Card](https://goreportcard.com/badge/github.com/storj-thirdparty/connector-ipfs)](https://goreportcard.com/report/github.com/storj-thirdparty/connector-ipfs)
[![Cloud Build](https://storage.googleapis.com/storj-utropic-services-badges/builds/connector-ipfs/branches/master.svg)


## Overview

The IPFS Connector connects to an IPFS server, takes a backup of the specified files and uploads the backup data on Storj network.

```bash
Usage:
  connector-ipfs [command] <flags>

Available Commands:
  help        Help about any command
  store       Command to upload data to a Storj V3 network.
  version     Prints the version of the tool

```

`store` - Connect to the specified IPFS (default: `ipfs_property.json`). Back-up of the IPFS is generated using tooling provided by IPFS and then uploaded to the Storj network.  Connect to a Storj v3 network using the access specified in the Storj configuration file (default: `storj_config.json`).

Sample configuration files are provided in the `./config` folder.

## Requirements and Install

To build from scratch, [install the latest Go](https://golang.org/doc/install#install).

> Note: Ensure go modules are enabled (GO111MODULE=on)


#### Option #1: clone this repo (most common)

To clone the repo

```
git clone https://github.com/storj-thirdparty/connector-ipfs.git
```

Then, build the project using the following:

```
cd connector-ipfs
go build
```

#### Option #2:  ``go get`` into your gopath

 To download the project inside your GOPATH use the following command:

```
go get github.com/storj-thirdparty/connector-ipfs
```

## Connect to IPFS Server

Make sure you are connected to IPFS server. If not, run the ipfs daemon in another `terminal` to join your node to the public network:

```
$ ipfs daemon
```

## Run (short version)

Once you have built the project run the following commands as per your requirement:

##### Get help

```
$ ./connector-ipfs --help
```

##### Check version

```
$ ./connector-ipfs --version
```

##### Create backup from ipfs and upload to Storj

```
$ ./connector-ipfs store
```

## Documentation

For more information on runtime flags, configuration, testing, and diagrams, check out the [Detail](//github.com/storj-thirdparty/wiki/Detail) or jump to:

* [Config Files](//github.com/storj-thirdparty/connector-ipfs/wiki/#config-files)
* [Run (long version)](//github.com/storj-thirdparty/connector-ipfs/wiki/#run)
* [Testing](//github.com/storj-thirdparty/connector-ipfs/wiki/#testing)
* [Flow Diagram](//github.com/storj-thirdparty/connector-ipfs/wiki/#flow-diagram)
