package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	shell "github.com/ipfs/go-ipfs-api"
)

// ConfigIpfs defines the variables and types for login.
type ConfigIpfs struct {
	HostName string `json:"hostName"`
	Port     string `json:"port"`
	Hash     string `json:"hash"`
}

// LoadIpfsProperty reads and parses the JSON file
// that contains a IPFS instance's properties
// and returns all the properties as an object.
func LoadIpfsProperty(fullFileName string) ConfigIpfs {

	var configIpfs ConfigIpfs

	// Open and read the file.
	fileHandle, err := os.Open(filepath.Clean(fullFileName))
	if err != nil {
		log.Fatal(err)
	}

	jsonParser := json.NewDecoder(fileHandle)
	if err = jsonParser.Decode(&configIpfs); err != nil {
		log.Fatal(err)
	}

	err = fileHandle.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Display Information about IPFS Instance.
	fmt.Println("Read IPFS configuration from the ", fullFileName, " file")
	fmt.Println("Host Name\t: ", configIpfs.HostName)
	fmt.Println("Port\t\t: ", configIpfs.Port)
	fmt.Println("Hash\t\t:", configIpfs.Hash)

	return configIpfs
}

// ConnectToIpfs will connect to an IPFS instance
// based on the read property from an external file.
// It returns a reference to shell with IPFS instance's information.
func ConnectToIpfs(configIpfs ConfigIpfs) *shell.Shell {

	fmt.Println("Connecting to IPFS...")

	if configIpfs.HostName == "ipfsHostName" || configIpfs.HostName == "" {
		err1 := errors.New("Invalid HostName")
		log.Fatal("Invalid Hostname error : ", err1)
	}
	// Connect IPFS deamon to IPFS node.
	sh := shell.NewShell(configIpfs.HostName + ":" + configIpfs.Port)
	_, _, errVer := sh.Version()
	if errVer != nil {
		err1 := errors.New("Could not find Daemon running")
		log.Fatal("Daemon error : ", err1)
	}

	return sh
}

// AllHash is a list to store hashes of all files to store from IPFS.
var AllHash []string

// AllFilesWithPaths is a list to store complete path of all the hash in the IPFS.
// It will be used for direct transfer of files from IPFS to Storj.
var AllFilesWithPaths []string

// GetHash retrieve the hash with the exact
// file structure from IPFS to the System
func GetHash(sh *shell.Shell, hash string, path string) {

	files, err := sh.FileList(hash)
	if err != nil {
		log.Fatal(err)
	}
	if files.Type == "Directory" {
		if len(files.Links) > 0 {
			for i := 0; i < len(files.Links); i++ {
				if files.Links[i].Type == "Directory" {
					GetHash(sh, files.Links[i].Hash, path+"/"+files.Links[i].Hash)
				} else {
					AllHash = append(AllHash, files.Links[i].Hash)
					AllFilesWithPaths = append(AllFilesWithPaths, path+"/"+files.Links[i].Hash)
				}
			}
		}
	} else {
		AllHash = append(AllHash, hash)
		if AllFilesWithPaths != nil {
			AllFilesWithPaths = append(AllFilesWithPaths, path+"/"+hash)
		} else {
			AllFilesWithPaths = append(AllFilesWithPaths, hash)
		}
	}
}

// GetReader returns a Reader of corresponding file whose path is specified.
// io.ReadCloser type of object returned is used to perform transfer of file to Storj.
func GetReader(sh *shell.Shell, path string) io.ReadCloser {

	ipfsReader, err := sh.Cat(path)
	if err != nil {
		log.Fatal("Read file error: ", err)
	}

	return ipfsReader
}
