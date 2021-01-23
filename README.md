# iouti.ReadFile
func ReadFile(filename string) ([]byte, error)
// ReadFile reads the file named by filename and returns the contents. A successful call returns err == nil, not err == EOF. 
// Because ReadFile reads the whole file, it does not treat an EOF from Read as an error to be reported.

string(content))---> converting bytes of data return by ReadFile into strings..
