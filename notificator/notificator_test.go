package notificator

import (
	"reflect"
	"testing"
)

func Test_notificator_Notify(t *testing.T) {
	notifi := notificator{}
	t1 := NewDbnotificator("m")
	t2 := NewTextfilenotificator("b")
	notifi.Subscribe(t1)
	notifi.Subscribe(t2)
	wantErr := false
	notifi.Notify("test", "test", "test", "test")

	if (notifi.errAux == nil) && wantErr {
		t.Fail()
	}

}

func Test_notificator_Subscribe(t *testing.T) {
	notifi := notificator{}
	t1 := NewDbnotificator("m")
	t2 := NewTextfilenotificator("b")
	notifi.Subscribe(t1)
	notifi.Subscribe(t2)

	if len(notifi.notificators) != 2 {
		t.Fail()
	}
}

func Test_notificator_Unsubscribe(t *testing.T) {
	notifi := notificator{}
	t1 := NewTextfilenotificator("m")
	t2 := NewDbnotificator("b")
	notifi.Subscribe(t1)
	notifi.Subscribe(t2)
	notifi.Unsubscribe(t2)

	if len(notifi.notificators) != 1 {
		t.Fail()
	}

	for _, notificator := range notifi.notificators {
		if notificator.GetID() == t2.GetID() {
			t.Fail()
		}
	}
}

func Test_notificator_UnsubscribeNoData(t *testing.T) {
	notifi := notificator{}
	t2 := NewDbnotificator("b")
	notifi.Unsubscribe(t2)

	if len(notifi.notificators) != 0 {
		t.Fail()
	}

	for _, notificator := range notifi.notificators {
		if notificator.GetID() == t2.GetID() {
			t.Fail()
		}
	}
}

func TestNewNotificator(t *testing.T) {
	tests := []struct {
		name string
		want *notificator
	}{
		{name: "", want: &notificator{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNotificator(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNotificator() = %v, want %v", got, tt.want)
			}
		})
	}
}