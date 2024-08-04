package controls

type Control interface {
	GetTemplateName() string
	ActionType() string
}

var controls []Control

func InitializeControls() {
	controls = []Control{&AlertControl{}}
}

func GetAllControls() []Control {
	return controls
}
