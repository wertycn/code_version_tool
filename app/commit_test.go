package app_test

import (
	"F10-CLI/app"
	"testing"
)

func TestCheckCommitTypePass(t *testing.T) {
	var typeList = []string{"feat", "fix", "refactor", "docs", "style", "test", "chore", "pref", "build", "ci", "revert"}
	for _, commitType := range typeList {
		if app.CheckCommitType(commitType) {
			//t.Log("check " + commitType + " pass")
		} else {
			t.Error("check type " + commitType + " not pass")
		}
	}
}

func TestCheckCommitTypeNotPass(t *testing.T) {
	var typeList = []string{"","00","asdf","qe","12341"}
	for _, commitType := range typeList {
		if app.CheckCommitType(commitType) {
			t.Error("check type " + commitType + " not pass")
		}
	}
}

func TestCheckCommitSubjectPass(t *testing.T) {
	var subjectList = []string{"中文有八个汉字呢","english is number should gt 24"}
	for _, subject := range subjectList {
		if false == app.CheckCommitSubjectLength(subject) {
			t.Error("check subject " + subject + " pass")
		}
	}
}

func TestCheckCommitSubjectNotPass(t *testing.T) {
	var subjectList = []string{"feat", "fix",  "docs", "style", "test", "chore", "pref", "build", "ci", "revert","中文字数有几个"}
	for _, subject := range subjectList {
		if app.CheckCommitSubjectLength(subject) {
			t.Error("check subject " + subject + " not pass")
		}
	}
}

func TestGitCommit(t *testing.T) {
	// TODO : 完整测试应该先添加变动再执行git add ,并再提交后回滚
	if app.GitCommit("feat","submit commit message test") {

	}else{
		t.Error("git commit failed ")
	}
}

func TestGitAdd(t *testing.T) {
	if app.GitAddAll(){

	}else{
		t.Error("git add all failed ")
	}
}