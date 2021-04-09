package main

func Auth(login, password string) bool {
	db.First(&u, "name = ? AND password_hash = ?", login, password)
	return u.ID != 0
}

func Logout() {
	u = User{}
}

func FindChannels(s string) []Channel {
	channels := []Channel{}
	db.Find(&channels, "name like ?", "%"+s+"%")
	return channels
}

func FindUsers() []User {
	users := []User{}
	db.Find(&users, "name != ?", u.Name)
	return users
}

func FindMessagesByChannel(c Channel) []Message {
	messages := []Message{}
	db.Preload("User").Find(&messages, "channel_id = ?", c.ID)
	return messages
}

func CreateMessage(u User, c Channel, content string) bool {
	db.Create(&Message{
		Content: content, User: u, Channel: c,
		UserId: u.ID, ChannelId: c.ID,
	})
	return true
}

func GetRolesByUser(u User) []Role {
	rs := []Role{}
	db.Find(&rs, "user_id = ?", u.ID)
	return rs
}

func CheckHasRole(u User, role string) bool {
	rs := GetRolesByUser(u)
	for _, r := range rs {
		if r.Name == role {
			return true
		}
	}
	return false
}

func DeleteMessage(message Message) bool {
	db.Delete(&message)
	return true
}

func SetRole(u User, role string) bool {
	db.Create(&Role{User: u, Name: role})
	return true
}

func UnsetRole(u User, role string) bool {
	r := Role{}
	db.Delete(&r, "user_id = ? AND name = ?", u.ID, role)
	return true
}
