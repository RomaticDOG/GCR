package validation

import (
	"context"

	"github.com/RomaticDOG/GCR/FastGO/internal/pkg/errorsx"
	apiV1 "github.com/RomaticDOG/GCR/FastGO/pkg/api/v1"
)

func (v *Validator) ValidateCreatePostReq(ctx context.Context, req *apiV1.CreatePostReq) error {
	if req.Title == "" {
		return errorsx.ErrTitleEmpty
	}
	if req.Content == "" {
		return errorsx.ErrContentEmpty
	}
	return nil
}

func (v *Validator) ValidateUpdatePostReq(ctx context.Context, req *apiV1.UpdatePostReq) error {
	return nil
}
