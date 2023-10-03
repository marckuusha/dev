# Concurrency.md

### race condition

Состояние гонки возникает тогда, когда несколько потоков многопоточного приложения пытаются одновременно получить доступ к данным, причем хотя бы один поток выполняет запись.
Для обнаружение race condition можно использовать флаг -race.

* https://go.dev/blog/race-detector - (не для новичков)
* https://medium.com/@deckarep/the-go-1-19-atomic-wrappers-and-why-to-use-them-ae14c1177ad8 - (uber.atomic)
* https://medium.com/trendyol-tech/race-conditions-in-golang-511314c0b85 - (основы синхронизации)
* https://medium.com/@yardenlaif/go-sync-or-go-home-errgroup-f91a0ee72d3f - errgroup

### piplines
* https://go.dev/blog/pipelines - (не для новичков, для общего)
* https://medium.com/statuscode/pipeline-patterns-in-go-a37bb3a7e61d - (дополнение к статье выше)