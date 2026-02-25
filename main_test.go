package main

import "testing"

func TestToBrowserURL(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		// ssh:// format
		{
			name:  "ssh with user and .git suffix",
			input: "ssh://git@github.com/hoge/fuga.git",
			want:  "https://github.com/hoge/fuga",
		},
		{
			name:  "ssh with user without .git suffix",
			input: "ssh://git@github.com/hoge/fuga",
			want:  "https://github.com/hoge/fuga",
		},
		{
			name:  "ssh with custom port",
			input: "ssh://git@gitlab.example.com:2222/hoge/fuga.git",
			want:  "https://gitlab.example.com:2222/hoge/fuga",
		},
		{
			name:  "ssh without user",
			input: "ssh://github.com/hoge/fuga.git",
			want:  "https://github.com/hoge/fuga",
		},

		// git@ format
		{
			name:  "git@ github.com with .git suffix",
			input: "git@github.com:user/repo.git",
			want:  "https://github.com/user/repo",
		},
		{
			name:  "git@ github.com without .git suffix",
			input: "git@github.com:user/repo",
			want:  "https://github.com/user/repo",
		},
		{
			name:  "git@ gitlab.org",
			input: "git@gitlab.org:user/repo.git",
			want:  "https://gitlab.org/user/repo",
		},
		{
			name:  "git@ bitbucket.net",
			input: "git@bitbucket.net:user/repo.git",
			want:  "https://bitbucket.net/user/repo",
		},
		{
			name:  "git@ .io TLD",
			input: "git@gitea.example.io:user/repo.git",
			want:  "https://gitea.example.io/user/repo",
		},
		{
			name:  "git@ .dev TLD",
			input: "git@gitlab.example.dev:user/repo.git",
			want:  "https://gitlab.example.dev/user/repo",
		},
		{
			name:  "git@ self-hosted with subdomain",
			input: "git@git.company.co.jp:team/project.git",
			want:  "https://git.company.co.jp/team/project",
		},
		{
			name:    "git@ missing colon",
			input:   "git@github.com/user/repo.git",
			wantErr: true,
		},

		// https format (passthrough)
		{
			name:  "https URL passed through",
			input: "https://github.com/user/repo",
			want:  "https://github.com/user/repo",
		},
		{
			name:  "https URL with .git suffix passed through",
			input: "https://github.com/user/repo.git",
			want:  "https://github.com/user/repo.git",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toBrowserURL(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("toBrowserURL(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("toBrowserURL(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
