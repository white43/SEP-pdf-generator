package jobs

import (
	"database/sql"
	"github.com/white43/sep401/pkg/database"
	"github.com/white43/sep401/pkg/errors"
)

type JobRepository struct {
	database *database.Database
}

func NewJobRepository(db *database.Database) *JobRepository {
	return &JobRepository{db}
}

func (jr JobRepository) GetByID(id string) (Job, error) {
	stmt, err := jr.database.Prepare("SELECT * FROM jobs WHERE id = ?")
	if err != nil {
		return Job{}, err
	}

	var e Job

	err = stmt.QueryRow(id).Scan(&e.ID, &e.Payload, &e.Result, &e.Status, &e.Type, &e.UserId)
	if err == sql.ErrNoRows {
		return Job{}, errors.JobNotFound
	} else if err != nil {
		return Job{}, err
	}

	return e, nil
}

func (jr JobRepository) InsertJob(id, payload string, userId int, jobType string) (Job, error) {
	stmt, err := jr.database.Prepare("INSERT INTO jobs (id, payload, status, type, user_id) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return Job{}, err
	}

	_, err = stmt.Exec(id, payload, "pending", jobType, userId)
	if err != nil {
		return Job{}, err
	}

	return jr.GetByID(id)
}

func (jr JobRepository) GetNextJob() (Job, error) {
	row := jr.database.QueryRow("SELECT * FROM jobs WHERE status = \"pending\"")

	var e Job

	if err := row.Scan(&e.ID, &e.Payload, &e.Result, &e.Status, &e.Type, &e.UserId); err != nil {
		return Job{}, err
	}

	return e, nil
}

func (jr JobRepository) UpdateJobStatus(job Job, status string) error {
	stmt, err := jr.database.Prepare("UPDATE jobs SET status = ? WHERE id = ?")
	if err != nil {
		return err
	}

	result, err := stmt.Exec(status, job.ID)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.NoRowsAffected
	}

	return nil
}

func (jr JobRepository) UpdateJobResult(job Job, result string) error {
	stmt, err := jr.database.Prepare("UPDATE jobs SET result = ? WHERE id = ?")
	if err != nil {
		return err
	}

	rslt, err := stmt.Exec(result, job.ID)
	if err != nil {
		return err
	}

	affected, err := rslt.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.NoRowsAffected
	}

	return nil
}
