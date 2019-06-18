package keybindings

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/style"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

////////////////////////////////////////////////////////////////////
type QuitHandler struct {
	GlobalHandler
}

func NewQuitHandler() *QuitHandler {
	handler := &QuitHandler{}
	handler.Index = 0
	return handler
}

func (h QuitHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		return gocui.ErrQuit
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type CopyHandler struct {
	GlobalHandler
	Content   *views.ItemWidget
	StatusBar *views.StatusbarWidget
}

func NewCopyHandler(content *views.ItemWidget, statusbar *views.StatusbarWidget) *CopyHandler {
	handler := &CopyHandler{
		Content:   content,
		StatusBar: statusbar,
	}
	handler.Index = 1
	return handler
}

func (h CopyHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		err := clipboard.WriteAll(h.Content.GetContent())
		if err != nil {
			h.StatusBar.Status("Failed to copy to clipboard", false)
			return nil
		}
		h.StatusBar.Status("Current resource's JSON copied to clipboard", false)
		return nil
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type FullscreenHandler struct {
	GlobalHandler
	List         *views.ListWidget
	IsFullscreen *bool
	Content      *views.ItemWidget
}

func NewFullscreenHandler(list *views.ListWidget, content *views.ItemWidget, isFullscreen *bool) *FullscreenHandler {
	handler := &FullscreenHandler{
		List:         list,
		Content:      content,
		IsFullscreen: isFullscreen,
	}
	handler.Index = 3
	return handler
}

func (h FullscreenHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		tmp := toggle(*h.IsFullscreen)
		h.IsFullscreen = &tmp // memory leak?
		if *h.IsFullscreen {
			g.Cursor = true
			maxX, maxY := g.Size()
			v, _ := g.SetView("fullscreenContent", 0, 0, maxX, maxY)
			v.Editable = true
			v.Frame = false
			v.Wrap = true
			keyBindings := GetKeyBindingsAsStrings()
			v.Title = fmt.Sprintf("JSON Response - Fullscreen (%s to exit)", strings.ToUpper(keyBindings["fullscreen"]))

			content := h.Content.GetContent()
			fmt.Fprint(v, style.ColorJSON(content))

			g.SetCurrentView("fullscreenContent")
		} else {
			g.Cursor = false
			g.DeleteView("fullscreenContent")
			g.SetCurrentView("listWidget")
		}
		return nil
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type HelpHandler struct {
	GlobalHandler
	ShowHelp *bool
}

func NewHelpHandler(showHelp *bool) *HelpHandler {
	handler := &HelpHandler{
		ShowHelp: showHelp,
	}
	handler.Index = 4
	return handler
}

func (h HelpHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		tmp := toggle(*h.ShowHelp)
		h.ShowHelp = &tmp // memory leak?

		// If we're up and running clear and redraw the view
		// if w.g != nil {
		if *h.ShowHelp {
			v, err := g.SetView("helppopup", 1, 1, 145, 40)
			g.SetCurrentView("helppopup")
			if err != nil && err != gocui.ErrUnknownView {
				panic(err)
			}
			keyBindings := GetKeyBindingsAsStrings()
			views.DrawHelp(keyBindings, v)
		} else {
			g.DeleteView("helppopup")
			g.SetCurrentView("listWidget")
		}
		return nil
	}
}

////////////////////////////////////////////////////////////////////


////////////////////////////////////////////////////////////////////
type EditHandler struct {
	GlobalHandler
	Content   *views.ItemWidget
	StatusBar *views.StatusbarWidget
}

func NewEditHandler(content *views.ItemWidget, statusbar *views.StatusbarWidget) *CopyHandler {
	handler := &CopyHandler{
		Content:   content,
		StatusBar: statusbar,
	}
	handler.Index = 24
	return handler
}

func (h EditHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		tmpFile, err := ioutil.TempFile(os.TempDir(), "azbrowse-")
		if err != nil {
			log.Panicln("Cannot create temporary file", err)
		}

		err = tmpFile.Close()
		if err != nil {
			log.Panicln("Cannot close temporary file", err)
		}

		newFilename := tmpFile.Name() + ".json"
		err = os.Rename(tmpFile.Name(), newFilename)
		if err != nil {
			log.Panicln("Cannot rename temporary file", err)
		}

		tmpFile, err = os.OpenFile(newFilename, os.O_RDWR, 0)
		if err != nil {
			log.Panicln("Cannot open renamed temporary file", err)
		}

		// Remember to clean up the file afterwards
		// defer os.Remove(newFilename)

		clipboard.WriteAll(h.Content.GetContent())

		tmpFile.WriteString(h.Content.GetContent())
		tmpFile.Close()

		h.StatusBar.Status("Opening JSON in editor...", false)
		err = OpenEditor(tmpFile.Name())
		if err != nil {
			log.Panicln("Cannot open editor", err)
		}
		h.StatusBar.Status("Opened JSON in editor", false)

		return nil
	}
}

////////////////////////////////////////////////////////////////////


////////////////////////////////////////////////////////////////////
type ConfirmDeleteHandler struct {
	GlobalHandler
	notificationWidget *views.NotificationWidget
}

func NewConfirmDeleteHandler(notificationWidget *views.NotificationWidget) *ConfirmDeleteHandler {
	handler := &ConfirmDeleteHandler{
		notificationWidget: notificationWidget,
	}
	handler.Index = 22
	return handler
}

func (h *ConfirmDeleteHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.notificationWidget.ConfirmDelete()
		return nil
	}
}

////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
type ClearPendingDeleteHandler struct {
	GlobalHandler
	notificationWidget *views.NotificationWidget
}

func NewClearPendingDeleteHandler(notificationWidget *views.NotificationWidget) *ClearPendingDeleteHandler {
	handler := &ClearPendingDeleteHandler{
		notificationWidget: notificationWidget,
	}
	handler.Index = 23
	return handler
}

func (h *ClearPendingDeleteHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		h.notificationWidget.ClearPendingDeletes()
		return nil
	}
}

////////////////////////////////////////////////////////////////////
