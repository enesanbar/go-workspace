package dbrepo

import (
	"context"
	"log"
	"time"

	"github.com/enesanbar/workspace/golang/projects/vigilate/internal/models"
)

// InsertHost inserts a host into the database
func (m *postgresDBRepo) InsertHost(host models.Host) (int, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	query := `
		insert into hosts (host_name, canonical_name, url, ip, ipv6, location, os, active, created_at, updated_at) 
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) returning id
	`

	var newID int
	err := m.DB.QueryRowContext(ctx, query,
		host.HostName,
		host.CanonicalName,
		host.URL,
		host.IP,
		host.IPV6,
		host.Location,
		host.OS,
		host.Active,
		time.Now(),
		time.Now(),
	).Scan(&newID)
	if err != nil {
		log.Println(err)
		return newID, err
	}

	query = "select id from services"
	serviceRows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer serviceRows.Close()

	for serviceRows.Next() {
		var svcID int
		err := serviceRows.Scan(&svcID)
		if err != nil {
			log.Println(err)
			return 0, err
		}

		stmt := `
		insert into host_services 
		    (host_id, service_id, active, schedule_number, schedule_unit, status, created_at, updated_at)
		values ($1, $2, 0, 3, 'm', 'pending', $3, $4)
	`

		_, err = m.DB.ExecContext(ctx, stmt,
			newID,
			svcID,
			time.Now(),
			time.Now(),
		)
		if err != nil {
			return newID, err
		}
	}

	return newID, nil
}

