package mutations

import (
    "ascent/specimens"
    "container/list"
    "math/rand"
)

type Mutation func(specimen specimens.Specimen)

type MutationRepository interface {
    Register(mutation Mutation, chance float32)

    IdentityChance() float32
    SetIdentityChance(chance float32)

    Repeat() int
    SetRepeat(count int)

    Mutate(specimen specimens.Specimen)
}

type registeredMutation struct {
    mutation Mutation
    chance float32
}

type repository struct {
    mutations *list.List
    identityChance float32
    repeat int
}

func NewRepository() MutationRepository {
    return &repository{list.New(), 0.0, 1}
}

func (r *repository) Register(mutation Mutation, chance float32) {
    r.mutations.PushBack(registeredMutation{mutation, chance})
}

func (r *repository) IdentityChance() float32 { return r.identityChance }
func (r *repository) SetIdentityChance(chance float32) { r.identityChance = chance }

func (r *repository) Repeat() int { return r.repeat }
func (r *repository) SetRepeat(count int) { r.repeat = count }

func (r *repository) Mutate(specimen specimens.Specimen) {
    for i := 0; i < r.repeat; i++ {
        r.mutateOnce(specimen)
    }
}

func (r *repository) mutateOnce(specimen specimens.Specimen) {
    point := rand.Float32() * r.sumChances()
    if point <= r.identityChance {
        return
    }

    point -= r.identityChance

    for e := r.mutations.Front(); e != nil; e = e.Next() {
        value := e.Value.(registeredMutation)
        if point <= value.chance {
            value.mutation(specimen)
            return
        }
        point -= value.chance
    }
}

func (r *repository) sumChances() float32 {
    sum := r.identityChance

    for e := r.mutations.Front(); e != nil; e = e.Next() {
        sum += e.Value.(registeredMutation).chance
    }

    return sum
}
