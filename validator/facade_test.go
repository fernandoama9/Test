package validator

import (
	"pruebaLDev180222/notificator"
	"reflect"
	"testing"
)

func Test_validatorfacade_Validate(t *testing.T) {
	type fields struct {
		notificator notificator.INotificator
		validator   CredentialValidator
	}
	noti := notificator.NewNotificator()
	validator := GetInstance()
	type args struct {
		ID       string
		Password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    User
		wantErr bool
	}{
		{name: "success", fields: fields{
			notificator: noti,
			validator:   validator,
		}, args: args{
			ID:       "test",
			Password: "test",
		}, want: User{
			ID:    "test",
			pass:  "test",
			email: "test@test",
		}, wantErr: false},
		{name: "fallo", fields: fields{
			notificator: noti,
			validator:   validator,
		}, args: args{
			ID:       "fallo",
			Password: "test",
		}, want: User{}, wantErr: true},
		{name: "fallo_no_validator_set", fields: fields{
			notificator: noti,
		}, args: args{
			ID:       "fallo",
			Password: "test",
		}, want: User{}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &validatorfacade{
				notificator: tt.fields.notificator,
				validator:   tt.fields.validator,
			}
			got, err := b.Validate(tt.args.ID, tt.args.Password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Validate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewValidatorFacade1(t *testing.T) {
	type args struct {
		notificator notificator.INotificator
		validator   CredentialValidator
	}
	validator := GetInstance()
	notificador := notificator.NewNotificator()
	tests := []struct {
		name string
		args args
		want *validatorfacade
	}{
		{name: "", args: args{validator: validator, notificator: notificador}, want: &validatorfacade{validator: validator, notificator: notificador}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewValidatorFacade(tt.args.notificator, tt.args.validator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewValidatorFacade() = %v, want %v", got, tt.want)
			}
		})
	}
}