// GetHosts get all host from the database
func (m *postgresDBRepo) GetHosts() ([]*models.Host, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	var hosts []*models.Host

	query := `
		select id, host_name, canonical_name, url, ip, ipv6, location, os, active, created_at, updated_at from hosts
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return hosts, err
	}
	defer rows.Close()

	for rows.Next() {
		var host models.Host
		err := rows.Scan(
			&host.ID,
			&host.HostName,
			&host.CanonicalName,
			&host.URL,
			&host.IP,
			&host.IPV6,
			&host.Location,
			&host.OS,
			&host.Active,
			&host.CreatedAt,
			&host.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		serviceQuery := `
		select 
		       hs.id, hs.host_id, hs.service_id, hs.active, hs.schedule_number, hs.schedule_unit, 
		       hs.last_check, hs.last_message, hs.status, hs.created_at, hs.updated_at,
		       s.id, s.service_name, s.active, s.icon, s.created_at, s.updated_at
		from host_services hs
		left join services s on (s.id = hs.service_id)
		where host_id = $1
		`

		serviceRows, err := m.DB.QueryContext(ctx, serviceQuery, host.ID)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		var hostServices []models.HostService
		for serviceRows.Next() {
			var hs models.HostService
			err := serviceRows.Scan(
				&hs.ID,
				&hs.HostID,
				&hs.ServiceID,
				&hs.Active,
				&hs.ScheduleNumber,
				&hs.ScheduleUnit,
				&hs.LastCheck,
				&hs.LastMessage,
				&hs.Status,
				&hs.CreatedAt,
				&hs.UpdatedAt,
				&hs.Service.ID,
				&hs.Service.ServiceName,
				&hs.Service.Active,
				&hs.Service.Icon,
				&hs.Service.CreatedAt,
				&hs.Service.UpdatedAt,
			)
			if err != nil {
				log.Println(err)
				return nil, err
			}

			hostServices = append(hostServices, hs)
		}
		serviceRows.Close()

		host.HostServices = hostServices
		hosts = append(hosts, &host)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
		return hosts, err
	}

	if err != nil {
		log.Println(err)
		return hosts, err
	}

	return hosts, nil
}

// GetHostByID get a host from the database
func (m *postgresDBRepo) GetHostByID(id int) (models.Host, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	query := `
		select 
		       id, host_name, canonical_name, url, ip, ipv6, location, os, active, created_at, updated_at 
		from hosts
		where id = $1
	`

	var host models.Host
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&host.ID,
		&host.HostName,
		&host.CanonicalName,
		&host.URL,
		&host.IP,
		&host.IPV6,
		&host.Location,
		&host.OS,
		&host.Active,
		&host.CreatedAt,
		&host.UpdatedAt,
	)

	if err != nil {
		log.Println(err)
		return host, err
	}

	// get all services for host
	query = `
		select 
		       hs.id, hs.host_id, hs.service_id, hs.active, hs.schedule_number, hs.schedule_unit, 
		       hs.last_check, hs.last_message, hs.status, hs.created_at, hs.updated_at,
		       s.id, s.service_name, s.active, s.icon, s.created_at, s.updated_at
		from host_services hs
		left join services s on (s.id = hs.service_id)
		where host_id = $1
		order by s.service_name
	`

	rows, err := m.DB.QueryContext(ctx, query, host.ID)
	if err != nil {
		log.Println(err)
		return models.Host{}, err
	}
	defer rows.Close()

	var hostServices []models.HostService
	for rows.Next() {
		var hs models.HostService
		err := rows.Scan(
			&hs.ID,
			&hs.HostID,
			&hs.ServiceID,
			&hs.Active,
			&hs.ScheduleNumber,
			&hs.ScheduleUnit,
			&hs.LastCheck,
			&hs.LastMessage,
			&hs.Status,
			&hs.CreatedAt,
			&hs.UpdatedAt,
			&hs.Service.ID,
			&hs.Service.ServiceName,
			&hs.Service.Active,
			&hs.Service.Icon,
			&hs.Service.CreatedAt,
			&hs.Service.UpdatedAt,
		)
		if err != nil {
			log.Println(err)
			return models.Host{}, err
		}

		hostServices = append(hostServices, hs)
	}

	host.HostServices = hostServices
	return host, nil

}

func (m *postgresDBRepo) UpdateHost(host models.Host) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	query := `
		update hosts set host_name = $1, canonical_name = $2, url = $3, ip = $4, ipv6 = $5, 
		                 location = $6, os = $7, active = $8, updated_at = $9 
		where id = $10
	`

	_, err := m.DB.ExecContext(ctx, query,
		host.HostName,
		host.CanonicalName,
		host.URL,
		host.IP,
		host.IPV6,
		host.Location,
		host.OS,
		host.Active,
		time.Now(),
		host.ID,
	)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// UpdateHostServiceStatus updates the active status of host service
func (m *postgresDBRepo) UpdateHostServiceStatus(hostID, serviceID, active int) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	query := "update host_services set active = $1 where host_id = $2 and service_id = $3"

	_, err := m.DB.ExecContext(ctx, query, active, hostID, serviceID)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// UpdateHostService updates a host service
func (m *postgresDBRepo) UpdateHostService(hs models.HostService) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	query := `
		update 
		    host_services 
		set 
		    host_id = $1, service_id = $2, active = $3, schedule_number = $4, schedule_unit = $5,
		    last_check = $6, last_message = $7, status = $8, updated_at = $9
		where id = $10
	`

	_, err := m.DB.ExecContext(ctx, query,
		hs.HostID,
		hs.ServiceID,
		hs.Active,
		hs.ScheduleNumber,
		hs.ScheduleUnit,
		hs.LastCheck,
		hs.LastMessage,
		hs.Status,
		time.Now(),
		hs.ID,
	)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (m *postgresDBRepo) GetAllServiceStatusCounts() (int, int, int, int, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	//query := "select status, count(id) from host_services where active = 1 group by status"
	query := `
		select
			(select count(id) from host_services where active = 1 and status = 'pending') as pending,
			(select count(id) from host_services where active = 1 and status = 'healthy') as healthy,
			(select count(id) from host_services where active = 1 and status = 'warning') as warning,
			(select count(id) from host_services where active = 1 and status = 'problem') as problem
	`
	var pending, healthy, warning, problem int
	err := m.DB.QueryRowContext(ctx, query).Scan(
		&pending,
		&healthy,
		&warning,
		&problem,
	)
	if err != nil {
		return 0, 0, 0, 0, err
	}

	return pending, healthy, warning, problem, err
}

func (m *postgresDBRepo) GetServicesByStatus(status string) ([]models.HostService, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	query := `
		select 
		       hs.id, hs.host_id, hs.service_id, hs.active, hs.schedule_number, hs.schedule_unit, 
		       hs.last_check, hs.last_message, hs.status, hs.created_at, hs.updated_at,
		       h.host_name,
		       s.service_name
		from host_services hs
		left join hosts h on h.id = hs.host_id
		left join services s on hs.service_id = s.id
		where status = $1 and hs.active = 1
		order by host_name, service_name;
	`

	rows, err := m.DB.QueryContext(ctx, query, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hostServices []models.HostService
	for rows.Next() {
		var hs models.HostService
		err := rows.Scan(
			&hs.ID,
			&hs.HostID,
			&hs.ServiceID,
			&hs.Active,
			&hs.ScheduleNumber,
			&hs.ScheduleUnit,
			&hs.LastCheck,
			&hs.LastMessage,
			&hs.Status,
			&hs.CreatedAt,
			&hs.UpdatedAt,
			&hs.HostName,
			&hs.Service.ServiceName,
		)
		if err != nil {
			return nil, err
		}
		hostServices = append(hostServices, hs)
	}

	return hostServices, nil
}

func (m *postgresDBRepo) GetHostServiceByID(id int) (models.HostService, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	query := `
		select 
		       hs.id, hs.host_id, hs.service_id, hs.active, hs.schedule_number, hs.schedule_unit, 
		       hs.last_check, hs.last_message, hs.status, hs.created_at, hs.updated_at,
		       s.id, s.service_name, s.active, s.icon, s.created_at, s.updated_at,
		       h.host_name
		from host_services hs
		left join services s on (s.id = hs.service_id)
		left join hosts h on (h.id = hs.host_id)
		where hs.id = $1
	`

	var hs models.HostService
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&hs.ID,
		&hs.HostID,
		&hs.ServiceID,
		&hs.Active,
		&hs.ScheduleNumber,
		&hs.ScheduleUnit,
		&hs.LastCheck,
		&hs.LastMessage,
		&hs.Status,
		&hs.CreatedAt,
		&hs.UpdatedAt,
		&hs.Service.ID,
		&hs.Service.ServiceName,
		&hs.Service.Active,
		&hs.Service.Icon,
		&hs.Service.CreatedAt,
		&hs.Service.UpdatedAt,
		&hs.HostName,
	)
	if err != nil {
		log.Println(err)
		return models.HostService{}, err
	}

	return hs, nil
}

func (m *postgresDBRepo) GetServicesToMonitor() ([]models.HostService, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	query := `
		select 
		       hs.id, hs.host_id, hs.service_id, hs.active, hs.schedule_number, hs.schedule_unit, 
		       hs.last_check, hs.last_message, hs.status, hs.created_at, hs.updated_at,
		       s.id, s.service_name, s.active, s.icon, s.created_at, s.updated_at,
		       h.host_name
		from host_services hs
		left join services s on (s.id = hs.service_id)
		left join hosts h on (h.id = hs.host_id)
		where h.active = 1 and hs.active = 1
	`

	var hostServices []models.HostService
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		log.Println(err)
		return hostServices, err
	}
	defer rows.Close()

	for rows.Next() {
		var hostService models.HostService
		err := rows.Scan(
			&hostService.ID,
			&hostService.HostID,
			&hostService.ServiceID,
			&hostService.Active,
			&hostService.ScheduleNumber,
			&hostService.ScheduleUnit,
			&hostService.LastCheck,
			&hostService.LastMessage,
			&hostService.Status,
			&hostService.CreatedAt,
			&hostService.UpdatedAt,
			&hostService.Service.ID,
			&hostService.Service.ServiceName,
			&hostService.Service.Active,
			&hostService.Service.Icon,
			&hostService.Service.CreatedAt,
			&hostService.Service.UpdatedAt,
			&hostService.HostName,
		)
		if err != nil {
			log.Println(err)
			return hostServices, err
		}

		hostServices = append(hostServices, hostService)
	}

	return hostServices, nil
}

func (m *postgresDBRepo) GetHostServiceByHostIDServiceID(hostID int, serviceID int) (models.HostService, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	query := `
		select 
			hs.id, hs.host_id, hs.service_id, hs.active, hs.schedule_number, hs.schedule_unit, 
		       hs.last_check, hs.last_message, hs.status, hs.created_at, hs.updated_at,
		       s.id, s.service_name, s.active, s.icon, s.created_at, s.updated_at, 
		       h.host_name
		from host_services hs
		left join services s on (s.id = hs.service_id)
		left join hosts h on (h.id = hs.host_id)
		where hs.host_id = $1 and hs.service_id = $2
	`

	var hs models.HostService
	err := m.DB.QueryRowContext(ctx, query, hostID, serviceID).Scan(
		&hs.ID,
		&hs.HostID,
		&hs.ServiceID,
		&hs.Active,
		&hs.ScheduleNumber,
		&hs.ScheduleUnit,
		&hs.LastCheck,
		&hs.LastMessage,
		&hs.Status,
		&hs.CreatedAt,
		&hs.UpdatedAt,
		&hs.Service.ID,
		&hs.Service.ServiceName,
		&hs.Service.Active,
		&hs.Service.Icon,
		&hs.Service.CreatedAt,
		&hs.Service.UpdatedAt,
		&hs.HostName,
	)
	if err != nil {
		log.Println(err)
		return hs, err
	}

	return hs, nil
}
