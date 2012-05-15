package ascent

import "testing"
import "ascent/specimens"
import "ascent/mutations"
import "math/rand"

type mockspecimen struct {
}

func (s *mockspecimen) Clone() specimens.Specimen {
    return &(*s) // hopefully a shallow copy
}

type mockmutation struct {
    Used int
}

func (m *mockmutation) Mutate(specimen specimens.Specimen) {
    m.Used++
}

func assertInRange(t *testing.T, message string, expectedLow, expectedHigh, actual int) {
    if actual >= expectedLow && actual <= expectedHigh {
            return
    }
    t.Fatalf(message + ": expected range [%d, %d] but was %d", expectedLow, expectedHigh, actual)
}

func TestMutationRepository(t *testing.T) {
    rand.Seed(12345)

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

    print(mockmutations[0].Used)

    assertInRange(t, "Unexpected frequency", 2800, 3200, mockmutations[0].Used)
    assertInRange(t, "Unexpected frequency", 2300, 2700, mockmutations[1].Used)
    assertInRange(t, "Unexpected frequency", 1800, 2200, mockmutations[2].Used)
    assertInRange(t, "Unexpected frequency", 1300, 1700, mockmutations[3].Used)
    assertInRange(t, "Unexpected frequency",  800, 1200, mockmutations[4].Used)
}
