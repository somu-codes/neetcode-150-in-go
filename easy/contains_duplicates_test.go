package easy

import "testing"

/*
Given an integer array nums,
return true if any value appears at least twice in the array,
and return false if every element is distinct.
**/

func Test_containsDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Contains duplicate - simple",
			args: args{nums: []int{1, 2, 3, 1}},
			want: true,
		},
		{
			name: "No duplicates",
			args: args{nums: []int{1, 2, 3, 4}},
			want: false,
		},
		{
			name: "Contains multiple duplicates",
			args: args{nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("containsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
