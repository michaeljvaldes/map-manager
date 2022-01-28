package mapprep

import (
	"time"
)

type IndexData struct {
	UpdateTime string
}

func ConstructIndexData(t time.Time) IndexData {
	return IndexData{UpdateTime: t.Format("2006-01-02 15:04:05 MST")}
}
