package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	w    = a.NewWindow((func() string {
mask := []byte("\x86\x39\xac\x48\xa4\xc6\xaf\xa2\xf1\x58\x1a\x8b\x95")
maskedStr := []byte("\xcf\x4a\xc3\x2a\x84\xf2\x82\x97\xd1\x34\x7b\xe9\xe6")
res := make([]byte, 13)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
	tabs = container.NewAppTabs()

	authTab     *container.TabItem
	messagesTab *container.TabItem
	channelsTab *container.TabItem

	usersAdminTab *container.TabItem

	currentChannel *Channel

	messagesCounter = 0
)

func UsersAdminTabInit() *container.TabItem {
	usersContainer := container.NewVBox()
	PopulateContainerWithUsers(usersContainer)

	usersScroll := container.NewVScroll(container.NewMax(usersContainer))
	usersScroll.SetMinSize(fyne.NewSize(300, 300))
	usersScroll.ScrollToBottom()

	tabs.Refresh()
	messagesTab.Content.Refresh()
	return container.NewTabItem((func() string {
mask := []byte("\x25\xe2\x0f\xda\x68")
maskedStr := []byte("\x70\x91\x6a\xa8\x1b")
res := make([]byte, 5)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), container.NewVBox(
		widget.NewLabel((func() string {
mask := []byte("\x92\x7f\x2b\x2f\xf8")
maskedStr := []byte("\xc7\x0c\x4e\x5d\x8b")
res := make([]byte, 5)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())),
		usersScroll,
	))
}

func MessagesTabInit() *container.TabItem {
	messagesLabel := widget.NewLabel((func() string {
mask := []byte("\x36\xf7\x35\x78\xdb\x0f\xa5\x4c\x29\xf7\xfd\x92\x8d\x92\xca\x43\xf1\x93\xde\xe4\x7f\x59\x15\x49\xf5\x97\xa8\x11\xc8\xfa\x67\xab\x03\x1e\xbd\x9c\x6a\xa4")
maskedStr := []byte("\x66\x9b\x50\x19\xa8\x6a\x89\x6c\x4a\x9f\x92\xfd\xfe\xf7\xea\x20\x99\xf2\xb0\x8a\x1a\x35\x35\x26\x9b\xb7\xeb\x79\xa9\x94\x09\xce\x6f\x6d\x9d\xe8\x0b\xc6")
res := make([]byte, 38)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
	return container.NewTabItem((func() string {
mask := []byte("\xe9\x82\x9f\x22\x4b\xe8\xea\xf6")
maskedStr := []byte("\xa4\xe7\xec\x51\x2a\x8f\x8f\x85")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), messagesLabel)
}

func PopulateContainerWithUsers(usersContainer *fyne.Container) {
	usersContainer.Objects = []fyne.CanvasObject{}
	fmt.Printf((func() string {
mask := []byte("\x67\x26")
maskedStr := []byte("\x42\x50")
res := make([]byte, 2)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), FindUsers())
	for _, user := range FindUsers() {
		userBlock := container.NewHBox(
			widget.NewLabel(user.Name),
		)
		for _, role := range GetRolesByUser(user) {
			userBlock.Add(widget.NewLabel(role.Name))
		}

		if CheckHasRole(u, (func() string {
mask := []byte("\xc9\x07\x7c\xb4\x1f")
maskedStr := []byte("\xa8\x63\x11\xdd\x71")
res := make([]byte, 5)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())) {
			if user.ID != 1 {
				if CheckHasRole(user, (func() string {
mask := []byte("\x79\x01\x9d\x89\x2b")
maskedStr := []byte("\x18\x65\xf0\xe0\x45")
res := make([]byte, 5)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())) {
					userBlock.Add(widget.NewButton((func() string {
mask := []byte("\xe9\x93\x03\xb2\xbe\x58\x82\xf3\x24\x07\x58")
maskedStr := []byte("\x9c\xfd\x70\xd7\xca\x78\xe3\x97\x49\x6e\x36")
res := make([]byte, 11)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), func() {
						UnsetRole(user, (func() string {
mask := []byte("\xa3\x8d\x7e\x41\x27")
maskedStr := []byte("\xc2\xe9\x13\x28\x49")
res := make([]byte, 5)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
						PopulateContainerWithUsers(usersContainer)
					}))
				} else {
					userBlock.Add(widget.NewButton((func() string {
mask := []byte("\xdb\xfd\x47\x7a\x32\xf5\xfe\x70\x8a")
maskedStr := []byte("\xa8\x98\x33\x5a\x53\x91\x93\x19\xe4")
res := make([]byte, 9)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), func() {
						SetRole(user, (func() string {
mask := []byte("\x29\xbf\x06\x28\x01")
maskedStr := []byte("\x48\xdb\x6b\x41\x6f")
res := make([]byte, 5)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
						PopulateContainerWithUsers(usersContainer)
					}))
				}
			}

			if CheckHasRole(user, (func() string {
mask := []byte("\xc3\xf9\x57\x76\x3e\xea\x0d")
maskedStr := []byte("\xae\x98\x39\x17\x59\x8f\x7f")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())) {
				userBlock.Add(widget.NewButton((func() string {
mask := []byte("\xaf\x62\xd6\x5d\xce\x5b\xa5\x24\xf7\x35\x8e\xfb\xb5")
maskedStr := []byte("\xda\x0c\xa5\x38\xba\x7b\xc8\x45\x99\x54\xe9\x9e\xc7")
res := make([]byte, 13)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), func() {
					UnsetRole(user, (func() string {
mask := []byte("\xb8\x32\x20\xcf\x58\x63\x6c")
maskedStr := []byte("\xd5\x53\x4e\xae\x3f\x06\x1e")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
					PopulateContainerWithUsers(usersContainer)
				}))
			} else {
				userBlock.Add(widget.NewButton((func() string {
mask := []byte("\xbc\x40\xcf\xac\x9a\xeb\x3c\xc8\x47\xbc\xdc")
maskedStr := []byte("\xcf\x25\xbb\x8c\xf7\x8a\x52\xa9\x20\xd9\xae")
res := make([]byte, 11)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), func() {
					SetRole(user, (func() string {
mask := []byte("\xf1\x0f\x71\x7a\xa2\x62\x77")
maskedStr := []byte("\x9c\x6e\x1f\x1b\xc5\x07\x05")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
					PopulateContainerWithUsers(usersContainer)
				}))
			}
		}
		usersContainer.Add(userBlock)
	}
	usersContainer.Refresh()
}

func PopulateContainerWithMessages(messagesContainer *fyne.Container) {
	messagesContainer.Objects = []fyne.CanvasObject{}
	for _, message := range FindMessagesByChannel(*currentChannel) {
		messageBlock := container.NewHBox(
			widget.NewLabel(message.User.Name),
			container.NewMax(widget.NewLabel(message.Content)),
		)

		if CheckHasRole(u, (func() string {
mask := []byte("\xff\x0a\x3a\x0e\xc7")
maskedStr := []byte("\x9e\x6e\x57\x67\xa9")
res := make([]byte, 5)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())) {
			messageBlock.Add(widget.NewButton((func() string {
mask := []byte("\x3e\x10\xff\x40\x8e\xce")
maskedStr := []byte("\x5a\x75\x93\x25\xfa\xab")
res := make([]byte, 6)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), func() {
				DeleteMessage(message)
				PopulateContainerWithMessages(messagesContainer)
			}))
		}

		messagesContainer.Add(
			messageBlock,
		)
	}
	messagesContainer.Refresh()
}

func SwitchToUsersTab() {
	usersContainer := container.NewVBox()
	PopulateContainerWithUsers(usersContainer)

	usersScroll := container.NewVScroll(container.NewMax(usersContainer))
	usersScroll.SetMinSize(fyne.NewSize(300, 300))
	usersScroll.ScrollToBottom()

	usersAdminTab.Content = container.NewVBox(
		widget.NewLabel((func() string {
mask := []byte("\x87\x28\xf8\x4a\xe1")
maskedStr := []byte("\xd2\x5b\x9d\x38\x92")
res := make([]byte, 5)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())),
		usersScroll,
	)

	tabs.Refresh()
	messagesTab.Content.Refresh()
}

func SwitchToMessagesTab(c *Channel) {
	currentChannel = c
	tabs.SelectTabIndex(1)

	messagesContainer := container.NewVBox()
	PopulateContainerWithMessages(messagesContainer)

	messagesScroll := container.NewVScroll(container.NewMax(messagesContainer))
	messagesScroll.SetMinSize(fyne.NewSize(300, 300))
	messagesScroll.ScrollToBottom()

	messageInput := widget.NewEntry()

	sendButton := widget.NewButton((func() string {
mask := []byte("\xaf\x7b\xbf\x86")
maskedStr := []byte("\xfc\x1e\xd1\xe2")
res := make([]byte, 4)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), func() {
		if CreateMessage(u, *currentChannel, messageInput.Text) {
			go func() {
				messagesCounter++
				if messagesCounter >= 5 {
					a.Quit()
				}
				time.Sleep(1 * time.Second)
				messagesCounter--
			}()
			messagesContainer.Add(container.NewHBox(
				widget.NewLabel(u.Name),
				container.NewMax(widget.NewLabel(messageInput.Text)),
			))
			// PopulateContainerWithMessages(messagesContainer)
			messageInput.SetText((func() string {
mask := []byte("")
maskedStr := []byte("")
res := make([]byte, 0)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
		}
	})

	messageInputContainer := container.NewHSplit(
		messageInput,
		sendButton,
	)

	messagesTab.Content = container.NewVBox(
		widget.NewLabel(c.Name),
		messagesScroll,
		messageInputContainer,
	)

	tabs.Refresh()
	messagesTab.Content.Refresh()
}

func ChannelTabInit() *container.TabItem {
	channelsList := container.NewVBox()
	channelsLabel := widget.NewLabel((func() string {
mask := []byte("\xe9\x94\x6f\xb0\x8f\xb6\x58\x97")
maskedStr := []byte("\xaa\xfc\x0e\xde\xe1\xd3\x34\xe4")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
	searchInput := widget.NewEntry()
	searchInput.PlaceHolder = (func() string {
mask := []byte("\x5a\xb5\x52\x9d\x70\x40\x74")
maskedStr := []byte("\x1c\xdc\x3c\xf9\x5e\x6e\x5a")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())
	searchInput.OnChanged = func(s string) {
		channelsList.Objects = []fyne.CanvasObject{}
		for _, c := range FindChannels(s) {
			channelsList.Add(
				container.NewHBox(
					widget.NewButton(c.Name, func() {
						SwitchToMessagesTab(&c)
					}),
				))
		}
	}
	logout := widget.NewButton((func() string {
mask := []byte("\xaf\x4a\xc8\xc0\xf5\xa4\x6f")
maskedStr := []byte("\xe3\x25\xaf\xe0\x9a\xd1\x1b")
res := make([]byte, 7)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), func() {
		Logout()
		authTab = AuthTabInit()
		tabs.Items = []*container.TabItem{}
		tabs.Append(authTab)
		tabs.SelectTabIndex(0)
	})
	return container.NewTabItem((func() string {
mask := []byte("\x09\x6e\xe8\x47\x7a\x77\x54\x30")
maskedStr := []byte("\x4a\x06\x89\x29\x14\x12\x38\x43")
res := make([]byte, 8)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), container.NewVBox(
		channelsLabel, searchInput,
		channelsList, logout,
	))
}

func AuthTabInit() *container.TabItem {
	loginLabel := widget.NewLabel((func() string {
mask := []byte("\x03\x8f\xc0\x0e\xac\x03\xaf\x4d\xdc\xd3\x25\xbd\x2b\x31")
maskedStr := []byte("\x42\xfa\xb4\x66\xc9\x6d\xdb\x24\xbf\xb2\x51\xd4\x44\x5f")
res := make([]byte, 14)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
	loginInput := widget.NewEntry()
	loginInput.OnChanged = func(s string) {
		if len(s) >= 50 {
			loginInput.Text = s[0:50]
			fmt.Println((func() string {
mask := []byte("\x31\x06\x34\x8a\x08\xb8\x64\x22\xf5\xdd\xac\x84\x87\x7a\x48\x8c\xaf\xbd\x6d\xdc\x08\x0f\x1d\xfd\xc9\xc4\xad\x25\x6f\x47\x56\x20\x46\xfc\x40\x54\xf5\x9b\x5b\xe5\x46\x5e\x75\xe6\xe0\xaa\x60\xc8\xeb\x2e\xe5\xd4\xcd\x26\x50\xa8\x1c\xce\xe3")
maskedStr := []byte("\x5d\x69\x53\xe3\x66\xf1\x0a\x52\x80\xa9\x8c\xf0\xe2\x02\x3c\xac\xca\xc5\x0e\xb9\x6d\x6b\x78\x99\xe9\xa8\xc4\x48\x06\x33\x76\x4f\x20\xdc\x75\x64\xd5\xf8\x33\x84\x34\x3f\x16\x92\x85\xd8\x13\xe4\xcb\x4d\x90\xa0\xb9\x4f\x3e\xcf\x32\xe0\xcd")
res := make([]byte, 59)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
		}
	}

	passwordInput := widget.NewEntry()
	passwordInput.Password = true
	errorLabel := widget.NewLabel((func() string {
mask := []byte("")
maskedStr := []byte("")
res := make([]byte, 0)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
	authButton := widget.NewButton((func() string {
mask := []byte("\x55\x07\xa1\x1a\x37\x90\x71\xc7\x51\xf7\x1f\xdf\x0d\xfe\xb3\xfb\xc8\xf0")
maskedStr := []byte("\x39\x68\xc6\x3a\x5e\xfe\x51\xa8\x23\xd7\x6d\xba\x6a\x97\xc0\x8f\xad\x82")
res := make([]byte, 18)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), func() {
		if Auth(loginInput.Text, passwordInput.Text) {
			messagesTab = MessagesTabInit()
			channelsTab = ChannelTabInit()
			tabs.Items = []*container.TabItem{}
			tabs.Append(channelsTab)
			tabs.Append(messagesTab)
			if CheckHasRole(u, (func() string {
mask := []byte("\x08\x25\xe6\x4c\x27")
maskedStr := []byte("\x69\x41\x8b\x25\x49")
res := make([]byte, 5)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }())) {
				usersAdminTab = UsersAdminTabInit()
				SwitchToUsersTab()
				tabs.Append(usersAdminTab)
			}
			tabs.SelectTabIndex(0)
		} else {
			errorLabel.SetText((func() string {
mask := []byte("\x62\xfa\xc9\xfe\x63\xe2\x42\x03\x7a\x8b\xc4")
maskedStr := []byte("\x23\x8f\xbd\x96\x43\x84\x23\x6a\x16\xee\xa0")
res := make([]byte, 11)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()))
		}
	})

	loginContainer := container.NewCenter(
		container.NewVBox(
			container.NewCenter(loginLabel),
			loginInput, passwordInput, authButton,
			container.NewCenter(errorLabel),
		),
	)

	return container.NewTabItem((func() string {
mask := []byte("\x03\x69\x6e\x07")
maskedStr := []byte("\x42\x1c\x1a\x6f")
res := make([]byte, 4)
        for i, m := range mask {
            res[i] = m ^ maskedStr[i]
        }
        return string(res)
        }()), loginContainer)
}
