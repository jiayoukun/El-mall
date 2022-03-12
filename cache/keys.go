package cache

import (
	"fmt"
)

const(
	TodayKey = "todayrank"
	BookKey = "bookrank"
	Electronic = "electronicrank"
)

func ProductViewKey(id string) string {
	return fmt.Sprintf("view:productid:"+id)
}

