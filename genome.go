package NEAT

import (
  "sort"
  "math/rand"
  "math"
)

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
  mutationCopy := make(map[string]float64)
  for k,v := range c.mutations{
    mutationCopy[k] = v
  }
 return &Genome{make([]Gene,c.maxNodes + c.nbOutputs),0,0,make([]Neuron,0),0,0,mutationCopy}
 }

func CopyGenome(g *Genome, c ConstContainer) *Genome{
  res := NewGenome(c)
  for _,x := range g.genes {
    res.genes = append(res.genes,*CopyGene(&x))
  }
  res.maxneuron = g.maxneuron
  return res
}

func BasicGenome(nbInputs int,mutations map[string]float64, c ConstContainer) *Genome{
  res := NewGenome(c)
  res.maxneuron = nbInputs
  mutate(res, c)
  return res
}

// Generates the full neuronic network from genes
func GenerateNetwork(g *Genome, c ConstContainer){
  for i := 0; i < c.nbInputs; i++{
    g.network[i] = *NewNeuron()
  }
  for i := 0; i < c.nbOutputs; i++{
    g.network[c.maxNodes + i] = *NewNeuron()
  }
  sort.Sort(GeneSlice(g.genes))

  seen := make(map[int]bool)
  for _,x := range g.genes {
    if x.enabled {
      if _,ok := seen[x.to]; !ok {
         g.network[x.to] = *NewNeuron()
      }
      seen[x.to] = true
      neuron := g.network[x.to]
      neuron.incoming = append(neuron.incoming,x)
      if _,ok := seen[x.from]; !ok {
        g.network[x.from] = *NewNeuron()
      }
      seen[x.from] = true
    }
  }
}

func crossover(g1, g2 *Genome, c ConstContainer) *Genome{
  // g1 has the highest fitness
  if (g2.fitness > g1.fitness){
    g2,g1 = g1,g2
  }

  // Make a child genome
  child := NewGenome(c)

  // Start a record of the innovations. May be put to a slice if i get the time later
  inno := make(map[int]Gene)
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

func evaluateNetwork(g *Genome, c ConstContainer, inputs []float64) []bool{

  // Add value for bias
  inputs = append(inputs,1)

  for i := 0; i < c.nbInputs; i++ {
    g.network[i].value = inputs[i]
  }

  for _, neuron := range g.network  {
    sum := 0.0
    // For now, we have sigmoid transformation neurons
    for _, gene := range neuron.incoming {
      other := g.network[gene.from]
      sum = sum + gene.weight * other.value
    }
    if sum > 0 {
      neuron.value = 1/(1 + math.Exp(-sum))
    }
  }

  outputs := make([]bool,c.nbOutputs)
  for o := 0; o < c.nbOutputs; o++ {
    outputs[o] = g.network[c.maxNodes+o].value > 0
  }

  return outputs
}

func weightMutate(g *Genome, c ConstContainer){
  step := c.mutations["step"]
  pertubation := c.mutations["pertubation"]
  for _,gene := range g.genes {
    if rand.Float64() < pertubation {
      gene.weight += rand.Float64()*step*2 - step
    } else {
      gene.weight = rand.Float64() * 4 - 2
    }
  }
}

func linkMutate(g *Genome, forceBias bool, c ConstContainer){
  neuron1 := randomNeuron(g,true,c)
  neuron2 := randomNeuron(g,false,c)

  newLink := NewGene()

  newLink.from = neuron1
  newLink.to = neuron2
  if forceBias {
    newLink.from = c.nbInputs
  }
}

func neuronMutate(g *Genome, globalInno *int){
  if len(g.genes) == 0 {
    return
  }

  g.maxneuron++
  geneDestroyed := g.genes[rand.Intn(len(g.genes))]
  if !geneDestroyed.enabled {
    return
  }
  *globalInno += 1
  gene1 := CopyGene(&geneDestroyed)
  gene1.to = g.maxneuron
  gene1.weight = 1.0
  gene1.innovation = *globalInno
  g.genes = append(g.genes,*gene1)

  *globalInno += 1
  gene2 := CopyGene(&geneDestroyed)
  gene2.from = g.maxneuron
  g.genes = append(g.genes,*gene2)

  geneDestroyed.enabled = false
}

func toggleAbleMutate(g *Genome, disable bool){
  possible := make([]Gene,0)
  for _,gene := range g.genes{
    if gene.enabled == disable {
      possible = append(possible,gene)
    }
  }

  if len(possible) == 0 {
    return
  }

  possible[rand.Intn(len(possible))].enabled = !disable
}

func mutate(g *Genome, c ConstContainer){
  for mutation, _ := range g.mutationRates{
    if rand.Intn(2) == 0 {
      g.mutationRates[mutation] *=0.95
    } else {
      g.mutationRates[mutation] *=1.0522
    }
  }

  if rand.Float64() < g.mutationRates["mConnection"]{
    weightMutate(g,c)
  }

  proba := g.mutationRates["mLink"]
  for proba > 0 {
    if rand.Float64() < proba {
      linkMutate(g,false,c)
    }
    proba -= 1
  }

  proba = g.mutationRates["mBias"]
  for proba > 0 {
    if rand.Float64() < proba {
      linkMutate(g,true,c)
    }
    proba -= 1
  }

  proba = g.mutationRates["enable"]
  for proba > 0 {
    if rand.Float64() < proba {
       toggleAbleMutate(g,false)
    }
    proba -= 1
  }

  proba = g.mutationRates["disable"]
  for proba > 0 {
    if rand.Float64() < proba {
       toggleAbleMutate(g,true)
    }
    proba -= 1
  }
}
