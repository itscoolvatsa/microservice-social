package natshandler

import (
	"context"
	"fmt"
	"microservice/graph/internal/controller"
	"microservice/graph/internal/repository/natsmsg"
	"microservice/graph/pkg/model"
	"time"
)

type natsMsg struct {
	ctrl *controller.Controller
}

func New(ctrl *controller.Controller) (*natsMsg, error) {
	return &natsMsg{
		ctrl: ctrl,
	}, nil
}

func (h *natsMsg) ReceivesMessage(data string) {
	h.ctrl.ReceivesMessage(data, func(res *natsmsg.NatsUser) {
		str := fmt.Sprint(res)
		fmt.Println(str + "handler function wala")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
		defer cancel()
		var user model.User

		user.UserId = res.UserId.Hex()
		user.Name = res.Name

		_, err := h.ctrl.AddUser(ctx, &user)

		if err != nil {
			panic(err)
		}
	})
}
