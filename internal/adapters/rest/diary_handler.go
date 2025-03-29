package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kritpi/arom-web-services/domain/requests"
	"github.com/kritpi/arom-web-services/domain/usecases"
)

type DiaryHandler interface {
	CreateDiary(c *fiber.Ctx) error
	GetAllDiary(c *fiber.Ctx) error
	GetDiaryByDate(c *fiber.Ctx) error
	GetDiaryByID(c *fiber.Ctx) error
	GetDiaryByUserID(c *fiber.Ctx) error
	UpdateDiary(c *fiber.Ctx) error
}

type diaryHandler struct {
	service usecases.DiaryUseCase
}

// CreateDiary implements DiaryHandler.
func (d *diaryHandler) CreateDiary(c *fiber.Ctx) error {
	var req requests.CreateDiaryRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse form data"})
	}

	files := form.File["images"]

	req.UserID = form.Value["user_id"][0]
	diary, err := d.service.CreateDiary(c.Context(), &req, files)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(diary)
}

// GetAllDiary implements DiaryHandler.
func (d *diaryHandler) GetAllDiary(c *fiber.Ctx) error {
	diaries, err := d.service.GetAllDiary(c.Context())
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Diary not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(diaries)
}

// GetDiaryByDate implements DiaryHandler.
func (d *diaryHandler) GetDiaryByDate(c *fiber.Ctx) error {
	date := c.Params("date")
	diaries, err := d.service.GetDiaryByDate(c.Context(), date)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Diary not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(diaries)
}

// GetDiaryByID implements DiaryHandler.
func (d *diaryHandler) GetDiaryByID(c *fiber.Ctx) error {
	id := c.Params("id")
	diary, err := d.service.GetDiaryByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Diary not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(diary)
}

// GetDiaryByUserID implements DiaryHandler.
func (d *diaryHandler) GetDiaryByUserID(c *fiber.Ctx) error {
	userID := c.Params("userID")
	diaries, err := d.service.GetDiaryByUserID(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Diary not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(diaries)
}

// UpdateDiary implements DiaryHandler.
func (d *diaryHandler) UpdateDiary(c *fiber.Ctx) error {
	date := c.Params("date")
	var req requests.UpdateDiaryRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse form data"})
	}
	req.ImagesURL = form.Value["images_url"]
	files := form.File["images"]

	req.UserID = form.Value["user_id"][0]
	errUpdate := d.service.UpdateDiary(c.Context(), &req, date, files)
	if errUpdate != nil {
		return errUpdate
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Diary updated successfully",
	})
}

func NewDiaryHandler(service usecases.DiaryUseCase) DiaryHandler {
	return &diaryHandler{
		service: service,
	}
}
