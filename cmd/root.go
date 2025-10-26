package cmd

import (
	"fmt"
	"os"

	"github.com/AstraBert/generate-action-readme/parsing"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gar",
	Short: "gar (generate-action-readme) is a CLI tool to convert YAML spefications for GitHub Actions into a READMEs.",
	Long:  "gar (generate-action-readme) is a CLI tool to convert YAML spefications for GitHub Actions (such as action.yml) into a well-structured README written in GitHub-flavored markdown.",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing anydocs '%s'\n", err)
		os.Exit(1)
	}
}

var actionPath string
var readmePath string

var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"g", "gen"},
	Short:   "Generate README for action",
	Long:    "Generate the README for a GitHub Action, optionally passing its path (defaults to `action.yml`) and the path to the output README (defaults to `README.md`)",
	Run: func(cmd *cobra.Command, args []string) {
		pathToRead := "action.yml"
		pathToWrite := "README.md"
		if actionPath != "" {
			pathToRead = actionPath
		}
		if readmePath != "" {
			pathToWrite = readmePath
		}
		data, err := parsing.ParseYml(pathToRead)
		if err != nil {
			fmt.Println("Error while parsing the input file: " + err.Error() + "\n")
			os.Exit(1)
		}
		mdStr, err := parsing.ParseActionData(data)
		if err != nil {
			fmt.Println("Error while converting parsed YAML data to markdown: " + err.Error() + "\n")
			os.Exit(1)
		}
		err = os.WriteFile(pathToWrite, []byte(mdStr), 0777)
		if err != nil {
			fmt.Println("An error occurred while writing the README at the specified location: " + err.Error() + "\nPlease check path and write permissions.")
		}
		fmt.Printf("Successfully written %s!\n", pathToWrite)
	},
}

func init() {
	generateCmd.Flags().StringVarP(&actionPath, "action", "a", "", "The path to the action file to convert to README")
	generateCmd.Flags().StringVarP(&readmePath, "readme", "r", "", "The path to the README file to write")

	rootCmd.AddCommand(generateCmd)
}
