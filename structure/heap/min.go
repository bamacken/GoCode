package heap

/*
Class MinHeap
    Define 'heap' as an empty list

    Function insert(value):
        Append value to heap
        SiftUp(length of heap - 1)

    Function siftUp(index):
        While index > 0 and heap[parent(index)] > heap[index]:
            Swap heap[index] with heap[parent(index)]
            index = parent(index)

    Function extractMin():
        If heap is empty:
            Return "Heap is empty"
        minElement = heap[0]
        heap[0] = heap[-1]
        heap.pop()
        SiftDown(0)
        Return minElement

    Function siftDown(index):
        While hasLeftChild(index):
            smallerChildIndex = getSmallerChildIndex(index)
            If heap[index] < heap[smallerChildIndex]:
                Break
            Swap heap[index] with heap[smallerChildIndex]
            index = smallerChildIndex

    Function getSmallerChildIndex(index):
        If hasRightChild(index) and heap[rightChild(index)] < heap[leftChild(index)]:
            Return rightChild(index)
        Return leftChild(index)

*/
