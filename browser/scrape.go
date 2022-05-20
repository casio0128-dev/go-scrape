package browser

import (
	"fmt"
	"github.com/sclevine/agouti"
	act "go-scrape/action"
	"go-scrape/profile"
	"time"
)

func Do(page *agouti.Page, prof *profile.Profile) error {
	operation := prof.Operation
	if wakeUpTime, err := operation.WakeUpTime.DateTime(); err != nil {
		return err
	} else {
		if time.Now().Before(wakeUpTime) {
			<-time.After(time.Until(wakeUpTime))
		}
	}

	target := operation.Url
	if err := page.Navigate(target); err != nil {
		return err
	}

	for _, ctrls := range operation.Control {
		for actName, content := range ctrls {
			fmt.Printf("act.ParseAction(%v, %v, %v)\n", actName, prof, content)
			action := act.ParseAction(actName, prof, content)
			fmt.Println("action", action)
			if err := action.Do(page); err != nil {
				return err
			}
		}
	}
	return nil
}
