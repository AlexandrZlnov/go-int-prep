# Go Interview Preparation Repository

## Введение
Этот репозиторий содержит набор задач и решений на языке **Go**, предназначенных для подготовки к собеседованиям на позицию Go‑разработчика. Задачи сгруппированы по темам, которые часто встречаются в интервью: работа с массивами, каналами, горутинами, синхронизацией, базами данных и т.д. В каждой задаче есть небольшое описание, ссылка на папку с кодом и, при необходимости, примеры тестов.

### Что делает репозиторий
* **Содержит готовые решения** – можно сразу запускать и проверять, как работает код.
* **Покрывает широкий спектр тем** – от простых структур данных до сложных паттернов конкурентного программирования.
* **Поддерживает поиск по работодателям** – в некоторых задачах указаны комментарии вида `// собес: Авито`, что позволяет быстро собрать все задания, которые встречались в интервью конкретной компании.
* **Лёгко расширяется** – добавляйте новые папки и задачи, и они автоматически появятся в списке.

## Структура репозитория
```
├── patterns/          – примеры паттернов проектирования
├── tasks/             – задачи по темам
│   ├── arr/           – массивы
│   ├── chan/          – каналы
│   ├── concurrency/   – конкурентность
│   ├── db/            – работа с БД
│   ├── gr/            – горутины
│   ├── interfaces/    – интерфейсы
│   ├── map/           – карты
│   ├── other/         – прочие задачи
│   │   ├── courses/   – курсы
│   │   └── t-bank/    – задачи T‑Bank
│   │   ├── balancers/ – балансировщики
│   │   └── cycles/    – циклы
│   ├── pointers/      – указатели
│   ├── slice/         – срезы
│   ├── strings/       – строки
│   ├── structs/       – структуры
│   ├── unresolved/    – нерешённые задачи
│   └── wip/           – what‑it‑print (что выведет)
└── README.md
```

## Задачи по темам

<details>
<summary>Arr (массивы)</summary>

* [task1](tasks/arr/task1/): получить из массива k наиболее часто встречающихся элементов
* [task2](tasks/arr/task2/): слить 2 массива в один отсортированный (по возрастанию)
* …
</details>

<details>
<summary>Chan (каналы)</summary>

* [task1](tasks/chan/!task1/): <span style="color:red">ПУСТАЯ</span>
* [task2](tasks/chan/task2/): дернуть N урлов с лимитом K
* [task3](tasks/chan/task3/): pipeline — writer генерирует 1..10, double умножает значения на 2, reader выводит результат
* [task4](tasks/chan/task4/): демонстрация select с таймерами и context (выбор источника завершения)
* [task5](tasks/chan/task5/): предотвращение утечек горутин через context и корректное прерывание отправки в канал
* [task6](tasks/chan/task6/): обёртка для долгой функции — прервать выполнение и вернуть ошибку, если время > 3s
* [task7](tasks/chan/task7/): merge n каналов; при закрытии любого входного канала отменить/закрыть весь пайплайн
* [task8](tasks/chan/task8/): объединить (fan‑in) несколько входных каналов в один выходной канал
* [task9](tasks/chan/task9/): дождаться закрытия двух каналов (варианты: WaitGroup и select)
* [task10](tasks/chan/task10/): запустить N воркеров (пример — 5), обрабатывать processData с общим таймаутом 5s; при тайм‑ауте отдавать 0
* …
</details>

<details>
<summary>Concurrency (конкурентность)</summary>

* [task1](tasks/concurrency/task1/): // собес: авито – пример синхронизации
* [task2](tasks/concurrency/task2!/): …
* [task3](tasks/concurrency/task3/): …
* [task4](tasks/concurrency/task4/): …
* …
</details>

<details>
<summary>Db (база данных)</summary>

* [task1](tasks/db/sql/task1/): запрос к БД
* [task2](tasks/db/sql/task2/): …
* [task3](tasks/db/sql/task3/): …
* [task4](tasks/db/sql/task4/): …
* [task5](tasks/db/sql/task5/): …
* …
</details>

<details>
<summary>Goroutines (горутиныr)</summary>

* [task1](tasks/gr/task1/): …
* [task2](tasks/gr/task2/): …
* [task3](tasks/gr/task3/): …
* …
</details>

<details>
<summary>Interfaces (интерфейсы)</summary>

* [task1](tasks/interfaces/task1/): …
* …
</details>

<details>
<summary>Maps (мапы)</summary>

* [task1](tasks/map/task1/): …
* …
</details>

