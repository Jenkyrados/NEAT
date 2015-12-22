package NEAT

// Just a singleton containing all the constants necessary to the network
type ConstContainer struct{
  // Note : nbInputs is 1 more than the actual number of inputs, due to the need of a bias
  nbInputs int
  nbOutputs int
  population int
  deltaDisjoint float64
  deltaWeight float64
  deltaThreshold float64
  deltaExcess float64
  stale int
  mutations map[string]float64
  maxNodes int
}

func NewConstContainer(nbInputs, nbOutputs, population int, deltaDis, deltaWei, deltaThres, deltaExc float64,
stale int, step, pertubation, crossover, mLink, mNode, mBias, mDisable, mEnable, mConnection float64, maxNodes int) *ConstContainer{
  mutations := make(map[string]float64)
  mutations["mConnection"] = mConnection
  mutations["pertubation"] = pertubation
  mutations["crossover"] = crossover
  mutations["mLink"] = mLink
  mutations["mNode"] = mNode
  mutations["mBias"] = mBias
  mutations["mDisable"] = mDisable
  mutations["mEnable"] = mEnable
  mutations["step"] = step
  return &ConstContainer{nbInputs,nbOutputs,population,deltaDis,deltaWei,deltaThres,deltaExc,stale,mutations,maxNodes}
}
