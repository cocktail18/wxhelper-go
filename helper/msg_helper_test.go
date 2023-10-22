package helper

import (
	"fmt"
	"github.com/beevik/etree"
	"testing"
)

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

func Test_decodeAtMsg(t *testing.T) {
	source := "<msgsource>\n\t<atuserlist><![CDATA[wxid_5yvnwqyxf3e421]]></atuserlist>\n\t<signature>v1_LW5k2iMD</signature>\n\t<tmp_node>\n\t\t<publisher-id></publisher-id>\n\t</tmp_node>\n</msgsource>\n"
	doc := etree.NewDocument()
	if err := doc.ReadFromString(source); err == nil {
		atWxIds := doc.FindElement("/msgsource/atuserlist").Text()
		fmt.Println(atWxIds)
	} else {
		t.Error(err)
	}
}
