package helper

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, errCode *fiber.Error, internalCode int, message string) (err error) {
	// Return HTTP response
	ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
	ctx.Status(errCode.Code)

	// Render default error view
	err = ctx.Render("errors/"+strconv.Itoa(errCode.Code), fiber.Map{"message": message})
	if err != nil {
		return ctx.JSON(fiber.Map{
			"status":  errCode.Code,
			"error":   errCode.Message,
			"message": message,
			"code":    internalCode,
		})
	}
	return err
}
