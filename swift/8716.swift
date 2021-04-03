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

func intersect(_ x1 : Int, _ x2 : Int, _ x3 : Int, _ x4 : Int) -> Int{
  var a = x1;
  var b = x2
  var c = x3
  var d = x4

  if(c < a){
    swap(&a, &c);
    swap(&b, &d);
  }

  // if(x1 <= x3){
  if( b <= c) {
    return 0
  }

  if(b <= d){
    return b - c 
  } else { 
    return d - c 
  }
}


var sc = Scanner()
var pr = [String]()

let x1 = sc.nextInt()
let y1 = sc.nextInt()
let x2 = sc.nextInt()
let y2 = sc.nextInt()

let x3 = sc.nextInt()
let y3 = sc.nextInt()
let x4 = sc.nextInt()
let y4 = sc.nextInt()


let x =  intersect(x1, x2, x3, x4)
let y =  intersect(y2, y1, y4, y3)

print( Int64(x) * Int64(y))
