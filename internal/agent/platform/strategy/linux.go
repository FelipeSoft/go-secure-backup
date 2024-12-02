package platform

type LinuxStrategy struct{}

func (s *LinuxStrategy) GetContentFromPath(path string) *[]byte {
	return nil
}
