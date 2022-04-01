package routes

import (
	"math"
	"net/http"
	"os"
	"strconv"
	"thailephan/flashcard-echo-api/entities"
	"thailephan/flashcard-echo-api/pkg/utils"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/teris-io/shortid"
)

var users []*entities.User

func userRoutes(e *echo.Echo) {
	var ex string
	ex, err := os.Getwd();
	if err != nil {
		panic(err)
	}

	var ab = ex + "\\mocks\\users.json";
	if err := utils.ReadJson(ab, &users); err != nil {
		panic(err)
	}

	e.GET("/users", getAllUsers)		
	e.POST("/users", createUser)		
	e.GET("/users/:id", getUserById)		
	e.PUT("/users/:id", updateUser)		
	e.DELETE("/users/:id", deleteUser)
}

type Map map[string]interface{}

func getAllUsers(c echo.Context) error {
	var err error
	var limit, offset int
	var limitParam, offsetParam string

	limitParam = c.QueryParam("limit")
	offsetParam = c.QueryParam("offset")

	if len(limitParam) == 0 {
		limit = 10
	} else {
		limit, err = strconv.Atoi(limitParam)
		if (limit < 0) {
			return c.JSON(http.StatusBadRequest, Map{
				"error": Map{
					"message": "`offset` is less than 0: " + limitParam,
					"code": 0,
				},
			})
		}
		if err != nil {
			return c.JSON(http.StatusBadRequest, Map{
				"error": Map{
					"message": "Error to parse `limit`: " + limitParam,
					"code": 0,
				},
			})
		}
	}

	if len(offsetParam) == 0 {
		offset = 0
	} else {
		offset, err = strconv.Atoi(offsetParam)
		if (offset < 0) {
			return c.JSON(http.StatusBadRequest, Map{
				"error": Map{
					"message": "`offset` is less than 0: " + offsetParam,
					"code": 0,
				},
			})
		}
		if err != nil {
			return c.JSON(http.StatusBadRequest, Map{
				"error": Map{
					"message": "Error to parse `offset`: " + offsetParam,
					"code": 0,
				},
			})
		}
	}

	var len64 = float64(len(users))
	var lo = int(math.Max(0, math.Min(float64(offset), len64)))
	var hi = int(math.Max(0, math.Min(float64(offset+limit), len64)))
	var filteredUsers = users[lo:hi]

	return c.JSON(http.StatusOK, Map{
		"data": Map{
			"items": filteredUsers,
			"count": len(filteredUsers),
			"query": Map{
				"limit": limitParam,
				"offset": offsetParam,
			},
		},
		"statusCode": http.StatusOK,
	})	
}

func createUser(c echo.Context) error {
	var user = &entities.User{}

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, Map{
			"error": Map{
				"message": err.Error(),
				"code": 0,
			},
		})
	}

	user.ID, _ = shortid.Generate()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	users = append(users, user)

	return c.JSON(http.StatusOK, Map{
		"data": Map{
			"items": user,
			"count": 1,
		},
	})
}

func getUserById(c echo.Context) error {
	id := c.Param("id")
	var u *entities.User

	for _, user := range users {
		if user.ID == id {
			u = user
			break
		}
	}
	return c.JSON(http.StatusOK, Map{
		"data": Map{
			"items": u,
			"count": 1,
		},
	})
}

func updateUser(c echo.Context) error {
	id := c.Param("id")
	var u = &entities.User{}

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, Map{
			"error": Map{
				"message": err.Error(),
				"code": 0,
			},
		})
	}

	for i, user := range users {
		if user.ID == id {
			u.ID = id
			u.UpdatedAt = time.Now()

			users = append(users[:i], users[i+1:]...)
			users = append(users, u)
			break
		}
	}
	return c.JSON(http.StatusOK, Map{
		"data": Map{
			"items": u,
			"count": 1,
		},
	})
}

func deleteUser(c echo.Context) error {
	id := c.Param("id")
	var u = &entities.User{}

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, Map{
			"error": Map{
				"message": err.Error(),
				"code": 0,
			},
		})
	}

	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			break
		}
	}
	return c.JSON(http.StatusOK, Map{
		"data": Map{
			"items": u,
			"count": 1,
		},
	})
}