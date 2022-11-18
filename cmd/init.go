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
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Ansybl by answering a series of questions to get started.",
	Long: `Enter the following information so we can alert you via PagerDuty when your validator is missing blocks.
	
	1) PagerDuty service ID
	2) PagerDuty email
	3) PagerDuty API token`,
	Run: func(cmd *cobra.Command, args []string) {
		get_consensus_address()

		fmt.Println("Enter your PagerDuty service ID:")
		var pd_service_id string
		fmt.Scanln(&pd_service_id)
		os.Setenv("PD_SERVICE_ID", pd_service_id)

		fmt.Println("Enter your email that's associated with your PagerDuty account:")
		var pd_email string
		fmt.Scanln(&pd_email)
		os.Setenv("PD_EMAIL", pd_email)

		fmt.Println("Enter your PagerDuty API token:")
		var pd_api_key string
		fmt.Scanln(&pd_api_key)
		os.Setenv("PD_API_KEY", pd_api_key)

		fmt.Println("Press enter to send a test alert.")
		fmt.Scanln()
		trigger_alarm()
		fmt.Println("If an alert got triggered, setup is complete.")
		fmt.Println("If you did not see an alert come through, please try again.")
		get_consensus_address()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func trigger_alarm() {
	var email = os.Getenv("PD_EMAIL")
	var service_id = os.Getenv("PD_SERVICE_ID")
	var authtoken = os.Getenv("PD_API_KEY")

	client := pagerduty.NewClient(authtoken)
	service := pagerduty.APIReference{
		ID:   service_id,
		Type: "service_reference",
	}
	body := pagerduty.APIDetails{
		Type:    "incident_body",
		Details: "Canto node is missing blocks!",
	}
	incident := pagerduty.CreateIncidentOptions{
		Type:    "incident",
		Title:   "Canto node is missing blocks",
		Service: &service,
		Body:    &body,
	}

	resp, err := client.CreateIncident(email, &incident)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

func get_consensus_address() {
	out, err := exec.Command("cantod tendermint show-address").Output()

	if err != nil {
		log.Fatal(err)
	}
	os.Setenv("CONSENSUS_ADDRESS", string(out))
	fmt.Println("Retrieved and set consensus address", string(out))
}
