package telegram

const msgHelp = `I can save and keep your pages. Also I can offer to you them to read.

In order to save the page, just send me all links to it.

In order to get a random page from your list, send me command /rnd
Caution! After that, this page will be removes from your list!`

const msgHello = "Hi there! \n\n" + msgHelp

const (
	msgUnknownCommand = "Unknown command"
	magNoSavedPages   = "You have no saved pages"
	msgSaved          = "Saved!"
	msgAlreadyExists  = "You have already have this page in your list"
)
