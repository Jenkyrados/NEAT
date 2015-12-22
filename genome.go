package NEAT

import "sort"

type Genome struct{
  genes []Gene
  fitness int
  adjustedFitness int
  network []Neuron
  maxneuron int
  globalRank int
  mutationRates map[string]float64
}

func NewGenome(c ConstContainer) *Genome{
  return &Genome(make([]Gene,c.maxNodes + c.nbOutputs),0,0,[]Neuron{},0,0,c.mutations)
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

func GenerateNetwork(g *Genome, c ConstContainer){
  for i := range c.nbInputs{
    g.network[i] = NewNeuron()
  }
  for i := range c.nbOutputs{
    g.network[c.maxNodes + i] = NewNeuron()
  }
  sort.Sort(GeneSlice(g.genes))
  for x := range g.genes{
    if x.enabled {
      if g.network[x.out] == nil{
         g.network[x.out] = NewNeuron()
      }
      neuron := g.network[x.out]
      neuron.incoming = append(neuron.incoming,x)
      if g.network[x.into] == nil {
        g.network[x.into] = NewNeuron()
      }
    }
  }
}

func mutate(g *Genome){
}
