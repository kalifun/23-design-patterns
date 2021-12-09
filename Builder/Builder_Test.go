package builder

import (
	"reflect"
	"testing"
)

func TestWorker_buildMotorVehicles(t *testing.T) {
	tests := []struct {
		name string
		w    *Worker
		want MotorVehicles
	}{
		{
			name: "test",
			w: &Worker{
				b: newCar(),
			},
			want: MotorVehicles{
				Tires:    4,
				Cylinder: 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.w.buildMotorVehicles(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Worker.buildMotorVehicles() = %v, want %v", got, tt.want)
			}
		})
	}
}
