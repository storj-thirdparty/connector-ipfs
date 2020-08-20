## Flow Diagram

![](https://github.com/storj-thirdparty/connector-ipfs/blob/master/README.assets/arch.drawio.png)

## Config Files

There are two config files that contain Storj network and IPFS connection information. The tool is designed so you can specify a config file as part of your tooling/workflow.



##### `ipfs_property.json`

Inside the `./config` directory there is a `ipfs_property.json` file, with following information about your IPFS instance:

* hostname :- Host Name connect to IPFS
* port	   :- Port Number connect to IPFS
* hash	   :- IPFS Hash of file/folder

##### `storj_config.json`

Inside the `./config` directory a `storj_config.json` file, with Storj network configuration information in JSON format:

* apiKey - API Key created in Storj Satellite GUI (mandatory)
* satelliteURL - Storj Satellite URL (mandatory)
* encryptionPassphrase - Storj Encryption Passphrase (mandatory)
* bucketName - Name of the bucket to upload data into (mandatory)
* uploadPath - Path on Storj Bucket to store data (optional) or "" or "/" (mandatory)
* serializedAccess - Serialized access shared while uploading data used to access bucket without API Key (mandatory)
* allowDownload - Set *true* to create serialized access with restricted download (mandatory while using *share* flag)
* allowUpload - Set *true* to create serialized access with restricted upload (mandatory while using *share* flag)
* allowList - Set *true* to create serialized access with restricted list access
* allowDelete - Set *true* to create serialized access with restricted delete
* notBefore - Set time that is always before *notAfter*
* notAfter - Set time that is always after *notBefore*


## Run

The following flags can be used with the `store` command:

* `accesskey` - Connects to the Storj network using a serialized access key instead of an API key, satellite url and encryption passphrase.
* `shared` - Generates a restricted shareable serialized access with the restrictions specified in the Storj configuration file.

Once you have built the project run the following commands:

##### Get help

```
$ ./connector-ipfs --help
```

##### Check version

```
$ ./connector-ipfs --version
```

##### Connect to IPFS and upload the files using their corresponding hash to Storj

```
$ ./connector-ipfs store --ipfs <path_to_ipfs_config_file> --storj <path_to_storj_config_file>
```

##### Connect to IPFS and upload the files using their corresponding hash to Storj bucket using Access Key

```
$ ./connector-ipfs store --accesskey
```

##### Connect to IPFS and upload the files using their corresponding hash to Storj and generate a Shareable Access Key based on restrictions in `storj_config.json`

```
$ ./connector-ipfs store --share
```



##  Testing

The project has been tested on the following operating systems:

```
* Windows
	* Version: 10 Pro
	* Processor: Intel(R) Core(TM) i3-5005U CPU @ 2.00GHz 2.00GHz

* macOS Catalina
	* Version: 10.15.4
	* Processor: 2.5 GHz Dual-Core Intel Core i5

* ubuntu
	* Version: 16.04 LTS
	* Processor: AMD A6-7310 APU with AMD Radeon R4 Graphics × 4
```
