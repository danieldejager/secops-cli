/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package commands

import (
	"github.com/danieldejager/secops-cli/internal/wrappers"
	"github.com/spf13/cobra"
)

// VirustotalCmd represents the virustotal command
func NewVirustotalCmd(virustotalWrapper wrappers.VirusTotalWrapper) *cobra.Command {
	vtCmd := &cobra.Command{
		Use:   "virustotal",
		Short: "The virustotal command allows you to get intellgence about a file hash, domain or IP Address",
		Long:  `The SECOPS CLI is a fully functional Command Line Interface (CLI) that interacts with a variety of security related applications`,
	}
	return vtCmd
}
