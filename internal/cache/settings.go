package cache

const (
	SettingName_AssetDirectory string = "SettingName_AssetDirectory"
)

/** settings **/

var settings map[string]interface{} = make(map[string]interface{})

func GetSetting[T int | float32 | string](name string) T {
	val, foundInSetting := settings[name]
	var result T

	if !foundInSetting {
		return result
	}

	typedV, ok := val.(T)
	if ok {
		result = typedV
	}

	return result
}

func SetSetting[T int | float32 | string](name string, value T) T {
	settings[name] = value
	return value
}

func ResetSettings() {
	for k := range settings {
		delete(settings, k)
	}
}
