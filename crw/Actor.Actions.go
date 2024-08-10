package crw

import (
	"crowform/internal/tools"
	"sort"
	"time"
)

func (actor *Actor) ActionsSetAsReady() {
	actor.actorActionState = ActorActionState_READY
}

func (actor *Actor) ActionsSetAsProcessing() {
	actor.actorActionState = ActorActionState_PROCESSING
}

func (actor *Actor) ActionsSetAsStop() {
	actor.actorActionState = ActorActionState_STOP
}

func (actor *Actor) runActions(deltaTime time.Duration) {
	if actor.actorActionState != ActorActionState_READY {
		return
	}

	actor.ActionsSetAsProcessing()

	actions := tools.FilterSlice(actor.actions, func(a ActorAction) bool { return a.when(actor) })

	if len(actions) == 0 {
		actor.ActionsSetAsReady()
		return
	}

	sort.Slice(actions, func(i, j int) bool {
		return actions[i].index < actions[j].index
	})

	actor.doActions(deltaTime, actions, 0, func() {
		actor.ActionsSetAsReady()
	})
}

func (actor *Actor) doActions(deltaTime time.Duration, allActions []ActorAction, idx int, onComplete func()) {
	if idx > len(allActions)-1 {
		onComplete()
		return
	}

	// Check if still valid when - can change based on other actions run
	if allActions[idx].when(actor) {
		allActions[idx].do(deltaTime, actor, func() {
			actor.doActions(deltaTime, allActions, idx+1, onComplete)
		})
	} else {
		actor.doActions(deltaTime, allActions, idx+1, onComplete)
	}
}
