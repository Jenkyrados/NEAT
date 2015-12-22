package NEAT

type Species struct{
  topFitness int
  staleness int
  genomes []Genomes
  averageFitness int
}

func NewSpecies() *Species{
  return &Species{0,0,[]Genome{},0}
}
