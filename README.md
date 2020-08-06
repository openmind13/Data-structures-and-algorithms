# Data-structures-and-algorithms
Some data structures and algorithms in GO


## Stack

```
stack := stack.New()
stack.Push(10)
stack.Pop()
stack.IsEmpty()
stack.Print()
stack.Size()
stack.Top()
```

## Queue

```
queue := queue.New()
queue.Add(10)
queue.Print()
queue.Get()
queue.IsEmpty()
queue.Size()
queue.Dequeue()
```

## Binary tree

```
tree := binarytree.New()
tree.Add(100)
tree.Print()
tree.PrefixPrint()
tree.InfixPrint()
tree.PostfixPrint()
tree.Clear()
tree.IsEmpty()

// tree.RemoveItem(10)
```

## Linked list

```
list := linkedlist.New()
list.AddHead(10)
list.AddTail(20)
list.Print()
list.Clear()
list.IsEmpty()
list.Size()

// list.AddItem(2, 500)
```
