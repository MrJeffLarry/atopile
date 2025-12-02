package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	buildEntry          string
	selectedBuilds      []string
	buildTargets        []string
	excludeTargets      []string
	frozen              bool
	keepPickedParts     bool
	keepNetNames        bool
	keepDesignators     bool
	standalone          bool
	openLayout          bool
)

var buildCmd = &cobra.Command{
	Use:   "build [ENTRY]",
	Short: "Build the specified target(s) or targets from build config",
	Long: `Build the specified --target(s) or the targets specified by the build config.
Optionally specify a different entrypoint with the argument ENTRY.
Example: ato build --target my_target path/to/source.ato:module.path`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Parse entry argument
		if len(args) > 0 {
			buildEntry = args[0]
		}

		fmt.Printf("Building project...\n")
		fmt.Printf("Entry: %s\n", buildEntry)
		fmt.Printf("Selected builds: %v\n", selectedBuilds)
		fmt.Printf("Targets: %v\n", buildTargets)
		fmt.Printf("Exclude targets: %v\n", excludeTargets)
		
		// TODO: Implement actual build logic
		// This will involve:
		// 1. Loading configuration
		// 2. Checking and installing dependencies
		// 3. Running the build pipeline
		// 4. Optionally opening the layout in KiCad
		
		return fmt.Errorf("build command not yet fully implemented")
	},
}

func init() {
	buildCmd.Flags().StringSliceVarP(&selectedBuilds, "build", "b", []string{}, "Select specific build(s)")
	buildCmd.Flags().StringSliceVarP(&buildTargets, "target", "t", []string{}, "Build specific target(s)")
	buildCmd.Flags().StringSliceVarP(&excludeTargets, "exclude-target", "x", []string{}, "Exclude specific target(s)")
	buildCmd.Flags().BoolVar(&frozen, "frozen", false, "PCB must be rebuilt without changes (useful in CI)")
	buildCmd.Flags().BoolVar(&keepPickedParts, "keep-picked-parts", false, "Keep picked parts")
	buildCmd.Flags().BoolVar(&keepNetNames, "keep-net-names", false, "Keep net names")
	buildCmd.Flags().BoolVar(&keepDesignators, "keep-designators", false, "Keep designators")
	buildCmd.Flags().BoolVar(&standalone, "standalone", false, "Build in standalone mode")
	buildCmd.Flags().BoolVar(&openLayout, "open", false, "Open layout after build")
	
	rootCmd.AddCommand(buildCmd)
}
