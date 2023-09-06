package langs

type LangID string

// 定数名のフォーマット： `言語名_バージョン_処理系`
// 定数値ののフォーマット： `言語名/バージョン/処理系`
//
// 整数値の enum で実装すると、将来、enum 列挙の行の途中に新たに言語を追加したとき、
// その行以降の言語の ID が変わってしまう。
// そのため文字列型で実装した。
const (
	C_11_GCC13         = LangID("c/11/gcc")
	CPP_20_GCC13       = LangID("cpp/20/gcc")
	JAVA_21_OPENJDK    = LangID("java/21/openjdk")
	PYTHON_311_CPYTHON = LangID("python/3.11/cpython")
)

type Meta struct {
	ID          LangID
	Name        string
	Active      bool
	DockerImage string
	SourceFile  string
	CompileCmd  []string
	ExecCmd     []string
}

const ImagePrefix = "szpp-judge-image-"

var langMetas = []Meta{
	{
		ID:          C_11_GCC13,
		Name:        "C11 (GCC 13.2)",
		Active:      true,
		DockerImage: ImagePrefix + "gcc13.2",
		SourceFile:  "main.c",
		CompileCmd: []string{
			"gcc",
			"-std=c11",
			"-I/opt/include",
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
	{
		ID:          CPP_20_GCC13,
		Name:        "C++20 (GCC 13.2)",
		Active:      true,
		DockerImage: ImagePrefix + "gcc13.2",
		SourceFile:  "main.cpp",
		CompileCmd: []string{
			"g++",
			"-std=c++20",
			"-I/opt/include",
			"-Wall",
			"-Wextra",
			"-DSZPP_JUDGE",
			"-O2",
			"-march=native",
			"-mtune=native",
			"main.cpp",
		},
		ExecCmd: []string{"./a.out"},
	},
	{
		ID:          JAVA_21_OPENJDK,
		Name:        "Java (OpenJDK 21)",
		Active:      true,
		DockerImage: ImagePrefix + "openjdk21",
		SourceFile:  "Main.java",
		CompileCmd:  []string{"javac", "Main.java"},
		ExecCmd:     []string{"java", "-Xss1G", "-Xmx1G", "Main"},
	},
	{
		ID:          PYTHON_311_CPYTHON,
		Name:        "Python (CPython 3.11)",
		Active:      true,
		DockerImage: ImagePrefix + "cpython3.11",
		SourceFile:  "main.py",
		CompileCmd:  []string{"python3", "-m", "compileall", "-q", "main.py"},
		ExecCmd:     []string{"python3", "main.py"},
	},
}

var langMetaIndex map[LangID]*Meta

func init() {
	langMetaIndex = make(map[LangID]*Meta, len(langMetas))

	// create langMetaIndex
	for i := range langMetas {
		langMetaIndex[langMetas[i].ID] = &langMetas[i]
	}
}

func Get(id LangID) (*Meta, bool) {
	m, ok := langMetaIndex[id]
	return m, ok
}
