package zllock

import (
	"errors"
	"fmt"
	"path"
	"sort"
	"strings"
	"time"

	"github.com/go-zookeeper/zk"
)

type ZKLock struct {
	rootPath string
	pathName string
	conn     *zk.Conn
}

// func main() {
// 	zl := new(ZKLock)
// 	err := zl.initConn("/lockTest13", []string{"127.0.0.1:2181"})
// 	zl.AttempLock()
// 	time.Sleep(time.Second * 10)
// 	zl.conn.Close()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// }

func (zl *ZKLock) Init(rootPath string, servers []string) error {
	zl.rootPath = rootPath
	var err error
	zl.conn, _, err = zk.Connect(servers, time.Second*10)
	if err != nil {
		panic(err)
	}
	fmt.Println("check if root dir exist: " + rootPath)
	isExist, _, err := zl.conn.Exists(zl.rootPath)
	if !isExist {
		fmt.Println("Dir not exist: " + rootPath)
		if _, err := zl.conn.Create(zl.rootPath, nil, 0, zk.WorldACL(zk.PermAll)); err != nil {
			return err
		}
	}
	return nil
}

func (zl *ZKLock) CreateLock(rootPath string, servers []string) error {
	zl.rootPath = rootPath
	var err error
	zl.conn, _, err = zk.Connect(servers, time.Second*10)
	if err != nil {
		panic(err)
	}

	zl.pathName, err = zl.conn.Create(zl.rootPath+"/", nil, zk.FlagEphemeral|zk.FlagSequence, zk.WorldACL(zk.PermAll))
	if err != nil {
		return err
	}
	fmt.Println("CreateLock: " + zl.pathName)
	return nil
}

func (zl *ZKLock) AttempLock() error {
	lockPaths, _, err := zl.conn.Children(zl.rootPath)
	if err != nil {
		return err
	}
	sort.Strings(lockPaths)
	fmt.Println("Child of rootPath: " + strings.Join(lockPaths, " "))
	lowestPath := lockPaths[0]
	if path.Join(zl.rootPath, lowestPath) == zl.pathName {
		fmt.Printf("%s acquire lock success\n", zl.pathName)
		return nil
	}
	wtachPath, err := zl.getWatchPath(lockPaths)
	if err != nil {
		return err
	}

	return zl.waitLock(wtachPath)
}

func (zl *ZKLock) getWatchPath(lockPaths []string) (string, error) {
	for index, v := range lockPaths {
		if zl.rootPath+"/"+v == zl.pathName {
			return zl.rootPath + "/" + lockPaths[index-1], nil
		}
	}
	return "", errors.New(fmt.Sprintf("NotFound path: %s\n", zl.pathName))
}

func (zl *ZKLock) waitLock(path string) error {
	fmt.Println("start monitoring path: " + path)
	isExist, _, nodeEvent, err := zl.conn.ExistsW(path)
	if err != nil {
		return err
	}
	if !isExist {
		fmt.Println("Monitoring node does not exist, acquire lock")
		return nil
	}
	// should have retry times
	for {
		select {
		case event := <-nodeEvent:
			{
				if event.Type == zk.EventNodeDeleted {
					fmt.Printf("%s Monitor node goes offline, get lock\n", zl.pathName)
					return nil
				}
			}
		}
	}
}

func (zl *ZKLock) Unlock() error {
	fmt.Println("Release lock: " + zl.pathName)
	_, s, err := zl.conn.Get(zl.pathName)
	if err != nil {
		return err
	}
	return zl.conn.Delete(zl.pathName, s.Version)
}
