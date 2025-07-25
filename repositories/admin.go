package repositories

import (
	"MentalHealthCare/database"
	"MentalHealthCare/models"
)


type AdminRepositoryImpl struct{}

// func InitAdminRepository() AdminRepository {
// 	return &AdminRepositoryImpl{}
// }

func (ar *AdminRepositoryImpl) GetAll() ([]models.Admin, error) {
	var admins []models.Admin

	if err := database.DB.Preload("User").Find(&admins).Error; err != nil {
		return []models.Admin{}, err
	}

	return admins, nil
}

func (ar *AdminRepositoryImpl) GetByID(id string) (models.Admin, error) {
	var admin models.Admin

	if err := database.DB.Preload("User").First(&admin, "id = ?", id).Error; err != nil {
		return models.Admin{}, err
	}

	return admin, nil
}

func (ar *AdminRepositoryImpl) GetByEmail(loginReq models.LoginRequest) (models.Admin, error) {
	var admin models.Admin

	if err := database.DB.First(&admin, "email = ?", loginReq.Email).Error; err != nil {
		return models.Admin{}, err
	}

	return admin, nil
}

func (ar *AdminRepositoryImpl) Create(adminReq models.AdminRequest) (models.Admin, error) {
	var admin models.Admin = models.Admin{
		UserID: adminReq.UserID,
		ManageUsers: adminReq.ManageUsers,
	}

	result := database.DB.Create(&admin)

	if err := result.Error; err != nil {
		return models.Admin{}, err
	}

	if err := result.Last(&admin).Error; err != nil {
		return models.Admin{}, err
	}

	return admin, nil
}

func (ar *AdminRepositoryImpl) Update(adminReq models.AdminRequest, id string) (models.Admin, error) {
	admin, err := ar.GetByID(id)

	if err != nil {
		return models.Admin{}, err
	}

	admin.UserID = adminReq.UserID
	admin.ManageUsers = adminReq.ManageUsers

	if err := database.DB.Save(&admin).Error; err != nil {
		return models.Admin{}, err
	}

	return admin, nil
}

func (ar *AdminRepositoryImpl) Delete(id string) error {
	admin, err := ar.GetByID(id)

	if err != nil {
		return err
	}

	if err := database.DB.Delete(&admin).Error; err != nil {
		return err
	}

	return nil
}
