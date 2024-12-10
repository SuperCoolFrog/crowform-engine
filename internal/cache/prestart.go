package cache

var preloadQueue []func() = make([]func(), 0)
var hasRunPreload = false

// Run before game starts; if game started runs immediate
func QueueForPreload(qItem func()) {
	if hasRunPreload {
		qItem()
		return
	}
	preloadQueue = append(preloadQueue, qItem)
}

func RunPreload() {
	for i := range preloadQueue {
		preloadQueue[i]()
	}

	RestPreload()
	hasRunPreload = true
}

func RestPreload() {
	preloadQueue = nil
}
