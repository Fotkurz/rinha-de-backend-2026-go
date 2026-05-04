package ready

var (
	ready bool = false
)

func IsReady() bool {
	return ready
}

func ToggleReady() {
	ready = !ready
}
