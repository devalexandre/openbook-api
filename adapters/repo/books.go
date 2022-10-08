package repo

type Books struct {
	ID          int    `ksql:"id" json:"id"`
	Title       string `ksql:"title" json:"title"`
	Thumbnail   string `ksql:"thumbnail" json:"thumbnail"`
	Author      string `ksql:"author" json:"author"`
	Description string `ksql:"description" json:"description"`
	Year        int    `ksql:"year" json:"year"`
	Category    string `ksql:"category" json:"category"`
}
