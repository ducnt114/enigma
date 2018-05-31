package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Enigma",
	Long:  `All software has versions. This is Enigma's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Enigma ProtoBuf File Generator v1.0")
	},
}
