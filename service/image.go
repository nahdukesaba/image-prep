package service

import (
	"context"
	"image-prep/service/request"
)

type image struct {
}

func NewImageService() ImageService {
	return &image{}
}

type ImageService interface {
	ConvertImage(ctx context.Context, form *request.ConvertImageRequest) error
	ResizeImage(ctx context.Context, form *request.ResizeImageRequest) error
	CompressImage(ctx context.Context, form *request.CompressImageRequest) error
}

func (c *image) ConvertImage(ctx context.Context, form *request.ConvertImageRequest) error {
	return nil
}

func (c *image) ResizeImage(ctx context.Context, form *request.ResizeImageRequest) error {
	return nil
}

func (c *image) CompressImage(ctx context.Context, form *request.CompressImageRequest) error {
	return nil
}
