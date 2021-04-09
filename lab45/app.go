package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	a   = app.New()
	u   User
	err error
	db  *gorm.DB
)

func ui() {
	w.Resize(fyne.NewSize(600, 300))
	var (
		authTab = AuthTabInit()
	)
	tabs.Append(authTab)
	tabs.SelectTab(authTab)

	w.SetContent(container.NewVBox(
		tabs,
	))

	w.ShowAndRun()
}

func orm() {
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	setup(db)
}

func main() {
	orm()
	ui()
}
