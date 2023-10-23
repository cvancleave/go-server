package database

import (
	"encoding/json"
	"go-server/internal/models"

	_ "github.com/lib/pq"
)

// GetTimesheetsByUserId - get all timesheets by user id
func (db *DB) GetTimesheetsByUserId(userId int) ([]models.Timesheet, error) {

	stmt, err := db.Client.Prepare("SELECT * FROM timesheets WHERE user_id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var timesheets []models.Timesheet
	for rows.Next() {
		var timesheet models.Timesheet
		var dataJson string
		if err = rows.Scan(
			&timesheet.ID,
			&timesheet.UserID,
			&timesheet.DateWeek,
			&timesheet.DateSubmitted,
			&dataJson,
		); err != nil {
			return nil, err
		}
		// unmarshal the json blob into the timesheet data field
		if err := json.Unmarshal([]byte(dataJson), &timesheet.Data); err != nil {
			return nil, err
		}
		timesheets = append(timesheets, timesheet)
	}

	return timesheets, nil
}

// InsertTimesheet - get all timesheets by user id
func (db *DB) InsertTimesheet(timesheet models.Timesheet) error {

	stmt, err := db.Client.Prepare(`
		INSERT INTO timesheets (user_id, date_week, date_submitted, data)
		VALUES ($1, $2, $3, $4)
		RETURNING id`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	dataJson, _ := json.Marshal(timesheet.Data)
	_, err = stmt.Exec(timesheet.UserID, timesheet.DateWeek, timesheet.DateSubmitted, dataJson)
	if err != nil {
		return err
	}

	return nil
}
