package getembedurl

import "testing"

func TestGetEmbedUrl(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{url: "https://www.youtube.com/watch?v=mA4McnyPRZk"},
			want: "https://www.youtube.com/embed/mA4McnyPRZk?feature=oembed",
		},
		{
			name: "test2",
			args: args{url: "https://www.youtube.com/watch?v=D0wudnBXQGY"},
			want: "https://www.youtube.com/embed/D0wudnBXQGY?feature=oembed",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEmbedUrl(tt.args.url); got != tt.want {
				t.Errorf("GetEmbedUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
