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
    picksHighestScore(t, 1)
}

func TestPicksHighestScoreMultiThreaded(t *testing.T) {
    picksHighestScore(t, 4)
}


func picksHighestScore(t *testing.T, threads int) {
    engine := New()

    engine.Mutations().SetIdentityChance(1.0)

    var counter int = 0

    seed := &idspecimen{counter, func(this *idspecimen) *idspecimen {
        counter += 1
        return &idspecimen{counter, this.clone}
    }}

    engine.SetGenerationCallback(func(winner specimens.Specimen) (bool) {
        if (int(winner.(*idspecimen).id) != counter) {
            t.Fatalf("Highest score did not win! Winner had score %d but counter is %d", winner.(*idspecimen).id, counter-1)
        }
        return counter <= 100
    })

    engine.Run(threads, seed, func(specimen specimens.Specimen) (float32) {
        // use id as score
        return float32(specimen.(*idspecimen).id)
    })
}
