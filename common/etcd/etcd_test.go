package etcd

import (
	"testing"
)

func TestGetValue(t *testing.T) {
	_conn, _ := GetValue("/banerwai/mongo/conn")

	if _conn != "localhost:27017" {
		t.Errorf("etcd GetValue error")
	}
}

func TestGetService(t *testing.T) {
	results, _ := GetService("/banerwai/mongo")

	if results.Node.Nodes[0].Value != "localhost:27017" {
		t.Errorf("etcd TestGetService error")
	}

	if results.Node.Nodes[1].Value != "banerwai" {
		t.Errorf("etcd TestGetService error")
	}
}
