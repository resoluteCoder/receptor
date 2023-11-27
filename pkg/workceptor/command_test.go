package workceptor_test

import (
	"context"
	"testing"

	"github.com/ansible/receptor/pkg/workceptor"
	"github.com/ansible/receptor/pkg/workceptor/mock_workceptor"
	"github.com/golang/mock/gomock"
)

func createTestWorkUnit(t *testing.T) workceptor.WorkUnit {

	cwc := &workceptor.CommandWorkerCfg{}
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockNetceptor := mock_workceptor.NewMockNetceptorForWorkceptor(ctrl)
	mockNetceptor.EXPECT().NodeID().Return("NodeID")
	w, err := workceptor.New(ctx, mockNetceptor, "/tmp")
	if err != nil {
		t.Errorf("Error while creating Workceptor: %v", err)
	}
	return cwc.NewWorker(w, "", "")
}

func TestSetFromParams2(t *testing.T) {
	wu := createTestWorkUnit(t)

	paramsTestCases := []struct {
		name       string
		params     map[string]string
		errorCatch func(error, *testing.T)
	}{
		{
			name:   "one",
			params: map[string]string{"": ""},
			errorCatch: func(err error, t *testing.T) {
				if err != nil {
					t.Error(err)
				}
			},
		},
		{
			name:   "two",
			params: map[string]string{"params": "param"},
			errorCatch: func(err error, t *testing.T) {
				if err == nil {
					t.Error(err)
				}
			},
		},
	}

	for _, testCase := range paramsTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			err := wu.SetFromParams(testCase.params)
			testCase.errorCatch(err, t)
		})
	}

}

func TestUnredactedStatus(t *testing.T) {
	wu := createTestWorkUnit(t)
	wu.UnredactedStatus()
}

func TestStart(t *testing.T) {
	wu := createTestWorkUnit(t)
	err := wu.Start()
	if err != nil {
		t.Error(err)
	}
}
