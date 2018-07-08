package cbdatastore

import (
	"fmt"
	"strings"

	"github.com/chvck/gocb"
)

// CBDataStore is a type of DataStore backing onto a Couchbase Database
type CBDataStore struct {
	cluster *gocb.Cluster
	bucket  *gocb.Bucket
}

// NewCBDataStore creates and returns a new CBDataStore
func NewCBDataStore(host string, port uint, bucketName, username, password string) (*CBDataStore, error) {
	connString := fmt.Sprintf("http://%s:%d", host, port)
	cluster, err := gocb.Connect(connString)
	if err != nil {
		return nil, err
	}

	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: username,
		Password: password,
	})

	bucket, err := cluster.OpenBucket(bucketName, "")
	if err != nil {
		return nil, err
	}

	dataStore := CBDataStore{
		cluster: cluster,
		bucket:  bucket,
	}

	return &dataStore, nil
}

func checkModelID(id, userID string) bool {
	splitID := strings.Split(id, "::")
	user := splitID[1]

	return user == userID
}
