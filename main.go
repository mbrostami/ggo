package main

// a tool to manage cronjobs
// keep output of each cron job
// keep duration of each cronjob
// alarm for duration changes
// alarm for errors (job level | critical, normal)
// add log level for success output and failed output (-vvv , -vv , -v)
// web interface for view logs and rename jobs by creating md5 for command and mappigng with given name
// ability to connect all ggo runners together and define a main one to keep db there or sync db
// a config for log all things into the specific file
// log all ggo errors into the specific
// a command to list jobs
// a flag to run one instance of this job
import (
	"fmt"
	// "log"
	"os"
	"os/exec"
	//"strconv"
	"strings"
	"time"
)

func main() {
	//configs := loadConfigs()
	//fmt.Println(configs.Db)
	args := os.Args[1:]
	// fix me os.Args escaping qutes :()
	stringArgs := strings.Join(args[:], " ")
	fmt.Println(args)
	runJob("sample job", stringArgs)
}

func runJob(jobName string, command string) {
	cmd := make(chan *exec.Cmd)
	startTime := time.Now()
	go func() { cmd <- exec.Command("/bin/bash", "-c", command) }()
	eventRunning(jobName)
	newCmd := <-cmd
	stdo, err := newCmd.CombinedOutput()
	eventFinished(jobName, time.Since(startTime))
	if err != nil {
		eventHasError(jobName)
		handleError(jobName, stdo, err.Error())
	} else {
		handleSuccess(jobName, stdo)
	}
}

func handleSuccess(jobName string, output []byte) {
	fmt.Printf("--- SUCCESS --- \nJob name: %s\nOutput: \n%s", jobName, output)
}

func handleError(jobName string, error []byte, errorCode string) {
	fmt.Printf("--- ERROR --- \nJob name: %s\nError code: %s\nError: %s", jobName, errorCode, error)
}

func eventRunning(jobName string) {
	fmt.Println("Job:event \"" + jobName + "\" is running")
}

func eventFinished(jobName string, duration time.Duration) {
	fmt.Println("Job:event \"" + jobName + "\" is finished after " + duration.String())
}

func eventHasError(jobName string) {
	fmt.Println("Job:event \"" + jobName + "\" has failed")
}

func eventNeedSyncDb() {
	storeInDb()
}

func storeInDb() {
	fmt.Printf("USING SQLITE")
}

func loadConfigs() ConfigStruct {
	configs := Configuration()
	return configs
}
