package flopsituationlist

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
	ImageURL                    string
	ImageDescription            string
}
