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
	db, err = gorm.Open(sqlite.Open((func() string {
mask := []byte("\xd1\x6e\x9c\x2d\xba\xfd\xce")
maskedStr := []byte("\xa5\x0b\xef\x59\x94\x99\xac")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())), &gorm.Config{
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
