# Concurrency.md

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
* http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang - паттерн воркер\