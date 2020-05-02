package usecase

import "task-api/src/entity/model"

type MiddlewareInteractor interface {
	CanAccessProject(*CanAccessProjectInputDS) (bool, error)
	CanWriteProject(*CanWriteProjectInputDS) (bool, error)
}

type middlewareInteractor struct {
	userRepository UserRepository
}

func NewMiddlewareInteractor(uRepo UserRepository) MiddlewareInteractor {
	return &middlewareInteractor{uRepo}
}

type CanAccessProjectInputDS struct {
	UID int64
	PID int64
}

func (mi *middlewareInteractor) CanAccessProject(in *CanAccessProjectInputDS) (bool, error) {
	// プロジェクト参加ユーザー一覧を取得
	// TODO: リポジトリの取得メソッドを変更する
	projectUsers, err := mi.userRepository.FindByProjectID(in.PID)
	if err != nil {
		return false, err
	}

	// プロジェクト参加ユーザー一覧にinputのuserIDが存在するか確認
	for _, u := range *projectUsers {
		if u.ID == in.UID {
			return true, nil
		}
	}

	return false, nil
}

type CanWriteProjectInputDS struct {
	UID int64
	PID int64
}

func (mi *middlewareInteractor) CanWriteProject(in *CanWriteProjectInputDS) (bool, error) {
	// プロジェクト参加ユーザー一覧を取得
	// TODO: リポジトリの取得メソッドを変更する
	projectUsers, err := mi.userRepository.FindByProjectID(in.PID)
	if err != nil {
		return false, err
	}

	// プロジェクトに参加してるかつ、Write以上の権限があるか確認
	for _, u := range *projectUsers {
		if u.ID == in.UID && (u.Role == model.Admin || u.Role == model.Write) {
			return true, nil
		}
	}

	return false, nil
}
