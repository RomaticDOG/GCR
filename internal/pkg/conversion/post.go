package conversion

import (
	"github.com/RomaticDOG/GCR/FastGO/internal/model"
	apiV1 "github.com/RomaticDOG/GCR/FastGO/pkg/api/v1"
	"github.com/onexstack/onexstack/pkg/core"
)

// PostModelToPostV1 将模型层的 Post（博客模型对象）转换为 Protobuf 层的 Post（v1 博客对象）.
func PostModelToPostV1(postModel *model.Post) *apiV1.Post {
	var protoPost apiV1.Post
	_ = core.CopyWithConverters(&protoPost, postModel)
	return &protoPost
}

// PostV1ToPostModel 将 Protobuf 层的 Post（v1 博客对象）转换为模型层的 Post（博客模型对象）.
func PostV1ToPostModel(protoPost *apiV1.Post) *model.Post {
	var postModel model.Post
	_ = core.CopyWithConverters(&postModel, protoPost)
	return &postModel
}
