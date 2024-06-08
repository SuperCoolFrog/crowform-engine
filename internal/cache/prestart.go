package cache

var preloadQueue []func() = make([]func(), 0)

func QueueForPreload(qItem func()) {
	preloadQueue = append(preloadQueue, qItem)
}

func RunPreload() {
	for i := range preloadQueue {
		preloadQueue[i]()
	}

	RestPreload()
}

func RestPreload() {
	preloadQueue = nil
}
