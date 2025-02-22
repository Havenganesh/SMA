package service

import (
	"context"
	"fmt"
	"sma/cErrors"
	"sma/db"
	"sma/dto"
	"sma/validate"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CallRecordService struct{}

func (c *CallRecordService) CreateCallRecord(ctx context.Context, req *dto.CreateCallRecordRequest) (*dto.CreateCallRecordResponse, error) {
	if err := validate.Validate.Struct(req); err != nil {
		return nil, fmt.Errorf("%s : %s", cErrors.INVALID_REQUEST, err)
	}
	sType := ctx.Value("serviceType").(string)
	var record dto.CallRecord
	record.Caller = req.Caller
	record.Receiver = req.Receiver
	record.CallType = req.CallType
	record.Duration = req.Duration
	record.StartTime = req.StartTime
	record.EndTime = req.EndTime
	record.Quality = req.Quality
	record.NetworkIssues = req.NetworkIssues
	record.Cost = calcCost(record.Duration, sType)
	record.SearchFilter = append(record.SearchFilter, "CALLER:"+record.Caller.Name)
	record.SearchFilter = append(record.SearchFilter, "RECEIVER:"+record.Receiver.Name)
	record.SearchFilter = append(record.SearchFilter, "CALLER_NO:"+record.Caller.PhoneNumber)
	record.SearchFilter = append(record.SearchFilter, "RECEIVER_NO:"+record.Receiver.PhoneNumber)
	record.SearchFilter = append(record.SearchFilter, "QUALITY:"+record.Quality)
	record.SearchFilter = append(record.SearchFilter, "CALL_TYPE:"+record.CallType)
	record.SearchFilter = append(record.SearchFilter, "START_TIME:"+record.StartTime.GoString())
	if record.NetworkIssues {
		record.SearchFilter = append(record.SearchFilter, "NET_ISSUES:YES")
	} else {
		record.SearchFilter = append(record.SearchFilter, "NET_ISSUES:NO")
	}
	err := db.DB.InsertOne(record)
	if err != nil {
		return nil, err
	}
	return &dto.CreateCallRecordResponse{Message: "Record Saved Successfully"}, nil
}

func (c *CallRecordService) ListCallRecord(ctx context.Context, req *dto.ListCallRecordRequest) (*dto.ListCallRecordResponse, error) {
	var records []*dto.CallRecord
	var searchFilter []string
	if req.CallerName != "" {
		searchFilter = append(searchFilter, "CALLER:"+req.CallerName)
	}
	if req.ReceiverName != "" {
		searchFilter = append(searchFilter, "RECEIVER:"+req.ReceiverName)
	}
	if req.PhoneNumber != "" {
		searchFilter = append(searchFilter, "CALLER_NO:"+req.PhoneNumber)
		searchFilter = append(searchFilter, "RECEIVER_NO:"+req.PhoneNumber)
	}
	filter := bson.M{}
	if len(searchFilter) > 0 {
		filter["searchFilter"] = bson.M{"$in": searchFilter}
	}
	if !req.StartTime.IsZero() {
		filter["startTime"] = bson.M{"$gte": req.StartTime}
		if !req.EndTime.IsZero() {
			if req.EndTime.Before(req.StartTime) {
				return nil, cErrors.END_TIME_MUST_BE_AFTER_START_TIME
			}
			filter["startTime"] = bson.M{"$gte": req.StartTime, "$lte": req.EndTime}
		}
	}
	if err := db.DB.FindAll(filter, &records); err != nil {
		fmt.Println("List call record find all error : ", err)
		return nil, err
	}
	if len(records) == 0 {
		return &dto.ListCallRecordResponse{Message: "No records found"}, nil
	}
	return &dto.ListCallRecordResponse{Records: records}, nil
}

func (c *CallRecordService) GetCallSummary(ctx context.Context, req *dto.GetCallSummaryRequest) (*dto.GetCallSummaryResponse, error) {
	if err := validate.Validate.Struct(req); err != nil {
		return nil, fmt.Errorf("%s : %s", cErrors.INVALID_REQUEST, err)
	}
	match := bson.D{{Key: "caller.phoneNumber", Value: req.PhoneNumber}}
	if !req.StartTime.IsZero() {
		t := bson.D{
			{Key: "$gte", Value: req.StartTime},
		}
		if !req.EndTime.IsZero() {
			if req.EndTime.Before(req.StartTime) {
				return nil, cErrors.END_TIME_MUST_BE_AFTER_START_TIME
			}
			t = append(t, bson.E{Key: "$lte", Value: req.EndTime})
		}
		match = append(match, bson.E{
			Key: "startTime", Value: t,
		})
	}
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: match}},
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: nil},
			{Key: "totalCalls", Value: bson.D{{Key: "$sum", Value: 1}}},
			{Key: "totalDuration", Value: bson.D{{Key: "$sum", Value: "$duration"}}},
			{Key: "totalCallCost", Value: bson.D{{Key: "$sum", Value: "$cost"}}},
			{Key: "videoCalls", Value: bson.D{{Key: "$sum", Value: bson.D{
				{Key: "$cond", Value: bson.A{
					bson.D{{Key: "$eq", Value: bson.A{"$callType", "Video"}}}, 1, 0,
				}},
			}}}},
			{Key: "voiceCalls", Value: bson.D{{Key: "$sum", Value: bson.D{
				{Key: "$cond", Value: bson.A{
					bson.D{{Key: "$eq", Value: bson.A{"$callType", "Voice"}}}, 1, 0,
				}},
			}}}},
			{Key: "networkIssues", Value: bson.D{{Key: "$sum", Value: bson.D{
				{Key: "$cond", Value: bson.A{
					bson.D{{Key: "$eq", Value: bson.A{"$networkIssues", true}}}, 1, 0,
				}},
			}}}},
		}}}}

	var summary dto.CallSummary
	err := db.DB.Aggregate(pipeline, &dto.CallRecord{}, &summary)
	if err != nil {
		return nil, fmt.Errorf("%s : %s", cErrors.DATABASE_AGGREAGATION_FAILED, err)
	}
	return &dto.GetCallSummaryResponse{CallSummary: &summary}, nil
}

