class MyQueue(object):

    def __init__(self):
        self.my_queue = []

    def push(self, x):
        self.my_queue.insert(0, x)

    def pop(self):
        if self.my_queue != []:
            result = self.my_queue[-1]
            del self.my_queue[-1]
            return result

    def peek(self):
        if self.my_queue != []:
            result = self.my_queue[-1]
            # def self.my_queue[-1]
            return result

    def empty(self):
        if self.my_queue == []:
            return True
        else:
            return False


myQueue = MyQueue()
myQueue.push(1)
myQueue.push(2)
print(myQueue.peek())
myQueue.pop()
print(myQueue.peek())