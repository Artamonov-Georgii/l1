// Есть несколько проблем
// 1. Не деаллоцируется память, используемая переменной v

// 2. На justString аллоцируется столько же памяти, сколько использует v, 
// хотя в итоге большая часть этой памяти не юзается 

// Fixed:

var justString string

func someFunc() {
  v := createHugeString(1 << 10)
  // Создаем буфер с фиксированным размером
  justString = make([]byte, 100)
  copy(justString, v[:100])
  // Освобождаем память, занимаемую v
  v = ""
}

func main() {
  someFunc()
}
