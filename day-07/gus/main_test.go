package main

import (
	"testing"
)

const Given string = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log(
5626152 d.ext
7214296 k`

func TestIsDir(t *testing.T) {
	t.Run("Is a Dir", func(t *testing.T) {
		given := "dir drzllllv"

		actual := checkIsDir(given)

		if !actual {
			t.Errorf("got %v given %s", actual, given)
		}
	})
}

func TestCheckIsFile(t *testing.T) {
	t.Run("Is a File", func(t *testing.T) {
		given := "114478 gtpgsvv"

		actual := checkIsFile(given)

		if !actual {
			t.Errorf("got %v given %s", actual, given)
		}
	})
	t.Run("Is not a File", func(t *testing.T) {
		given := "dir folder-1"

		actual := checkIsFile(given)

		if actual {
			t.Errorf("got %v given %s", actual, given)
		}
	})
}

func TestBuildDirectory(t *testing.T) {
	given := `$ ls
1 data.txt
dir folder-1
dir folder-2
$ cd folder-1
$ ls
2 data.txt
$ cd ..`

	actual := buildDirectory(given)
	actualChildOne := actual.childs[0]
	actualChildTwo := actual.childs[1]

	if actual.Data.files[0].name != "data.txt" && actual.Data.files[0].size != 1 {
		t.Errorf("got %v want %v given %s", actual.Data.files[0], "data.txt", given)
	}
	if actualChildOne.Data.name != "folder-1" {
		t.Errorf("got %v want %v given %s", actualChildOne.Data.name, "folder-1", given)
	}
	if actualChildTwo.Data.name != "folder-2" {
		t.Errorf("got %v want %v given %s", actualChildTwo.Data.name, "folder-2", given)
	}
	if actualChildOne.Data.files[0].name != "data.txt" && actualChildOne.Data.files[0].size != 2 {
		t.Errorf("got %v want %v given %s", actualChildOne.Data.files[0], "data.txt", given)
	}
}

func TestCalculateSizePerDirectory(t *testing.T) {
	dir := `$ ls
dir folder-1
dir folder-2
$ cd folder-1
$ ls
dir folder-11
$ cd folder-11
$ ls
1 data.txt
2 data.txt
$ cd ..`
	node := buildDirectory(dir)
	expected := 3
	node.calculateSizePerDirectory()

	if node.Data.size != expected {
		t.Errorf("got %d want %d", node.Data.size, expected)
	}
}

func TestDeleteDirectoryGivenCondition(t *testing.T) {
	given := `$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`
	expected := 95437

	actual := getTotalOfFolderBySizeLessThan(given, 100000)

	if actual != expected {
		t.Errorf("got %d want %d", actual, expected)
	}

}

/*
Total 	 70.000.000
used 	 41.072.511
required 30.000.000
free	 28.927.489
missing   1.072.511
option	  
*/

func TestFindFolderSizeNearTo(t *testing.T) {
	//given := 8381165

	t.Errorf("got %d want %d", 1, 2)
}
