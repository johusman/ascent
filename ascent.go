package ascent

import (
    "ascent/mutations"
    "ascent/specimens"
)

type AscentEngine interface {
    Mutations() mutations.MutationRepository
    SetGenerationCallback(func(specimens.Specimen))
    Run(start specimens.Specimen, fitness func(specimens.Specimen) (float32, bool)) specimens.Specimen
}

type engine struct {
    mutations mutations.MutationRepository
    generationCallback func(specimens.Specimen)
}

func New() AscentEngine {
    return &engine{ mutations.NewRepository(), func(specimens.Specimen){} }
}

func (this *engine) Mutations() mutations.MutationRepository {
    return this.mutations
}

func (this *engine) SetGenerationCallback(callback func(specimens.Specimen)) {
    this.generationCallback = callback
}

func (this *engine) Run(start specimens.Specimen, fitness func(specimens.Specimen) (float32, bool)) specimens.Specimen {

    seed := start

    for true {
        nextseed := seed
        highscore, _ := fitness(seed)

        for i := 0; i < 100; i++ {
            offspring := seed.Clone()
            this.mutations.Mutate(offspring)
            score, stop := fitness(offspring)

            if score > highscore {
                nextseed = offspring
                highscore = score
            }

            if stop {
                return offspring
            }
        }

        seed = nextseed
        this.generationCallback(seed)
    }

    return nil
}
