package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"context"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	DB *sql.DB
}

func (h *Handler) HandleCreatePingRequest(c *fiber.Ctx) error {

	_, err := h.DB.ExecContext(context.Background(), "insert into ping (create_at) values (now())")

	if err != nil {
		c.SendString(fmt.Sprintf("Failed to insert into the database: %w", err))
		return err
	}

	c.JSON(http.StatusOK)
	return nil
}

func (h *Handler) HandleListPingRequest(c *fiber.Ctx) error {

	rows, err := h.DB.Query("select * from ping")

	if err != nil {
		// it is not a good practice to expose the error message to the client
		c.SendString(fmt.Sprintf("Failed to query the ping table: %w", err))
		return err
	}

	defer rows.Close()

	var pings []map[string]interface{}
	for rows.Next() {
		var id int
		var createAt string
		if err := rows.Scan(&id, &createAt); err != nil {
			c.SendString(fmt.Sprintf("Failed to scan row: %v", err))
			return err
		}
		pings = append(pings, map[string]interface{}{
			"id":        id,
			"create_at": createAt,
		})
	}

	if err := rows.Err(); err != nil {
		c.SendString(fmt.Sprintf("Error iterating rows: %v", err))
		return err
	}

	c.JSON(pings)
	return nil

}
