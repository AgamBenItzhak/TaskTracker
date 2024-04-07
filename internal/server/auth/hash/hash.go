package hash

type HashSalt struct {
	Hash *[]byte
	Salt *[]byte
}

type HashIface interface {
	GenerateSalt() (*[]byte, error)
	GenerateHash(data *[]byte, salt *[]byte) (*HashSalt, error)
	GenerateSaltAndHash(data *[]byte) (*HashSalt, error)
	CompareHash(data *[]byte, hash HashSalt) error
}
