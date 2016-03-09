package main

//import (
//	"github.com/blablacar/dgr/utils"
//	"github.com/n0rad/go-erlog/data"
//	"github.com/n0rad/go-erlog/logs"
//	"io/ioutil"
//	"os"
//	"path/filepath"
//)
//
//func discoverAndRunInitType(path string) {
//	fields := data.WithField("path", path)
//	if _, err := os.Stat(path); err != nil {
//		if err := os.MkdirAll(path, 0755); err != nil {
//			logs.WithEF(err, fields).Fatal("Cannot create path directory")
//		}
//	}
//
//	empty, err := utils.IsDirEmpty(path)
//	if err != nil {
//		logs.WithEF(err, fields).Fatal("Cannot read path directory")
//	}
//	if !initForce {
//		if !empty {
//			logs.WithEF(err, fields).Fatal("Path is not empty cannot init")
//		}
//	}
//
//	logs.WithEF(err, fields).Info("Init project")
//
//	files := make(map[string]string)
//
//	files[PATH_RUNLEVELS+PATH_PRESTART_EARLY+"/10.prestart-early.sh"] = `#!/dgr/bin/busybox sh
//source /dgr/bin/functions.sh
//isLevelEnabled "debug" && set -x
//
//echo "I'm a prestart early script that is run before templating"
//`
//	files[PATH_RUNLEVELS+PATH_PRESTART_LATE+"/10.prestart-late.sh"] = `#!/dgr/bin/busybox sh
//source /dgr/bin/functions.sh
//isLevelEnabled "debug" && set -x
//
//echo "I'm a prestart late script that is run after templating"
//`
//	files[PATH_RUNLEVELS+PATH_BUILD+"/10.install.sh"] = `#!/dgr/bin/busybox sh
//source /dgr/bin/functions.sh
//isLevelEnabled "debug" && set -x
//
//echo "I'm a build script that is run to install applications"
//`
//	files[PATH_RUNLEVELS+PATH_BUILD_SETUP+"/10.setup.sh"] = `#!/bin/sh
//source ${TARGET}/rootfs/dgr/bin/functions.sh
//isLevelEnabled "debug" && set -x
//
//echo "I'm build setup script file that is run from $BASEDIR to prepare $TARGET/rootfs before running build scripts"
//`
//	files[PATH_RUNLEVELS+PATH_BUILD_LATE+"/10.setup.sh"] = `#!/dgr/bin/busybox sh
//source /dgr/bin/functions.sh
//isLevelEnabled "debug" && set -x
//
//echo "I'm a build late script that is run to install applications after the copy of files,template,etc..."
//`
//	files[PATH_RUNLEVELS+PATH_INHERIT_BUILD_EARLY+"/10.inherit-build-early.sh"] = `#!/dgr/bin/busybox sh
//source /dgr/bin/functions.sh
//isLevelEnabled "debug" && set -x
//
//echo "I'm a inherit build early script that is run on this image and all images that have me as From during build"
//`
//	files[PATH_RUNLEVELS+PATH_INHERIT_BUILD_LATE+"/10.inherit-build-early.sh"] = `#!/dgr/bin/busybox sh
//source /dgr/bin/functions.sh
//isLevelEnabled "debug" && set -x
//
//echo "I'm a inherit build late script that is run on this image and all images that have me as From during build"
//`
//	files[PATH_FILES+"/etc/dummy"] = `Dummy file
//`
//	files[PATH_ATTRIBUTES+"/attributes.yml"] = `default:
//  dummy: world
//`
//	files[PATH_TEMPLATES+"/etc/templated.tmpl.cfg"] = `---
//uid: 0
//gid: 0
//mode: 0644
//checkCmd: true
//`
//	files[PATH_TEMPLATES+"/etc/templated.tmpl"] = `Hello {{ .dummy }}
//`
//	files[PATH_TEMPLATES+"/header.partial"] = `{{define "header"}}
//world
//{{end}}
//`
//	files[".gitignore"] = `target/
//`
//	files["aci-manifest.yml"] = `from: ""
//name: aci.example.com/aci-dummy:1
//aci:
//  app:
//    exec: [ "/dgr/bin/busybox", "sh" ]
//`
//	files[PATH_TESTS+"/dummy.bats"] = `#!/dgr/bin/bats
//
//@test "Prestart should template" {
//  result="$(cat /etc/templated)"
//  [ "$result" == "Hello world" ]
//}
//
//@test "dgr should copy files" {
//  result="$(cat /etc/dummy)"
//  [ "$result" == "Dummy file" ]
//}
//`
//	files[PATH_TESTS+"/wait.sh"] = `exit 0`
//
//	for filePath, data := range files {
//		fpath := path + "/" + filePath
//		os.MkdirAll(filepath.Dir(fpath), 0777)
//		ioutil.WriteFile(fpath, []byte(data), 0777)
//	}
//}