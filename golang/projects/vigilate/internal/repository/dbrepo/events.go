package dbrepo

import (
	"context"
	"log"
	"time"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/models"
)

func (m *postgresDBRepo) InsertEvent(e models.Event) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	query := `
		insert into events (event_type, host_service_id, host_id, service_name, host_name, message, created_at, updated_at) 
		values ($1, $2, $3, $4, $5, $6, $7, $8);
	`

	_, err := m.DB.ExecContext(ctx, query,
		e.EventType,
		e.HostServiceID,
		e.HostID,
		e.ServiceName,
		e.HostName,
		e.Message,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (m *postgresDBRepo) GetAllEvents() ([]models.Event, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	query := `
		select 
			id, event_type, host_service_id, host_id, service_name, host_name, message, created_at, updated_at
		from events
		order by created_at
	`

	var events []models.Event
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		log.Println(err)
		return events, err
	}
	defer rows.Close()

	for rows.Next() {
		var event models.Event
		err := rows.Scan(
			&event.ID,
			&event.EventType,
			&event.HostServiceID,
			&event.HostID,
			&event.ServiceName,
			&event.HostName,
			&event.Message,
			&event.CreatedAt,
			&event.UpdatedAt,
		)
		if err != nil {
			log.Println(err)
			return events, err
		}

		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	return events, err
}
