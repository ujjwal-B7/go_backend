package usecase

import (
	"backend/model"
)

type UseCase struct {
	Repo model.FutsalRepositoryInterface
}

func NewUsecase(repo model.FutsalUsecaseInterface) model.FutsalUsecaseInterface {
	return &UseCase{Repo: repo}
}

func (u *UseCase) SaveFutsal(futsal model.Futsal) error {
	err := u.Repo.SaveFutsal(futsal)

	if err != nil {
		return err
	}
	return nil
}

func (u *UseCase) GetAllFutsals() (*[]model.Futsal, error) {
	data, err := u.Repo.GetAllFutsals()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *UseCase) UpdateFutsal(id uint64, futsal model.Futsal) (*model.Futsal, error) {

	data, err := u.Repo.UpdateFutsal(id, futsal)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *UseCase) UpdateFutsalFields(id uint64, futsal model.Futsal) (*model.Futsal, error) {

	data, err := u.Repo.UpdateFutsalFields(id, futsal)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *UseCase) DeleteFutsal(id uint64) error {
	err := u.Repo.DeleteFutsal(id)

	if err != nil {
		return err
	}

	return nil
}
