package example

import (
	"github.com/pachyderm/pachyderm/src/pfs"
	"github.com/pachyderm/pachyderm/src/pps"
)

func CreateJobRequest() *pps.CreateJobRequest {
	return &pps.CreateJobRequest{
		Transform: &pps.Transform{
			Cmd: []string{"cmd", "args..."},
		},
		Shards: 1,
		InputCommit: []*pfs.Commit{
			{
				Repo: &pfs.Repo{Name: "in_repo"},
				Id:   "10cf676b626044f9a405235bf7660959",
			},
		},
		OutputParent: &pfs.Commit{
			Repo: &pfs.Repo{Name: "out_repo"},
			Id:   "a951ca06cfda4377b8ffaa050d1074df",
		},
	}
}

func CreatePipelineRequest() *pps.CreatePipelineRequest {
	return &pps.CreatePipelineRequest{
		Pipeline: &pps.Pipeline{
			Name: "name",
		},
		Transform: &pps.Transform{
			Cmd: []string{"cmd", "args..."},
		},
		Shards: 1,
		InputRepo: []*pfs.Repo{
			{Name: "in_repo"},
		},
		OutputRepo: &pfs.Repo{Name: "out_repo"},
	}
}
