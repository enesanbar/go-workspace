import time
import numpy as np

# create an array with 10^7 elements
numpy_array = np.arange(1e7)

# convert ndarray to list
python_list = numpy_array.tolist()

start_time = time.time()
numpy_array *= 1.1
numpy_execution_time = time.time() - start_time

start_time = time.time()
for i, val in enumerate(python_list):
    python_list[i] = val * 1.1
list_execution_time = time.time() - start_time

print("{} s (NumPy Multiplication - arr * 1.1)".format(numpy_execution_time))
print("{} s (List Multiplication with loop)".format(list_execution_time))
print("NumPy's ndarray operation is {} times faster than Python loop"
      .format(round(list_execution_time/numpy_execution_time, 2)))
