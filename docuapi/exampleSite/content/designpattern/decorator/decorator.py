def use_logging(func):
    def wrapper(*args, **kwargs):
        print("%s is running" % func.__name__)
        return wrapper(*args, **kwargs)
    return wrapper

def bar():
    print('i am bar')

bar = use_logging(bar)
bar()