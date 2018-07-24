package lessonType

import "time"

type Lesson struct {
	ID           string    `json:"id" datastore:"-"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	DurationSec  float64   `json:"duration_sec"`
	ViewCount    int64     `json:"view_count"`
	ThumbnailURL string    `json:"thumbnail_url"`
	GraphicIDs   []string  `json:"graphic_ids"`
	Published    time.Time `json:"published"`
	Updated      time.Time `json:"updated"`
}

type LessonAuthor struct {
	ID       string `json:"id" datastore:"-"`
	LessonID string `json:"lesson_id"`
	UserID   string `json:"user_id"`
	Role     string `json:"role"`
}

type Graphic struct {
	ID     string   `json:"id" datastore:"-"`
	UserID string   `json:"user_id"`
	TeamID []string `json:"owner_ids"`
}

/* The following structs is for json.Unmarshall */

type LessonMaterial struct {
  DurationSec  float64            `json:"durationSec"`
	Timelines    []LessonTimeline   `json:"timelines"`
	Poses        []LessonAvatarPose `json:"poses"`
	Published    time.Time          `json:"published"`
	Updated      time.Time          `json:"updated"`
}

type LessonTimeline struct {
	TimeSec  float64                   `json:"timeSec"`
	Text     LessonText                `json:"text"`
	Voice    LessonVoice               `json:"voice"`
	Graphic  []LessonGraphic           `json:"graphics"`
	SPAction LessonAvatarSpecialAction `json:"spAction"`
}

type LessonText struct {
	DurationSec  float64 `json:"durationSec"`
	DelaySec     float64 `json:"delaySec"`
	Text         string  `json:"text"`
	Position     string  `json:"position"`
	Style        string  `json:"style"`
	Size         uint8   `json:"size"`
	Color        string  `json:"color"`
}

type LessonVoice struct {
	FileID      string  `json:"fileID"`
	DurationSec float64 `json:"durationSec"`
	DelaySec    float64 `json:"delaySec"`
}

type LessonGraphic struct {
	DelaySec  float64 `json:"delaySec"`
	GraphicID string  `json:"graphicID"`
	Action    string  `json:"action"`
	WidthPx   uint16  `json:"widthPx"`
	HeightPx  uint16  `json:"heightPx"`
	Position  string  `json:"position"`
}

type LessonAvatarSpecialAction struct {
	Action           string `json:"action"`
	FacialExpression string `json:"facialExpression"`
}

type LessonAvatarPose struct {
	TimeSec    float64                    `json:"timeSec"`
	LeftHand   LessonAvatarPositionVector `json:"leftHand"`
	RightHand  LessonAvatarPositionVector `json:"rightHand"`
	LeftElbow  LessonAvatarPositionVector `json:"leftElbow"`
	RightElbow LessonAvatarPositionVector `json:"rightElbow"`
	LookAt     LessonAvatarPositionVector `json:"lookAt"`
	CoreBody   LessonAvatarPositionVector `json:"coreBody"`
}

type LessonAvatarPositionVector struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}