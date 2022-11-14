import numpy as np
from PIL import Image

# numpy array from python list
alist = list(range(10))
arr = np.array(alist)

# numpy array in range 0-N
arr = np.arange(100)

# numpy array in range M-N
arr = np.arange(10, 100)

# numpy array with steps. 100 steps between 0 and 1
arr = np.linspace(0, 1, 100)

# Creating a 50x50 array of zeros (an image)
data = np.zeros((500, 500))
img = Image.fromarray(data, 'RGB')
# img.show()


# Creating a 5x5x5 cube of 1's
# The astype() method sets the array with integer elements.
cube = np.zeros((5, 5, 5)).astype(int) + 1

# Or even simpler with 16-bit floating-point precision...
cube = np.ones((5, 5, 5)).astype(np.float16)

# when creating arrays, 64-bit precision is not always necessary
# we can specify the precision with dtype
# Array of zero integers
arr = np.zeros(2, dtype=int)
# Array of zero floats
arr = np.zeros(2, dtype=np.float32)


# RESHAPING ARRAYS

# Creating an array with elements from 0 to 999
arr1d = np.arange(100)

# Now reshaping the array to a 10x10x10 3D array
# original size must not change: 10 * 5 * 2 = 100
arr3d = arr1d.reshape((10, 5, 2))
# The reshape command can alternatively be called this way
# arr3d = np.reshape(arr1d, (10, 10, 10))

# Inversely, we can flatten arrays
arr4d = np.zeros((10, 10, 10, 10))
arr1d = arr4d.ravel()

# slicing arrays
alist=[['00', '01', '02'],
       ['10', '11', '12'],
       ['20', '21', '22']]

# Converting the list defined above into an array
arr = np.array(alist)

# To return the (0,1) element we use ...
print(arr[0, 1])

# Now to access the second column, we simply use ...
print('2nd column: ', arr[:, 1])

# Accessing the rows is achieved in the same way,
print('2nd row:    ', arr[1, :])