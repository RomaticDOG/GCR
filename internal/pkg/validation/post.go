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

// ValidateUpdatePostReq 更新博文请求体校验
func (v *Validator) ValidateUpdatePostReq(ctx context.Context, req *apiV1.UpdatePostReq) error {
	return nil
}

// ValidateDeletePostReq 删除博文请求体校验
func (v *Validator) ValidateDeletePostReq(ctx context.Context, req *apiV1.DeletePostReq) error {
	return nil
}

// ValidateGetPostReq 获取博文请求体校验
func (v *Validator) ValidateGetPostReq(ctx context.Context, req *apiV1.GetPostReq) error {
	return nil
}

// ValidateListPostReq 获取博文列表请求体校验
func (v *Validator) ValidateListPostReq(ctx context.Context, req *apiV1.ListPostReq) error {
	return nil
}
