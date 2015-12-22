package NEAT

type Neuron struct{
  incoming []Gene
  value float64
}

func NewNeuron() *Neuron{
  return &Neuron{[]Gene{},0}
}
