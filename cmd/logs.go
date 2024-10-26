package cmd

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// lgCmd represents the lg command
var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Logs all lines containing Infof, Errorf, and Warningf",
	Long:  `Arguments are each a path to a single file.`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		fmt.Printf("attempting to get log lines for %d file(s)\n", len(args))

		for _, fn := range args {
			fmt.Printf("getting log lines for file %s\n", fn)

			var fb []byte
			if fb, err = ioutil.ReadFile(fn); nil != err {
				fmt.Printf("error reading file with name %s, continuing on to next file, err: %s\n", fn, err)
				continue
			}

			goCodeLines := strings.Split(string(fb), "\n")

			green := color.New(color.FgGreen).SprintFunc()
			var logLines []string
			for i, ln := range goCodeLines {
				ln = strings.TrimSpace(ln)
				if strings.Contains(ln, "Infof") || strings.Contains(ln, "Warningf") || strings.Contains(ln, "Errorf") {
					c := fmt.Sprintf("%s:%d	%s", fn, i+1, green(ln))
					logLines = append(logLines, c)
				}
			}

			if len(logLines) == 0 {
				fmt.Printf("no comments found in %d lines of code, continuing on\n", len(goCodeLines))
				continue
			}

			fmt.Printf("found %d comment lines in %d lines of code, will log out log lines for spell checking\n\n", len(logLines), len(goCodeLines))

			for _, c := range logLines {
				fmt.Println(c)
			}
			fmt.Println()
			fmt.Println("* spell checking service: https://grademiners.com/spell-checker")
		}
	},
}

func init() {
	rootCmd.AddCommand(logsCmd)
}
