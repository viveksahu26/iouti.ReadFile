# iouti.ReadFile
1) func ReadFile(filename string) ([]byte, error)

// ReadFile reads the file named by filename and returns the contents. A successful call returns err == nil, not err == EOF. 

// Because ReadFile reads the whole file, it does not treat an EOF from Read as an error to be reported.

2) string(content))---> converting bytes of data return by ReadFile into strings..

3) func WriteFile(filename string, data []byte, perm os.FileMode) error

//WriteFile writes data to a file named by filename. If the file does not exist, 

//WriteFile creates it with permissions perm (before umask); otherwise WriteFile truncates it before writing, without changing permissions

4) func OpenFile(name string, flag int, perm FileMode) (*File, error)

//OpenFile is the generalized open call; most users will use Open or Create instead. It opens the named file with specified flag (O_RDONLY etc.). 

//If the file does not exist, and the O_CREATE flag is passed, it is created with mode perm (before umask). 

//If successful, methods on the returned File can be used for I/O. If there is an error, it will be of type *PathError.
