package utils

import "testing"

func TestToJson(t *testing.T) {
	type args struct {
		o interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test nil",
			args: args{
				o: nil,
			},
			want:    "{}",
			wantErr: false,
		},
		{
			name: "test common",
			args: args{
				o: struct {
					Name string
				}{
					Name: "test",
				},
			},
			want: `{"Name":"test"}`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJson(tt.args.o)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToJson() got = %v, want %v", got, tt.want)
			}
		})
	}
}
