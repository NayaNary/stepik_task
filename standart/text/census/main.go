package main

import (
	"fmt"
	mathrand "math/rand"
	"os"
	"path/filepath"
)

// алфавит планеты Нибиру
const alphabet = "aeiourtnsl"

// Census реализует перепись населения.
// Записи о рептилоидах хранятся в каталоге census, в отдельных файлах,
// по одному файлу на каждую букву алфавита.
// В каждом файле перечислены рептилоиды, чьи имена начинаются
// на соответствующую букву, по одному рептилоиду на строку.
type Census struct {
	mapFiles map[byte]*os.File
	count    int
}

// Count возвращает общее количество переписанных рептилоидов.
func (c *Census) Count() int {
	return c.count
}

// Add записывает сведения о рептилоиде.
func (c *Census) Add(name string) {
	if len(name) == 0 {
		return
	}
	key := name[0]
	file, ok := c.mapFiles[key]
	var err error
	if !ok {
		file, err = os.OpenFile(filepath.Join("census", fmt.Sprintf("%s.txt", string(key))), os.O_CREATE|os.O_APPEND, 0666)
		fmt.Println(err)
	}
	_, _ = file.WriteString(name)
	// fmt.Println(n, err)
	_, _ = file.WriteString("\n")
	// fmt.Println(n, err)
	c.mapFiles[key] = file
	c.count++
}

// Close закрывает файлы, использованные переписью.
func (c *Census) Close() {
	for _, file := range c.mapFiles {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}
}

// NewCensus создает новую перепись и пустые файлы
// для будущих записей о населении.
func NewCensus() *Census {
	return &Census{
		mapFiles: make(map[byte]*os.File, len(alphabet)),
	}
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

var rand = mathrand.New(mathrand.NewSource(0))

// randomName возвращает имя очередного рептилоида.
func randomName(n int) string {
	chars := make([]byte, n)
	for i := range chars {
		chars[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(chars)
}

func main() {
	census := NewCensus()
	defer census.Close()
	for i := 0; i < 1024; i++ {
		reptoid := randomName(5)
		census.Add(reptoid)
	}
	fmt.Println(census.Count())
}
