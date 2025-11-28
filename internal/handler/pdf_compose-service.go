package handler

import (
	"bytes"
	"context"
	"io"

	"service-pdf-compose/api"
	"service-pdf-compose/pkg/composer"
)

type PdfComposeServer struct {
	api.UnimplementedPdfComposeServiceServer
}

func (s *PdfComposeServer) ComposeFromFiles(ctx context.Context, req *api.ComposeRequest) (*api.ComposeResponse, error) {
	readers := []io.ReadCloser{}

	if len(req.Upfile1) > 0 {
		readers = append(readers, io.NopCloser(bytes.NewReader(req.Upfile1)))
	}
	if len(req.Upfile2) > 0 {
		readers = append(readers, io.NopCloser(bytes.NewReader(req.Upfile2)))
	}
	if len(req.Upfile3) > 0 {
		readers = append(readers, io.NopCloser(bytes.NewReader(req.Upfile3)))
	}

	pdfReader, err := composer.ComposeFromFiles(readers)
	if err != nil {
		return nil, err
	}

	pdfBytes, _ := io.ReadAll(pdfReader)
	return &api.ComposeResponse{PdfFile: pdfBytes}, nil
}
