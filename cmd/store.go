package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// storeCmd represents the store command.
var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Command to upload data to storj V3 network.",
	Long:  `Command to connect to desired IPFS account and back-up the complete data to given Storj Bucket.`,
	Run:   ipfsStore,
}

func init() {

	// Setup the store command with its flags.
	rootCmd.AddCommand(storeCmd)
	var defaultIpfsFile string
	var defaultStorjFile string
	storeCmd.Flags().BoolP("accesskey", "a", false, "Connect to storj using access key(default connection method is by using API Key).")
	storeCmd.Flags().BoolP("share", "s", false, "For generating share access of the uploaded backup file.")
	storeCmd.Flags().StringVarP(&defaultIpfsFile, "ipfs", "i", "././config/ipfs_property.json", "full filepath contaning IPFS configuration.")
	storeCmd.Flags().StringVarP(&defaultStorjFile, "storj", "u", "././config/storj_config.json", "full filepath contaning storj V3 configuration.")
}

func ipfsStore(cmd *cobra.Command, args []string) {

	// Process arguments from the CLI.
	ipfsConfigfilePath, _ := cmd.Flags().GetString("ipfs")
	fullFileNameStorj, _ := cmd.Flags().GetString("storj")
	useAccessKey, _ := cmd.Flags().GetBool("accesskey")
	useAccessShare, _ := cmd.Flags().GetBool("share")

	// Read IPFS instance's configurations from an external file and create an IPFS configuration object.
	configIpfs := LoadIpfsProperty(ipfsConfigfilePath)

	// Read storj network configurations from and external file and create a storj configuration object.
	storjConfig := LoadStorjConfiguration(fullFileNameStorj)

	// Connect to storj network using the specified credentials.
	access, project := ConnectToStorj(storjConfig, useAccessKey)

	// Connect to IPFS using the specified credentials
	ipfsShell := ConnectToIpfs(configIpfs)

	// Create the slices of hash names and file hash paths of files to be uploaded.
	GetHash(ipfsShell, configIpfs.Hash, configIpfs.Hash)

	fmt.Printf("\nInitiating back-up.\n")
	// Fetch all hash from IPFS instance and simultaneously store them into desired Storj bucket.
	for i := 0; i < len(AllHash); i++ {
		file := AllFilesWithPaths[i]
		ipfsReader := GetReader(ipfsShell, AllHash[i])
		UploadData(project, storjConfig, file, ipfsReader)
	}
	fmt.Printf("\nBack-up complete.\n\n")

	// Create restricted shareable serialized access if share is provided as argument.
	if useAccessShare {
		ShareAccess(access, storjConfig)
	}
}
