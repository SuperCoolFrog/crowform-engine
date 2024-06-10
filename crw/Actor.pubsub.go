package crw

func (actor *Actor) Subscribe(key string, listener func()) {
	sub, exist := actor.subscribers[key]

	if !exist {
		sub = make([]func(), 0)
	}

	actor.subscribers[key] = append(sub, listener)
}

func (actor *Actor) Publish(key string) {
	listeners, exist := actor.subscribers[key]

	if !exist {
		return
	}

	for i := range listeners {
		l := listeners[i]
		l()
	}
}
