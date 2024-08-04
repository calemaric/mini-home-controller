package actions

import (
	"errors"
)

type Action interface {
	GetType() string
	GetTemplateName() string
	Execute() *ActionResult
}

type ActionResult struct {
}

var actions []Action

func InitializeActions() {
	actions = []Action{&AlertAction{}}
}

func GetActionByType(actionType string) (Action, error) {
	for _, action := range actions {
		if action.GetType() == actionType {
			return action, nil
		}
	}

	return nil, errors.New("No action found for actionType " + actionType)

}