func (c *CallRecordService) GetCallQualityReport(ctx context.Context, req *dto.GetCallQualityReportReq) (*dto.GetCallQualityReportRes, error) {
	if err := validate.Validate.Struct(req); err != nil {
		return nil, fmt.Errorf("%s : %s", cErrors.INVALID_REQUEST, err)
	}
	match := bson.D{{Key: "caller.phoneNumber", Value: req.PhoneNumber}}
	if !req.StartTime.IsZero() {
		t := bson.D{
			{Key: "$gte", Value: req.StartTime},
		}
		if !req.EndTime.IsZero() {
			if req.EndTime.Before(req.StartTime) {
				return nil, cErrors.END_TIME_MUST_BE_AFTER_START_TIME
			}
			t = append(t, bson.E{Key: "$lte", Value: req.EndTime})
		}
		match = append(match, bson.E{
			Key: "startTime", Value: t,
		})
	}
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: match}},
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: nil},
			{Key: "totalCalls", Value: bson.D{{Key: "$sum", Value: 1}}},
			{Key: "excellent", Value: bson.D{{Key: "$sum", Value: bson.D{
				{Key: "$cond", Value: bson.A{
					bson.D{{Key: "$eq", Value: bson.A{"$quality", "Excellent"}}}, 1, 0,
				}},
			}}}},
			{Key: "good", Value: bson.D{{Key: "$sum", Value: bson.D{
				{Key: "$cond", Value: bson.A{
					bson.D{{Key: "$eq", Value: bson.A{"$quality", "Good"}}}, 1, 0,
				}},
			}}}},
			{Key: "fair", Value: bson.D{{Key: "$sum", Value: bson.D{
				{Key: "$cond", Value: bson.A{
					bson.D{{Key: "$eq", Value: bson.A{"$quality", "Fair"}}}, 1, 0,
				}},
			}}}},
		}}}}
	var report dto.QualityReport
	err := db.DB.Aggregate(pipeline, &dto.CallRecord{}, &report)
	if err != nil {
		return nil, fmt.Errorf("%s : %s", cErrors.DATABASE_AGGREAGATION_FAILED, err)
	}
	return &dto.GetCallQualityReportRes{QualityReport: &report}, nil
}

func (c *CallRecordService) GetCallBillReport(ctx context.Context, req *dto.GetCallBillReportReq) (*dto.GetCallBillReportRes, error) {
	if err := validate.Validate.Struct(req); err != nil {
		return nil, fmt.Errorf("%s : %s", cErrors.INVALID_REQUEST, err)
	}
	var records []*dto.CallRecord
	var searchFilter []string
	searchFilter = append(searchFilter, "CALLER_NO:"+req.PhoneNumber)
	filter := bson.M{"searchFilter": bson.M{"$in": searchFilter}}
	if !req.StartTime.IsZero() {
		filter["startTime"] = bson.M{"$gte": req.StartTime}
		if !req.EndTime.IsZero() {
			if req.EndTime.Before(req.StartTime) {
				return nil, cErrors.END_TIME_MUST_BE_AFTER_START_TIME
			}
			filter["startTime"] = bson.M{"$gte": req.StartTime, "$lte": req.EndTime}
		}
	}
	if err := db.DB.FindAll(filter, &records); err != nil {
		fmt.Println("List call record find all error : ", err)
		return nil, err
	}
	var billReport []*dto.BillReport
	var totalCost float64
	var totalDuration int64
	for _, rec := range records {
		br := &dto.BillReport{}
		br.CallerName = rec.Caller.Name
		br.CallerNumber = rec.Caller.PhoneNumber
		br.ReceiverName = rec.Receiver.Name
		br.Cost = rec.Cost
		br.Duration = rec.Duration
		billReport = append(billReport, br)
		totalCost += rec.Cost
		totalDuration += int64(rec.Duration)
	}
	return &dto.GetCallBillReportRes{BillReport: billReport, TotalCost: totalCost, TotalDuration: totalDuration}, nil
}

func calcCost(duration int64, sType string) float64 {
	var cost float64
	switch sType {
	case string(dto.NORMAL):
		cost = float64(duration) * 1.5
	case string(dto.ELITE):
		cost = float64(duration) * 2.5
	case string(dto.PREMIUM):
		cost = float64(duration) * 3.5
	}
	return cost
}
