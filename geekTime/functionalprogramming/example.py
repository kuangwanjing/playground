def example1():
    def inc(x):
        def incx(y):
            return x + y
        return incx

    inc2 = inc(2)
    inc5 = inc(5)

    print inc2(5)
    print inc5(8)

example1()

# use imperative programming to simulate a cars race.
def car_race1():
    from random import random

    time = 5
    car_positions = [1, 1, 1]

    while time:
        time -= 1

        print ''
        for i in range(len(car_positions)):
            # move car
            if random() > 0.3:
                car_positions[i] += 1

            # draw car

            print '-' * car_positions[i]

# car_race1()

time = 5
car_positions = [1, 1, 1]

# A stateful implementation of car race:
# Encapsulate the drawing and moving functionality into modules like normal encapsulation.
# This snippet of codes is constantly sharing the state of car_positions among different modules.
# This makes it hard for programers to understand the codes by switching context between modules..
def car_race2():
    from random import random

    def move_cars():
        for i in range(len(car_positions)):
            # move car
            if random() > 0.3:
                car_positions[i] += 1

    def draw_car(car_position):
        print '-' * car_position

    def run_step_of_race():
        global time
        time -= 1
        move_cars()

    def draw():
        print ''
        for car_position in car_positions:
            draw_car(car_position)

    while time:
        run_step_of_race()
        draw()

#car_race2()

# A stateless implementation of car race:
def car_race3():
    from random import random

    def move_cars(car_positions):
        return map(lambda x: x + 1 if random() > 0.3 else x, car_positions)

    def output_car(car_position):
        return '-' * car_position

    def run_step_of_race(state):
        return {'time': state['time'] - 1, 'car_positions': move_cars(state['car_positions'])}

    def draw(state):
        print ''
        print '\n'.join(map(output_car, state['car_positions']))

    def race(state):
        draw(state)
        if state['time']:
            race(run_step_of_race(state))

    race({'time': 5, 'car_positions': [1, 1, 1]})

#car_race3()

# pipeline
def pipeline1():
    def even_filter(nums):
        return filter(lambda x: x % 2 == 0, nums)

    def multiply_by_three(nums):
        return map(lambda x: x * 3, nums)

    def convert_to_string(nums):
        return map(lambda x: 'The Number: %s' % x, nums)

    nums = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

    pipeline = convert_to_string(multiply_by_three(even_filter(nums)))

    for num in pipeline:
        print num

    def pipeline_func(data, fns):
        return reduce(lambda a, x: x(a), fns, data)

    # be careful of the order of arguments of reduce.
    for num in pipeline_func(nums, [even_filter, multiply_by_three, convert_to_string]):
        print num

#pipeline1()

def pipeline2():
    class Pipe(object):
        def __init__(self, func):
            self.func = func

        def __ror__(self, other):
            def generator():
                for obj in other:
                    if obj is not None:
                        yield self.func(obj)
            return generator()
    @Pipe
    def even_filter(num):
        return num if num % 2 == 0 else None

    @Pipe
    def multiply_by_three(num):
        return num * 3

    @Pipe
    def convert_to_string(num):
        return 'The Number: %s' % num

    @Pipe
    def echo(item):
        print item
        return item

    def force(sqs):
        for item in sqs: pass

    nums = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

    force(nums | even_filter | multiply_by_three | convert_to_string | echo)

pipeline2()
