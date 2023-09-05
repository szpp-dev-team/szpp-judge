package langs

type LangID string

const (
	ImagePrefix = "szpp-judge-image-"
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

var langMetas = []Meta{
	{
		ID:          "c/11/gcc13.2",
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
		ID:          "cpp/20/gcc13.2",
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
		ID:          "java/21/openjdk",
		Name:        "Java (OpenJDK 21)",
		Active:      true,
		DockerImage: ImagePrefix + "openjdk21",
		SourceFile:  "Main.java",
		CompileCmd:  []string{"javac", "Main.java"},
		ExecCmd:     []string{"java", "-Xss1G", "-Xmx1G", "Main"},
	},
	{
		ID:          "python/3.11/cpython",
		Name:        "Python (CPython 3.11)",
		Active:      true,
		DockerImage: ImagePrefix + "cpython3.11",
		SourceFile:  "main.py",
		CompileCmd:  []string{"python3", "-m", "compileall", "main.py"},
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
