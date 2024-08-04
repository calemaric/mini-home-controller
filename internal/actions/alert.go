package actions

const ALERT = "alert"

type AlertAction struct {
}

func (a *AlertAction) GetType() string {
	return ALERT
}
func (a *AlertAction) GetTemplateName() string {
	return ALERT
}

func (a *AlertAction) Execute() *ActionResult {
	return &ActionResult{}

}
