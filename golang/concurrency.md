# Concurrency.md

### concurrency. intro
* https://medium.com/nuances-of-programming/%D0%BA%D0%BE%D0%BD%D0%BA%D1%83%D1%80%D0%B5%D0%BD%D1%82%D0%BD%D0%BE%D1%81%D1%82%D1%8C-%D0%B8-%D0%BF%D0%B0%D1%80%D0%B0%D0%BB%D0%BB%D0%B5%D0%BB%D0%B8%D0%B7%D0%BC-%D0%B2-golang-go-%D0%BF%D1%80%D0%BE%D1%86%D0%B5%D0%B4%D1%83%D1%80%D1%8B-82bae0f92e81 - конкуренция в го
* https://habr.com/ru/sandbox/152886/ - отличия конкуренции и параллелизма

### goroutine
* https://habr.com/ru/companies/otus/articles/527748/ - что такое горутина
* https://habr.com/ru/articles/141853/ - внетреннее устройство горутины
* https://dave.cheney.net/2013/06/02/why-is-a-goroutines-stack-infinite - просто статья про горутины

### race condition

Состояние гонки возникает тогда, когда несколько потоков многопоточного приложения пытаются одновременно получить доступ к данным, причем хотя бы один поток выполняет запись.
Для обнаружение race condition можно использовать флаг -race.


* https://go.dev/blog/race-detector - (не для новичков)
* https://medium.com/@deckarep/the-go-1-19-atomic-wrappers-and-why-to-use-them-ae14c1177ad8 - (uber.atomic)
* https://medium.com/trendyol-tech/race-conditions-in-golang-511314c0b85 - (основы синхронизации)
* https://medium.com/@yardenlaif/go-sync-or-go-home-errgroup-f91a0ee72d3f - errgroup

### pattern concurrency
* https://github.com/golang/go/wiki/CommonMistakes - очень крутая статься. Распространенные ошибки
* https://go.dev/blog/pipelines - (не для новичков, для общего) piplines pattern
* https://medium.com/statuscode/pipeline-patterns-in-go-a37bb3a7e61d - (дополнение к статье выше) piplines pattern
* https://blog.golang.org/go-concurrency-patterns-timing-out-and - timeout pattern
* http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang - паттерн воркер