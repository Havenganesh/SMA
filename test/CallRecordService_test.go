package test

import (
	"context"
	"sma/db"
	"sma/dto"
	"sma/service"
	"sma/validate"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateCallRecord(t *testing.T) {
	validate.Init()
	db.Init()
	// Hardcoded test cases with unique names
	testCases := []struct {
		name      string
		ctx       context.Context
		request   *dto.CreateCallRecordRequest
		expectErr bool
	}{
		{
			name: "Valid Request - User 1",
			ctx:  context.WithValue(context.Background(), "serviceType", "NORMAL"),
			request: &dto.CreateCallRecordRequest{
				Caller: &dto.Contact{
					Name:        "Alice Johnson",
					PhoneNumber: "+1234567890",
					Prefix:      "+1",
				},
				Receiver: &dto.Contact{
					Name:        "Bob Williams",
					PhoneNumber: "+1987654321",
					Prefix:      "+1",
				},
				Duration:      300, // 5 minutes
				CallType:      "Voice",
				StartTime:     timePtr(time.Now()),
				EndTime:       timePtr(time.Now().Add(5 * time.Minute)),
				Quality:       "Good",
				NetworkIssues: true,
			},
			expectErr: false,
		},
		{
			name: "InValid Request - User 2",
			ctx:  context.WithValue(context.Background(), "serviceType", "ELITE"),
			request: &dto.CreateCallRecordRequest{
				Caller: nil,
				Receiver: &dto.Contact{
					Name:        "Daisy Miller",
					PhoneNumber: "+1777888999",
					Prefix:      "+1",
				},
				Duration:      450, // 7.5 minutes
				CallType:      "Video",
				StartTime:     timePtr(time.Now()),
				EndTime:       timePtr(time.Now().Add(7*time.Minute + 30*time.Second)),
				Quality:       "Excellent",
				NetworkIssues: true,
			},
			expectErr: true,
		},
		{
			name: "InValid Request - User 2",
			ctx:  context.WithValue(context.Background(), "serviceType", "PREMIUM"),
			request: &dto.CreateCallRecordRequest{
				Caller: &dto.Contact{
					Name:        "Charlie Brown",
					PhoneNumber: "+1444555666",
					Prefix:      "+1",
				},
				Receiver:      nil,
				Duration:      450, // 7.5 minutes
				CallType:      "Video",
				StartTime:     timePtr(time.Now()),
				EndTime:       timePtr(time.Now().Add(7*time.Minute + 30*time.Second)),
				Quality:       "Excellent",
				NetworkIssues: true,
			},
			expectErr: true,
		},
		{
			name: "Valid Request - User 3",
			ctx:  context.WithValue(context.Background(), "serviceType", "NORMAL"),
			request: &dto.CreateCallRecordRequest{
				Caller: &dto.Contact{
					Name:        "Ethan Scott",
					PhoneNumber: "+1222333444",
					Prefix:      "+1",
				},
				Receiver: &dto.Contact{
					Name:        "Fiona Green",
					PhoneNumber: "+1666777888",
					Prefix:      "+1",
				},
				Duration:      0,
				CallType:      "Voice",
				StartTime:     timePtr(time.Now()),
				EndTime:       timePtr(time.Now()),
				Quality:       "Fair",
				NetworkIssues: false,
			},
			expectErr: false,
		},
	}

	// Create an instance of the service
	callRecordService := &service.CallRecordService{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := callRecordService.CreateCallRecord(tc.ctx, tc.request)

			// Check for expected error cases
			if tc.expectErr {
				assert.Error(t, err, "Expected an error but got nil")
			} else {
				assert.NoError(t, err, "Expected no error")
				assert.NotNil(t, resp, "Response should not be nil")
			}
		})
	}
}

func TestListCallRecord(t *testing.T) {
	validate.Init()
	db.Init()
	// Hardcoded test cases with unique names
	testCases := []struct {
		name      string
		ctx       context.Context
		request   *dto.ListCallRecordRequest
		expectErr bool
	}{
		{
			name:      "Valid Request - User 1",
			ctx:       context.WithValue(context.Background(), "userName", "test_user1"),
			request:   &dto.ListCallRecordRequest{},
			expectErr: false,
		},
		{
			name:      "Valid Request - User 2",
			ctx:       context.WithValue(context.Background(), "userName", "test_user2"),
			request:   &dto.ListCallRecordRequest{PhoneNumber: "+1234567890"},
			expectErr: false,
		},
		{
			name:      "Valid Request - User 2",
			ctx:       context.WithValue(context.Background(), "userName", "test_user2"),
			request:   &dto.ListCallRecordRequest{CallerName: "Alice Johnson"},
			expectErr: false,
		},
		{
			name:      "Valid Request - User 2",
			ctx:       context.WithValue(context.Background(), "userName", "test_user2"),
			request:   &dto.ListCallRecordRequest{ReceiverName: "Bob Williams"},
			expectErr: false,
		},
		{
			name: "Ivalid Valid Request - Use wrong times 1",
			ctx:  context.WithValue(context.Background(), "userName", "test_user1"),
			request: &dto.ListCallRecordRequest{
				StartTime: time.Now().Add(time.Hour * 1),
				EndTime:   time.Now()},
			expectErr: true,
		},
	}

	// Create an instance of the service
	callRecordService := &service.CallRecordService{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := callRecordService.ListCallRecord(tc.ctx, tc.request)

			// Check for expected error cases
			if tc.expectErr {
				assert.Error(t, err, "Expected an error but got nil")
			} else {
				assert.NoError(t, err, "Expected no error")
				assert.NotNil(t, resp, "Response should not be nil")
			}
		})
	}
}

