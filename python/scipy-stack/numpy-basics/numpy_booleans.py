import numpy as np
import numpy.random as rand

# Filtering with conditional
arr = np.array([1, 2, 3, 4, 5, 6, 7])

# returns boolean array: [False, False, True, True, T...]
index = arr > 2

# Creating the desired array
new_arr = arr[index]
# new_arr = arr[arr > 2]
print(new_arr)

#########################################################

img1 = np.zeros((20, 20)) + 3
img1[4:-4, 4:-4] = 6
img1[7:-7, 7:-7] = 9

img2 = np.copy(img1)
# Let's filter out all values larger than 2 and less than 7.
img2[(img1 > 3) & (img1 < 7)] = 0
print(img2)

################################################################

# Operate on specific elements in an array,
a = rand.randn(100)

# Here we generate an index for filtering out undesired elements.
index = a > 0.2
b = a[index]

# We execute some operation on the desired elements.
b = b ** 2 - 2

# Then we put the modified elements back into the original array.
a[index] = b
