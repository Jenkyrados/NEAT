package NEAT

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

// Implementation to sort a slice of Genes
// Go has the bad habit of naming a slice type ._. No generecism sucks

type GeneSlice []Gene

func (a GeneSlice) Len() int {return len(a)}
func (a GeneSlice) Swap(i,j int) {a[i],a[j] = a[j],a[i]}
func (a GeneSlice) Less(i,j int) bool {return a[i].to < a[j].to}
