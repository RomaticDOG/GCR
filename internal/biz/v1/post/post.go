package post

import (
	"context"
	"fmt"

	"github.com/RomaticDOG/GCR/FastGO/internal/model"
	"github.com/RomaticDOG/GCR/FastGO/internal/pkg/contextx"
	"github.com/RomaticDOG/GCR/FastGO/internal/pkg/conversion"
	"github.com/RomaticDOG/GCR/FastGO/internal/store"
	apiV1 "github.com/RomaticDOG/GCR/FastGO/pkg/api/v1"
	"github.com/jinzhu/copier"
	"github.com/onexstack/onexstack/pkg/store/where"
)

// PostBiz 博文业务层接口
type PostBiz interface {
	Create(ctx context.Context, req *apiV1.CreatePostReq) (*apiV1.CreatePostResp, error)
	Update(ctx context.Context, req *apiV1.UpdatePostReq) (*apiV1.UpdatePostResp, error)
	Delete(ctx context.Context, req *apiV1.DeletePostReq) (*apiV1.DeletePostResp, error)
	Get(ctx context.Context, req *apiV1.GetPostReq) (*apiV1.GetPostResp, error)
	List(ctx context.Context, req *apiV1.ListPostReq) (*apiV1.ListPostResp, error)

	PostExpansion
}

// PostExpansion 博文业务层扩展接口
type PostExpansion interface {
}

// postBiz 博文业务层接口具体实现
type postBiz struct {
	store store.IStore
}

// New 返回一个业务层接口具体实现
func New(store store.IStore) PostBiz {
	return &postBiz{store: store}
}

// Create 创建博文请求
func (pb *postBiz) Create(ctx context.Context, req *apiV1.CreatePostReq) (*apiV1.CreatePostResp, error) {
	var post model.Post
	_ = copier.Copy(&post, req)
	if err := pb.store.Post().Create(ctx, &post); err != nil {
		return nil, err
	}
	return &apiV1.CreatePostResp{PostID: post.PostID}, nil
}

// Update 更新博文请求
func (pb *postBiz) Update(ctx context.Context, req *apiV1.UpdatePostReq) (*apiV1.UpdatePostResp, error) {
	whr := where.F("userID", contextx.UserID(ctx), "postID", req.PostID)
	post, err := pb.store.Post().Get(ctx, whr)
	if err != nil {
		return nil, err
	}
	if req.Content != nil {
		post.Content = *req.Content
	}
	if req.Title != nil {
		post.Title = *req.Title
	}
	if err = pb.store.Post().Update(ctx, post); err != nil {
		return nil, err
	}
	return &apiV1.UpdatePostResp{}, nil
}

// Delete 删除博文请求
func (pb *postBiz) Delete(ctx context.Context, req *apiV1.DeletePostReq) (*apiV1.DeletePostResp, error) {
	whr := where.F("userID", contextx.UserID(ctx), "postID", req.PostIDs)
	if err := pb.store.Post().Delete(ctx, whr); err != nil {
		return nil, err
	}
	return &apiV1.DeletePostResp{}, nil
}

// Get 获取博文请求
func (pb *postBiz) Get(ctx context.Context, req *apiV1.GetPostReq) (*apiV1.GetPostResp, error) {
	whr := where.F("userID", contextx.UserID(ctx), "postID", req.PostID)
	post, err := pb.store.Post().Get(ctx, whr)
	if err != nil {
		return nil, err
	}
	return &apiV1.GetPostResp{Post: conversion.PostModelToPostV1(post)}, nil
}

// List 获取博文列表
func (pb *postBiz) List(ctx context.Context, req *apiV1.ListPostReq) (*apiV1.ListPostResp, error) {
	whr := where.F("userID", contextx.UserID(ctx)).P(int(req.Offset), int(req.Limit))
	if req.Title != nil {
		whr = whr.Q("title LIKE ?", fmt.Sprintf("%%%s%%", *req.Title))
	}
	cnt, posts, err := pb.store.Post().List(ctx, whr)
	if err != nil {
		return nil, err
	}
	ret := make([]*apiV1.Post, 0, len(posts))
	for _, post := range posts {
		ret = append(ret, conversion.PostModelToPostV1(post))
	}
	return &apiV1.ListPostResp{TotalCount: cnt, Posts: ret}, nil
}
