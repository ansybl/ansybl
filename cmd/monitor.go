/*
Copyright © 2022 Brian <brian@ansybl.io>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Checks for missed blocks",
	Long: `This command checks if the validator running on this machine has missed signing any blocks. 
	If it has missed blocks it will trigger a PagerDuty alert.`,
	Run: func(cmd *cobra.Command, args []string) {
		info := get_signing_info()
		b, err := json.Marshal(info)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(b)
	},
}

func init() {
	rootCmd.AddCommand(monitorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// monitorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// monitorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func get_signing_info() []byte {
	arg0 := "query"
	arg1 := "slashing"
	arg2 := "signing-infos"
	arg3 := "--limit"
	arg4 := "200"
	arg5 := "--output"
	arg6 := "json"
	out, err := exec.Command("cantod", arg0, arg1, arg2, arg3, arg4, arg5, arg6).Output()

	if err != nil {
		log.Fatal(err)
	}
	return out
}
