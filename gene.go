package NEAT

type Gene struct{
  into int
  out int
  weight float64
  enabled bool
  innovation int
}

func NewGene() *Gene{
  return &Gene{0,0,0,true,0}
}

func CopyGene(g *Gene) *Gene{
  return &Gene{g.into,g.out,g.weight,g.enabled,g.innovation}
}

func containsLink(genes []Gene, link Gene) bool{
  for _,x := range(genes){
    if x.into == link.into && x.out == link.out {
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
func (a GeneSlice) Less(i,j int) bool {return a[i].out < a[j].out}
