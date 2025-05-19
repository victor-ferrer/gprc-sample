package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/victor-ferrer/gprc-sample/internal/models"
)

type TicketRepository struct {
	db *sql.DB
}

func NewTicketRepository(db *sql.DB) *TicketRepository {
	return &TicketRepository{
		db: db,
	}
}
func (tr *TicketRepository) CreateTicket(ticket *models.Ticket) error {
	_, err := tr.db.Exec("INSERT INTO tickets (created_at, updated_at, purchase_date, amount, currency, labels, file) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		ticket.CreatedAt,
		ticket.UpdatedAt,
		ticket.PurchaseDate,
		ticket.Amount,
		ticket.Currency,
		strings.Join(ticket.Labels, ","),
		ticket.File,
	)
	if err != nil {
		return fmt.Errorf("failed to create ticket: %w", err)
	}
	return nil
}
func (tr *TicketRepository) GetTicket(id int) (*models.Ticket, error) {
	row := tr.db.QueryRow("SELECT id, created_at, updated_at, purchase_date, amount, currency, labels, file FROM tickets WHERE id = $1", id)
	ticket := &models.Ticket{}
	err := row.Scan(&ticket.ID, &ticket.CreatedAt, &ticket.UpdatedAt, &ticket.PurchaseDate, &ticket.Amount, &ticket.Currency, &ticket.Labels, &ticket.File)
	if err != nil {
		return nil, fmt.Errorf("failed to get ticket: %w", err)
	}
	return ticket, nil
}
