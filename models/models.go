package models

//go:generate reform

// News represents a record in the "news" table.
//reform:news
type News struct {
	ID      int64  `reform:"id,pk"`
	Title   string `reform:"title"`
	Content string `reform:"content"`
}

type NewsCategory struct {
	NewsID     int64 `reform:"news_id,pk"`
	CategoryID int64 `reform:"category_id,pk"`
}
