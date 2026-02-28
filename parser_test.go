package hlsdl

import (
	"net/url"
	"testing"

	"github.com/grafov/m3u8"
)

func Test(t *testing.T) {
	_, err := parseHlsSegments("https://cdn.theoplayer.com/video/big_buck_bunny_encrypted/stream-800/index.m3u8", nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestResolveMasterPlaylist(t *testing.T) {
	baseURL, _ := url.Parse("https://example.com/path/master.m3u8")

	master := m3u8.NewMasterPlaylist()
	master.Append("stream-400/index.m3u8", nil, m3u8.VariantParams{Bandwidth: 400000})
	master.Append("stream-800/index.m3u8", nil, m3u8.VariantParams{Bandwidth: 800000})
	master.Append("stream-1500/index.m3u8", nil, m3u8.VariantParams{Bandwidth: 1500000})

	resolved, err := resolveMasterPlaylist(master, baseURL)
	if err != nil {
		t.Fatal(err)
	}

	expected := "https://example.com/path/stream-1500/index.m3u8"
	if resolved != expected {
		t.Fatalf("expected %s, got %s", expected, resolved)
	}
}

func TestResolveMasterPlaylistAbsoluteURL(t *testing.T) {
	baseURL, _ := url.Parse("https://example.com/path/master.m3u8")

	master := m3u8.NewMasterPlaylist()
	master.Append("https://cdn.example.com/stream-400/index.m3u8", nil, m3u8.VariantParams{Bandwidth: 400000})
	master.Append("https://cdn.example.com/stream-800/index.m3u8", nil, m3u8.VariantParams{Bandwidth: 800000})

	resolved, err := resolveMasterPlaylist(master, baseURL)
	if err != nil {
		t.Fatal(err)
	}

	expected := "https://cdn.example.com/stream-800/index.m3u8"
	if resolved != expected {
		t.Fatalf("expected %s, got %s", expected, resolved)
	}
}

func TestResolveMasterPlaylistNoVariants(t *testing.T) {
	baseURL, _ := url.Parse("https://example.com/path/master.m3u8")

	master := m3u8.NewMasterPlaylist()

	_, err := resolveMasterPlaylist(master, baseURL)
	if err == nil {
		t.Fatal("expected error for empty master playlist")
	}
}
