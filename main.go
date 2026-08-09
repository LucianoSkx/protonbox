package main

import (
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type gui struct {
	app  fyne.App
	win  fyne.Window
	lang string

	list         *widget.List
	search       *widget.Entry
	logo         *canvas.Image
	langRadio    *widget.RadioGroup
	configBtn    *widget.Button
	aboutBtn     *widget.Button
	detailTitle  *widget.Label
	detailCat    *widget.Label
	detailCompat *widget.Label
	detailDesc   *widget.Label
	detailCmd    *widget.Label
	copyBtn      *widget.Button
	status       *widget.Label

	langHeader  *fyne.MenuItem
	langPT      *fyne.MenuItem
	langEN      *fyne.MenuItem
	themeSystem *fyne.MenuItem
	themeLight  *fyne.MenuItem
	themeDark   *fyne.MenuItem
	copyOnClick *fyne.MenuItem

	combLabel   *widget.Label
	combCount   *widget.Label
	combCopyBtn *widget.Button
	clearBtn    *widget.Button

	all      []Command
	filtered []int
	selected map[int]bool
	current  int
	selID    int
}

func main() {
	a := app.NewWithID("br.com.protoncommands")
	a.SetIcon(resourceIcon)
	w := a.NewWindow("Proton Command")
	w.Resize(fyne.NewSize(980, 640))

	g := &gui{app: a, win: w, all: commands(), lang: "pt", selected: map[int]bool{}}
	if l := a.Preferences().StringWithFallback("lang", "pt"); l == "en" || l == "pt" {
		g.lang = l
	}
	g.filtered = make([]int, len(g.all))
	for i := range g.all {
		g.filtered[i] = i
	}
	g.build()
	w.ShowAndRun()
}

func (g *gui) tr(key string) string {
	if g.lang == "en" {
		if s, ok := enTexts[key]; ok {
			return s
		}
	}
	return ptTexts[key]
}

func (g *gui) t(l Localized) string {
	if g.lang == "en" {
		return l.EN
	}
	return l.PT
}

func (g *gui) cmd(c Command) string {
	if g.lang == "en" && c.CommandEN != "" {
		return c.CommandEN
	}
	return c.Command
}

func (g *gui) build() {
	g.search = widget.NewEntry()
	g.search.OnChanged = func(_ string) { g.applyFilter() }

	g.list = widget.NewList(
		func() int { return len(g.filtered) },
		func() fyne.CanvasObject {
			check := widget.NewCheck("", nil)
			return container.NewHBox(
				check,
				container.NewVBox(
					widget.NewLabel(""), // título
					widget.NewLabel(""), // categoria
				),
			)
		},
		func(id int, obj fyne.CanvasObject) {
			idx := g.filtered[id]
			c := g.all[idx]
			box := obj.(*fyne.Container)
			check := box.Objects[0].(*widget.Check)
			inner := box.Objects[1].(*fyne.Container)
			title := inner.Objects[0].(*widget.Label)
			cat := inner.Objects[1].(*widget.Label)
			title.SetText(g.t(c.Title))
			title.TextStyle = fyne.TextStyle{Bold: true}
			cat.SetText(g.t(c.Category) + " · " + g.t(c.Compat))
			cat.TextStyle = fyne.TextStyle{Italic: true}
			check.OnChanged = nil
			check.SetChecked(g.selected[idx])
			check.OnChanged = func(v bool) { g.toggle(idx, v) }
		},
	)
	g.list.OnSelected = func(id widget.ListItemID) {
		g.selID = id
		g.selectCommand(id)
	}
	g.list.OnUnselected = func(_ widget.ListItemID) {
		g.selID = -1
		g.selectCommand(-1)
	}

	g.detailTitle = widget.NewLabel("")
	g.detailTitle.TextStyle = fyne.TextStyle{Bold: true}
	g.detailTitle.Wrapping = fyne.TextWrapWord

	g.detailCat = widget.NewLabel("")
	g.detailCat.TextStyle = fyne.TextStyle{Italic: true}
	g.detailCat.Wrapping = fyne.TextWrapWord

	g.detailCompat = widget.NewLabel("")
	g.detailCompat.TextStyle = fyne.TextStyle{Italic: true}
	g.detailCompat.Wrapping = fyne.TextWrapWord

	g.detailCmd = widget.NewLabel("")
	g.detailCmd.TextStyle = fyne.TextStyle{Monospace: true, Bold: true}
	g.detailCmd.Wrapping = fyne.TextWrapWord

	g.detailDesc = widget.NewLabel("")
	g.detailDesc.Wrapping = fyne.TextWrapWord

	g.copyBtn = widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
		if g.current < 0 {
			return
		}
		g.copyCurrent()
	})
	g.copyBtn.Importance = widget.HighImportance

	g.status = widget.NewLabel("")
	g.status.TextStyle = fyne.TextStyle{Italic: true}

	detail := container.NewBorder(
		container.NewVBox(
			g.detailTitle,
			g.detailCat,
			g.detailCompat,
			widget.NewSeparator(),
			g.detailCmd,
			widget.NewSeparator(),
		),
		container.NewVBox(g.copyBtn, g.status),
		nil, nil,
		container.NewVScroll(g.detailDesc),
	)
	detail.Resize(fyne.NewSize(520, 560))

	left := container.NewBorder(
		g.search, nil, nil, nil,
		g.list,
	)

	split := container.NewHSplit(left, detail)
	split.Offset = 0.42

	configMenu := g.buildConfigMenu()

	g.logo = canvas.NewImageFromResource(resourceLogo)
	g.logo.FillMode = canvas.ImageFillContain
	g.logo.SetMinSize(fyne.NewSize(280, 80))
	g.logo.Resize(fyne.NewSize(280, 80))

	g.langRadio = widget.NewRadioGroup([]string{"Português", "English"}, func(v string) {
		if v == "English" {
			g.setLang("en", true)
		} else if v != "" {
			g.setLang("pt", true)
		}
	})
	g.langRadio.Horizontal = true

	var configBtn *widget.Button
	configBtn = widget.NewButton("", func() {
		pop := widget.NewPopUpMenu(configMenu, g.win.Canvas())
		pos := fyne.CurrentApp().Driver().AbsolutePositionForObject(configBtn)
		pop.ShowAtPosition(pos.AddXY(0, configBtn.Size().Height))
	})
	g.configBtn = configBtn

	g.aboutBtn = widget.NewButton("", func() {
		dialog.NewInformation(g.tr("about"), g.tr("aboutBody"), g.win).Show()
	})

	titleBar := container.NewBorder(
		nil, nil, nil,
		container.NewHBox(g.configBtn, g.aboutBtn),
		container.NewCenter(g.logo),
	)

	langBar := container.NewCenter(g.langRadio)

	g.combLabel = widget.NewLabel("")
	g.combLabel.Wrapping = fyne.TextWrapWord

	g.combCount = widget.NewLabel("")
	g.combCount.TextStyle = fyne.TextStyle{Italic: true}

	g.combCopyBtn = widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
		g.copyCombination()
	})
	g.combCopyBtn.Importance = widget.HighImportance

	g.clearBtn = widget.NewButton("", func() {
		g.clearSelection()
	})

	combBar := container.NewBorder(
		container.NewHBox(
			g.combCount,
			layout.NewSpacer(),
			g.clearBtn,
		),
		nil, nil, nil,
		container.NewVBox(
			g.combLabel,
			g.combCopyBtn,
		),
	)

	g.win.SetContent(container.NewBorder(
		container.NewVBox(titleBar, langBar),
		combBar, nil, nil, split,
	))

	g.applyLang()
	g.copyOnClick.Checked = g.app.Preferences().BoolWithFallback("copyOnClick", false)
	g.setTheme(g.app.Preferences().StringWithFallback("theme", "system"), false)
	g.updateCombination()
	if len(g.filtered) > 0 {
		g.list.Select(0)
	}
}

