package database

import (
	sql "database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	// Go postgres driver
	_ "github.com/lib/pq"
)

// Job represents a job in the database
type Job struct {
	ID, Frequency                                          int
	Source, Destination, Interval, JobType, Name, Schedule string
	NumberOfRuns                                           sql.NullInt32
	LastRunDuration, LastRun, NextRun                      sql.NullString
}

// Connection returns an open db connection
func Connection() (db *sql.DB) {
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
		var id, frequency int
		var source, destination, interval, jobType string
		var numberOfRuns sql.NullInt32
		var lastRunDuration, lastRun, nextRun sql.NullString

		err = rows.Scan(&id, &source, &destination, &interval, &frequency, &jobType, &lastRun, &lastRunDuration, &numberOfRuns, &nextRun)
		if err != nil {
			panic(err)
		}

		job.Name = source + " to " + destination
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
