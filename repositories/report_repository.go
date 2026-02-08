package repositories

import (
	"database/sql"
	"kasir-api/models"
	"time"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (repo *ReportRepository) GetDailyReport(date time.Time) (*models.DailyReport, error) {
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	var totalRevenue, totalTransaksi int

	err := repo.db.QueryRow(`
		SELECT 
			COALESCE(SUM(total_amount), 0) as total_revenue,
			COUNT(*) as total_transaksi
		FROM transactions
		WHERE created_at >= $1 AND created_at < $2
	`, startOfDay, endOfDay).Scan(&totalRevenue, &totalTransaksi)

	if err != nil {
		return nil, err
	}

	var bestProduct *models.BestProduct
	var nama string
	var qtyTerjual int

	err = repo.db.QueryRow(`
		SELECT 
			p.name,
			SUM(td.quantity) as qty_terjual
		FROM transaction_details td
		JOIN transactions t ON td.transaction_id = t.id
		JOIN products p ON td.product_id = p.id
		WHERE t.created_at >= $1 AND t.created_at < $2
		GROUP BY p.id, p.name
		ORDER BY qty_terjual DESC
		LIMIT 1
	`, startOfDay, endOfDay).Scan(&nama, &qtyTerjual)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == nil {
		bestProduct = &models.BestProduct{
			Nama:       nama,
			QtyTerjual: qtyTerjual,
		}
	}

	return &models.DailyReport{
		TotalRevenue:   totalRevenue,
		TotalTransaksi: totalTransaksi,
		ProdukTerlaris: bestProduct,
	}, nil
}

func (repo *ReportRepository) GetRangeReport(startDate, endDate time.Time) (*models.DailyReport, error) {
	var totalRevenue, totalTransaksi int

	err := repo.db.QueryRow(`
		SELECT 
			COALESCE(SUM(total_amount), 0) as total_revenue,
			COUNT(*) as total_transaksi
		FROM transactions
		WHERE created_at >= $1 AND created_at < $2
	`, startDate, endDate).Scan(&totalRevenue, &totalTransaksi)

	if err != nil {
		return nil, err
	}

	var bestProduct *models.BestProduct
	var nama string
	var qtyTerjual int

	err = repo.db.QueryRow(`
		SELECT 
			p.name,
			SUM(td.quantity) as qty_terjual
		FROM transaction_details td
		JOIN transactions t ON td.transaction_id = t.id
		JOIN products p ON td.product_id = p.id
		WHERE t.created_at >= $1 AND t.created_at < $2
		GROUP BY p.id, p.name
		ORDER BY qty_terjual DESC
		LIMIT 1
	`, startDate, endDate).Scan(&nama, &qtyTerjual)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == nil {
		bestProduct = &models.BestProduct{
			Nama:       nama,
			QtyTerjual: qtyTerjual,
		}
	}

	return &models.DailyReport{
		TotalRevenue:   totalRevenue,
		TotalTransaksi: totalTransaksi,
		ProdukTerlaris: bestProduct,
	}, nil
}
