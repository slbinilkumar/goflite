package goflite

import "testing"

func TestSynthesisWithInvalidVoice(t *testing.T) {
	voicename := "invalid"
	_, err := TextToWave("Testing", voicename)
	if err == nil {
		t.Errorf("Synthesis should fail when voicename is invalid")
	}
}

func TestSynthesisWithDefaultVoice(t *testing.T) {
	voicename := defaultVoiceName
	_, err := TextToWave("Hello World", voicename)
	if err != nil {
		t.Errorf("Synthesis with default voice should not fail")
	}

}