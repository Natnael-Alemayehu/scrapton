package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetH1FromHtml(t *testing.T) {
	cases := []struct {
		name      string
		inputhtml string
		expected  string
	}{
		{
			name: "normal operation",
			inputhtml: `
			<html>
			<body>
				<h1>Welcome to scrapton</h1>
				<main>
				<p>Learn to code by building real projects.</p>
				<p>This is the second paragraph.</p>
				</main>
			</body>
			</html>
			`,
			expected: "Welcome to scrapton",
		},
		{
			name: "Empty H1",
			inputhtml: `
			<html>
			<body>
				<main>
				<p>Learn to code by building real projects.</p>
				<p>This is the second paragraph.</p>
				</main>
			</body>
			</html>
			`,
			expected: "",
		},
		{
			name: "Multiple H1 tags",
			inputhtml: `
			<html>
			<body>
				<h1>Welcome to scrapton</h1>
				<main>
				<h1>Welcome to main</h1>
				<p>Learn to code by building real projects.</p>
				<p>This is the second paragraph.</p>
				</main>
			</body>
			</html>
			`,
			expected: "Welcome to scrapton",
		},
	}
	for i, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := getH1FromHTML(tc.inputhtml)
			if actual != tc.expected {
				t.Errorf("Test: %v - FAIL: expected stirng: %v, actual: %v", i, tc.expected, actual)
			}
		})
	}
}

func TestGetFirstParagraphFromHtml(t *testing.T) {
	cases := []struct {
		name      string
		inputhtml string
		expected  string
	}{
		{
			name: "normal operation",
			inputhtml: `
			<html>
			<body>
				<h1>Welcome to scrapton</h1>
				<main>
				<p>Learn to code by building real projects.</p>
				<p>This is the second paragraph.</p>
				</main>
			</body>
			</html>
			`,
			expected: "Learn to code by building real projects.",
		},
		{
			name: "Empty H1",
			inputhtml: `
			<html>
			<body>
				<main>
				<h1>Learn to code by building real projects.</h1>
				</main>
			</body>
			</html>
			`,
			expected: "",
		},
		{
			name: "Empty H1",
			inputhtml: `
			<html>
			<body>
				<main>
				<h1>This is actually an h1.</h1>
				<p>Learn to code by building real projects.</p>
				</main>
			</body>
			</html>
			`,
			expected: "Learn to code by building real projects.",
		},
	}
	for i, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := getFirstParagraphFromHTML(tc.inputhtml)
			if actual != tc.expected {
				t.Errorf("Test: %v - FAIL: expected stirng: %v, actual: %v", i, tc.expected, actual)
			}
		})
	}
}

func TestGetURLsFromHTMLAbsolute(t *testing.T) {
	cases := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:      "normal operation",
			inputURL:  `https://blog.boot.dev`,
			inputBody: `<html><body><a href="https://blog.boot.dev">Learn Development</a></body></html>`,
			expected:  []string{"https://blog.boot.dev"},
		},
	}
	for i, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			baseURL, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Errorf("Test: %v - FAIL: unexpected err: %v, input: %v", i, err, tc.inputURL)
			}

			actual, err := getURLsFromHTML(tc.inputBody, baseURL)
			if err != nil {
				t.Errorf("Test: %v - FAIL: unexpected err: %v, input: %v", i, err, tc.inputBody)
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test: %v - FAIL: expected stirng: %v, actual: %v", i, tc.expected, actual)
			}
		})
	}
}

func TestGetImagesFromHTML(t *testing.T) {
	cases := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:      "normal operation",
			inputURL:  `https://blog.boot.dev`,
			inputBody: `<html><body><img src="/logo.png" alt="Logo"></body></html>`,
			expected:  []string{"https://blog.boot.dev/logo.png"},
		},
		{
			name:     "multiple images",
			inputURL: `https://blog.boot.dev`,
			inputBody: `<html><body>
		<img src="/logo.png" alt="Logo">
		<img src="https://cdn.boot.dev/banner.jpg">
	</body></html>`,
			expected: []string{
				"https://blog.boot.dev/logo.png",
				"https://cdn.boot.dev/banner.jpg",
			},
		},
	}
	for i, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			baseURL, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Errorf("Test: %v - FAIL: unexpected err: %v, input: %v", i, err, tc.inputURL)
			}

			actual, err := getImagesFromHTML(tc.inputBody, baseURL)
			if err != nil {
				t.Errorf("Test: %v - FAIL: unexpected err: %v, input: %v", i, err, tc.inputBody)
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test: %v - FAIL: expected stirng: %v, actual: %v", i, tc.expected, actual)
			}
		})
	}
}
