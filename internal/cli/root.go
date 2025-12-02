package cli

import (
	"fmt"
	"os"

	"github.com/atopile/atopile/internal/version"
	"github.com/spf13/cobra"
)

var (
	verbose        int
	nonInteractive bool
	debug          bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ato",
	Short: "atopile - Design circuit boards with code",
	Long: `atopile is a language, compiler, and toolchain for electronics.
Write hardware like software with declarative .ato files, deep validation, 
and layout that works natively with KiCad.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Set up logging based on verbosity
		// TODO: implement proper logging setup
		if debug {
			fmt.Println("Debug mode enabled")
			// TODO: setup debugger waiting
		}
		
		// Check for updates
		if cmd.Name() != "version" && cmd.Name() != "self-check" {
			version.CheckForUpdate()
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Global flags
	rootCmd.PersistentFlags().IntVarP(&verbose, "verbose", "v", 0, "Increase verbosity (-v, -vv, -vvv)")
	rootCmd.PersistentFlags().BoolVar(&nonInteractive, "non-interactive", false, "Run in non-interactive mode")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Wait to attach debugger on start")
	
	// Version flags
	rootCmd.Flags().BoolP("version", "", false, "Output version string")
	rootCmd.Flags().BoolP("semver", "", false, "Output semver-compliant version string")
	
	// Hidden flags for internal use
	rootCmd.Flags().Bool("python-path", false, "Print Python interpreter path")
	rootCmd.Flags().MarkHidden("python-path")
	
	rootCmd.Flags().Bool("atopile-path", false, "Print atopile source path")
	rootCmd.Flags().MarkHidden("atopile-path")
	
	// Handle version flags in PreRun
	rootCmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		if v, _ := cmd.Flags().GetBool("version"); v {
			fmt.Println(version.GetVersion())
			os.Exit(0)
		}
		if s, _ := cmd.Flags().GetBool("semver"); s {
			fmt.Println(version.Parse(version.GetVersion()))
			os.Exit(0)
		}
		if p, _ := cmd.Flags().GetBool("python-path"); p {
			fmt.Println("Go version does not use Python interpreter")
			os.Exit(0)
		}
		if a, _ := cmd.Flags().GetBool("atopile-path"); a {
			if ex, err := os.Executable(); err == nil {
				fmt.Println(ex)
			}
			os.Exit(0)
		}
		return nil
	}
}
