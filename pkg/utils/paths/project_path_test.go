package paths

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"reflect"
	"testing"
)

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
}

func TestAsProjectPath(t *testing.T) {
	tmpDir := t.TempDir()
	tmpSubDir := filepath.Join(tmpDir, "/any/path")

	wd, err := os.Getwd()
	if err != nil {
		t.Fatal("could not load current working directory")
	}

	type args struct {
		dir string
	}
	tests := []struct {
		name string
		args args
		want *ProjectPath
	}{
		{
			name: "Test any absolute dir",
			args: args{dir: tmpSubDir},
			want: &ProjectPath{
				Path:   tmpSubDir,
				Name:   "path",
				Parent: filepath.Join(tmpDir, "/any"),
			},
		},
		{
			name: "Test current directory",
			args: args{dir: "."},
			want: &ProjectPath{
				Path:   wd,
				Name:   filepath.Base(wd),
				Parent: filepath.Dir(wd),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AsProjectPath(tt.args.dir); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsProjectPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProjectPath_MkDirAll(t *testing.T) {
	tmpDir := t.TempDir()
	tmpSubDir := filepath.Join(tmpDir, "/any/path")
	err := os.MkdirAll(tmpSubDir, 0o777)
	if err != nil {
		t.Fatalf("MkdirAll %q: %s", tmpSubDir, err)
	}

	type fields struct {
		Path   string
		Name   string
		Parent string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Path already exists",
			fields: fields{
				Path:   tmpSubDir,
				Name:   "path",
				Parent: filepath.Join(tmpDir, "/any"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ProjectPath{
				Path:   tt.fields.Path,
				Name:   tt.fields.Name,
				Parent: tt.fields.Parent,
			}
			p.MkDirAll()
		})
	}
}

func TestProjectPath_IsEmpty(t *testing.T) {
	// make temp paths for tests
	tmpDir := t.TempDir()
	tmpSubDir := filepath.Join(tmpDir, "/folder1")
	tmpSubSubDir := filepath.Join(tmpSubDir, "/folder2")

	// make subdirectories
	err := os.MkdirAll(tmpSubSubDir, 0o777)
	if err != nil {
		t.Fatalf("MkdirAll %q: %s", tmpSubSubDir, err)
	}

	// make file
	tmpFile, errCreateTemp := os.CreateTemp(tmpSubDir, "*")
	if errCreateTemp != nil {
		t.Fatalf("create %q: %s", tmpFile.Name(), err)
	}
	defer closeFile(tmpFile)

	type fields struct {
		Path   string
		Name   string
		Parent string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Not empty: has child directory",
			fields: fields{
				Path:   tmpDir,
				Name:   filepath.Base(tmpDir),
				Parent: filepath.Dir(tmpDir),
			},
			want: false,
		},
		{
			name: "Not empty: has child directory and file",
			fields: fields{
				Path:   tmpSubDir,
				Name:   "folder1",
				Parent: tmpDir,
			},
			want: false,
		},
		{
			name: "Empty directory",
			fields: fields{
				Path:   tmpSubSubDir,
				Name:   "folder2",
				Parent: tmpSubDir,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ProjectPath{
				Path:   tt.fields.Path,
				Name:   tt.fields.Name,
				Parent: tt.fields.Parent,
			}
			if got := p.IsEmpty(); got != tt.want {
				t.Errorf("ProjectPath.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProjectPath_Exists(t *testing.T) {
	// make temp paths for tests
	tmpDir := t.TempDir()
	tmpSubDir := filepath.Join(tmpDir, "/folder1")
	tmpSubSubDir := filepath.Join(tmpSubDir, "/folder2")
	tmpDirNotExists := filepath.Join(tmpSubSubDir, "/doesnotexist")

	// make subdirectories
	err := os.MkdirAll(tmpSubSubDir, 0o777)
	if err != nil {
		t.Fatalf("MkdirAll %q: %s", tmpSubSubDir, err)
	}

	// make file
	tmpFile, errCreateTemp := os.CreateTemp(tmpSubDir, "*")
	if errCreateTemp != nil {
		t.Fatalf("create %q: %s", tmpFile.Name(), err)
	}
	defer closeFile(tmpFile)

	type fields struct {
		Path   string
		Name   string
		Parent string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Exists: has child dir",
			fields: fields{
				Path:   tmpDir,
				Name:   filepath.Base(tmpDir),
				Parent: filepath.Dir(tmpDir),
			},
			want: true,
		},
		{
			name: "Exists: has child directory and file",
			fields: fields{
				Path:   tmpSubDir,
				Name:   "folder1",
				Parent: tmpDir,
			},
			want: true,
		},
		{
			name: "Exists: Empty directory",
			fields: fields{
				Path:   tmpSubSubDir,
				Name:   "folder2",
				Parent: tmpSubDir,
			},
			want: true,
		},
		{
			name: "Does not exist",
			fields: fields{
				Path:   tmpDirNotExists,
				Name:   "doesnotexist",
				Parent: tmpSubSubDir,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ProjectPath{
				Path:   tt.fields.Path,
				Name:   tt.fields.Name,
				Parent: tt.fields.Parent,
			}
			if got := p.Exists(); got != tt.want {
				t.Errorf("ProjectPath.Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_resolveTildePaths(t *testing.T) {
	// make temp paths for tests
	tmpDir := t.TempDir()
	tmpSubDirWithTilde := filepath.Join(tmpDir, "/~folder1")

	usr, err := user.Current()
	if err != nil {
		t.Fatalf("user.Current(): %s", err)
	}

	homeDir := usr.HomeDir

	// make subdirectories
	err = os.MkdirAll(tmpSubDirWithTilde, 0o777)
	if err != nil {
		t.Fatalf("MkdirAll %q: %s", tmpSubDirWithTilde, err)
	}

	type args struct {
		dir string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Tilde not home dir",
			args: args{
				dir: tmpSubDirWithTilde,
			},
			want: tmpSubDirWithTilde,
		},
		{
			name: "Tilde is home dir",
			args: args{
				dir: "~",
			},
			want: homeDir,
		},
		{
			name: "Tilde is home dir with slash",
			args: args{
				dir: "~/",
			},
			want: homeDir,
		},
		{
			name: "Tilde is home dir with slash and child dir",
			args: args{
				dir: "~/folder1",
			},
			want: filepath.Join(homeDir, "/folder1"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resolveTildePaths(tt.args.dir); got != tt.want {
				t.Errorf("resolveTildePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
