package nameday

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/avu12/golangwebpage/webpagego/internal/services"
)

func GetNamedayNow() map[string]interface{} {
	request, err := http.NewRequest("GET", "https://api.abalin.net/today", nil)
	if err != nil {
		fmt.Println("ERROR in request!")
	}

	result, err := services.NamedayService.GetNameday(request)
	if err != nil {
		fmt.Println("ERROR in result!")
	}
	fmt.Println(result)

	var helperslice []string
	helperslice = strings.Split(result.Data.Namedays.Hungarian, ",")

	datas := map[string]interface{}{
		"day":      result.Data.Dates.Day,
		"month":    result.Data.Dates.Month,
		"namedays": helperslice,
	}
	return datas
}
