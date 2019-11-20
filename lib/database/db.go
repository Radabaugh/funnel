package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	// Go postgres driver
	_ "github.com/lib/pq"
)

// Job represents a job in the database
type Job struct {
	ID, Frequency, NumberOfRuns                                                               int
	Source, Destination, Interval, JobType, LastRunDuration, LastRun, NextRun, Name, Schedule string
}

// Connection returns an open db connection
func Connection() (db *sql.DB) {
	env := os.Getenv("FUNNEL_ENV")
	if "" == env {
		env = "development"
	}

	godotenv.Load(".env." + env)
	godotenv.Load()

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	return db
}

// SelectJobs returns everything from the jobs table
func SelectJobs() []Job {
	db := Connection()
	rows, err := db.Query("select * from jobs")
	if err != nil {
		panic(err)
	}

	job := Job{}
	jobs := []Job{}

	for rows.Next() {
		var id, frequency, numberOfRuns int
		var source, destination, interval, jobType, lastRunDuration, lastRun, nextRun string

		err = rows.Scan(&id, &source, &destination, &interval, &frequency, &jobType, &lastRun, &lastRunDuration, &numberOfRuns, &nextRun)
		if err != nil {
			panic(err)
		}

		job.Name = source + "to" + destination
		job.Schedule = strconv.Itoa(frequency) + " " + interval
		job.JobType = jobType
		job.LastRunDuration = lastRunDuration
		job.LastRun = lastRun
		job.NextRun = nextRun
		job.NumberOfRuns = numberOfRuns
		jobs = append(jobs, job)
	}
	defer db.Close()

	return jobs
}
