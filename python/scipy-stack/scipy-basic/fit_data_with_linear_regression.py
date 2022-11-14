import numpy as np
from scipy.optimize import curve_fit
import matplotlib.pyplot as plt


################################################
# Fitting noisy data with a linear equation.
################################################


# Creating a function to model and create data
def func(x, a, b):
    return a * x + b


# Generating 100 clean data between 0-10
x = np.linspace(0, 10, 100)
y = func(x, 1, 2)

# Adding noise to the data
yn = y + 0.9 * np.random.normal(size=len(x))

# Executing curve_fit on noisy data
popt, pcov = curve_fit(func, x, yn)

plt.figure()
plt.plot(x, yn, 'ro')
plt.plot(x, y, label="Function")
plt.plot(x, func(x, *popt), label="Best Fit")
plt.legend()
# plt.show()

# popt returns the best fit values for parameters of the given model (func).
print(popt)


################################################
# Fitting noisy data with a Gaussian equation.
################################################


# Creating Gaussian equation function to model and create data.
def func(x, a, b, c):
    return a * np.exp(-(x - b) ** 2 / (2 * c ** 2))


# Generating clean data
x = np.linspace(0, 10, 100)
y = func(x, 1, 5, 2)

# Adding noise to the data
yn = y + 0.2 * np.random.normal(size=len(x))

# Executing curve_fit on noisy data. popt has the best fit results.
popt, pcov = curve_fit(func, x, yn)

plt.figure()
plt.plot(x, yn, 'ro')
plt.plot(x, y, label="Function")
plt.plot(x, func(x, *popt), label="Best Fit")
plt.legend()
# plt.show()


#######################################################
#  Fitting noisy data with multiple Gaussian equations.
#######################################################


# Two-Gaussian model
def func(x, a0, b0, c0, a1, b1, c1):
    return a0 * np.exp(-(x - b0) ** 2 / (2 * c0 ** 2)) \
           + a1 * np.exp(-(x - b1) ** 2 / (2 * c1 ** 2))


# Generating clean data
x = np.linspace(0, 20, 200)
y = func(x, 1, 3, 1, -2, 15, 0.5)
# Adding noise to the data

yn = y + 0.2 * np.random.normal(size=len(x))

# Since we are fitting a more complex function,
# providing guesses for the fitting will lead to better results.
guesses = [1, 3, 1, 1, 15, 1]

# Executing curve_fit on noisy data
popt, pcov = curve_fit(func, x, yn, p0=guesses)
plt.figure()
plt.plot(x, yn, 'ro')
plt.plot(x, y, label="Function")
plt.plot(x, func(x, *popt), label="Best Fit")
plt.legend()
plt.show()