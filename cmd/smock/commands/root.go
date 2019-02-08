package commands

import (
	"github.com/spf13/cobra"
	"path/filepath"
	"smock/smock"
)

var rootCmd = cobra.Command{
	Use:  "smock <TypeName>",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		typeName := args[0]

		pkg, err := cmd.Flags().GetString("package")
		if err != nil {
			return err
		}

		loc, err := cmd.Flags().GetString("outputFile")
		if err != nil {
			return err
		}

		if len(loc) == 0 {
			loc, err = filepath.Abs(filepath.Join("mocks", typeName+"_on.go"))
			if err != nil {
				return err
			}
		}

		return GenerateMockOnMethods(pkg, args[0], loc)
	},
}

func init() {
	rootCmd.PersistentFlags().StringP("package", "p", ".", "Package (defaults to current package)")
	rootCmd.PersistentFlags().StringP("outputFile", "o", "", "Output file location")
}

func ExecuteRootCommand() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func GenerateMockOnMethods(pkg, typeName, fileName string) error {
	g, err := smock.NewGeneratorForType(pkg, typeName)
	if err != nil {
		return err
	}

	return g.WriteMockOnMethodsToFile(fileName)
}
