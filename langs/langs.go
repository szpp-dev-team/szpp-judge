package langs

import (
	"fmt"
	"strings"
)

type LangID string

// 定数名のフォーマット： `言語名_バージョン_処理系`
// 定数値のフォーマット： `言語名/バージョン/処理系`
//
// 整数値の enum で実装すると、将来、enum 列挙の行の途中に新たに言語を追加したとき、
// その行以降の言語の ID が変わってしまう。
// そのため文字列型で実装した。
const (
	C_11_GCC           = LangID("c/11/gcc")
	CPP_20_GCC         = LangID("cpp/20/gcc")
	SCRATCH_3_GCC      = LangID("scratch/3/gcc")
	JAVA_21_OPENJDK    = LangID("java/21/openjdk")
	PYTHON_311_CPYTHON = LangID("python/3.11/cpython")
	CPP_20_GCC_TESTLIB = LangID("cpp(testlib)/20/gcc")
)

type Meta struct {
	ID   LangID
	Name string

	// 提出可能な言語はこれでフィルタリングする。
	// 提出言語リストから削除したい場合はこのフィールドを false にする
	// (この langs.go から削除することはしないように！！)。
	Active bool

	// cpplib など。true にするとフロントエンド側に露出させない
	// (フロントエンドのコード生成の段階で除外する)。
	Internal bool

	DockerImage string
	SourceFile  string
	CompileCmd  []string
	ExecCmd     []string
}

const ImagePrefix = "szpp-judge-image-"

const gccVer = "13.2"
const gccDockerImage = ImagePrefix + "gcc" + gccVer

func name(langName, langVer, implName, implVer string) string {
	langName = strings.TrimSpace(langName)
	langVer = strings.TrimSpace(langVer)
	implName = strings.TrimSpace(implName)
	implVer = strings.TrimSpace(implVer)

	pad := ""
	if implVer != "" {
		pad = " "
	}
	return fmt.Sprintf("%s%s (%s%s%s)", langName, langVer, implName, pad, implVer)
}

func genCcCompileCmd(compiler, std, file string) []string {
	return []string{
		compiler,
		std,
		"-I/opt/include",
		"-lm",
		"-Wall",
		"-Wextra",
		"-DSZPP_JUDGE",
		"-O2",
		"-march=native",
		"-mtune=native",
		file,
	}
}

var LangMetas = []Meta{
	{
		ID:          C_11_GCC,
		Name:        name("C", "11", "GCC", gccVer),
		Active:      true,
		DockerImage: gccDockerImage,
		SourceFile:  "main.c",
		CompileCmd:  genCcCompileCmd("gcc", "-std=c11", "main.c"),
		ExecCmd:     []string{"./a.out"},
	},
	{
		ID:          CPP_20_GCC,
		Name:        name("C++", "20", "GCC", gccVer),
		Active:      true,
		DockerImage: gccDockerImage,
		SourceFile:  "main.cpp",
		CompileCmd:  genCcCompileCmd("g++", "-std=c++20", "main.cpp"),
		ExecCmd:     []string{"./a.out"},
	},
	{
		ID:          SCRATCH_3_GCC,
		Name:        name("Scratch", "3.0", "convert to C++20; GCC", gccVer),
		Active:      true,
		DockerImage: gccDockerImage,
		SourceFile:  "main.cpp",
		CompileCmd:  genCcCompileCmd("g++", "-std=c++20", "main.cpp"),
		ExecCmd:     []string{"./a.out"},
	},
	{
		ID:          JAVA_21_OPENJDK,
		Name:        name("Java", "21", "OpenJDK", ""),
		Active:      true,
		DockerImage: ImagePrefix + "openjdk21",
		SourceFile:  "Main.java",
		CompileCmd:  []string{"javac", "Main.java"},
		ExecCmd:     []string{"java", "-Xss1G", "-Xmx1G", "Main"},
	},

	{
		ID:          PYTHON_311_CPYTHON,
		Name:        name("Python", "3.11", "CPython", ""),
		Active:      true,
		DockerImage: ImagePrefix + "cpython3.11",
		SourceFile:  "main.py",
		CompileCmd:  []string{"python3", "-m", "compileall", "-q", "main.py"},
		ExecCmd:     []string{"python3", "main.py"},
	},
	{
		ID:          CPP_20_GCC_TESTLIB,
		Name:        name("C++(testlib)", "20", "GCC", gccVer),
		Active:      true,
		Internal:    true,
		DockerImage: gccDockerImage,
		SourceFile:  "checker.cpp",
		CompileCmd: []string{
			"g++",
			"-std=c++20",
			"-I/opt/include",
			"-lm",
			"-DSZPP_JUDGE",
			"-O2",
			"-march=native",
			"-mtune=native",
			"-o",
			"checker",
			"checker.cpp",
		},
		ExecCmd: []string{
			"./checker",
			"testcase_input.txt",
			"testcase_output.txt",
			"user_output.txt",
		},
	},
}

var langMetaIndex map[LangID]*Meta

func init() {
	langMetaIndex = make(map[LangID]*Meta, len(LangMetas))

	// create langMetaIndex
	for i := range LangMetas {
		langMetaIndex[LangMetas[i].ID] = &LangMetas[i]
	}
}

func Get(id LangID) (*Meta, bool) {
	m, ok := langMetaIndex[id]
	return m, ok
}
