append的常见操作：
将切片 b 的元素追加到切片 a 之后：a = append(a, b...)
复制切片 a 的元素到新的切片 b 上：
b = make([]T, len(a))
copy(b, a)
删除位于索引 i 的元素：a = append(a[:i], a[i+1:]...)
切除切片 a 中从索引 i 至 j 位置的元素：a = append(a[:i], a[j:]...)
为切片 a 扩展 j 个元素长度：a = append(a, make([]T, j)...)
在索引 i 的位置插入元素 x：a = append(a[:i], append([]T{x}, a[i:]...)...)
在索引 i 的位置插入长度为 j 的新切片：a = append(a[:i], append(make([]T, j), a[i:]...)...)
在索引 i 的位置插入切片 b 的所有元素：a = append(a[:i], append(b, a[i:]...)...)
取出位于切片 a 最末尾的元素 x：x, a = a[len(a)-1], a[:len(a)-1]
将元素 x 追加到切片 a：a = append(a, x)
