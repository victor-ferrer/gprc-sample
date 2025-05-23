package repository

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/victor-ferrer/gprc-sample/internal/models"
)

var testDB *sql.DB
var ticketRepo *TicketRepository

func TestMain(m *testing.M) {
	// Setup in-memory SQLite database
	var err error
	testDB, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("failed to open test database: %v", err)
	}

	// Create the Ticket table
	_, err = testDB.Exec(`
		CREATE TABLE Ticket (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			created_at DATETIME,
			updated_at DATETIME,
			purchase_date DATETIME,
			amount REAL,
			currency TEXT,
			labels TEXT,
			file TEXT
		)
	`)
	if err != nil {
		log.Fatalf("failed to create Ticket table: %v", err)
	}

	// Initialize repository
	ticketRepo = NewTicketRepository(testDB)

	// Run tests
	code := m.Run()

	// Cleanup
	testDB.Close()

	os.Exit(code)
}

func TestCreateTicket(t *testing.T) {
	ticket := &models.Ticket{
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		PurchaseDate: time.Now(),
		Amount:       100.50,
		Currency:     "USD",
		Labels:       []string{"concert", "vip"},
		File:         "ticket.pdf",
	}

	err := ticketRepo.CreateTicket(ticket)
	if err != nil {
		t.Fatalf("failed to create ticket: %v", err)
	}

	var count int
	err = testDB.QueryRow("SELECT COUNT(*) FROM Ticket").Scan(&count)
	if err != nil {
		t.Fatalf("failed to count tickets: %v", err)
	}

	if count != 1 {
		t.Fatalf("expected 1 ticket, got %d", count)
	}
}

func TestGetTicket(t *testing.T) {
	// Insert a ticket directly into the database
	_, err := testDB.Exec(`
		INSERT INTO Ticket (created_at, updated_at, purchase_date, amount, currency, labels, file)
		VALUES (?, ?, ?, ?, ?, ?, ?)`,
		time.Now(), time.Now(), time.Now(), 50.00, "EUR", "event,standard", "event.pdf",
	)
	if err != nil {
		t.Fatalf("failed to insert ticket: %v", err)
	}

	// Retrieve the ticket
	ticket, err := ticketRepo.GetTicket(1)
	if err != nil {
		t.Fatalf("failed to get ticket: %v", err)
	}

	if ticket.Amount != 50.00 || ticket.Currency != "EUR" || ticket.File != "event.pdf" {
		t.Fatalf("retrieved ticket does not match expected values")
	}
}
