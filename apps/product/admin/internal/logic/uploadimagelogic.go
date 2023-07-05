package logic

import (
	"context"
	"net/http"

	"github.com/liujingkaiai/go-zero-mall/apps/product/admin/internal/svc"
	"github.com/liujingkaiai/go-zero-mall/apps/product/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	imageFileName = "image"
	bucketName    = "lebron-mall"
)

type UploadImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewUploadImageLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *UploadImageLogic {
	return &UploadImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *UploadImageLogic) UploadImage() (resp *types.UploadImageResponse, err error) {
	// todo: add your logic here and delete this line
	file, header, err := l.r.FormFile(imageFileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	bucket, err := l.svcCtx.OssClient.Bucket(bucketName)
	if err != nil {
		return nil, err
	}

	if err = bucket.PutObject(header.Filename, file); err != nil {
		return nil, err
	}
	return &types.UploadImageResponse{
		Success: true,
	}, nil
}
