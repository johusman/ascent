package ascent

import (
    "ascent/mutations"
    "ascent/specimens"
)

type AscentEngine interface {
    Mutations() mutations.MutationRepository
    SetGenerationCallback(func([]specimens.Specimen))
    Run(start specimens.Specimen, fitness func(specimens.Specimen) (float32, bool)) specimens.Specimen
}

type engine struct {
    mutations mutations.MutationRepository
    generationCallback func([]specimens.Specimen)
}

func New() AscentEngine {
    return &engine{ mutations.NewRepository(), func([]specimens.Specimen){} }
}

func (this *engine) Mutations() mutations.MutationRepository {
    return this.mutations
}

func (this *engine) SetGenerationCallback(callback func([]specimens.Specimen)) {
    this.generationCallback = callback
}

func (this *engine) Run(start specimens.Specimen, fitness func(specimens.Specimen) (float32, bool)) specimens.Specimen {

    seed := make([]specimens.Specimen, 1, 1)
    seed[0] = start

    //pool := make([]specimens.Specimen, 100, 100)

    for true {
        nextseed := seed[0]
        highscore, _ := fitness(seed[0])

        for i := 0; i < 100; i++ {
            //pool[i] = seed[0].Clone()
            //this.mutations.Mutate(pool[i])
            offspring := seed[0].Clone()
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

        seed[0] = nextseed
        this.generationCallback(seed)
    }

    return nil
}
