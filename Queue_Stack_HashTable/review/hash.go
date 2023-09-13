package main
import "fmt"
const arraySize=7
type HashTable struct {
	array [arraySize]*bucket

}
type bucket struct{
	next *bucketNode
}
type bucketNode struct{
	key string
	value string
	next *bucketNode
}
func Init() *HashTable{
	result:=&HashTable{}
	for i range result.array{
		result.array[i]=&bucket{}
	}
}
func Hash(key string)int{
	sum:=0
	for _,v:=range key{
		sum+=int(v)
	}
	return sum%arraySize
}

func (h *HashTable) insert(key string,value string)  {
	hash:=Hash(key)
	if !h.array==search(key){
		h.array=Insert(key,value)
	}
	
}
func (h *HashTable) search(key string,value string)  {
	hash:=Hash(key)
}
func (b *Bucket) Insert(k string,v string)  {
	newNode:=new(k)
}
func main(){

}