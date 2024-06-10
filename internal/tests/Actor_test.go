package tests

import (
	"crowform/crw"
	"log"
	"testing"
	"time"
)

func TestActorQueryIncludesSelf(t *testing.T) {
	a1 := crw.BuildActor().Build()
	a2 := crw.BuildActor().Build()
	a3 := crw.BuildActor().Build()

	var Solid crw.QueryAttribute = "TEST_SOLID"
	var Attacks crw.QueryAttribute = "TEST_ATTACKS"

	a1.AddQueryAttr(Solid)
	a2.AddQueryAttr(Solid)
	a3.AddQueryAttr(Attacks)

	a1.AddChild(a2)
	a1.AddChild(a3)

	q := []crw.QueryAttribute{Solid}
	res := a1.QueryAny(q)

	if len(res) != 2 {
		t.Fatalf("Test query returns incorrect results actual %d, expected %d", len(res), 2)
	}

	log.Output(1, "[PASS]: TestActorQueryIncludesSelf")
}

func TestActorQueryChildrenWithoutSelf(t *testing.T) {
	a1 := crw.BuildActor().Build()
	a2 := crw.BuildActor().Build()
	a3 := crw.BuildActor().Build()

	var Solid crw.QueryAttribute = "TEST_SOLID"
	var Attacks crw.QueryAttribute = "TEST_ATTACKS"

	a2.AddQueryAttr(Solid)
	a3.AddQueryAttr(Attacks)

	a1.AddChild(a2)
	a1.AddChild(a3)

	q := []crw.QueryAttribute{Solid}
	res := a1.QueryAny(q)

	if len(res) != 1 {
		t.Fatalf("Test query returns incorrect results actual %d, expected %d", len(res), 1)
	}

	log.Output(1, "[PASS]: TestActorQueryChildrenWithoutSelf")
}

func TestActorAddChildZOrderAsc(t *testing.T) {
	a1 := crw.BuildActor().Build()
	a2 := crw.BuildActor().WithPosition(0, 0, 1).Build()
	a3 := crw.BuildActor().WithPosition(0, 0, 2).Build()

	a1.AddChild(a2)
	a1.AddChild(a3)

	if a1.Children[0] != a2 {
		t.Fatalf("Test ChildZ returns incorrect results: expected [a2] at index [0]")
	}
	if a1.Children[1] != a3 {
		t.Fatalf("Test ChildZ returns incorrect results: expected [a3] at index [1]")
	}

	log.Output(1, "[PASS]: TestActorAddChildZOrderAsc")
}

func TestActorAddChildZOrderDesc(t *testing.T) {
	a1 := crw.BuildActor().Build()
	a2 := crw.BuildActor().WithPosition(0, 0, 2).Build()
	a3 := crw.BuildActor().WithPosition(0, 0, 1).Build()

	a1.AddChild(a2)
	a1.AddChild(a3)

	if a1.Children[0] != a3 {
		t.Fatalf("Test ChildZ returns incorrect results: expected [a3] at index [0]")
	}

	if a1.Children[1] != a2 {
		t.Fatalf("Test ChildZ returns incorrect results: expected [a2] at index [1]")
	}

	log.Output(1, "[PASS]: TestActorAddChildZOrderDesc")
}

func TestActorAddChildZDefault(t *testing.T) {
	a1 := crw.BuildActor().Build()
	a2 := crw.BuildActor().WithPosition(0, 0, 0).Build()
	a3 := crw.BuildActor().WithPosition(0, 0, 0).Build()

	a1.AddChild(a2)
	a1.AddChild(a3)

	if a1.Children[0] != a2 {
		t.Fatalf("Test ChildZ returns incorrect results: expected [a2] at index [0]")
	}

	if a1.Children[1] != a3 {
		t.Fatalf("Test ChildZ returns incorrect results: expected [a3] at index [1]")
	}

	log.Output(1, "[PASS]: TestActorAddChildZDefault")
}

func TestActorActionsRun(t *testing.T) {
	expected := "ACTION RAN"
	actual := "ACTION DID NOT RUN"

	a1 := crw.BuildActor().
		WithAction(
			func(actor *crw.Actor) bool {
				return true
			},
			func(deltaTime time.Duration, actor *crw.Actor, done crw.ActorActionDone) {
				actual = expected
				done()
			}).
		Build()

	a1.ActionsSetAsReady()
	a1.Update(time.Second)

	if actual != expected {
		t.Fatalf("Actions were not run actual %s, expected %s", actual, expected)
	}

	log.Output(1, "[PASS]: TestActorActionsRun")
}

