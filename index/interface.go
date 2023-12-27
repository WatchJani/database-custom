package index

type CustomIndexer interface {
	Search(key string)
	Insert(key string, value []byte)
}
