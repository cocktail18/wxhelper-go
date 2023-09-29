package helper

import "testing"

func Test_getWxIdAndContentFromMsgContent(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name            string
		args            args
		wantWxId        string
		wantRealContent string
	}{
		// TODO: Add test cases.
		{
			"test",
			args{content: "wxid_qyxf3e421:\n123"},
			"wxid_qyxf3e421",
			"123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWxId, gotRealContent := getWxIdAndContentFromMsgContent(tt.args.content)
			if gotWxId != tt.wantWxId {
				t.Errorf("getWxIdAndContentFromMsgContent() gotWxId = %v, want %v", gotWxId, tt.wantWxId)
			}
			if gotRealContent != tt.wantRealContent {
				t.Errorf("getWxIdAndContentFromMsgContent() gotRealContent = %v, want %v", gotRealContent, tt.wantRealContent)
			}
		})
	}
}
