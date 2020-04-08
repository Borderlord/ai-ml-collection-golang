package decision_tree

const (
	CONTINUOUS  = "continuous"
	CATEGORICAL = "categorical"
)

type DecisionTreeDataTrain struct {
	Features    []string
	Data        [][]interface{}
	Type        []string
	TargetClass []string
}