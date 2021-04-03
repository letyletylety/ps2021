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

for scenario in 1 ..< 123_123_123 {
    let n = sc.nextInt()
    let d = sc.nextInt()

    if (n, d) == (0, 0) {
        break
    }

    print("Scenario \(scenario)")

    let br = sc.nextInt()
    let cr = sc.nextInt()

    for days in 1 ... d {
        var a = sc.nextInt()
        var b = sc.nextInt()

        if br > 0, br <= a {
            a += 1
        }

        if cr > 0, cr <= b {
            b -= 1
        }

        // print(a, b)
        var result = ""
        if a + b == n + 1 {
            result = "ALERT"
        } else {
            result = "OK"
        }
        print("Day \(days) \(result)")
    }
}
