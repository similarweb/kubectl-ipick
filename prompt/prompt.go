package prompt

import (
	"errors"

	"github.com/AlecAivazis/survey/v2"
	log "github.com/sirupsen/logrus"
)

// PickSelectionFromData return selected data
func PickSelectionFromData(text string, data []string) (int, error) {

	selected := -1
	prompt := &survey.Select{
		Message:       text,
		Options:       data,
		PageSize:      50,
		FilterMessage: "*",
	}
	err := survey.AskOne(prompt, &selected)
	if err != nil {
		return selected, err
	}

	if selected == -1 {
		return selected, errors.New("promp canceled")
	}

	log.WithField("selection", selected).Debug("user selection")
	return selected, nil
}
