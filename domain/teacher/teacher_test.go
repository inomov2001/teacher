package teacher

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestFactory_New(t *testing.T) {
	type args struct {
		firstName   string
		lastName    string
		email       string
		phonenumber string
		subjectID   uuid.UUID
	}
	tests := []struct {
		name    string
		args    args
		want    Teacher
		wantErr bool
	}{
		struct {
			name    string
			args    args
			want    Teacher
			wantErr bool
		}{
			name: "should pass",
			args: args{
				firstName:   "Po'lat",
				lastName:    "Nazarov",
				email:       "polat@nazarov.com",
				phonenumber: "+998977777777",
				subjectID:   testSubjectID,
			},
			want: Teacher{
				iD:         testTeacherID,
				firstName:   "Po'lat",
				lastName:    "Nazarov",
				email:       "polat@nazarov.com",
				phonenumber: "+998977777777",
				subjectID:   testSubjectID,
			},
			wantErr: false,
		},

		struct {
			name    string
			args    args
			want    Teacher
			wantErr bool
		}{
			name: "should failed with empty first name",
			args: args{
				firstName:   "",
				lastName:    "Nazarov",
				email:       "polat@nazarov.com",
				phonenumber: "+998977777777",
				subjectID:   testSubjectID,
			},
			want:    Teacher{},
			wantErr: true,
		},
		struct {
			name    string
			args    args
			want    Teacher
			wantErr bool
		}{
			name: "should failed with empty last name",
			args: args{
				firstName:   "Po'lat",
				lastName:    "",
				email:       "polat@nazarov.com",
				phonenumber: "+998977777777",
				subjectID:   testSubjectID,
			},
			want:    Teacher{},
			wantErr: true,
		},
		struct {
			name    string
			args    args
			want    Teacher
			wantErr bool
		}{
			name: "should failed with empty invalid email",
			args: args{
				firstName:   "Po'lat",
				lastName:    "Nazarov",
				email:       "polat.com",
				phonenumber: "+998977777777",
				subjectID:   testSubjectID,
			},
			want:    Teacher{},
			wantErr: true,
		},
		struct {
			name    string
			args    args
			want    Teacher
			wantErr bool
		}{
			name: "should failed with empty invalid phone number",
			args: args{
				firstName:   "Po'lat",
				lastName:    "Nazarov",
				email:       "polat@nazarov.com",
				phonenumber: "+99897777777",
				subjectID:   testSubjectID,
			},
			want:    Teacher{},
			wantErr: true,
		},
	}

	f := NewFactory(testIDGenerator{})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := f.NewTeacher (tt.args.firstName, tt.args.lastName, tt.args.email, tt.args.phonenumber, tt.args.subjectID)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

var (
	testSubjectID = uuid.New()
	testTeacherID = uuid.New()
)

type testIDGenerator struct{}

func (g testIDGenerator) GeneratorID() uuid.UUID {
	return testTeacherID
}
