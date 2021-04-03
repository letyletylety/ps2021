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

let a = sc.nextInt()
let b = sc.nextInt()

let m = max(a + b, a - b, a * b)

var aa = "\(a)"
var bb = "\(b)"
var mm = "\(m)"

if a < 0 {
    aa = "(\(a))"
}

if b < 0 {
    bb = "(\(b))"
}

if m < 0 {
    mm = "(\(m))"
}

var cnt = 0

// print(m)

if m == a + b {
    cnt += 1
} 
if m == a - b {
    cnt += 1
} 
if m == a * b{
    cnt += 1
}

if cnt > 1 {
    print("NIE")
} else if m == a + b {
    print("\(aa)+\(bb)=\(mm)")
} else if m == a - b {
    print("\(aa)-\(bb)=\(mm)")
} else {
    print("\(aa)*\(bb)=\(mm)")
}