func (g *gui) buildConfigMenu() *fyne.Menu {
	g.langHeader = fyne.NewMenuItem("", nil)
	g.langHeader.Disabled = true
	g.langPT = fyne.NewMenuItem("Português", func() {
		g.setLang("pt", true)
	})
	g.langEN = fyne.NewMenuItem("English", func() {
		g.setLang("en", true)
	})
	g.themeSystem = fyne.NewMenuItem("", func() {
		g.setTheme("system", true)
	})
	g.themeLight = fyne.NewMenuItem("", func() {
		g.setTheme("light", true)
	})
	g.themeDark = fyne.NewMenuItem("", func() {
		g.setTheme("dark", true)
	})
	g.copyOnClick = fyne.NewMenuItem("", func() {
		g.copyOnClick.Checked = !g.copyOnClick.Checked
		g.app.Preferences().SetBool("copyOnClick", g.copyOnClick.Checked)
	})
	return fyne.NewMenu("",
		g.langHeader,
		g.langPT,
		g.langEN,
		fyne.NewMenuItemSeparator(),
		g.themeSystem,
		g.themeLight,
		g.themeDark,
		fyne.NewMenuItemSeparator(),
		g.copyOnClick,
	)
}

func (g *gui) setLang(lang string, persist bool) {
	if lang == "en" || lang == "pt" {
		g.lang = lang
	}
	if persist {
		g.app.Preferences().SetString("lang", g.lang)
	}
	g.applyLang()
	if g.selID >= 0 && g.selID < len(g.filtered) {
		g.list.Select(g.selID)
	} else if len(g.filtered) > 0 {
		g.list.Select(0)
	}
	g.list.Refresh()
}

func (g *gui) applyLang() {
	if g.lang == "en" {
		g.langRadio.SetSelected("English")
	} else {
		g.langRadio.SetSelected("Português")
	}
	g.search.SetPlaceHolder(g.tr("searchPlaceholder"))
	g.configBtn.Text = g.tr("settings")
	g.configBtn.Refresh()
	g.aboutBtn.Text = g.tr("about")
	g.aboutBtn.Refresh()
	g.copyBtn.Text = g.tr("copyCommand")
	g.copyBtn.Refresh()
	g.langHeader.Label = g.tr("language")
	g.langPT.Checked = g.lang == "pt"
	g.langEN.Checked = g.lang == "en"
	g.themeSystem.Label = g.tr("themeSystem")
	g.themeLight.Label = g.tr("themeLight")
	g.themeDark.Label = g.tr("themeDark")
	g.copyOnClick.Label = g.tr("copyOnClick")
	g.combCopyBtn.Text = g.tr("copyCombination")
	g.combCopyBtn.Refresh()
	g.clearBtn.Text = g.tr("clearSelection")
	g.clearBtn.Refresh()
	g.updateCombination()
	g.refreshDetail()
}

