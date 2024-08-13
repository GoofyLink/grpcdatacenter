package handler

import (
	"ucenter-api/internal/svc"
)

func RegisterHandlers(r *Routes, serverCtx *svc.ServiceContext) {
	register := NewRegisterHandler(serverCtx)
	registerRouter := r.Group()
	registerRouter.Get("/uc/register/phone", register.Register)
}
