# WTF

WTF IS GOLANG

## Usage

```sh
> // Usage
> wtf is [...args]
> // Example
> wtf is "Python Programming" Rust "Progressive Web Apps"
```

```sh
❯ ./wtf is rust
1.> rust : Rust is an iron oxide, a usually reddish-brown oxide formed by the reaction of iron and oxygen in the catalytic presence of water or air moisture. Rust consists of hydrous iron(III) oxides (Fe2O3·nH2O) and iron(III) oxide-hydroxide (FeO(OH), Fe(OH)3), and is typically associated with the corrosion of refined iron. Given sufficient time, any iron mass, in the presence of water and oxygen, could eventually convert entirely to rust.

```

```sh
❯ ./wtf is "rust language" golang
1.> rust language : Rust is a general-purpose programming language emphasizing performance, type safety, and concurrency. It enforces memory safety, meaning that all references point to valid memory. It does so without a traditional garbage collector; instead,  memory safety errors and data races are prevented by the "borrow checker", which tracks the object lifetime of references at compile time.

2.> golang : Go is a statically typed, compiled high-level general purpose programming language. It was designed at Google in 2009 by Robert Griesemer, Rob Pike, and Ken Thompson. It is syntactically similar to C, but also has memory safety, garbage collection, structural typing, and CSP-style concurrency.

```

```sh
❯ ./wtf IS meow
1.> meow or Meow : A meow or miaow is a cat vocalization. Meows may have diverse tones in terms of their sound, and what is heard can vary from being chattered to calls, murmurs, and whispers. Adult cats rarely meow to each other.
```

### Uses Wikipedia
### Pretty simple
### Makes a total N+1 api calls to wikipedia where N is the number of query items you provide


## Build and Run
```sh
go build && ./wtf is rust
```

## Run Locally
```sh
go run main.go is rust
```
