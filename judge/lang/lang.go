package lang

import "fmt"

type LangID uint

type LangMeta struct {
	LangID         LangID
	Name           string
	ImageName      string
	SourceFileName string
	CompileCmd     []string
	ExecCmd        []string
}

const (
	C LangID = iota + 1
	CPP
	JAVA
	PYTHON
)

var langMetas map[LangID]*LangMeta = map[LangID]*LangMeta{
	C: {
		LangID:         C,
		Name:           "C11 (GCC 13.2)",
		ImageName:      "szpp-judge-images-gcc",
		SourceFileName: "main.c",
		CompileCmd: []string{
			"gcc",
			"-std=c11",
			"-Wall",
			"-Wextra",
			"-DSZPP_JUDGE",
			"-O2",
			"-march=native",
			"-mtune=native",
			"main.c",
		},
		ExecCmd: []string{"./a.out"},
	},
	CPP: {
		LangID:         CPP,
		Name:           "C++20 (GCC 13.2)",
		ImageName:      "szpp-judge-images-gcc",
		SourceFileName: "main.cpp",
		CompileCmd: []string{
			"g++",
			"-std=c++20",
			"-Wall",
			"-Wextra",
			"-DSZPP_JUDGE",
			"-I/opt/include",
			"-O2",
			"-march=native",
			"-mtune=native",
			"main.cpp",
		},
		ExecCmd: []string{"./a.out"},
	},
	JAVA: {
		LangID:         JAVA,
		Name:           "Java (OpenJDK 21)",
		ImageName:      "szpp-judge-images-java",
		SourceFileName: "Main.java",
		CompileCmd:     []string{"javac", "Main.java"},
		ExecCmd:        []string{"java", "-Xss1G", "-Xmx1G", "Main"},
	},
	PYTHON: {
		LangID:         PYTHON,
		Name:           "Python (CPython 3.11)",
		ImageName:      "szpp-judge-images-python",
		SourceFileName: "main.py",
		CompileCmd:     []string{"python3", "-c", "print(1)"},
		ExecCmd:        []string{"python3", "main.py"},
	},
}

func GetLangMeta(langID LangID) (*LangMeta, error) {
	if meta, ok := langMetas[langID]; ok {
		return meta, nil
	}

	return nil, fmt.Errorf("GetJudgeLangMeta(): unknown langID `%d`", langID)
}
