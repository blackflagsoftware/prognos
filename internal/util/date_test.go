package util

import (
	"reflect"
	"testing"
	"time"
)

func TestGetLastMonth(t *testing.T) {
	type args struct {
		now time.Time
	}
	tests := []struct {
		name          string
		args          args
		wantStartDate time.Time
		wantEndDate   time.Time
	}{
		{
			"08-05-2022",
			args{time.Date(2022, time.August, 5, 0, 0, 0, 0, time.UTC)},
			time.Date(2022, time.July, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2022, time.July, 31, 0, 0, 0, 0, time.UTC),
		},
		{
			"10-26-2022",
			args{time.Date(2022, time.October, 26, 0, 0, 0, 0, time.UTC)},
			time.Date(2022, time.September, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2022, time.September, 30, 0, 0, 0, 0, time.UTC),
		},
		{
			"01-13-2022",
			args{time.Date(2022, time.January, 13, 0, 0, 0, 0, time.UTC)},
			time.Date(2021, time.December, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2021, time.December, 31, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStartDate, gotEndDate := GetLastMonth(tt.args.now)
			if !reflect.DeepEqual(gotStartDate, tt.wantStartDate) {
				t.Errorf("GetLastMonth() gotStartDate = %v, want %v", gotStartDate, tt.wantStartDate)
			}
			if !reflect.DeepEqual(gotEndDate, tt.wantEndDate) {
				t.Errorf("GetLastMonth() gotEndDate = %v, want %v", gotEndDate, tt.wantEndDate)
			}
		})
	}
}
