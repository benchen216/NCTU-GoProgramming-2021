package cw

type Article struct {
	Article_id    string
	Article_title string
	Author        string
	Date          string
}

type PTTArticles struct {
	Articles []PTTArticle
}

type PTTArticle struct {
	Article
	Message_count PTTMessageCount
	Ip            string
	Url           string
}

type PTTMessageCount struct {
	Push    int
	Neutral int
	Boo     int
}

type FBArticles struct {
	Articles []FBArticle
}

type FBArticle struct {
	Article
	Message_count FBMessageCount
}

type FBMessageCount struct {
	Like    int
	Dislike int
}
