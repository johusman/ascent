package mutations

import (
    "ascent/specimens"
    "container/list"
)

type Mutation interface {
    Mutate(specimen *specimens.Specimen)
}

type MutationRepository interface {
    Register(mutation *Mutation, chance float32)

    IdentityChance() float32
    SetIdentityChance(chance float32)

    Repeat() int
    SetRepeat(count int)

    Mutate(specimen *specimens.Specimen)
}

type registeredMutation struct {
    mutation *Mutation
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

func (r *repository) Register(mutation *Mutation, chance float32) {
    r.mutations.PushBack(registeredMutation{mutation, chance})
}

func (r *repository) IdentityChance() float32 { return r.identityChance }
func (r *repository) SetIdentityChance(chance float32) { r.identityChance = chance }

func (r *repository) Repeat() int { return r.repeat }
func (r *repository) SetRepeat(count int) { r.repeat = count }

func (r *repository) Mutate(specimen *specimens.Specimen) {

}
