package meta

import (
	"fmt"
	base "github.com/sjwhitworth/golearn/base"
	eval "github.com/sjwhitworth/golearn/evaluation"
	filters "github.com/sjwhitworth/golearn/filters"
	trees "github.com/sjwhitworth/golearn/trees"
	"math/rand"
	"testing"
	"time"
)

func TestRandomForest1(testEnv *testing.T) {
	inst, err := base.ParseCSVToInstances("../examples/datasets/iris_headers.csv", true)
	if err != nil {
		panic(err)
	}

	rand.Seed(time.Now().UnixNano())
	insts := base.InstancesTrainTestSplit(inst, 0.6)
	filt := filters.NewBinningFilter(insts[0], 10)
	filt.AddAllNumericAttributes()
	filt.Build()
	filt.Run(insts[1])
	filt.Run(insts[0])
	rf := new(BaggedModel)
	rf.RandomFeatures = 2
	rf.SelectedFeatures = make(map[int][]base.Attribute)
	for i := 0; i < 10; i++ {
		rf.AddModel(trees.NewRandomTree(2))
	}
	rf.Fit(insts[0])
	fmt.Println(rf)
	predictions := rf.Predict(insts[1])
	fmt.Println(predictions)
	confusionMat := eval.GetConfusionMatrix(insts[1], predictions)
	fmt.Println(confusionMat)
	fmt.Println(eval.GetMacroPrecision(confusionMat))
	fmt.Println(eval.GetMacroRecall(confusionMat))
	fmt.Println(eval.GetSummary(confusionMat))
}
