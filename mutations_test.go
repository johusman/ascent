package ascent

import "testing"
import "ascent/specimens"
import "ascent/mutations"
import "math/rand"

type mockspecimen struct {
}

func (s *mockspecimen) Clone() specimens.Specimen {
    clone := *s
    return &clone
}

func (s *mockspecimen) ToString() string {
    return "A featureless specimen"
}

type mutationCounter struct {
    used int
    mutation func(specimen specimens.Specimen)
}

func mockmutation() *mutationCounter {
    mock := mutationCounter{0, nil}
    mock.mutation = func(specimen specimens.Specimen) {
        mock.used++
    }
    return &mock
}

func assertInRange(t *testing.T, message string, expectedLow, expectedHigh, actual int) {
    if actual >= expectedLow && actual <= expectedHigh {
            return
    }
    t.Fatalf(message + ": expected range [%d, %d] but was %d", expectedLow, expectedHigh, actual)
}

func TestMutationRepository(t *testing.T) {
    rand.Seed(95873)

    repo := mutations.NewRepository()

    counters := []*mutationCounter{ mockmutation(), mockmutation(), mockmutation(), mockmutation(), mockmutation() }

    repo.Register(counters[0].mutation, 0.3)
    repo.Register(counters[1].mutation, 0.25)
    repo.Register(counters[2].mutation, 0.2)
    repo.Register(counters[3].mutation, 0.15)
    repo.Register(counters[4].mutation, 0.1)

    specimen := mockspecimen{}

    for i := 0; i < 10000; i++ {
        repo.Mutate(&specimen)
    }

    assertInRange(t, "Unexpected frequency", 2800, 3200, counters[0].used)
    assertInRange(t, "Unexpected frequency", 2300, 2700, counters[1].used)
    assertInRange(t, "Unexpected frequency", 1800, 2200, counters[2].used)
    assertInRange(t, "Unexpected frequency", 1300, 1700, counters[3].used)
    assertInRange(t, "Unexpected frequency",  800, 1200, counters[4].used)
}
