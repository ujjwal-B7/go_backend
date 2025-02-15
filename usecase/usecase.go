package usecase

import "backend/model"

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
