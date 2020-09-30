//Design your implementation of the circular double-ended queue (deque).
//
// Your implementation should support following operations:
//
//
// MyCircularDeque(k): Constructor, set the size of the deque to be k.
// insertFront(): Adds an item at the front of Deque. Return true if the operati
//on is successful.
// insertLast(): Adds an item at the rear of Deque. Return true if the operation
// is successful.
// deleteFront(): Deletes an item from the front of Deque. Return true if the op
//eration is successful.
// deleteLast(): Deletes an item from the rear of Deque. Return true if the oper
//ation is successful.
// getFront(): Gets the front item from the Deque. If the deque is empty, return
// -1.
// getRear(): Gets the last item from Deque. If the deque is empty, return -1.
// isEmpty(): Checks whether Deque is empty or not.
// isFull(): Checks whether Deque is full or not.
//
//
//
//
// Example:
//
//
//MyCircularDeque circularDeque = new MycircularDeque(3); // set the size to be
//3
//circularDeque.insertLast(1);			// return true
//circularDeque.insertLast(2);			// return true
//circularDeque.insertFront(3);			// return true
//circularDeque.insertFront(4);			// return false, the queue is full
//circularDeque.getRear();  			// return 2
//circularDeque.isFull();				// return true
//circularDeque.deleteLast();			// return true
//circularDeque.insertFront(4);			// return true
//circularDeque.getFront();			// return 4
//
//
//
//
// Note:
//
//
// All values will be in the range of [0, 1000].
// The number of operations will be in the range of [1, 1000].
// Please do not use the built-in Deque library.
//
// Related Topics Design Queue
// 👍 287 👎 42

//
//641. 设计循环双端队列
//        设计实现双端队列。
//        你的实现需要支持以下操作：
//
//        MyCircularDeque(k)：构造函数,双端队列的大小为k。
//        insertFront()：将一个元素添加到双端队列头部。 如果操作成功返回 true。
//        insertLast()：将一个元素添加到双端队列尾部。如果操作成功返回 true。
//        deleteFront()：从双端队列头部删除一个元素。 如果操作成功返回 true。
//        deleteLast()：从双端队列尾部删除一个元素。如果操作成功返回 true。
//        getFront()：从双端队列头部获得一个元素。如果双端队列为空，返回 -1。
//        getRear()：获得双端队列的最后一个元素。 如果双端队列为空，返回 -1。
//        isEmpty()：检查双端队列是否为空。
//        isFull()：检查双端队列是否满了。
//        示例：
//
//        MyCircularDeque circularDeque = new MycircularDeque(3); // 设置容量大小为3
//        circularDeque.insertLast(1);			        // 返回 true
//        circularDeque.insertLast(2);			        // 返回 true
//        circularDeque.insertFront(3);			        // 返回 true
//        circularDeque.insertFront(4);			        // 已经满了，返回 false
//        circularDeque.getRear();  				// 返回 2
//        circularDeque.isFull();				        // 返回 true
//        circularDeque.deleteLast();			        // 返回 true
//        circularDeque.insertFront(4);			        // 返回 true
//        circularDeque.getFront();				// 返回 4
//
//
//
//        提示：
//
//        所有值的范围为 [1, 1000]
//        操作次数的范围为 [1, 1000]
//        请不要使用内置的双端队列库。
//

package leetcode.editor.en;

import org.junit.jupiter.api.Assertions;

public class DesignCircularDeque {
    public static void main(String[] args) {
        MyCircularDeque circularDeque = new DesignCircularDeque().new MyCircularDeque(3);

        Assertions.assertTrue(circularDeque.insertLast(1));            // return true
        Assertions.assertTrue(circularDeque.insertLast(2));            // return true
        Assertions.assertTrue(circularDeque.insertFront(3));            // return true
        Assertions.assertFalse(circularDeque.insertFront(4));            // return false, the queue is full
        Assertions.assertEquals(2, circularDeque.getRear());            // return 2
        Assertions.assertTrue(circularDeque.isFull());                // return true
        Assertions.assertTrue(circularDeque.deleteLast());            // return true
        Assertions.assertTrue(circularDeque.insertFront(4));            // return true
        Assertions.assertEquals(4, circularDeque.getFront());            // return 4
    }

    //leetcode submit region begin(Prohibit modification and deletion)
    class MyCircularDeque {
        private int[] buffers;
        private int front;
        private int rear;
        private int capacity;

        /**
         * Initialize your data structure here. Set the size of the deque to be k.
         */
        public MyCircularDeque(int k) {
            capacity = k + 1;
            buffers = new int[capacity];
        }

        /**
         * Adds an item at the front of Deque. Return true if the operation is successful.
         */
        public boolean insertFront(int value) {
            if (isFull()) {
                return false;
            } else {
                front = (front - 1 + capacity) % capacity;
                buffers[front] = value;
                return true;
            }
        }

        /**
         * Adds an item at the rear of Deque. Return true if the operation is successful.
         */
        public boolean insertLast(int value) {
            if (isFull()) {
                return false;
            } else {
                buffers[rear] = value;
                rear = (rear + 1) % capacity;
                return true;
            }
        }

        /**
         * Deletes an item from the front of Deque. Return true if the operation is successful.
         */
        public boolean deleteFront() {
            if (isEmpty()) {
                return false;
            }
            front = (front + 1) % capacity;
            return true;
        }

        /**
         * Deletes an item from the rear of Deque. Return true if the operation is successful.
         */
        public boolean deleteLast() {
            if (isEmpty()) {
                return false;
            }
            rear = (rear + capacity - 1) % capacity;
            return true;
        }

        /**
         * Get the front item from the deque.
         */
        public int getFront() {
            if (isEmpty()) {
                return -1;
            }
            return buffers[front];
        }

        /**
         * Get the last item from the deque.
         */
        public int getRear() {
            if (isEmpty()) {
                return -1;
            }
            return buffers[(rear -1 + capacity) % capacity];
        }

        /**
         * Checks whether the circular deque is empty or not.
         */
        public boolean isEmpty() {
            return front == rear;
        }

        /**
         * Checks whether the circular deque is full or not.
         */
        public boolean isFull() {
            return (rear + 1) % capacity == front;
        }
    }

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * MyCircularDeque obj = new MyCircularDeque(k);
 * boolean param_1 = obj.insertFront(value);
 * boolean param_2 = obj.insertLast(value);
 * boolean param_3 = obj.deleteFront();
 * boolean param_4 = obj.deleteLast();
 * int param_5 = obj.getFront();
 * int param_6 = obj.getRear();
 * boolean param_7 = obj.isEmpty();
 * boolean param_8 = obj.isFull();
 */
//leetcode submit region end(Prohibit modification and deletion)

}
