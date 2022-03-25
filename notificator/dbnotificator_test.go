package notificator

import (
	"reflect"
	"testing"
)

func TestNewDbnotificator(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want *dbnotificator
	}{
		{name: "", args: args{id: "test"}, want:&dbnotificator{id: "test"} },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDbnotificator(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDbnotificator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dbnotificator_GetID(t *testing.T) {
	type fields struct {
		id     string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "", fields: fields{
			id:     "test",
		} , want:"test" },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &dbnotificator{
				id:     tt.fields.id,
			}
			if got := d.GetID(); got != tt.want {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dbnotificator_Notify(t *testing.T) {
	type fields struct {
		id     string
	}
	type args struct {
		ID    string
		pass  string
		fecha string
		hora  string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want string
		wantErr bool
	}{
		{name: "", fields: fields{id: "test"}, args: args{
			ID:    "fjama",
			pass:  "12345678",
			fecha: "18-02-2022",
			hora:  "3.46",
		},want: "fjama1234567818-02-20223.46", wantErr: false},//no hay manera de retornar Error
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &dbnotificator{
				id:     tt.fields.id,
			}
			d.Notify(tt.args.ID, tt.args.pass, tt.args.fecha, tt.args.hora)
			if (d.aux!= tt.want) && tt.wantErr{
				t.Fail()
			}
		})
	}
}
