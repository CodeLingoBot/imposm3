package main

import (
	"flag"
	"fmt"
	"goposm/cache"
	"goposm/element"
	"goposm/parser"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
)

func parse(cache *cache.OSMCache, filename string) {
	nodes := make(chan []element.Node)
	coords := make(chan []element.Node)
	ways := make(chan []element.Way)
	relations := make(chan []element.Relation)

	positions := parser.PBFBlockPositions(filename)

	waitParser := sync.WaitGroup{}
	for i := 0; i < runtime.NumCPU(); i++ {
		waitParser.Add(1)
		go func() {
			for pos := range positions {
				parser.ParseBlock(pos, coords, nodes, ways, relations)
			}
			waitParser.Done()
		}()
	}

	waitCounter := sync.WaitGroup{}

	for i := 0; i < runtime.NumCPU(); i++ {
		waitCounter.Add(1)
		go func() {
			wayCounter := 0
			for ws := range ways {
				cache.Ways.PutWays(ws)
				wayCounter += len(ws)
			}
			fmt.Println("ways", wayCounter)
			waitCounter.Done()
		}()
	}
	for i := 0; i < runtime.NumCPU(); i++ {
		waitCounter.Add(1)
		go func() {
			relationCounter := 0
			for rels := range relations {
				cache.Relations.PutRelations(rels)
				relationCounter += len(rels)
			}
			fmt.Println("relations", relationCounter)
			waitCounter.Done()
		}()
	}
	for i := 0; i < runtime.NumCPU(); i++ {
		waitCounter.Add(1)
		go func() {
			nodeCounter := 0
			for nds := range coords {
				cache.Coords.PutCoords(nds)
				nodeCounter += len(nds)
			}
			fmt.Println("coords", nodeCounter)
			waitCounter.Done()
		}()
	}
	for i := 0; i < 2; i++ {
		waitCounter.Add(1)
		go func() {
			nodeCounter := 0
			for nds := range nodes {
				n, _ := cache.Nodes.PutNodes(nds)
				nodeCounter += n
			}
			fmt.Println("nodes", nodeCounter)
			waitCounter.Done()
		}()
	}

	waitParser.Wait()
	close(coords)
	close(nodes)
	close(ways)
	close(relations)
	waitCounter.Wait()
}

func main() {
	f, err := os.Create("/tmp/goposm.pprof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	log.SetFlags(log.LstdFlags | log.Llongfile)
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()
	cache, err := cache.NewOSMCache("/tmp/goposm")
	if err != nil {
		log.Fatal(err)
	}
	defer cache.Close()

	parse(cache, flag.Arg(0))

	rel := cache.Relations.Iter()
	for r := range rel {
		fmt.Println(r)
	}

	way := cache.Ways.Iter()
	i := 0
	for w := range way {
		i += 1
		cache.Coords.FillWay(w)
		//fmt.Println(i)
	}
	fmt.Println(i)
	//parser.PBFStats(os.Args[1])
	fmt.Println("done")
}
