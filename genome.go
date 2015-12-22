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
  for _,x := range g.genes {
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
  for i := 0; i < c.nbInputs; i++{
    g.network[i] = NewNeuron()
  }
  for i := 0; i < c.nbOutputs; i++{
    g.network[c.maxNodes + i] = NewNeuron()
  }
  sort.Sort(GeneSlice(g.genes))
  for _,x := range g.genes {
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

func crossover(g1 *Genome, g2 *Genome){
  // g1 has the highest fitness
  if (g2.fitness > g1.fitness){
    g2,g1 = g1,g2
  }

  // Make a child genome
  child := NewGenome()

  // Start a record of the innovations. May be put to a slice if i get the time later
  inno := make(map[int]*Gene)
  for _,gene := range g2.genes{
    inno[gene.innovation] = gene
  }

  if g1.fitness == g2.fitness { // If fitnesses are equal, all disjoint/excess are kept
    for _,gene1 := range g1.genes{
      gene2, ok := inno[gene1.innovation] // The matching gene in g2, if any
      if !ok || rand.Intn(2) == 0 { // There's no matching gene, or it failed its luck roll
        inno[gene1.innovation] = gene1
      } else {
         inno[gene1.innovation] = gene2
      }
    }
    // inno is a map of genes now, assigned by their innovation numbers
    for key := 0; key < len(inno); key ++ {
      child.genes = append(child.genes,inno[key])
    }
    child.maxneuron = len(child.genes)
  } else { // Fitnesses different : only disjoint excess from fittest parent are kept
    for _,gene1 := range g1.genes{
      gene2, ok := inno[gene1.innovation] // The matching gene in g2, if any
      if ok && rand.Intn(2) == 0 { // There's a matching gene, and it won its luck roll
        child.genes = append(child.genes,gene2)
      } else {
        child.genes = append(child.genes,gene1)
      }
    }
    child.maxneuron = g1.maxneuron // Since all the neurons of g1 are added, and no more
  }

  for mutation,rate := range g1.mutationRates {
     child.mutationRates[mutation] = rate
  }

  return child
}
