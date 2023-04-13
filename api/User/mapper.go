package User

import (
	serviceModel "api/model"
	"api/user/model"
)

func MapClientToServiceUserModel(model model.User) serviceModel.User {
	return serviceModel.User{
		Gender:         model.Gender,
		Title:          model.Name.Title,
		First:          model.Name.First,
		Last:           model.Name.Last,
		Email:          model.Email,
		Date:           model.DOB.Date,
		Age:            model.DOB.Age,
		RegisteredDate: model.Registered.Date,
		RegisteredAge:  model.Registered.Age,
		Phone:          model.Phone,
		Cell:           model.Cell,
		IdName:         model.Id.Name,
		IdValue:        model.Id.Value,
		Nat:            model.Nat,
	}
}

func MapClientToServiceLocationModel(model model.User) serviceModel.Location {
	return serviceModel.Location{
		Number:      model.Location.Street.Number,
		Name:        model.Location.Street.Name,
		City:        model.Location.City,
		State:       model.Location.State,
		Country:     model.Location.Country,
		Postcode:    model.Location.Postcode,
		Latitude:    model.Location.Coordinates.Latitude,
		Longitude:   model.Location.Coordinates.Longitude,
		Offset:      model.Location.Timezone.Offset,
		Description: model.Location.Timezone.Description,
	}
}

func MapClientToServiceLoginModel(model model.User) serviceModel.Login {
	return serviceModel.Login{
		Uuid:     model.Login.Uuid,
		Username: model.Login.Username,
		Password: model.Login.Password,
		Salt:     model.Login.Salt,
		MD5:      model.Login.MD5,
		SHA1:     model.Login.SHA1,
		SHA256:   model.Login.SHA256,
	}
}

func MapClientToServicePictureModel(model model.User) serviceModel.Picture {
	return serviceModel.Picture{
		Large:     model.Picture.Large,
		Medium:    model.Picture.Medium,
		Thumbnail: model.Picture.Thumbnail,
	}
}
