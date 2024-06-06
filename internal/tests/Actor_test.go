package tests

import (
	"crowform/pkg"
	"log"
	"testing"
	"time"
)

func TestActorQueryIncludesSelf(t *testing.T) {
	a1 := pkg.BuildActor().Build()
	a2 := pkg.BuildActor().Build()
	a3 := pkg.BuildActor().Build()

	var Solid pkg.QueryAttribute = "TEST_SOLID"
	var Attacks pkg.QueryAttribute = "TEST_ATTACKS"

	a1.AddQueryAttr(Solid)
	a2.AddQueryAttr(Solid)
	a3.AddQueryAttr(Attacks)

	a1.AddChild(a2)
	a1.AddChild(a3)

	q := []pkg.QueryAttribute{Solid}
	res := a1.QueryAny(q)

	if len(res) != 2 {
		t.Fatalf("Test query returns incorrect results actual %d, expected %d", len(res), 2)
	}

	log.Output(1, "[PASS]: TestActorQueryIncludesSelf")
}

func TestActorQueryChildrenWithoutSelf(t *testing.T) {
	a1 := pkg.BuildActor().Build()
	a2 := pkg.BuildActor().Build()
	a3 := pkg.BuildActor().Build()

	var Solid pkg.QueryAttribute = "TEST_SOLID"
	var Attacks pkg.QueryAttribute = "TEST_ATTACKS"

	a2.AddQueryAttr(Solid)
	a3.AddQueryAttr(Attacks)

	a1.AddChild(a2)
	a1.AddChild(a3)

	q := []pkg.QueryAttribute{Solid}
	res := a1.QueryAny(q)

	if len(res) != 1 {
		t.Fatalf("Test query returns incorrect results actual %d, expected %d", len(res), 1)
	}

	log.Output(1, "[PASS]: TestActorQueryChildrenWithoutSelf")
}

func TestActorAddChildZOrderAsc(t *testing.T) {
	a1 := pkg.BuildActor().Build()
	a2 := pkg.BuildActor().WithPosition(0, 0, 1).Build()
	a3 := pkg.BuildActor().WithPosition(0, 0, 2).Build()

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
	a1 := pkg.BuildActor().Build()
	a2 := pkg.BuildActor().WithPosition(0, 0, 2).Build()
	a3 := pkg.BuildActor().WithPosition(0, 0, 1).Build()

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
	a1 := pkg.BuildActor().Build()
	a2 := pkg.BuildActor().WithPosition(0, 0, 0).Build()
	a3 := pkg.BuildActor().WithPosition(0, 0, 0).Build()

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

	a1 := pkg.BuildActor().
		WithAction(
			func(actor *pkg.Actor) bool {
				return true
			},
			func(deltaTime time.Duration, actor *pkg.Actor, done pkg.ActorActionDone) {
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

	a1 := pkg.BuildActor().
		WithAction(
			func(actor *pkg.Actor) bool {
				return true
			},
			func(deltaTime time.Duration, actor *pkg.Actor, done pkg.ActorActionDone) {
				actual = append(actual, 1)
				done()
			}).
		WithAction(
			func(actor *pkg.Actor) bool {
				return true
			},
			func(deltaTime time.Duration, actor *pkg.Actor, done pkg.ActorActionDone) {
				actual = append(actual, 2)
				done()
			}).
		WithAction(
			func(actor *pkg.Actor) bool {
				return true
			},
			func(deltaTime time.Duration, actor *pkg.Actor, done pkg.ActorActionDone) {
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

	a1 := pkg.BuildActor().Build()
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

	a1 := pkg.BuildActor().Build()
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

	a1 := pkg.BuildActor().Build()
	c1 := pkg.BuildActor().Build()

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

	a1 := pkg.BuildActor().Build()
	c1 := pkg.BuildActor().Build()

	a1.AddChild(c1)
	c1.RemoveSelf()

	actual := len(a1.Children)

	if actual != expected {
		t.Fatalf("Failed Remove self actual parent children count %d, expected %d", actual, expected)
	}

	log.Output(1, "[PASS]: TestRemoveSelf")
}
