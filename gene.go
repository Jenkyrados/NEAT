package NEAT

import (
  "math"
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

// Returns the number of genes who are in one slice, not the other, and have an innovation
// number inferior to the other's (disjoint)

// Also returnes the number of genes who are in one slice, not the other, and
// are not disjoint (excess)

// Both are normalized by the number of genes in the larger genome

// TODO : check if the genes are guaranteed to be always sorted (avoids a test on gene)
func disjointExcess(g1, g2 []Gene) (disjoint, excess float64){
  disjoint = float64(len(g1) + len(g2))
  excess = 0.0
  innos := make(map[int]bool)

  maxinno1 := 0
  for _,g := range g1 {
    innos[g.innovation] = true
    if maxinno1 < g.innovation {
      maxinno1 = g.innovation
    }
  }

  maxinno2 := 0
  for _,g := range g2 {
    if g.innovation > maxinno1 {
      excess++
    } else if _,ok := innos[g.innovation]; ok {
      disjoint -= 2 // Two genes match
    }
    if maxinno2 < g.innovation {
       maxinno2 = g.innovation
    }
  }

  if maxinno2 < maxinno1 {
    for k,_ := range innos {
      if k > maxinno2 {
        disjoint-- // this is an excess, not a disjoint
        excess++
      }
    }
  }

  maxGene := math.Max(float64(len(g1)),float64(len(g2)))
  disjoint /= maxGene
  excess /= maxGene

  return
}

func weights(g1,g2 []Gene) float64{
  inno := make(map[int]Gene)

  for _,g := range g1 {
    inno[g.innovation] = g
  }

  result := 0.0
  matches := 0

  for _, g := range g2 {
    if v,ok := inno[g.innovation]; ok {
      result += math.Abs(g.weight-v.weight)
      matches++
    }
  }

  return result / float64(matches)
}

// Implementation to sort a slice of Genes
// Go has the bad habit of naming a slice type ._. No generecism sucks

type GeneSlice []Gene

func (a GeneSlice) Len() int {return len(a)}
func (a GeneSlice) Swap(i,j int) {a[i],a[j] = a[j],a[i]}
func (a GeneSlice) Less(i,j int) bool {return a[i].to < a[j].to}
