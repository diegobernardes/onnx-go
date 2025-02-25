package onnx

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding"
	"gonum.org/v1/gonum/graph/simple"
	"gorgonia.org/tensor"
)

// testBackend is a simple directed graph
type testBackend struct {
	g *simple.WeightedDirectedGraph
}

// Node of the graph
type nodeTest struct {
	id          int64
	name        string
	description string
	value       tensor.Tensor
	opType      string
	//attributes  []*pb.Attribute
}

func (n *nodeTest) SetTensor(t tensor.Tensor) error {
	n.value = t
	return nil
}

func (n *nodeTest) GetTensor() tensor.Tensor {
	return n.value
}

// ID to fulfil the graph.Node interface
func (n *nodeTest) ID() int64 {
	return n.id
}

// Attributes it a method to fulfil the encoding/dot package
func (n *nodeTest) Attributes() []encoding.Attribute {
	var value string
	value = fmt.Sprintf(`%v`, n.name)
	if n.opType != "" {
		value = fmt.Sprintf(`%v|{Op|%v}`, value, n.opType)
	}
	if n.value != nil {
		value = fmt.Sprintf(`%v|{shape|%v}|{type|%v}`, value, n.value.Shape(), n.value.Dtype())
	}
	value = fmt.Sprintf(`"%v"`, value)
	return []encoding.Attribute{
		encoding.Attribute{
			Key:   "shape",
			Value: "Mrecord",
		},
		encoding.Attribute{
			Key:   "label",
			Value: value,
		},
	}
}

// SetName to fulfil the Namer interface
func (n *nodeTest) SetName(desc string) {
	n.name = desc
}

// GetName to fulfil the Namer interface
func (n *nodeTest) GetName() string {
	return n.name
}

// SetDescription to fulfil the Namer interface
func (n *nodeTest) SetDescription(desc string) {
	n.description = desc
}

// GetDescription to fulfil the Namer interface
func (n *nodeTest) GetDescription() string {
	return n.description
}

// ApplyTensor to fulfil the TensorCarrier interface
func (n *nodeTest) ApplyTensor(t tensor.Tensor) error {
	n.value = t
	return nil
}

// NewSimpleGraph ...
func newTestBackend() *testBackend {
	return &testBackend{
		g: simple.NewWeightedDirectedGraph(math.MaxFloat64, -1),
	}
}

func (g *testBackend) SetWeightedEdge(e graph.WeightedEdge) {
	g.g.SetWeightedEdge(e)

}
func (g *testBackend) NewWeightedEdge(from, to graph.Node, w float64) graph.WeightedEdge {
	return g.g.NewWeightedEdge(from, to, w)

}
func (g *testBackend) AddNode(n graph.Node) {
	g.g.AddNode(n)

}
func (g *testBackend) NewNode() graph.Node {
	n := g.g.NewNode()
	return &nodeTest{
		id: n.ID(),
	}
}

func (g *testBackend) Node(id int64) graph.Node {
	return g.g.Node(id)
}

func (g *testBackend) Nodes() graph.Nodes {
	return g.g.Nodes()
}

func (g *testBackend) From(id int64) graph.Nodes {
	return g.g.From(id)
}

func (g *testBackend) HasEdgeBetween(xid, yid int64) bool {
	return g.g.HasEdgeBetween(xid, yid)
}

func (g *testBackend) Edge(uid, vid int64) graph.Edge {
	return g.g.Edge(uid, vid)
}

func (g *testBackend) HasEdgeFromTo(uid, vid int64) bool {
	return g.g.HasEdgeFromTo(uid, vid)
}

func (g *testBackend) To(id int64) graph.Nodes {
	return g.g.To(id)
}

func (g *testBackend) ApplyOperation(_ Operation, _ ...graph.Node) error {
	return nil
}

// WeightedEdge returns the weighted edge from u to v if such an edge exists and nil otherwise.
// The node v must be directly reachable from u as defined by the From method.
func (g *testBackend) WeightedEdge(uid, vid int64) graph.WeightedEdge {
	return g.g.WeightedEdge(uid, vid)
}
