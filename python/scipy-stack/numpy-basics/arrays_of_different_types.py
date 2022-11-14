import numpy as np

# Creating an array of zeros and defining column types
# i4: 32-bit integer, f4: 32-bit float, and a10: string of 10 characters.
recarr = np.zeros((2,), dtype='i4,f4,a10')

# Now creating the columns we want to put in the recarray
col1 = np.arange(2) + 1
col2 = np.arange(2, dtype=np.float32) + 2
col3 = ['Hello', 'World']

# put columns together:  [(1,2.,'Hello'), (2,3.,"World")]
toadd = list(zip(col1, col2, col3))
recarr[:] = toadd

# Assign names to each column: by default f0, f1, f2
recarr.dtype.names = ('Integers', 'Floats', 'Strings')

# Now we can access a column by its name
print(recarr['Integers'])
print(recarr['Floats'])
print(recarr['Strings'])
