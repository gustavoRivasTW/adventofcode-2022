package main

import (
	"adventofcode/utils"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Node struct {
	childs []*Node
	parent *Node
	Data   Data
}

type Data struct {
	files   []File
	folders []string
	name    string
	size    int
}

type File struct {
	name string
	size int
}

func New(data Data, parent *Node) *Node {
	return &Node{
		Data:   data,
		parent: parent,
	}
}

func NewFile(name string, size int) File {
	return File{name: name, size: size}
}

func (d *Data) addFile(data string) {
	info := strings.Split(data, " ")
	s, _ := strconv.Atoi(info[0])
	file := NewFile(info[1], int(s))
	d.files = append(d.files, file)
}

func (d *Data) addName(name string) {
	d.name = name
}
func (d *Data) addFolder(folder string) {
	d.folders = append(d.folders, folder)
}
func (t *Node) addChild(n *Node) {
	t.childs = append(t.childs, n)
}

func (t *Node) print() {
	if t.parent != nil {
		fmt.Println("name " + t.parent.Data.name + "/" + t.Data.name + ", files : ")
	} else {
		fmt.Println("name " + t.Data.name + ", files : ")
	}
	for _, v := range t.childs {
		v.print()
	}
}

func checkIsDir(name string) bool {
	return strings.HasPrefix(name, "dir")
}
func checkIsFile(name string) bool {
	match, _ := regexp.MatchString(`^\d`, name)
	return match
}

func checkIsChangeDir(name string) bool {
	return strings.HasPrefix(name, "$ cd")
}

func buildDirectory(data string) *Node {
	cmds := strings.Split(data, "\n")
	node := New(Data{name: "root"}, nil)
	for _, v := range cmds {
		if checkIsFile(v) {
			node.Data.addFile(v)
		}
		if checkIsDir(v) {
			node.Data.folders = append(node.Data.folders, v)
			child := New(Data{name: v[4:]}, node)
			node.addChild(child)
		}
		if checkIsChangeDir(v) {
			if strings.Contains(v, "..") {
				node = node.parent
			} else {
				for _, n := range node.childs {
					if n.Data.name == v[5:] {
						node = n
					}
				}
			}
		}
	}

	for true {
		if node.parent != nil {
			node = node.parent
		} else {
			break
		}
	}
	return node
}

func (n *Node) calculateSizePerDirectory() {
	n.Data.size = 0
	for _, v := range n.Data.files {
		n.Data.size += v.size
	}
	for _, v := range n.childs {
		v.calculateSizePerDirectory()
		n.Data.size += v.Data.size
	}
}

func filterFolderBySizeLessThan(n *Node, size int) []*Node {
	var folders []*Node
	for _, c := range n.childs {
		if c.Data.size < size {
			folders = append(folders, c)
		}
		folders = append(folders, filterFolderBySizeLessThan(c, size)...)
	}
	return folders
}

func filterFolderBySizeMoreEqualThan(n *Node, minSize int) []*Node {
	var folders []*Node
	for _, c := range n.childs {
		//	folders = append(folders, c)
		if c.Data.size > minSize {
			folders = append(folders, c)
		}
		folders = append(folders, filterFolderBySizeMoreEqualThan(c, minSize)...)
	}
	return folders
}

func getTotalOfFolderBySizeLessThan(data string, amount int) int {
	node := buildDirectory(data)
	node.calculateSizePerDirectory()
	folder := filterFolderBySizeLessThan(node, amount)

	var total int
	for _, k := range folder {
		total += k.Data.size
	}
	return total
}

func findFolderSizeNearTo(n *Node, number int) int {

	folder := filterFolderBySizeMoreEqualThan(n, number)

	var sizes []int
	for _, k := range folder {
		sizes = append(sizes, k.Data.size)
	}
	sort.Ints(sizes[:])
	fmt.Println(sizes)
	return sizes[0]
}
func partOne() {
	//Remover primer $ cd /
	data := utils.ReadFile("../input-data.txt")
	size := getTotalOfFolderBySizeLessThan(data, 100000)
	fmt.Printf("Total Size %d", size)
}

func partTwo() {
	//Remover primer $ cd /
	data := utils.ReadFile("../input-data.txt")
	node := buildDirectory(data)
	node.calculateSizePerDirectory()
	fmt.Printf("\n Total Size: %d ", node.Data.size)
	required := (70000000 - node.Data.size - 30000000) * -1 // (Total - Used - Required)
	smallest := findFolderSizeNearTo(node, required)
	fmt.Printf("\nSmallest Elimination %d", smallest)
}
func main() {
	partTwo()
}
