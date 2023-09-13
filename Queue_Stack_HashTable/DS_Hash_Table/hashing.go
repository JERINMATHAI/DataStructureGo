package main

import "fmt"

const ArraySize = 7

type HashTable struct {
	array [ArraySize]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key   string
	value string
	next  *bucketNode
}

func (h *HashTable) Insert(key string, value string) {

	index := hash(key)
	if !h.array[index].search(key) {
		h.array[index].insert(key, value)
	} else {
		fmt.Println("The key is alreay exist...")
	}

}

func (h *HashTable) Search(key string, value string) bool {
	index := hash(key)
	return h.array[index].search(value)
}

func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)

}

func (b *bucket) insert(k string, v string) {
	if !b.search(k) {
		newNode := new(bucketNode)
		newNode.key = k
		newNode.value = v

		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, "Already exist.....!")
	}

}

func (b *bucket) search(k string) bool {
	temp := b.head

	for temp != nil {
		if temp.key == k {
			return true
		}

		temp = temp.next
	}

	return false
}

func (b *bucket) delete(k string) {

	if b.head.key == k {
		b.head = b.head.next
		return
	}

	temp := b.head

	for temp.next != nil {
		if temp.next.key == k {
			temp.next = temp.next.next
			return
		}

		temp = temp.next
	}

	fmt.Println(k, "Doesn`t exist....!")

}

func hash(key string) int {
	sum := 0

	for _, v := range key {
		sum += int(v)
	}

	return sum % ArraySize
}

func Init() *HashTable {
	result := &HashTable{}

	for i := range result.array {
		result.array[i] = &bucket{}
	}

	return result
}

func main() {

	hashTable := Init()

	hashTable.Insert("name", "Jerin")

	fmt.Println(hashTable.Search("name", "Sanika"))

	for i := 0; i < len(hashTable.array); i++ {
		fmt.Println(hashTable.array[i])
		temp := hashTable.array[i].head

		fmt.Println(i)

		for temp != nil {
			fmt.Println(temp.key)
			fmt.Println(temp.value)
			temp = temp.next

		}

	}

}