func TestActorActionsRunInSequence(t *testing.T) {
	expected := []int{1, 2, 3}
	actual := make([]int, 0)

	a1 := crw.BuildActor().
		WithAction(
			func(actor *crw.Actor) bool {
				return true
			},
			func(deltaTime time.Duration, actor *crw.Actor, done crw.ActorActionDone) {
				actual = append(actual, 1)
				done()
			}).
		WithAction(
			func(actor *crw.Actor) bool {
				return true
			},
			func(deltaTime time.Duration, actor *crw.Actor, done crw.ActorActionDone) {
				actual = append(actual, 2)
				done()
			}).
		WithAction(
			func(actor *crw.Actor) bool {
				return true
			},
			func(deltaTime time.Duration, actor *crw.Actor, done crw.ActorActionDone) {
				actual = append(actual, 3)
				done()
			}).
		Build()

	a1.ActionsSetAsReady()
	a1.Update(time.Second)

	if actual[0] != 1 || actual[1] != 2 || actual[2] != 3 {
		t.Fatalf("Actions were not run in sequence actual %v, expected %v", actual, expected)
	}

	log.Output(1, "[PASS]: TestActorActionsRunInSequence")
}

func TestPubSub(t *testing.T) {
	expected := "SUB CALLED"
	actual := "SUB NOT CALLED"

	a1 := crw.BuildActor().Build()
	a1.Subscribe("1", func() {
		actual = expected
	})

	a1.Publish("1")

	if actual != expected {
		t.Fatalf("Failed PubSub actual '%s', expected '%s'", actual, expected)
	}

	log.Output(1, "[PASS]: TestPubSub")
}

func TestPubAllSubs(t *testing.T) {
	expected := 3
	actual := 0

	a1 := crw.BuildActor().Build()
	a1.Subscribe("1", func() {
		actual++
	})
	a1.Subscribe("1", func() {
		actual++
	})
	a1.Subscribe("1", func() {
		actual++
	})

	a1.Publish("1")

	if actual != expected {
		t.Fatalf("Failed Pub AllSubs actual times %d, expected times %d", actual, expected)
	}

	log.Output(1, "[PASS]: TestPubAllSubs")
}

func TestRemoveChild(t *testing.T) {
	expected := 0

	a1 := crw.BuildActor().Build()
	c1 := crw.BuildActor().Build()

	a1.AddChild(c1)
	a1.RemoveChild(c1)

	actual := len(a1.Children)

	if actual != expected {
		t.Fatalf("Failed Remove children actual children count %d, expected %d", actual, expected)
	}

	log.Output(1, "[PASS]: TestRemoveChild")
}

func TestRemoveSelf(t *testing.T) {
	expected := 0

	a1 := crw.BuildActor().Build()
	c1 := crw.BuildActor().Build()

	a1.AddChild(c1)
	c1.RemoveSelf()

	actual := len(a1.Children)

	if actual != expected {
		t.Fatalf("Failed Remove self actual parent children count %d, expected %d", actual, expected)
	}

	log.Output(1, "[PASS]: TestRemoveSelf")
}

func TestWindowPositionForChildActorZero(t *testing.T) {
	a1 := crw.BuildActor().WithPosition(0, 0, 0).Build()
	a2 := crw.BuildActor().WithPosition(0, 0, 0).Build()

	a1.AddChild(a2)

	if a1.GetWindowPosition().X != 0 {
		t.Fatalf("Expected parent X to remain 0, actual %f", a1.GetWindowPosition().X)
	}
	if a1.GetWindowPosition().Y != 0 {
		t.Fatalf("Expected parent Y to remain 0, actual %f", a1.GetWindowPosition().Y)
	}
	if a2.GetWindowPosition().X != 0 {
		t.Fatalf("Expected child X to remain 0, actual %f", a2.GetWindowPosition().X)
	}
	if a2.GetWindowPosition().Y != 0 {
		t.Fatalf("Expected child Y to remain 0, actual %f", a2.GetWindowPosition().Y)
	}

	log.Output(1, "[PASS]: TestWindowPositionForChildActorZero")
}

func TestWindowPositionForChildActorWhenAdded(t *testing.T) {
	a1 := crw.BuildActor().WithPosition(10, 20, 0).Build()
	a2 := crw.BuildActor().WithPosition(5, 10, 0).Build()

	a1.AddChild(a2)

	if a1.GetWindowPosition().X != 10 {
		t.Fatalf("Expected parent X to remain 10, actual %f", a1.GetWindowPosition().X)
	}
	if a1.GetWindowPosition().Y != 20 {
		t.Fatalf("Expected parent Y to remain 20, actual %f", a1.GetWindowPosition().Y)
	}
	if a2.GetWindowPosition().X != 15 {
		t.Fatalf("Expected child X be 15, actual %f", a2.GetWindowPosition().X)
	}
	if a2.GetWindowPosition().Y != 30 {
		t.Fatalf("Expected child Y to be 30, actual %f", a2.GetWindowPosition().Y)
	}

	log.Output(1, "[PASS]: TestWindowPositionForChildActorWhenAdded")
}

func TestWindowPositionForChildActorWhenUpdated(t *testing.T) {
	a1 := crw.BuildActor().WithPosition(10, 20, 0).Build()
	a2 := crw.BuildActor().WithPosition(5, 10, 0).Build()

	a1.AddChild(a2)
	a1.SetX(15)
	a1.SetY(25)

	if a1.GetWindowPosition().X != 15 {
		t.Fatalf("Expected parent X to remain 15, actual %f", a1.GetWindowPosition().X)
	}
	if a1.GetWindowPosition().Y != 25 {
		t.Fatalf("Expected parent Y to remain 20, actual %f", a1.GetWindowPosition().Y)
	}
	if a2.GetWindowPosition().X != 20 {
		t.Fatalf("Expected child X be 20, actual %f", a2.GetWindowPosition().X)
	}
	if a2.GetWindowPosition().Y != 35 {
		t.Fatalf("Expected child Y to be 35, actual %f", a2.GetWindowPosition().Y)
	}

	log.Output(1, "[PASS]: TestWindowPositionForChildActorWhenUpdated")
}