func (g *gui) refreshDetail() {
	if g.current >= 0 {
		c := g.all[g.current]
		g.detailTitle.SetText(g.t(c.Title))
		g.detailCat.SetText(g.tr("category") + g.t(c.Category))
		g.detailCompat.SetText(g.tr("compatible") + g.t(c.Compat))
		g.detailCmd.SetText(g.cmd(c))
		g.detailDesc.SetText(g.t(c.Description))
	} else {
		g.detailTitle.SetText(g.tr("noCommand"))
		g.detailCat.SetText("")
		g.detailCompat.SetText("")
		g.detailCmd.SetText("")
		g.detailDesc.SetText("")
	}
}

func (g *gui) setTheme(name string, persist bool) {
	switch name {
	case "light":
		g.app.Settings().SetTheme(theme.LightTheme())
	case "dark":
		g.app.Settings().SetTheme(theme.DarkTheme())
	default:
		g.app.Settings().SetTheme(theme.DefaultTheme())
	}
	g.themeSystem.Checked = name == "system"
	g.themeLight.Checked = name == "light"
	g.themeDark.Checked = name == "dark"
	if persist {
		g.app.Preferences().SetString("theme", name)
	}
}

func (g *gui) toggle(idx int, v bool) {
	if v {
		g.selected[idx] = true
	} else {
		delete(g.selected, idx)
	}
	g.updateCombination()
}

func (g *gui) isWrapper(c Command) bool {
	for _, w := range []string{"mangohud", "gamemoderun", "gamescope"} {
		if strings.HasPrefix(c.Command, w) {
			return true
		}
	}
	return false
}

func (g *gui) buildCombination() string {
	var wrappers, envs []string
	for i := range g.all {
		if !g.selected[i] {
			continue
		}
		c := g.all[i]
		base := strings.TrimSpace(strings.TrimSuffix(g.cmd(c), "%command%"))
		if g.isWrapper(c) {
			wrappers = append(wrappers, base+"%command%")
		} else {
			envs = append(envs, base)
		}
	}
	var parts []string
	parts = append(parts, wrappers...)
	parts = append(parts, envs...)
	if len(parts) == 0 {
		return ""
	}
	return strings.Join(parts, " ") + " %command%"
}

func (g *gui) updateCombination() {
	n := len(g.selected)
	comb := g.buildCombination()
	if n == 0 {
		g.combLabel.SetText(g.tr("noCommandSelected"))
		g.combCount.SetText("")
		g.combCopyBtn.Disable()
		return
	}
	g.combLabel.SetText(comb)
	g.combCount.SetText(g.tr("selectedCount") + strconv.Itoa(n))
	g.combCopyBtn.Enable()
}

func (g *gui) copyCombination() {
	comb := g.buildCombination()
	if comb == "" {
		return
	}
	g.win.Clipboard().SetContent(comb)
	g.status.SetText(g.tr("copied") + comb)
	g.status.Refresh()
}

func (g *gui) clearSelection() {
	g.selected = map[int]bool{}
	g.list.Refresh()
	g.updateCombination()
	g.status.SetText("")
}

func (g *gui) copyCurrent() {
	if g.current < 0 {
		return
	}
	g.win.Clipboard().SetContent(g.cmd(g.all[g.current]))
	g.status.SetText(g.tr("copied") + g.cmd(g.all[g.current]))
	g.status.Refresh()
}

func (g *gui) applyFilter() {
	query := strings.ToLower(strings.TrimSpace(g.search.Text))
	g.filtered = g.filtered[:0]
	for i, c := range g.all {
		hay := strings.ToLower(c.Command + " " + c.CommandEN + " " + g.t(c.Title) + " " + g.t(c.Category) + " " + g.t(c.Description) + " " + g.t(c.Compat))
		if query == "" || strings.Contains(hay, query) {
			g.filtered = append(g.filtered, i)
		}
	}
	g.selID = -1
	g.list.Refresh()
	if len(g.filtered) > 0 {
		g.list.Select(0)
	} else {
		g.clearDetail()
	}
}

func (g *gui) selectCommand(id widget.ListItemID) {
	if id < 0 || id >= len(g.filtered) {
		g.clearDetail()
		return
	}
	g.current = g.filtered[id]
	g.refreshDetail()
	g.status.SetText("")
	g.copyBtn.Enable()
	if g.copyOnClick.Checked {
		g.copyCurrent()
	}
}

func (g *gui) clearDetail() {
	g.current = -1
	g.refreshDetail()
	g.status.SetText("")
	g.copyBtn.Disable()
}
