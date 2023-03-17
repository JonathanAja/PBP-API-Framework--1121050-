package controllers

import (
	"LatihanRevel/app/models"

	"github.com/revel/revel"
)

type AppRev struct {
	*revel.Controller
}

func (c AppRev) Index() revel.Result {
	return c.Render()
}

func (c AppRev) GetUsers() revel.Result {
	db := connect()
	defer db.Close()

	query := "select * from users"
	rows, err := db.Query((query))
	if err != nil {
		revel.AppLog.Fatal("%s", err)
	}
	var user models.User
	var users []models.User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Password, &user.Email, &user.UserType); err != nil {
			revel.AppLog.Info(err.Error())
		} else {
			users = append(users, user)
		}
	}

	return c.RenderJSON(users)
}

func (c AppRev) InsertUsers() revel.Result {
	db := connect()
	defer db.Close()

	var user models.User
	err := c.Params.BindJSON(&user)
	if err != nil {
		return c.RenderError(err)
	}

	query := "INSERT INTO users (Name, Age, Address, email,password) value (?, ?, ?, ?, ?)"
	_, errQuery := db.Exec(query, user.Name, user.Age, user.Addres, user.Email, user.Password)

	var response models.UserResponse
	if errQuery == nil {

		response.Message = "succes"
		response.Status = 200
	} else {
		revel.AppLog.Info(errQuery.Error())
		response.Message = "Failed"
		response.Status = 400
	}

	return c.RenderJSON(response)
}

func (c AppRev) DeleteUser() revel.Result {
	db := connect()
	defer db.Close()

	id := c.Params.Get("id")
	revel.AppLog.Info(id)
	query := "delete from users where id = ?"
	_, errQuery := db.Exec(query, id)

	var response models.UserResponse
	if errQuery == nil {

		response.Message = "succes"
		response.Status = 200
	} else {
		revel.AppLog.Info(errQuery.Error())
		response.Message = "Failed"
		response.Status = 400
	}
	return c.RenderJSON(response)
}

func (c AppRev) UpdateUsers() revel.Result {
	db := connect()
	defer db.Close()
	var user models.User
	id := c.Params.Get("id")
	revel.AppLog.Info(id)
	query := "UPDATE users SET name = ?, age = ?, address = ?,email =?,password =? WHERE id = ?"
	_, errQuery := db.Exec(query, user.Name, user.Age, user.Addres, user.Email, user.Password)

	var response models.UserResponse
	if errQuery == nil {

		response.Message = "success"
		response.Status = 200
	} else {
		revel.AppLog.Info(errQuery.Error())
		response.Message = "Failed"
		response.Status = 400
	}
	return c.RenderJSON(response)
}
