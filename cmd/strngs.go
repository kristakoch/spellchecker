package cmd

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// strngsCmd represents the strngs command
var strngsCmd = &cobra.Command{
	Use:   "strngs",
	Short: "Returns all lines containing string punctuation",
	Long:  `Arguments are each a path to a single file.`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		fmt.Printf("attempting to get lines with strings for %d file(s)\n", len(args))

		for _, fn := range args {
			fmt.Printf("getting lines with strings for file %s\n", fn)

			var fb []byte
			if fb, err = ioutil.ReadFile(fn); nil != err {
				fmt.Printf("error reading file with name %s, continuing on to next file, err: %s\n", fn, err)
				continue
			}

			goCodeLines := strings.Split(string(fb), "\n")

			green := color.New(color.FgGreen).SprintFunc()
			var strLines []string
			for i, ln := range goCodeLines {
				ln = strings.TrimSpace(ln)
				if strings.Contains(ln, "\"") || strings.Contains(ln, "`") {
					c := fmt.Sprintf("%s:%d	%s", fn, i+1, green(ln))
					strLines = append(strLines, c)
				}
			}

			if len(strLines) == 0 {
				fmt.Printf("no lines with strings found in %d lines of code, continuing on\n", len(goCodeLines))
				continue
			}

			fmt.Printf("found %d string lines in %d lines of code, will log out lines with strings for spell checking\n\n", len(strLines), len(goCodeLines))

			for _, c := range strLines {
				fmt.Println(c)
			}
			fmt.Println()
			fmt.Println("* spell checking service: https://grademiners.com/spell-checker")
		}
	},
}

func init() {
	rootCmd.AddCommand(strngsCmd)
}
