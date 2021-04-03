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

while (true) {
  let n = sc.nextInt()
  if(n == 0) {break}

  let a = sc.nextInts(n)

  var cnt = 0
  var pre = 0  

  var pr = [String]()
  for i in a {
    cnt+=1
    for _  in pre ..< i {
        pr.append("\(cnt)")
    }
    pre = i
  }
  print(pr.joined(separator: " "))
}
