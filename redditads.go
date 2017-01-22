package main

import (
	"github.com/jzelinskie/geddit"
)

func main() {
	// Login to reddit
	session, err := geddit.NewLoginSession(
		"username",
		"password",
		"Test",
	)

	if err != nil {
		panic(err)
	}

	post := geddit.NewSubmission{
		Content:"https://github.com/",
		Subreddit:"test",
		Title:"Github Test",
		Self:false,
		Save:true,
		SendReplies:true,
		Captcha: nil,
	}

	err = session.Submit(&post)

	if err != nil {
		panic(err)
	}
}
