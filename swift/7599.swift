struct Scanner {
    var elements = [String]()
    var index = 0

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

for _ in 0 ..< 987_654_321 {
    let lib = sc.next()
    let f = sc.nextInt()
 
    if f == 0 {
        break
    }

    let n = sc.nextInt()

    print("\(lib) Library")

    for i in 1 ... n {
        let c = sc.nextInt()
        let w = sc.next()

        var answer = ""

        if 2 + f * w.count <= c {
            answer = "horizontal"
        } else {
            answer = "vertical"
        }

        print("Book \(i) \(answer)")
    }
}
