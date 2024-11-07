package crw

import "crowform/internal/tools"

func (game *Game) subscribe(key GameStateEventKey, listener func()) (listenerId int) {
	sub, exist := game.subscribers[key]

	if !exist {
		sub = make([]gameSubListener, 0)
	}

	id := game.lastSubId + 1
	handler := gameSubListener{
		id:      id,
		handler: listener,
	}
	game.lastSubId = id

	game.subscribers[key] = append(sub, handler)

	return id
}
func (game *Game) unsubscribe(key GameStateEventKey, listenerId int) {
	sub, exist := game.subscribers[key]

	if !exist {
		return
	}

	game.subscribers[key] = tools.FilterSlice(sub, func(listener gameSubListener) bool {
		return listener.id != listenerId
	})
}

func (game *Game) publish(key GameStateEventKey) {
	listeners, exist := game.subscribers[key]

	if !exist {
		return
	}

	for i := range listeners {
		l := listeners[i]
		l.handler()
	}
}
