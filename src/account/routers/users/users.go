package users

import (
	"cos-backend-com/src/account/routers"
	"cos-backend-com/src/common/flake"
	"cos-backend-com/src/libs/apierror"
	"cos-backend-com/src/libs/models/users"
	"cos-backend-com/src/libs/sdk/account"

	"github.com/wujiu2020/strip/utils/apires"
)

type Users struct {
	routers.Base
	Uid flake.ID `inject:"uid"`
}

func (h *Users) Me() (res interface{}) {
	var input account.LoginInput
	if err := h.Params.BindJsonBody(&input); err != nil {
		h.Log.Warn(err)
		res = apierror.ErrBadRequest.WithData(err)
		return
	}

	var user account.UserResult
	if err := users.Users.Get(h.Ctx, h.Uid, &user); err != nil {
		h.Log.Warn(err)
		res = apierror.HandleError(err)
		return
	}

	res = apires.With(user)
	return
}
