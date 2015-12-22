package NEAT

import (
  "math/rand"
)

type Gene struct{
  from int
  to int
  weight float64
  enabled bool
  innovation int
}

func NewGene() *Gene{
  return &Gene{0,0,0,true,0}
}

func CopyGene(g *Gene) *Gene{
  return &Gene{g.from,g.to,g.weight,g.enabled,g.innovation}
}

func containsLink(genes []Gene, link Gene) bool{
  for _,x := range(genes){
    if x.from == link.from && x.to == link.to {
      return true
    }
  }
  return false
}

func randomNeuron(g *Genome, notInput bool, c ConstContainer) int {
  if notInput {
     return rand.Intn(len(g.network)-c.nbInputs) + c.nbInputs
  }
  return rand.Intn(len(g.network))
}


func disjoint(g1, g2 []Gene) (res int){
  res = len(g1) + len(g2)
  innos := make(map[int]bool)
  for _,g := range g1 {
    innos[g.innovation] = true
  }
  for _,g := range g2 {
    if _,ok := innos[g.innovation]; ok {
      res--
    }
  }
  return
}

// Implementation to sort a slice of Genes
// Go has the bad habit of naming a slice type ._. No generecism sucks

type GeneSlice []Gene

func (a GeneSlice) Len() int {return len(a)}
func (a GeneSlice) Swap(i,j int) {a[i],a[j] = a[j],a[i]}
func (a GeneSlice) Less(i,j int) bool {return a[i].to < a[j].to}
