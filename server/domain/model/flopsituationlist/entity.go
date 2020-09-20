package flopsituationlist

import (
	"database/sql"
)

// Entity はフロップシチュエーションのリストの要素を表す。
type Entity struct {
	InPositionBetFrequency      float32
	OutOfPositionBetFrequency   float32
	InPositionCheckFrequency    float32
	OutOfPositionCheckFrequency float32
	InPosition33BetFrequency    float32
	OutOfPosition33BetFrequency float32
	InPosition67BetFrequency    float32
	OutOfPosition67BetFrequency float32
	InPositionEquity            float32
	OutOfPositionEquity         float32
	ImageURL                    sql.NullString
	ImageName                   sql.NullString
	ImageDescription            sql.NullString
}

// Entities は Entity のリストを表す。
type Entities []*Entity

func (es Entities) avg(values []float32) float32 {
	total := float32(0.0)
	if len(values) == 0 {
		return total
	}
	for _, v := range values {
		total += v
	}
	return total / float32(len(values))
}

// AvgInPositionBetFrequency は InPositionBetFrequency の平均値を返す。
func (es Entities) AvgInPositionBetFrequency() float32 {
	var values []float32
	for _, e := range es {
		values = append(values, e.InPositionBetFrequency)
	}
	return es.avg(values)
}

// AvgOutOfPositionBetFrequency は OutOfPositionBetFrequency の平均値を返す。
func (es Entities) AvgOutOfPositionBetFrequency() float32 {
	var values []float32
	for _, e := range es {
		values = append(values, e.OutOfPositionBetFrequency)
	}
	return es.avg(values)
}

// AvgInPositionCheckFrequency は InPositionCheckFrequency の平均値を返す。
func (es Entities) AvgInPositionCheckFrequency() float32 {
	var values []float32
	for _, e := range es {
		values = append(values, e.InPositionCheckFrequency)
	}
	return es.avg(values)
}

// AvgOutOfPositionCheckFrequency は OutOfPositionCheckFrequency の平均値を返す。
func (es Entities) AvgOutOfPositionCheckFrequency() float32 {
	var values []float32
	for _, e := range es {
		values = append(values, e.OutOfPositionCheckFrequency)
	}
	return es.avg(values)
}

// AvgInPosition33BetFrequency は InPosition33BetFrequency の平均値を返す。
func (es Entities) AvgInPosition33BetFrequency() float32 {
	var values []float32
	for _, e := range es {
		values = append(values, e.InPosition33BetFrequency)
	}
	return es.avg(values)
}

// AvgOutOfPosition33BetFrequency は OutOfPosition33BetFrequency の平均値を返す。
func (es Entities) AvgOutOfPosition33BetFrequency() float32 {
	var values []float32
	for _, e := range es {
		values = append(values, e.OutOfPosition33BetFrequency)
	}
	return es.avg(values)
}

// AvgInPosition67BetFrequency は InPosition67BetFrequency の平均値を返す。
func (es Entities) AvgInPosition67BetFrequency() float32 {
	var values []float32
	for _, e := range es {
		values = append(values, e.InPosition67BetFrequency)
	}
	return es.avg(values)
}

// AvgOutOfPosition67BetFrequency は OutOfPosition67BetFrequency の平均値を返す。
func (es Entities) AvgOutOfPosition67BetFrequency() float32 {
	var values []float32
	for _, e := range es {
		values = append(values, e.OutOfPosition67BetFrequency)
	}
	return es.avg(values)
}

// AvgInPositionEquity は InPositionEquity の平均値を返す。
func (es Entities) AvgInPositionEquity() float32 {
	var values []float32
	for _, e := range es {
		values = append(values, e.InPositionEquity)
	}
	return es.avg(values)
}

// AvgOutOfPositionEquity は OutOfPositionEquity の平均値を返す。
func (es Entities) AvgOutOfPositionEquity() float32 {
	var values []float32
	for _, e := range es {
		values = append(values, e.OutOfPositionEquity)
	}
	return es.avg(values)
}

// Images は画像のリストを返す。
func (es Entities) Images() []Image {
	var images []Image
	for _, e := range es {
		if !e.ImageURL.Valid {
			continue
		}
		images = append(images, Image{
			URL:         e.ImageURL.String,
			Name:        e.ImageName.String,
			Description: e.ImageDescription.String,
		})
	}
	return images
}
