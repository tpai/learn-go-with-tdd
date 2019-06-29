Difference Between Concurrency
===

There are two ways to handle goroutine job: `mutexes` and `channels`, they can achieve the same result just like the sample code in this folder, according to [this amazing book for learning Go](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/sync#wrapping-up), there are two recommendation scenarios belongs to both ways individually.

- `mutexes`: Managing state, such as counter or wallet.
- `channels`: Passing data, such as API response or other async behaviours.
