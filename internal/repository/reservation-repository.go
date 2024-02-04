package repository

import (
	"time"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
)

func (r *postgresRepository) InsertReservation(reservation models.Reservation) (int, error) {
	var reservationID int

	query := `
						INSERT INTO reservations (first_name, last_name, email, phone, start_date, end_date, room_id, created_at, updated_at)
						VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
						RETURNING id
					`

	err := r.DB.SQL.QueryRow(
		query,
		reservation.FirstName,
		reservation.LastName,
		reservation.Email,
		reservation.Phone,
		reservation.StartDate,
		reservation.EndDate,
		reservation.RoomID,
		time.Now(),
		time.Now()).Scan(&reservationID)
	if err != nil {
		return 0, err
	}

	return reservationID, nil
}

func (r *postgresRepository) GetAllReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation

	query := `
		SELECT 
			re.id,
			re.first_name,
			re.last_name,
			re.email,
			re.phone,
			re.start_date,
			re.end_date,
			re.room_id,
			ro.id,
			ro.room_name,
			re.processed
		FROM reservations re
		LEFT JOIN rooms ro ON (ro.id = re.room_id)
		ORDER BY re.start_date asc
	`

	rows, err := r.DB.SQL.Query(query)
	if err != nil {
		return reservations, err
	}

	defer rows.Close()

	for rows.Next() {
		var i models.Reservation
		err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Phone,
			&i.StartDate,
			&i.EndDate,
			&i.RoomID,
			&i.ID,
			&i.Room.RoomName,
			&i.Processed,
		)
		if err != nil {
			return reservations, err
		}

		reservations = append(reservations, i)
	}

	return reservations, nil
}

func (r *postgresRepository) GetAllNewReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation

	query := `
		SELECT 
			re.id,
			re.first_name,
			re.last_name,
			re.email,
			re.phone,
			re.start_date,
			re.end_date,
			re.room_id,
			ro.id,
			ro.room_name,
			re.processed
		FROM reservations re
		LEFT JOIN rooms ro ON (ro.id = re.room_id)
		WHERE re.processed = 0
		ORDER BY re.start_date asc
	`

	rows, err := r.DB.SQL.Query(query)
	if err != nil {
		return reservations, err
	}

	defer rows.Close()

	for rows.Next() {
		var i models.Reservation
		err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Phone,
			&i.StartDate,
			&i.EndDate,
			&i.RoomID,
			&i.ID,
			&i.Room.RoomName,
			&i.Processed,
		)
		if err != nil {
			return reservations, err
		}

		reservations = append(reservations, i)
	}

	return reservations, nil
}
