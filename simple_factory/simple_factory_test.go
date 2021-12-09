package simplefactory

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_traffic_SetTraffic(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		tr   *traffic
		args args
	}{
		{
			name: "test1",
			tr: &traffic{
				Name:  "宝马",
				Parts: []Part{},
			},
			args: args{
				name: "宝马",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tr.SetTraffic(tt.args.name)
		})
	}
}

func Test_traffic_SetPart(t *testing.T) {
	type args struct {
		p Part
	}
	tests := []struct {
		name string
		tr   *traffic
		args args
	}{
		{
			name: "test1",
			tr: &traffic{
				Name: "",
				Parts: []Part{
					{
						Name: "Tires",
						Num:  2,
					},
				},
			},
			args: args{
				p: Part{
					Name: "Tires",
					Num:  2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tr.SetPart(tt.args.p)
		})
	}
}

func TestNewCar(t *testing.T) {
	bmw := NewCar("宝马")
	bmw.SetPart(Part{
		Name: "轮胎",
		Num:  4,
	})
	bmw.SetPart(Part{
		Name: "刹车盘",
		Num:  4,
	})
	fmt.Println(bmw)
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want TheTrafficTools
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCar(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBicycle(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want TheTrafficTools
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBicycle(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBicycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
