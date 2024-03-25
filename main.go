package main

import (
	"api-login/controllers"
	"api-login/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// Get spesifik data
func GetUsers(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, id)
	}
	var getUser *models.User
	for _, u := range models.Users {
		if u.ID == id {
			getUser = u
			break
		}
	}
	if getUser == nil {
		return c.JSON(http.StatusNotFound, id)
	}
	return c.JSON(http.StatusOK, getUser)
}

// Semua data array pakai slice 0
func GetAllUsers(c echo.Context) error {
	allUsers := make([]*models.User, 0, len(models.Users))
	for _, u := range models.Users {
		allUsers = append(allUsers, u)
	}
	return c.JSON(http.StatusOK, allUsers)
}
func GetUsersWithPagination(c echo.Context) error {
	// Mengambil parameter pagination dari query string
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	// Menghitung indeks awal dan akhir pengguna yang harus ditampilkan berdasarkan halaman dan batasan
	startIndex := (page - 1) * limit
	endIndex := startIndex + limit

	// Memastikan indeks akhir tidak melebihi jumlah total pengguna
	if endIndex > len(models.Users) {
		endIndex = len(models.Users)
	}
	// Mendapatkan pengguna yang sesuai dengan halaman dan batasan
	paginatedUsers := make([]*models.User, 0)
	for _, u := range models.Users {
		paginatedUsers = append(paginatedUsers, u)
	}
	return c.JSON(http.StatusOK, paginatedUsers[startIndex:endIndex])
}

func main() {
	e := echo.New()
	for i := 0; i < 40; i++ {
		newUser := &models.User{
			ID:   models.Seq,
			Name: "User" + strconv.Itoa(models.Seq),
		}
		models.Users[models.Seq] = newUser
		models.Seq++
	}

	e.POST("/users", controllers.CreateUser)
	e.GET("/users/:id", GetUsers)
	e.GET("/users", GetAllUsers)
	e.GET("/users/pagination", GetUsersWithPagination)
	e.PUT("/users", controllers.UpdateUser)
	e.DELETE("/users", controllers.DeleteUser)

	e.Logger.Fatal(e.Start(":9000"))

}
