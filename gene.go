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
