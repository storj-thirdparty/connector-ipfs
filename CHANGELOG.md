# connector-ipfs Changelog

## [1.0.5] - 27-05-2020
### Changelog:
* Added cobra cli for user interface.
* Restructured project based on the requirements for cobra cli.
* Changed arguments to optional flags.
* Added `--accesskey` flag and removed `key`, `test`, `parse` and `debug` flags.
* Added `--storj` flag to set storj config file path and `--ipfs` to set ipfs config file path.
* Fixed bugs


## [1.0.7] - 04-12-2019
### Changelog:
* Added Macroon functionality.
* Added option to access storj using Serialized Scope Key. 
* Added keyword `key` to access storj using API key rather than Serialized Scope Key (defalt).
* Added keyword `restrict` to apply restrictions on API key and provide shareable Serialized Scope Key for users.


## [1.0.6] - 29-11-2019
### Changelog:
* Optimized storj-download.
* Exception handling in storj connection.
* Made changes according to updated libuplink package.


## [1.0.5] - 25-11-2019
### Changelog:
* Removed parse and parse debug commands.
* Error handling for various events.


## [1.0.4] - 22-11-2019
### Changelog:
* Added download debug command.
* Updated and optimized test command.
* Optimized upload process by using io.reader instead to bytes.reader.
* The program shows error if daemon is not running.
* fileHash has renamed to shareableHash in ipfs_download.json.
* Retry mechanism is added for upload function.
* Insufficent memory(system freezing) issue when uploading large data resolved.


## [1.0.3] - 15-11-2019
### Changelog:
* Made storj-ipfs-connector able to upload large files.
* Changes made in json fies.


## [1.0.2] - 12-11-2019
### Changelog:
* Made ipfsHostName and PortNumber configurable.
* Changes made according to updated cli package.


## [1.0.1] - 08-11-2019
### Changelog:
* Added download
* Added upload debug
* Added test
* Added storj-ipfs as module
* Changed main project folder to storj-ipfs-connector


## [1.0.0] - 05-11-2019