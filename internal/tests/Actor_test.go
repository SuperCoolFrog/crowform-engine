package tests

import (
	"crowform/pkg"
	"log"
	"testing"
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
