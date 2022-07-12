package sqlstorage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/storage"

	"github.com/g4web/otus_go/hw12_13_14_15_calendar/configs"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var ErrRowsAffected = errors.New("The number of affected rows is not equal to one")

type Storage struct {
	db  *sqlx.DB
	ctx context.Context
}

func New(c *configs.Config) (*Storage, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", c.DbUser, c.DbPassword, c.DbName)

	s := &Storage{}
	err := s.connect(dsn, context.Background())
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Storage) connect(dsn string, ctx context.Context) error {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return err
	}

	s.db = db
	s.ctx = ctx

	return err
}

func (s *Storage) Insert(e *storage.EventDTO) error {
	query := `
				INSERT INTO event
					(user_id, title, description, start_date, end_date, notification_before)
				VALUES
					($1, $2, $3, $4, $5, $6)
				;
	`

	_, err := s.db.ExecContext(
		s.ctx,
		query,
		e.UserID(),
		e.Title(),
		e.Description(),
		e.StartDate().Round(time.Microsecond),
		e.EndDate().Round(time.Microsecond),
		e.NotificationBefore().Round(time.Second).Seconds(),
	)

	return err
}

func (s *Storage) Update(eventID int, e *storage.EventDTO) (bool, error) {
	query := `
				UPDATE 
					event
				SET
					user_id = $2,
					title = $3,
					description = $4, 
					start_date = $5, 
					end_date = $6,
				    notification_before = $7
				WHERE 
					id = $1
	`
	_, err := s.db.ExecContext(
		s.ctx,
		query,
		eventID,
		e.UserID(),
		e.Title(),
		e.Description(),
		e.StartDate(),
		e.EndDate(),
		int(e.NotificationBefore().Round(time.Second).Seconds()),
	)

	return err == nil, err
}

func (s *Storage) Delete(id int) (bool, error) {
	query := `
				DELETE
				FROM
					event
				WHERE
					id = $1
	`

	result, err := s.db.ExecContext(s.ctx, query, id)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	if rowsAffected != 1 {
		return false, ErrRowsAffected
	}

	return true, nil
}

func (s *Storage) FindOneById(eventId int) (*storage.EventDTO, error) {
	query := `
		SELECT
		 id,
		 title,
		 description,
		 user_id,
		 start_date,
		 end_date,
		 notification_before
		FROM
		  event
		WHERE
		  id = $1
		;
	`
	row := s.db.QueryRowContext(s.ctx, query, eventId)

	var title, description string
	var id, user_id, notification_before int
	var start_date, end_date time.Time
	err := row.Scan(
		&id,
		&title,
		&description,
		&user_id,
		&start_date,
		&end_date,
		&notification_before,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return storage.NewEventDTO(
		id,
		title,
		description,
		user_id,
		start_date,
		end_date,
		time.Duration(notification_before*1e9),
	), nil
}

func (s *Storage) FindListByPeriod(startDate time.Time, endDate time.Time, userID int) ([]*storage.EventDTO, error) {
	sqlQuery := `
		SELECT
		 id,
		 title,
		 description,
		 user_id,
		 start_date,
		 end_date,
		 notification_before
		FROM
		  event
		WHERE
		  $1 < end_date AND $2 > start_date
		  AND user_id = $3	
		;
	`
	rows, err := s.db.QueryxContext(s.ctx, sqlQuery, startDate, endDate, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := make([]*storage.EventDTO, 0)
	for rows.Next() {
		var event *storage.EventDTO

		var title, description string
		var id, user_id, notification_before int
		var start_date, end_date time.Time

		err = rows.Scan(
			&id,
			&title,
			&description,
			&user_id,
			&start_date,
			&end_date,
			&notification_before,
		)

		if err != nil {
			return nil, err
		}

		event = storage.NewEventDTO(
			id,
			title,
			description,
			user_id,
			start_date,
			end_date,
			time.Duration(notification_before*1e9),
		)
		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}
