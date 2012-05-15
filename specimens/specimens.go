package specimens

type Specimen interface {
    Clone() Specimen
    ToString() string
}
