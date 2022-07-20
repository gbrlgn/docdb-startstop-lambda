/*
This module comprises the main function and its related handler function for
AWS Lambda, StartStop. A type Event is declared with Event and Db fields, to
be passed into the function as an JSON event by the CloudWatch event trigger.
A configurable REGION constant is also declared.
*/

package main

/*
Imports log, os, strconv and time from the standard library, for fatal errors,
getting environment variables, converting said variables to integers and
comparing them with the current time; aws, session and docdb from the AWS Go
SDK for connecting into the DocumentDB cluster; the lambda module from AWS
Lambda is also imported for required integration.
*/
import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/docdb"
)

/*
Event to be passed into the function, containing an Event field, which might
be a string that is either "start" or "start", and a Db field, which must
contain the DocumentDB cluster identifier.
*/
type Event struct {
	Event string `json:"event"`
	Db    string `json:"db"`
}

// AWS DocumentDB region.
const REGION string = "us-east-1"

/*
StartStop receives an Event as a parameter and returns a string and an error.
First, it instatiates an AWS session, which is then used when creating a new
DocumentDB client. Before checking each of the events (start or stop), the
current time is compared with the start and stop time environment variables
set in the Terraform module. Incoherent times log a fatal error. If the event
field of the Event parameter contains "start", the function starts the cluster
using StartDBCluster, handling any possible errors. The opposite happens if
the field contains "stop". If the event field contains any other value, the
function exits with an informative string and a nil value.
*/
func StartStop(evt Event) (string, error) {
	// Creating a new session with SharedConfigState.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Creates a new DocumentDB client.
	db := docdb.New(sess, aws.NewConfig().WithRegion(REGION))

	if evt.Event == "start" {
		// Cluster start time from the scheduler cron job.
		var start, _ = strconv.Atoi(os.Getenv("START_TIME"))

		exec, _, _ := time.Now().Clock()

		/*
			Checks if current time is the same as the module's cron schedule
			setting. If a hour has passed, an fatal error is logged.
		*/
		if exec > start {
			log.Fatal("Wrong execution time.")
			return "Wrong execution time.", nil
		}

		/*
			Creates a StartDBClusterInput for the cluster identifier passed in
			the event by CloudWatch.
		*/
		startDb := docdb.StartDBClusterInput{
			DBClusterIdentifier: &evt.Db,
		}

		// Starts the DocumentDB cluster.
		_, err := db.StartDBCluster(&startDb)

		// Checks for errors when starting the cluster.
		if err != nil {
			log.Fatal(err)
			return "Error starting DocumentDB cluster.", err
		}

		// Exit 0.
		return "Cluster started successfully.", nil
	}

	if evt.Event == "stop" {
		// Cluster stop time from the scheduler cron job.
		var stop, _ = strconv.Atoi(os.Getenv("STOP_TIME"))

		exec, _, _ := time.Now().Clock()

		/*
			Checks if current time is the same as the module's cron schedule
			setting. If a hour has passed, an fatal error is logged.
		*/
		if exec > stop {
			log.Fatal("Wrong execution time.")
			return "Wrong execution time.", nil
		}

		/*
			Creates a StopDBClusterInput for the cluster identifier passed in
			the event by CloudWatch.
		*/
		stopDb := docdb.StopDBClusterInput{
			DBClusterIdentifier: &evt.Db,
		}

		// Stops the DocumentDB cluster.
		_, err := db.StopDBCluster(&stopDb)

		// Checks for errors when stopping the cluster.
		if err != nil {
			log.Fatal(err)
			return "Error starting DocumentDB cluster.", err
		}

		// Exit 0.
		return "Cluster stopped successfully", nil
	}

	// Exit 0 with no effect caused by any event with unwanted data.
	return "No start or stop operation executed.", nil
}

// Simply starts the handler function StartStop.
func main() {
	lambda.Start(StartStop)
}
