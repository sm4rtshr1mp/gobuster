package cmd

import (
	"fmt"

	"github.com/sm4rtshr1mp/gobuster/v3/cli"
	"github.com/sm4rtshr1mp/gobuster/v3/gobusters3"
	"github.com/sm4rtshr1mp/gobuster/v3/libgobuster"
	"github.com/spf13/cobra"
)

var cmdS3 *cobra.Command

func runS3(cmd *cobra.Command, args []string) error {
	globalopts, pluginopts, err := parseS3Options()
	if err != nil {
		return fmt.Errorf("error on parsing arguments: %w", err)
	}

	plugin, err := gobusters3.NewGobusterS3(mainContext, globalopts, pluginopts)
	if err != nil {
		return fmt.Errorf("error on creating gobusters3: %w", err)
	}

	if err := cli.Gobuster(mainContext, globalopts, plugin); err != nil {
		return fmt.Errorf("error on running gobuster: %w", err)
	}
	return nil
}

func parseS3Options() (*libgobuster.Options, *gobusters3.OptionsS3, error) {
	globalopts, err := parseGlobalOptions()
	if err != nil {
		return nil, nil, err
	}

	plugin := gobusters3.NewOptionsS3()

	httpOpts, err := parseBasicHTTPOptions(cmdS3)
	if err != nil {
		return nil, nil, err
	}

	plugin.UserAgent = httpOpts.UserAgent
	plugin.Proxy = httpOpts.Proxy
	plugin.Timeout = httpOpts.Timeout

	plugin.MaxFilesToList, err = cmdS3.Flags().GetInt("maxfiles")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for maxfiles: %w", err)
	}

	plugin.NoTLSValidation, err = cmdS3.Flags().GetBool("no-tls-validation")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for no-tls-validation: %w", err)
	}

	return globalopts, plugin, nil
}

func init() {
	cmdS3 = &cobra.Command{
		Use:   "s3",
		Short: "Uses aws bucket enumeration mode",
		RunE:  runS3,
	}

	addBasicHTTPOptions(cmdS3)
	cmdS3.Flags().IntP("maxfiles", "m", 5, "max files to list when listing buckets (only shown in verbose mode)")
	cmdS3.Flags().BoolP("no-tls-validation", "k", false, "Skip TLS certificate verification")

	cmdS3.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		configureGlobalOptions()
	}

	rootCmd.AddCommand(cmdS3)
}
