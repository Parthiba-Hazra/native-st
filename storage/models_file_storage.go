package storage

import (
	"io/ioutil"
	"mime/multipart"

	"github.com/ceph/go-ceph/rados"
)

type CephStorage struct {
	conn  *rados.Conn
	ioctx *rados.IOContext
}

func NewCephStorage() (*CephStorage, error) {

	conn, err := rados.NewConn()
	if err != nil {
		return nil, err
	}

	err = conn.ReadConfigFile("/etc/ceph/ceph.conf")
	if err != nil {
		return nil, err
	}
	err = conn.Connect()
	if err != nil {
		return nil, err
	}

	ioctx, err := conn.OpenIOContext("files")
	if err != nil {
		return nil, err
	}

	return &CephStorage{conn: conn, ioctx: ioctx}, nil
}

func (s *CephStorage) Close() error {

	s.ioctx.Destroy()
	s.conn.Shutdown()

	return nil
}

func (s *CephStorage) UploadFile(file multipart.File, filename string) error {

	// Read contents from file
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	// Create a ceph object
	obj, err := s.OpenIOContext("native-st")
	if err != nil {
		return err
	}
	defer obj.Destroy()
	obj.WriteFull(filename, fileContents)

	return nil
}

func DeleteFile(filename string) error {
	// Create a Ceph connection
	conn, err := rados.NewConn()
	if err != nil {
		return err
	}

	err = conn.ReadConfigFile("/etc/ceph/ceph.conf")
	if err != nil {
		return err
	}
	err = conn.Connect()
	if err != nil {
		return err
	}

	obj, err := conn.OpenIOContext("native-st")
	if err != nil {
		return err
	}
	defer obj.Destroy()
	obj.Delete(filename)

	return nil
}

func DownloadFile(filename string) ([]byte, error) {
	conn, err := rados.NewConn()
	if err != nil {
		return nil, err
	}

	err = conn.ReadConfigFile("/etc/ceph/ceph.conf")
	if err != nil {
		return nil, err
	}
	err = conn.Connect()
	if err != nil {
		return nil, err
	}

	ioctx, err := conn.OpenIOContext("native-st")
	if err != nil {
		return nil, err
	}
	defer ioctx.Destroy()

	fileStat, err := ioctx.Stat(filename)
	if err != nil {
		return nil, err
	}

	fileSize := fileStat.Size
	fileCounts := make([]byte, fileSize)
	_, err = ioctx.Read(filename, fileCounts, 0)
	if err != nil {
		return nil, err
	}

	return fileCounts, nil
}
