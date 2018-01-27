package facade

var db = map[string]string{
	"hoge@example.com": "hoge",
	"fuga@example.com": "fuga",
	"piyo@example.com": "piyo",
}

type Database struct {
}

func (self *Database) GetUserNameByEmail(email string) string {
	return db[email]
}

type MarkdownWriter struct {
}

func (self *MarkdownWriter) title(userName string) string {
	return "# Welcome to " + userName + "`s page!!"
}

type PageMaker struct {
}

func (self *PageMaker) MakeWelcomePage(email string) string {
	db := Database{}
	writer := MarkdownWriter{}

	name := db.GetUserNameByEmail(email)
	page := writer.title(name)

	return page
}
