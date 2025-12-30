# Репозиторий: набор задач и решений на Go для подготовки к собеседованиям и практики.

В этом README задачи сгруппированы по темам (папки в каталоге `tasks`). Для каждой задачи:
- приведено краткое описание,
- указана ссылка на папку с задачей,

---

## Структура репозитория (ключевые папки): 
### tasks - содержит задачи с разных собесов, сгруппировано по темам:
- `arr` — массивы
- `chan` — каналы
- `slice` — слайсы (срезы)
- `strings` — строки
- `structs` — структуры
- `pointers` — указатели
- `gr` — горутины
- `other` — прочие задачи
- `unresolved` — нерешённые
- `wip` — что выведет код?
- `сoncurrency` — конкурентность

## tasks/arr:
- [task1](https://github.com/AlexandrZlnov/go-int-prep/tree/main/tasks/arr/task1): получить из массива k наиболее часто встречающихся элементов
- [task2](https://github.com/AlexandrZlnov/go-int-prep/tree/main/tasks/arr/task2): слить 2 массива один отсортированный (по возрастанию)

## tasks/chan:
- [task1](https://github.com/AlexandrZlnov/go-int-prep/tree/main/tasks/chan/task1): <span style="color:red">ПУСТАЯ</span>
- [task2](https://github.com/AlexandrZlnov/go-int-prep/tree/main/tasks/chan/task2): дернуть N урлов с лимитом K 
- [task3](https://github.com/AlexandrZlnov/go-int-prep/tree/main/tasks/chan/task3): pipeline — writer генерирует 1..10, double умножает значения на 2, reader выводит результат
- [task4](https://github.com/AlexandrZlnov/go-int-prep/tree/main/tasks/chan/task4): демонстрация select с таймерами и context (выбор источника завершения)
- [task5](https://github.com/AlexandrZlnov/go-int-prep/tree/main/tasks/chan/task5): предотвращение утечек горутин через context и корректное прерывание отправки в канал
- [task6](https://github.com/AlexandrZlnov/go-int-prep/tree/main/tasks/chan/task6): обёртка для долгой функции — прервать выполнение и вернуть ошибку, если время > 3s
- [task7](https://github.com/AlexandrZlnov/go-int-prep/tree/main/tasks/chan/task7): merge n каналов; при закрытии любого входного канала отменить/закрыть весь пайплайн
- [task8](https://github.com/AlexandrZlnov/go-int-prep/tree/main/tasks/chan/task8): объединить (fan‑in) несколько входных каналов в один выходной канал
- [task9](https://github.com/AlexandrZlnov/go-int-prep/tree/main/tasks/chan/task9): дождаться закрытия двух каналов (варианты: WaitGroup и select)  
- [task10](https://github.com/AlexandrZlnov/go-int-prep/tree/main/tasks/chan/task10): запустить N воркеров (пример — 5), обрабатывать processData с общим таймаутом 5s; при тайм‑ауте отдавать 0
