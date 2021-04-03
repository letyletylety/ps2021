struct Scanner {
    private var elements = [String]()
    private var index = 0

    mutating func peek() -> String {
        while elements.count == index {
            elements = readLine()!.split(separator: " ").map(String.init)
            index = 0
        }
        return elements[index]
    }

    mutating func next() -> String {
        defer { index += 1 }
        return peek()
    }

    mutating func nextInt() -> Int {
        return Int(next())!
    }

    mutating func nextInts(_ n: Int) -> [Int] {
        return (0 ..< n).map { _ in nextInt() }
    }

    mutating func nextDouble() -> Double {
        return Double(next())!
    }
}

var sc = Scanner()
var pr = [String]()

let x = sc.nextInt()

for _ in 0 ..< x {
  var y = Int64(sc.nextInt())
  let w = sc.nextInt()

  var answer = 0
  for cnt in 0 ..< 1000000 {
    if y >= w {
      answer = cnt
      break
    }

    y *= 2
  }

  print(answer)
}
