package cmd

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// cmntsCmd represents the cmnts command
var cmntsCmd = &cobra.Command{
	Use:   "cmnts",
	Short: "Returns all single-line go comments in target files",
	Long:  `Arguments are each a path to a single file.`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		fmt.Printf("attempting to get comments for %d file(s)\n", len(args))

		for _, fn := range args {
			fmt.Printf("getting comments for file %s\n", fn)

			var fb []byte
			if fb, err = ioutil.ReadFile(fn); nil != err {
				fmt.Printf("error reading file with name %s, continuing on to next file, err: %s\n", fn, err)
				continue
			}

			goCodeLines := strings.Split(string(fb), "\n")

			green := color.New(color.FgGreen).SprintFunc()
			var comments []string
			for i, ln := range goCodeLines {
				ln = strings.TrimSpace(ln)
				if strings.HasPrefix(ln, "//") {
					c := fmt.Sprintf("%s:%d	%s", fn, i+1, green(ln))
					comments = append(comments, c)
				}
			}

			if len(comments) == 0 {
				fmt.Printf("no comments found in %d lines of code, continuing on\n", len(goCodeLines))
				continue
			}

			fmt.Printf("found %d comment lines in %d lines of code, will log out comments for spell checking\n\n", len(comments), len(goCodeLines))

			for _, c := range comments {
				fmt.Println(c)
			}
			fmt.Println()
			fmt.Println("* spell checking service: https://grademiners.com/spell-checker")
		}
	},
}

func init() {
	rootCmd.AddCommand(cmntsCmd)
}
