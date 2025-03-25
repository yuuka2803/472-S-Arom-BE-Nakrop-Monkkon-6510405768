package rest

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/kritpi/arom-web-services/domain/usecases"
)

func Test_diaryHandler_CreateDiary(t *testing.T) {
	type fields struct {
		service usecases.DiaryUseCase
	}
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &diaryHandler{
				service: tt.fields.service,
			}
			if err := d.CreateDiary(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("diaryHandler.CreateDiary() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_diaryHandler_GetAllDiary(t *testing.T) {
	type fields struct {
		service usecases.DiaryUseCase
	}
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &diaryHandler{
				service: tt.fields.service,
			}
			if err := d.GetAllDiary(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("diaryHandler.GetAllDiary() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_diaryHandler_GetDiaryByDate(t *testing.T) {
	type fields struct {
		service usecases.DiaryUseCase
	}
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &diaryHandler{
				service: tt.fields.service,
			}
			if err := d.GetDiaryByDate(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("diaryHandler.GetDiaryByDate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_diaryHandler_GetDiaryByID(t *testing.T) {
	type fields struct {
		service usecases.DiaryUseCase
	}
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &diaryHandler{
				service: tt.fields.service,
			}
			if err := d.GetDiaryByID(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("diaryHandler.GetDiaryByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_diaryHandler_GetDiaryByUserID(t *testing.T) {
	type fields struct {
		service usecases.DiaryUseCase
	}
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &diaryHandler{
				service: tt.fields.service,
			}
			if err := d.GetDiaryByUserID(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("diaryHandler.GetDiaryByUserID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_diaryHandler_UpdateDiary(t *testing.T) {
	type fields struct {
		service usecases.DiaryUseCase
	}
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &diaryHandler{
				service: tt.fields.service,
			}
			if err := d.UpdateDiary(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("diaryHandler.UpdateDiary() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
