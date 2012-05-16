package ascent

import (
    "testing"
    "ascent/specimens"
)

type idspecimen struct {
    id int
    clone func(this *idspecimen) *idspecimen
}

func (s *idspecimen) Clone() specimens.Specimen {
    return s.clone(s)
}

func (s *idspecimen) ToString() string {
    return string(s.id)
}

func TestPicksHighestScore(t *testing.T) {
    engine := New()

    engine.Mutations().SetIdentityChance(1.0)

    var counter int = 0

    seed := &idspecimen{counter, func(this *idspecimen) *idspecimen {
        counter += 1
        return &idspecimen{counter, this.clone}
    }}

    engine.SetGenerationCallback(func(pool []specimens.Specimen) {
        if (int(pool[0].(*idspecimen).id) != counter) {
            t.Fatalf("Highest score did not win! Winner had score %d but counter is %d", pool[0].(*idspecimen).id, counter-1)
        }
    })

    engine.Run(seed, func(specimen specimens.Specimen) (float32, bool) {
        // use counter as score
        return float32(specimen.(*idspecimen).id), counter > 100
    })
}
