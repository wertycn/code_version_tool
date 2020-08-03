package app_test

import (
	"F10-CLI/app"
	"testing"
)

func TestIsExist(t *testing.T) {
	if app.IsExist("==========/") {
		t.Error("error")
	}

	if app.IsExist("app") {
		t.Log("pass")
	}
}

func TestCreateFile(t *testing.T) {
	app.CreateFile("test", "Test")
	if app.IsExist("./version/Test.md") {
		t.Log("pass")
	} else {
		t.Error("error")
	}
}

func TestGetTemplateContent(t *testing.T) {
	content := app.GetTemplateContent()
	t.Log(content)
}

func TestReplaceContent(t *testing.T) {
	content := app.GetTemplateContent()
	var version app.VersionInfo
	version.CodeChangeCommitInfo = "代码改动commit 日志"
	version.CodeChangeFileCount = "代码改动统计数据"
	version.CodeChangeFileInfo = "代码变动文件清单"
	version.Version = "TEST_ONE"
	version.Development = "debug.icu"
	version.ProjectGitUrl = "git url"
	version.DateTime = "2020-08-04 18:00:00"
	cont := app.ReplaceContent(content, version)
	app.CreateFile(cont, "TEST")
	t.Log("pass")
}
