# iouti.ReadFile
1) func ReadFile(filename string) ([]byte, error)

// ReadFile reads the file named by filename and returns the contents. A successful call returns err == nil, not err == EOF. 

// Because ReadFile reads the whole file, it does not treat an EOF from Read as an error to be reported.

2) string(content))---> converting bytes of data return by ReadFile into strings..

3) func WriteFile(filename string, data []byte, perm os.FileMode) error

//WriteFile writes data to a file named by filename. If the file does not exist, 

//WriteFile creates it with permissions perm (before umask); otherwise WriteFile truncates it before writing, without changing permissions
