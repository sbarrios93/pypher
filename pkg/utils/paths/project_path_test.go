package paths

import (
	"os"
	"os/user"
	"path/filepath"
	"reflect"
	"testing"
)

func TestAsProjectPath(t *testing.T) {
	tmpDir := t.TempDir()
	tmpSubDir := tmpDir + "/any/path"

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
			name: "Test root directory '/'",
			args: args{dir: "/"},
			want: &ProjectPath{
				Path:   "/",
				Name:   "/",
				Parent: "/",
			},
		},
		{
			name: "Test any absolute dir",
			args: args{dir: tmpSubDir},
			want: &ProjectPath{
				Path:   tmpSubDir,
				Name:   "path",
				Parent: tmpDir + "/any",
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
	tmpSubDir := tmpDir + "/any/path"
	err := os.MkdirAll(tmpSubDir, 0777)
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
				Parent: tmpDir + "/any",
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
	tmpSubDir := tmpDir + "/folder1"
	tmpSubSubDir := tmpSubDir + "/folder2"
	tmpFile := tmpSubDir + "/file.txt"

	// make subdirectories
	err := os.MkdirAll(tmpSubSubDir, 0777)
	if err != nil {
		t.Fatalf("MkdirAll %q: %s", tmpSubSubDir, err)
	}

	// make file
	f, err := os.Create(tmpFile)
	if err != nil {
		t.Fatalf("create %q: %s", tmpFile, err)
	}
	defer f.Close()

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
	tmpSubDir := tmpDir + "/folder1"
	tmpSubSubDir := tmpSubDir + "/folder2"
	tmpFile := tmpSubDir + "/file.txt"
	tmpDirNotExists := tmpSubSubDir + "/doesnotexist"

	// make subdirectories
	err := os.MkdirAll(tmpSubSubDir, 0777)
	if err != nil {
		t.Fatalf("MkdirAll %q: %s", tmpSubSubDir, err)
	}

	// make file
	f, err := os.Create(tmpFile)
	if err != nil {
		t.Fatalf("create %q: %s", tmpFile, err)
	}
	defer f.Close()

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
	tmpSubDirWithTilde := tmpDir + "/~folder1"

	usr, err := user.Current()
	if err != nil {
		t.Fatalf("user.Current(): %s", err)
	}

	homeDir := usr.HomeDir

	// make subdirectories
	err = os.MkdirAll(tmpSubDirWithTilde, 0777)
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
			want: homeDir + "/folder1",
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
