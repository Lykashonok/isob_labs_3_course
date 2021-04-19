package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	w    = a.NewWindow("Isob 4-5 labs")
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
	return container.NewTabItem("Users", container.NewVBox(
		widget.NewLabel("Users"),
		usersScroll,
	))
}

func MessagesTabInit() *container.TabItem {
	messagesLabel := widget.NewLabel("Please, choose channel on Channels tab")
	return container.NewTabItem("Messages", messagesLabel)
}

func PopulateContainerWithUsers(usersContainer *fyne.Container) {
	usersContainer.Objects = []fyne.CanvasObject{}
	fmt.Printf("%v", FindUsers())
	for _, user := range FindUsers() {
		userBlock := container.NewHBox(
			widget.NewLabel(user.Name),
		)
		for _, role := range GetRolesByUser(user) {
			userBlock.Add(widget.NewLabel(role.Name))
		}

		if CheckHasRole(u, "admin") {
			if user.ID != 1 {
				if CheckHasRole(user, "admin") {
					userBlock.Add(widget.NewButton("unset admin", func() {
						UnsetRole(user, "admin")
						PopulateContainerWithUsers(usersContainer)
					}))
				} else {
					userBlock.Add(widget.NewButton("set admin", func() {
						SetRole(user, "admin")
						PopulateContainerWithUsers(usersContainer)
					}))
				}
			}

			if CheckHasRole(user, "manager") {
				userBlock.Add(widget.NewButton("unset manager", func() {
					UnsetRole(user, "manager")
					PopulateContainerWithUsers(usersContainer)
				}))
			} else {
				userBlock.Add(widget.NewButton("set manager", func() {
					SetRole(user, "manager")
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

		if CheckHasRole(u, "admin") {
			messageBlock.Add(widget.NewButton("delete", func() {
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
		widget.NewLabel("Users"),
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

	sendButton := widget.NewButton("Send", func() {
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
			messageInput.SetText("")
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
	channelsLabel := widget.NewLabel("Channels")
	searchInput := widget.NewEntry()
	searchInput.PlaceHolder = "Find..."
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
	logout := widget.NewButton("Log out", func() {
		Logout()
		authTab = AuthTabInit()
		tabs.Items = []*container.TabItem{}
		tabs.Append(authTab)
		tabs.SelectTabIndex(0)
	})
	return container.NewTabItem("Channels", container.NewVBox(
		channelsLabel, searchInput,
		channelsList, logout,
	))
}

func AuthTabInit() *container.TabItem {
	loginLabel := widget.NewLabel("Authentication")
	loginInput := widget.NewEntry()
	loginInput.OnChanged = func(s string) {
		if len(s) >= 50 {
			loginInput.Text = s[0:50]
			fmt.Println("loginInput text exceeded limit of 50 characters, cutting...")
		}
	}

	passwordInput := widget.NewEntry()
	passwordInput.Password = true
	errorLabel := widget.NewLabel("")
	authButton := widget.NewButton("log in or register", func() {
		if Auth(loginInput.Text, passwordInput.Text) {
			messagesTab = MessagesTabInit()
			channelsTab = ChannelTabInit()
			tabs.Items = []*container.TabItem{}
			tabs.Append(channelsTab)
			tabs.Append(messagesTab)
			if CheckHasRole(u, "admin") {
				usersAdminTab = UsersAdminTabInit()
				SwitchToUsersTab()
				tabs.Append(usersAdminTab)
			}
			tabs.SelectTabIndex(0)
		} else {
			errorLabel.SetText("Auth failed")
		}
	})

	loginContainer := container.NewCenter(
		container.NewVBox(
			container.NewCenter(loginLabel),
			loginInput, passwordInput, authButton,
			container.NewCenter(errorLabel),
		),
	)

	return container.NewTabItem("Auth", loginContainer)
}
