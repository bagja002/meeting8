package models


type Admin struct{
	Id_admin uint   `gorm:"primaryKey:autoIncrement"`
	Nama string
	Email string
	Password string
	No_telpon int
	 //Deklarasi one to many 
}

type Users struct{
	Id_users uint `gorm:"primaryKey:autoIncrement"`
	Nama string
	Email string
	Password string
	No_telpon string
	Alamat string
}