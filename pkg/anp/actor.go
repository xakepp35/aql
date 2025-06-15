package anp

import (
	"context"

	"github.com/xakepp35/aql/pkg/aql"
)

type Actor struct {
	ID     string   // Имя или UUID
	Owner  *Client  // Кто запустил (может быть nil)
	VM     *aql.VM  // Собственная виртуальная машина
	Inbox  chan any // Приёмник сообщений
	Cancel context.CancelFunc
}

func (a *Actor) Run(ctx context.Context) {
	defer a.Cancel()
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-a.Inbox:
			a.VM.Push(msg)
			a.VM.Run()
			if a.VM.Err != nil {
				fmt.Println("actor:", a.ID, "err:", a.VM.Err)
			}
		}
	}
}
