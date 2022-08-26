package sqlstorage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/config"
	"github.com/g4web/otus_go/hw12_13_14_15_calendar/internal/storage"
	"github.com/jmoiron/sqlx"
	// PG driver.
	_ "github.com/lib/pq"
)

var ErrRowsAffected = errors.New("the number of affected rows is not equal to one")

type Storage struct {
	db  *sqlx.DB
	ctx context.Context
}

func New(c *config.Config) (*Storage, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", c.DBHost, c.DBUser, c.DBPassword, c.DBName)

	s := &Storage{}
	err := s.connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Storage) connect(ctx context.Context, dsn string) error {
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
		e.NotificationBefore().Seconds(),
	)

	return err
}

func (s *Storage) Update(eventID int32, e *storage.EventDTO) error {
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
		e.NotificationBefore().Seconds(),
	)

	return err
}

func (s *Storage) MarkNotificationAsSent(eventID int32) error {
	query := `
				UPDATE 
					event
				SET
				    notification_is_sent = true
				WHERE 
					id = $1
	`
	_, err := s.db.ExecContext(
		s.ctx,
		query,
		eventID,
	)

	return err
}

func (s *Storage) Delete(id int32) error {
	query := `
				DELETE
				FROM
					event
				WHERE
					id = $1
	`

	result, err := s.db.ExecContext(s.ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return ErrRowsAffected
	}

	return nil
}

func (s *Storage) DeleteOld(endDate time.Time) (int32, error) {
	query := `
				DELETE
				FROM
					event
				WHERE
				    start_date < $1
	`
	result, err := s.db.ExecContext(s.ctx, query, endDate)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return int32(rowsAffected), err
	}

	return int32(rowsAffected), nil
}

func (s *Storage) FindOneByID(eventID int32) (*storage.EventDTO, error) {
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
	row := s.db.QueryRowContext(s.ctx, query, eventID)

	var title, description string
	var id, userID, notificationBefore int32
	var startDate, endDate time.Time
	err := row.Scan(
		&id,
		&title,
		&description,
		&userID,
		&startDate,
		&endDate,
		&notificationBefore,
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
		userID,
		startDate,
		endDate,
		time.Duration(notificationBefore*1e9),
	), nil
}

func (s *Storage) FindListByPeriod(startDate time.Time, endDate time.Time, userID int32) ([]*storage.EventDTO, error) {
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

	return s.rowsToEvents(rows)
}

func (s *Storage) FindNotificationByPeriod(startDate time.Time, endDate time.Time) ([]*storage.EventDTO, error) {
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
		  notification_is_sent = false AND  
		  start_date - (notification_before || ' seconds')::interval between $1 AND $2
		;
	`
	rows, err := s.db.QueryxContext(s.ctx, sqlQuery, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return s.rowsToEvents(rows)
}

func (s *Storage) rowsToEvents(rows *sqlx.Rows) ([]*storage.EventDTO, error) {
	events := make([]*storage.EventDTO, 0)
	for rows.Next() {
		var event *storage.EventDTO

		var title, description string
		var id, userID, notificationBefore int32
		var startDate, endDate time.Time

		err := rows.Scan(
			&id,
			&title,
			&description,
			&userID,
			&startDate,
			&endDate,
			&notificationBefore,
		)
		if err != nil {
			return nil, err
		}

		event = storage.NewEventDTO(
			id,
			title,
			description,
			userID,
			startDate,
			endDate,
			time.Duration(notificationBefore*1e9),
		)
		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}
