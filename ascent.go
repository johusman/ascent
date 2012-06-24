package ascent

import (
    "ascent/mutations"
    "ascent/specimens"
)

type AscentEngine interface {
    Mutations() mutations.MutationRepository
    SetGenerationCallback(func(specimens.Specimen)(bool))
    Run(threads int, start specimens.Specimen, fitness func(specimens.Specimen) (float32)) specimens.Specimen
}

type engine struct {
    mutations mutations.MutationRepository
    generationCallback func(specimens.Specimen) (bool)
}

func New() AscentEngine {
    return &engine{ mutations.NewRepository(), func(specimens.Specimen) (bool) { return true } }
}

func (this *engine) Mutations() mutations.MutationRepository {
    return this.mutations
}

func (this *engine) SetGenerationCallback(callback func(specimens.Specimen) (bool)) {
    this.generationCallback = callback
}

type rated struct {
    specimen specimens.Specimen
    score float32
}

func (this *engine) Run(threads int, start specimens.Specimen, fitness func(specimens.Specimen) (float32)) specimens.Specimen {

    perGeneration := 2 * 2 * 3 * 3 * 5 // Some nice divisors
    perThread := perGeneration / threads

    winner := rated{ start, fitness(start) }

    for true {
        winnerChan := make(chan rated)
        for i := 0; i < threads; i++ {
            go func() {
                currentTop := winner
                for i := 0; i < perThread; i++ {
                    offspring := winner.specimen.Clone()
                    this.mutations.Mutate(offspring)
                    score := fitness(offspring)

                    if score > currentTop.score {
                        currentTop.specimen = offspring
                        currentTop.score = score
                    }
                }
                winnerChan <- currentTop
            }()
        }

        nextWinner := winner
        for i := 0; i < threads; i++ {
            subWinner := <-winnerChan
            if subWinner.score > nextWinner.score {
                nextWinner = subWinner
            }
        }

        winner = nextWinner
        if !this.generationCallback(winner.specimen) {
            return winner.specimen
        }
    }

    return nil
}
