package ascent

import "testing"
import "ascent/specimens"
import "ascent/mutations"

type mockspecimen struct {
}

type mockmutation struct {
    Used int
}

func (m *mockmutation) Mutate(specimen *specimens.Specimen) {
    m.Used++
}

func assertInRange(t *testing.T, message string, expectedLow, expectedHigh, actual float32) {
    if actual >= expectedLow && actual <= expectedHigh {
            return
    }
    t.Errorf(message + ": expected range [%g, %g] but was %g", expectedLow, expectedHigh, actual)
}

func TestMutationRepository(t *testing.T) {
    repo := mutations.NewRepository()

    mockmutations := []mockmutation{ mockmutation{}, mockmutation{}, mockmutation{}, mockmutation{}, mockmutation{} }

    repo.Register(&mockmutations[0], 0.3)
    repo.Register(&mockmutations[1], 0.25)
    repo.Register(&mockmutations[2], 0.2)
    repo.Register(&mockmutations[3], 0.15)
    repo.Register(&mockmutations[4], 0.1)

    specimen := mockspecimen{}

    for i := 0; i < 10000; i++ {
        repo.Mutate(&specimen)
    }

    t.assertInRange(t, "Unexpected frequency", 2500, 3500, mockmutations[0].Used)
    t.assertInRange(t, "Unexpected frequency", 2000, 3000, mockmutations[1].Used)
    t.assertInRange(t, "Unexpected frequency", 1500, 2500, mockmutations[2].Used)
    t.assertInRange(t, "Unexpected frequency", 1000, 2000, mockmutations[3].Used)
    t.assertInRange(t, "Unexpected frequency",  500, 1500, mockmutations[4].Used)
}
