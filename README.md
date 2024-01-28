# Protobuf demo

### 字段基本类型
- int 族、uint 族和 sint 族。PS：sint 族是纯纯垃圾设计， 它的初衷是大概率是负数的，就用 sint。
- bytes：字节数组（在 Go 里面是切片）。
- float 和 double：典型的浮点数。
- fixed 族和 sfixed 族：固定长度的数字类型。
- bool：bool 类型。
- string： 类型。
- map 基本语法：map<KeyType, ValueType> 。PS：map 的 key 必须是基本类型，而且必须是唯一的。
- enum：枚举类型。
- repeated：重复的，即数组。
- oneof：一组备选字段，最多只能有一个生效。
- any：任意类型，可以用来做扩展。
- oneof 和 any 是 Protobuf 3.0 新增的两个特性。