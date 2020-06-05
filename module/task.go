package module

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

const (
	GoroutineNum  = ConnNum
	PartSizeLimit = 200 * 1024 // 200KB
)

type Task struct {
	downloadURL string
	fileName    string
	fileSize    int64
	file        *os.File

	downloadedParts     chan *part
	downloadFailedParts chan *part
	unDownloadedParts   chan *part

	wg *sync.WaitGroup
}

type part struct {
	start, end int64

	data []byte
}

func (task *Task) DownloadFile() {

}

func (task *Task) splitFile() {
	for start, end := int64(0), int64(0); start < task.fileSize; start += PartSizeLimit {
		if task.fileSize-start >= PartSizeLimit {
			end = start + PartSizeLimit - 1
		} else {
			end = task.fileSize - 1
		}

		p := &part{
			start: start,
			end:   end,
		}
		task.unDownloadedParts <- p
	}
}

func (task *Task) concurrentDownload() {
	for i := 0; i < GoroutineNum; i++ {
		task.wg.Add(1)

		go func() {
			defer task.wg.Done()

			for p := range task.unDownloadedParts {
				if err := task.downloadPart(p); err != nil {
					// todo: 网络错误如何处理?
					fmt.Printf("download part failed: start: %d, end: %d\n",
						p.start, p.end)
					task.downloadFailedParts <- p
				} else {
					fmt.Printf("downoad part succeeded: start: %d, end: %d\n",
						p.start, p.end)
					task.downloadedParts <- p
				}
			}
		}()
	}
}

func (task *Task) downloadPart(p *part) error {
	req, err := SetHTTPReqHeaderRange(task.downloadURL, p.start, p.end)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	p.data = data
	return nil
}

func (task *Task) fillParts() {
	for p := range task.downloadedParts {
		if err := task.writePart(p); err != nil {
			fmt.Printf("fill part failed, start: %d, end: %d, err: %s\n",
				p.start, p.end, err)

			// todo: 磁盘写入失败如何处理?
			//task.unDownloadedParts <- p
			break
		}

		fmt.Printf("fill part succeeded, start: %d, end: %d\n",
			p.start, p.end)
	}
}

func (task *Task) writePart(p *part) error {
	file := task.file
	_, err := file.Seek(p.start, 0)
	if err != nil {
		return err
	}

	_, err = file.Write(p.data)
	return err
}
