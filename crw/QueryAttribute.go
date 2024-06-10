package crw

type QueryAttribute string

const (
	// Internal Attribute for handling pause
	queryAttribute_UPDATES_WHEN_PAUSED QueryAttribute = "queryAttribute_UPDATES_WHEN_PAUSED"
	// Internal Attribute for handling mouse down
	queryAttribute_RECEIVES_MOUSE_DOWN_EVENT QueryAttribute = "queryAttribute_RECEIVES_MOUSE_DOWN_EVENT"
	// Internal Attribute for handling mouse up
	queryAttribute_RECEIVES_MOUSE_UP_EVENT QueryAttribute = "queryAttribute_RECEIVES_MOUSE_UP_EVENT"
	// Internal Attribute for handling mouse move
	queryAttribute_RECEIVES_MOUSE_MOVE_EVENT QueryAttribute = "queryAttribute_RECEIVES_MOUSE_MOVE_EVENTS"
	// Internal Attribute for handling mouse enter
	queryAttribute_RECEIVES_MOUSE_ENTER_EVENT QueryAttribute = "queryAttribute_RECEIVES_MOUSE_ENTER_EVENT"
	// Internal Attribute for handling mouse exit
	queryAttribute_RECEIVES_MOUSE_EXIT_EVENT QueryAttribute = "queryAttribute_RECEIVES_MOUSE_EXIT_EVENT"
	// Internal Attribute for handling key input
	queryAttribute_RECEIVES_KEY_EVENT QueryAttribute = "queryAttribute_RECEIVES_KEY_EVENT"
)
