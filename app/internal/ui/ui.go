package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/gopherzz/gg/app/pkg/config"
	"github.com/gopherzz/gg/app/pkg/pkggo"
	"github.com/gopherzz/gg/app/pkg/pkggo/models"
	"github.com/rivo/tview"
	"golang.design/x/clipboard"
)

type Ui struct {
	Config *config.Config
	app    *tview.Application
}

func NewUi(config *config.Config) *Ui {
	ui := &Ui{}
	ui.app = tview.NewApplication()
	ui.Config = config
	return ui
}

func (ui Ui) Render() error {
	var query string
	var packages []models.GoPackage
	var err error

	err = clipboard.Init()
	if err != nil {
		return err
	}

	// initrialize tui elements
	pages := tview.NewPages()
	form := tview.NewForm()
	flex := tview.NewFlex()
	packagesList := tview.NewList().ShowSecondaryText(false)
	packageInfo := tview.NewTextView()
	debugInfo := tview.NewTextView()
	debugInfo.SetText(fmt.Sprintf("%v", ui.Config))
	// setup pages
	pages.AddPage("Search", form, true, true)
	pages.AddPage("List", flex, true, false)

	// TODO: fix thing when cannot roll list after change page
	// setup packages list
	packagesList.SetSelectedFunc(func(i int, s1, s2 string, r rune) {
		pkg := &packages[i]
		packageInfo.Clear()
		packageInfo.SetText(fmt.Sprintf("Name: %s\nUrl: %s\nDescription: %s", pkg.Name, pkg.Url, pkg.ShortDesc))
		clipboard.Write(clipboard.FmtText, []byte(fmt.Sprintf("go get %s", pkg.Url[1:len(pkg.Url)-1])))
	})

	// setup search form
	form.AddInputField("Search", "", 20, nil, func(input string) {
		query = input
	})
	form.AddButton("Find", func() {
		// packagesList.Clear()
		pages.SwitchToPage("List")
		packagesList.SetCurrentItem(0)
		packages, err = pkggo.FindPackages(query)
		if err != nil {
			return
		}
		for i, pkg := range packages {
			packagesList.AddItem(pkg.Name+" "+pkg.Url, "", rune(49+i), nil)
		}
		// packageInfo.SetText(fmt.Sprintf("Name: %s\nUrl: %s\nDescription: %s", packages[0].Name, packages[0].Url, packages[0].ShortDesc))
	})

	// setup packages screen
	flex.SetDirection(tview.FlexRow).
		AddItem(tview.NewFlex().
			AddItem(packagesList, 0, 4, true).
			AddItem(packageInfo, 0, 5, false), 0, 6, false).
		AddItem(debugInfo, 0, 1, false)

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 13:
			ui.app.Stop()
		}
		return event
	})

	return ui.app.SetRoot(pages, true).EnableMouse(ui.Config.Local.MouseEnabled).ForceDraw().Run()
}
