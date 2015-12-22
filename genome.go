package NEAT

type Genome struct{
  genes []Gene
  fitness int
  adjustedFitness int
  network []Neuron
  maxneuron int
  globalRank int
  mutationRates map[string]float64
}

func NewGenome(Mutations map[string]float64) *Genome{
  return &Genome([]Gene{},0,0,[]Neuron{},0,0,Mutations)
}

func CopyGenome(g *Genome) *Genome{
  res := NewGenome(g.mutationRates)
  for x := range g.genes {
    res.genes = append(res.genes,CopyGene(x))
  }
  res.maxneuron = g.maxneuron
  return res
}

func BasicGenome(nbInputs int,mutations map[string]float64) *Genome{
  res := NewGenome(mutations)
  res.maxneuron = nbInputs
  mutate(res)
  return res
}

func GenerateNetwork(g *Genome, constContainer ConstContainer){
  for i := range constContainer.nbInputs{
    g.network = append(g.network,NewNeuron())
  }
  for i := range constContainer.nbOutputs
}

func mutate(g *Genome){
}
