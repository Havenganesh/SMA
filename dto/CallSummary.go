package dto

type CallSummary struct {
	TotalCalls      int     `json:"totalCalls,omitempty" bson:"totalCalls,omitempty"`
	TotalDuration   int     `json:"totalDuration,omitempty" bson:"totalDuration,omitempty"`
	AverageDuration float64 `json:"averageDuration,omitempty" bson:"averageDuration,omitempty"`
	VoiceCalls      int     `json:"voiceCalls,omitempty" bson:"voiceCalls,omitempty"`
	VideoCalls      int     `json:"videoCalls,omitempty" bson:"videoCalls,omitempty"`
	NetworkIssues   int     `json:"networkIssues,omitempty" bson:"networkIssues,omitempty"`
	TotalCallCost   float64 `json:"totalCallCost,omitempty" bson:"totalCallCost,omitempty"`
}
