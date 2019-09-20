package myMysql

type Fruit struct {
	Id         uint   `db:"id"`
	FruitName  string `db:"fruit_name"`
	FruitPrice string `db:"fruit_price"`
	CreateTs   string `db:"create_ts"`
	UpdateTs   string `db:"update_ts"`
}
