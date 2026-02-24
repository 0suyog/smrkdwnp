package parser_test

import (
	"io"
	"log"
	"strings"
	"testing"

	"github.com/0suyog/smrkdwnp/parser"
)

func TestParse(t *testing.T) {
	log.SetOutput(io.Discard)
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		f    io.Reader
		want string
	}{
		struct {
			name string
			f    io.Reader
			want string
		}{
			name: "indented codeblock, heading1, paragraph",
			f: strings.NewReader(`
    apple
   # apple
        ball
    cat
dog
apple
				`),
			want: `<body><pre><code>apple</code></pre><h1>apple</h1><pre><code>    ball
cat</code></pre><p>dog apple</p></body>`,
		},
		struct {
			name string
			f    io.Reader
			want string
		}{
			name: "paragraph, thematic break",
			f:    strings.NewReader("cow\n___"),
			want: "<body><p>cow</p><hr /></body>",
		},
		struct {
			name string
			f    io.Reader
			want string
		}{
			name: "setex heading 2",
			f:    strings.NewReader("cow\n-"),
			want: "<body><h2>cow</h2></body>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parser.Parse(tt.f)
			if tt.want != got {
				t.Errorf("Parse() =\n%v,\n want =\n%v", got, tt.want)
			}
		})
	}
}
