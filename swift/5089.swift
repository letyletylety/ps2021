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

var week = 0
while true {
    week += 1
    let n = sc.nextInt()
    if n == 0 {
        break
    }

    var visited = [String: Int]()

    for _ in 0 ..< n {
        let a = readLine()!
        visited[a, default: 0] += 1
    }

    print("Week \(week) \(visited.count)")
}
