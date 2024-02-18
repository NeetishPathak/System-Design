package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Job struct {
	id     string
	status int
}

var mutex1 sync.Mutex

var jobMap = make(map[string]Job)
var tickerDone chan bool
var jobReady = make(chan string, 1)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func queueJob(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	jobId := strconv.FormatInt(time.Now().Unix(), 10)
	mutex1.Lock()
	jobMap[jobId] = Job{id: jobId}
	mutex1.Unlock()
	fmt.Println("Received jobId ", jobId)
	go fireUpdateJobs(jobMap[jobId])
	fmt.Fprintf(w, jobId)
}

func checkJobStatus(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	v := r.URL.Query().Get("jobid")
	if v == "" {
		fmt.Println("Job Id not provided")
		return
	} else {
		mutex1.Lock()
		job, ok := jobMap[v]
		mutex1.Unlock()
		if ok {
			fmt.Println("Will check for ", job.id)

			for {
				select {
				case jid := <-jobReady:
					// while the status is not 100, so not send and just loop
					// fmt.Println("jobready receieved for ", jid)
					if jid == job.id {
						fmt.Fprintf(w, "jobId:status -> "+job.id+":"+fmt.Sprintf("%d", jobMap[v].status))
						return
					}
				default:
					mutex1.Lock()
					st := jobMap[v].status
					mutex1.Unlock()
					if st >= 300 {
						fmt.Fprintf(w, "jobId:status -> "+job.id+":"+fmt.Sprintf("%d", st))
						return
					}
				}
			}
		}

	}
}

func fireUpdateJobs(j Job) {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-tickerDone:
			ticker.Stop()
			return
		case <-ticker.C:
			mutex1.Lock()
			j.status += 10
			jobMap[j.id] = j
			mutex1.Unlock()
			if j.status >= 300 {
				jobReady <- j.id
				tickerDone <- true
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
