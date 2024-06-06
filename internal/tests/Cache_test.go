package tests

import (
	"crowform/internal/cache"
	"log"
	"testing"
)

func TestCacheSettingsInt(t *testing.T) {

	var val int = 1
	var name string = "s1"

	cache.SetSetting(name, val)

	res := cache.GetSetting[int](name)

	if res != val {
		t.Fatalf("incorrect int value cached actual %d, expected %d", res, val)
	}

	log.Output(1, "[PASS]: TestCacheSettingsInt")
}

func TestCacheSettingsFloat32(t *testing.T) {

	var val float32 = 1
	var name string = "s1"

	cache.SetSetting(name, val)

	res := cache.GetSetting[float32](name)

	if res != val {
		t.Fatalf("incorrect float value cached actual %f, expected %f", res, val)
	}

	log.Output(1, "[PASS]: TestCacheSettingsFloat32")
}

func TestCacheSettingsString(t *testing.T) {

	var val string = "my test string"
	var name string = "s1"

	cache.SetSetting(name, val)

	res := cache.GetSetting[string](name)

	if res != val {
		t.Fatalf("incorrect string value cached actual %s, expected %s", res, val)
	}

	log.Output(1, "[PASS]: TestCacheSettingsString")
}
