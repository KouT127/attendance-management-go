package payloads

import "testing"


func TestUserPayload_Validate(t *testing.T) {
	type fields struct {
		Name  string
		Email string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "Should validate",
			fields:  fields{
				Name:  "sato taro",
				Email: "sato@taro.com",
			},
			wantErr: false,
		},
		{
			name:    "Should not validate when name is max length",
			fields:  fields{
				Name:  "012345678901234567890123456789012345678901234567890123456789",
				Email: "sato",
			},
			wantErr: true,
		},
		{
			name:    "Should not validate when email isn't Email",
			fields:  fields{
				Name:  "sato taro",
				Email: "sato",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserPayload{
				Name:  tt.fields.Name,
				Email: tt.fields.Email,
			}
			if err := u.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
