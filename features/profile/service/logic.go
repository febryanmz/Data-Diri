package service

import (
	"github.com/febryanmz/Data-Diri/features/profile"
)

type profileService struct {
	profileRepository profile.RepositoryInterface
}

func New(repo profile.RepositoryInterface) profile.ServiceInterface {
	return &profileService{
		profileRepository: repo,
	}
}

func (service *profileService) Create(input profile.Core) (err error) {
	_, errCreate := service.profileRepository.Create(input)
	if errCreate != nil {
		return errors.New("failed to insert data, error logic")
	}
	return nil
}

//// Get by (ID)
//func (service *profileService) GetById(id int) (data profile.Core, err error) {
//	data, errGet := service.profileRepository.GetById(id)
//	if errGet != nil {
//		return data, errors.New("Get data by ID failed, error logic")
//	}
//	return data, nil
//}
//
//// Update
//func (service *profileService) Update(dataCore profile.Core, id int) (err error) {
//	errUpdate := service.profileRepository.Update(dataCore, id)
//	if errUpdate != nil {
//		return errors.New("update failed, error logic")
//	}
//	return nil
//
//}
//
//// Delete
//func (service *profileService) Delete(id int) (err error) {
//	_, errDel := service.profileRepository.Delete(id)
//	if errDel != nil {
//		return errors.New("delete failed, error logic")
//	}
//	return nil
//}
