package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/resourceexplorer2"
	"github.com/spf13/cobra"
	"github.com/user/tfsearch/pkg/config"
)

var searchCmd = &cobra.Command{
	Use:   "search [name]",
	Short: "Search for the given name",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig("main.hcl")
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}

		name := args[0]
		var s config.Search
		for _, search := range cfg.Searches {
			if search.Name == name {
				s = search
				break
			}
		}

		// Get the provider from the config
		provider, ok := cfg.Providers[s.Provider]
		if !ok {
			fmt.Println("Invalid provider")
			fmt.Println(provider)
			os.Exit(1)
		}

		// Initialize Resource Explorer Client
		sess := session.Must(session.NewSession())
		//		sess := session.Must(session.NewSession(&aws.Config{
		//			LogLevel: aws.LogLevel(aws.LogDebugWithHTTPBody),
		//		}))
		client := resourceexplorer2.New(sess)

		// Loop through the queries
		for _, query := range s.Queries {
			queryType := query.Type // get the query type

			if queryType != "explorer" && queryType != "raw" {
				log.Fatalf("Query type '%s' is not supported for the provider 'aws'. Only 'explorer' and 'raw' are supported.", queryType)
			}

			// Ensure the query type is "explorer"
			if queryType == "explorer" {
				var tags []config.AwsTag
				for _, tag := range query.Query.Tags {
					splitTag := strings.Split(tag, ":")
					tags = append(tags, config.AwsTag{Key: splitTag[0], Value: splitTag[1]})
				}

				// Construct the query string
				var tagStrings []string
				for _, tag := range tags {
					tagString := fmt.Sprintf("-tag:%s=%s", tag.Key, tag.Value)
					tagStrings = append(tagStrings, tagString)
				}

				queryString := strings.Join(tagStrings, " ")
				region := query.Region // replace with the actual region variable if different
				queryString += fmt.Sprintf(" -region:%s", region)

				fmt.Println(queryString)

				// Construct the request
				request := &resourceexplorer2.SearchInput{
					MaxResults:  aws.Int64(1000),
					QueryString: aws.String(queryString),
				}

				// Execute the request
				response, err := client.Search(request)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				// Print the results
				fmt.Println("Found", len(response.Resources), "resources")
				for _, resource := range response.Resources {
					fmt.Println("Resource ARN: ", *resource.Arn)
					fmt.Println("Last Reported At: ", resource.LastReportedAt.Format(time.RFC3339))
					fmt.Println("Owning Account ID: ", *resource.OwningAccountId)
					fmt.Println("Region: ", *resource.Region)
					fmt.Println("Resource Type: ", *resource.ResourceType)
					fmt.Println("Service: ", *resource.Service)

					for _, property := range resource.Properties {
						fmt.Println("Property Name: ", *property.Name)
						fmt.Println("Property Last Reported At: ", property.LastReportedAt.Format(time.RFC3339))
					}

					fmt.Println("===================================")
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Args = cobra.ExactArgs(1)
}
