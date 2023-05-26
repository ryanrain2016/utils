# 封装golang的一些容器，提供类似python中对应结构的方法
# 写了一些golang版的python的一些内置方法
# 作者比较懒为了不写文档，实现的效果几乎和python的一致


# 增加了通用avltree的实现
这里使用类型参数，不必使用any传参，减少类型转换代码
```golang
import "github.com/ryanrain2016/utils/collections/tree"
type user struct {
    ID   int
    Name string
}
tree := tree.NewAVLTreeByKey(func(u *user) int { return u.ID })
tree.Insert(&user{5, "alice"})
tree.Insert(&user{2, "bob"})
tree.Insert(&user{7, "tom"})
tree.Insert(&user{1, "sam"})
tree.Insert(&user{8, "grace"})
tree.Insert(&user{3, "lily"})
tree.Insert(&user{6, "jim"})
tree.Insert(&user{4, "candy"})
tree.Insert(&user{5, "bob"})
users := make([]*user, 0)
tree.Traverse(func(u *user) { users = append(users, u) })
u, err := tree.Search(&user{ID: 4})
u // &user{4, "candy"}
```
也可以传入cmp方法
```golang
import "github.com/ryanrain2016/utils/collections/tree"

tree := tree.NewAVLTree(func(a, b int) int {
    if a < b {
        return -1
    } else if a > b {
        return 1
    } else {
        return 0
    }
})
tree.Insert(5)
tree.Insert(2)
tree.Insert(7)
tree.Insert(1)
tree.Insert(8)
tree.Insert(3)
tree.Insert(6)
tree.Insert(4)
val, err := tree.Search(6)
val // 6
```