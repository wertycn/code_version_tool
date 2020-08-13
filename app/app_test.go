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
	app.GetTemplateContent()
	t.Log("pass")
}

func TestReplaceContent(t *testing.T) {
	content := app.GetTemplateContent()
	var version app.VersionInfoType
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

func TestGetCommitLogInfo(t *testing.T) {
	info := app.GetCommitLogInfo("dev")
	t.Log(info)
}

func TestFormatCommitLog(t *testing.T) {
	info := app.GetCommitLogInfo("dev")
	commit, authors := app.FormatCommitLog(info)
	t.Log(commit)
	t.Log(authors)
}

func TestCreateVersionFile(t *testing.T) {
	app.SetGitMainBranch("dev")
	app.CreateVersionFile("LOCAL_TEST")
	t.Log("PASS...")
}

func TestGetRemoteUrl(t *testing.T) {
	app.SetGitRemote("mygitlab")
	url, name := app.GetProjectNameAndRemoteUrl()
	t.Log("pass", url, name)
}

func TestGetChangeFileInfo(t *testing.T) {
	info := app.GetChangeFileInfo("dev")
	t.Log(info)
}

func TestGetMainDir(t *testing.T) {
	projectPath, err := app.GetMainDir()
	if err != nil {
		t.Log(projectPath)
	}
	t.Log(projectPath)
}
