package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
	"time"
)

type ReportService struct {
	repo *repositories.ReportRepository
}

func NewReportService(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetDailyReport() (*models.DailyReport, error) {
	now := time.Now()
	return s.repo.GetDailyReport(now)
}

func (s *ReportService) GetRangeReport(startDate, endDate time.Time) (*models.DailyReport, error) {
	return s.repo.GetRangeReport(startDate, endDate)
}
