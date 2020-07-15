import threading


import threading

class Foo:
    def __init__(self):
        self.b1 = threading.Event()
        self.b2 = threading.Event()

    def first(self, printFirst: 'Callable[[], None]') -> None:
        printFirst()
        self.b1.set()

    def second(self, printSecond: 'Callable[[], None]') -> None:
        self.b1.wait()
        printSecond()
        self.b2.set()

    def third(self, printThird: 'Callable[[], None]') -> None:
        self.b2.wait()
        printThird()
