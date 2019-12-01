package job

import (
	"fmt"

	db "github.com/Radabaugh/funnel/internal/database"
)

// Save inserts a new Job into the database
func Save(source string, destination string, interval string, frequency int) {
	connection := db.Connection()

	sqlStatement := fmt.Sprintf("INSERT INTO jobs (source, destination, interval, frequency) VALUES ('%s', '%s', '%s', %d);", source, destination, interval, frequency)

	connection.Exec(sqlStatement)
}

// DeleteByID deletes a record from the jobs table in the database
func DeleteByID(jobID int) {
	connection := db.Connection()

	sqlStatement := fmt.Sprintf("DELETE FROM jobs WHERE jobs.id = %d;", jobID)

	connection.Exec(sqlStatement)
}
