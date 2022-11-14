from scipy.optimize import fsolve
import numpy as np
import matplotlib.pyplot as plt


#####################################################
# Approximate the root of a linear function at y = 0.
#####################################################


def func(x):
    return x + 3


# Generating 100 clean data between 0-10
x = np.linspace(-10, 10, 100)
y = func(x)

solution = fsolve(func, -2)

plt.figure()
plt.plot(x, y, label="Function")
plt.plot(solution, func(solution), label="Root", marker='o')
plt.legend()
plt.show()

print(solution)


########################################################
# Finding the intersection points between two equations
########################################################


# Defining function to simplify intersection solution
def find_intersection(func1, func2, x0):
    return fsolve(lambda x: func1(x) - func2(x), x0)

# Defining functions that will intersect
funky = lambda x: np.cos(x / 5) * np.sin(x / 2)
line = lambda x: 0.01 * x - 0.5

# Defining range and getting solutions on intersection points
x = np.linspace(0, 45, 10000)
result = find_intersection(funky, line, [15, 20, 30, 35, 40, 45])

plt.figure()
plt.plot(x, funky(x), label="Cosine")
plt.plot(x, line(x), label="Line")
plt.plot(result, line(result), 'ro', label="Intersections")
plt.legend()
plt.show()


# Printing out results for x and y
print(result, line(result))
