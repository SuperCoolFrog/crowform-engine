package crw

import (
	"crowform/internal/mog"
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

	mog.Debug("%v -- runActions::actions length = %d", actor.Tags, len(actions))
	mog.Debug("%v -- runActions::actions slice = %v", actor.Tags, actions)

	sort.Slice(actions, func(i, j int) bool {
		return actions[i].index < actions[j].index
	})
	mog.Debug("%v -- runActions::actions sorted = %v", actor.Tags, actions)

	actor.doActions(deltaTime, actions, 0, func() {
		mog.Debug("%v -- runActions::SetAsReady", actor.Tags)
		actor.ActionsSetAsReady()
	})
}

func (actor *Actor) doActions(deltaTime time.Duration, allActions []ActorAction, idx int, onComplete func()) {
	if idx > len(allActions)-1 {
		mog.Debug("%v -- doActions:: onComplete called idx = %d, len = %d", actor.Tags, idx, len(allActions))
		onComplete()
		return
	}

	// Check if still valid when - can change based on other actions run
	if allActions[idx].when(actor) {
		mog.Debug("%v -- doActions:: when TRUE idx = %d", actor.Tags, idx)
		allActions[idx].do(deltaTime, actor, func() {
			mog.Debug("%v -- doActions:: when TRUE done called for idx = %d", actor.Tags, idx)
			actor.doActions(deltaTime, allActions, idx+1, onComplete)
		})
	} else {
		mog.Debug("%v -- doActions:: when FALSE idx = %d", actor.Tags, idx)
		actor.doActions(deltaTime, allActions, idx+1, onComplete)
	}
}
