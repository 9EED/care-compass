package services

import (
	"github.com/lai0xn/hackiwna-backend/internal/models"
	"github.com/lai0xn/hackiwna-backend/internal/storage"
	"github.com/lai0xn/hackiwna-backend/internal/transport/types"
	uuid "github.com/satori/go.uuid"
)

type AppointmentService struct{}

func (a *AppointmentService) GetAppointments(id uuid.UUID) []models.Appointment {
	var appointments []models.Appointment
	storage.DB.Preload("Doctor").
		Preload("Patient").
		Preload("Doctor").
		Where("patient_id = ?", id).
		Find(&appointments)
	return appointments
}

func (a *AppointmentService) CreateApointment(payload types.AppointmentPayload) error {
	db := storage.DB.Create(&models.Appointment{
		DoctorID:  payload.DoctorID,
		PatientID: payload.PatientID,
		Date:      payload.Date,
	})
	if db.Error != nil {
		panic(db.Error)
	}
	return nil
}
