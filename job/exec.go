package job

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"strconv"

	"github.com/kpango/glg"
)

type JobExec struct {
	Hash      []byte      `json:"hash"`
	Timestamp int64       `json:"timestamp"`
	Duration  int64       `json:"duaration"` //saved in nanoseconds
	Err       interface{} `json:"err"`
	Result    interface{} `json:"result"`
	By        []byte      `json:"by"` //! ID of the worker node that ran this
}

func (j JobExec) GetHash() []byte {
	return j.Hash
}

func (j *JobExec) setHash() {
	headers := bytes.Join(
		[][]byte{
			[]byte(strconv.FormatInt(j.GetTimestamp(), 10)),
			[]byte(strconv.FormatInt(j.GetDuration(), 10)),
			j.GetErr().([]byte),
			j.GetResult().([]byte),
			j.GetBy(),
		},
		[]byte{},
	)
	hash := sha256.Sum256(headers)
	j.Hash = hash[:]
}

func (j JobExec) GetTimestamp() int64 {
	return j.Timestamp
}

func (j *JobExec) SetTimestamp(t int64) {
	j.Timestamp = t
}

func (j JobExec) GetDuration() int64 {
	return j.Duration
}

func (j *JobExec) SetDuration(t int64) {
	j.Duration = t
}

func (j JobExec) GetErr() interface{} {
	return j.Err
}

func (j *JobExec) SetErr(e interface{}) {
	j.Err = e
}

func (j JobExec) GetResult() interface{} {
	return j.Result
}

func (j *JobExec) SetResult(r interface{}) {
	j.Result = r
}

func (j JobExec) GetBy() []byte {
	return j.By
}

func (j *JobExec) SetBy(by []byte) {
	j.By = by
}

func (j JobExec) Serialize() []byte {
	temp, err := json.Marshal(j)
	if err != nil {
		glg.Error(err)
	}
	return temp
}