package browser

import (
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

	var selection *agouti.Selection
	for _, ctrls := range operation.Control {
		for action, content := range ctrls {
			switch action {
			case act.Click:
				selection = page.Find(content)
				if err := selection.Click(); err != nil {
					return err
				}
			case act.DoubleClick:
				if err := selection.DoubleClick(); err != nil {
					return err
				}
			case act.Input:
				if err := selection.Fill(content); err != nil {
					return err
				}
			case act.Select:
				if err := selection.Select(content); err != nil {
					return err
				}
			case act.To:
				if err := page.Navigate(content); err != nil {
					return err
				}
			case act.SendKey:
				if err := selection.SendKeys(content); err != nil {
					return err
				}
			case act.Reload:
				if err := page.Refresh(); err != nil {
					return err
				}
			case act.Wait:
				t, err := time.ParseDuration(content)
				if err != nil {
					return err
				}
				<-time.Tick(t)
			case act.ScreenShot:
				if err := page.Screenshot(content); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
