package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Job struct {
	id     string
	status int
}

var jobMap = make(map[string]Job)
var procDone chan bool

func queueJob(w http.ResponseWriter, r *http.Request) {
	jobId := strconv.FormatInt(time.Now().Unix(), 10)
	jobMap[jobId] = Job{id: jobId}
	fmt.Println("Received jobId ", jobId)
	go fireUpdateJobs(jobMap[jobId])
	fmt.Fprintf(w, jobId)
}

func checkJobStatus(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query().Get("jobid")
	if v == "" {
		fmt.Println("Job Id not provided")
		return
	} else {
		job, ok := jobMap[v]
		if ok {
			fmt.Fprintf(w, "jobId:status -> "+job.id+":"+fmt.Sprintf("%d", job.status))
		}
	}
}

func fireUpdateJobs(j Job) {
	ticker := time.NewTicker(3 * time.Second)
	for {
		select {
		case <-procDone:
			ticker.Stop()
			return
		case <-ticker.C:
			j.status += 10
			jobMap[j.id] = j
			if j.status >= 100 {
				procDone <- true
			}
		}
	}
}

func main() {
	http.HandleFunc("/submit", queueJob)
	http.HandleFunc("/checkjob", checkJobStatus)
	fmt.Println("Start listening at :8080")
	http.ListenAndServe(":8080", nil)
}
