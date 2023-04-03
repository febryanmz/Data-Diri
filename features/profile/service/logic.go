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
