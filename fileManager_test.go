package fm

import (
	"os"
	"reflect"
	"testing"
)

func TestNewFileManager(t *testing.T) {
	tests := []struct {
		name string
		want Manager
	}{
		{"1", &FileManager{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFileManager(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFileManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileManager_OpenFile(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		fm      *FileManager
		args    args
		want    string
		wantErr bool
	}{
		{"1", &FileManager{}, args{"fixtures/readfile.txt"}, "line1\nline2\nline3\nline4", false},
		{"2", &FileManager{}, args{"fixtures/nofile.txt"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fm.OpenFile(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileManager.OpenFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FileManager.OpenFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFileManager_OpenFile(t *testing.B) {
	fm := &FileManager{}
	t.ResetTimer()
	for index := 0; index < t.N; index++ {
		fm.OpenFile("fixtures/readfile.txt")
	}
}

func TestFileManager_WriteFile(t *testing.T) {
	type args struct {
		file        string
		data        []byte
		permissions uint32
	}
	tests := []struct {
		name    string
		fm      *FileManager
		args    args
		wantErr bool
	}{
		{"1", &FileManager{}, args{"fixtures/writefile.txt", []byte("testline\ntestline2"), 0}, false},
		{"2", &FileManager{}, args{"fixtures/newfile.txt", []byte("firstline\nsecondline"), 0}, false},
		{"3", &FileManager{}, args{"fixtures/writefile.txt", []byte("testline\ntestline2"), 0777}, false},
		{"4", &FileManager{}, args{"fixtures/newfilewithpermissions.txt", []byte("testline\ntestline2"), 0777}, false},
		{"5", &FileManager{}, args{"fixtures////", []byte("invalid\nfile\nname"), 0777}, true},
		{"6", &FileManager{}, args{"fixtures////", []byte("invalid\nfile\nname"), 0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fm.WriteFile(tt.args.file, tt.args.data, tt.args.permissions); (err != nil) != tt.wantErr {
				t.Errorf("FileManager.WriteFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	err := os.Remove("fixtures/newfile.txt")
	if err != nil {
		t.Fatalf("removing file: fixtures/newfile.txt")
	}
	err = os.Remove("fixtures/newfilewithpermissions.txt")
	if err != nil {
		t.Fatalf("removing file: fixtures/newfilewithpermissions.txt")
	}
}

func BenchmarkFileManager_WriteFile(t *testing.B) {
	fm := &FileManager{}
	t.ResetTimer()
	for index := 0; index < t.N; index++ {
		fm.WriteFile("fixtures/benchwritefile.txt", []byte("testline\ntestline2"), 0777)
	}
	t.StopTimer()
	err := os.Remove("fixtures/benchwritefile.txt")
	if err != nil {
		t.Fatalf("removing file: fixtures/benchwritefile.txt")
	}
}

func TestFileManager_ExistsFile(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		fm   *FileManager
		args args
		want bool
	}{
		{"1", &FileManager{}, args{"fixtures/readfile.txt"}, true},
		{"2", &FileManager{}, args{"fixtures/nofile.txt"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fm.ExistsFile(tt.args.file)
			if got != tt.want {
				t.Errorf("FileManager.ExistsFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFileManager_ExistsFile(t *testing.B) {
	fm := &FileManager{}
	t.ResetTimer()
	for index := 0; index < t.N; index++ {
		fm.ExistsFile("fixtures/readfile.txt")
	}
}