<details>
<summary>Other (прочие)</summary>

  <details>
  <summary>Сourses (задачи из курсов)</summary>

  * [task1](tasks/other/courses/task1/): …
  * [task2](tasks/other/courses/task2/): …
  * [task3](tasks/other/courses/task3/): …
  * [task4](tasks/other/courses/task4/): …
  * [task5](tasks/other/courses/task5/): …
  * [task6](tasks/other/courses/task6/): …
  * [task7](tasks/other/courses/task7/): …
  * [tast8](tasks/other/courses/tast8/): …
  </details>

  <details>
  <summary>T-bank (задачи T‑Bank)</summary>

  * [task1](tasks/other/t-bank/task1/): …
  * [task2](tasks/other/t-bank/task2/): …
  * [task3](tasks/other/t-bank/task3/): …
  * [task4](tasks/other/t-bank/task4/): …
  * [task5](tasks/other/t-bank/task5!/): …
  * [task6](tasks/other/t-bank/task6!/): …
  * [task7](tasks/other/t-bank/task7!/): …
  </details>  

  <details>
  <summary>Balancers (балансировщики)</summary>

  * [task1](tasks/unresolved/balancers/round‑robin/main.go): …
  </details>  

  <details>
  <summary>Cycles (циклы)</summary>

  * [task1](tasks/unresolved/cycles/task1/): …
  </details>  

* [task1](tasks/other/task1/): …
* [task2](tasks/other/task2/): …
* [task3](tasks/other/task3/): …
* [task4](tasks/other/task4/): …
* [task5](tasks/other/task5/): …
* [task6](tasks/other/task6/): …
* [task7](tasks/other/task7/): …
* [task8](tasks/other/task8/): …
</details>

<details>
<summary>Pointers (указатели)</summary>

* [task1](tasks/pointers/task1/): …
* …
</details>

<details>
<summary>Slices (срезы)</summary>

* [task1](tasks/slice/task1/): …
* [task2](tasks/slice/task2/): …
* [task3](tasks/slice/task3/): …
* …
</details>

<details>
<summary>Strings (строки)</summary>

* [task1](tasks/strings/task1/): …
* [task2](tasks/strings/task2/): …
* …
</details>

<details>
<summary>Structs (структуры)</summary>

* [task1](tasks/structs/task1/): …
* …
</details>

<details>
<summary>Unresolved (нерешённые задачи)</summary>

* [db](tasks/unresolved/db/!task1/): …
* [task1](tasks/unresolved/db/!task1/main.go): …
* …
</details>

<details>
<summary>What it print (что выведет)</summary>

* [task1](tasks/wip/task1/): …
* [task2](tasks/wip/task2/): …
* [task3](tasks/wip/task3/): …
* [task4](tasks/wip/task4/): …
* [task5](tasks/wip/task5/): …
* [task6](tasks/wip/task6/): …
* [task7](tasks/wip/task7/): …
* [task8](tasks/wip/task8/): …
* [task9](tasks/wip/task9/): …
* [task10](tasks/wip/task10/): …
* [task11](tasks/wip/task11/): …
* [task12](tasks/wip/task12/): …
* [task13](tasks/wip/task13/): …
* [task14](tasks/wip/task14/): …
* [task15](tasks/wip/task15/): …
* [task16](tasks/wip/task16/): …
* [task17](tasks/wip/task17/): …
* [task18](tasks/wip/task18/): …
* [task19](tasks/wip/task19/): …
* [task20](tasks/wip/task20/): …
* [task21](tasks/wip/task21/): …
* [task22](tasks/wip/task22/): …
</details>

## Задачи по работодателям

<details>
<summary>Авито</summary>

* [task1](tasks/concurrency/task1/): // собес: авито
* [task7](tasks/other/task7/): // собес: Авито Платформа
* [task2](tasks/gr/task2/): …
* [task3](tasks/gr/task3/): …
* …
</details>

<details>
<summary>МТС</summary>

* …
</details>

<details>
<summary>X5</summary>

* [task2](tasks/concurrency/task2!/): …
</details>

<details>
<summary>Магнит</summary>

* [task4](tasks/concurrency/task4/): … 
</details>

<details>
<summary>Яндекс</summary>

* [task1](tasks/map/task1/): …
</details>

## Быстрый старт
Клонируй репозиторий:  
  ```git clone https://github.com/AlexandrZlnov/go-int-prep.git```  
  ```cd go-int-prep```  

Запуск решения:  
  ```go run tasks/arr/task1/main.go```

Или выполни `go run main.go` в соответствующей задаче.


## Как внести вклад
Контрибьюции приветствуются!

**Как добавить изменения**
1. Сделай fork репозитория
2. Создай новую ветку:  
    ```git checkout -b feature/add-task```  
3. Создай папку в `tasks/<тема>/<taskN>`  
4. Добавь:
 - `main.go` с решением задачи. 
 - `README.md` обнови основной README (добавь задачу в список).
4. Сделай commit:  
  ```git commit -m "add new task"```
5. Запушь изменения:  
  ```git push origin feature/add-task```
6. Открой Pull Request

### Правила контрибьюции
**Структура задачи**
- Каждая задача должна быть в отдельной папке
- Формат названия: taskN
- Обязательно наличие:
  - `main.go` - задача с решением 
  - `README.md` - с описанием твоей задачи
**Нейминг**
- Используй task1, task2 и т.д.
- Не используй спецсимволы
**Что можно добавлять**
- Новые задачи
- Улучшения решений
- Исправления и документацию
**Что не принимается**
- Дубликаты задач
- Задачи без описания
- Нерабочий код

### Стек технологий
Go (Golang)

## Поддержка
Если репозиторий оказался полезным поддержи меня — поставь звезду ⭐

---
**Автор:** Alex Zel
