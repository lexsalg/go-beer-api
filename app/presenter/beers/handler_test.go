package beers

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestPresenter_Register(t *testing.T) {
	type args struct {
		fiberContext *fiber.Ctx
	}
	tests := []struct {
		name    string
		handler *Presenter
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.handler.Register(tt.args.fiberContext); (err != nil) != tt.wantErr {
				t.Errorf("Presenter.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPresenter_Update(t *testing.T) {
	type args struct {
		fiberContext *fiber.Ctx
	}
	tests := []struct {
		name    string
		handler *Presenter
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.handler.Update(tt.args.fiberContext); (err != nil) != tt.wantErr {
				t.Errorf("Presenter.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPresenter_All(t *testing.T) {
	type args struct {
		fiberContext *fiber.Ctx
	}
	tests := []struct {
		name    string
		handler *Presenter
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.handler.All(tt.args.fiberContext); (err != nil) != tt.wantErr {
				t.Errorf("Presenter.All() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPresenter_GetById(t *testing.T) {
	type args struct {
		fiberContext *fiber.Ctx
	}
	tests := []struct {
		name    string
		handler *Presenter
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.handler.GetById(tt.args.fiberContext); (err != nil) != tt.wantErr {
				t.Errorf("Presenter.GetById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPresenter_BeerBox(t *testing.T) {
	type args struct {
		fiberContext *fiber.Ctx
	}
	tests := []struct {
		name    string
		handler *Presenter
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.handler.BeerBox(tt.args.fiberContext); (err != nil) != tt.wantErr {
				t.Errorf("Presenter.BeerBox() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
