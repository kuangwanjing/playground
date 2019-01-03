def example():
    def hello(fn):
        def wrapper():
            print "hello, %s" % fn.__name__
            fn()
            print "goodbye, %s" % fn.__name__
        return wrapper

    @hello
    def Hao():
        print "i am Hao Chen"

    Hao()

#example()

def html():
    def makeHtmlTag(tag, *args, **kwds):
        def real_decorator(fn):
            css_class = " class='{0}'".format(kwds["css_class"]) \
                    if "css_class" in kwds else ""
            # if wrapped is missing, the signature of fn is different from the runtime call
            # fn takes no parameter but 1 is given
            def wrapped(*args, **kwds):
                return "<" + tag + css_class + ">" + fn(*args, **kwds) + "</" + tag + ">"
            return wrapped
        return real_decorator
    # hello = makeHtmlTag(arg1, arg2) however, hello takes no parameter
    # therefore, the example takes the advantage of closure to remember the parameters of
    # makeHtmlTag and return a decorator.
    @makeHtmlTag(tag="b", css_class="bold_css")
    @makeHtmlTag(tag="i", css_class="italic_css")
    def hello():
        return "hello world"

    print hello()

#html()

def cacheExample():
    # https://docs.python.org/2/library/functools.html
    from functools import wraps

    def memorization(fn):
        cache = {}
        miss = None

        @wraps(fn)
        def wrapper(*args):
            result = cache.get(args, miss)
            if result is miss:
                result = fn(*args)
                cache[args] = result
            return result
        
        return wrapper

    @memorization
    def fib(n):
        if n < 2:
            return n
        return fib(n-1) + fib(n-2)

    print fib(10)
    print fib(8)
    print fib(12)

#cacheExample()

def router():
    class MyApp():
        def __init__(self):
            self.func_map = {}

        # remember 
        # "@func
        #  def fn():"          
        # is equals to "fn = func(fn)", therefore, register first wraps up a returning function
        # returns fn itself. In this function, func is registered as a method of MyApp. And the 
        # function's name is referred by the reigster's parameter "name".
        def register(self, name):
            def func_wrapper(func):
                self.func_map[name] = func
                return func
            return func_wrapper

        def call_method(self, name=None):
            func = self.func_map.get(name, None)
            if func is None:
                raise Exception("No function registered - " + str(name))
            return func()

    app = MyApp()
    
    @app.register('/')
    def main_page_func():
        return "This is the main page."

    @app.register('/next_page')
    def next_page_func():
        return "This is the next page."

    print app.call_method("/")
    print app.call_method("/next_page")

router()
