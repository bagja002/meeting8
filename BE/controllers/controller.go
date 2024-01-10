package controllers

import (
	"belajar-database/database"
	"belajar-database/models"

	"github.com/gofiber/fiber/v2"
)

func GetDataAll(c *fiber.Ctx)error{


	var user  []models.Users

	//select * from users where 1
	database.DB.Find(&user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Pesan":"Berhasil Get data",
		"data":user,
	})
}

func GetOneData(c *fiber.Ctx)error{

	var user models.Users

	id_user := c.Query("id_user")

	// select * from users where id_user = id_uses
	database.DB.Where("id_users = ?", id_user).Find(&user)


	return c.JSON(fiber.Map{
		"Pesan":"Berasil get data",
		"data": user ,
	})
}

func UpdateData(c *fiber.Ctx)error{

	var user models.Users

	var data map[string]string

	err:= c.BodyParser(&data)
	if err!= nil {
		return c.JSON(fiber.Map{
			"pesan":"Isi data yang benar",
		})
	}
	id_user := c.Query("id_user")

	users:= models.Users{
		Nama: data["nama"],
		No_telpon: data["no_telpon"],
		Email: data["email"],
		Password: data["password"],
	}

	//perubahan data 

	database.DB.Model(&user).Where("id_users = ?", id_user).Updates(&users)


	return c.JSON(fiber.Map{
		"Pesan":"Data Telah di update",
		"data yang di Update yaitu":user,
	})
}


func CreateData(c *fiber.Ctx)error{

	var data map[string]string

	err:= c.BodyParser(&data)
	if err!= nil {
		return c.JSON(fiber.Map{
			"pesan":"Isi data yang benar",
		})
	}

	user:= models.Users{
		Nama: data["nama"],
		Email: data["email"],
		Password: data["password"],
		No_telpon: data["no_telpon"],
		Alamat: data["alamat"],
	}

	database.DB.Create(&user)

	return c.JSON(fiber.Map{
		"Pesan":"Data Telah di tambah",
		"data yang di tambahkan yatu":user,
	})
}