func TestGetCallSummary(t *testing.T) {
	validate.Init()
	db.Init()
	// Hardcoded test cases with unique names
	testCases := []struct {
		name      string
		ctx       context.Context
		request   *dto.GetCallSummaryRequest
		expectErr bool
	}{
		{
			name:      "Valid Request - User 1",
			ctx:       context.WithValue(context.Background(), "userName", "test_user1"),
			request:   &dto.GetCallSummaryRequest{PhoneNumber: "+1222333444"},
			expectErr: false,
		},
		{
			name:      "Valid Request - User 1",
			ctx:       context.WithValue(context.Background(), "userName", "test_user1"),
			request:   &dto.GetCallSummaryRequest{PhoneNumber: "+1234567890"},
			expectErr: false,
		},
		{
			name: "Ivalid Valid Request - Use wrong times 1",
			ctx:  context.WithValue(context.Background(), "userName", "test_user1"),
			request: &dto.GetCallSummaryRequest{PhoneNumber: "+1234567890",
				StartTime: time.Now().Add(time.Hour * 1),
				EndTime:   time.Now()},
			expectErr: true,
		},
		{
			name:      "Ivalid Valid Request - no phone Number",
			ctx:       context.WithValue(context.Background(), "userName", "test_user1"),
			request:   &dto.GetCallSummaryRequest{},
			expectErr: true,
		},
	}

	// Create an instance of the service
	callRecordService := &service.CallRecordService{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := callRecordService.GetCallSummary(tc.ctx, tc.request)

			// Check for expected error cases
			if tc.expectErr {
				assert.Error(t, err, "Expected an error but got nil")
			} else {
				assert.NoError(t, err, "Expected no error")
				assert.NotNil(t, resp, "Response should not be nil")
			}
		})
	}
}
func TestQualityReport(t *testing.T) {
	validate.Init()
	db.Init()
	// Hardcoded test cases with unique names
	testCases := []struct {
		name      string
		ctx       context.Context
		request   *dto.GetCallQualityReportReq
		expectErr bool
	}{
		{
			name:      "Valid Request - User 1",
			ctx:       context.WithValue(context.Background(), "userName", "test_user1"),
			request:   &dto.GetCallQualityReportReq{PhoneNumber: "+1222333444"},
			expectErr: false,
		},
		{
			name:      "Valid Request - User 2",
			ctx:       context.WithValue(context.Background(), "userName", "test_user1"),
			request:   &dto.GetCallQualityReportReq{PhoneNumber: "+1234567890"},
			expectErr: false,
		},
		{
			name: "Ivalid Valid Request - Use wrong times 1",
			ctx:  context.WithValue(context.Background(), "userName", "test_user1"),
			request: &dto.GetCallQualityReportReq{PhoneNumber: "+1234567890",
				StartTime: time.Now().Add(time.Hour * 1),
				EndTime:   time.Now()},
			expectErr: true,
		},
		{
			name:      "Ivalid Valid Request - no phone Number",
			ctx:       context.WithValue(context.Background(), "userName", "test_user1"),
			request:   &dto.GetCallQualityReportReq{},
			expectErr: true,
		},
	}

	// Create an instance of the service
	callRecordService := &service.CallRecordService{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := callRecordService.GetCallQualityReport(tc.ctx, tc.request)

			// Check for expected error cases
			if tc.expectErr {
				assert.Error(t, err, "Expected an error but got nil")
			} else {
				assert.NoError(t, err, "Expected no error")
				assert.NotNil(t, resp, "Response should not be nil")
			}
		})
	}
}

func TestGetBillReport(t *testing.T) {
	validate.Init()
	db.Init()
	// Hardcoded test cases with unique names
	testCases := []struct {
		name      string
		ctx       context.Context
		request   *dto.GetCallBillReportReq
		expectErr bool
	}{
		{
			name:      "Valid Request - User 1",
			ctx:       context.WithValue(context.Background(), "userName", "test_user1"),
			request:   &dto.GetCallBillReportReq{PhoneNumber: "+1222333444"},
			expectErr: false,
		},
		{
			name: "Ivalid Valid Request - Use wrong times 1",
			ctx:  context.WithValue(context.Background(), "userName", "test_user1"),
			request: &dto.GetCallBillReportReq{PhoneNumber: "+1234567890",
				StartTime: time.Now().Add(time.Hour * 1),
				EndTime:   time.Now()},
			expectErr: true,
		},
		{
			name:      "Ivalid Valid Request - no phone Number",
			ctx:       context.WithValue(context.Background(), "userName", "test_user1"),
			request:   &dto.GetCallBillReportReq{},
			expectErr: true,
		},
	}

	// Create an instance of the service
	callRecordService := &service.CallRecordService{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := callRecordService.GetCallBillReport(tc.ctx, tc.request)

			// Check for expected error cases
			if tc.expectErr {
				assert.Error(t, err, "Expected an error but got nil")
			} else {
				assert.NoError(t, err, "Expected no error")
				assert.NotNil(t, resp, "Response should not be nil")
			}
		})
	}
}

// Helper function to return a pointer to a time.Time value
func timePtr(t time.Time) *time.Time {
	return &t
}
