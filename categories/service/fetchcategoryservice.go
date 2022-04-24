package service

import (
	"context"

	"qwik.in/categories/log"
	"qwik.in/categories/proto"
	"qwik.in/categories/repository"
)

type CategoryFetchService struct {
	proto.UnimplementedFetchcategoryServer
}

var (
	Repository repository.CategoryRepository
	Service    CategoryService
)

func NewCategoryFetchService() *CategoryFetchService {
	Repository = repository.NewDynamoRepository()
	err := Repository.Connect()
	if err != nil {
		log.Error("Error while Connecting to DB: ", err)
		return nil
	}
	Service = NewCategoryService(Repository)
	return &CategoryFetchService{}
}

func (s *CategoryFetchService) GetCategoryById(ctx context.Context, in *proto.Request) (*proto.Response, error) {
	log.Info("gRPC received message: ", in.GetId())
	category, err1 := Service.SearchCategory(in.GetId())
	if err1 != nil {
		return nil, err1
	}
	return &proto.Response{
		Id:              in.GetId(),
		Name:            string(category.Category_description[0].Name),
		Description:     string(category.Category_description[0].Description),
		MetaDescription: string(category.Category_description[0].Meta_description),
		MetaKeyword:     string(category.Category_description[0].Meta_keyword),
		MetaTitle:       string(category.Category_description[0].Meta_title),
	}, nil
}
