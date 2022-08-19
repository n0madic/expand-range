package expandrange

import (
	"reflect"
	"testing"
)

func TestRange(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		rng     Range
		wantErr bool
	}{
		{
			name:    "empty",
			str:     "",
			rng:     Range{},
			wantErr: false,
		},
		{
			name:    "simple",
			str:     "1,3",
			rng:     Range{1, 3},
			wantErr: false,
		},
		{
			name:    "advanced",
			str:     "2,3-5,6,8-10,13",
			rng:     Range{2, 3, 4, 5, 6, 8, 9, 10, 13},
			wantErr: false,
		},
		{
			name:    "dupcheck",
			str:     "1-3,2",
			rng:     Range{1, 2, 3},
			wantErr: false,
		},
		{
			name:    "sorting",
			str:     "3,2,1",
			rng:     Range{1, 2, 3},
			wantErr: false,
		},
		{
			name:    "check convertation",
			str:     "1-e",
			wantErr: true,
		},
		{
			name:    "check greater",
			str:     "3-2",
			wantErr: true,
		},
		{
			name:    "check values count",
			str:     "1-2-3",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rng, err := Parse(tt.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("Range.Parse() error = %v, wantErr %v", err, tt.wantErr)
			} else if !tt.wantErr {
				rng.Sort()
				if !reflect.DeepEqual(tt.rng, rng) {
					t.Errorf("Expected range %v, got %v", tt.rng, rng)
				}
				for _, r := range tt.rng {
					if !rng.InRange(r) {
						t.Errorf("Error in the entry %v into a range %v", r, rng)
					}
				}
				if rng.InRange(999) {
					t.Errorf("Occurrence of non-existent element")
				}
			}
		})
	}
}
