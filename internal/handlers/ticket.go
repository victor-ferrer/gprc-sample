package handlers

import (
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/victor-ferrer/gprc-sample/internal/models"
)

type TicketHandler struct {
	tr TicketRepository
}

type TicketRepository interface {
	CreateTicket(ticket *models.Ticket) error
	GetTicket(id int) (*models.Ticket, error)
}

func NewTicketHandler(tr TicketRepository) *TicketHandler {
	return &TicketHandler{
		tr: tr,
	}
}

func (t *TicketHandler) UploadTicket(c *fiber.Ctx) error {
	// Parse the form data
	//	if err := c.Request().ParseMultipartForm(10 << 20); err != nil {
	//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	//			"error": "Failed to parse form data",
	//		})
	//	}

	// Get the file from the form data
	file, err := c.FormFile("ticket")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to get file from form data",
		})
	}

	// Optional metadata
	labels := c.FormValue("labels")
	purchaseDateStr := c.FormValue("purchase_date")
	amount := c.FormValue("amount")
	currency := c.FormValue("currency")

	amountFloat, _ := strconv.ParseFloat(amount, 64)
	purchaseDate, _ := time.Parse(time.RFC3339, purchaseDateStr)

	// Save the file to the server
	//err = c.SaveFile(file, fmt.Sprintf("./uploads/%s", file.Filename))
	//if err != nil {
	//	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	//		"error": "Failed to save file",
	//	})
	//}

	ticket := &models.Ticket{
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		PurchaseDate: purchaseDate,
		Amount:       amountFloat,
		Currency:     currency,
		Labels:       strings.Split(labels, ""),
		File:         file.Filename,
	}

	err = t.tr.CreateTicket(ticket)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save ticket to DB",
		})
	}

	return c.JSON(fiber.Map{
		"message": "File uploaded successfully",
	})

}
