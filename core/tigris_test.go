package core

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	s4 "github.com/aws/aws-sdk-go/service/s3"
	"github.com/stretchr/testify/require"
	"github.com/yandex-cloud/geesefs/core/cfg"
)

func TigrisDetected(flags *cfg.FlagStorage) bool {
	r, err := http.Get(flags.Endpoint + "/")
	if err != nil {
		return false
	}

	return strings.Contains(r.Header.Get("Server"), "Tigris")
}

func TestListIncludeMetadataAndContent(t *testing.T) {
	flags := cfg.DefaultFlags()

	conf := selectTestConfig(flags)
	conf.EnableSpecials = true
	flags.Backend = &conf
	flags.TigrisListContent = true

	if !TigrisDetected(flags) {
		t.Skip("Tigris not detected")
	}

	bucket := fmt.Sprintf("test-metadata-bucket-1-%x", rand.Int63())

	s3, err := NewS3(bucket, flags, &conf)
	require.NoError(t, err)

	_, err = s3.CreateBucket(&s4.CreateBucketInput{Bucket: aws.String(bucket)})
	require.NoError(t, err)

	blobs := []struct {
		name     string
		content  []byte
		metadata map[string]*string
	}{
		{"blob1", []byte("content of blob1"), nil},
		{"blob2", make([]byte, 16384+2048), map[string]*string{"key2": aws.String("value2"), "key5": aws.String("value5")}},
		{"blob3", []byte("content of blob3"), map[string]*string{"key3": aws.String("value3")}},
	}

	for _, blob := range blobs {
		input := &PutBlobInput{
			Key:      blob.name,
			Body:     bytes.NewReader(blob.content),
			Metadata: blob.metadata,
		}
		_, err := s3.PutBlob(input)
		require.NoError(t, err)
	}

	input := &ListBlobsInput{}
	listedBlobs, err := s3.ListBlobs(input)
	require.NoError(t, err)

	require.Equal(t, len(blobs), len(listedBlobs.Items))
	for i, blob := range blobs {
		require.Equal(t, blob.name, *listedBlobs.Items[i].Key, i)
		require.Equal(t, blob.metadata, listedBlobs.Items[i].Metadata, i)
		if len(blob.content) < 1024 {
			require.Equal(t, blob.content, listedBlobs.Items[i].Content, i)
		} else {
			require.Nil(t, listedBlobs.Items[i].Content, i)
		}
	}
}
