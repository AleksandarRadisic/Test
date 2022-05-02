package app

import "github.com/AleksandarRadisic/Test/GoPaket/domain"

func Data() domain.User {
	user := domain.User{
		Username: "sad",
	}
	return user
}
