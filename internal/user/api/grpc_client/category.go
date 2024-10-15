package grpcclient

import (
	"context"
	"log"

	pb "github.com/achmad-dev/internal/pkg/common/genproto"
	"github.com/achmad-dev/internal/user/internal/middleware"
	"github.com/achmad-dev/internal/user/internal/service"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

type CategoryClient struct {
	client pb.CategoryServiceClient
}

func NewCategoryClient(conn *grpc.ClientConn) *CategoryClient {
	return &CategoryClient{
		client: pb.NewCategoryServiceClient(conn),
	}
}

func (c *CategoryClient) GetCategory(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	req := &pb.GetCategoryRequest{Id: id}
	resp, err := c.client.GetCategory(context.Background(), req)
	if err != nil {
		log.Printf("Error calling GetCategory: %v", err)
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Category not found"})
	}
	return ctx.Status(fiber.StatusOK).JSON(resp.Category)
}

func (c *CategoryClient) CreateCategory(ctx *fiber.Ctx) error {
	var req pb.CreateCategoryRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	_, err := c.client.CreateCategory(context.Background(), &req)
	if err != nil {
		log.Printf("Error calling CreateCategory: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Category created successfully"})
}

func (c *CategoryClient) GetCategoryByName(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	req := &pb.GetCategoryByNameRequest{Name: name}
	resp, err := c.client.GetCategoryByName(context.Background(), req)
	if err != nil {
		log.Printf("Error calling GetCategoryByName: %v", err)
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Category not found"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"category": resp.Category})
}

func (c *CategoryClient) UpdateCategory(ctx *fiber.Ctx) error {
	var req pb.UpdateCategoryRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	_, err := c.client.UpdateCategory(context.Background(), &req)
	if err != nil {
		log.Printf("Error calling UpdateCategory: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Category updated successfully"})
}

func (c *CategoryClient) DeleteCategory(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	req := &pb.DeleteCategoryRequest{Id: id}
	_, err := c.client.DeleteCategory(context.Background(), req)
	if err != nil {
		log.Printf("Error calling DeleteCategory: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{"message": "Category deleted successfully"})
}

func (c *CategoryClient) ListCategories(ctx *fiber.Ctx) error {
	req := &pb.ListCategoriesRequest{}
	resp, err := c.client.ListCategories(context.Background(), req)
	if err != nil {
		log.Printf("Error calling ListCategories: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"categories": resp.Categories})
}

func RegisterCategoryRoutes(api fiber.Router, conn *grpc.ClientConn, secret string, userService service.UserService) {
	client := NewCategoryClient(conn)
	api.Get("/category/:id", middleware.AuthAdminMiddleware(secret, userService), client.GetCategory)
	api.Get("/category/name/:name", middleware.AuthAdminMiddleware(secret, userService), client.GetCategoryByName)
	api.Post("/category", middleware.AuthAdminMiddleware(secret, userService), client.CreateCategory)
	api.Put("/category", middleware.AuthAdminMiddleware(secret, userService), client.UpdateCategory)
	api.Delete("/category/:id", middleware.AuthAdminMiddleware(secret, userService), client.DeleteCategory)
	api.Get("/categories", middleware.AuthAdminMiddleware(secret, userService), client.ListCategories)
}
