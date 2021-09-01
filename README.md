> ### ⚠️ This is a proof of concept
>
> As this is a proof of concept,  it won't be supported by the k6 team.
> It may also break in the future as xk6 evolves. USE AT YOUR OWN RISK!
> Any issues with the tool should be raised [here](https://github.com/SkynSoul/xk6-tcp/issues).

</br>
</br>

<div align="center">

# xk6-tcp
TCP xk6 plugin. Built for [k6](https://go.k6.io/k6) using [xk6](https://github.com/k6io/xk6).

</div>

## Build

To build a `k6` binary with this extension, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git

Then:

1. Download `xk6`:
  ```bash
  $ go install github.com/k6io/xk6/cmd/xk6@v0.33.0
  ```

2. Build the binary:
  ```bash
  $ xk6 build --with github.com/SkynSoul/xk6-tcp@latest
  ```

3. If you're working with a local directory run the following:
  ```bash
  $ xk6 build --with github.com/SkynSoul/xk6-tcp@latest="/absolute/path/to/xk6-redis"
  ```