// Sprites
func TestWindowPositionForChildSpriteZero(t *testing.T) {
	a1 := crw.BuildActor().WithPosition(0, 0, 0).Build()
	sprite1 := crw.BuildSprite().WithDestRect(0, 0, 1, 1).Build()

	a1.AddSprite(sprite1)

	if sprite1.GetWindowDestRect().X != 0 {
		t.Fatalf("Expected child sprite X to remain 0, actual %f", sprite1.GetWindowDestRect().X)
	}
	if sprite1.GetWindowDestRect().Y != 0 {
		t.Fatalf("Expected child sprite Y to remain 0, actual %f", sprite1.GetWindowDestRect().Y)
	}

	log.Output(1, "[PASS]: TestWindowPositionForChildSpriteZero")
}

func TestWindowPositionForChildSpriteWhenAdded(t *testing.T) {
	a1 := crw.BuildActor().WithPosition(10, 20, 0).Build()
	sprite1 := crw.BuildSprite().WithDestRect(5, 10, 1, 1).Build()

	a1.AddSprite(sprite1)

	if sprite1.GetWindowDestRect().X != 15 {
		t.Fatalf("Expected child sprite X to be 15, actual %f", sprite1.GetWindowDestRect().X)
	}
	if sprite1.GetWindowDestRect().Y != 30 {
		t.Fatalf("Expected child sprite Y to be 30, actual %f", sprite1.GetWindowDestRect().Y)
	}

	log.Output(1, "[PASS]: TestWindowPositionForChildSpriteWhenAdded")
}

func TestWindowPositionForChildSpriteWhenUpdated(t *testing.T) {
	a1 := crw.BuildActor().WithPosition(0, 0, 0).Build()
	sprite1 := crw.BuildSprite().WithDestRect(5, 10, 1, 1).Build()

	a1.AddSprite(sprite1)
	a1.SetX(10)
	a1.SetY(20)

	if sprite1.GetWindowDestRect().X != 15 {
		t.Fatalf("Expected child sprite X to be 15, actual %f", sprite1.GetWindowDestRect().X)
	}
	if sprite1.GetWindowDestRect().Y != 30 {
		t.Fatalf("Expected child sprite Y to be 30, actual %f", sprite1.GetWindowDestRect().Y)
	}

	log.Output(1, "[PASS]: TestWindowPositionForChildSpriteWhenUpdated")
}

// Animation
func TestWindowPositionForChildAnimationZero(t *testing.T) {
	a1 := crw.BuildActor().WithPosition(0, 0, 0).Build()
	anim1 := crw.BuildAnimation().WithDestRect(0, 0, 1, 1).Build()

	a1.AddAnimation(anim1)

	if anim1.GetWindowDestRect().X != 0 {
		t.Fatalf("Expected child animation X to remain 0, actual %f", anim1.GetWindowDestRect().X)
	}
	if anim1.GetWindowDestRect().Y != 0 {
		t.Fatalf("Expected child animation Y to remain 0, actual %f", anim1.GetWindowDestRect().Y)
	}

	log.Output(1, "[PASS]: TestWindowPositionForChildAnimationZero")
}

func TestWindowPositionForChildAnimationWhenAdded(t *testing.T) {
	a1 := crw.BuildActor().WithPosition(10, 20, 0).Build()
	anim1 := crw.BuildAnimation().WithDestRect(5, 10, 1, 1).Build()

	a1.AddAnimation(anim1)

	if anim1.GetWindowDestRect().X != 15 {
		t.Fatalf("Expected child animation X to be 15, actual %f", anim1.GetWindowDestRect().X)
	}
	if anim1.GetWindowDestRect().Y != 30 {
		t.Fatalf("Expected child animation Y to be 30, actual %f", anim1.GetWindowDestRect().Y)
	}

	log.Output(1, "[PASS]: TestWindowPositionForChildAnimationWhenAdded")
}

func TestWindowPositionForChildAnimationWhenUpdated(t *testing.T) {
	a1 := crw.BuildActor().WithPosition(0, 0, 0).Build()
	animation1 := crw.BuildAnimation().WithDestRect(5, 10, 1, 1).Build()

	a1.AddAnimation(animation1)
	a1.SetX(10)
	a1.SetY(20)

	if animation1.GetWindowDestRect().X != 15 {
		t.Fatalf("Expected child animation X to be 15, actual %f", animation1.GetWindowDestRect().X)
	}
	if animation1.GetWindowDestRect().Y != 30 {
		t.Fatalf("Expected child animation Y to be 30, actual %f", animation1.GetWindowDestRect().Y)
	}

	log.Output(1, "[PASS]: TestWindowPositionForChildAnimationWhenUpdated")
}
