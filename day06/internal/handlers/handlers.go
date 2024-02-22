package handlers

import (
	"day06/internal/service"
	"day06/models"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"time"
)

type Handler struct {
	Db     service.Store
	Logger *zap.Logger
}

//TODO
//service to handle delete that is not exists,

var limitPerPage = 3

func New(db service.Store, logger *zap.Logger) *Handler {
	return &Handler{
		Db:     db,
		Logger: logger,
	}
}

func (h *Handler) InitRoutes(app *fiber.App) {
	app.Post("/article", h.requireAuthorization, h.AddArticle)
	app.Get("/article/:id", h.GetArticle)
	app.Get("/articles", h.GetArticles)
	app.Delete("/article/:id", h.requireAuthorization, h.RemoveArticle)
}

func (h *Handler) AddArticle(c *fiber.Ctx) error {
	article := models.Articles{}

	err := c.BodyParser(&article)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Success": false, "Error": "invalid data input", "Message": err})
	}

	article.Date = time.Now()
	err = h.Db.AddArticle(article)
	if err != nil {
		h.Logger.Info(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Success": false, "Error": "article was not created", "Message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Success": true, "Id": article.Id})
}

func (h *Handler) GetArticle(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", -1)
	if err != nil || id < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Success": false, "Error": "invalid id input"})
	}

	article, err := h.Db.GetArticle(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"Success": false, "Article": nil, "Message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Success": true, "Article": article})

}

func (h *Handler) GetArticles(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	size := c.QueryInt("size", limitPerPage)

	searchParams := models.SearchParams{
		Offset: (page - 1) * limitPerPage,
		Limit:  size,
	}

	articles, err := h.Db.GetArticles(searchParams)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"Success": false, "Articles": nil, "Message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Success": true, "Articles": articles})
}

func (h *Handler) RemoveArticle(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", -1)
	if err != nil || id < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Success": false, "Error": "invalid id input"})
	}

	err = h.Db.RemoveArticle(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"Success": false, "Error": "article was not deleted", "Message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Success": true, "Id": id})
}

//func (h *Handler) JWTAuth(c *fiber.Ctx) error {
//
//}

func (h *Handler) requireAuthorization(c *fiber.Ctx) error {
	// Получаем токен из заголовка Authorization
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	//// Проверяем токен
	//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	//	// Убедитесь, что метод подписи соответствует вашему ожиданию
	//	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	//		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	//	}
	//	return jwtKey, nil
	//})
	//if err != nil {
	//	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	//		"error": "Unauthorized",
	//	})
	//}
	//if !token.Valid {
	//	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	//		"error": "Unauthorized",
	//	})
	//}
	//
	//// Продолжаем выполнение цепочки middleware
	return c.Next()
}
