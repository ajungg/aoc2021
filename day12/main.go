package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var possiblePaths int = 0

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	graph := graph{
		nodes: make(map[string][]string),
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		edge := strings.Split(scanner.Text(), "-")
		graph.addEdge(edge[0], edge[1])
	}

	p := pathFinder{visited: make(map[string]bool), graph: &graph}
	p.DFS("start")
	fmt.Println(p.numPaths)
}

type pathFinder struct {
	visited  map[string]bool
	graph    *graph
	numPaths int
}

func (p *pathFinder) DFS(node string) {
	if p.visited[node] && strings.ToLower(node) == node {
		return
	}

	if node == "end" {
		p.numPaths++
		return
	}

	p.visited[node] = true

	for _, e := range p.graph.getEdges(node) {
		p.DFS(e)
	}

	p.visited[node] = false
}

type graph struct {
	nodes map[string][]string
}

func (g *graph) addEdge(start string, end string) {
	g.nodes[start] = append(g.nodes[start], end)
	g.nodes[end] = append(g.nodes[end], start)
}

func (g *graph) getEdges(node string) []string {
	return g.nodes[node]
}
