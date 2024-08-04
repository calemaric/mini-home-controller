package controls

import "calemaric/mini-home-controller/internal/actions"

type AlertControl struct {
}

func (c *AlertControl) GetTemplateName() string {
	return "Alert"

}

func (c *AlertControl) ActionType() string {
	return actions.ALERT
}
