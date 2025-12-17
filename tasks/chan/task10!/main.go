// По видео канала Skill Issue не тему каналов
// https://www.youtube.com/watch?v=k-1OEYl7N8Q&t=130s
// Запустить 5 воркеров для обработки входных данных и процессит с processData
// И результать processData писать в канал out
// И вся операция должна выполняться не более 5 секунд
// Если более 5 секунд то дропаем и возвращаем 0
//
// func processData(val int) int {
//     time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
//     return val * 2
// }

// func main() {
//     in := make(chan int)
//     out := make(chan int)

//     go func() {
//         for i := range 100 {
//             in <- i
//         }
//         close(in)
//     }()

//     now := time.Now()
//     processParallel(in, out, 5)

//     for val := range out {
//         fmt.Println(val)
//     }
//     fmt.Println(time.Since(now))
// }
//