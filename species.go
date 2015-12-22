package NEAT

type Species struct{
  topFitness int
  staleness int
  genomes []Genome
  averageFitness int
}

func NewSpecies() *Species{
  return &Species{0,0,[]Genome{},0}
}
